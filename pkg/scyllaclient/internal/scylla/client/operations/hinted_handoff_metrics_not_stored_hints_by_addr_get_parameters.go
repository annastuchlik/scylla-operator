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

// NewHintedHandoffMetricsNotStoredHintsByAddrGetParams creates a new HintedHandoffMetricsNotStoredHintsByAddrGetParams object
// with the default values initialized.
func NewHintedHandoffMetricsNotStoredHintsByAddrGetParams() *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	var ()
	return &HintedHandoffMetricsNotStoredHintsByAddrGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHintedHandoffMetricsNotStoredHintsByAddrGetParamsWithTimeout creates a new HintedHandoffMetricsNotStoredHintsByAddrGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHintedHandoffMetricsNotStoredHintsByAddrGetParamsWithTimeout(timeout time.Duration) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	var ()
	return &HintedHandoffMetricsNotStoredHintsByAddrGetParams{

		timeout: timeout,
	}
}

// NewHintedHandoffMetricsNotStoredHintsByAddrGetParamsWithContext creates a new HintedHandoffMetricsNotStoredHintsByAddrGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewHintedHandoffMetricsNotStoredHintsByAddrGetParamsWithContext(ctx context.Context) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	var ()
	return &HintedHandoffMetricsNotStoredHintsByAddrGetParams{

		Context: ctx,
	}
}

// NewHintedHandoffMetricsNotStoredHintsByAddrGetParamsWithHTTPClient creates a new HintedHandoffMetricsNotStoredHintsByAddrGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHintedHandoffMetricsNotStoredHintsByAddrGetParamsWithHTTPClient(client *http.Client) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	var ()
	return &HintedHandoffMetricsNotStoredHintsByAddrGetParams{
		HTTPClient: client,
	}
}

/*
HintedHandoffMetricsNotStoredHintsByAddrGetParams contains all the parameters to send to the API endpoint
for the hinted handoff metrics not stored hints by addr get operation typically these are written to a http.Request
*/
type HintedHandoffMetricsNotStoredHintsByAddrGetParams struct {

	/*Addr
	  The peer address

	*/
	Addr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) WithTimeout(timeout time.Duration) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) WithContext(ctx context.Context) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) WithHTTPClient(client *http.Client) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddr adds the addr to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) WithAddr(addr string) *HintedHandoffMetricsNotStoredHintsByAddrGetParams {
	o.SetAddr(addr)
	return o
}

// SetAddr adds the addr to the hinted handoff metrics not stored hints by addr get params
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) SetAddr(addr string) {
	o.Addr = addr
}

// WriteToRequest writes these params to a swagger request
func (o *HintedHandoffMetricsNotStoredHintsByAddrGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addr
	if err := r.SetPathParam("addr", o.Addr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
