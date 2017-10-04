package resource

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api/v1"
	helper "k8s.io/kubernetes/pkg/api/v1/resource"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
)

type NamespaceResourceMap map[string]*schedulercache.Resource

type NamespaceCPUMap map[string]int64
type NamespaceMemoryMap map[string]float64
type NamespaceGPUMap map[string]int64


func ResourceListerforPod(pods []*v1.Pod, nrm NamespaceResourceMap, ncm NamespaceCPUMap, nmm NamespaceMemoryMap, ngm NamespaceGPUMap) NamespaceResourceMap {
	for _, pod := range pods {
		var namespace = pod.Namespace
		podResource := &schedulercache.Resource{
			MilliCPU:  0,
			Memory:    0,
			NvidiaGPU: 0,
		}
		req, _, err := helper.PodRequestsAndLimits(pod)
		if err != nil {
			fmt.Printf("Error computing resource usage of pod, ignoring: %#v\n", pod.Name)
			continue
		}
		for name, quantity := range req {
			if name == v1.ResourceCPU {
				podResource.MilliCPU += quantity.Value()
				ncm[namespace] += quantity.Value() 
			} else if name == v1.ResourceMemory {
				podResource.Memory += quantity.Value()
				ncm[namespace] += quantity.Value()
			} else if name == v1.ResourceNvidiaGPU {
				podResource.NvidiaGPU += quantity.Value()
				ngm[namespace] += quantity.Value()
			}
		}
		nrm[namespace] = podResource
	}
	return nrm
}
