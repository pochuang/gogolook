package model

// Response is the response model return from our api.
type Response struct {
	Result interface{} `json:"result"`
}
