// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCRLocation           = "ibm_dns_custom_resolver_location"
	pdnsResolverID          = "resolver_id"
	pdnsCRLocationID        = "location_id"
	pdnsCRLocationSubnetCRN = "subnet_crn"
	pdnsCRLocationEnable    = "enabled"
	pdnsCRLocationServerIP  = "dns_server_ip"
)

func resourceIBMPrivateDNSCRLocation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPrivateDNSLocationCreate,
		ReadContext:   resourceIBMPrivateDNSLocationRead,
		UpdateContext: resourceIBMPrivateDNSLocationUpdate,
		DeleteContext: resourceIBMPrivateDNSLocationDelete,
		Importer:      &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance ID",
			},

			pdnsResolverID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Custom Resolver ID",
			},
			pdnsCRLocationID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRLocation ID",
			},

			pdnsCRLocationSubnetCRN: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CRLocation Subnet CRN",
			},

			pdnsCRLocationEnable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "CRLocation Enabled",
			},

			pdnsCRLocationHealthy: {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "CRLocation Healthy",
			},

			pdnsCRLocationServerIP: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRLocation Server IP",
			},
		},
	}
}
func resourceIBMPrivateDNSLocationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsResolverID).(string)

	opt := sess.NewAddCustomResolverLocationOptions(instanceID, resolverID)

	if subnetcrn, ok := d.GetOk(pdnsCRLocationSubnetCRN); ok {
		opt.SetSubnetCrn(subnetcrn.(string))
	}
	if enable, ok := d.GetOkExists(pdnsCRLocationEnable); ok {
		opt.SetEnabled(enable.(bool))
	}
	result, resp, err := sess.AddCustomResolverLocationWithContext(context, opt)

	if err != nil || result == nil {
		return diag.FromErr(fmt.Errorf("Error creating the custom resolver location %s:%s", err, resp))
	}
	d.SetId(convertCisToTfThreeVar(*result.ID, resolverID, instanceID))
	_, err = waitForPDNSCustomResolverHealthy(d, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	optCr := sess.NewUpdateCustomResolverOptions(instanceID, resolverID)
	optCr.SetEnabled(true)
	resultCr, respCr, errCr := sess.UpdateCustomResolverWithContext(context, optCr)
	if errCr != nil || resultCr == nil {
		return diag.FromErr(fmt.Errorf("Error updating the  custom resolver %s:%s", errCr, respCr))
	}

	return resourceIBMPrivateDNSLocationRead(context, d, meta)
}

func resourceIBMPrivateDNSLocationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
func resourceIBMPrivateDNSLocationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	locationID, resolverID, instanceID, err := convertTfToCisThreeVar(d.Id())
	updatelocation := sess.NewUpdateCustomResolverLocationOptions(instanceID, resolverID, locationID)

	if d.HasChange(pdnsCRLocationSubnetCRN) ||
		d.HasChange(pdnsCRLocationEnable) {
		if scrn, ok := d.GetOk(pdnsCRLocationSubnetCRN); ok {
			updatelocation.SetSubnetCrn(scrn.(string))
		}
		if e, ok := d.GetOkExists(pdnsCRLocationEnable); ok {
			updatelocation.SetEnabled(e.(bool))
		}
		result, resp, err := sess.UpdateCustomResolverLocationWithContext(context, updatelocation)
		if err != nil || result == nil {
			return diag.FromErr(fmt.Errorf("Error updating the custom resolver location %s:%s", err, resp))
		}
	}
	return resourceIBMPrivateDNSLocationRead(context, d, meta)
}
func resourceIBMPrivateDNSLocationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	locationID, resolverID, instanceID, err := convertTfToCisThreeVar(d.Id())
	updatelocation := sess.NewUpdateCustomResolverLocationOptions(instanceID, resolverID, locationID)
	updatelocation.SetEnabled(false)
	result, resp, err := sess.UpdateCustomResolverLocationWithContext(context, updatelocation)
	if err != nil || result == nil {
		return diag.FromErr(fmt.Errorf("Error updating the custom resolver location %s:%s", err, resp))
	}
	deleteCRlocation := sess.NewDeleteCustomResolverLocationOptions(instanceID, resolverID, locationID)
	resp, errDel := sess.DeleteCustomResolverLocationWithContext(context, deleteCRlocation)
	if errDel != nil {
		if resp != nil && resp.StatusCode == 404 {
			return nil
		}
		return diag.FromErr(fmt.Errorf("Error deleting the custom resolver location %s:%s", errDel, resp))
	}
	d.SetId("")
	return nil
}

func waitForPDNSCustomResolverHealthy(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	sess, err := meta.(ClientSession).PrivateDNSClientSession()
	if err != nil {
		return nil, err
	}
	_, customResolverID, crn, _ := convertTfToCisThreeVar(d.Id())
	opt := sess.NewGetCustomResolverOptions(crn, customResolverID)

	stateConf := &resource.StateChangeConf{
		Pending: []string{pdnsCustomResolverCritical},
		Target:  []string{pdnsCustomResolverDegraded, pdnsCustomResolverHealthy},
		Refresh: func() (interface{}, string, error) {
			res, detail, err := sess.GetCustomResolver(opt)
			if err != nil {
				if detail != nil && detail.StatusCode == 404 {
					return nil, "", fmt.Errorf("The custom resolver %s does not exist anymore: %v", customResolverID, err)
				}
				return nil, "", fmt.Errorf("Get the custom resolver %s failed with resp code: %s, err: %v", customResolverID, detail, err)
			}
			return res, *res.Health, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForState()
}
