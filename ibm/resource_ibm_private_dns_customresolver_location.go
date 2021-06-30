// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCRLocation           = "ibm_dns_cr_locations"
	pdnsResolverID          = "resolver_id"
	pdnsCRLocationID        = "location_id"
	pdnsCRLocationSubnetCRN = "subnet_crn"
	pdnsCRLocationEnable    = "enabled"
	pdnsCRLocationServerIP  = "dns_server_ip"
)

func resourceIBMPrivateDNSCRLocation() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDNSLocationCreate,
		Read:     resourceIBMPrivateDNSLocationRead,
		Update:   resourceIBMPrivateDNSLocationUpdate,
		Delete:   resourceIBMPrivateDNSLocationDelete,
		Importer: &schema.ResourceImporter{},
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
				Description: "Zone ID",
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
func resourceIBMPrivateDNSLocationCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return fmt.Errorf("Error connecting to session :%s", err)
	}
	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsResolverID).(string)

	opt := sess.NewAddCustomResolverLocationOptions(instanceID, resolverID)

	if subnetcrn, ok := d.GetOk(pdnsCRLocationSubnetCRN); ok {
		opt.SetSubnetCrn(subnetcrn.(string))
		log.Printf("[DEBUG] *********** SetSubnetCrn ********* ")
	}
	if enable, ok := d.GetOkExists(pdnsCRLocationEnable); ok {
		opt.SetEnabled(enable.(bool))
		log.Printf("[DEBUG] *********** SetEnabled ********* ")
	}
	log.Printf("[DEBUG] *********** Before AddCustomResolverLocation ********* %v", opt)
	result, resp, err := sess.AddCustomResolverLocation(opt)

	log.Printf("[DEBUG] *********** AddCustomResolverLocation ********* %v", result)
	if err != nil || result == nil {
		return fmt.Errorf("Error while adding customer resolver location :%s %s", err, resp)
	}
	d.SetId(convertCisToTfThreeVar(*result.ID, resolverID, instanceID))
	return resourceIBMPrivateDNSLocationRead(d, meta)
}

func resourceIBMPrivateDNSLocationRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceIBMPrivateDNSLocationUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
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
		_, resp, err := sess.UpdateCustomResolverLocation(updatelocation)
		if err != nil {
			return fmt.Errorf("Error updating custom resolver location :%s\n%s", err, resp)
		}
	}
	return resourceIBMPrivateDNSLocationRead(d, meta)
}
func resourceIBMPrivateDNSLocationDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}
	locationID, resolverID, instanceID, err := convertTfToCisThreeVar(d.Id())
	deleteCRlocation := sess.NewDeleteCustomResolverLocationOptions(instanceID, resolverID, locationID)
	resp, err := sess.DeleteCustomResolverLocation(deleteCRlocation)
	if err != nil {
		return fmt.Errorf("Error deleting custom resolver locations :%s\n%s", err, resp)
	}
	d.SetId("")
	return nil
}
