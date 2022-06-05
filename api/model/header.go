package model

// Headers represents the model for header params
// @HeaderParameters Headers
type Headers struct {
	Authorization string `parse:"Authorization" json:"Authorization" example:"Bearer thisisatest" skip:"true"`
}
