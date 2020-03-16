package misc

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
)

func Imports_Add_Grouped(c kubernetes.Interface) {
	c.CoreV1().Pods("").Create(context.TODO(), nil, metav1.CreateOptions{})
}
