// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package restapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

const (
	defaultRequestTimeout = 60 * time.Second
	maxLoggedBodyLength   = 2048
)

// restRequest describes a single arbitrary REST call issued by the
// ibm_restapi_request resource or the ibm_restapi_data data source.
type restRequest struct {
	Method      string
	URL         string
	Body        string
	Headers     map[string]string
	QueryParams map[string]string
	Timeout     time.Duration
	UseIAMAuth  bool
}

// restResponse holds the outcome of a restRequest.
type restResponse struct {
	StatusCode int
	Body       string
	Headers    map[string]string
}

// isSuccess reports whether the status code is a 2xx, or is explicitly
// listed by the practitioner in accept_status_codes.
func (r *restResponse) isSuccess(accepted []int) bool {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return true
	}
	for _, code := range accepted {
		if code == r.StatusCode {
			return true
		}
	}
	return false
}

// buildURL joins the request URL with any additional query parameters. Query
// parameters already present on the URL are preserved; entries in
// query_params win when the same key is given twice.
func buildURL(rawURL string, params map[string]string) (string, error) {
	if strings.TrimSpace(rawURL) == "" {
		return "", fmt.Errorf("url must not be empty")
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid url %q: %w", rawURL, err)
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("invalid url %q: scheme must be http or https", rawURL)
	}
	if parsed.Host == "" {
		return "", fmt.Errorf("invalid url %q: missing host", rawURL)
	}
	if len(params) > 0 {
		query := parsed.Query()
		// Sort the keys so the generated URL is stable across runs.
		keys := make([]string, 0, len(params))
		for key := range params {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			query.Set(key, params[key])
		}
		parsed.RawQuery = query.Encode()
	}
	return parsed.String(), nil
}

// doRequest performs the REST call. Authentication is taken from the
// configured provider credentials unless the practitioner opted out with
// use_iam_auth = false.
func doRequest(ctx context.Context, meta interface{}, req restRequest) (*restResponse, error) {
	fullURL, err := buildURL(req.URL, req.QueryParams)
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = bytes.NewBufferString(req.Body)
	}

	timeout := req.Timeout
	if timeout <= 0 {
		timeout = defaultRequestTimeout
	}
	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(reqCtx, strings.ToUpper(req.Method), fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to build %s request for %s: %w", req.Method, fullURL, err)
	}

	// Defaults that most IBM Cloud APIs expect. They are only applied when the
	// practitioner did not set the header explicitly.
	httpReq.Header.Set("Accept", "application/json")
	if req.Body != "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	if req.UseIAMAuth {
		session, err := meta.(conns.ClientSession).Authenticator()
		if err != nil {
			return nil, fmt.Errorf("failed to obtain IBM Cloud credentials, set use_iam_auth to false to send an unauthenticated request: %w", err)
		}
		if session == nil {
			return nil, fmt.Errorf("no IBM Cloud credentials are configured on the provider, set use_iam_auth to false to send an unauthenticated request")
		}
		if err := session.Authenticate(httpReq); err != nil {
			return nil, fmt.Errorf("failed to authenticate request to %s: %w", fullURL, err)
		}
	}

	log.Printf("[DEBUG] restapi: %s %s", httpReq.Method, fullURL)
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("%s request to %s failed: %w", httpReq.Method, fullURL, err)
	}
	defer httpResp.Body.Close()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body from %s: %w", fullURL, err)
	}

	headers := make(map[string]string, len(httpResp.Header))
	for key, values := range httpResp.Header {
		headers[key] = strings.Join(values, ", ")
	}

	resp := &restResponse{
		StatusCode: httpResp.StatusCode,
		Body:       string(rawBody),
		Headers:    headers,
	}
	log.Printf("[DEBUG] restapi: %s %s returned %d", httpReq.Method, fullURL, resp.StatusCode)
	return resp, nil
}

