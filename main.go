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
	k8 "gok8/api"
	handler "gok8/api/handler"
	_ "gok8/api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Starting API server...")
	// k8Auth()
	k8.K8AuthExternal()

	router := mux.NewRouter().StrictSlash(true)
	registerV1Routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerV1Routes(r *mux.Router) {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swagger/").Handler(sh)
	r.HandleFunc("/namespace", handler.ListNamespaces).Methods("GET")
	r.HandleFunc("/namespace/{namespace}", handler.CreateNamespace).Methods("POST")
}
