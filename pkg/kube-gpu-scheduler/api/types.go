package api

import (
	"k8s.io/kubernetes/pkg/api/v1"
	"github.com/swiftdiaries/kube-gpu-scheduler/pkg/resource"
)

type NamespaceResourceMap map[v1.Namespace] resource.*Resource