// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// UserNameUpdateReader is a Reader for the UserNameUpdate structure.
type UserNameUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserNameUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserNameUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserNameUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserNameUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserNameUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserNameUpdateOK creates a UserNameUpdateOK with default headers values
func NewUserNameUpdateOK() *UserNameUpdateOK {
	return &UserNameUpdateOK{}
}

/* UserNameUpdateOK describes a response with status code 200, with default header values.

OK
*/
type UserNameUpdateOK struct {
	Payload *models.Response
}

func (o *UserNameUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /user/username/{username}][%d] userNameUpdateOK  %+v", 200, o.Payload)
}
func (o *UserNameUpdateOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserNameUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserNameUpdateBadRequest creates a UserNameUpdateBadRequest with default headers values
func NewUserNameUpdateBadRequest() *UserNameUpdateBadRequest {
	return &UserNameUpdateBadRequest{}
}

/* UserNameUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserNameUpdateBadRequest struct {
	Payload *models.Response
}

func (o *UserNameUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /user/username/{username}][%d] userNameUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *UserNameUpdateBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserNameUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserNameUpdateUnauthorized creates a UserNameUpdateUnauthorized with default headers values
func NewUserNameUpdateUnauthorized() *UserNameUpdateUnauthorized {
	return &UserNameUpdateUnauthorized{}
}

/* UserNameUpdateUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type UserNameUpdateUnauthorized struct {
	Payload *models.Response
}

func (o *UserNameUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user/username/{username}][%d] userNameUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *UserNameUpdateUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserNameUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserNameUpdateInternalServerError creates a UserNameUpdateInternalServerError with default headers values
func NewUserNameUpdateInternalServerError() *UserNameUpdateInternalServerError {
	return &UserNameUpdateInternalServerError{}
}

/* UserNameUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UserNameUpdateInternalServerError struct {
	Payload *models.Response
}

func (o *UserNameUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user/username/{username}][%d] userNameUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *UserNameUpdateInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *UserNameUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}