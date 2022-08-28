package clusteringressissuer

import (
	"github.com/xmh19936688/ingress-cert-operator/pkg/global"
	"github.com/xmh19936688/ingress-cert-operator/pkg/utils"
)

type Handler struct{}

func (h Handler) ObjectCreated(obj interface{}) (err error) {
	issuer, ok := obj.(*ClusterIngressIssuer)
	if !ok {
		utils.Logger.Printf("object created with type: %T \n", obj)
		return
	}

	global.ClusterIssuerName = issuer.Spec.IssuerName
	utils.Logger.Printf("issuer created: %v, %v \n", issuer.Name, issuer.Spec.IssuerName)
	return
}

func (h Handler) ObjectDeleted(obj interface{}) (err error) {
	issuer, ok := obj.(*ClusterIngressIssuer)
	if !ok {
		utils.Logger.Printf("object deleted with type: %T \n", obj)
		return
	}

	global.ClusterIssuerName = ""
	utils.Logger.Printf("issuer deleted: %v \n", issuer.Name)
	return
}
