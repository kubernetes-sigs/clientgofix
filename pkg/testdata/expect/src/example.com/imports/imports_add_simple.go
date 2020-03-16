package misc

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Imports_Add_Simple(c kubernetes.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(context.TODO(), pod, metav1.CreateOptions{})
}
