package ibm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
)

const (
	cisRLThreshold   = "threshold"
	cisRLPeriod      = "period"
	cisRLDescription = "description"
	cisRLTimeout     = "timeout"
	cisRLBody        = "body"
	cisRLURL         = "url"
)

func resourceIBMCISRateLimit() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCISRateLimitCreate,
		Read:     resourceIBMCISRateLimitRead,
		Update:   resourceIBMCISRateLimitUpdate,
		Delete:   resourceIBMCISRateLimitDelete,
		Exists:   resourceIBMCISRateLimitExists,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CIS Intance CRN",
			},
			"domain_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CIS Domain ID",
			},
			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this rate limiting rule is currently disabled.",
			},
			cisRLDescription: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_cis_rate_limit", cisRLDescription),
				Description:  "A note that you can use to describe the reason for a rate limiting rule.",
			},
			"bypass": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Bypass URL",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "url",
							Description: "bypass URL name",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "bypass URL value",
						},
					},
				},
			},
			cisRLThreshold: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_cis_rate_limit", cisRLThreshold),
				Description:  "Rate Limiting Threshold",
			},
			cisRLPeriod: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_cis_rate_limit", cisRLPeriod),
				Description:  "Rate Limiting Period",
			},
			"correlate": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Ratelimiting Correlate",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"by": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "nat",
							ValidateFunc: InvokeValidator("ibm_cis_rate_limit", "by"),
							Description:  "Whether to enable NAT based rate limiting",
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Rate Limiting Action",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: InvokeValidator("ibm_cis_rate_limit", "mode"),
							Description:  "Type of action performed.Valid values are: 'simulate', 'ban', 'challenge', 'js_challenge'.",
						},
						"timeout": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: InvokeValidator("ibm_cis_rate_limit", cisRLTimeout),
							Description:  "The time to perform the mitigation action. Timeout be the same or greater than the period.",
						},
						"response": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Rate Limiting Action Response",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: InvokeValidator("ibm_cis_rate_limit", "content_type"),
										Description:  "Custom content-type and body to return. It must be one of following 'text/plain', 'text/xml', 'application/json'.",
									},
									"body": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: InvokeValidator("ibm_cis_rate_limit", cisRLBody),
										Description:  "The body to return. The content here must confirm to the 'content_type'",
									},
								},
							},
						},
					},
				},
			},
			"match": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Rate Limiting Match",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type:        schema.TypeList,
							Optional:    true,
							MinItems:    1,
							MaxItems:    1,
							Description: "Rate Limiting Match Request",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"methods": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "HTTP Methos of matching request. It can be one or many. Example methods 'POST', 'PUT'",
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: InvokeValidator("ibm_cis_rate_limit", "methods"),
										},
									},
									"schemes": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "HTTP Schemes of matching request. It can be one or many. Example schemes 'HTTP', 'HTTPS'.",
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: InvokeValidator("ibm_cis_rate_limit", "schemes"),
										},
									},
									"url": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "URL pattern of matching request",
										ValidateFunc: InvokeValidator("ibm_cis_rate_limit", cisRLURL),
									},
								},
							},
						},
						"response": {
							Type:        schema.TypeList,
							Optional:    true,
							MinItems:    1,
							MaxItems:    1,
							Description: "Rate Limiting Response",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "HTTP Status Codes of matching response. It can be one or many. Example status codes '403', '401",
										Elem:        &schema.Schema{Type: schema.TypeInt},
									},
									"origin_traffic": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Origin Traffic of matching response.",
									},
									"headers": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The name of the response header to match.",
												},
												"op": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The operator when matching. Valid values are 'eq' and 'ne'.",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value of the header, which is exactly matched.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Rate Limit rule Id",
			},
		},
	}
}
func resourceIBMCISRateLimitValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	byValues := "nat"
	modeValues := "simulate, ban, challenge, js_challenge"
	ctypeValues := "text/plain, text/xml, application/json"
	methodValues := "GET, POST, PUT, DELETE, PATCH, HEAD, _ALL_"
	schemeValues := "HTTP, HTTPS, _ALL_"

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisRLDescription,
			ValidateFunctionIdentifier: StringLenBetween,
			Type:                       TypeString,
			Optional:                   true,
			MinValueLength:             0,
			MaxValueLength:             1024})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisRLThreshold,
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			Required:                   true,
			MinValue:                   "1",
			MaxValue:                   "1000000"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisRLPeriod,
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			Required:                   true,
			MinValue:                   "1",
			MaxValue:                   "86400"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "by",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              byValues})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "mode",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              modeValues})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "content_type",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              ctypeValues})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "methods",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              methodValues})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "schemes",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              schemeValues})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisRLTimeout,
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			Optional:                   true,
			MinValue:                   "1",
			MaxValue:                   "86400"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisRLBody,
			ValidateFunctionIdentifier: StringLenBetween,
			Type:                       TypeString,
			Optional:                   true,
			MinValueLength:             0,
			MaxValueLength:             10240})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisRLURL,
			ValidateFunctionIdentifier: StringLenBetween,
			Type:                       TypeString,
			Optional:                   true,
			MinValueLength:             0,
			MaxValueLength:             1024})

	ibmCISRateLimitResourceValidator := ResourceValidator{ResourceName: "ibm_cis_rate_limit", Schema: validateSchema}
	return &ibmCISRateLimitResourceValidator
}
func resourceIBMCISRateLimitCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	cisID := d.Get("cis_id").(string)
	zoneID, _, err := convertTftoCisTwoVar(d.Get("domain_id").(string))
	if err != nil {
		return err
	}
	//payload to create a rate limit rule
	newRecord := v1.RateLimitRecord{
		Threshold: d.Get(cisRLThreshold).(int),
		Period:    d.Get(cisRLPeriod).(int),
	}

	if description, ok := d.GetOk(cisRLDescription); ok {
		newRecord.Description = description.(string)
	}

	if disabled, ok := d.GetOk("disabled"); ok {
		newRecord.Disabled = disabled.(bool)
	}

	action, err := expandRateLimitAction(d)
	if err != nil {
		return fmt.Errorf("Error in getting action from expandRateLimitAction %s", err)
	}
	newRecord.Action = action

	match, err := expandRateLimitMatch(d)
	if err != nil {
		return fmt.Errorf("Error in getting match from expandRateLimitMatch %s", err)
	}
	newRecord.Match = match

	correlate, err := expandRateLimitCorrelate(d)
	if err != nil {
		return fmt.Errorf("Error in getting correlate from expandRateLimitCorrelate %s", err)
	}
	newRecord.Correlate = &correlate

	byPass, err := expandRateLimitBypass(d)
	if err != nil {
		return fmt.Errorf("Error in getting bypass from expandRateLimitBypass %s", err)
	}
	newRecord.Bypass = byPass

	//creating rate limit rule
	recordPtr, err := cisClient.RateLimit().CreateRateLimit(cisID, zoneID, newRecord)
	if err != nil {
		return fmt.Errorf("Failed to create RateLimit: %v", err)
	}
	record := *recordPtr
	if record.ID == "" {
		return fmt.Errorf("Failed to find record in Create response; Record was empty")
	}
	d.SetId(convertCisToTfThreeVar(record.ID, zoneID, cisID))

	return resourceIBMCISRateLimitRead(d, meta)
}

func resourceIBMCISRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	recordID, zoneID, cisID, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	rateLimit, err := cisClient.RateLimit().GetRateLimit(cisID, zoneID, recordID)
	if err != nil {
		if strings.Contains(err.Error(), "Request failed with status code: 404") {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Failed to read RateLimit: %v", err)
	}

	rule := *rateLimit
	d.Set("cis_id", cisID)
	d.Set("domain_id", convertCisToTfTwoVar(zoneID, cisID))
	d.Set("rule_id", recordID)
	d.Set("disabled", rule.Disabled)
	d.Set(cisRLDescription, rule.Description)
	d.Set(cisRLThreshold, rule.Threshold)
	d.Set(cisRLPeriod, rule.Period)
	d.Set("action", flattenRateLimitAction(rule.Action))
	d.Set("match", flattenRateLimitMatch(rule.Match))
	d.Set("correlate", flattenRateLimitCorrelate(*rule.Correlate))
	d.Set("bypass", flattenRateLimitByPass(rule.Bypass))

	return nil
}

func resourceIBMCISRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	recordID, zoneID, cisID, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	updateRecord := v1.RateLimitRecord{}
	if d.HasChange("disabled") || d.HasChange(cisRLThreshold) || d.HasChange(cisRLPeriod) || d.HasChange(cisRLDescription) || d.HasChange("action") || d.HasChange("match") || d.HasChange("correlate") || d.HasChange("bypass") {

		updateRecord.Threshold = d.Get(cisRLThreshold).(int)
		updateRecord.Period = d.Get(cisRLPeriod).(int)

		if description, ok := d.GetOk(cisRLDescription); ok {
			updateRecord.Description = description.(string)
		}

		if disabled, ok := d.GetOk("disabled"); ok {
			updateRecord.Disabled = disabled.(bool)
		}

		action, err := expandRateLimitAction(d)
		if err != nil {
			return fmt.Errorf("Error in getting action from expandRateLimitAction %s", err)
		}
		updateRecord.Action = action

		match, err := expandRateLimitMatch(d)
		if err != nil {
			return fmt.Errorf("Error in getting match from expandRateLimitMatch %s", err)
		}
		updateRecord.Match = match

		correlate, err := expandRateLimitCorrelate(d)
		if err != nil {
			return fmt.Errorf("Error in getting correlate from expandRateLimitCorrelate %s", err)
		}
		updateRecord.Correlate = &correlate

		byPass, err := expandRateLimitBypass(d)
		if err != nil {
			return fmt.Errorf("Error in getting bypass from expandRateLimitBypass %s", err)
		}
		updateRecord.Bypass = byPass

	}
	_, err = cisClient.RateLimit().UpdateRateLimit(cisID, zoneID, recordID, updateRecord)
	if err != nil {
		return fmt.Errorf("Failed to update RateLimit: %v", err)
	}
	return resourceIBMCISRateLimitRead(d, meta)
}

func resourceIBMCISRateLimitDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	recordID, zoneID, cisID, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	err = cisClient.RateLimit().DeleteRateLimit(cisID, zoneID, recordID)
	if err != nil && !strings.Contains(err.Error(), "Request failed with status code: 404") {
		return fmt.Errorf("Failed to delete RateLimit: %v", err)
	}
	return nil
}

func resourceIBMCISRateLimitExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return false, err
	}
	recordID, zoneID, cisID, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return false, err
	}
	_, err = cisClient.RateLimit().GetRateLimit(cisID, zoneID, recordID)
	if err != nil {
		if strings.Contains(err.Error(), "Request failed with status code: 404") {
			return false, nil
		}
		return false, fmt.Errorf("Failed to getting existing RateLimit: %v", err)
	}
	return true, nil
}

func expandRateLimitAction(d *schema.ResourceData) (action v1.RateLimitAction, err error) {
	actionRecord := d.Get("action").([]interface{})[0].(map[string]interface{})
	mode := actionRecord["mode"].(string)
	timeout := actionRecord["timeout"].(int)
	if timeout == 0 {
		if mode == "simulate" || mode == "ban" {
			return action, fmt.Errorf("For the mode 'simulate' and 'ban' timeout must be set.. valid range for timeout is 10 - 86400. %s", err)
		}
	} else {
		if mode == "challenge" || mode == "js_challenge" {
			return action, fmt.Errorf("Timeout field is only valid for 'simulate' and 'ban' modes. %s", err)
		}
	}
	action.Mode = mode
	action.Timeout = timeout

	if _, ok := actionRecord["response"]; ok && len(actionRecord["response"].([]interface{})) > 0 {
		actionResponse := actionRecord["response"].([]interface{})[0].(map[string]interface{})
		action.Response = &v1.ActionResponse{
			ContentType: actionResponse["content_type"].(string),
			Body:        actionResponse["body"].(string),
		}
	}

	return action, nil
}

func expandRateLimitMatch(d *schema.ResourceData) (match v1.RateLimitMatch, err error) {
	m, ok := d.GetOk("match")
	if !ok {
		return
	}
	matchRecord := m.([]interface{})[0].(map[string]interface{})

	if matchReqRecord, ok := matchRecord["request"]; ok && len(matchReqRecord.([]interface{})) > 0 {
		matchRequestRecord := matchReqRecord.([]interface{})[0].(map[string]interface{})

		matchRequest := v1.MatchRequest{
			URL: matchRequestRecord["url"].(string),
		}
		if methodsRecord, ok := matchRequestRecord["methods"]; ok {
			methods := make([]string, methodsRecord.(*schema.Set).Len())
			for i, m := range methodsRecord.(*schema.Set).List() {
				methods[i] = m.(string)
			}
			matchRequest.Methods = methods
		}
		if schemesRecord, ok := matchRequestRecord["schemes"]; ok {
			schemes := make([]string, schemesRecord.(*schema.Set).Len())
			for i, s := range schemesRecord.(*schema.Set).List() {
				schemes[i] = s.(string)
			}
			matchRequest.Schemes = schemes
		}

		match.Request = matchRequest
	}
	if matchResRecord, ok := matchRecord["response"]; ok && len(matchResRecord.([]interface{})) > 0 {
		matchResponseRecord := matchResRecord.([]interface{})[0].(map[string]interface{})
		matchResponse := v1.MatchResponse{}
		if statusRecord, ok := matchResponseRecord["status"]; ok {
			statuses := make([]int, statusRecord.(*schema.Set).Len())
			for i, s := range statusRecord.(*schema.Set).List() {
				statuses[i] = s.(int)
			}
			matchResponse.Statuses = statuses
		}
		if originRecord, ok := matchResponseRecord["origin_traffic"]; ok {
			originTraffic := originRecord.(bool)
			matchResponse.OriginTraffic = &originTraffic
		}
		if headersRecord, ok := matchResponseRecord["headers"]; ok && len(headersRecord.([]interface{})) > 0 {
			matchResponseHeaders := headersRecord.([]interface{})

			responseHeaders := make([]v1.MatchResponseHeader, 0)

			for _, h := range matchResponseHeaders {
				header, _ := h.(map[string]interface{})
				headerRecord := v1.MatchResponseHeader{}
				headerRecord.Name = header["name"].(string)
				headerRecord.Op = header["op"].(string)
				headerRecord.Value = header["value"].(string)
				responseHeaders = append(responseHeaders, headerRecord)
			}
			matchResponse.Headers = responseHeaders

		}
		match.Response = matchResponse
	}

	return match, nil
}

