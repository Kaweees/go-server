package main

// Represents the response types of GetExample()
type GetExampleResponse struct {
	Result string `json:"value"`
}


// Represets the response type of PostExample()
type PostExampleResponse struct {
	Value bool `json:"value"`
	AnotherValue string `json:"another_value"`
}
