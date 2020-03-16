package misc

import (
	"context"

	authenticationapi "k8s.io/api/authentication/v1"
	certificatesapi "k8s.io/api/certificates/v1beta1"
	coreapi "k8s.io/api/core/v1"
	extensionsapi "k8s.io/api/extensions/v1beta1"
	policyapi "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	certificatesclient "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
	extensionsclient "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

// - UpdateApproval(                     certificateSigningRequest *certificates.CertificateSigningRequest                           ) (result *certificates.CertificateSigningRequest, err error)
func UpdateApproval_117(c certificatesclient.CertificateSigningRequestInterface) {
	c.UpdateApproval(context.TODO(), nil, metav1.UpdateOptions{})
	c.UpdateApproval(context.TODO(), &certificatesapi.CertificateSigningRequest{}, metav1.UpdateOptions{})
}

// + UpdateApproval(ctx context.Context, certificateSigningRequest *certificates.CertificateSigningRequest, opts metav1.UpdateOptions) (result *certificates.CertificateSigningRequest, err error)
func UpdateApproval_118(c certificatesclient.CertificateSigningRequestInterface) {
	c.UpdateApproval(context.TODO(), nil, metav1.UpdateOptions{})
	c.UpdateApproval(context.TODO(), &certificatesapi.CertificateSigningRequest{}, metav1.UpdateOptions{})
}

// - Finalize(                     item *v1.Namespace                           ) (*v1.Namespace, error)
func Finalize_117(c coreclient.NamespaceInterface) {
	c.Finalize(context.TODO(), nil, metav1.UpdateOptions{})
	c.Finalize(context.TODO(), &coreapi.Namespace{}, metav1.UpdateOptions{})
}

// + Finalize(ctx context.Context, item *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error)
func Finalize_118(c coreclient.NamespaceInterface) {
	c.Finalize(context.TODO(), nil, metav1.UpdateOptions{})
	c.Finalize(context.TODO(), &coreapi.Namespace{}, metav1.UpdateOptions{})
}

// - PatchStatus(                     nodeName string, data []byte) (*v1.Node, error)
func PatchStatus_117(c coreclient.NodeInterface) {
	c.PatchStatus(context.TODO(), "", []byte{})
}

// + PatchStatus(ctx context.Context, nodeName string, data []byte) (*v1.Node, error)
func PatchStatus_118(c coreclient.NodeInterface) {
	c.PatchStatus(context.TODO(), "", []byte{})
}

// - Bind(                     binding *v1.Binding                           ) error
func Bind_117(c coreclient.PodInterface) {
	c.Bind(context.TODO(), nil, metav1.CreateOptions{})
	c.Bind(context.TODO(), &coreapi.Binding{}, metav1.CreateOptions{})
}

// + Bind(ctx context.Context, binding *v1.Binding, opts metav1.CreateOptions) error
func Bind_118(c coreclient.PodInterface) {
	c.Bind(context.TODO(), nil, metav1.CreateOptions{})
	c.Bind(context.TODO(), &coreapi.Binding{}, metav1.CreateOptions{})
}

// - Evict(                     eviction *policy.Eviction) error
func Evict_117(c coreclient.PodInterface) {
	c.Evict(context.TODO(), &policyapi.Eviction{})
}

// + Evict(ctx context.Context, eviction *policy.Eviction) error
func Evict_118(c coreclient.PodInterface) {
	c.Evict(context.TODO(), &policyapi.Eviction{})
}

// - CreateToken(                     name string, tokenRequest *authenticationv1.TokenRequest                           ) (*authenticationv1.TokenRequest, error)
func CreateToken_117(c coreclient.ServiceAccountInterface) {
	c.CreateToken(context.TODO(), "", &authenticationapi.TokenRequest{}, metav1.CreateOptions{})
}

// + CreateToken(ctx context.Context, name string, tokenRequest *authenticationv1.TokenRequest, opts metav1.CreateOptions) (*authenticationv1.TokenRequest, error)
func CreateToken_118(c coreclient.ServiceAccountInterface) {
	c.CreateToken(context.TODO(), "", &authenticationapi.TokenRequest{}, metav1.CreateOptions{})
}

// - Rollback(                 *v1beta1.DeploymentRollback                      ) error
func Rollback_117(c extensionsclient.DeploymentInterface) {
	c.Rollback(context.TODO(), &extensionsapi.DeploymentRollback{}, metav1.CreateOptions{})
}

// + Rollback(context.Context, *v1beta1.DeploymentRollback, metav1.CreateOptions) error
func Rollback_118(c extensionsclient.DeploymentInterface) {
	c.Rollback(context.TODO(), &extensionsapi.DeploymentRollback{}, metav1.CreateOptions{})
}
