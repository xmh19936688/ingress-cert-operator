package clusteringressissuer

import "k8s.io/apimachinery/pkg/runtime"

func (in *ClusterIngressIssuer) DeepCopyInto(out *ClusterIngressIssuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

func (in *ClusterIngressIssuer) DeepCopy() *ClusterIngressIssuer {
	if in == nil {
		return nil
	}

	out := new(ClusterIngressIssuer)
	in.DeepCopyInto(out)
	return out
}

func (in *ClusterIngressIssuer) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}

	return in.DeepCopy()
}

func (in *ClusterIngressIssuerList) DeepCopyInto(out *ClusterIngressIssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)

	if in.Items != nil {
		ins, outs := &in.Items, &out.Items
		*outs = make([]ClusterIngressIssuer, len(*ins))
		for i := range *ins {
			(*ins)[i].DeepCopyInto(&(*outs)[i])
		}
	}

	return
}

func (in *ClusterIngressIssuerList) DeepCopy() *ClusterIngressIssuerList {
	if in == nil {
		return nil
	}

	out := new(ClusterIngressIssuerList)
	in.DeepCopyInto(out)
	return out
}

func (in *ClusterIngressIssuerList) DeepCopyObject() runtime.Object {
	if in == nil {
		return nil
	}

	return in.DeepCopy()
}
