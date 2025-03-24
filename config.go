package baseapi

import "fmt"

var ErrFailedToImportRoutes = fmt.Errorf("failed to import the routes JSON file. check the logs for more details")
var ErrFailedToImportCodes = fmt.Errorf("failed to import the codes JSON file. check the logs for more details")
var ErrNoIndexRoute = fmt.Errorf("there must be a index route with the GET method at the routes JSON file")
var ErrNoRequiredCode = fmt.Errorf("a required application code is not set at the codes JSON file")

var requiredCodes = []string{"OK", "I001", "I002", "I003"}
