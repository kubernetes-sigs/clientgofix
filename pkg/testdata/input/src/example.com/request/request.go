package request

import (
	"context"

	"k8s.io/client-go/rest"
)

func Do(r *rest.Request) {
	r.Do()
	r.DoRaw()
	r.Stream()
	r.Watch()

	ctx := context.TODO()
	r.Do(ctx)
	r.DoRaw(ctx)
	r.Stream(ctx)
	r.Watch(ctx)

	r.Do(context.TODO())
	r.DoRaw(context.TODO())
	r.Stream(context.TODO())
	r.Watch(context.TODO())
}
