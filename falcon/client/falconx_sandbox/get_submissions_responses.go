// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// GetSubmissionsReader is a Reader for the GetSubmissions structure.
type GetSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubmissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSubmissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSubmissionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSubmissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSubmissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubmissionsOK creates a GetSubmissionsOK with default headers values
func NewGetSubmissionsOK() *GetSubmissionsOK {
	return &GetSubmissionsOK{}
}

/*GetSubmissionsOK handles this case with default header values.

OK
*/
type GetSubmissionsOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSubmissionV1Response
}

func (o *GetSubmissionsOK) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsOK  %+v", 200, o.Payload)
}

func (o *GetSubmissionsOK) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsBadRequest creates a GetSubmissionsBadRequest with default headers values
func NewGetSubmissionsBadRequest() *GetSubmissionsBadRequest {
	return &GetSubmissionsBadRequest{}
}

/*GetSubmissionsBadRequest handles this case with default header values.

Bad Request
*/
type GetSubmissionsBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSubmissionV1Response
}

func (o *GetSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubmissionsBadRequest) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsForbidden creates a GetSubmissionsForbidden with default headers values
func NewGetSubmissionsForbidden() *GetSubmissionsForbidden {
	return &GetSubmissionsForbidden{}
}

/*GetSubmissionsForbidden handles this case with default header values.

Forbidden
*/
type GetSubmissionsForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetSubmissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetSubmissionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSubmissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsTooManyRequests creates a GetSubmissionsTooManyRequests with default headers values
func NewGetSubmissionsTooManyRequests() *GetSubmissionsTooManyRequests {
	return &GetSubmissionsTooManyRequests{}
}

/*GetSubmissionsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetSubmissionsTooManyRequests struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
	/*Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetSubmissionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSubmissionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSubmissionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-RetryAfter
	xRateLimitRetryAfter, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-RetryAfter"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", response.GetHeader("X-RateLimit-RetryAfter"))
	}
	o.XRateLimitRetryAfter = xRateLimitRetryAfter

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsInternalServerError creates a GetSubmissionsInternalServerError with default headers values
func NewGetSubmissionsInternalServerError() *GetSubmissionsInternalServerError {
	return &GetSubmissionsInternalServerError{}
}

/*GetSubmissionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSubmissionsInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FalconxSubmissionV1Response
}

func (o *GetSubmissionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] getSubmissionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSubmissionsInternalServerError) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubmissionsDefault creates a GetSubmissionsDefault with default headers values
func NewGetSubmissionsDefault(code int) *GetSubmissionsDefault {
	return &GetSubmissionsDefault{
		_statusCode: code,
	}
}

/*GetSubmissionsDefault handles this case with default header values.

OK
*/
type GetSubmissionsDefault struct {
	_statusCode int

	Payload *models.FalconxSubmissionV1Response
}

// Code gets the status code for the get submissions default response
func (o *GetSubmissionsDefault) Code() int {
	return o._statusCode
}

func (o *GetSubmissionsDefault) Error() string {
	return fmt.Sprintf("[GET /falconx/entities/submissions/v1][%d] GetSubmissions default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubmissionsDefault) GetPayload() *models.FalconxSubmissionV1Response {
	return o.Payload
}

func (o *GetSubmissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FalconxSubmissionV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}