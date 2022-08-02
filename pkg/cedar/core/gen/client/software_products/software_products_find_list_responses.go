// Code generated by go-swagger; DO NOT EDIT.

package software_products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// SoftwareProductsFindListReader is a Reader for the SoftwareProductsFindList structure.
type SoftwareProductsFindListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareProductsFindListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareProductsFindListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSoftwareProductsFindListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSoftwareProductsFindListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSoftwareProductsFindListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSoftwareProductsFindListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSoftwareProductsFindListOK creates a SoftwareProductsFindListOK with default headers values
func NewSoftwareProductsFindListOK() *SoftwareProductsFindListOK {
	return &SoftwareProductsFindListOK{}
}

/* SoftwareProductsFindListOK describes a response with status code 200, with default header values.

OK
*/
type SoftwareProductsFindListOK struct {
	Payload *models.SoftwareProductsFindResponse
}

func (o *SoftwareProductsFindListOK) Error() string {
	return fmt.Sprintf("[GET /softwareProducts][%d] softwareProductsFindListOK  %+v", 200, o.Payload)
}
func (o *SoftwareProductsFindListOK) GetPayload() *models.SoftwareProductsFindResponse {
	return o.Payload
}

func (o *SoftwareProductsFindListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareProductsFindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareProductsFindListBadRequest creates a SoftwareProductsFindListBadRequest with default headers values
func NewSoftwareProductsFindListBadRequest() *SoftwareProductsFindListBadRequest {
	return &SoftwareProductsFindListBadRequest{}
}

/* SoftwareProductsFindListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SoftwareProductsFindListBadRequest struct {
	Payload *models.Response
}

func (o *SoftwareProductsFindListBadRequest) Error() string {
	return fmt.Sprintf("[GET /softwareProducts][%d] softwareProductsFindListBadRequest  %+v", 400, o.Payload)
}
func (o *SoftwareProductsFindListBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *SoftwareProductsFindListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareProductsFindListUnauthorized creates a SoftwareProductsFindListUnauthorized with default headers values
func NewSoftwareProductsFindListUnauthorized() *SoftwareProductsFindListUnauthorized {
	return &SoftwareProductsFindListUnauthorized{}
}

/* SoftwareProductsFindListUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type SoftwareProductsFindListUnauthorized struct {
	Payload *models.Response
}

func (o *SoftwareProductsFindListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /softwareProducts][%d] softwareProductsFindListUnauthorized  %+v", 401, o.Payload)
}
func (o *SoftwareProductsFindListUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *SoftwareProductsFindListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareProductsFindListNotFound creates a SoftwareProductsFindListNotFound with default headers values
func NewSoftwareProductsFindListNotFound() *SoftwareProductsFindListNotFound {
	return &SoftwareProductsFindListNotFound{}
}

/* SoftwareProductsFindListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SoftwareProductsFindListNotFound struct {
	Payload *models.Response
}

func (o *SoftwareProductsFindListNotFound) Error() string {
	return fmt.Sprintf("[GET /softwareProducts][%d] softwareProductsFindListNotFound  %+v", 404, o.Payload)
}
func (o *SoftwareProductsFindListNotFound) GetPayload() *models.Response {
	return o.Payload
}

func (o *SoftwareProductsFindListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareProductsFindListInternalServerError creates a SoftwareProductsFindListInternalServerError with default headers values
func NewSoftwareProductsFindListInternalServerError() *SoftwareProductsFindListInternalServerError {
	return &SoftwareProductsFindListInternalServerError{}
}

/* SoftwareProductsFindListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SoftwareProductsFindListInternalServerError struct {
	Payload *models.Response
}

func (o *SoftwareProductsFindListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /softwareProducts][%d] softwareProductsFindListInternalServerError  %+v", 500, o.Payload)
}
func (o *SoftwareProductsFindListInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *SoftwareProductsFindListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}