package baseapi

import (
	"io"
	"log"

	"github.com/ncastellani/baseutils"
)

// required interfaces to deal with API items
type API struct {
	writer   io.Writer
	methods  Methods
	hostData []string
	codes    map[string]Code
	routes   map[string]map[string]Resource

	// application middlewares
	PreRequest  func(r *Request)
	PostRequest func(r *Request)
}

// pull the config files, perform validations and return an API interface
func NewAPI(routes, codes string, methods Methods, writer io.Writer, hostData []string) (api API, err error) {

	api.writer = writer
	api.methods = methods
	api.hostData = hostData

	// setup a debug logger
	l := log.New(writer, "", log.LstdFlags|log.Lmsgprefix)

	l.Println("setting up a new API handler...")

	// load the API routes to the application
	l.Println("importing routes JSON file from the passed path...")

	err = baseutils.ParseJSONFile(routes, &api.routes)
	if err != nil {
		l.Printf("failed to import routes JSON file [err: %v]", err)

		err = ErrFailedToImportRoutes
		return
	}

	// import the API codes
	l.Println("importing codes JSON file from the passed path...")

	err = baseutils.ParseJSONFile(codes, &api.codes)
	if err != nil {
		l.Printf("failed to import codes JSON file [err: %v]", err)

		err = ErrFailedToImportCodes
		return
	}

	l.Println("configuration file parsed and imported! validating minimum requirements...")

	// check if there is a index route and if it has a GET method
	if v, ok := api.routes["index"]; !ok {
		l.Println("no index route")

		err = ErrNoIndexRoute
		return
	} else {
		if _, ok := v["GET"]; !ok {
			l.Println("no index route")

			err = ErrNoIndexRoute
			return
		}
	}

	// check for the codes used at the at lib
	for _, code := range requiredCodes {
		if _, ok := api.codes[code]; !ok {
			l.Printf("a required application code does not exist [code: %v]", code)

			err = ErrNoRequiredCode
			return
		}
	}

	l.Println("required index route and codes are available!")

	// set defaults pre and post request middlewares
	api.PreRequest = func(r *Request) {}
	api.PostRequest = func(r *Request) {}

	l.Println("successfully setted up this API handler!")

	return
}
