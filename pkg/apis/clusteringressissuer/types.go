package clusteringressissuer

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GroupName    = "crd.xmh"
	GroupVersion = "v1"
	ResourceName = "clusteringressissuers"
)

type ClusterIngressIssuer struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterIngressIssuerSpec `json:"spec"`
}

type ClusterIngressIssuerSpec struct {
	IssuerName string `json:"issuerName"`
}

type ClusterIngressIssuerList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []ClusterIngressIssuer `json:"items"`
}
