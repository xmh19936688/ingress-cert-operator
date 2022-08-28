package ingress

import (
	"context"

	networking_v1 "k8s.io/api/networking/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/xmh19936688/ingress-cert-operator/pkg/global"
	"github.com/xmh19936688/ingress-cert-operator/pkg/utils"
)

type Handler struct{}

// 对象已创建的回调
func (h Handler) ObjectCreated(obj interface{}) (err error) {
	ing, ok := obj.(*networking_v1.Ingress)
	if !ok {
		utils.Logger.Printf("object created with type: %T \n", obj)
		return
	}
	utils.Logger.Printf("ingress created: %v \n", ing.Name)

	if len(global.ClusterIssuerName) == 0 {
		return
	}

	// 忽略operator启动之前创建的ingress
	if ing.GetCreationTimestamp().Time.Before(global.Now) {
		return
	}

	// 忽略已设置集群颁发者的ingress
	if _, ok := ing.Annotations["cert-manager.io/cluster-issuer"]; ok {
		return
	}

	// 忽略已设置命名空间颁发者的ingress
	if _, ok := ing.Annotations["cert-manager.io/issuer"]; ok {
		return
	}

	// 给ingress添加issuer注解，设置颁发者
	utils.Logger.Printf("ingress [%v] issuing by [%v] \n", ing.Name, global.ClusterIssuerName)
	ing.Annotations["cert-manager.io/cluster-issuer"] = global.ClusterIssuerName
	_, err = utils.GetKubecli().NetworkingV1().Ingresses(ing.Namespace).Update(context.Background(), ing, meta_v1.UpdateOptions{})

	return
}

// 对象已删除的回调
func (h Handler) ObjectDeleted(obj interface{}) (err error) {
	ing, ok := obj.(*networking_v1.Ingress)
	if !ok {
		utils.Logger.Printf("object deleted with type: %T \n", obj)
		return
	}

	utils.Logger.Printf("object deleted: %v \n", ing.Name)
	return
}
