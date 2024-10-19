package baseapi

import "io"

// resource function name into a application func map
type Methods map[string]func(r *Request) (interface{}, string)

// required interfaces to deal with API items
type API struct {

	// logging
	Writer      io.Writer
	DebugWriter io.Writer

	// API handling
	methods Methods
	codes   map[string]Code
	routes  map[string]map[string]Resource

	//
}

// !! pull the config files, perform validations and return an API interface
func NewAPI(routes, codes string, methods Methods) (api API, err error) {
	return
}
