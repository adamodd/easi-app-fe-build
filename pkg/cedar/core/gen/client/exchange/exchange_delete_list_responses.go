// Code generated by go-swagger; DO NOT EDIT.

package exchange

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// ExchangeDeleteListReader is a Reader for the ExchangeDeleteList structure.
type ExchangeDeleteListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExchangeDeleteListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExchangeDeleteListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExchangeDeleteListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewExchangeDeleteListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExchangeDeleteListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExchangeDeleteListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExchangeDeleteListOK creates a ExchangeDeleteListOK with default headers values
func NewExchangeDeleteListOK() *ExchangeDeleteListOK {
	return &ExchangeDeleteListOK{}
}

/* ExchangeDeleteListOK describes a response with status code 200, with default header values.

OK
*/
type ExchangeDeleteListOK struct {
	Payload *models.Response
}

func (o *ExchangeDeleteListOK) Error() string {
	return fmt.Sprintf("[DELETE /exchange][%d] exchangeDeleteListOK  %+v", 200, o.Payload)
}
func (o *ExchangeDeleteListOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeDeleteListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeDeleteListBadRequest creates a ExchangeDeleteListBadRequest with default headers values
func NewExchangeDeleteListBadRequest() *ExchangeDeleteListBadRequest {
	return &ExchangeDeleteListBadRequest{}
}

/* ExchangeDeleteListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExchangeDeleteListBadRequest struct {
	Payload *models.Response
}

func (o *ExchangeDeleteListBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /exchange][%d] exchangeDeleteListBadRequest  %+v", 400, o.Payload)
}
func (o *ExchangeDeleteListBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeDeleteListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeDeleteListUnauthorized creates a ExchangeDeleteListUnauthorized with default headers values
func NewExchangeDeleteListUnauthorized() *ExchangeDeleteListUnauthorized {
	return &ExchangeDeleteListUnauthorized{}
}

/* ExchangeDeleteListUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type ExchangeDeleteListUnauthorized struct {
	Payload *models.Response
}

func (o *ExchangeDeleteListUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /exchange][%d] exchangeDeleteListUnauthorized  %+v", 401, o.Payload)
}
func (o *ExchangeDeleteListUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeDeleteListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeDeleteListNotFound creates a ExchangeDeleteListNotFound with default headers values
func NewExchangeDeleteListNotFound() *ExchangeDeleteListNotFound {
	return &ExchangeDeleteListNotFound{}
}

/* ExchangeDeleteListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExchangeDeleteListNotFound struct {
	Payload *models.Response
}

func (o *ExchangeDeleteListNotFound) Error() string {
	return fmt.Sprintf("[DELETE /exchange][%d] exchangeDeleteListNotFound  %+v", 404, o.Payload)
}
func (o *ExchangeDeleteListNotFound) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeDeleteListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExchangeDeleteListInternalServerError creates a ExchangeDeleteListInternalServerError with default headers values
func NewExchangeDeleteListInternalServerError() *ExchangeDeleteListInternalServerError {
	return &ExchangeDeleteListInternalServerError{}
}

/* ExchangeDeleteListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExchangeDeleteListInternalServerError struct {
	Payload *models.Response
}

func (o *ExchangeDeleteListInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /exchange][%d] exchangeDeleteListInternalServerError  %+v", 500, o.Payload)
}
func (o *ExchangeDeleteListInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *ExchangeDeleteListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}