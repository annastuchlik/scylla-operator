// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCommitlogMetricsPendingTasksGetParams creates a new CommitlogMetricsPendingTasksGetParams object
// with the default values initialized.
func NewCommitlogMetricsPendingTasksGetParams() *CommitlogMetricsPendingTasksGetParams {

	return &CommitlogMetricsPendingTasksGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCommitlogMetricsPendingTasksGetParamsWithTimeout creates a new CommitlogMetricsPendingTasksGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCommitlogMetricsPendingTasksGetParamsWithTimeout(timeout time.Duration) *CommitlogMetricsPendingTasksGetParams {

	return &CommitlogMetricsPendingTasksGetParams{

		timeout: timeout,
	}
}

// NewCommitlogMetricsPendingTasksGetParamsWithContext creates a new CommitlogMetricsPendingTasksGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCommitlogMetricsPendingTasksGetParamsWithContext(ctx context.Context) *CommitlogMetricsPendingTasksGetParams {

	return &CommitlogMetricsPendingTasksGetParams{

		Context: ctx,
	}
}

// NewCommitlogMetricsPendingTasksGetParamsWithHTTPClient creates a new CommitlogMetricsPendingTasksGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCommitlogMetricsPendingTasksGetParamsWithHTTPClient(client *http.Client) *CommitlogMetricsPendingTasksGetParams {

	return &CommitlogMetricsPendingTasksGetParams{
		HTTPClient: client,
	}
}

/*
CommitlogMetricsPendingTasksGetParams contains all the parameters to send to the API endpoint
for the commitlog metrics pending tasks get operation typically these are written to a http.Request
*/
type CommitlogMetricsPendingTasksGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the commitlog metrics pending tasks get params
func (o *CommitlogMetricsPendingTasksGetParams) WithTimeout(timeout time.Duration) *CommitlogMetricsPendingTasksGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the commitlog metrics pending tasks get params
func (o *CommitlogMetricsPendingTasksGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the commitlog metrics pending tasks get params
func (o *CommitlogMetricsPendingTasksGetParams) WithContext(ctx context.Context) *CommitlogMetricsPendingTasksGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the commitlog metrics pending tasks get params
func (o *CommitlogMetricsPendingTasksGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the commitlog metrics pending tasks get params
func (o *CommitlogMetricsPendingTasksGetParams) WithHTTPClient(client *http.Client) *CommitlogMetricsPendingTasksGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the commitlog metrics pending tasks get params
func (o *CommitlogMetricsPendingTasksGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CommitlogMetricsPendingTasksGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
