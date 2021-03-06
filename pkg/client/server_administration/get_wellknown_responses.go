// Code generated by go-swagger; DO NOT EDIT.

package server_administration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// GetWellknownReader is a Reader for the GetWellknown structure.
type GetWellknownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWellknownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWellknownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetWellknownNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWellknownOK creates a GetWellknownOK with default headers values
func NewGetWellknownOK() *GetWellknownOK {
	return &GetWellknownOK{}
}

/*GetWellknownOK handles this case with default header values.

Server discovery information.
*/
type GetWellknownOK struct {
	Payload *models.GetWellknownOKBody
}

func (o *GetWellknownOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/matrix/client][%d] getWellknownOK  %+v", 200, o.Payload)
}

func (o *GetWellknownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetWellknownOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWellknownNotFound creates a GetWellknownNotFound with default headers values
func NewGetWellknownNotFound() *GetWellknownNotFound {
	return &GetWellknownNotFound{}
}

/*GetWellknownNotFound handles this case with default header values.

No server discovery information available.
*/
type GetWellknownNotFound struct {
}

func (o *GetWellknownNotFound) Error() string {
	return fmt.Sprintf("[GET /.well-known/matrix/client][%d] getWellknownNotFound ", 404)
}

func (o *GetWellknownNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