// truncateBody keeps error messages readable when an API returns a very large
// error document.
func truncateBody(body string) string {
	if len(body) <= maxLoggedBodyLength {
		return body
	}
	return body[:maxLoggedBodyLength] + "... (truncated)"
}

// responseError builds a consistent error for an unexpected status code.
func responseError(method, requestURL string, resp *restResponse) error {
	if resp.Body == "" {
		return fmt.Errorf("%s %s returned unexpected status %d with an empty body", method, requestURL, resp.StatusCode)
	}
	return fmt.Errorf("%s %s returned unexpected status %d: %s", method, requestURL, resp.StatusCode, truncateBody(resp.Body))
}

// extractID walks a dot separated path through the JSON response body and
// returns the value found there as a string. An empty body, a non JSON body or
// a missing key are all reported as errors so that the practitioner is not
// left with a resource that has no usable identifier.
func extractID(body, idAttribute string) (string, error) {
	if strings.TrimSpace(body) == "" {
		return "", fmt.Errorf("cannot read id_attribute %q: the response body is empty", idAttribute)
	}

	var decoded interface{}
	if err := json.Unmarshal([]byte(body), &decoded); err != nil {
		return "", fmt.Errorf("cannot read id_attribute %q: the response body is not valid JSON: %w", idAttribute, err)
	}

	current := decoded
	segments := strings.Split(idAttribute, ".")
	for index, segment := range segments {
		object, ok := current.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("cannot read id_attribute %q: %q is not a JSON object", idAttribute, strings.Join(segments[:index], "."))
		}
		value, found := object[segment]
		if !found {
			return "", fmt.Errorf("cannot read id_attribute %q: key %q not found, available keys are [%s]", idAttribute, segment, strings.Join(sortedKeys(object), ", "))
		}
		current = value
	}

	switch value := current.(type) {
	case string:
		if value == "" {
			return "", fmt.Errorf("cannot read id_attribute %q: the value is an empty string", idAttribute)
		}
		return value, nil
	case float64:
		// JSON numbers decode to float64. Integral identifiers are far more
		// common than fractional ones, so render them without a decimal point.
		if value == float64(int64(value)) {
			return fmt.Sprintf("%d", int64(value)), nil
		}
		return fmt.Sprintf("%v", value), nil
	case bool:
		return fmt.Sprintf("%t", value), nil
	case nil:
		return "", fmt.Errorf("cannot read id_attribute %q: the value is null", idAttribute)
	default:
		return "", fmt.Errorf("cannot read id_attribute %q: the value is a %T, expected a string, number or boolean", idAttribute, current)
	}
}

func sortedKeys(object map[string]interface{}) []string {
	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// joinURL appends a path segment to a base URL, keeping any query string on
// the base URL intact.
func joinURL(base, segment string) (string, error) {
	parsed, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("invalid url %q: %w", base, err)
	}
	// Path holds the decoded form and RawPath the encoded one. Both are set so
	// that identifiers containing reserved characters are escaped exactly once.
	basePath := strings.TrimSuffix(parsed.Path, "/")
	baseRawPath := strings.TrimSuffix(parsed.EscapedPath(), "/")
	parsed.Path = basePath + "/" + segment
	parsed.RawPath = baseRawPath + "/" + url.PathEscape(segment)
	return parsed.String(), nil
}

// stringMap converts a schema map attribute into a plain map of strings,
// skipping entries with an empty key.
func stringMap(raw interface{}) map[string]string {
	result := map[string]string{}
	values, ok := raw.(map[string]interface{})
	if !ok {
		return result
	}
	for key, value := range values {
		if key == "" {
			continue
		}
		result[key] = fmt.Sprintf("%v", value)
	}
	return result
}

// intList converts a schema list attribute into a slice of ints.
func intList(raw interface{}) []int {
	values, ok := raw.([]interface{})
	if !ok {
		return nil
	}
	result := make([]int, 0, len(values))
	for _, value := range values {
		if code, ok := value.(int); ok {
			result = append(result, code)
		}
	}
	return result
}
