package exceptions

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	RawBody  []byte `json:"body"`
	ErrorMsg string `json:"error"`
	ModelData any    `json:"model"`
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.ErrorMsg
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.RawBody
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() any {
	return e.ModelData
}
