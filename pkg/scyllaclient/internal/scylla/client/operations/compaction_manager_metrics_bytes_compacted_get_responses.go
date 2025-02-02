// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla/models"
)

// CompactionManagerMetricsBytesCompactedGetReader is a Reader for the CompactionManagerMetricsBytesCompactedGet structure.
type CompactionManagerMetricsBytesCompactedGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompactionManagerMetricsBytesCompactedGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompactionManagerMetricsBytesCompactedGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCompactionManagerMetricsBytesCompactedGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCompactionManagerMetricsBytesCompactedGetOK creates a CompactionManagerMetricsBytesCompactedGetOK with default headers values
func NewCompactionManagerMetricsBytesCompactedGetOK() *CompactionManagerMetricsBytesCompactedGetOK {
	return &CompactionManagerMetricsBytesCompactedGetOK{}
}

/*
CompactionManagerMetricsBytesCompactedGetOK handles this case with default header values.

CompactionManagerMetricsBytesCompactedGetOK compaction manager metrics bytes compacted get o k
*/
type CompactionManagerMetricsBytesCompactedGetOK struct {
	Payload int32
}

func (o *CompactionManagerMetricsBytesCompactedGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CompactionManagerMetricsBytesCompactedGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompactionManagerMetricsBytesCompactedGetDefault creates a CompactionManagerMetricsBytesCompactedGetDefault with default headers values
func NewCompactionManagerMetricsBytesCompactedGetDefault(code int) *CompactionManagerMetricsBytesCompactedGetDefault {
	return &CompactionManagerMetricsBytesCompactedGetDefault{
		_statusCode: code,
	}
}

/*
CompactionManagerMetricsBytesCompactedGetDefault handles this case with default header values.

internal server error
*/
type CompactionManagerMetricsBytesCompactedGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the compaction manager metrics bytes compacted get default response
func (o *CompactionManagerMetricsBytesCompactedGetDefault) Code() int {
	return o._statusCode
}

func (o *CompactionManagerMetricsBytesCompactedGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CompactionManagerMetricsBytesCompactedGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CompactionManagerMetricsBytesCompactedGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
