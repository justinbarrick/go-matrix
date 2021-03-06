// Code generated by go-swagger; DO NOT EDIT.

package media

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// GetContentThumbnailReader is a Reader for the GetContentThumbnail structure.
type GetContentThumbnailReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetContentThumbnailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContentThumbnailOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 429:
		result := NewGetContentThumbnailTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContentThumbnailOK creates a GetContentThumbnailOK with default headers values
func NewGetContentThumbnailOK(writer io.Writer) *GetContentThumbnailOK {
	return &GetContentThumbnailOK{
		Payload: writer,
	}
}

/*GetContentThumbnailOK handles this case with default header values.

A thumbnail of the requested content.
*/
type GetContentThumbnailOK struct {
	/*The content type of the thumbnail.
	 */
	ContentType string

	Payload io.Writer
}

func (o *GetContentThumbnailOK) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailOK  %+v", 200, o.Payload)
}

func (o *GetContentThumbnailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentThumbnailTooManyRequests creates a GetContentThumbnailTooManyRequests with default headers values
func NewGetContentThumbnailTooManyRequests() *GetContentThumbnailTooManyRequests {
	return &GetContentThumbnailTooManyRequests{}
}

/*GetContentThumbnailTooManyRequests handles this case with default header values.

This request was rate-limited.
*/
type GetContentThumbnailTooManyRequests struct {
	Payload *models.GetContentThumbnailTooManyRequestsBody
}

func (o *GetContentThumbnailTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /_matrix/media/unstable/thumbnail/{serverName}/{mediaId}][%d] getContentThumbnailTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentThumbnailTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetContentThumbnailTooManyRequestsBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
