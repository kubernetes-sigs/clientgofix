package generated

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"

	"example.com/clientset/versioned"
)

func UpdateStatus_117(c kubernetes.Interface, customClient versioned.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").UpdateStatus(context.TODO(), pod, metav1.UpdateOptions{})
	c.CoreV1().Pods("").UpdateStatus(context.TODO(), &corev1.Pod{}, metav1.UpdateOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").UpdateStatus(context.TODO(), pod, metav1.UpdateOptions{})
	customClient.SamplecontrollerV1alpha1().Foos("").UpdateStatus(context.TODO(), &corev1.Pod{}, metav1.UpdateOptions{})
}

func UpdateStatus_118(c kubernetes.Interface, customClient versioned.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	opts := metav1.UpdateOptions{}
	c.CoreV1().Pods("").UpdateStatus(ctx, pod, opts)
	c.CoreV1().Pods("").UpdateStatus(context.TODO(), &corev1.Pod{}, metav1.UpdateOptions{})

	customClient.SamplecontrollerV1alpha1().Foos("").UpdateStatus(ctx, pod, opts)
	customClient.SamplecontrollerV1alpha1().Foos("").UpdateStatus(context.TODO(), &corev1.Pod{}, metav1.UpdateOptions{})
}

func UpdateStatus_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").UpdateStatus(ctx, pod, metav1.UpdateOptions{})

	opts := metav1.UpdateOptions{}
	c.CoreV1().Pods("").UpdateStatus(context.TODO(), pod, opts)

	c.CoreV1().Pods("").UpdateStatus(context.TODO(), pod, metav1.UpdateOptions{})
	c.CoreV1().Pods("").UpdateStatus(context.TODO(), pod, metav1.UpdateOptions{})
	c.CoreV1().Pods("").UpdateStatus(context.TODO(), pod, metav1.UpdateOptions{})
}
