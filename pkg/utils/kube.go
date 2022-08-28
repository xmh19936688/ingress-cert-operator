package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var kubeClientSet *kubernetes.Clientset
var Logger *logrus.Entry

func InitKubecli(cfg *rest.Config) error {

	if cfg.UserAgent == "" {
		cfg.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	cliSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return fmt.Errorf("failed to init kubecli: %v", err)
	}

	kubeClientSet = cliSet
	return nil
}

func GetKubecli() *kubernetes.Clientset {
	if kubeClientSet == nil {
		panic("Call InitKubecli first!")
	}
	return kubeClientSet
}
