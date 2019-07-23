// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// CreateAwsCloudAccountReader is a Reader for the CreateAwsCloudAccount structure.
type CreateAwsCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAwsCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAwsCloudAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAwsCloudAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateAwsCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAwsCloudAccountCreated creates a CreateAwsCloudAccountCreated with default headers values
func NewCreateAwsCloudAccountCreated() *CreateAwsCloudAccountCreated {
	return &CreateAwsCloudAccountCreated{}
}

/*CreateAwsCloudAccountCreated handles this case with default header values.

successful operation
*/
type CreateAwsCloudAccountCreated struct {
	Payload *models.CloudAccountAws
}

func (o *CreateAwsCloudAccountCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws][%d] createAwsCloudAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateAwsCloudAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountAws)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAwsCloudAccountBadRequest creates a CreateAwsCloudAccountBadRequest with default headers values
func NewCreateAwsCloudAccountBadRequest() *CreateAwsCloudAccountBadRequest {
	return &CreateAwsCloudAccountBadRequest{}
}

/*CreateAwsCloudAccountBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateAwsCloudAccountBadRequest struct {
}

func (o *CreateAwsCloudAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws][%d] createAwsCloudAccountBadRequest ", 400)
}

func (o *CreateAwsCloudAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateAwsCloudAccountForbidden creates a CreateAwsCloudAccountForbidden with default headers values
func NewCreateAwsCloudAccountForbidden() *CreateAwsCloudAccountForbidden {
	return &CreateAwsCloudAccountForbidden{}
}

/*CreateAwsCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type CreateAwsCloudAccountForbidden struct {
}

func (o *CreateAwsCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws][%d] createAwsCloudAccountForbidden ", 403)
}

func (o *CreateAwsCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
