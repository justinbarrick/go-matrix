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

// SendMessageReader is a Reader for the SendMessage structure.
type SendMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendMessageOK creates a SendMessageOK with default headers values
func NewSendMessageOK() *SendMessageOK {
	return &SendMessageOK{}
}

/*SendMessageOK handles this case with default header values.

An ID for the sent event.
*/
type SendMessageOK struct {
	Payload *models.SendMessageOKBody
}

func (o *SendMessageOK) Error() string {
	return fmt.Sprintf("[PUT /_matrix/client/unstable/rooms/{roomId}/send/{eventType}/{txnId}][%d] sendMessageOK  %+v", 200, o.Payload)
}

func (o *SendMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SendMessageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
