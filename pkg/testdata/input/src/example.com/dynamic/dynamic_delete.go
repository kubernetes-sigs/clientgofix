package dynamic

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dynamic "k8s.io/client-go/dynamic"
)

func Delete_117(c dynamic.Interface) {
	opts := &metav1.DeleteOptions{}
	subresources := []string{}
	c.Resource(schema.GroupVersionResource{}).Delete("", opts)
	c.Resource(schema.GroupVersionResource{}).Delete("", opts, "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete("", opts, subresources...)
	c.Resource(schema.GroupVersionResource{}).Delete("", makeDeleteOptionsPtr())
	c.Resource(schema.GroupVersionResource{}).Delete("", makeDeleteOptionsPtr(), "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete("", makeDeleteOptionsPtr(), subresources...)
	c.Resource(schema.GroupVersionResource{}).Delete("", nil)
	c.Resource(schema.GroupVersionResource{}).Delete("", nil, "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete("", nil, subresources...)
	c.Resource(schema.GroupVersionResource{}).Delete("", &metav1.DeleteOptions{})
	c.Resource(schema.GroupVersionResource{}).Delete("", &metav1.DeleteOptions{}, "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete("", &metav1.DeleteOptions{}, subresources...)
}

func Delete_118(c dynamic.Interface) {
	ctx := context.TODO()
	opts := metav1.DeleteOptions{}
	subresources := []string{}
	c.Resource(schema.GroupVersionResource{}).Delete(ctx, "", opts)
	c.Resource(schema.GroupVersionResource{}).Delete(ctx, "", opts, "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete(ctx, "", opts, subresources...)
	c.Resource(schema.GroupVersionResource{}).Delete(ctx, "", makeDeleteOptions())
	c.Resource(schema.GroupVersionResource{}).Delete(ctx, "", makeDeleteOptions(), "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete(ctx, "", makeDeleteOptions(), subresources...)
	c.Resource(schema.GroupVersionResource{}).Delete(context.TODO(), "", metav1.DeleteOptions{})
	c.Resource(schema.GroupVersionResource{}).Delete(context.TODO(), "", metav1.DeleteOptions{}, "subresource")
	c.Resource(schema.GroupVersionResource{}).Delete(context.TODO(), "", metav1.DeleteOptions{}, subresources...)
}

func makeDeleteOptionsPtr() *metav1.DeleteOptions {
	return &metav1.DeleteOptions{}
}

func makeDeleteOptions() metav1.DeleteOptions {
	return metav1.DeleteOptions{}
}
