// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

// Package rhaiiv1 is a small client for the Red Hat AI Inference on IBM Cloud API
// (https://cloud.ibm.com/docs/apis/inference).
//
// The public API spec marks the inference operations with x-sdk-exclude, so no
// published Go SDK covers them. This package follows the layout of the generated
// IBM Cloud Go SDKs (go-sdk-core BaseService, options structs, WithContext
// variants) so that it can be swapped for an official SDK later with few changes.
// Only the read operations that the Terraform provider needs are implemented.
package rhaiiv1

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
)

const (
	// DefaultServiceURL is the default URL of the Red Hat AI Inference API.
	DefaultServiceURL = "https://us-east.rhai.ibm.com/v1"
	// DefaultServiceName is the default key used to find external configuration information.
	DefaultServiceName = "rhaii"
)

// RhaiiV1 : Red Hat AI Inference on IBM Cloud API.
type RhaiiV1 struct {
	Service *core.BaseService
}

// RhaiiV1Options : Service options
type RhaiiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewRhaiiV1UsingExternalConfig : constructs an instance of RhaiiV1 with passed in options and external configuration.
func NewRhaiiV1UsingExternalConfig(options *RhaiiV1Options) (rhaii *RhaiiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	rhaii, err = NewRhaiiV1(options)
	if err != nil {
		return
	}

	err = rhaii.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = rhaii.Service.SetServiceURL(options.URL)
	}
	return
}

