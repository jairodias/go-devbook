package routes

import (
	"api/src/controllers"
	"net/http"
)

var authenticateRoute = Route{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controllers.Authenticate,
	RequiresAuthentication: false,
}
