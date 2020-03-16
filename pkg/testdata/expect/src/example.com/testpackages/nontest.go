package testpackages

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

func HandleGet(c dynamic.Interface) {
	c.Resource(schema.GroupVersionResource{}).Get(context.TODO(), "", metav1.GetOptions{})
}

func HandleList(c dynamic.Interface) {
	listOpts := metav1.ListOptions{
		LabelSelector: action.GetListRestrictions().Labels.String(),
		FieldSelector: action.GetListRestrictions().Fields.String(),
	}
	c.Resource(schema.GroupVersionResource{}).List(context.TODO(), listOpts)
}
