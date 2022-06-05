package model

type namespaceQuery struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type namespaceResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
