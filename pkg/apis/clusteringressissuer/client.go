package clusteringressissuer

import (
	"context"
	"fmt"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type Client struct {
	cli *rest.RESTClient
}

func NewClient(cfg *rest.Config) (cli Client, err error) {
	if cfg.UserAgent == "" {
		cfg.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	cfg.APIPath = "/apis" // 解决报错`the server could not find the requested resource`
	cfg.GroupVersion = &SchemeGroupVersion
	cfg.NegotiatedSerializer = Codecs.WithoutConversion() // 解决报错`no kind \"ClusterIngressIssuerList\" is registered for version \"crd.xmh/v1\"`

	httpClient, err := rest.HTTPClientFor(cfg)
	if err != nil {
		err = fmt.Errorf("failed to init kubecli: %v", err)
		return
	}

	restCli, err := rest.RESTClientForConfigAndClient(cfg, httpClient)
	if err != nil {
		err = fmt.Errorf("failed to init kubecli: %v", err)
		return
	}

	cli = Client{cli: restCli}
	return
}

func (c Client) List(options meta_v1.ListOptions) (runtime.Object, error) {
	res := &ClusterIngressIssuerList{}
	err := c.cli.
		Get().
		Resource(ResourceName).
		VersionedParams(&options, ParameterCodec).
		Do(context.Background()).
		Into(res)
	return res, err
}

func (c Client) Watch(options meta_v1.ListOptions) (watch.Interface, error) {
	options.Watch = true
	return c.cli.
		Get().
		Resource(ResourceName).
		VersionedParams(&options, ParameterCodec).
		Watch(context.Background())
}
