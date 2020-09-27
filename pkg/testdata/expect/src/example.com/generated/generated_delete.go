package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"example.com/clientset/versioned"
)

func Delete_117(c kubernetes.Interface, customClient versioned.Interface) {
	opts := &metav1.DeleteOptions{}
	c.CoreV1().Pods("").Delete(context.TODO(), "", *opts)
	c.CoreV1().Pods("").Delete(context.TODO(), "", *makeDeleteOptionsPtr())
	c.CoreV1().Pods("").Delete(context.TODO(), "", metav1.DeleteOptions{})
	c.CoreV1().Pods("").Delete(context.TODO(), "", metav1.DeleteOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").Delete(context.TODO(), "", *opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Delete(context.TODO(), "", *makeDeleteOptionsPtr())
	customClient.SamplecontrollerV1alpha1().Foos("").Delete(context.TODO(), "", metav1.DeleteOptions{})
	customClient.SamplecontrollerV1alpha1().Foos("").Delete(context.TODO(), "", metav1.DeleteOptions{})
}

func Delete_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	opts := metav1.DeleteOptions{}
	c.CoreV1().Pods("").Delete(ctx, "", opts)
	c.CoreV1().Pods("").Delete(ctx, "", makeDeleteOptions())
	c.CoreV1().Pods("").Delete(context.TODO(), "", metav1.DeleteOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").Delete(ctx, "", opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Delete(ctx, "", makeDeleteOptions())
	customClient.SamplecontrollerV1alpha1().Foos("").Delete(context.TODO(), "", metav1.DeleteOptions{})
}

func makeDeleteOptionsPtr() *metav1.DeleteOptions {
	return &metav1.DeleteOptions{}
}

func makeDeleteOptions() metav1.DeleteOptions {
	return metav1.DeleteOptions{}
}
