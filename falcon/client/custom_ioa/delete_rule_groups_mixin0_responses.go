// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// DeleteRuleGroupsMixin0Reader is a Reader for the DeleteRuleGroupsMixin0 structure.
type DeleteRuleGroupsMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuleGroupsMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRuleGroupsMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRuleGroupsMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRuleGroupsMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRuleGroupsMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRuleGroupsMixin0Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRuleGroupsMixin0OK creates a DeleteRuleGroupsMixin0OK with default headers values
func NewDeleteRuleGroupsMixin0OK() *DeleteRuleGroupsMixin0OK {
	return &DeleteRuleGroupsMixin0OK{}
}

/*DeleteRuleGroupsMixin0OK handles this case with default header values.

OK
*/
type DeleteRuleGroupsMixin0OK struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DeleteRuleGroupsMixin0OK) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0OK  %+v", 200, o.Payload)
}

func (o *DeleteRuleGroupsMixin0OK) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0Forbidden creates a DeleteRuleGroupsMixin0Forbidden with default headers values
func NewDeleteRuleGroupsMixin0Forbidden() *DeleteRuleGroupsMixin0Forbidden {
	return &DeleteRuleGroupsMixin0Forbidden{}
}

/*DeleteRuleGroupsMixin0Forbidden handles this case with default header values.

Forbidden
*/
type DeleteRuleGroupsMixin0Forbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DeleteRuleGroupsMixin0Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRuleGroupsMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0NotFound creates a DeleteRuleGroupsMixin0NotFound with default headers values
func NewDeleteRuleGroupsMixin0NotFound() *DeleteRuleGroupsMixin0NotFound {
	return &DeleteRuleGroupsMixin0NotFound{}
}

/*DeleteRuleGroupsMixin0NotFound handles this case with default header values.

Not Found
*/
type DeleteRuleGroupsMixin0NotFound struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DeleteRuleGroupsMixin0NotFound) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0NotFound  %+v", 404, o.Payload)
}

func (o *DeleteRuleGroupsMixin0NotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0TooManyRequests creates a DeleteRuleGroupsMixin0TooManyRequests with default headers values
func NewDeleteRuleGroupsMixin0TooManyRequests() *DeleteRuleGroupsMixin0TooManyRequests {
	return &DeleteRuleGroupsMixin0TooManyRequests{}
}

/*DeleteRuleGroupsMixin0TooManyRequests handles this case with default header values.

Too Many Requests
*/
type DeleteRuleGroupsMixin0TooManyRequests struct {
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

func (o *DeleteRuleGroupsMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] deleteRuleGroupsMixin0TooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRuleGroupsMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRuleGroupsMixin0Default creates a DeleteRuleGroupsMixin0Default with default headers values
func NewDeleteRuleGroupsMixin0Default(code int) *DeleteRuleGroupsMixin0Default {
	return &DeleteRuleGroupsMixin0Default{
		_statusCode: code,
	}
}

/*DeleteRuleGroupsMixin0Default handles this case with default header values.

OK
*/
type DeleteRuleGroupsMixin0Default struct {
	_statusCode int

	Payload *models.MsaReplyMetaOnly
}

// Code gets the status code for the delete rule groups mixin0 default response
func (o *DeleteRuleGroupsMixin0Default) Code() int {
	return o._statusCode
}

func (o *DeleteRuleGroupsMixin0Default) Error() string {
	return fmt.Sprintf("[DELETE /ioarules/entities/rule-groups/v1][%d] delete-rule-groupsMixin0 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRuleGroupsMixin0Default) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRuleGroupsMixin0Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}