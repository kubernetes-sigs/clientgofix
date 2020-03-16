package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Watch_117(c kubernetes.Interface) {
	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").Watch(context.TODO(), opts)
	c.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
}

func Watch_118(c kubernetes.Interface) {
	ctx := context.TODO()
	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").Watch(ctx, opts)
	c.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
}

func Watch_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	c.CoreV1().Pods("").Watch(ctx, metav1.ListOptions{})

	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").Watch(context.TODO(), opts)

	c.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
	c.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})

	c.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
}
