package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func DeleteCollection_117(c kubernetes.Interface) {
	opts := &metav1.DeleteOptions{}
	listOpts := metav1.ListOptions{}
	c.CoreV1().Pods("").DeleteCollection(context.TODO(), *opts, listOpts)
	c.CoreV1().Pods("").DeleteCollection(context.TODO(), *makeDeleteOptionsPtr(), listOpts)
	c.CoreV1().Pods("").DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
	c.CoreV1().Pods("").DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
}

func DeleteCollection_118(c kubernetes.Interface) {
	ctx := context.TODO()
	opts := metav1.DeleteOptions{}
	listOpts := metav1.ListOptions{}
	c.CoreV1().Pods("").DeleteCollection(ctx, opts, listOpts)
	c.CoreV1().Pods("").DeleteCollection(ctx, makeDeleteOptions(), listOpts)
	c.CoreV1().Pods("").DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
}

func makeDeleteOptionsPtr() *metav1.DeleteOptions {
	return &metav1.DeleteOptions{}
}

func makeDeleteOptions() metav1.DeleteOptions {
	return metav1.DeleteOptions{}
}
