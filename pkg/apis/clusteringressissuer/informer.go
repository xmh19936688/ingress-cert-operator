package clusteringressissuer

import (
	"k8s.io/client-go/tools/cache"
)

func NewInformer(c Client) cache.SharedIndexInformer {
	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc:  c.List,
			WatchFunc: c.Watch,
		},
		&ClusterIngressIssuer{},
		0,
		cache.Indexers{},
	)

	return informer
}
