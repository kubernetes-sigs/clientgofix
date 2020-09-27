package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"example.com/clientset/versioned"
)

func Watch_117(c kubernetes.Interface, customClient versioned.Interface) {
	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").Watch(opts)
	c.CoreV1().Pods("").Watch(metav1.ListOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").Watch(opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Watch(metav1.ListOptions{})
}

func Watch_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").Watch(ctx, opts)
	c.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").Watch(ctx, opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Watch(context.TODO(), metav1.ListOptions{})
}

func Watch_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	c.CoreV1().Pods("").Watch(ctx)

	opts := metav1.ListOptions{}
	c.CoreV1().Pods("").Watch(opts)

	c.CoreV1().Pods("").Watch()
	c.CoreV1().Pods("").Watch(metav1.ListOptions{})

	c.CoreV1().Pods("").Watch(context.TODO())
}
