package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	println("Hello, world!")
	router := mux.NewRouter().StrictSlash(true)
	registerV1Routes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func registerV1Routes(r *mux.Router) {
	sh := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/")))
	r.PathPrefix("/swagger/").Handler(sh)
}

// @Version 1.0.0
// @Title Infinity Servers API
// @Description API for Infinity Servers panel for interactions to kubernetes.
// @ContactName Infinity Servers
// @ContactEmail shaun@infinityservers.co.uk
// @ContactURL https://www.infinityservers.co.uk
// @TermsOfServiceUrl https://www.infinityservers.co.uk/terms
// @LicenseName MIT
// @LicenseURL https://en.wikipedia.org/wiki/MIT_License
// @Server https://www.fake.com Server-1
// @Security AuthorizationHeader read write
// @SecurityScheme AuthorizationHeader http bearer test
