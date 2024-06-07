package pkg

import (
	"gitee.com/maoyuting0503109/k8s/consts"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd" //用于处理 Kubernetes 客户端配置
)

func InitK8s() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", consts.ConfigPath)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
