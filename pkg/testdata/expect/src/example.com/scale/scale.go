package scale

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/scale"
	scalefake "k8s.io/client-go/scale/fake"
)

func Scale_117() {
	var c scale.ScaleInterface
	s := scale.New(nil, nil, nil, nil)
	f := &scalefake.FakeScaleClient{}

	c.Get(context.TODO(), schema.GroupResource{}, "", metav1.GetOptions{})
	s.Get(context.TODO(), schema.GroupResource{}, "", metav1.GetOptions{})
	f.Get(context.TODO(), schema.GroupResource{}, "", metav1.GetOptions{})

	c.Update(context.TODO(), schema.GroupResource{}, &autoscalingv1.Scale{}, metav1.UpdateOptions{})
	s.Update(context.TODO(), schema.GroupResource{}, &autoscalingv1.Scale{}, metav1.UpdateOptions{})
	f.Update(context.TODO(), schema.GroupResource{}, &autoscalingv1.Scale{}, metav1.UpdateOptions{})

	c.Patch(context.TODO(), schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{})
	s.Patch(context.TODO(), schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{})
	f.Patch(context.TODO(), schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{}, metav1.PatchOptions{})
}

func Scale_118() {
	ctx := context.TODO()
	getOpts := metav1.GetOptions{}
	updateOpts := metav1.UpdateOptions{}
	patchOpts := metav1.PatchOptions{}

	var c scale.ScaleInterface
	s := scale.New(nil, nil, nil, nil)
	f := &scalefake.FakeScaleClient{}

	c.Get(ctx, schema.GroupResource{}, "", getOpts)
	s.Get(ctx, schema.GroupResource{}, "", getOpts)
	f.Get(ctx, schema.GroupResource{}, "", getOpts)

	c.Update(ctx, schema.GroupResource{}, &autoscalingv1.Scale{}, updateOpts)
	s.Update(ctx, schema.GroupResource{}, &autoscalingv1.Scale{}, updateOpts)
	f.Update(ctx, schema.GroupResource{}, &autoscalingv1.Scale{}, updateOpts)

	c.Patch(ctx, schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{}, patchOpts)
	s.Patch(ctx, schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{}, patchOpts)
	f.Patch(ctx, schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{}, patchOpts)
}
