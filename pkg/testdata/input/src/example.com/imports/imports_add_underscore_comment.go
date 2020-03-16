package misc

import (
	// before archive comment
	_ "archive/tar" // archive/tar comment
	"fmt"           // fmt comment

	corev1 "k8s.io/api/core/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Imports_Add_Grouped(c kubernetes.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(pod)
	fmt.Println()
}
