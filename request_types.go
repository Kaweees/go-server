package main

// Represets the request type of PostExample()
type PostExampleRequest struct {
	Value bool `json:"value"`
	AnotherValue string `json:"another_value"`
}
