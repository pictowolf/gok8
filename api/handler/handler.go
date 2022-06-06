package handler

import (
	"context"
	"encoding/json"
	"fmt"
	k8 "gok8/api"
	_ "gok8/api/model"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Tag namespace
// @Title List namespaces
// @Description Get a list of namespaces.
// @Success 200 {object} model.NamespaceResponse "The response from the kubernetes cluster"
// @Router /namespace [get]
func ListNamespaces(w http.ResponseWriter, r *http.Request) {
	// list of namespaces
	namespaces, err := k8.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// make a map
	namespaceMap := make(map[string]string)
	for _, namespace := range namespaces.Items {
		namespaceMap[namespace.Name] = namespace.Name
	}

	fmt.Println(namespaceMap)

	// encode into json
	json.NewEncoder(w).Encode(namespaceMap)
}

// @Tag namespace
// @Title Get namespace
// @Description Get a namespace
// @Param namespace path string true "The namespace name"
// @Success 200 {object} model.NamespaceResponse "The response from the kubernetes cluster"
// @Router /namespace/{namespace} [get]
func getNamespace() {
	return
}

// @Tag namespace
// @Title Create namespace
// @Description Create a namespace
// @Param request body model.NamespaceRequest true "The request to create a namespace"
// @Success 200 {object} model.NamespaceResponse "The response from the kubernetes cluster"
// @Router /namespace [post]
func createNamespace() {
	return
}

// @Tag namespace
// @Title Update namespace
// @Description Update a namespace
// @Param namespace path string true "The namespace name"
// @Param request body model.NamespaceRequest true "The request to update a namespace"
// @Success 200 {object} model.NamespaceResponse "The response from the kubernetes cluster"
// @Router /namespace/{namespace} [put]
func updateNamespace() {
	return
}

// @Tag namespace
// @Title Delete namespace
// @Description Delete a namespace
// @Param namespace path string true "The namespace name"
// @Success 200 {object} model.NamespaceResponse "The response from the kubernetes cluster"
// @Router /namespace/{namespace} [delete]
func deleteNamespace() {
	return
}
