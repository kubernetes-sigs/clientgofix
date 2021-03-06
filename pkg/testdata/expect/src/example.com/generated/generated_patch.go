package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kubernetes "k8s.io/client-go/kubernetes"

	"example.com/clientset/versioned"
)

func Patch_117(c kubernetes.Interface, customClient versioned.Interface) {
	data := []byte{}
	subresources := []string{}
	c.CoreV1().Pods("").Patch(context.TODO(), "", types.ApplyPatchType, data, metav1.PatchOptions{})
	c.CoreV1().Pods("").Patch(context.TODO(), "", types.ApplyPatchType, data, metav1.PatchOptions{}, "status")
	c.CoreV1().Pods("").Patch(context.TODO(), "", types.ApplyPatchType, data, metav1.PatchOptions{}, subresources...)

	customClient.SamplecontrollerV1alpha1().Foos("").Patch(context.TODO(), "", types.ApplyPatchType, data, metav1.PatchOptions{})
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(context.TODO(), "", types.ApplyPatchType, data, metav1.PatchOptions{}, "status")
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(context.TODO(), "", types.ApplyPatchType, data, metav1.PatchOptions{}, subresources...)
}

func Patch_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	data := []byte{}
	opts := metav1.PatchOptions{}
	subresources := []string{}
	c.CoreV1().Pods("").Patch(ctx, "", types.ApplyPatchType, data, opts)
	c.CoreV1().Pods("").Patch(ctx, "", types.ApplyPatchType, data, opts, "status")
	c.CoreV1().Pods("").Patch(ctx, "", types.ApplyPatchType, data, opts, subresources...)
	c.CoreV1().Pods("").Patch(context.TODO(), "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{})
	c.CoreV1().Pods("").Patch(context.TODO(), "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{}, "status")
	c.CoreV1().Pods("").Patch(context.TODO(), "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{}, "status", "status2")

	customClient.SamplecontrollerV1alpha1().Foos("").Patch(ctx, "", types.ApplyPatchType, data, opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(ctx, "", types.ApplyPatchType, data, opts, "status")
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(ctx, "", types.ApplyPatchType, data, opts, subresources...)
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(context.TODO(), "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{})
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(context.TODO(), "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{}, "status")
	customClient.SamplecontrollerV1alpha1().Foos("").Patch(context.TODO(), "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{}, "status", "status2")
}
