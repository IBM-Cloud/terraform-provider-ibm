// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package restapi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// ResourceIBMRestApiRequest manages an object that is created and destroyed
// through an arbitrary IBM Cloud REST API. It exists for services that do not
// yet have a dedicated Terraform resource.
func ResourceIBMRestApiRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMRestApiRequestCreate,
		ReadContext:   resourceIBMRestApiRequestRead,
		UpdateContext: resourceIBMRestApiRequestUpdate,
		DeleteContext: resourceIBMRestApiRequestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceIBMRestApiRequestImport,
		},

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The URL the object is created against. For a collection style API this is the collection URL, for example https://resource-controller.cloud.ibm.com/v2/resource_instances.",
			},
			"create_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "POST",
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"POST", "PUT", "PATCH"}, false),
				Description:  "The HTTP method used to create the object. Defaults to POST.",
			},
			"read_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "GET",
				ValidateFunc: validation.StringInSlice([]string{"GET", "POST"}, false),
				Description:  "The HTTP method used to read the object back. Defaults to GET.",
			},
			"update_method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "PUT",
				ValidateFunc: validation.StringInSlice([]string{"PUT", "PATCH", "POST"}, false),
				Description:  "The HTTP method used to update the object. Defaults to PUT.",
			},
			"destroy_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"DELETE", "POST"}, false),
				Default:      "DELETE",
				Description:  "The HTTP method used to destroy the object. Defaults to DELETE.",
			},
			"request_body": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "The JSON body sent when the object is created, and when it is updated unless update_request_body is set.",
			},
			"update_request_body": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "The JSON body sent when the object is updated. When it is not set, request_body is sent instead.",
			},
			"destroy_request_body": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "The JSON body sent when the object is destroyed. Most APIs do not need one.",
			},
			"id_attribute": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Dot separated path to the identifier inside the create response, for example id or metadata.guid. When it is set, the object URL used for read, update and destroy is the create URL with the identifier appended. When it is not set, the create URL is used unchanged.",
			},
			"object_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL used to read, update and destroy the object.",
			},
			"object_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The identifier read out of the create response using id_attribute. Empty when id_attribute is not set.",
			},
			"headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Additional HTTP headers sent with every request.",
			},
			"sensitive_headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Sensitive:   true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Additional HTTP headers sent with every request, kept out of the Terraform output. Values here override values with the same key in headers.",
			},
			"query_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Query parameters appended to every request URL.",
			},
			"accept_status_codes": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Status codes outside the 2xx range that are treated as success instead of an error.",
			},
			"use_iam_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the IBM Cloud credentials configured on the provider are used to authenticate the request. Set it to false when the target API is not IAM protected or when an Authorization header is supplied through sensitive_headers.",
			},
			"read_on_create": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the object is read back after it is created. Set it to false for APIs that do not expose a read operation, in which case the create response is kept as the object state.",
			},
			"ignore_read_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether a failing read is ignored instead of failing the run. The last known response body is kept in state.",
			},
			"timeout_seconds": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      60,
				ValidateFunc: validation.IntBetween(1, 3600),
				Description:  "How long a single HTTP request may take before it is cancelled.",
			},
			"response_body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The body of the most recent read response, or of the create response when read_on_create is false.",
			},
			"create_response_body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The body returned by the create request. Useful for APIs that only return generated values once.",
			},
			"response_status_code": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status code of the most recent response.",
			},
			"response_headers": {
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The headers of the most recent response.",
			},
		},
	}
}

// requestHeaders merges headers and sensitive_headers, with the sensitive map
// taking precedence.
func requestHeaders(d *schema.ResourceData) map[string]string {
	headers := stringMap(d.Get("headers"))
	for key, value := range stringMap(d.Get("sensitive_headers")) {
		headers[key] = value
	}
	return headers
}

func baseRequest(d *schema.ResourceData) restRequest {
	return restRequest{
		Headers:     requestHeaders(d),
		QueryParams: stringMap(d.Get("query_params")),
		Timeout:     time.Duration(d.Get("timeout_seconds").(int)) * time.Second,
		UseIAMAuth:  d.Get("use_iam_auth").(bool),
	}
}

func setResponseAttributes(d *schema.ResourceData, resp *restResponse) error {
	if err := d.Set("response_body", resp.Body); err != nil {
		return fmt.Errorf("error setting response_body: %w", err)
	}
	if err := d.Set("response_status_code", resp.StatusCode); err != nil {
		return fmt.Errorf("error setting response_status_code: %w", err)
	}
	if err := d.Set("response_headers", resp.Headers); err != nil {
		return fmt.Errorf("error setting response_headers: %w", err)
	}
	return nil
}

func resourceIBMRestApiRequestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createURL := d.Get("url").(string)
	method := d.Get("create_method").(string)

	req := baseRequest(d)
	req.Method = method
	req.URL = createURL
	req.Body = d.Get("request_body").(string)

	resp, err := doRequest(ctx, meta, req)
	if err != nil {
		return diag.FromErr(err)
	}
	accepted := intList(d.Get("accept_status_codes"))
	if !resp.isSuccess(accepted) {
		return diag.FromErr(responseError(method, createURL, resp))
	}

	objectURL := createURL
	objectID := ""
	if idAttribute := d.Get("id_attribute").(string); idAttribute != "" {
		objectID, err = extractID(resp.Body, idAttribute)
		if err != nil {
			// The object was created remotely, so surface the identifier
			// problem rather than silently orphaning it.
			return diag.FromErr(fmt.Errorf("the object was created by %s %s but its identifier could not be determined, it may need to be removed manually: %w", method, createURL, err))
		}
		objectURL, err = joinURL(createURL, objectID)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(objectURL)
	if err := d.Set("object_url", objectURL); err != nil {
		return diag.FromErr(fmt.Errorf("error setting object_url: %w", err))
	}
	if err := d.Set("object_id", objectID); err != nil {
		return diag.FromErr(fmt.Errorf("error setting object_id: %w", err))
	}
	if err := d.Set("create_response_body", resp.Body); err != nil {
		return diag.FromErr(fmt.Errorf("error setting create_response_body: %w", err))
	}
	if err := setResponseAttributes(d, resp); err != nil {
		return diag.FromErr(err)
	}

	if !d.Get("read_on_create").(bool) {
		return nil
	}
	return resourceIBMRestApiRequestRead(ctx, d, meta)
}

func resourceIBMRestApiRequestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	objectURL := d.Get("object_url").(string)
	if objectURL == "" {
		objectURL = d.Id()
	}
	method := d.Get("read_method").(string)

	req := baseRequest(d)
	req.Method = method
	req.URL = objectURL

	resp, err := doRequest(ctx, meta, req)
	if err != nil {
		if d.Get("ignore_read_errors").(bool) {
			return nil
		}
		return diag.FromErr(err)
	}

	if resp.StatusCode == 404 || resp.StatusCode == 410 {
		// The object is gone remotely, let Terraform plan a replacement.
		d.SetId("")
		return nil
	}

	if !resp.isSuccess(intList(d.Get("accept_status_codes"))) {
		if d.Get("ignore_read_errors").(bool) {
			return nil
		}
		return diag.FromErr(responseError(method, objectURL, resp))
	}

	if err := setResponseAttributes(d, resp); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceIBMRestApiRequestUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Attribute only changes, such as a different ignore_read_errors, do not
	// need a call to the API.
	if !d.HasChanges("request_body", "update_request_body", "headers", "sensitive_headers", "query_params") {
		return resourceIBMRestApiRequestRead(ctx, d, meta)
	}

	objectURL := d.Get("object_url").(string)
	if objectURL == "" {
		objectURL = d.Id()
	}
	method := d.Get("update_method").(string)

	body := d.Get("update_request_body").(string)
	if body == "" {
		body = d.Get("request_body").(string)
	}

	req := baseRequest(d)
	req.Method = method
	req.URL = objectURL
	req.Body = body

	resp, err := doRequest(ctx, meta, req)
	if err != nil {
		return diag.FromErr(err)
	}
	if !resp.isSuccess(intList(d.Get("accept_status_codes"))) {
		return diag.FromErr(responseError(method, objectURL, resp))
	}
	if err := setResponseAttributes(d, resp); err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMRestApiRequestRead(ctx, d, meta)
}

func resourceIBMRestApiRequestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	objectURL := d.Get("object_url").(string)
	if objectURL == "" {
		objectURL = d.Id()
	}
	method := d.Get("destroy_method").(string)

	req := baseRequest(d)
	req.Method = method
	req.URL = objectURL
	req.Body = d.Get("destroy_request_body").(string)

	resp, err := doRequest(ctx, meta, req)
	if err != nil {
		return diag.FromErr(err)
	}

	// An object that is already gone is not an error for a destroy.
	if resp.StatusCode == 404 || resp.StatusCode == 410 {
		d.SetId("")
		return nil
	}

	if !resp.isSuccess(intList(d.Get("accept_status_codes"))) {
		return diag.FromErr(responseError(method, objectURL, resp))
	}

	d.SetId("")
	return nil
}

// resourceIBMRestApiRequestImport accepts either the object URL on its own, or
// the create URL and the object URL separated by a comma. The two part form is
// needed when id_attribute is set, because the create URL is the collection
// URL and cannot be derived from the object URL.
func resourceIBMRestApiRequestImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), ",")
	for index := range parts {
		parts[index] = strings.TrimSpace(parts[index])
	}

	var createURL, objectURL string
	switch len(parts) {
	case 1:
		createURL = parts[0]
		objectURL = parts[0]
	case 2:
		createURL = parts[0]
		objectURL = parts[1]
	default:
		return nil, fmt.Errorf("invalid import id %q, expected either <object_url> or <url>,<object_url>", d.Id())
	}
	if createURL == "" || objectURL == "" {
		return nil, fmt.Errorf("invalid import id %q, expected either <object_url> or <url>,<object_url>", d.Id())
	}
	if _, err := buildURL(objectURL, nil); err != nil {
		return nil, err
	}

	d.SetId(objectURL)
	if err := d.Set("url", createURL); err != nil {
		return nil, fmt.Errorf("error setting url: %w", err)
	}
	if err := d.Set("object_url", objectURL); err != nil {
		return nil, fmt.Errorf("error setting object_url: %w", err)
	}
	if createURL != objectURL {
		if err := d.Set("object_id", strings.TrimPrefix(objectURL, strings.TrimSuffix(createURL, "/")+"/")); err != nil {
			return nil, fmt.Errorf("error setting object_id: %w", err)
		}
	}
	return []*schema.ResourceData{d}, nil
}
