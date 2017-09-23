package main

import (
	"fmt"
	
	"github.com/swiftdiaries/gpuscheduler/pkg/client"
	"github.com/swiftdiaries/gpuscheduler/pkg/pod"
)

func main() {
	
	clientset := client.CreateClient()

	pods := pod.ListPodsByNamespace(clientset, "default")
}