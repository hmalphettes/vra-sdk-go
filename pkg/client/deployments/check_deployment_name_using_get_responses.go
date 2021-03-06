// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CheckDeploymentNameUsingGETReader is a Reader for the CheckDeploymentNameUsingGET structure.
type CheckDeploymentNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckDeploymentNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckDeploymentNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCheckDeploymentNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckDeploymentNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckDeploymentNameUsingGETOK creates a CheckDeploymentNameUsingGETOK with default headers values
func NewCheckDeploymentNameUsingGETOK() *CheckDeploymentNameUsingGETOK {
	return &CheckDeploymentNameUsingGETOK{}
}

/*CheckDeploymentNameUsingGETOK handles this case with default header values.

OK
*/
type CheckDeploymentNameUsingGETOK struct {
}

func (o *CheckDeploymentNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/names/{name}][%d] checkDeploymentNameUsingGETOK ", 200)
}

func (o *CheckDeploymentNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckDeploymentNameUsingGETUnauthorized creates a CheckDeploymentNameUsingGETUnauthorized with default headers values
func NewCheckDeploymentNameUsingGETUnauthorized() *CheckDeploymentNameUsingGETUnauthorized {
	return &CheckDeploymentNameUsingGETUnauthorized{}
}

/*CheckDeploymentNameUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type CheckDeploymentNameUsingGETUnauthorized struct {
}

func (o *CheckDeploymentNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/names/{name}][%d] checkDeploymentNameUsingGETUnauthorized ", 401)
}

func (o *CheckDeploymentNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckDeploymentNameUsingGETNotFound creates a CheckDeploymentNameUsingGETNotFound with default headers values
func NewCheckDeploymentNameUsingGETNotFound() *CheckDeploymentNameUsingGETNotFound {
	return &CheckDeploymentNameUsingGETNotFound{}
}

/*CheckDeploymentNameUsingGETNotFound handles this case with default header values.

Not Found
*/
type CheckDeploymentNameUsingGETNotFound struct {
}

func (o *CheckDeploymentNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/names/{name}][%d] checkDeploymentNameUsingGETNotFound ", 404)
}

func (o *CheckDeploymentNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
