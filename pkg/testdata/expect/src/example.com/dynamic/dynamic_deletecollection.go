package dynamic

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dynamic "k8s.io/client-go/dynamic"
)

func DeleteCollection_117(c dynamic.Interface) {
	opts := &metav1.DeleteOptions{}
	listOpts := metav1.ListOptions{}
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(context.TODO(), *opts, listOpts)
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(context.TODO(), *makeDeleteOptionsPtr(), listOpts)
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
}

func DeleteCollection_118(c dynamic.Interface) {
	ctx := context.TODO()
	opts := metav1.DeleteOptions{}
	listOpts := metav1.ListOptions{}
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(ctx, opts, listOpts)
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(ctx, makeDeleteOptions(), listOpts)
	c.Resource(schema.GroupVersionResource{}).DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{})
}

func makeDeleteOptionsPtr() *metav1.DeleteOptions {
	return &metav1.DeleteOptions{}
}

func makeDeleteOptions() metav1.DeleteOptions {
	return metav1.DeleteOptions{}
}
