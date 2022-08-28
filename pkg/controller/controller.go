package controller

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	util_runtime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/xmh19936688/ingress-cert-operator/pkg/handler"
	"github.com/xmh19936688/ingress-cert-operator/pkg/utils"
)

var maxRetries = 3

type Controller struct {
	// 用来写log
	logger *logrus.Entry
	// 用来调k8s接口
	kubeClient kubernetes.Interface
	// 用来缓存需要处理的数据
	queue workqueue.RateLimitingInterface
	// 用来跟k8s交互
	informer cache.SharedIndexInformer
	// 用来处理事件
	eventHandler handler.Handler
}

func (c *Controller) SetLogger(l *logrus.Entry) *Controller {
	c.logger = l
	return c
}

func (c *Controller) SetQueue(queue workqueue.RateLimitingInterface) *Controller {
	c.queue = queue
	return c
}

func (c *Controller) SetKubeClient(cli *kubernetes.Clientset) *Controller {
	c.kubeClient = cli
	return c
}

func (c *Controller) SetHandler(handler handler.Handler) *Controller {
	c.eventHandler = handler
	return c
}

func (c *Controller) SetInformer(i cache.SharedIndexInformer) *Controller {
	c.informer = i

	c.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		// 注册创建回调
		AddFunc: addObjectFunc(c.queue),
		// 注册删除回调
		DeleteFunc: deleteObjectFunc(c.queue),
	})

	return c
}

// controller的启动入口
func (c *Controller) Run(stopCh <-chan struct{}) {
	// 处理panic
	defer util_runtime.HandleCrash()
	// 关闭队列
	defer c.queue.ShutDown()

	c.logger.Info("starting controller")

	// 启动SharedInformer
	go c.informer.Run(stopCh)

	// runWorker之前确保缓存已同步
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		util_runtime.HandleError(fmt.Errorf("time out waiting for caches to sync"))
		return
	}
	c.logger.Info("kubewatch controller synced and ready")

	// 每秒一次循环调用runWorker
	wait.Until(c.runWorker, time.Second, stopCh)
}

func (c *Controller) runWorker() {
	// 循环处理队列中的对象
	for c.processNextItem() {
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	// 返回时标记为完成，失败的逻辑中会再次添加到队列
	defer c.queue.Done(key)

	// 重试时重新添加到队列
	// 成功和失败后执行forget
	err := c.processItem(key.(string))
	if err == nil {
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		c.logger.Errorln("will retry:", err.Error())
		// 重新添加到队列中来重试
		c.queue.AddRateLimited(key)
	} else {
		c.logger.Errorln("give up:", err.Error())
		c.queue.Forget(key)
		util_runtime.HandleError(err)
	}
	return true
}

func (c *Controller) processItem(key string) error {
	c.logger.Infoln("process:", key)

	// 根据标识符获取对象
	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("failed get obj by key: %v", err)
	}

	// 不存在就执行删除回调
	if !exists {
		return c.eventHandler.ObjectDeleted(obj)
	}

	// 执行创建回调
	return c.eventHandler.ObjectCreated(obj)
}

func addObjectFunc(queue workqueue.RateLimitingInterface) func(interface{}) {
	return func(obj interface{}) {
		// 获取标识符并添加到队列
		key, err := cache.MetaNamespaceKeyFunc(obj)
		if err != nil {
			return
		}

		utils.Logger.Println(key)
		queue.Add(key)
	}
}

func deleteObjectFunc(queue workqueue.RateLimitingInterface) func(interface{}) {
	return func(obj interface{}) {
		// 获取标识符并添加到队列
		key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
		if err != nil {
			return
		}

		utils.Logger.Println(key)
		queue.Add(key)
	}
}
