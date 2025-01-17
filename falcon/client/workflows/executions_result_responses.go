// Code generated by go-swagger; DO NOT EDIT.

package workflows

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

// ExecutionsResultReader is a Reader for the ExecutionsResult structure.
type ExecutionsResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecutionsResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecutionsResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExecutionsResultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecutionsResultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExecutionsResultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExecutionsResultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecutionsResultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /workflows/entities/execution-results/v1] executions.result", response, response.Code())
	}
}

// NewExecutionsResultOK creates a ExecutionsResultOK with default headers values
func NewExecutionsResultOK() *ExecutionsResultOK {
	return &ExecutionsResultOK{}
}

/*
ExecutionsResultOK describes a response with status code 200, with default header values.

OK
*/
type ExecutionsResultOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this executions result o k response has a 2xx status code
func (o *ExecutionsResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this executions result o k response has a 3xx status code
func (o *ExecutionsResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this executions result o k response has a 4xx status code
func (o *ExecutionsResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this executions result o k response has a 5xx status code
func (o *ExecutionsResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this executions result o k response a status code equal to that given
func (o *ExecutionsResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the executions result o k response
func (o *ExecutionsResultOK) Code() int {
	return 200
}

func (o *ExecutionsResultOK) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultOK  %+v", 200, o.Payload)
}

func (o *ExecutionsResultOK) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultOK  %+v", 200, o.Payload)
}

func (o *ExecutionsResultOK) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *ExecutionsResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionsResultBadRequest creates a ExecutionsResultBadRequest with default headers values
func NewExecutionsResultBadRequest() *ExecutionsResultBadRequest {
	return &ExecutionsResultBadRequest{}
}

/*
ExecutionsResultBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExecutionsResultBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this executions result bad request response has a 2xx status code
func (o *ExecutionsResultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this executions result bad request response has a 3xx status code
func (o *ExecutionsResultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this executions result bad request response has a 4xx status code
func (o *ExecutionsResultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this executions result bad request response has a 5xx status code
func (o *ExecutionsResultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this executions result bad request response a status code equal to that given
func (o *ExecutionsResultBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the executions result bad request response
func (o *ExecutionsResultBadRequest) Code() int {
	return 400
}

func (o *ExecutionsResultBadRequest) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultBadRequest  %+v", 400, o.Payload)
}

func (o *ExecutionsResultBadRequest) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultBadRequest  %+v", 400, o.Payload)
}

func (o *ExecutionsResultBadRequest) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *ExecutionsResultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionsResultForbidden creates a ExecutionsResultForbidden with default headers values
func NewExecutionsResultForbidden() *ExecutionsResultForbidden {
	return &ExecutionsResultForbidden{}
}

/*
ExecutionsResultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExecutionsResultForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this executions result forbidden response has a 2xx status code
func (o *ExecutionsResultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this executions result forbidden response has a 3xx status code
func (o *ExecutionsResultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this executions result forbidden response has a 4xx status code
func (o *ExecutionsResultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this executions result forbidden response has a 5xx status code
func (o *ExecutionsResultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this executions result forbidden response a status code equal to that given
func (o *ExecutionsResultForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the executions result forbidden response
func (o *ExecutionsResultForbidden) Code() int {
	return 403
}

func (o *ExecutionsResultForbidden) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultForbidden  %+v", 403, o.Payload)
}

func (o *ExecutionsResultForbidden) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultForbidden  %+v", 403, o.Payload)
}

func (o *ExecutionsResultForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExecutionsResultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionsResultNotFound creates a ExecutionsResultNotFound with default headers values
func NewExecutionsResultNotFound() *ExecutionsResultNotFound {
	return &ExecutionsResultNotFound{}
}

/*
ExecutionsResultNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExecutionsResultNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this executions result not found response has a 2xx status code
func (o *ExecutionsResultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this executions result not found response has a 3xx status code
func (o *ExecutionsResultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this executions result not found response has a 4xx status code
func (o *ExecutionsResultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this executions result not found response has a 5xx status code
func (o *ExecutionsResultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this executions result not found response a status code equal to that given
func (o *ExecutionsResultNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the executions result not found response
func (o *ExecutionsResultNotFound) Code() int {
	return 404
}

func (o *ExecutionsResultNotFound) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultNotFound  %+v", 404, o.Payload)
}

func (o *ExecutionsResultNotFound) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultNotFound  %+v", 404, o.Payload)
}

func (o *ExecutionsResultNotFound) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *ExecutionsResultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionsResultTooManyRequests creates a ExecutionsResultTooManyRequests with default headers values
func NewExecutionsResultTooManyRequests() *ExecutionsResultTooManyRequests {
	return &ExecutionsResultTooManyRequests{}
}

/*
ExecutionsResultTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ExecutionsResultTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this executions result too many requests response has a 2xx status code
func (o *ExecutionsResultTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this executions result too many requests response has a 3xx status code
func (o *ExecutionsResultTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this executions result too many requests response has a 4xx status code
func (o *ExecutionsResultTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this executions result too many requests response has a 5xx status code
func (o *ExecutionsResultTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this executions result too many requests response a status code equal to that given
func (o *ExecutionsResultTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the executions result too many requests response
func (o *ExecutionsResultTooManyRequests) Code() int {
	return 429
}

func (o *ExecutionsResultTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExecutionsResultTooManyRequests) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExecutionsResultTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ExecutionsResultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecutionsResultInternalServerError creates a ExecutionsResultInternalServerError with default headers values
func NewExecutionsResultInternalServerError() *ExecutionsResultInternalServerError {
	return &ExecutionsResultInternalServerError{}
}

/*
ExecutionsResultInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ExecutionsResultInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIExecutionResultsResponse
}

// IsSuccess returns true when this executions result internal server error response has a 2xx status code
func (o *ExecutionsResultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this executions result internal server error response has a 3xx status code
func (o *ExecutionsResultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this executions result internal server error response has a 4xx status code
func (o *ExecutionsResultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this executions result internal server error response has a 5xx status code
func (o *ExecutionsResultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this executions result internal server error response a status code equal to that given
func (o *ExecutionsResultInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the executions result internal server error response
func (o *ExecutionsResultInternalServerError) Code() int {
	return 500
}

func (o *ExecutionsResultInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecutionsResultInternalServerError) String() string {
	return fmt.Sprintf("[GET /workflows/entities/execution-results/v1][%d] executionsResultInternalServerError  %+v", 500, o.Payload)
}

func (o *ExecutionsResultInternalServerError) GetPayload() *models.APIExecutionResultsResponse {
	return o.Payload
}

func (o *ExecutionsResultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIExecutionResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
