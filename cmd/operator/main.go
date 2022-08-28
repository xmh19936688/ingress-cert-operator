package main

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"

	"github.com/xmh19936688/ingress-cert-operator/pkg/apis/clusteringressissuer"
	"github.com/xmh19936688/ingress-cert-operator/pkg/apis/ingress"
	"github.com/xmh19936688/ingress-cert-operator/pkg/controller"
	"github.com/xmh19936688/ingress-cert-operator/pkg/utils"
)

func main() {
	utils.Logger = logrus.NewEntry(logrus.New())

	// 读取kubeconfig
	// TODO use rest.InClusterConfig() in pod
	cfg, err := clientcmd.BuildConfigFromFlags("", "kubeconfig.yaml")
	if err != nil {
		utils.Logger.Println("failed to get kubeconfig:", err.Error())
		return
	}

	// 初始化kube-cli
	err = utils.InitKubecli(cfg)
	if err != nil {
		utils.Logger.Println("failed to init kubecli:", err.Error())
		return
	}

	ciiCli, err := clusteringressissuer.NewClient(cfg)
	if err != nil {
		utils.Logger.Println("failed to init kubecli:", err.Error())
		return
	}

	// 初始化controller
	ingressCtl := (&controller.Controller{}).
		SetLogger(utils.Logger).
		SetQueue(workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())).
		SetKubeClient(utils.GetKubecli()).
		SetHandler(ingress.Handler{}).
		SetInformer(ingress.NewInformer())
	issuerCtl := (&controller.Controller{}).
		SetLogger(utils.Logger).
		SetQueue(workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())).
		SetKubeClient(utils.GetKubecli()).
		SetHandler(clusteringressissuer.Handler{}).
		SetInformer(clusteringressissuer.NewInformer(ciiCli))

	// 启动controller
	chStop := make(chan struct{})
	defer close(chStop)
	go ingressCtl.Run(chStop)
	go issuerCtl.Run(chStop)

	select {}
}
