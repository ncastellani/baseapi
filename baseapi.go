package baseapi

import (
	"io"
	"log"

	"github.com/ncastellani/baseutils"
)

// resource function name into a application func map
type Methods map[string]func(r *Request) (interface{}, string)

// required interfaces to deal with API items
type API struct {
	writer  io.Writer
	methods Methods
	codes   map[string]Code
	routes  map[string]map[string]Resource
}

// pull the config files, perform validations and return an API interface
func NewAPI(routes, codes string, methods Methods, writer io.Writer) (api API, err error) {

	api.writer = writer
	api.methods = methods

	// setup a debug logger
	logger := log.New(writer, "", log.LstdFlags|log.Lmsgprefix)

	logger.Println("setting up a new API handler...")

	// load the API routes to the application
	logger.Println("importing routes JSON file from config...")

	err = baseutils.ParseJSONFile(routes, &api.routes)
	if err != nil {
		logger.Printf("failed to import routes JSON file [err: %v]", err)
		return API{}, ErrFailedToImportRoutes
	}

	// import the API codes
	logger.Println("importing codes JSON file from config...")

	err = baseutils.ParseJSONFile(codes, &api.codes)
	if err != nil {
		logger.Printf("failed to import codes JSON file [err: %v]", err)
		return API{}, ErrFailedToImportCodes
	}

	logger.Println("configuration file parsed and imported! validating minimum requirements...")

	// check if there is a index in the GET method route
	if v, ok := api.routes["index"]; !ok {
		return API{}, ErrNoIndexRoute
	} else {
		if _, ok := v["GET"]; !ok {
			return API{}, ErrNoIndexRoute
		}
	}

	// check for the codes used at the at lib
	for _, code := range requiredCodes {
		if _, ok := api.codes[code]; !ok {
			logger.Printf("required application code does not exist [code: %v]", code)
			return API{}, ErrNoRequiredCode
		}
	}

	return
}
