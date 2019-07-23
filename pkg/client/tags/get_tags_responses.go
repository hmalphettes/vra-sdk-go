// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetTagsReader is a Reader for the GetTags structure.
type GetTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTagsOK creates a GetTagsOK with default headers values
func NewGetTagsOK() *GetTagsOK {
	return &GetTagsOK{}
}

/*GetTagsOK handles this case with default header values.

successful operation
*/
type GetTagsOK struct {
	Payload *models.TagResult
}

func (o *GetTagsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/tags][%d] getTagsOK  %+v", 200, o.Payload)
}

func (o *GetTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagsForbidden creates a GetTagsForbidden with default headers values
func NewGetTagsForbidden() *GetTagsForbidden {
	return &GetTagsForbidden{}
}

/*GetTagsForbidden handles this case with default header values.

Forbidden
*/
type GetTagsForbidden struct {
}

func (o *GetTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/tags][%d] getTagsForbidden ", 403)
}

func (o *GetTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
