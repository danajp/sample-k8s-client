package main

import (
	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	"context"
	"fmt"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	argov1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

var scheme = runtime.NewScheme()

const namespace = "argocd"

func main() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(argov1alpha1.AddToScheme(scheme))

	c, _ := client.New(config.GetConfigOrDie(), client.Options{
		Scheme: scheme,
	})

	apps := argov1alpha1.ApplicationList{}

	if err := c.List(context.TODO(), &apps, &client.ListOptions{Namespace: namespace}); err != nil {
		fmt.Printf("error listing applications: %s\n", err)
		return
	}

	for _, app := range(apps.Items) {
		fmt.Printf("%s %s/%s\n", app.Kind, app.Namespace, app.Name)
	}
}
