// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePolicyUsingDELETEReader is a Reader for the DeletePolicyUsingDELETE structure.
type DeletePolicyUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePolicyUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePolicyUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePolicyUsingDELETENoContent creates a DeletePolicyUsingDELETENoContent with default headers values
func NewDeletePolicyUsingDELETENoContent() *DeletePolicyUsingDELETENoContent {
	return &DeletePolicyUsingDELETENoContent{}
}

/*DeletePolicyUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeletePolicyUsingDELETENoContent struct {
}

func (o *DeletePolicyUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /policy/api/policies/{id}][%d] deletePolicyUsingDELETENoContent ", 204)
}

func (o *DeletePolicyUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}