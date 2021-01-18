// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// CreateCSPMAwsAccountReader is a Reader for the CreateCSPMAwsAccount structure.
type CreateCSPMAwsAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCSPMAwsAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCSPMAwsAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCreateCSPMAwsAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCSPMAwsAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCSPMAwsAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateCSPMAwsAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCSPMAwsAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCSPMAwsAccountCreated creates a CreateCSPMAwsAccountCreated with default headers values
func NewCreateCSPMAwsAccountCreated() *CreateCSPMAwsAccountCreated {
	return &CreateCSPMAwsAccountCreated{}
}

/*CreateCSPMAwsAccountCreated handles this case with default header values.

Created
*/
type CreateCSPMAwsAccountCreated struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

func (o *CreateCSPMAwsAccountCreated) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-aws/entities/account/v1][%d] createCSPMAwsAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateCSPMAwsAccountCreated) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *CreateCSPMAwsAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCSPMAwsAccountMultiStatus creates a CreateCSPMAwsAccountMultiStatus with default headers values
func NewCreateCSPMAwsAccountMultiStatus() *CreateCSPMAwsAccountMultiStatus {
	return &CreateCSPMAwsAccountMultiStatus{}
}

/*CreateCSPMAwsAccountMultiStatus handles this case with default header values.

Multi-Status
*/
type CreateCSPMAwsAccountMultiStatus struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

func (o *CreateCSPMAwsAccountMultiStatus) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-aws/entities/account/v1][%d] createCSPMAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateCSPMAwsAccountMultiStatus) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *CreateCSPMAwsAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCSPMAwsAccountBadRequest creates a CreateCSPMAwsAccountBadRequest with default headers values
func NewCreateCSPMAwsAccountBadRequest() *CreateCSPMAwsAccountBadRequest {
	return &CreateCSPMAwsAccountBadRequest{}
}

/*CreateCSPMAwsAccountBadRequest handles this case with default header values.

Bad Request
*/
type CreateCSPMAwsAccountBadRequest struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

func (o *CreateCSPMAwsAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-aws/entities/account/v1][%d] createCSPMAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCSPMAwsAccountBadRequest) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *CreateCSPMAwsAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCSPMAwsAccountForbidden creates a CreateCSPMAwsAccountForbidden with default headers values
func NewCreateCSPMAwsAccountForbidden() *CreateCSPMAwsAccountForbidden {
	return &CreateCSPMAwsAccountForbidden{}
}

/*CreateCSPMAwsAccountForbidden handles this case with default header values.

Forbidden
*/
type CreateCSPMAwsAccountForbidden struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *CreateCSPMAwsAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-aws/entities/account/v1][%d] createCSPMAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateCSPMAwsAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateCSPMAwsAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCSPMAwsAccountTooManyRequests creates a CreateCSPMAwsAccountTooManyRequests with default headers values
func NewCreateCSPMAwsAccountTooManyRequests() *CreateCSPMAwsAccountTooManyRequests {
	return &CreateCSPMAwsAccountTooManyRequests{}
}

/*CreateCSPMAwsAccountTooManyRequests handles this case with default header values.

Too Many Requests
*/
type CreateCSPMAwsAccountTooManyRequests struct {
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

func (o *CreateCSPMAwsAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-aws/entities/account/v1][%d] createCSPMAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateCSPMAwsAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateCSPMAwsAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateCSPMAwsAccountInternalServerError creates a CreateCSPMAwsAccountInternalServerError with default headers values
func NewCreateCSPMAwsAccountInternalServerError() *CreateCSPMAwsAccountInternalServerError {
	return &CreateCSPMAwsAccountInternalServerError{}
}

/*CreateCSPMAwsAccountInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateCSPMAwsAccountInternalServerError struct {
	/*Request limit per minute.
	 */
	XRateLimitLimit int64
	/*The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

func (o *CreateCSPMAwsAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-cspm-aws/entities/account/v1][%d] createCSPMAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCSPMAwsAccountInternalServerError) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *CreateCSPMAwsAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}