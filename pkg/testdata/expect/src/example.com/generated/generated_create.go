package generated

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	samplev1a1 "example.com/apis/samplecontroller/v1alpha1"
	"example.com/clientset/versioned"
)

func Create_117(c kubernetes.Interface, customClient versioned.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(context.TODO(), pod, metav1.CreateOptions{})
	c.CoreV1().Pods("").Create(context.TODO(), &corev1.Pod{}, metav1.CreateOptions{})

	foo := &samplev1a1.Foo{}
	customClient.SamplecontrollerV1alpha1().Foos("").Create(context.TODO(), foo, metav1.CreateOptions{})
	customClient.SamplecontrollerV1alpha1().Foos("").Create(context.TODO(), &samplev1a1.Foo{}, metav1.CreateOptions{})
}

func Create_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	opts := metav1.CreateOptions{}
	c.CoreV1().Pods("").Create(ctx, pod, opts)
	c.CoreV1().Pods("").Create(context.TODO(), &corev1.Pod{}, metav1.CreateOptions{})

	foo := &samplev1a1.Foo{}
	customClient.SamplecontrollerV1alpha1().Foos("").Create(ctx, foo, opts)
	customClient.SamplecontrollerV1alpha1().Foos("").Create(context.TODO(), &samplev1a1.Foo{}, metav1.CreateOptions{})
}

func Create_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(ctx, pod, metav1.CreateOptions{})

	opts := metav1.CreateOptions{}
	c.CoreV1().Pods("").Create(context.TODO(), pod, opts)

	c.CoreV1().Pods("").Create(context.TODO(), pod, metav1.CreateOptions{})
	c.CoreV1().Pods("").Create(context.TODO(), pod, metav1.CreateOptions{})
	c.CoreV1().Pods("").Create(context.TODO(), pod, metav1.CreateOptions{})
}
