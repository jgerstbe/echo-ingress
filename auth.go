package main

import (
    "net/http"
	"log"
    "github.com/labstack/echo/v4"
    "github.com/patrickmn/go-cache"
)

func validateRequestMiddleware(ccache cache.Cache) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            // Get the request method, path, and bearer token from the request
            method := c.Request().Method
            path := c.Request().URL.Path
            token := c.Request().Header.Get("Authorization")

            // Match a route from cached config
            hasMatchedRoute, matchedRoute := matchConfigRoute(ccache, path, method)
            if !hasMatchedRoute {
                return c.String(http.StatusNotFound, "No matching route configured.")
            }

            c.Set("matchedRoute", matchedRoute)

            if matchedRoute.LocalAuth == "none" {
                return next(c)
            }

            // Make an HTTP request to an external service
            // Replace this with your actual validation logic
            if validateRequestWithExternalService(method, path, token) {
                return next(c) // Proceed with the request handling
            }

            // If the request is not authorized, return a 401 Unauthorized response
            return c.String(http.StatusUnauthorized, "Unauthorized")
        }
    }
}

// Replace this with your actual validation logic
func validateRequestWithExternalService(method, path, token string) bool {
    // Return true if the request is authorized, or false if it's not
    log.Println(method, path, token)

    // Valildate against external auth service
    response, err := http.Get("http://httpbin.org/status/200%2C401")
	if err != nil {
		return false
	}
	if response.StatusCode != http.StatusOK {
        return false
    }
	defer response.Body.Close()

    return true
}

func matchConfigRoute(c cache.Cache, path string, method string) (bool, Route) {
    var matchedRoute Route
    routeKey := "ROUTE_"+method+"_"+path
    cachedRoute, found := c.Get(routeKey)

    if found {
        matchedRoute = cachedRoute.(Route)
        log.Printf("Match found: Route Name: %s, Path: %s\n", matchedRoute.Name, matchedRoute.Path)
    } else {
        log.Println("No match found.")
    }

	return found, matchedRoute
}