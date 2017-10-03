package pod

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
)

func ListPodsOnANode(client clientset.Interface, node *v1.Node) ([]*v1.Pod, error) {
	podList, err := client.CoreV1().Pods(v1.NamespaceAll).List(
		metav1.ListOptions{FieldSelector: fields.SelectorFromSet(fields.Set{"spec.nodeName": node.Name}).String()})
	if err != nil {
		return []*v1.Pod{}, err
	}

	pods := make([]*v1.Pod, 0)
	for i := range podList.Items {
		pods = append(pods, &podList.Items[i])
	}

	return pods, nil
}
