// Code generated by go-swagger; DO NOT EDIT.

package request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetRequestTrackersReader is a Reader for the GetRequestTrackers structure.
type GetRequestTrackersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestTrackersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRequestTrackersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetRequestTrackersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRequestTrackersOK creates a GetRequestTrackersOK with default headers values
func NewGetRequestTrackersOK() *GetRequestTrackersOK {
	return &GetRequestTrackersOK{}
}

/*GetRequestTrackersOK handles this case with default header values.

successful operation
*/
type GetRequestTrackersOK struct {
	Payload *models.RequestTrackerResult
}

func (o *GetRequestTrackersOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker][%d] getRequestTrackersOK  %+v", 200, o.Payload)
}

func (o *GetRequestTrackersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTrackerResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestTrackersForbidden creates a GetRequestTrackersForbidden with default headers values
func NewGetRequestTrackersForbidden() *GetRequestTrackersForbidden {
	return &GetRequestTrackersForbidden{}
}

/*GetRequestTrackersForbidden handles this case with default header values.

Forbidden
*/
type GetRequestTrackersForbidden struct {
}

func (o *GetRequestTrackersForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker][%d] getRequestTrackersForbidden ", 403)
}

func (o *GetRequestTrackersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}