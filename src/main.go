package main

import (
	"bytes"
	"context"
	"io"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	var podName []string
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	pods, _ := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	podLog := make([]*rest.Request, len(nodes.Items))
	log := make([]string, len(nodes.Items))
	line := int64(len(nodes.Items) - 1)
	pingRegex := regexp.MustCompile(`monlat-agent-`)
	for _, pod := range pods.Items {
		if pingRegex.MatchString(pod.Name) {
			podName = append(podName, pod.Name)
		}
	}
	for i := 1; i <= len(nodes.Items); i++ {
		podLog[i-1] = clientset.CoreV1().Pods("default").GetLogs(podName[i-1], &corev1.PodLogOptions{
			TailLines: &line,
		})
	}
	e := echo.New()
	e.GET("/metrics", func(c echo.Context) error {
		for i := 0; i < len(nodes.Items); i++ {
			podLogs, err := podLog[i].Stream(context.TODO())
			if err != nil {
				panic(err.Error())
			}
			defer podLogs.Close()
			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, podLogs)
			if err != nil {
				panic(err.Error())
			}
			str := buf.String()
			lines := strings.Split(str, "\n")
			log[i] = strings.Join(lines, "\n")
		}
		laslog := strings.Join(log, "")
		return c.String(http.StatusOK, laslog)
	})
	e.Logger.Fatal(e.Start(":9090"))
}
