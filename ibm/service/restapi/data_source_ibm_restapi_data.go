// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package restapi

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// DataSourceIBMRestApiData reads data from an arbitrary IBM Cloud REST API.
func DataSourceIBMRestApiData() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMRestApiDataRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL that is read, for example https://resource-controller.cloud.ibm.com/v2/resource_instances.",
			},
			"method": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "GET",
				ValidateFunc: validation.StringInSlice([]string{"GET", "POST", "HEAD"}, false),
				Description:  "The HTTP method used to read the data. Defaults to GET. POST is allowed for search style APIs that take a query body.",
			},
			"request_body": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "The JSON body sent with the request. Only useful when method is POST.",
			},
			"headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Additional HTTP headers sent with the request.",
			},
			"sensitive_headers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Sensitive:   true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Additional HTTP headers sent with the request, kept out of the Terraform output. Values here override values with the same key in headers.",
			},
			"query_params": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Query parameters appended to the request URL.",
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
			"timeout_seconds": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      60,
				ValidateFunc: validation.IntBetween(1, 3600),
				Description:  "How long the HTTP request may take before it is cancelled.",
			},
			"response_body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The raw response body. Use the jsondecode function to work with it in the configuration.",
			},
			"response_status_code": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The status code of the response.",
			},
			"response_headers": {
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The headers of the response.",
			},
			"is_json": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the response body parsed as JSON. Configurations should check it before calling jsondecode on response_body.",
			},
		},
	}
}

func dataSourceIBMRestApiDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	requestURL := d.Get("url").(string)
	method := d.Get("method").(string)

	headers := stringMap(d.Get("headers"))
	for key, value := range stringMap(d.Get("sensitive_headers")) {
		headers[key] = value
	}

	req := restRequest{
		Method:      method,
		URL:         requestURL,
		Body:        d.Get("request_body").(string),
		Headers:     headers,
		QueryParams: stringMap(d.Get("query_params")),
		Timeout:     time.Duration(d.Get("timeout_seconds").(int)) * time.Second,
		UseIAMAuth:  d.Get("use_iam_auth").(bool),
	}

	resp, err := doRequest(ctx, meta, req)
	if err != nil {
		return diag.FromErr(err)
	}
	if !resp.isSuccess(intList(d.Get("accept_status_codes"))) {
		return diag.FromErr(responseError(method, requestURL, resp))
	}

	// The identifier is derived from the request so that two data sources
	// reading different URLs never collide, and so that it stays stable
	// between runs.
	fullURL, err := buildURL(requestURL, req.QueryParams)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s/%x", fullURL, sha256.Sum256([]byte(method+req.Body))))

	if err := d.Set("response_body", resp.Body); err != nil {
		return diag.FromErr(fmt.Errorf("error setting response_body: %w", err))
	}
	if err := d.Set("response_status_code", resp.StatusCode); err != nil {
		return diag.FromErr(fmt.Errorf("error setting response_status_code: %w", err))
	}
	if err := d.Set("response_headers", resp.Headers); err != nil {
		return diag.FromErr(fmt.Errorf("error setting response_headers: %w", err))
	}

	isJSON := false
	if resp.Body != "" {
		var decoded interface{}
		isJSON = json.Unmarshal([]byte(resp.Body), &decoded) == nil
	}
	if err := d.Set("is_json", isJSON); err != nil {
		return diag.FromErr(fmt.Errorf("error setting is_json: %w", err))
	}

	return nil
}
