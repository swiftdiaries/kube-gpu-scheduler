package resource

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api/v1"
	helper "k8s.io/kubernetes/pkg/api/v1/resource"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
)

type NamespaceResourceMap map[string]*schedulercache.Resource

func ResourceListerforPod(pods []*v1.Pod, nrm NamespaceResourceMap) NamespaceResourceMap {
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
				podResource.MilliCPU += quantity.MilliValue()
			} else if name == v1.ResourceMemory {
				podResource.Memory += quantity.Value()
			} else if name == v1.ResourceNvidiaGPU {
				podResource.NvidiaGPU += quantity.Value()
			}
		}
		nrm[namespace] = podResource
	}
	return nrm
}
