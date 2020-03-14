/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import "strings"

type SelectorMatcher func(selectorType, method string) bool

func matchTypeAndMethod(expectedType, expectedMethod string) SelectorMatcher {
	return func(selectorType, method string) bool {
		return expectedType == selectorType && expectedMethod == method
	}
}
func matchTypePrefixAndMethod(prefixes []string, expectedMethod string) SelectorMatcher {
	return func(selectorType, method string) bool {
		if expectedMethod != method {
			return false
		}
		for _, prefix := range prefixes {
			if strings.HasPrefix(selectorType, prefix) {
				return true
			}
		}
		return false
	}
}

var dynamicClientPrefixes = []string{
	"k8s.io/client-go/dynamic.",
	"k8s.io/client-go/metadata.",
}
var generatedClientPrefixes = []string{
	"k8s.io/client-go/kubernetes/typed/",
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/",
	"k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/typed/",
}

var transforms = []struct {
	name      string
	matcher   SelectorMatcher
	transform Transformer
}{
	// Expansions
	// git diff upstream/release-1.17 upstream/release-1.18 --name-only -- staging/src/k8s.io/client-go/kubernetes/typed/ | grep expansion | egrep -v 'generated|fake|authorization|authentication' | xargs -n 1 git diff upstream/release-1.17 upstream/release-1.18 -- | egrep '^[-+]\t[A-Z]'

	// - UpdateApproval(                     certificateSigningRequest *certificates.CertificateSigningRequest                           ) (result *certificates.CertificateSigningRequest, err error)
	// + UpdateApproval(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts metav1.UpdateOptions) (result *certificates.CertificateSigningRequest, err error)
	{
		"UpdateApproval",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "UpdateApproval"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.UpdateOptions", makeMetav1OptionsArg("UpdateOptions")),
		},
	},

	// - Finalize(                     item *v1.Namespace                           ) (*v1.Namespace, error)
	// + Finalize(ctx context.Context, item *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error)
	{
		"Finalize",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "Finalize"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.UpdateOptions", makeMetav1OptionsArg("UpdateOptions")),
		},
	},

	// - PatchStatus(                     nodeName string, data []byte) (*v1.Node, error)
	// + PatchStatus(ctx context.Context, nodeName string, data []byte) (*v1.Node, error)
	{
		"PatchStatus",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "PatchStatus"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},

	// - Bind(                     binding *v1.Binding                           ) error
	// + Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error
	{
		"Bind",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "Bind"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.CreateOptions", makeMetav1OptionsArg("CreateOptions")),
		},
	},

	// - Evict(                     eviction *policy.Eviction) error
	// + Evict(ctx context.Context, eviction *policy.Eviction) error
	{
		"Evict",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "Evict"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},

	// - CreateToken(                     name string, tokenRequest *authenticationv1.TokenRequest                           ) (*authenticationv1.TokenRequest, error)
	// + CreateToken(ctx context.Context, name string, tokenRequest *authenticationv1.TokenRequest, opts metav1.CreateOptions) (*authenticationv1.TokenRequest, error)
	{
		"CreateToken",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "CreateToken"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.CreateOptions", makeMetav1OptionsArg("CreateOptions")),
		},
	},

	// - Rollback(                 *v1beta1.DeploymentRollback                      ) error
	// + Rollback(context.Context, *v1beta1.DeploymentRollback, metav1.CreateOptions) error
	{
		"Rollback",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/kubernetes/typed/"}, "Rollback"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.CreateOptions", makeMetav1OptionsArg("CreateOptions")),
		},
	},

	// Scale
	// git diff upstream/release-1.17 upstream/release-1.18 -- staging/src/k8s.io/client-go/scale/interfaces.go

	// - Get(                     resource schema.GroupResource, name string                        ) (*autoscalingapi.Scale, error)
	// + Get(ctx context.Context, resource schema.GroupResource, name string, opts metav1.GetOptions) (*autoscalingapi.Scale, error)
	{
		"Get",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/scale"}, "Get"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.GetOptions", makeMetav1OptionsArg("GetOptions")),
		},
	},
	// - Update(                     resource schema.GroupResource, scale *autoscalingapi.Scale                           ) (*autoscalingapi.Scale, error)
	// + Update(ctx context.Context, resource schema.GroupResource, scale *autoscalingapi.Scale, opts metav1.UpdateOptions) (*autoscalingapi.Scale, error)
	{
		"Update",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/scale"}, "Update"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.UpdateOptions", makeMetav1OptionsArg("UpdateOptions")),
		},
	},
	// - Patch(                     gvr schema.GroupVersionResource, name string, pt types.PatchType, data []byte                          ) (*autoscalingapi.Scale, error)
	// + Patch(ctx context.Context, gvr schema.GroupVersionResource, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions) (*autoscalingapi.Scale, error)
	{
		"Patch",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/scale"}, "Patch"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.PatchOptions", makeMetav1OptionsArg("PatchOptions")),
		},
	},

	// Request Do/DoRaw

	// - func (r *Request) Do(                   ) Result {
	// + func (r *Request) Do(ctx context.Context) Result {
	{
		"Do",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/rest."}, "Do"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// - DoRaw(               ) ([]byte, error)
	// + DoRaw(context.Context) ([]byte, error)
	{
		"DoRaw",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/rest."}, "DoRaw"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// - Stream(               ) (io.ReadCloser, error)
	// + Stream(context.Context) (io.ReadCloser, error)
	{
		"Stream",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/rest."}, "Stream"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// - Watch(               ) (watch.Interface, error)
	// + Watch(context.Context) (watch.Interface, error)
	{
		"Watch",
		matchTypePrefixAndMethod([]string{"k8s.io/client-go/rest."}, "Watch"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},

	// Generated clientset methods
	{
		"Get",
		matchTypePrefixAndMethod(generatedClientPrefixes, "Get"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.GetOptions", makeMetav1OptionsArg("GetOptions")),
		},
	},
	{
		"List",
		matchTypePrefixAndMethod(generatedClientPrefixes, "List"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.ListOptions", makeMetav1OptionsArg("ListOptions")),
		},
	},
	{
		"Watch",
		matchTypePrefixAndMethod(generatedClientPrefixes, "Watch"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.ListOptions", makeMetav1OptionsArg("ListOptions")),
		},
	},
	{
		"Create",
		matchTypePrefixAndMethod(generatedClientPrefixes, "Create"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.CreateOptions", makeMetav1OptionsArg("CreateOptions")),
		},
	},
	{
		"Update",
		matchTypePrefixAndMethod(generatedClientPrefixes, "Update"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.UpdateOptions", makeMetav1OptionsArg("UpdateOptions")),
		},
	},
	{
		"UpdateStatus",
		matchTypePrefixAndMethod(generatedClientPrefixes, "UpdateStatus"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureLastArg("k8s.io/apimachinery/pkg/apis/meta/v1.UpdateOptions", makeMetav1OptionsArg("UpdateOptions")),
		},
	},
	{
		"Patch",
		matchTypePrefixAndMethod(generatedClientPrefixes, "Patch"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			ensureArgAtIndex(4, "k8s.io/apimachinery/pkg/apis/meta/v1.PatchOptions", makeMetav1OptionsArg("PatchOptions")),
		},
	},
	{
		"Delete",
		matchTypePrefixAndMethod(generatedClientPrefixes, "Delete"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			dereferenceArgAtIndexIfPointer(2, "*k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions"),
			replaceArgAtIndexIfNil(2, "k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions", makeMetav1OptionsArg("DeleteOptions")),
		},
	},
	{
		"DeleteCollection",
		matchTypePrefixAndMethod(generatedClientPrefixes, "DeleteCollection"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			dereferenceArgAtIndexIfPointer(1, "*k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions"),
			replaceArgAtIndexIfNil(1, "k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions", makeMetav1OptionsArg("DeleteOptions")),
		},
	},

	// Dynamic/Metadata client methods

	// dynamic:
	// - Get(                     name string, options metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error)
	// + Get(ctx context.Context, name string, options metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error)
	// metadata:
	// - Get(                     name string, options metav1.GetOptions, subresources ...string) (*metav1.PartialObjectMetadata, error)
	// + Get(ctx context.Context, name string, options metav1.GetOptions, subresources ...string) (*metav1.PartialObjectMetadata, error)
	{
		"Get",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "Get"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic:
	// - List(                     opts metav1.ListOptions) (*unstructured.UnstructuredList, error)
	// + List(ctx context.Context, opts metav1.ListOptions) (*unstructured.UnstructuredList, error)
	// metadata:
	// - List(                     opts metav1.ListOptions) (*metav1.PartialObjectMetadataList, error)
	// + List(ctx context.Context, opts metav1.ListOptions) (*metav1.PartialObjectMetadataList, error)
	{
		"List",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "List"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic:
	// - Watch(                     opts metav1.ListOptions) (watch.Interface, error)
	// + Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	// metadata:
	// - Watch(                     opts metav1.ListOptions) (watch.Interface, error)
	// + Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	{
		"Watch",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "Watch"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic:
	// - Create(                     obj *unstructured.Unstructured, options metav1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error)
	// + Create(ctx context.Context, obj *unstructured.Unstructured, options metav1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error)
	{
		"Create",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "Create"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic:
	// - Update(                     obj *unstructured.Unstructured, options metav1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error)
	// + Update(ctx context.Context, obj *unstructured.Unstructured, options metav1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error)
	{
		"Update",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "Update"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic:
	// - UpdateStatus(                     obj *unstructured.Unstructured, options metav1.UpdateOptions) (*unstructured.Unstructured, error)
	// + UpdateStatus(ctx context.Context, obj *unstructured.Unstructured, options metav1.UpdateOptions) (*unstructured.Unstructured, error)
	{
		"UpdateStatus",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "UpdateStatus"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic:
	// - Patch(                     name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*unstructured.Unstructured, error)
	// + Patch(ctx context.Context, name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*unstructured.Unstructured, error)
	// metadata:
	// - Patch(                     name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*metav1.PartialObjectMetadata, error)
	// + Patch(ctx context.Context, name string, pt types.PatchType, data []byte, options metav1.PatchOptions, subresources ...string) (*metav1.PartialObjectMetadata, error)
	{
		"Patch",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "Patch"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
		},
	},
	// dynamic client:
	// - Delete(                     name string, options metav1.DeleteOptions, subresources ...string) error
	// + Delete(ctx context.Context, name string, options metav1.DeleteOptions, subresources ...string) error
	// metadata client:
	// - Delete(                     name string, options *metav1.DeleteOptions, subresources ...string) error
	// + Delete(ctx context.Context, name string, options metav1.DeleteOptions, subresources ...string) error
	{
		"Delete",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "Delete"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			dereferenceArgAtIndexIfPointer(2, "*k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions"),
			replaceArgAtIndexIfNil(2, "k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions", makeMetav1OptionsArg("DeleteOptions")),
		},
	},
	// dynamic client:
	// - DeleteCollection(                     options metav1.DeleteOptions, listOptions metav1.ListOptions) error
	// + DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error
	// metadata client:
	// - DeleteCollection(                     options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	// + DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error
	{
		"DeleteCollection",
		matchTypePrefixAndMethod(dynamicClientPrefixes, "DeleteCollection"),
		Transforms{
			ensureArgAtIndex(0, "context.Context", makeContextArg),
			dereferenceArgAtIndexIfPointer(1, "*k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions"),
			replaceArgAtIndexIfNil(1, "k8s.io/apimachinery/pkg/apis/meta/v1.DeleteOptions", makeMetav1OptionsArg("DeleteOptions")),
		},
	},
}
