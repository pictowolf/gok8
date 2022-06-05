package handler

// @Tag namespace
// @Title List namespaces
// @Description Get a list of namespaces.
// @Header model.Headers
// @Param request body model.namespaceQuery true "The query to filter namespaces"
// @Success 200 {object} model.namespaceResponse "The response from the kubernetes cluster"
// @Router /namespace [get]
func listNamespaces() {
	return
}

// @Tag namespace
// @Title Get namespace
// @Description Get a namespace
// @Param namespace path string true "The namespace name"
// @Success 200 {object} model.namespaceResponse "The response from the kubernetes cluster"
// @Router /namespace/{namespace} [get]
func getNamespace() {
	return
}

// @Tag namespace
// @Title Create namespace
// @Description Create a namespace
// @Param request body model.namespaceRequest true "The request to create a namespace"
// @Success 200 {object} model.namespaceResponse "The response from the kubernetes cluster"
// @Router /namespace [post]
func createNamespace() {
	return
}

// @Tag namespace
// @Title Update namespace
// @Description Update a namespace
// @Param namespace path string true "The namespace name"
// @Param request body model.namespaceRequest true "The request to update a namespace"
// @Success 200 {object} model.namespaceResponse "The response from the kubernetes cluster"
// @Router /namespace/{namespace} [put]
func updateNamespace() {
	return
}

// @Tag namespace
// @Title Delete namespace
// @Description Delete a namespace
// @Param namespace path string true "The namespace name"
// @Success 200 {object} model.namespaceResponse "The response from the kubernetes cluster"
// @Router /namespace/{namespace} [delete]
func deleteNamespace() {
	return
}
