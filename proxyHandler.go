package main

import (
    "net/http"
	"fmt"
	"io/ioutil"
    "github.com/labstack/echo/v4"
)

func handleRequest(c echo.Context) error {
    path := c.Request().URL.Path
    method := c.Request().Method

    // URL to make an HTTP GET request to
    // matchedRoute Route := c.Get("matchedRoute")
    matchedRoute, ok := c.Get("matchedRoute").(Route)
    if !ok {
        return c.String(http.StatusInternalServerError, "Error retrieving custom route")
    }
    url := fmt.Sprint(matchedRoute.UpstreamURL, path)

    // Create a new HTTP request with the specified method
    req, err := http.NewRequest(method, url, nil)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error creating HTTP request: "+err.Error())
    }

    // Make the HTTP request
    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error making HTTP request: "+err.Error())
    }
    defer response.Body.Close()

    // Read the response body
    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error reading response body: "+err.Error())
    }
    
    // Set the response content type
    c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

    // Write the response body to the Echo response
    _, err = c.Response().Write(responseBody)
    if err != nil {
        return c.String(http.StatusInternalServerError, "Error writing response: "+err.Error())
    }

    return nil
}