// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// RoleDeleteListReader is a Reader for the RoleDeleteList structure.
type RoleDeleteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleDeleteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleDeleteListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoleDeleteListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRoleDeleteListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRoleDeleteListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRoleDeleteListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoleDeleteListOK creates a RoleDeleteListOK with default headers values
func NewRoleDeleteListOK() *RoleDeleteListOK {
	return &RoleDeleteListOK{}
}

/* RoleDeleteListOK describes a response with status code 200, with default header values.

OK
*/
type RoleDeleteListOK struct {
	Payload *models.Response
}

func (o *RoleDeleteListOK) Error() string {
	return fmt.Sprintf("[DELETE /role][%d] roleDeleteListOK  %+v", 200, o.Payload)
}
func (o *RoleDeleteListOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *RoleDeleteListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleDeleteListBadRequest creates a RoleDeleteListBadRequest with default headers values
func NewRoleDeleteListBadRequest() *RoleDeleteListBadRequest {
	return &RoleDeleteListBadRequest{}
}

/* RoleDeleteListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RoleDeleteListBadRequest struct {
	Payload *models.Response
}

func (o *RoleDeleteListBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /role][%d] roleDeleteListBadRequest  %+v", 400, o.Payload)
}
func (o *RoleDeleteListBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *RoleDeleteListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleDeleteListUnauthorized creates a RoleDeleteListUnauthorized with default headers values
func NewRoleDeleteListUnauthorized() *RoleDeleteListUnauthorized {
	return &RoleDeleteListUnauthorized{}
}

/* RoleDeleteListUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type RoleDeleteListUnauthorized struct {
	Payload *models.Response
}

func (o *RoleDeleteListUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /role][%d] roleDeleteListUnauthorized  %+v", 401, o.Payload)
}
func (o *RoleDeleteListUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *RoleDeleteListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleDeleteListNotFound creates a RoleDeleteListNotFound with default headers values
func NewRoleDeleteListNotFound() *RoleDeleteListNotFound {
	return &RoleDeleteListNotFound{}
}

/* RoleDeleteListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RoleDeleteListNotFound struct {
	Payload *models.Response
}

func (o *RoleDeleteListNotFound) Error() string {
	return fmt.Sprintf("[DELETE /role][%d] roleDeleteListNotFound  %+v", 404, o.Payload)
}
func (o *RoleDeleteListNotFound) GetPayload() *models.Response {
	return o.Payload
}

func (o *RoleDeleteListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoleDeleteListInternalServerError creates a RoleDeleteListInternalServerError with default headers values
func NewRoleDeleteListInternalServerError() *RoleDeleteListInternalServerError {
	return &RoleDeleteListInternalServerError{}
}

/* RoleDeleteListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type RoleDeleteListInternalServerError struct {
	Payload *models.Response
}

func (o *RoleDeleteListInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /role][%d] roleDeleteListInternalServerError  %+v", 500, o.Payload)
}
func (o *RoleDeleteListInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *RoleDeleteListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}