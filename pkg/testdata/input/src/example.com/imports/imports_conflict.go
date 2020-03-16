package misc

import (
	context "fmt"

	policyapi "k8s.io/api/policy/v1beta1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

const context2 = ""

var context3 = ""

func context4() {}

type context5 string

func AddContextImport(c coreclient.PodInterface, context6 string) {
	var context7 string
	context8 := ""
	c.Evict(&policyapi.Eviction{})
}

func X() {
	context.Println()
}
