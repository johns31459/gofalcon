// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// GetLatestIntelRuleFileReader is a Reader for the GetLatestIntelRuleFile structure.
type GetLatestIntelRuleFileReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestIntelRuleFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestIntelRuleFileOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetLatestIntelRuleFileNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetLatestIntelRuleFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestIntelRuleFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLatestIntelRuleFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLatestIntelRuleFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestIntelRuleFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLatestIntelRuleFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLatestIntelRuleFileOK creates a GetLatestIntelRuleFileOK with default headers values
func NewGetLatestIntelRuleFileOK(writer io.Writer) *GetLatestIntelRuleFileOK {
	return &GetLatestIntelRuleFileOK{

		Payload: writer,
	}
}

/* GetLatestIntelRuleFileOK describes a response with status code 200, with default header values.

OK
*/
type GetLatestIntelRuleFileOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload io.Writer
}

func (o *GetLatestIntelRuleFileOK) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileOK  %+v", 200, o.Payload)
}
func (o *GetLatestIntelRuleFileOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetLatestIntelRuleFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestIntelRuleFileNotModified creates a GetLatestIntelRuleFileNotModified with default headers values
func NewGetLatestIntelRuleFileNotModified() *GetLatestIntelRuleFileNotModified {
	return &GetLatestIntelRuleFileNotModified{}
}

/* GetLatestIntelRuleFileNotModified describes a response with status code 304, with default header values.

Not Modified
*/
type GetLatestIntelRuleFileNotModified struct {
}

func (o *GetLatestIntelRuleFileNotModified) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileNotModified ", 304)
}

func (o *GetLatestIntelRuleFileNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLatestIntelRuleFileBadRequest creates a GetLatestIntelRuleFileBadRequest with default headers values
func NewGetLatestIntelRuleFileBadRequest() *GetLatestIntelRuleFileBadRequest {
	return &GetLatestIntelRuleFileBadRequest{}
}

/* GetLatestIntelRuleFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLatestIntelRuleFileBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetLatestIntelRuleFileBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileBadRequest  %+v", 400, o.Payload)
}
func (o *GetLatestIntelRuleFileBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetLatestIntelRuleFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestIntelRuleFileForbidden creates a GetLatestIntelRuleFileForbidden with default headers values
func NewGetLatestIntelRuleFileForbidden() *GetLatestIntelRuleFileForbidden {
	return &GetLatestIntelRuleFileForbidden{}
}

/* GetLatestIntelRuleFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLatestIntelRuleFileForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetLatestIntelRuleFileForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileForbidden  %+v", 403, o.Payload)
}
func (o *GetLatestIntelRuleFileForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetLatestIntelRuleFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLatestIntelRuleFileNotFound creates a GetLatestIntelRuleFileNotFound with default headers values
func NewGetLatestIntelRuleFileNotFound() *GetLatestIntelRuleFileNotFound {
	return &GetLatestIntelRuleFileNotFound{}
}

/* GetLatestIntelRuleFileNotFound describes a response with status code 404, with default header values.

Bad Request
*/
type GetLatestIntelRuleFileNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetLatestIntelRuleFileNotFound) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileNotFound  %+v", 404, o.Payload)
}
func (o *GetLatestIntelRuleFileNotFound) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetLatestIntelRuleFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestIntelRuleFileTooManyRequests creates a GetLatestIntelRuleFileTooManyRequests with default headers values
func NewGetLatestIntelRuleFileTooManyRequests() *GetLatestIntelRuleFileTooManyRequests {
	return &GetLatestIntelRuleFileTooManyRequests{}
}

/* GetLatestIntelRuleFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetLatestIntelRuleFileTooManyRequests struct {

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

func (o *GetLatestIntelRuleFileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetLatestIntelRuleFileTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetLatestIntelRuleFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLatestIntelRuleFileInternalServerError creates a GetLatestIntelRuleFileInternalServerError with default headers values
func NewGetLatestIntelRuleFileInternalServerError() *GetLatestIntelRuleFileInternalServerError {
	return &GetLatestIntelRuleFileInternalServerError{}
}

/* GetLatestIntelRuleFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLatestIntelRuleFileInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetLatestIntelRuleFileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] getLatestIntelRuleFileInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLatestIntelRuleFileInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetLatestIntelRuleFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestIntelRuleFileDefault creates a GetLatestIntelRuleFileDefault with default headers values
func NewGetLatestIntelRuleFileDefault(code int) *GetLatestIntelRuleFileDefault {
	return &GetLatestIntelRuleFileDefault{
		_statusCode: code,
	}
}

/* GetLatestIntelRuleFileDefault describes a response with status code -1, with default header values.

OK
*/
type GetLatestIntelRuleFileDefault struct {
	_statusCode int
}

// Code gets the status code for the get latest intel rule file default response
func (o *GetLatestIntelRuleFileDefault) Code() int {
	return o._statusCode
}

func (o *GetLatestIntelRuleFileDefault) Error() string {
	return fmt.Sprintf("[GET /intel/entities/rules-latest-files/v1][%d] GetLatestIntelRuleFile default ", o._statusCode)
}

func (o *GetLatestIntelRuleFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
