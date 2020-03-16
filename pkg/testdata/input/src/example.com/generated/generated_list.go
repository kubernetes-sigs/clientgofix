package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func List_117(c kubernetes.Interface) {
	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").List(opts)
	c.CoreV1().Pods("").List(metav1.ListOptions{})
}

func List_118(c kubernetes.Interface) {
	ctx := context.TODO()
	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").List(ctx, opts)
	c.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
}

func List_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	c.CoreV1().Pods("").List(ctx)

	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").List(opts)

	c.CoreV1().Pods("").List()
	c.CoreV1().Pods("").List(metav1.ListOptions{})

	c.CoreV1().Pods("").List(context.TODO())
}
