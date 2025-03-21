// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-civitai/models"
)

// GetModelsReader is a Reader for the GetModels structure.
type GetModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetModelsOK creates a GetModelsOK with default headers values
func NewGetModelsOK() *GetModelsOK {
	return &GetModelsOK{}
}

/*
GetModelsOK describes a response with status code 200, with default header values.

A list of models.
*/
type GetModelsOK struct {
	Payload *models.ModelsResponse
}

// IsSuccess returns true when this get models o k response has a 2xx status code
func (o *GetModelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get models o k response has a 3xx status code
func (o *GetModelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get models o k response has a 4xx status code
func (o *GetModelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get models o k response has a 5xx status code
func (o *GetModelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get models o k response a status code equal to that given
func (o *GetModelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get models o k response
func (o *GetModelsOK) Code() int {
	return 200
}

func (o *GetModelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /models][%d] getModelsOK %s", 200, payload)
}

func (o *GetModelsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /models][%d] getModelsOK %s", 200, payload)
}

func (o *GetModelsOK) GetPayload() *models.ModelsResponse {
	return o.Payload
}

func (o *GetModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelsDefault creates a GetModelsDefault with default headers values
func NewGetModelsDefault(code int) *GetModelsDefault {
	return &GetModelsDefault{
		_statusCode: code,
	}
}

/*
GetModelsDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetModelsDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// IsSuccess returns true when this get models default response has a 2xx status code
func (o *GetModelsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get models default response has a 3xx status code
func (o *GetModelsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get models default response has a 4xx status code
func (o *GetModelsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get models default response has a 5xx status code
func (o *GetModelsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get models default response a status code equal to that given
func (o *GetModelsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get models default response
func (o *GetModelsDefault) Code() int {
	return o._statusCode
}

func (o *GetModelsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /models][%d] getModels default %s", o._statusCode, payload)
}

func (o *GetModelsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /models][%d] getModels default %s", o._statusCode, payload)
}

func (o *GetModelsDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
