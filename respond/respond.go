package respond

import (
	"net/http"

	"github.com/go-chi/render"
)

//Response holds the API response
type Response struct {
	Data           interface{} `json:"data,omitempty"`
	HTTPStatusCode int         `json:"-"` // http response status code
}

// Render is used to implement ErrResponse to Renderer
func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, resp.HTTPStatusCode)
	return nil
}

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

// Render is used to implement ErrResponse to Renderer
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrBadRequest is sent when the request is not correct
func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// OK is returned in a success response
func OK(data interface{}) render.Renderer {
	return &Response{
		Data:           data,
		HTTPStatusCode: 200,
	}
}

// ErrInternalServer is sent when the request is not correct
func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Something went wrong",
		ErrorText:      err.Error(),
	}
}

// ErrNotFound is returned when a resource is not found
var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
