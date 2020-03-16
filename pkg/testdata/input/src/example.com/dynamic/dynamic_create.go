package dynamic

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	dynamic "k8s.io/client-go/dynamic"
)

func Create_117(c dynamic.Interface) {
	u := &unstructured.Unstructured{}
	opts := metav1.CreateOptions{}
	subresources := []string{}
	c.Resource(schema.GroupVersionResource{}).Create(u, opts)
	c.Resource(schema.GroupVersionResource{}).Create(u, metav1.CreateOptions{})
	c.Resource(schema.GroupVersionResource{}).Create(u, metav1.CreateOptions{}, "subresource")
	c.Resource(schema.GroupVersionResource{}).Create(u, metav1.CreateOptions{}, subresources...)
}

func Create_118(c dynamic.Interface) {
	ctx := context.TODO()
	u := &unstructured.Unstructured{}
	opts := metav1.CreateOptions{}
	subresources := []string{}
	c.Resource(schema.GroupVersionResource{}).Create(ctx, u, opts)
	c.Resource(schema.GroupVersionResource{}).Create(ctx, u, opts, "subresource")
	c.Resource(schema.GroupVersionResource{}).Create(ctx, u, opts, subresources...)
	c.Resource(schema.GroupVersionResource{}).Create(context.TODO(), &unstructured.Unstructured{}, metav1.CreateOptions{})
	c.Resource(schema.GroupVersionResource{}).Create(context.TODO(), &unstructured.Unstructured{}, metav1.CreateOptions{}, "subresource")
	c.Resource(schema.GroupVersionResource{}).Create(context.TODO(), &unstructured.Unstructured{}, metav1.CreateOptions{}, subresources...)
}
