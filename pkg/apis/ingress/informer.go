package ingress

import (
	"context"

	networking_v1 "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	"github.com/xmh19936688/ingress-cert-operator/pkg/utils"
)

// 初始化SharedInformer
func NewInformer() cache.SharedIndexInformer {
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc:  listIngress,
			WatchFunc: watchIngress,
		},
		&networking_v1.Ingress{},
		0,
		cache.Indexers{},
	)
	return informer
}

func listIngress(options meta_v1.ListOptions) (runtime.Object, error) {
	// 指定其list接口
	return utils.GetKubecli().NetworkingV1().Ingresses(meta_v1.NamespaceAll).List(context.Background(), options)
}

func watchIngress(options meta_v1.ListOptions) (watch.Interface, error) {
	// 指定其watch接口
	return utils.GetKubecli().NetworkingV1().Ingresses(meta_v1.NamespaceAll).Watch(context.Background(), options)
}
