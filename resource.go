package baseapi

import "log"

// define an operation result code, its messages by language and its HTTP code
type Code struct {
	HTTPCode int               `json:"status"`  // HTTP return code
	Message  map[string]string `json:"message"` // messages from the code
}

// define an incoming request, with its metadata, payload and useful contents.
type Request struct {
	logger *log.Logger
	api    *API

	// general request data
	ID      string
	IP      string
	Headers map[string]string
	Query   map[string]string
	Path    string
	Method  string
	Input   []byte

	// asserted data
	Resource   Resource
	Parameters *map[string]interface{}

	// method response
	ResultData interface{}
	ResultCode string
}

// define an API method within a route
type Resource struct {
	ResourceMethod string              `json:"function"`
	Parameters     []ResourceParameter `json:"parameters"` // acceptable parameters for this action
}

// define an parameter specification for the resource
type ResourceParameter struct {
	Name      string   `json:"name"`       // parameter name
	Kind      string   `json:"kind"`       // parameter type (string/number/enum)
	Required  bool     `json:"required"`   // is required
	Options   []string `json:"options"`    // if type ENUM, this is a list of the available options
	MaxLength int      `json:"max_length"` // !! if type STRING, validate its length
}
