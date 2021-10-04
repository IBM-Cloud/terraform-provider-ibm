// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
)

func resourceIBMCbrZone() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCbrZoneCreate,
		ReadContext:   resourceIBMCbrZoneRead,
		UpdateContext: resourceIBMCbrZoneUpdate,
		DeleteContext: resourceIBMCbrZoneDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_cbr_zone", "name"),
				Description:  "The name of the zone.",
			},
			"account_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_cbr_zone", "account_id"),
				Description:  "The id of the account owning this zone.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_cbr_zone", "description"),
				Description:  "The description of the zone.",
			},
			"addresses": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of addresses in the zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of address.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IP address.",
						},
						"ref": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "A service reference value.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The id of the account owning the service.",
									},
									"service_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The service type.",
									},
									"service_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The service name.",
									},
									"service_instance": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The service instance.",
									},
								},
							},
						},
					},
				},
			},
			"excluded": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of excluded addresses in the zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of address.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IP address.",
						},
						"ref": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "A service reference value.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The id of the account owning the service.",
									},
									"service_type": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The service type.",
									},
									"service_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The service name.",
									},
									"service_instance": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The service instance.",
									},
								},
							},
						},
					},
				},
			},
			"transaction_id": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator("ibm_cbr_zone", "transaction_id"),
				Description:  "The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a transaction ID in the response.**Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each request.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone CRN.",
			},
			"address_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of addresses in the zone.",
			},
			"excluded_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of excluded addresses in the zone.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The href link to the resource.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time the resource was created.",
			},
			"created_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM ID of the user or service which created the resource.",
			},
			"last_modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last time the resource was modified.",
			},
			"last_modified_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM ID of the user or service which modified the resource.",
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMCbrZoneValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 \\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		ValidateSchema{
			Identifier:                 "account_id",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\\-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[\\x20-\\xFE]*$`,
			MinValueLength:             0,
			MaxValueLength:             300,
		},
		ValidateSchema{
			Identifier:                 "transaction_id",
			ValidateFunctionIdentifier: ValidateRegexpLen,
			Type:                       TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)

	resourceValidator := ResourceValidator{ResourceName: "ibm_cbr_zone", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMCbrZoneCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	contextBasedRestrictionsClient, err := meta.(ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createZoneOptions := &contextbasedrestrictionsv1.CreateZoneOptions{}

	if _, ok := d.GetOk("name"); ok {
		createZoneOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("account_id"); ok {
		createZoneOptions.SetAccountID(d.Get("account_id").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createZoneOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("addresses"); ok {
		var addresses []contextbasedrestrictionsv1.Address
		for _, e := range d.Get("addresses").([]interface{}) {
			value := e.(map[string]interface{})
			addressesItem := resourceIBMCbrZoneMapToAddress(value)
			addresses = append(addresses, addressesItem)
		}
		createZoneOptions.SetAddresses(addresses)
	}
	if _, ok := d.GetOk("excluded"); ok {
		var excluded []contextbasedrestrictionsv1.Address
		for _, e := range d.Get("excluded").([]interface{}) {
			value := e.(map[string]interface{})
			excludedItem := resourceIBMCbrZoneMapToAddress(value)
			excluded = append(excluded, excludedItem)
		}
		createZoneOptions.SetExcluded(excluded)
	}
	if _, ok := d.GetOk("transaction_id"); ok {
		createZoneOptions.SetTransactionID(d.Get("transaction_id").(string))
	}

	zone, response, err := contextBasedRestrictionsClient.CreateZoneWithContext(context, createZoneOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateZoneWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateZoneWithContext failed %s\n%s", err, response))
	}

	d.SetId(*zone.ID)

	return resourceIBMCbrZoneRead(context, d, meta)
}

func resourceIBMCbrZoneMapToAddress(addressMap map[string]interface{}) contextbasedrestrictionsv1.AddressIntf {
	address := contextbasedrestrictionsv1.Address{}

	if addressMap["type"] != nil {
		address.Type = core.StringPtr(addressMap["type"].(string))
	}
	if addressMap["value"] != nil {
		address.Value = core.StringPtr(addressMap["value"].(string))
	}
	if addressMap["ref"] != nil {
		// TODO: handle Ref of type ServiceRefValue -- not primitive type, not list
	}

	return &address
}

func resourceIBMCbrZoneMapToServiceRefValue(serviceRefValueMap map[string]interface{}) contextbasedrestrictionsv1.ServiceRefValue {
	serviceRefValue := contextbasedrestrictionsv1.ServiceRefValue{}

	serviceRefValue.AccountID = core.StringPtr(serviceRefValueMap["account_id"].(string))
	if serviceRefValueMap["service_type"] != nil {
		serviceRefValue.ServiceType = core.StringPtr(serviceRefValueMap["service_type"].(string))
	}
	if serviceRefValueMap["service_name"] != nil {
		serviceRefValue.ServiceName = core.StringPtr(serviceRefValueMap["service_name"].(string))
	}
	if serviceRefValueMap["service_instance"] != nil {
		serviceRefValue.ServiceInstance = core.StringPtr(serviceRefValueMap["service_instance"].(string))
	}

	return serviceRefValue
}

func resourceIBMCbrZoneMapToAddressIPAddress(addressIPAddressMap map[string]interface{}) contextbasedrestrictionsv1.AddressIPAddress {
	addressIPAddress := contextbasedrestrictionsv1.AddressIPAddress{}

	addressIPAddress.Type = core.StringPtr(addressIPAddressMap["type"].(string))
	addressIPAddress.Value = core.StringPtr(addressIPAddressMap["value"].(string))

	return addressIPAddress
}

func resourceIBMCbrZoneMapToAddressServiceRef(addressServiceRefMap map[string]interface{}) contextbasedrestrictionsv1.AddressServiceRef {
	addressServiceRef := contextbasedrestrictionsv1.AddressServiceRef{}

	addressServiceRef.Type = core.StringPtr(addressServiceRefMap["type"].(string))
	// TODO: handle Ref of type ServiceRefValue -- not primitive type, not list

	return addressServiceRef
}

func resourceIBMCbrZoneMapToAddressSubnet(addressSubnetMap map[string]interface{}) contextbasedrestrictionsv1.AddressSubnet {
	addressSubnet := contextbasedrestrictionsv1.AddressSubnet{}

	addressSubnet.Type = core.StringPtr(addressSubnetMap["type"].(string))
	addressSubnet.Value = core.StringPtr(addressSubnetMap["value"].(string))

	return addressSubnet
}

func resourceIBMCbrZoneMapToAddressIPAddressRange(addressIPAddressRangeMap map[string]interface{}) contextbasedrestrictionsv1.AddressIPAddressRange {
	addressIPAddressRange := contextbasedrestrictionsv1.AddressIPAddressRange{}

	addressIPAddressRange.Type = core.StringPtr(addressIPAddressRangeMap["type"].(string))
	addressIPAddressRange.Value = core.StringPtr(addressIPAddressRangeMap["value"].(string))

	return addressIPAddressRange
}

func resourceIBMCbrZoneMapToAddressVPC(addressVPCMap map[string]interface{}) contextbasedrestrictionsv1.AddressVPC {
	addressVPC := contextbasedrestrictionsv1.AddressVPC{}

	addressVPC.Type = core.StringPtr(addressVPCMap["type"].(string))
	addressVPC.Value = core.StringPtr(addressVPCMap["value"].(string))

	return addressVPC
}

func resourceIBMCbrZoneRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	contextBasedRestrictionsClient, err := meta.(ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{}

	getZoneOptions.SetZoneID(d.Id())

	zone, response, err := contextBasedRestrictionsClient.GetZoneWithContext(context, getZoneOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetZoneWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetZoneWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("transaction_id", getZoneOptions.TransactionID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting transaction_id: %s", err))
	}
	if err = d.Set("name", zone.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("account_id", zone.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}
	if err = d.Set("description", zone.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if zone.Addresses != nil {
		addresses := []map[string]interface{}{}
		for _, addressesItem := range zone.Addresses {
			addressesItemMap := resourceIBMCbrZoneAddressToMap(addressesItem)
			addresses = append(addresses, addressesItemMap)
		}
		if err = d.Set("addresses", addresses); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting addresses: %s", err))
		}
	}
	if zone.Excluded != nil {
		excluded := []map[string]interface{}{}
		for _, excludedItem := range zone.Excluded {
			excludedItemMap := resourceIBMCbrZoneAddressToMap(excludedItem)
			excluded = append(excluded, excludedItemMap)
		}
		if err = d.Set("excluded", excluded); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting excluded: %s", err))
		}
	}
	if err = d.Set("crn", zone.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("address_count", intValue(zone.AddressCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting address_count: %s", err))
	}
	if err = d.Set("excluded_count", intValue(zone.ExcludedCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting excluded_count: %s", err))
	}
	if err = d.Set("href", zone.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	if err = d.Set("created_at", dateTimeToString(zone.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by_id", zone.CreatedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by_id: %s", err))
	}
	if err = d.Set("last_modified_at", dateTimeToString(zone.LastModifiedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_at: %s", err))
	}
	if err = d.Set("last_modified_by_id", zone.LastModifiedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_by_id: %s", err))
	}
	if err = d.Set("version", response.Headers.Get("Etag")); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	return nil
}

func resourceIBMCbrZoneAddressToMap(address contextbasedrestrictionsv1.AddressIntf) map[string]interface{} {
	addressMap := map[string]interface{}{}

	// TODO: Add code here to convert a contextbasedrestrictionsv1.AddressIntf to map[string]interface{}

	return addressMap
}

func resourceIBMCbrZoneServiceRefValueToMap(serviceRefValue contextbasedrestrictionsv1.ServiceRefValue) map[string]interface{} {
	serviceRefValueMap := map[string]interface{}{}

	serviceRefValueMap["account_id"] = serviceRefValue.AccountID
	if serviceRefValue.ServiceType != nil {
		serviceRefValueMap["service_type"] = serviceRefValue.ServiceType
	}
	if serviceRefValue.ServiceName != nil {
		serviceRefValueMap["service_name"] = serviceRefValue.ServiceName
	}
	if serviceRefValue.ServiceInstance != nil {
		serviceRefValueMap["service_instance"] = serviceRefValue.ServiceInstance
	}

	return serviceRefValueMap
}

func resourceIBMCbrZoneAddressIPAddressToMap(addressIPAddress contextbasedrestrictionsv1.AddressIPAddress) map[string]interface{} {
	addressIPAddressMap := map[string]interface{}{}

	addressIPAddressMap["type"] = addressIPAddress.Type
	addressIPAddressMap["value"] = addressIPAddress.Value

	return addressIPAddressMap
}

func resourceIBMCbrZoneAddressServiceRefToMap(addressServiceRef contextbasedrestrictionsv1.AddressServiceRef) map[string]interface{} {
	addressServiceRefMap := map[string]interface{}{}

	addressServiceRefMap["type"] = addressServiceRef.Type
	RefMap := resourceIBMCbrZoneServiceRefValueToMap(*addressServiceRef.Ref)
	addressServiceRefMap["ref"] = []map[string]interface{}{RefMap}

	return addressServiceRefMap
}

func resourceIBMCbrZoneAddressSubnetToMap(addressSubnet contextbasedrestrictionsv1.AddressSubnet) map[string]interface{} {
	addressSubnetMap := map[string]interface{}{}

	addressSubnetMap["type"] = addressSubnet.Type
	addressSubnetMap["value"] = addressSubnet.Value

	return addressSubnetMap
}

func resourceIBMCbrZoneAddressIPAddressRangeToMap(addressIPAddressRange contextbasedrestrictionsv1.AddressIPAddressRange) map[string]interface{} {
	addressIPAddressRangeMap := map[string]interface{}{}

	addressIPAddressRangeMap["type"] = addressIPAddressRange.Type
	addressIPAddressRangeMap["value"] = addressIPAddressRange.Value

	return addressIPAddressRangeMap
}

func resourceIBMCbrZoneAddressVPCToMap(addressVPC contextbasedrestrictionsv1.AddressVPC) map[string]interface{} {
	addressVPCMap := map[string]interface{}{}

	addressVPCMap["type"] = addressVPC.Type
	addressVPCMap["value"] = addressVPC.Value

	return addressVPCMap
}

func resourceIBMCbrZoneUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	contextBasedRestrictionsClient, err := meta.(ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{}

	replaceZoneOptions.SetZoneID(d.Id())
	if _, ok := d.GetOk("name"); ok {
		replaceZoneOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("account_id"); ok {
		replaceZoneOptions.SetAccountID(d.Get("account_id").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		replaceZoneOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("addresses"); ok {
		var addresses []contextbasedrestrictionsv1.Address
		for _, e := range d.Get("addresses").([]interface{}) {
			value := e.(map[string]interface{})
			addressesItem := resourceIBMCbrZoneMapToAddress(value)
			addresses = append(addresses, addressesItem)
		}
		replaceZoneOptions.SetAddresses(addresses)
	}
	if _, ok := d.GetOk("excluded"); ok {
		var excluded []contextbasedrestrictionsv1.Address
		for _, e := range d.Get("excluded").([]interface{}) {
			value := e.(map[string]interface{})
			excludedItem := resourceIBMCbrZoneMapToAddress(value)
			excluded = append(excluded, excludedItem)
		}
		replaceZoneOptions.SetExcluded(excluded)
	}
	if _, ok := d.GetOk("transaction_id"); ok {
		replaceZoneOptions.SetTransactionID(d.Get("transaction_id").(string))
	}
	replaceZoneOptions.SetIfMatch(d.Get("version").(string))

	_, response, err := contextBasedRestrictionsClient.ReplaceZoneWithContext(context, replaceZoneOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceZoneWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ReplaceZoneWithContext failed %s\n%s", err, response))
	}

	return resourceIBMCbrZoneRead(context, d, meta)
}

func resourceIBMCbrZoneDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	contextBasedRestrictionsClient, err := meta.(ClientSession).ContextBasedRestrictionsV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteZoneOptions := &contextbasedrestrictionsv1.DeleteZoneOptions{}

	deleteZoneOptions.SetZoneID(d.Id())

	deleteZoneOptions.SetIfMatch(d.Get("version").(string))

	response, err := contextBasedRestrictionsClient.DeleteZoneWithContext(context, deleteZoneOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteZoneWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteZoneWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
