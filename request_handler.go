package main

type RequestHandler struct {}

func NewRequestHandler() (*RequestHandler, error) {
	requestHandler := &RequestHandler{}
	err := requestHandler.init()
	if err != nil {
		return nil, err
	}
	return requestHandler, nil
}

func (rh *RequestHandler) init() error {
	// Initialize the request handler
	return nil
}

func (rh *RequestHandler) GetExample() (*GetExampleResponse, error) {
	return &GetExampleResponse{
		Result: "Hello, world!",
	}, nil
}

func (rh *RequestHandler) PostExample(Value bool, AnotherValue string) (*PostExampleResponse, error) {
	return &PostExampleResponse{
		Value:        Value,
		AnotherValue: AnotherValue,
	}, nil
}
