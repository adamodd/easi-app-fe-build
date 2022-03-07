// Code generated by go-swagger; DO NOT EDIT.

package intake

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/intake/gen/models"
)

// IntakeStatusByCedarIDReader is a Reader for the IntakeStatusByCedarID structure.
type IntakeStatusByCedarIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntakeStatusByCedarIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntakeStatusByCedarIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIntakeStatusByCedarIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIntakeStatusByCedarIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIntakeStatusByCedarIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIntakeStatusByCedarIDOK creates a IntakeStatusByCedarIDOK with default headers values
func NewIntakeStatusByCedarIDOK() *IntakeStatusByCedarIDOK {
	return &IntakeStatusByCedarIDOK{}
}

/* IntakeStatusByCedarIDOK describes a response with status code 200, with default header values.

OK
*/
type IntakeStatusByCedarIDOK struct {
	Payload *models.IntakeStatus
}

func (o *IntakeStatusByCedarIDOK) Error() string {
	return fmt.Sprintf("[GET /intake/status/cedar/{id}][%d] intakeStatusByCedarIdOK  %+v", 200, o.Payload)
}
func (o *IntakeStatusByCedarIDOK) GetPayload() *models.IntakeStatus {
	return o.Payload
}

func (o *IntakeStatusByCedarIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntakeStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntakeStatusByCedarIDBadRequest creates a IntakeStatusByCedarIDBadRequest with default headers values
func NewIntakeStatusByCedarIDBadRequest() *IntakeStatusByCedarIDBadRequest {
	return &IntakeStatusByCedarIDBadRequest{}
}

/* IntakeStatusByCedarIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IntakeStatusByCedarIDBadRequest struct {
	Payload *models.Response
}

func (o *IntakeStatusByCedarIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /intake/status/cedar/{id}][%d] intakeStatusByCedarIdBadRequest  %+v", 400, o.Payload)
}
func (o *IntakeStatusByCedarIDBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *IntakeStatusByCedarIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntakeStatusByCedarIDUnauthorized creates a IntakeStatusByCedarIDUnauthorized with default headers values
func NewIntakeStatusByCedarIDUnauthorized() *IntakeStatusByCedarIDUnauthorized {
	return &IntakeStatusByCedarIDUnauthorized{}
}

/* IntakeStatusByCedarIDUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type IntakeStatusByCedarIDUnauthorized struct {
	Payload *models.Response
}

func (o *IntakeStatusByCedarIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /intake/status/cedar/{id}][%d] intakeStatusByCedarIdUnauthorized  %+v", 401, o.Payload)
}
func (o *IntakeStatusByCedarIDUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *IntakeStatusByCedarIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntakeStatusByCedarIDInternalServerError creates a IntakeStatusByCedarIDInternalServerError with default headers values
func NewIntakeStatusByCedarIDInternalServerError() *IntakeStatusByCedarIDInternalServerError {
	return &IntakeStatusByCedarIDInternalServerError{}
}

/* IntakeStatusByCedarIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type IntakeStatusByCedarIDInternalServerError struct {
	Payload *models.Response
}

func (o *IntakeStatusByCedarIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intake/status/cedar/{id}][%d] intakeStatusByCedarIdInternalServerError  %+v", 500, o.Payload)
}
func (o *IntakeStatusByCedarIDInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *IntakeStatusByCedarIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}