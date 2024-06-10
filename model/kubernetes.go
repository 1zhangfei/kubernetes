package model

import (
	"context"
	"github.com/1zhangfei/kubernetes/pkg"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPods() (*v1.PodList, error) {
	clientset := pkg.InitK8s()
	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

func GetPod(namespace, podName string) (*v1.Pod, error) {
	clientset := pkg.InitK8s()
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

func CreatePod(namespace string, pod *v1.Pod) (*v1.Pod, error) {
	clientset := pkg.InitK8s()
	createdPod, err := clientset.CoreV1().Pods(namespace).Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return createdPod, nil
}

func UpdatePod(namespace string, pod *v1.Pod) (*v1.Pod, error) {
	clientset := pkg.InitK8s()
	updatedPod, err := clientset.CoreV1().Pods(namespace).Update(context.Background(), pod, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return updatedPod, nil
}

func DeletePod(namespace, podName string) error {
	clientset := pkg.InitK8s()
	err := clientset.CoreV1().Pods(namespace).Delete(context.Background(), podName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func GetServices() (*v1.ServiceList, error) {
	clientset := pkg.InitK8s()
	services, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}

func GetService(namespace, serviceName string) (*v1.Service, error) {
	clientset := pkg.InitK8s()
	service, err := clientset.CoreV1().Services(namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return service, nil
}

func CreateService(namespace string, service *v1.Service) (*v1.Service, error) {
	clientset := pkg.InitK8s()
	createdService, err := clientset.CoreV1().Services(namespace).Create(context.Background(), service, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return createdService, nil
}

func UpdateService(namespace string, service *v1.Service) (*v1.Service, error) {
	clientset := pkg.InitK8s()
	updatedService, err := clientset.CoreV1().Services(namespace).Update(context.Background(), service, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return updatedService, nil
}

func DeleteService(namespace, serviceName string) error {
	clientset := pkg.InitK8s()
	err := clientset.CoreV1().Services(namespace).Delete(context.Background(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func GetDeployments() (*appsv1.DeploymentList, error) {
	clientset := pkg.InitK8s()
	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

func GetDeployment(namespace, deploymentName string) (*appsv1.Deployment, error) {
	clientset := pkg.InitK8s()
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

func CreateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	clientset := pkg.InitK8s()
	createdDeployment, err := clientset.AppsV1().Deployments(namespace).Create(context.Background(), deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return createdDeployment, nil
}

func UpdateDeployment(namespace string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	clientset := pkg.InitK8s()
	updatedDeployment, err := clientset.AppsV1().Deployments(namespace).Update(context.Background(), deployment, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return updatedDeployment, nil
}

func DeleteDeployment(namespace, deploymentName string) error {
	clientset := pkg.InitK8s()
	err := clientset.AppsV1().Deployments(namespace).Delete(context.Background(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func GetStatefulSets() (*appsv1.StatefulSetList, error) {
	clientset := pkg.InitK8s()
	statefulSets, err := clientset.AppsV1().StatefulSets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return statefulSets, nil
}

func GetStatefulSet(namespace, statefulSetName string) (*appsv1.StatefulSet, error) {
	clientset := pkg.InitK8s()
	statefulSet, err := clientset.AppsV1().StatefulSets(namespace).Get(context.Background(), statefulSetName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return statefulSet, nil
}

func CreateStatefulSet(namespace string, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	clientset := pkg.InitK8s()
	createdStatefulSet, err := clientset.AppsV1().StatefulSets(namespace).Create(context.Background(), statefulSet, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return createdStatefulSet, nil
}

func UpdateStatefulSet(namespace string, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	clientset := pkg.InitK8s()
	updatedStatefulSet, err := clientset.AppsV1().StatefulSets(namespace).Update(context.Background(), statefulSet, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return updatedStatefulSet, nil
}

func DeleteStatefulSet(namespace, statefulSetName string) error {
	clientset := pkg.InitK8s()
	err := clientset.AppsV1().StatefulSets(namespace).Delete(context.Background(), statefulSetName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func GetNodes() (*v1.NodeList, error) {
	clientset := pkg.InitK8s()
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func GetNode(nodeName string) (*v1.Node, error) {
	clientset := pkg.InitK8s()
	node, err := clientset.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return node, nil
}

func UpdateNode(node *v1.Node) (*v1.Node, error) {
	clientset := pkg.InitK8s()
	updatedNode, err := clientset.CoreV1().Nodes().Update(context.Background(), node, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return updatedNode, nil
}

func DeleteNode(nodeName string) error {
	clientset := pkg.InitK8s()
	err := clientset.CoreV1().Nodes().Delete(context.Background(), nodeName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
func AddNode(node *v1.Node) (*v1.Node, error) {
	clientset := pkg.InitK8s()
	createdNode, err := clientset.CoreV1().Nodes().Create(context.Background(), node, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return createdNode, nil
}
