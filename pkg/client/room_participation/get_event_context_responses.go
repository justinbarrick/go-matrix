// Code generated by go-swagger; DO NOT EDIT.

package room_participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// GetEventContextReader is a Reader for the GetEventContext structure.
type GetEventContextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventContextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventContextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventContextOK creates a GetEventContextOK with default headers values
func NewGetEventContextOK() *GetEventContextOK {
	return &GetEventContextOK{}
}

/*GetEventContextOK handles this case with default header values.

The events and state surrounding the requested event.
*/
type GetEventContextOK struct {
	Payload *models.GetEventContextOKBody
}

func (o *GetEventContextOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/client/unstable/rooms/{roomId}/context/{eventId}][%d] getEventContextOK  %+v", 200, o.Payload)
}

func (o *GetEventContextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetEventContextOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
