package generated

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Update_117(c kubernetes.Interface) {
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Update(context.TODO(), pod, metav1.UpdateOptions{})
	c.CoreV1().Pods("").Update(context.TODO(), &corev1.Pod{}, metav1.UpdateOptions{})
}

func Update_118(c kubernetes.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	opts := metav1.UpdateOptions{}
	c.CoreV1().Pods("").Update(ctx, pod, opts)
	c.CoreV1().Pods("").Update(context.TODO(), &corev1.Pod{}, metav1.UpdateOptions{})
}

func Update_Other(c kubernetes.Interface) {
	ctx := context.TODO()
	pod := &corev1.Pod{}
	c.CoreV1().Pods("").Update(ctx, pod, metav1.UpdateOptions{})

	opts := metav1.UpdateOptions{}
	c.CoreV1().Pods("").Update(context.TODO(), pod, opts)

	c.CoreV1().Pods("").Update(context.TODO(), pod, metav1.UpdateOptions{})
	c.CoreV1().Pods("").Update(context.TODO(), pod, metav1.UpdateOptions{})
	c.CoreV1().Pods("").Update(context.TODO(), pod, metav1.UpdateOptions{})
}
