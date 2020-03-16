package misc

import (
	policyapi "k8s.io/api/policy/v1beta1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

func AddContextImport(c coreclient.PodInterface) {
	c.Evict(&policyapi.Eviction{})
}
