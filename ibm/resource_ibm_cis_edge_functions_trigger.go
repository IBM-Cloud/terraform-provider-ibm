package ibm

import (
	"fmt"
	"log"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	cisEdgeFunctionsTriggerRouteID              = "route_id"
	cisEdgeFunctionsTriggerPattern              = "pattern"
	cisEdgeFunctionsTriggerScript               = "script"
	cisEdgeFunctionsTriggerRequestLimitFailOpen = "request_limit_fail_open"
)

func resourceIBMCISEdgeFunctionsTrigger() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCISEdgeFunctionsTriggerCreate,
		Read:     resourceIBMCISEdgeFunctionsTriggerRead,
		Update:   resourceIBMCISEdgeFunctionsTriggerUpdate,
		Delete:   resourceIBMCISEdgeFunctionsTriggerDelete,
		Exists:   resourceIBMCISEdgeFunctionsTriggerExists,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CIS Intance CRN",
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "CIS Domain ID",
				DiffSuppressFunc: suppressDataDiff,
			},
			cisEdgeFunctionsTriggerRouteID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CIS Edge Functions trigger route ID",
			},
			cisEdgeFunctionsTriggerPattern: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Edge function trigger pattern",
			},
			cisEdgeFunctionsTriggerScript: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Edge function trigger script name",
			},
			cisEdgeFunctionsTriggerRequestLimitFailOpen: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Edge function trigger request limit fail open",
			},
		},
	}
}

func resourceIBMCISEdgeFunctionsTriggerCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	zoneID, _, err := convertTftoCisTwoVar(d.Get(cisDomainID).(string))
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	script := d.Get(cisEdgeFunctionsTriggerScript).(string)
	pattern := d.Get(cisEdgeFunctionsTriggerPattern).(string)

	opt := cisClient.NewCreateEdgeFunctionsTriggerOptions()
	opt.SetPattern(pattern)
	opt.SetScript(script)

	result, _, err := cisClient.CreateEdgeFunctionsTrigger(opt)
	if err != nil {
		return fmt.Errorf("Error creating edge function trigger route : %v", err)
	}
	d.SetId(convertCisToTfThreeVar(*result.Result.ID, zoneID, crn))
	return resourceIBMCISEdgeFunctionsTriggerRead(d, meta)
}

func resourceIBMCISEdgeFunctionsTriggerUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return err
	}

	routeID, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	if d.HasChange(cisEdgeFunctionsTriggerScript) ||
		d.HasChange(cisEdgeFunctionsTriggerPattern) {
		script := d.Get(cisEdgeFunctionsTriggerScript).(string)
		pattern := d.Get(cisEdgeFunctionsTriggerPattern).(string)

		opt := cisClient.NewUpdateEdgeFunctionsTriggerOptions(routeID)
		opt.SetPattern(pattern)
		opt.SetScript(script)

		_, _, err := cisClient.UpdateEdgeFunctionsTrigger(opt)
		if err != nil {
			return fmt.Errorf("Error updating edge function trigger route : %v", err)
		}
	}
	return resourceIBMCISEdgeFunctionsTriggerRead(d, meta)
}

func resourceIBMCISEdgeFunctionsTriggerRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return err
	}

	routeID, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewGetEdgeFunctionsTriggerOptions(routeID)
	result, resp, err := cisClient.GetEdgeFunctionsTrigger(opt)
	if err != nil {
		return fmt.Errorf("Error: %v", resp)
	}
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisEdgeFunctionsTriggerRouteID, routeID)
	d.Set(cisEdgeFunctionsTriggerScript, result.Result.Script)
	d.Set(cisEdgeFunctionsTriggerPattern, result.Result.Pattern)
	d.Set(cisEdgeFunctionsTriggerRequestLimitFailOpen, result.Result.RequestLimitFailOpen)
	return nil
}

func resourceIBMCISEdgeFunctionsTriggerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cisClient, err := meta.(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return false, err
	}

	routeID, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewGetEdgeFunctionsTriggerOptions(routeID)
	_, response, err := cisClient.GetEdgeFunctionsTrigger(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("Edge functions trigger route is not found")
			return false, nil
		}
		return false, fmt.Errorf("Error: %v", response)
	}
	return true, nil
}

func resourceIBMCISEdgeFunctionsTriggerDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisEdgeFunctionClientSession()
	if err != nil {
		return fmt.Errorf("Error in creating CIS object")
	}

	routeID, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewDeleteEdgeFunctionsTriggerOptions(routeID)
	_, response, err := cisClient.DeleteEdgeFunctionsTrigger(opt)
	if err != nil {
		return fmt.Errorf("Error in edge function trigger route deletion: %v", response)
	}
	return nil
}
