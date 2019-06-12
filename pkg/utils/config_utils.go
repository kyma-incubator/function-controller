package utils

import (
	"k8s.io/client-go/rest"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

//var log = logf.Log.WithName("function_controller")

func InitKubeConfig() *rest.Config {

	// Get a config to talk to the apiserver
	log.Info("Build kube config")
	kubeConfig, err := config.GetConfig()
	if err != nil {
		log.Error(err, "Unable to build kubeconfig")
		os.Exit(1)
	}

	return kubeConfig
}
