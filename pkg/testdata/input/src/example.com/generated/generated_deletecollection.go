package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"example.com/clientset/versioned"
)

func DeleteCollection_117(c kubernetes.Interface, customClient versioned.Interface) {
	opts := &metav1.DeleteOptions{}
	listOpts := metav1.ListOptions{}
	c.CoreV1().Pods("").DeleteCollection(opts, listOpts)
	c.CoreV1().Pods("").DeleteCollection(makeDeleteOptionsPtr(), listOpts)
	c.CoreV1().Pods("").DeleteCollection(nil, metav1.ListOptions{})
	c.CoreV1().Pods("").DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(opts, listOpts)
	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(makeDeleteOptionsPtr(), listOpts)
	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(nil, metav1.ListOptions{})
	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{})
}

func DeleteCollection_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	opts := metav1.DeleteOptions{}
	listOpts := metav1.ListOptions{}
	c.CoreV1().Pods("").DeleteCollection(ctx, opts, listOpts)
	c.CoreV1().Pods("").DeleteCollection(ctx, makeDeleteOptions(), listOpts)
	c.CoreV1().Pods("").DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(ctx, opts, listOpts)
	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(ctx, makeDeleteOptions(), listOpts)
	customClient.SamplecontrollerV1alpha1().Foos("").DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
}

func makeDeleteOptionsPtr() *metav1.DeleteOptions {
	return &metav1.DeleteOptions{}
}

func makeDeleteOptions() metav1.DeleteOptions {
	return metav1.DeleteOptions{}
}
