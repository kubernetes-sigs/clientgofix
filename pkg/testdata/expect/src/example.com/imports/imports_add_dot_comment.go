package misc

import (
	// before archive comment
	. "archive/tar" // archive/tar comment
	"context"
	"fmt" // fmt comment

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Imports_Add_Grouped(c kubernetes.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(context.TODO(), pod, metav1.CreateOptions{})
	fmt.Println()
}

// TODO: fix
