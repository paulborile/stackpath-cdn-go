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

// UpdateScopeRuleConfigurationReader is a Reader for the UpdateScopeRuleConfiguration structure.
type UpdateScopeRuleConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScopeRuleConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateScopeRuleConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateScopeRuleConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewUpdateScopeRuleConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateScopeRuleConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateScopeRuleConfigurationOK creates a UpdateScopeRuleConfigurationOK with default headers values
func NewUpdateScopeRuleConfigurationOK() *UpdateScopeRuleConfigurationOK {
	return &UpdateScopeRuleConfigurationOK{}
}

/*UpdateScopeRuleConfigurationOK handles this case with default header values.

UpdateScopeRuleConfigurationOK update scope rule configuration o k
*/
type UpdateScopeRuleConfigurationOK struct {
	Payload *models.CdnUpdateScopeRuleConfigurationResponse
}

func (o *UpdateScopeRuleConfigurationOK) Error() string {
	return fmt.Sprintf("[PATCH /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] updateScopeRuleConfigurationOK  %+v", 200, o.Payload)
}

func (o *UpdateScopeRuleConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CdnUpdateScopeRuleConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScopeRuleConfigurationUnauthorized creates a UpdateScopeRuleConfigurationUnauthorized with default headers values
func NewUpdateScopeRuleConfigurationUnauthorized() *UpdateScopeRuleConfigurationUnauthorized {
	return &UpdateScopeRuleConfigurationUnauthorized{}
}

/*UpdateScopeRuleConfigurationUnauthorized handles this case with default header values.

Returned when an unauthorized request is attempted.
*/
type UpdateScopeRuleConfigurationUnauthorized struct {
	Payload *models.APIStatus
}

func (o *UpdateScopeRuleConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] updateScopeRuleConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateScopeRuleConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScopeRuleConfigurationInternalServerError creates a UpdateScopeRuleConfigurationInternalServerError with default headers values
func NewUpdateScopeRuleConfigurationInternalServerError() *UpdateScopeRuleConfigurationInternalServerError {
	return &UpdateScopeRuleConfigurationInternalServerError{}
}

/*UpdateScopeRuleConfigurationInternalServerError handles this case with default header values.

Internal server error.
*/
type UpdateScopeRuleConfigurationInternalServerError struct {
	Payload *models.APIStatus
}

func (o *UpdateScopeRuleConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] updateScopeRuleConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateScopeRuleConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScopeRuleConfigurationDefault creates a UpdateScopeRuleConfigurationDefault with default headers values
func NewUpdateScopeRuleConfigurationDefault(code int) *UpdateScopeRuleConfigurationDefault {
	return &UpdateScopeRuleConfigurationDefault{
		_statusCode: code,
	}
}

/*UpdateScopeRuleConfigurationDefault handles this case with default header values.

Default error structure.
*/
type UpdateScopeRuleConfigurationDefault struct {
	_statusCode int

	Payload *models.APIStatus
}

// Code gets the status code for the update scope rule configuration default response
func (o *UpdateScopeRuleConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateScopeRuleConfigurationDefault) Error() string {
	return fmt.Sprintf("[PATCH /cdn/v1/stacks/{stack_id}/sites/{site_id}/scopes/{scope_id}/rules/{rule_id}/configuration][%d] UpdateScopeRuleConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateScopeRuleConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}