package node

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	corelisters "k8s.io/kubernetes/pkg/client/listers/core/v1"
)

func ReadyNodes(client clientset.Interface, stopChannel <-chan struct{}) ([]*v1.Node, error) {
	nl := GetNodeLister(client, stopChannel)
	nodes, err := nl.List(labels.Everything())
	if err != nil {
		return []*v1.Node{}, err
	}

	if len(nodes) == 0 {
		var err error
		nItems, err := client.Core().Nodes().List(metav1.ListOptions{})
		if err != nil {
			return []*v1.Node{}, err
		}

		for i, _ := range nItems.Items {
			node := nItems.Items[i]
			nodes = append(nodes, &node)
		}
	}

	readyNodes := make([]*v1.Node, 0, len(nodes))
	for _, node := range nodes {
		if IsReady(node) {
			readyNodes = append(readyNodes, node)
		}
	}
	return readyNodes, nil
}

func GetNodeLister(client clientset.Interface, stopChannel <-chan struct{}) corelisters.NodeLister {
	listWatcher := cache.NewListWatchFromClient(client.Core().RESTClient(), "nodes", v1.NamespaceAll, fields.Everything())
	store := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	nodeLister := corelisters.NewNodeLister(store)
	reflector := cache.NewReflector(listWatcher, &v1.Node{}, store, time.Hour)
	reflector.RunUntil(stopChannel)

	return nodeLister
}

func IsReady(node *v1.Node) bool {
	for i := range node.Status.Conditions {
		cond := &node.Status.Conditions[i]
		if cond.Type == v1.NodeReady && cond.Status != v1.ConditionTrue {
			fmt.Printf("Ignoring node %v with %v condition status %v", node.Name, cond.Type, cond.Status)
			return false
		}
	}
	return true
}
