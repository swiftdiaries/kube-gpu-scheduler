package pod

import (
	"fmt"
	
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/pkg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ListPodsByNamespace(clientset *kubernetes.Clientset, namespace string) (*v1.PodList) {
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	if len(pods.Items) > 0 {
		fmt.Printf("There are %d pods in the %s\n", len(pods.Items), namespace)
		for _, pod := range pods.Items {
			fmt.Printf("pod: %s\n", pod.GetName())
		}
	} else {
		fmt.Println("No pods found!")
	}
	return pods	
}