// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"time"

	// "github.com/IBM/networking-go-sdk/dnssvcsv1"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.ibm.com/ibmcloud/networking-go-sdk/dnssvcsv1"
)

const (
	ibmDNSCustomResolver        = "ibm_dns_custom_resolver"
	pdnsCustomResolvers         = "custom_resolvers"
	pdnsCustomResolverLocations = "locations"
	pdnsCRId                    = "custom_resolver_id"
	pdnsCRName                  = "name"
	pdnsCRDescription           = "description"
	pdnsCRHealth                = "health"
	pdnsCREnabled               = "enabled"
	pdnsCRCreatedOn             = "created_on"
	pdnsCRModifiedOn            = "modified_on"
	pdnsCRLocationId            = "location_id"
	pdnsCRLocationSubnetCrn     = "subnet_crn"
	pdnsCRLocationEnabled       = "enabled"
	pdnsCRLocationHealthy       = "healthy"
	pdnsCRLocationDnsServerIp   = "dns_server_ip"
)

func resouceIBMPrivateDNSCustomResolver() *schema.Resource {
	return &schema.Resource{
		Create:   resouceIBMPrivateDNSCustomResolverCreate,
		Read:     resouceIBMPrivateDNSCustomResolverRead,
		Update:   resouceIBMPrivateDNSCustomResolverUpdate,
		Delete:   resouceIBMPrivateDNSCustomResolverDelete,
		Exists:   resouceIBMPrivateDNSCustomResolverExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance ID",
			},

			pdnsCRId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Identifier of the custom resolver",
			},
			pdnsCRName: {
				Type:     schema.TypeString,
				Required: true,
				// ForceNew:    true,
				Description: "Name of the custom resolver",
				// ValidateFunc: InvokeValidator(pdnsCustomResolvers,
				// 	pdnsCRName),
			},
			pdnsCRDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Descriptive text of the custom resolver.",
				// ValidateFunc: InvokeValidator(pdnsCustomResolvers,
				// 	pdnsCRDescription),
			},
			pdnsCREnabled: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the custom resolver is enabled",
				// ValidateFunc: InvokeValidator(pdnsCustomResolvers,
				// 	pdnsCREnabled),
			},
			pdnsCRHealth: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Healthy state of the custom resolver",
			},
			pdnsCustomResolverLocations: {
				Type:        schema.TypeSet,
				Description: "Locations on which the custom resolver will be running",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						pdnsCRLocationId: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Location ID",
						},
						pdnsCRLocationSubnetCrn: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Subnet CRN",
						},
						pdnsCRLocationEnabled: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the location is enabled for the custom resolver",
						},
						pdnsCRLocationHealthy: {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the DNS server in this location is healthy or not.",
						},
						pdnsCRLocationDnsServerIp: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip address of this dns server",
						},
					},
				},
			},

			pdnsCRCreatedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when a custom resolver is created",
			},

			pdnsCRModifiedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The recent time when a custom resolver is modified",
			},
		},
	}
}

func resouceIBMPrivateDNSCustomResolverCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return fmt.Errorf("Error while getting the PrivateDNSClientSession %s", err)
	}

	var crName, crDescription string

	// session options
	crn := d.Get(pdnsInstanceID).(string)
	if name, ok := d.GetOk(pdnsCRName); ok {
		crName = name.(string)
	}
	if des, ok := d.GetOk(pdnsCRDescription); ok {
		crDescription = des.(string)
	}

	crLocations := d.Get(pdnsCustomResolverLocations).(*schema.Set)
	customResolverOption := sess.NewCreateCustomResolverOptions(crn)
	customResolverOption.SetName(crName)
	customResolverOption.SetDescription(crDescription)
	customResolverOption.SetLocations(expandPdnsCRLocations(crLocations))

	result, response, err := sess.CreateCustomResolver(customResolverOption)
	if err != nil {
		return fmt.Errorf("Error creating pdns custom resolver:%s\n%s", err, response)
	}

	d.SetId(convertCisToTfTwoVar(*result.ID, crn))
	d.Set(pdnsCRId, *result.ID)

	return resouceIBMPrivateDNSCustomResolverRead(d, meta)
}

func resouceIBMPrivateDNSCustomResolverRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return fmt.Errorf("Error while getting the PrivateDNSClientSession %s", err)
	}

	customResolverID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}

	opt := sess.NewGetCustomResolverOptions(crn, customResolverID)
	result, response, err := sess.GetCustomResolver(opt)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("Error Custom Resolver not found ")
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding pdns custom resolver %q: %s", d.Id(), err)
	}
	d.Set(pdnsInstanceID, crn)
	d.Set(pdnsCRId, *result.ID)
	d.Set(pdnsCRName, *result.Name)
	d.Set(pdnsCRDescription, *result.Description)
	d.Set(pdnsCRHealth, *result.Health)
	d.Set(pdnsCREnabled, *result.Enabled)
	d.Set(pdnsCustomResolverLocations, flattenPdnsCRLocations(result.Locations))

	return nil
}

func resouceIBMPrivateDNSCustomResolverUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return fmt.Errorf("Error while getting the PrivateDNSClientSession %s", err)
	}

	customResolverID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange(pdnsCRName) ||
		d.HasChange(pdnsCRDescription) ||
		d.HasChange(pdnsCREnabled) {

		opt := sess.NewUpdateCustomResolverOptions(crn, customResolverID)
		if name, ok := d.GetOk(pdnsCRName); ok {
			crName := name.(string)
			opt.SetName(crName)
		}
		if des, ok := d.GetOk(pdnsCRDescription); ok {
			crDescription := des.(string)
			opt.SetDescription(crDescription)
		}
		if enabled, ok := d.GetOkExists(pdnsCREnabled); ok {
			crEnabled := enabled.(bool)
			opt.SetEnabled(crEnabled)
		}

		result, _, err := sess.UpdateCustomResolver(opt)
		if err != nil {
			return fmt.Errorf("Error updating pdns custom resolver: %s", err)
		}
		if *result.ID == "" {
			return fmt.Errorf("Error failed to find id in Update response; resource was empty")
		}
	}

	return resouceIBMPrivateDNSCustomResolverRead(d, meta)
}

func resouceIBMPrivateDNSCustomResolverDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}

	customResolverID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}

	opt := sess.NewDeleteCustomResolverOptions(crn, customResolverID)
	response, err := sess.DeleteCustomResolver(opt)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Error deleting pdns custom resolver:%s\n%s", err, response)
	}

	d.SetId("")
	return nil
}

func resouceIBMPrivateDNSCustomResolverExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return false, err
	}

	customResolverID, crn, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return false, err
	}
	opt := sess.NewGetCustomResolverOptions(crn, customResolverID)
	_, response, err := sess.GetCustomResolver(opt)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("Custom Resolver does not exist.")
			return false, nil
		}
		log.Printf("Custom Resolver failed: %v", response)
		return false, err
	}
	return true, nil
}

func flattenPdnsCRLocations(crLocation []dnssvcsv1.Location) interface{} {
	flattened := make([]interface{}, 0)
	for _, v := range crLocation {
		customLocations := map[string]interface{}{
			pdnsCRLocationId:          *v.ID,
			pdnsCRLocationSubnetCrn:   *v.SubnetCrn,
			pdnsCRLocationEnabled:     *v.Enabled,
			pdnsCRLocationHealthy:     *v.Healthy,
			pdnsCRLocationDnsServerIp: *v.DnsServerIp,
		}
		flattened = append(flattened, customLocations)
	}
	return flattened
}

func expandPdnsCRLocations(crLocList *schema.Set) (crLocations []dnssvcsv1.LocationInput) {
	for _, iface := range crLocList.List() {
		var locOpt dnssvcsv1.LocationInput
		loc := iface.(map[string]interface{})
		log.Println(loc[pdnsCRLocationSubnetCrn].(string))
		locOpt.SubnetCrn = core.StringPtr(loc[pdnsCRLocationSubnetCrn].(string))
		if val, ok := loc[pdnsCRLocationEnabled]; ok {
			locOpt.Enabled = core.BoolPtr(val.(bool))
		}
		crLocations = append(crLocations, locOpt)
	}
	return
}
