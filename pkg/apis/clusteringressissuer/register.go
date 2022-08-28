package clusteringressissuer

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	Scheme             = runtime.NewScheme()
	Codecs             = serializer.NewCodecFactory(Scheme)
	ParameterCodec     = runtime.NewParameterCodec(Scheme) // 解决`"unable to decode an event from the watch stream: unable to decode to metav1.Event"`
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}

	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder = runtime.SchemeBuilder{
		SchemeBuilder.AddToScheme,
	}
)

func init() {
	// 解决`no kind is registered for the type v1.ListOptions in scheme`
	utilruntime.Must(localSchemeBuilder.AddToScheme(Scheme))
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ClusterIngressIssuer{},
		&ClusterIngressIssuerList{},
	)

	// 解决`no kind is registered for the type v1.ListOptions in scheme`
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
