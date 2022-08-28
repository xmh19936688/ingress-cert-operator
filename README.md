# ingress-cert-operator

基于`cert-manager`自动为ingress颁发证书。

## Usage

- `git clone`
- `cd ingress-cert-operator`
- `go mod tidy && go mod vendor && go build ./cmd/operator`
- `cp /path/to/kubeconfig ./kubecofnig.yaml && ./operator`
- `kubectl apply -f ./files/crd.yaml`
- 修改`files/cr.yaml`中的`issuerName`为`kubectl get clusterissuer`获取的name
- `kubectl apply -f ./files/cr.yaml`
- 创建个ingress试试

## TODO

- [ ] 支持自定义kubeconfig路径
- [ ] 编写Dockerfile完成容器化并上传镜像到dockerHub
- [ ] 编写yaml并部署到k8s
- [ ] 支持在k8s集群内获取kubeconfig
- [ ] 打包operator并上传到operatorHub

## 参考引用

- [A deep dive into Kubernetes controllers](https://docs.bitnami.com/tutorials/a-deep-dive-into-kubernetes-controllers)
- [Kubewatch, an example of Kubernetes custom controller](https://docs.bitnami.com/tutorials/kubewatch-an-example-of-kubernetes-custom-controller/)
- [sample-controller](https://github.com/kubernetes/sample-controller)
