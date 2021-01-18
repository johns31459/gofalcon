// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRListPutFilesReader is a Reader for the RTRListPutFiles structure.
type RTRListPutFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRListPutFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRListPutFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRListPutFilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRListPutFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRListPutFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRListPutFilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRTRListPutFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRTRListPutFilesOK creates a RTRListPutFilesOK with default headers values
func NewRTRListPutFilesOK() *RTRListPutFilesOK {
	return &RTRListPutFilesOK{}
}

/*RTRListPutFilesOK handles this case with default header values.

OK
*/
type RTRListPutFilesOK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.BinservclientMsaPutFileResponse
}

func (o *RTRListPutFilesOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesOK  %+v", 200, o.Payload)
}

func (o *RTRListPutFilesOK) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListPutFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListPutFilesBadRequest creates a RTRListPutFilesBadRequest with default headers values
func NewRTRListPutFilesBadRequest() *RTRListPutFilesBadRequest {
	return &RTRListPutFilesBadRequest{}
}

/*RTRListPutFilesBadRequest handles this case with default header values.

Bad Request
*/
type RTRListPutFilesBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListPutFilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesBadRequest  %+v", 400, o.Payload)
}

func (o *RTRListPutFilesBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListPutFilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListPutFilesForbidden creates a RTRListPutFilesForbidden with default headers values
func NewRTRListPutFilesForbidden() *RTRListPutFilesForbidden {
	return &RTRListPutFilesForbidden{}
}

/*RTRListPutFilesForbidden handles this case with default header values.

Forbidden
*/
type RTRListPutFilesForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *RTRListPutFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesForbidden  %+v", 403, o.Payload)
}

func (o *RTRListPutFilesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListPutFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListPutFilesNotFound creates a RTRListPutFilesNotFound with default headers values
func NewRTRListPutFilesNotFound() *RTRListPutFilesNotFound {
	return &RTRListPutFilesNotFound{}
}

/*RTRListPutFilesNotFound handles this case with default header values.

Not Found
*/
type RTRListPutFilesNotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

func (o *RTRListPutFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesNotFound  %+v", 404, o.Payload)
}

func (o *RTRListPutFilesNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRListPutFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRListPutFilesTooManyRequests creates a RTRListPutFilesTooManyRequests with default headers values
func NewRTRListPutFilesTooManyRequests() *RTRListPutFilesTooManyRequests {
	return &RTRListPutFilesTooManyRequests{}
}

/*RTRListPutFilesTooManyRequests handles this case with default header values.

Too Many Requests
*/
type RTRListPutFilesTooManyRequests struct {
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

func (o *RTRListPutFilesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] rTRListPutFilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRListPutFilesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRListPutFilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRListPutFilesDefault creates a RTRListPutFilesDefault with default headers values
func NewRTRListPutFilesDefault(code int) *RTRListPutFilesDefault {
	return &RTRListPutFilesDefault{
		_statusCode: code,
	}
}

/*RTRListPutFilesDefault handles this case with default header values.

OK
*/
type RTRListPutFilesDefault struct {
	_statusCode int

	Payload *models.BinservclientMsaPutFileResponse
}

// Code gets the status code for the r t r list put files default response
func (o *RTRListPutFilesDefault) Code() int {
	return o._statusCode
}

func (o *RTRListPutFilesDefault) Error() string {
	return fmt.Sprintf("[GET /real-time-response/queries/put-files/v1][%d] RTR-ListPut-Files default  %+v", o._statusCode, o.Payload)
}

func (o *RTRListPutFilesDefault) GetPayload() *models.BinservclientMsaPutFileResponse {
	return o.Payload
}

func (o *RTRListPutFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BinservclientMsaPutFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}