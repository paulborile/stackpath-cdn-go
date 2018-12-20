// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/renderinc/stackpath-cdn-go/models"
)

// DeleteSiteScriptReader is a Reader for the DeleteSiteScript structure.
type DeleteSiteScriptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSiteScriptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSiteScriptNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteSiteScriptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteSiteScriptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSiteScriptDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSiteScriptNoContent creates a DeleteSiteScriptNoContent with default headers values
func NewDeleteSiteScriptNoContent() *DeleteSiteScriptNoContent {
	return &DeleteSiteScriptNoContent{}
}

/*DeleteSiteScriptNoContent handles this case with default header values.

No content
*/
type DeleteSiteScriptNoContent struct {
}

func (o *DeleteSiteScriptNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}][%d] deleteSiteScriptNoContent ", 204)
}

func (o *DeleteSiteScriptNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSiteScriptUnauthorized creates a DeleteSiteScriptUnauthorized with default headers values
func NewDeleteSiteScriptUnauthorized() *DeleteSiteScriptUnauthorized {
	return &DeleteSiteScriptUnauthorized{}
}

/*DeleteSiteScriptUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type DeleteSiteScriptUnauthorized struct {
	Payload *models.APIStatus
}

func (o *DeleteSiteScriptUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}][%d] deleteSiteScriptUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSiteScriptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSiteScriptInternalServerError creates a DeleteSiteScriptInternalServerError with default headers values
func NewDeleteSiteScriptInternalServerError() *DeleteSiteScriptInternalServerError {
	return &DeleteSiteScriptInternalServerError{}
}

/*DeleteSiteScriptInternalServerError handles this case with default header values.

Internal server error.
*/
type DeleteSiteScriptInternalServerError struct {
	Payload *models.APIStatus
}

func (o *DeleteSiteScriptInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}][%d] deleteSiteScriptInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSiteScriptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSiteScriptDefault creates a DeleteSiteScriptDefault with default headers values
func NewDeleteSiteScriptDefault(code int) *DeleteSiteScriptDefault {
	return &DeleteSiteScriptDefault{
		_statusCode: code,
	}
}

/*DeleteSiteScriptDefault handles this case with default header values.

Default error structure.
*/
type DeleteSiteScriptDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the delete site script default response
func (o *DeleteSiteScriptDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSiteScriptDefault) Error() string {
	return fmt.Sprintf("[DELETE /cdn/v1/stacks/{stack_id}/sites/{site_id}/scripts/{script_id}][%d] DeleteSiteScript default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSiteScriptDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}