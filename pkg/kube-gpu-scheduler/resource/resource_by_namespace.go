package resource

import(
	
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/v1"
	helper "k8s.io/kubernetes/pkg/api/v1/resource"
	"k8s.io/kubernetes/pkg/client/clientset_generated/clientset"

	"github.com/swiftdiaries/kube-gpu-scheduler/pkg/api"
)

func ResourceLister(node *v1.Node, pods []*v1.Pod) (namespace string,[] *helper.Resource) {
	
}
