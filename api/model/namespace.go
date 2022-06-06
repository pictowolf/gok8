package model

type NamespaceQuery struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type NamespaceResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type NamespaceRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