// NewRhaiiV1 : constructs an instance of RhaiiV1 with passed in options.
func NewRhaiiV1(options *RhaiiV1Options) (service *RhaiiV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &RhaiiV1{
		Service: baseService,
	}
	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region.
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"us-east": "https://us-east.rhai.ibm.com/v1",
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "rhaii" suitable for processing requests.
func (rhaii *RhaiiV1) Clone() *RhaiiV1 {
	if core.IsNil(rhaii) {
		return nil
	}
	clone := *rhaii
	clone.Service = rhaii.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL.
func (rhaii *RhaiiV1) SetServiceURL(url string) error {
	return rhaii.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL.
func (rhaii *RhaiiV1) GetServiceURL() string {
	return rhaii.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request.
func (rhaii *RhaiiV1) SetDefaultHeaders(headers http.Header) {
	rhaii.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field.
func (rhaii *RhaiiV1) SetEnableGzipCompression(enableGzip bool) {
	rhaii.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field.
func (rhaii *RhaiiV1) GetEnableGzipCompression() bool {
	return rhaii.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
func (rhaii *RhaiiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	rhaii.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (rhaii *RhaiiV1) DisableRetries() {
	rhaii.Service.DisableRetries()
}

// ListInferenceModels : List inference models
// Lists the currently available inference models of a project.
func (rhaii *RhaiiV1) ListInferenceModels(listInferenceModelsOptions *ListInferenceModelsOptions) (result *InferenceModelCollection, response *core.DetailedResponse, err error) {
	return rhaii.ListInferenceModelsWithContext(context.Background(), listInferenceModelsOptions)
}

// ListInferenceModelsWithContext is an alternate form of the ListInferenceModels method which supports a Context parameter
func (rhaii *RhaiiV1) ListInferenceModelsWithContext(ctx context.Context, listInferenceModelsOptions *ListInferenceModelsOptions) (result *InferenceModelCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listInferenceModelsOptions, "listInferenceModelsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listInferenceModelsOptions, "listInferenceModelsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *listInferenceModelsOptions.ProjectID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rhaii.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rhaii.Service.Options.URL, `/projects/{project_id}/inference/models`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listInferenceModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	result = &InferenceModelCollection{}
	response, err = rhaii.Service.Request(request, result)
	if err != nil {
		result = nil
	}
	return
}

// GetInferenceModel : Get an inference model
// Retrieves detailed information about a specific inference model.
func (rhaii *RhaiiV1) GetInferenceModel(getInferenceModelOptions *GetInferenceModelOptions) (result *InferenceModel, response *core.DetailedResponse, err error) {
	return rhaii.GetInferenceModelWithContext(context.Background(), getInferenceModelOptions)
}

// GetInferenceModelWithContext is an alternate form of the GetInferenceModel method which supports a Context parameter
func (rhaii *RhaiiV1) GetInferenceModelWithContext(ctx context.Context, getInferenceModelOptions *GetInferenceModelOptions) (result *InferenceModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getInferenceModelOptions, "getInferenceModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getInferenceModelOptions, "getInferenceModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"project_id": *getInferenceModelOptions.ProjectID,
		"model":      *getInferenceModelOptions.Model,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = rhaii.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(rhaii.Service.Options.URL, `/projects/{project_id}/inference/models/{model}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getInferenceModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	result = &InferenceModel{}
	response, err = rhaii.Service.Request(request, result)
	if err != nil {
		result = nil
	}
	return
}

// ListInferenceModelsOptions : The ListInferenceModels options.
type ListInferenceModelsOptions struct {
	// The Red Hat AI Inference project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListInferenceModelsOptions : Instantiate ListInferenceModelsOptions
func (*RhaiiV1) NewListInferenceModelsOptions(projectID string) *ListInferenceModelsOptions {
	return &ListInferenceModelsOptions{
		ProjectID: core.StringPtr(projectID),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *ListInferenceModelsOptions) SetProjectID(projectID string) *ListInferenceModelsOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListInferenceModelsOptions) SetHeaders(param map[string]string) *ListInferenceModelsOptions {
	options.Headers = param
	return options
}

// GetInferenceModelOptions : The GetInferenceModel options.
type GetInferenceModelOptions struct {
	// The Red Hat AI Inference project ID.
	ProjectID *string `json:"project_id" validate:"required,ne="`

	// The ID of the model.
	Model *string `json:"model" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetInferenceModelOptions : Instantiate GetInferenceModelOptions
func (*RhaiiV1) NewGetInferenceModelOptions(projectID string, model string) *GetInferenceModelOptions {
	return &GetInferenceModelOptions{
		ProjectID: core.StringPtr(projectID),
		Model:     core.StringPtr(model),
	}
}

// SetProjectID : Allow user to set ProjectID
func (_options *GetInferenceModelOptions) SetProjectID(projectID string) *GetInferenceModelOptions {
	_options.ProjectID = core.StringPtr(projectID)
	return _options
}

// SetModel : Allow user to set Model
func (_options *GetInferenceModelOptions) SetModel(model string) *GetInferenceModelOptions {
	_options.Model = core.StringPtr(model)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetInferenceModelOptions) SetHeaders(param map[string]string) *GetInferenceModelOptions {
	options.Headers = param
	return options
}

// InferenceModelCollection : Response of the list inference models operation.
type InferenceModelCollection struct {
	// The inference models.
	Data []InferenceModelSummary `json:"data"`

	// The object type, "list".
	Object *string `json:"object,omitempty"`
}

// InferenceModelSummary : An inference model in a list (OpenAI model object with IBM metadata).
type InferenceModelSummary struct {
	// The ID of the model.
	ID *string `json:"id"`

	// The object type, "model".
	Object *string `json:"object,omitempty"`

	// The Unix timestamp in seconds when the model was created.
	Created *int64 `json:"created,omitempty"`

	// The owner of the model.
	OwnedBy *string `json:"owned_by,omitempty"`

	// IBM metadata of the model.
	CustomMetadata *InferenceModelMetadata `json:"custom_metadata,omitempty"`
}

// InferenceModel : Details of an inference model (Llama Stack model object with IBM metadata).
type InferenceModel struct {
	// Unique identifier of the model.
	Identifier *string `json:"identifier"`

	// Unique identifier of the model in the provider.
	ProviderResourceID *string `json:"provider_resource_id,omitempty"`

	// ID of the provider that owns the model.
	ProviderID *string `json:"provider_id,omitempty"`

	// The resource type, "model".
	Type *string `json:"type,omitempty"`

	// The model type, for example "llm", "embedding" or "rerank".
	ModelType *string `json:"model_type,omitempty"`

	// IBM metadata of the model.
	Metadata *InferenceModelMetadata `json:"metadata,omitempty"`
}

// InferenceModelMetadata : IBM metadata of an inference model.
type InferenceModelMetadata struct {
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	State       *string `json:"state,omitempty"`
	Status      *string `json:"status,omitempty"`
	TotalParams *int64  `json:"total_params,omitempty"`

	// Model configuration with architecture details and hyperparameters. Free form.
	Config map[string]interface{} `json:"config,omitempty"`

	// Model card in markdown format.
	ModelCard *string `json:"model_card,omitempty"`

	// Pricing measures of the model.
	// Note: the published spec defines this under the key "validation" in
	// IBMMetadataPricingField, while the service and the spec examples return "pricing".
	Pricing *ModelPricing `json:"pricing,omitempty"`

	// Validation information of a custom model.
	Validation *ModelValidation `json:"validation,omitempty"`

	// Source of a custom model (huggingface, ibmcos, oci or s3). Free form.
	Source map[string]interface{} `json:"source,omitempty"`
}

// ModelPricing : Model pricing information.
type ModelPricing struct {
	InputMeasure  *string `json:"input_measure,omitempty"`
	OutputMeasure *string `json:"output_measure,omitempty"`
}

// ModelValidation : Model validation information.
type ModelValidation struct {
	Status      *string `json:"status,omitempty"`
	Error       *string `json:"error,omitempty"`
	CompletedAt *string `json:"completed_at,omitempty"`
}

// ProjectServiceURL returns the base URL of the API for a single project, for
// example https://us-east.rhai.ibm.com/v1/projects/<project_id>.
func ProjectServiceURL(serviceURL, projectID string) string {
	return strings.TrimSuffix(serviceURL, "/") + "/projects/" + projectID
}
