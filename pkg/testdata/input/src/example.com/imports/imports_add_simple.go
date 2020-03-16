package misc

import (
	corev1 "k8s.io/api/core/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Imports_Add_Simple(c kubernetes.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(pod)
}
