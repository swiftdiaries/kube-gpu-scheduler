package resource

import (
	"fmt"
	
	"k8s.io/api/core/v1"
)

type Resource struct {
	MilliCPU float64
	Memory float64
	NvidiaGPU int64
}

func EmptyResource() *Resource {
	return &Resource{
		MilliCPU: 0,
		Memory: 0,
		NvidiaGPU: 0,
	}
}

func (r *Resource) Clone() *Resource {
	clone := &Resource{
		MilliCPU: r.MilliCPU,
		Memory: r.Memory,
		NvidiaGPU: r.NvidiaGPU,
	}
	return clone
}

var minMilliCPU float64 = 10
var minMemory float64 = 10 * 1024 * 1024
var minNvidiaGPU int64 = 1

func NewResource(rl v1.ResourceList) *Resource {
	r := EmptyResource()
	for rName, rQuant := range rl {
		switch rName {
		case v1.ResourceCPU:
			r.MilliCPU += rQuant.MilliValue()
		case v1.ResourceMemory:
			r.Memory += rQuant.Value()
		case v1.ResourceNvidiaGPU:
			r.NvidiaGPU += rQuant.Value()	
		}
	}
	return r
}

func (r *Resource) IsEmpty() bool {
	return r.MilliCPU < minMilliCPU && r.Memory < minMemory
}
