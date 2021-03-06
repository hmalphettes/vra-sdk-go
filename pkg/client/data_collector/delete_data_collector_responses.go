// Code generated by go-swagger; DO NOT EDIT.

package data_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteDataCollectorReader is a Reader for the DeleteDataCollector structure.
type DeleteDataCollectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDataCollectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteDataCollectorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteDataCollectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDataCollectorNoContent creates a DeleteDataCollectorNoContent with default headers values
func NewDeleteDataCollectorNoContent() *DeleteDataCollectorNoContent {
	return &DeleteDataCollectorNoContent{}
}

/*DeleteDataCollectorNoContent handles this case with default header values.

No Content
*/
type DeleteDataCollectorNoContent struct {
}

func (o *DeleteDataCollectorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/data-collectors/{id}][%d] deleteDataCollectorNoContent ", 204)
}

func (o *DeleteDataCollectorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDataCollectorForbidden creates a DeleteDataCollectorForbidden with default headers values
func NewDeleteDataCollectorForbidden() *DeleteDataCollectorForbidden {
	return &DeleteDataCollectorForbidden{}
}

/*DeleteDataCollectorForbidden handles this case with default header values.

Forbidden
*/
type DeleteDataCollectorForbidden struct {
}

func (o *DeleteDataCollectorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/data-collectors/{id}][%d] deleteDataCollectorForbidden ", 403)
}

func (o *DeleteDataCollectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
