// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new real time response API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for real time response API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BatchActiveResponderCmd(params *BatchActiveResponderCmdParams) (*BatchActiveResponderCmdCreated, error)

	BatchCmd(params *BatchCmdParams) (*BatchCmdCreated, error)

	BatchGetCmd(params *BatchGetCmdParams) (*BatchGetCmdCreated, error)

	BatchGetCmdStatus(params *BatchGetCmdStatusParams) (*BatchGetCmdStatusOK, error)

	BatchInitSessions(params *BatchInitSessionsParams) (*BatchInitSessionsCreated, error)

	BatchRefreshSessions(params *BatchRefreshSessionsParams) (*BatchRefreshSessionsCreated, error)

	RTRAggregateSessions(params *RTRAggregateSessionsParams) (*RTRAggregateSessionsOK, error)

	RTRCheckActiveResponderCommandStatus(params *RTRCheckActiveResponderCommandStatusParams) (*RTRCheckActiveResponderCommandStatusOK, error)

	RTRCheckCommandStatus(params *RTRCheckCommandStatusParams) (*RTRCheckCommandStatusOK, error)

	RTRDeleteFile(params *RTRDeleteFileParams) (*RTRDeleteFileNoContent, error)

	RTRDeleteQueuedSession(params *RTRDeleteQueuedSessionParams) (*RTRDeleteQueuedSessionNoContent, error)

	RTRDeleteSession(params *RTRDeleteSessionParams) (*RTRDeleteSessionNoContent, error)

	RTRExecuteActiveResponderCommand(params *RTRExecuteActiveResponderCommandParams) (*RTRExecuteActiveResponderCommandCreated, error)

	RTRExecuteCommand(params *RTRExecuteCommandParams) (*RTRExecuteCommandCreated, error)

	RTRGetExtractedFileContents(params *RTRGetExtractedFileContentsParams) (*RTRGetExtractedFileContentsOK, error)

	RTRInitSession(params *RTRInitSessionParams) (*RTRInitSessionCreated, error)

	RTRListAllSessions(params *RTRListAllSessionsParams) (*RTRListAllSessionsOK, error)

	RTRListFiles(params *RTRListFilesParams) (*RTRListFilesOK, error)

	RTRListQueuedSessions(params *RTRListQueuedSessionsParams) (*RTRListQueuedSessionsOK, error)

	RTRListSessions(params *RTRListSessionsParams) (*RTRListSessionsOK, error)

	RTRPulseSession(params *RTRPulseSessionParams) (*RTRPulseSessionCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BatchActiveResponderCmd batches executes a r t r active responder command across the hosts mapped to the given batch ID
*/
func (a *Client) BatchActiveResponderCmd(params *BatchActiveResponderCmdParams) (*BatchActiveResponderCmdCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchActiveResponderCmdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BatchActiveResponderCmd",
		Method:             "POST",
		PathPattern:        "/real-time-response/combined/batch-active-responder-command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchActiveResponderCmdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchActiveResponderCmdCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchActiveResponderCmd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchCmd batches executes a r t r read only command across the hosts mapped to the given batch ID
*/
func (a *Client) BatchCmd(params *BatchCmdParams) (*BatchCmdCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchCmdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BatchCmd",
		Method:             "POST",
		PathPattern:        "/real-time-response/combined/batch-command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchCmdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchCmdCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchCmd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchGetCmd batches executes get command across hosts to retrieve files after this call is made g e t real time response combined batch get command v1 is used to query for the results
*/
func (a *Client) BatchGetCmd(params *BatchGetCmdParams) (*BatchGetCmdCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchGetCmdParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BatchGetCmd",
		Method:             "POST",
		PathPattern:        "/real-time-response/combined/batch-get-command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchGetCmdReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchGetCmdCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchGetCmd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchGetCmdStatus retrieves the status of the specified batch get command will return successful files when they are finished processing
*/
func (a *Client) BatchGetCmdStatus(params *BatchGetCmdStatusParams) (*BatchGetCmdStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchGetCmdStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BatchGetCmdStatus",
		Method:             "GET",
		PathPattern:        "/real-time-response/combined/batch-get-command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchGetCmdStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchGetCmdStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BatchGetCmdStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  BatchInitSessions batches initialize a r t r session on multiple hosts before any r t r commands can be used an active session is needed on the host
*/
func (a *Client) BatchInitSessions(params *BatchInitSessionsParams) (*BatchInitSessionsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchInitSessionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BatchInitSessions",
		Method:             "POST",
		PathPattern:        "/real-time-response/combined/batch-init-session/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchInitSessionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchInitSessionsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchInitSessions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BatchRefreshSessions batches refresh a r t r session on multiple hosts r t r sessions will expire after 10 minutes unless refreshed
*/
func (a *Client) BatchRefreshSessions(params *BatchRefreshSessionsParams) (*BatchRefreshSessionsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBatchRefreshSessionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BatchRefreshSessions",
		Method:             "POST",
		PathPattern:        "/real-time-response/combined/batch-refresh-session/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BatchRefreshSessionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BatchRefreshSessionsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BatchRefreshSessions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRAggregateSessions gets aggregates on session data
*/
func (a *Client) RTRAggregateSessions(params *RTRAggregateSessionsParams) (*RTRAggregateSessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRAggregateSessionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-AggregateSessions",
		Method:             "POST",
		PathPattern:        "/real-time-response/aggregates/sessions/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRAggregateSessionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRAggregateSessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRAggregateSessionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRCheckActiveResponderCommandStatus gets status of an executed active responder command on a single host
*/
func (a *Client) RTRCheckActiveResponderCommandStatus(params *RTRCheckActiveResponderCommandStatusParams) (*RTRCheckActiveResponderCommandStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRCheckActiveResponderCommandStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-CheckActiveResponderCommandStatus",
		Method:             "GET",
		PathPattern:        "/real-time-response/entities/active-responder-command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRCheckActiveResponderCommandStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRCheckActiveResponderCommandStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRCheckActiveResponderCommandStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRCheckCommandStatus gets status of an executed command on a single host
*/
func (a *Client) RTRCheckCommandStatus(params *RTRCheckCommandStatusParams) (*RTRCheckCommandStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRCheckCommandStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-CheckCommandStatus",
		Method:             "GET",
		PathPattern:        "/real-time-response/entities/command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRCheckCommandStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRCheckCommandStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRCheckCommandStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRDeleteFile deletes a r t r session file
*/
func (a *Client) RTRDeleteFile(params *RTRDeleteFileParams) (*RTRDeleteFileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRDeleteFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-DeleteFile",
		Method:             "DELETE",
		PathPattern:        "/real-time-response/entities/file/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRDeleteFileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRDeleteFileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-DeleteFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRDeleteQueuedSession deletes a queued session command
*/
func (a *Client) RTRDeleteQueuedSession(params *RTRDeleteQueuedSessionParams) (*RTRDeleteQueuedSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRDeleteQueuedSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-DeleteQueuedSession",
		Method:             "DELETE",
		PathPattern:        "/real-time-response/entities/queued-sessions/command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRDeleteQueuedSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRDeleteQueuedSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-DeleteQueuedSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRDeleteSession deletes a session
*/
func (a *Client) RTRDeleteSession(params *RTRDeleteSessionParams) (*RTRDeleteSessionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRDeleteSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-DeleteSession",
		Method:             "DELETE",
		PathPattern:        "/real-time-response/entities/sessions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRDeleteSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRDeleteSessionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-DeleteSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRExecuteActiveResponderCommand executes an active responder command on a single host
*/
func (a *Client) RTRExecuteActiveResponderCommand(params *RTRExecuteActiveResponderCommandParams) (*RTRExecuteActiveResponderCommandCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRExecuteActiveResponderCommandParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-ExecuteActiveResponderCommand",
		Method:             "POST",
		PathPattern:        "/real-time-response/entities/active-responder-command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRExecuteActiveResponderCommandReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRExecuteActiveResponderCommandCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-ExecuteActiveResponderCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRExecuteCommand executes a command on a single host
*/
func (a *Client) RTRExecuteCommand(params *RTRExecuteCommandParams) (*RTRExecuteCommandCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRExecuteCommandParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-ExecuteCommand",
		Method:             "POST",
		PathPattern:        "/real-time-response/entities/command/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRExecuteCommandReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRExecuteCommandCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-ExecuteCommand: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRGetExtractedFileContents gets r t r extracted file contents for specified session and sha256
*/
func (a *Client) RTRGetExtractedFileContents(params *RTRGetExtractedFileContentsParams) (*RTRGetExtractedFileContentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRGetExtractedFileContentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-GetExtractedFileContents",
		Method:             "GET",
		PathPattern:        "/real-time-response/entities/extracted-file-contents/v1",
		ProducesMediaTypes: []string{"application/json", "application/x-7z-compressed"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRGetExtractedFileContentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRGetExtractedFileContentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRGetExtractedFileContentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRInitSession initializes a new session with the r t r cloud
*/
func (a *Client) RTRInitSession(params *RTRInitSessionParams) (*RTRInitSessionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRInitSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-InitSession",
		Method:             "POST",
		PathPattern:        "/real-time-response/entities/sessions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRInitSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRInitSessionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-InitSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RTRListAllSessions gets a list of session ids
*/
func (a *Client) RTRListAllSessions(params *RTRListAllSessionsParams) (*RTRListAllSessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRListAllSessionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-ListAllSessions",
		Method:             "GET",
		PathPattern:        "/real-time-response/queries/sessions/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRListAllSessionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRListAllSessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRListAllSessionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRListFiles gets a list of files for the specified r t r session
*/
func (a *Client) RTRListFiles(params *RTRListFilesParams) (*RTRListFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRListFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-ListFiles",
		Method:             "GET",
		PathPattern:        "/real-time-response/entities/file/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRListFilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRListFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRListFilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRListQueuedSessions gets queued session metadata by session ID
*/
func (a *Client) RTRListQueuedSessions(params *RTRListQueuedSessionsParams) (*RTRListQueuedSessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRListQueuedSessionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-ListQueuedSessions",
		Method:             "POST",
		PathPattern:        "/real-time-response/entities/queued-sessions/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRListQueuedSessionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRListQueuedSessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRListQueuedSessionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRListSessions gets session metadata by session id
*/
func (a *Client) RTRListSessions(params *RTRListSessionsParams) (*RTRListSessionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRListSessionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-ListSessions",
		Method:             "POST",
		PathPattern:        "/real-time-response/entities/sessions/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRListSessionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRListSessionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RTRListSessionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RTRPulseSession refreshes a session timeout on a single host
*/
func (a *Client) RTRPulseSession(params *RTRPulseSessionParams) (*RTRPulseSessionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRTRPulseSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RTR-PulseSession",
		Method:             "POST",
		PathPattern:        "/real-time-response/entities/refresh-session/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RTRPulseSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RTRPulseSessionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RTR-PulseSession: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}