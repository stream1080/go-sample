package global

// import (
// 	"path/filepath"

// 	"go.uber.org/zap"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/rest"
// 	"k8s.io/client-go/tools/clientcmd"
// 	"k8s.io/client-go/util/homedir"
// )

// func InitKubeCli() {
// 	config, err := getKubeConfig()
// 	if err != nil {
// 		zap.L().Panic("load kube config err", zap.String("err", err.Error()))
// 	}

// 	KubeCli, err = kubernetes.NewForConfig(config)
// 	if err != nil {
// 		zap.L().Panic("creates a new client set err", zap.String("err", err.Error()))
// 	}

// }

// func getKubeConfig() (*rest.Config, error) {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		zap.L().Warn("load cluster kube config", zap.Error(err))
// 		kubeConfig := ""
// 		if home := homedir.HomeDir(); home != "" {
// 			kubeConfig = filepath.Join(home, ".kube/config")
// 		}
// 		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return config, nil
// }
