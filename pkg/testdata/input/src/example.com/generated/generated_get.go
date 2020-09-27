package generated

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"example.com/clientset/versioned"
)

func Get_117(c kubernetes.Interface, customClient versioned.Interface) {
	opts := metav1.GetOptions{}
	c.CoreV1().Pods("").Get("", opts)
	c.CoreV1().Pods("").Get("", metav1.GetOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").Get("", opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Get("", metav1.GetOptions{})
}

func Get_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	opts := metav1.GetOptions{}
	c.CoreV1().Pods("").Get(ctx, "", opts)
	c.CoreV1().Pods("").Get(context.TODO(), "", metav1.GetOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").Get(ctx, "", opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Get(context.TODO(), "", metav1.GetOptions{})
}

func Get_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	c.CoreV1().Pods("").Get(ctx, "")

	opts := metav1.GetOptions{}
	c.CoreV1().Pods("").Get("", opts)

	c.CoreV1().Pods("").Get("")
	c.CoreV1().Pods("").Get("", metav1.GetOptions{})
	c.CoreV1().Pods("").Get(context.TODO(), "")
}
