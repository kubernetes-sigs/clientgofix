package generated

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Create_117(c kubernetes.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(pod)
	c.CoreV1().Pods("").Create(&corev1.Pod{})
}

func Create_118(c kubernetes.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	opts := metav1.CreateOptions{}
	c.CoreV1().Pods("").Create(ctx, pod, opts)
	c.CoreV1().Pods("").Create(context.TODO(), &corev1.Pod{}, metav1.CreateOptions{})
}

func Create_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Create(ctx, pod)

	opts := metav1.CreateOptions{}
	c.CoreV1().Pods("").Create(pod, opts)

	c.CoreV1().Pods("").Create(pod)
	c.CoreV1().Pods("").Create(pod, metav1.CreateOptions{})
	c.CoreV1().Pods("").Create(context.TODO(), pod)
}
