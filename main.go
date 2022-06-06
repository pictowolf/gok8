// @Version 1.0.0
// @Title Infinity Servers API
// @Description API for Infinity Servers panel for interactions to kubernetes.
// @ContactName Infinity Servers
// @ContactEmail shaun@infinityservers.co.uk
// @ContactURL https://www.infinityservers.co.uk
// @TermsOfServiceUrl https://www.infinityservers.co.uk/terms
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader APIKey header X-API-KEY

package main

import (
	"context"
	"flag"
	"fmt"
	handler "gok8/api/handler"
	_ "gok8/api/model"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	println("Starting API server...")
	// k8Auth()
	k8AuthExternal()
	router := mux.NewRouter().StrictSlash(true)
	registerV1Routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerV1Routes(r *mux.Router) {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swagger/").Handler(sh)
	r.HandleFunc("/namespace", handler.ListNamespaces).Methods("GET")
}

// This should be used when API is hosted in K8 itself.
func k8Auth() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	}
}

// This should be used for local dev, when API is not hosted in K8.
func k8AuthExternal() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	}
}
