// package comment
package misc

// before imports comment
import (
	// top of import block
	_ "archive/tar"

	// between import groups

	corev1 "k8s.io/api/core/v1" // corev1
	// between corev1/kubernetes
	kubernetes "k8s.io/client-go/kubernetes" //kubernetes
)

// after imports

// CommentFuncA godoc
func CommentFuncA( /* before c */ c kubernetes.Interface /* after c */) { // CommentFuncA line
	// before pod
	pod := &corev1.Pod{} // pod
	// after pod
	c.CoreV1().Pods("").Update( /* before pod param*/ pod /* after pod param*/) // Update line
	// after Update line
	c.CoreV1().Pods("").Update(
		/* before pod param*/ pod, /* after pod param*/
	) // Update line

	/* block comment before function end*/
}
