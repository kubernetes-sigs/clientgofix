package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Delete_117(c kubernetes.Interface) {
	opts := &metav1.DeleteOptions{}
	c.CoreV1().Pods("").Delete("", opts)
	c.CoreV1().Pods("").Delete("", makeDeleteOptionsPtr())
	c.CoreV1().Pods("").Delete("", nil)
	c.CoreV1().Pods("").Delete("", &metav1.DeleteOptions{})
}

func Delete_118(c kubernetes.Interface) {
	ctx := context.TODO()
	opts := metav1.DeleteOptions{}
	c.CoreV1().Pods("").Delete(ctx, "", opts)
	c.CoreV1().Pods("").Delete(ctx, "", makeDeleteOptions())
	c.CoreV1().Pods("").Delete(context.TODO(), "", metav1.DeleteOptions{})
}

func makeDeleteOptionsPtr() *metav1.DeleteOptions {
	return &metav1.DeleteOptions{}
}

func makeDeleteOptions() metav1.DeleteOptions {
	return metav1.DeleteOptions{}
}
