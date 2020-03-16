package misc

import kubernetes "k8s.io/client-go/kubernetes"

func Imports_Add_Grouped(c kubernetes.Interface) {
	c.CoreV1().Pods("").Create(nil)
}
