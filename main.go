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
	"fmt"
	handler "gok8/api/handler"
	_ "gok8/api/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	println("Starting API server...")
	// k8Auth()

	router := mux.NewRouter().StrictSlash(true)
	registerV1Routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerV1Routes(r *mux.Router) {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swagger/").Handler(sh)
	r.HandleFunc("/namespace", handler.ListNamespaces).Methods("GET")
}

func k8Auth() {
	println("Starting k8 auth in cluster")
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

		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		_, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod example-xxxxx not found in default namespace\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found example-xxxxx pod in default namespace\n")
		}

		time.Sleep(10 * time.Second)
	}
}
