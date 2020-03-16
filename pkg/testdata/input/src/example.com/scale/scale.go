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

	c.Get(schema.GroupResource{}, "")
	s.Get(schema.GroupResource{}, "")
	f.Get(schema.GroupResource{}, "")

	c.Update(schema.GroupResource{}, &autoscalingv1.Scale{})
	s.Update(schema.GroupResource{}, &autoscalingv1.Scale{})
	f.Update(schema.GroupResource{}, &autoscalingv1.Scale{})

	c.Patch(schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{})
	s.Patch(schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{})
	f.Patch(schema.GroupVersionResource{}, "", types.ApplyPatchType, []byte{})
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
