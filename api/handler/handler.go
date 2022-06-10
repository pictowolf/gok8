package handler

import (
	"context"
	"encoding/json"
	"fmt"
	k8 "gok8/api"
	_ "gok8/api/model"
	"net/http"

	"github.com/gorilla/mux"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// @Tag namespace
// @Title List namespaces
// @Description Get a list of namespaces.
// @Success 200 {object} model.NamespaceResponse "The response from the kubernetes cluster"
// @Router /namespace [get]
func ListNamespaces(w http.ResponseWriter, r *http.Request) {
	namespaces, err := k8.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	namespaceSlice := []string{}
	for _, namespace := range namespaces.Items {
		namespaceSlice = append(namespaceSlice, namespace.Name)
	}

	json.NewEncoder(w).Encode(namespaceSlice)
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
// @Param namespace path string true "The namespace name"
// @Success 201 "Namespace created"
// @Failure 409 "Already exists"
// @Router /namespace/{namespace} [post]
func CreateNamespace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := k8.Clientset.CoreV1().Namespaces().Create(context.TODO(), &core.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: vars["namespace"],
		},
	}, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("namespace : " + vars["namespace"] + " already exists")
	} else if err != nil {
		fmt.Println(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("namespace : " + vars["namespace"] + " created")
	}
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
