package handlers

import "github.com/labstack/echo"

// RawResponse takes the response from handlers and returns the response is given format
func RawResponse(c echo.Context, response interface{}, httpCode int) error {
	var responseFunc func(int, interface{}) error
	switch c.Request().Header.Get("accept") {
	case "application/json", "text/json", "json":
		responseFunc = c.JSON
	case "text/xml", "application/xml", "xml":
		responseFunc = c.XML
	default:
		responseFunc = c.JSON
	}
	//metrics.Request(c.Get("RequestID").(string), c.Request().URL.Path, c.Request().Method, httpCode)
	return responseFunc(httpCode, response)
}

func Init() error {
	return nil
}