func expandRateLimitCorrelate(d *schema.ResourceData) (correlate v1.RateLimitCorrelate, err error) {
	c, ok := d.GetOk("correlate")
	if !ok {
		return
	}
	correlateRecord := c.([]interface{})[0].(map[string]interface{})
	correlate = v1.RateLimitCorrelate{
		By: correlateRecord["by"].(string),
	}

	return correlate, nil
}

func expandRateLimitBypass(d *schema.ResourceData) (byPass []v1.RateLimitByPass, err error) {
	b, ok := d.GetOk("bypass")
	if !ok {
		return
	}
	byPassKV := b.([]interface{})

	byPassRecord := make([]v1.RateLimitByPass, 0)

	for _, kv := range byPassKV {
		keyValue, _ := kv.(map[string]interface{})

		byPassKeyValue := v1.RateLimitByPass{}
		byPassKeyValue.Name = keyValue["name"].(string)
		byPassKeyValue.Value = keyValue["value"].(string)
		byPassRecord = append(byPassRecord, byPassKeyValue)
	}
	byPass = byPassRecord

	return byPass, nil
}

func flattenRateLimitAction(action v1.RateLimitAction) []map[string]interface{} {
	actionRecord := map[string]interface{}{
		"mode":    action.Mode,
		"timeout": action.Timeout,
	}

	if action.Response != nil {
		actionResponseRecord := *action.Response
		actionResponse := map[string]interface{}{
			"content_type": actionResponseRecord.ContentType,
			"body":         actionResponseRecord.Body,
		}
		actionRecord["response"] = []map[string]interface{}{actionResponse}
	}
	return []map[string]interface{}{actionRecord}
}

func flattenRateLimitMatch(match v1.RateLimitMatch) []map[string]interface{} {

	matchRecord := map[string]interface{}{}
	matchRecord["request"] = flattenRateLimitMatchRequest(match.Request)
	matchRecord["response"] = flattenRateLimitMatchResponse(match.Response)

	return []map[string]interface{}{matchRecord}
}

func flattenRateLimitMatchRequest(request v1.MatchRequest) []map[string]interface{} {

	requestRecord := map[string]interface{}{}
	methods := make([]string, 0)
	for _, m := range request.Methods {
		methods = append(methods, m)
	}
	requestRecord["methods"] = methods
	schemes := make([]string, 0)
	for _, s := range request.Schemes {
		schemes = append(schemes, s)
	}
	requestRecord["schemes"] = schemes

	requestRecord["url"] = request.URL
	return []map[string]interface{}{requestRecord}
}

func flattenRateLimitMatchResponse(response v1.MatchResponse) []interface{} {
	responseRecord := map[string]interface{}{}
	flag := false
	if response.OriginTraffic != nil {
		responseRecord["origin_traffic"] = *response.OriginTraffic
		flag = true
	}

	if len(response.Statuses) > 0 {
		statuses := make([]int, 0)
		for _, s := range response.Statuses {
			statuses = append(statuses, s)
		}
		responseRecord["status"] = statuses
		flag = true
	}

	if len(response.Headers) > 0 {
		headers := make([]map[string]interface{}, 0)
		for _, h := range response.Headers {
			header := map[string]interface{}{}
			header["name"] = h.Name
			header["op"] = h.Op
			header["value"] = h.Value
			headers = append(headers, header)

		}
		responseRecord["headers"] = headers
		flag = true
	}
	if flag == true {
		return []interface{}{responseRecord}
	}
	return []interface{}{}
}
func flattenRateLimitCorrelate(correlate v1.RateLimitCorrelate) []map[string]interface{} {
	correlateRecord := map[string]interface{}{}
	if correlate.By != "" {
		correlateRecord["by"] = correlate.By
	}
	return []map[string]interface{}{correlateRecord}
}

func flattenRateLimitByPass(byPass []v1.RateLimitByPass) []map[string]interface{} {
	byPassRecord := make([]map[string]interface{}, 0, len(byPass))
	if len(byPass) > 0 {
		for _, b := range byPass {
			byPassKV := map[string]interface{}{
				"name":  b.Name,
				"value": b.Value,
			}
			byPassRecord = append(byPassRecord, byPassKV)
		}
	}
	return byPassRecord
}
