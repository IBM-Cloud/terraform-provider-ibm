/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isDedicatedHostAvailable  = "available"
	isDedicatedHostDeleting   = "deleting"
	isDedicatedHostDeleteDone = "done"
	isDedicatedHostFailed     = "failed"
)

func resourceIbmIsDedicatedHost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmIsDedicatedHostCreate,
		ReadContext:   resourceIbmIsDedicatedHostRead,
		UpdateContext: resourceIbmIsDedicatedHostUpdate,
		DeleteContext: resourceIbmIsDedicatedHostDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_placement_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If set to true, instances can be placed on this dedicated host.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The profile to use for this dedicated host. Globally unique name for this dedicated host profile.",
			},
			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier for the resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.",
			},
			"host_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier of the dedicated host group for this dedicated host.",
			},
			"available_memory": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of memory in gibibytes that is currently available for instances.",
			},
			"available_vcpu": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The available VCPU for the dedicated host.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"architecture": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The VCPU architecture.",
						},
						"count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of VCPUs assigned.",
						},
					},
				},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the dedicated host was created.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this dedicated host.",
			},
			"group": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The dedicated host group this dedicated host is in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN for this dedicated host group.",
						},
						"deleted": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The URL for this dedicated host group.",
						},
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The unique identifier for this dedicated host group.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of resource referenced.",
						},
					},
				},
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dedicated host.",
			},
			"instances": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of instances that are allocated to this dedicated host.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN for this virtual server instance.",
						},
						"deleted": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "If present, this property indicates the referenced resource has been deleted and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Link to documentation about deleted resources.",
									},
								},
							},
						},
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The URL for this virtual server instance.",
						},
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The unique identifier for this virtual server instance.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The user-defined name for this virtual server instance (and default system hostname).",
						},
					},
				},
			},
			"lifecycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the dedicated host resource.",
			},
			"memory": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total amount of memory in gibibytes for this host.",
			},
			"provisionable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this dedicated host is available for instance creation.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of resource referenced.",
			},
			"socket_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of sockets for this host.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The administrative state of the dedicated host.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.",
			},
			"supported_instance_profiles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of instance profiles that can be used by instances placed on this dedicated host.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The URL for this virtual server instance profile.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The globally unique name for this virtual server instance profile.",
						},
					},
				},
			},
			"vcpu": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The total VCPU of the dedicated host.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"architecture": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The VCPU architecture.",
						},
						"count": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of VCPUs assigned.",
						},
					},
				},
			},
			"zone": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The zone this dedicated host resides in.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The URL for this zone.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The globally unique name for this zone.",
						},
					},
				},
			},
		},
	}
}

func resourceIbmIsDedicatedHostCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}
	log.Println("in the create method ******")
	createDedicatedHostOptions := &vpcv1.CreateDedicatedHostOptions{}
	dedicatedHostPrototype := vpcv1.DedicatedHostPrototype{}

	if dhname, ok := d.GetOk("name"); ok {

		namestr := dhname.(string)
		if namestr != "" {
			dedicatedHostPrototype.Name = &namestr
		}
	}
	if insplacementenabled, ok := d.GetOk("instance_placement_enabled"); ok {
		insplacementenabledbool := insplacementenabled.(bool)
		dedicatedHostPrototype.InstancePlacementEnabled = &insplacementenabledbool
	}

	if dhprofile, ok := d.GetOk("profile"); ok {
		dhprofilename := dhprofile.(string)
		dedicatedHostProfileIdentity := vpcv1.DedicatedHostProfileIdentity{
			Name: &dhprofilename,
		}
		dedicatedHostPrototype.Profile = &dedicatedHostProfileIdentity
	}

	if dhgroup, ok := d.GetOk("host_group"); ok {
		dhgroupid := dhgroup.(string)
		dedicatedHostGroupIdentity := vpcv1.DedicatedHostGroupIdentity{
			ID: &dhgroupid,
		}
		dedicatedHostPrototype.Group = &dedicatedHostGroupIdentity
	}

	if resgroup, ok := d.GetOk("resource_group"); ok {
		resgroupid := resgroup.(string)
		resourceGroupIdentity := vpcv1.ResourceGroupIdentity{
			ID: &resgroupid,
		}
		dedicatedHostPrototype.ResourceGroup = &resourceGroupIdentity
	}

	createDedicatedHostOptions.SetDedicatedHostPrototype(&dedicatedHostPrototype)

	//dedicatedHostPrototype := resourceIbmIsDedicatedHostMapToDedicatedHostPrototype(d.Get("dedicated_host_prototype").([]interface{})[0].(map[string]interface{}))
	//createDedicatedHostOptions.SetDedicatedHostPrototype(&dedicatedHostPrototype)
	dedicatedHost, response, err := vpcClient.CreateDedicatedHostWithContext(context, createDedicatedHostOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateDedicatedHostWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*dedicatedHost.ID)

	return resourceIbmIsDedicatedHostRead(context, d, meta)
}

func resourceIbmIsDedicatedHostMapToDedicatedHostPrototype(dedicatedHostPrototypeMap map[string]interface{}) vpcv1.DedicatedHostPrototype {
	dedicatedHostPrototype := vpcv1.DedicatedHostPrototype{}
	log.Println("resourceIbmIsDedicatedHostMapToDedicatedHostPrototype  ******* ", dedicatedHostPrototypeMap)
	if dedicatedHostPrototypeMap["instance_placement_enabled"] != nil {
		dedicatedHostPrototype.InstancePlacementEnabled = core.BoolPtr(dedicatedHostPrototypeMap["instance_placement_enabled"].(bool))
	}

	dhName, _ := dedicatedHostPrototypeMap["name"]
	dhNamestr := dhName.(string)
	if dhNamestr != "" {
		dedicatedHostPrototype.Name = core.StringPtr(dedicatedHostPrototypeMap["name"].(string))
	}

	if len(dedicatedHostPrototypeMap["profile"].([]interface{})) != 0 {
		log.Println("dedicatedHostPrototypeMap profile ******* ", dedicatedHostPrototypeMap["profile"])
		dedicatedHostPrototype.Profile = resourceIbmIsDedicatedHostMapToDedicatedHostProfileIdentity(dedicatedHostPrototypeMap["profile"].([]interface{})[0].(map[string]interface{}))
	}
	// TODO: handle Profile of type DedicatedHostProfileIdentity -- not primitive type, not list

	if len(dedicatedHostPrototypeMap["resource_group"].([]interface{})) != 0 {
		// TODO: handle ResourceGroup of type ResourceGroupIdentity -- not primitive type, not list
		log.Println("dedicatedHostPrototypeMap resource_group ******* ", dedicatedHostPrototypeMap["resource_group"])
		dedicatedHostPrototype.ResourceGroup = resourceIbmIsDedicatedHostMapToResourceGroupIdentity(dedicatedHostPrototypeMap["resource_group"].([]interface{})[0].(map[string]interface{}))
	}
	if len(dedicatedHostPrototypeMap["group"].([]interface{})) != 0 {
		// TODO: handle Group of type DedicatedHostGroupIdentity -- not primitive type, not list
		dedicatedHostPrototype.Group = resourceIbmIsDedicatedHostMapToDedicatedHostGroupIdentity(dedicatedHostPrototypeMap["group"].([]interface{})[0].(map[string]interface{}))
	}
	if len(dedicatedHostPrototypeMap["zone"].([]interface{})) != 0 {
		// TODO: handle Zone of type ZoneIdentity -- not primitive type, not list
		log.Printf("zone *************")
		dedicatedHostPrototype.Zone = resourceIbmIsDedicatedHostMapToZoneIdentity(dedicatedHostPrototypeMap["zone"].([]interface{})[0].(map[string]interface{}))
	}

	return dedicatedHostPrototype
}

func resourceIbmIsDedicatedHostMapToDedicatedHostProfileIdentity(dedicatedHostProfileIdentityMap map[string]interface{}) *vpcv1.DedicatedHostProfileIdentity {
	dedicatedHostProfileIdentity := vpcv1.DedicatedHostProfileIdentity{}

	name, _ := dedicatedHostProfileIdentityMap["name"]
	namestr := name.(string)
	if namestr != "" {
		dedicatedHostProfileIdentity.Name = core.StringPtr(dedicatedHostProfileIdentityMap["name"].(string))
	}
	href, _ := dedicatedHostProfileIdentityMap["href"]
	hrefstr := href.(string)
	if hrefstr != "" {
		dedicatedHostProfileIdentity.Href = core.StringPtr(dedicatedHostProfileIdentityMap["href"].(string))
	}

	return &dedicatedHostProfileIdentity
}

func resourceIbmIsDedicatedHostMapToDedicatedHostProfileIdentityByName(dedicatedHostProfileIdentityByNameMap map[string]interface{}) vpcv1.DedicatedHostProfileIdentityByName {
	dedicatedHostProfileIdentityByName := vpcv1.DedicatedHostProfileIdentityByName{}

	dedicatedHostProfileIdentityByName.Name = core.StringPtr(dedicatedHostProfileIdentityByNameMap["name"].(string))

	return dedicatedHostProfileIdentityByName
}

func resourceIbmIsDedicatedHostMapToDedicatedHostProfileIdentityByHref(dedicatedHostProfileIdentityByHrefMap map[string]interface{}) vpcv1.DedicatedHostProfileIdentityByHref {
	dedicatedHostProfileIdentityByHref := vpcv1.DedicatedHostProfileIdentityByHref{}

	dedicatedHostProfileIdentityByHref.Href = core.StringPtr(dedicatedHostProfileIdentityByHrefMap["href"].(string))

	return dedicatedHostProfileIdentityByHref
}

func resourceIbmIsDedicatedHostMapToResourceGroupIdentity(resourceGroupIdentityMap map[string]interface{}) *vpcv1.ResourceGroupIdentity {
	resourceGroupIdentity := vpcv1.ResourceGroupIdentity{}

	resourcegroup, _ := resourceGroupIdentityMap["id"]
	resourcegroupstr := resourcegroup.(string)
	if resourcegroupstr != "" {
		resourceGroupIdentity.ID = core.StringPtr(resourceGroupIdentityMap["id"].(string))
	}

	return &resourceGroupIdentity
}

func resourceIbmIsDedicatedHostMapToResourceGroupIdentityByID(resourceGroupIdentityByIDMap map[string]interface{}) vpcv1.ResourceGroupIdentityByID {
	resourceGroupIdentityByID := vpcv1.ResourceGroupIdentityByID{}

	resourceGroupIdentityByID.ID = core.StringPtr(resourceGroupIdentityByIDMap["id"].(string))

	return resourceGroupIdentityByID
}

func resourceIbmIsDedicatedHostMapToDedicatedHostGroupIdentity(dedicatedHostGroupIdentityMap map[string]interface{}) *vpcv1.DedicatedHostGroupIdentity {
	dedicatedHostGroupIdentity := vpcv1.DedicatedHostGroupIdentity{}

	hostGroupID, _ := dedicatedHostGroupIdentityMap["id"]
	hostGroupIDstr := hostGroupID.(string)
	if hostGroupIDstr != "" {
		dedicatedHostGroupIdentity.ID = core.StringPtr(dedicatedHostGroupIdentityMap["id"].(string))
	}
	hostGroupCRN, _ := dedicatedHostGroupIdentityMap["crn"]
	hostGroupCRNstr := hostGroupCRN.(string)
	if hostGroupCRNstr != "" {
		dedicatedHostGroupIdentity.CRN = core.StringPtr(dedicatedHostGroupIdentityMap["crn"].(string))
	}
	hostGrouphref, _ := dedicatedHostGroupIdentityMap["href"]
	hostGrouphrefstr := hostGrouphref.(string)
	if hostGrouphrefstr != "" {
		dedicatedHostGroupIdentity.Href = core.StringPtr(dedicatedHostGroupIdentityMap["href"].(string))
	}

	return &dedicatedHostGroupIdentity
}

func resourceIbmIsDedicatedHostMapToDedicatedHostGroupIdentityByID(dedicatedHostGroupIdentityByIDMap map[string]interface{}) vpcv1.DedicatedHostGroupIdentityByID {
	dedicatedHostGroupIdentityByID := vpcv1.DedicatedHostGroupIdentityByID{}

	dedicatedHostGroupIdentityByID.ID = core.StringPtr(dedicatedHostGroupIdentityByIDMap["id"].(string))

	return dedicatedHostGroupIdentityByID
}

func resourceIbmIsDedicatedHostMapToDedicatedHostGroupIdentityByCRN(dedicatedHostGroupIdentityByCRNMap map[string]interface{}) vpcv1.DedicatedHostGroupIdentityByCRN {
	dedicatedHostGroupIdentityByCRN := vpcv1.DedicatedHostGroupIdentityByCRN{}

	dedicatedHostGroupIdentityByCRN.CRN = core.StringPtr(dedicatedHostGroupIdentityByCRNMap["crn"].(string))

	return dedicatedHostGroupIdentityByCRN
}

func resourceIbmIsDedicatedHostMapToDedicatedHostGroupIdentityByHref(dedicatedHostGroupIdentityByHrefMap map[string]interface{}) vpcv1.DedicatedHostGroupIdentityByHref {
	dedicatedHostGroupIdentityByHref := vpcv1.DedicatedHostGroupIdentityByHref{}

	dedicatedHostGroupIdentityByHref.Href = core.StringPtr(dedicatedHostGroupIdentityByHrefMap["href"].(string))

	return dedicatedHostGroupIdentityByHref
}

func resourceIbmIsDedicatedHostMapToZoneIdentity(zoneIdentityMap map[string]interface{}) *vpcv1.ZoneIdentity {
	zoneIdentity := vpcv1.ZoneIdentity{}

	zoneName, _ := zoneIdentityMap["name"]
	zoneNameStr := zoneName.(string)
	if zoneNameStr != "" {
		zoneIdentity.Name = core.StringPtr(zoneIdentityMap["name"].(string))
	}
	zonehref, _ := zoneIdentityMap["href"]
	zonehrefstr := zonehref.(string)
	if zonehrefstr != "" {
		zoneIdentity.Href = core.StringPtr(zoneIdentityMap["href"].(string))
	}

	return &zoneIdentity
}

func resourceIbmIsDedicatedHostMapToZoneIdentityByName(zoneIdentityByNameMap map[string]interface{}) vpcv1.ZoneIdentityByName {
	zoneIdentityByName := vpcv1.ZoneIdentityByName{}

	zoneIdentityByName.Name = core.StringPtr(zoneIdentityByNameMap["name"].(string))

	return zoneIdentityByName
}

func resourceIbmIsDedicatedHostMapToZoneIdentityByHref(zoneIdentityByHrefMap map[string]interface{}) vpcv1.ZoneIdentityByHref {
	zoneIdentityByHref := vpcv1.ZoneIdentityByHref{}

	zoneIdentityByHref.Href = core.StringPtr(zoneIdentityByHrefMap["href"].(string))

	return zoneIdentityByHref
}

func resourceIbmIsDedicatedHostMapToDedicatedHostPrototypeDedicatedHostByGroup(dedicatedHostPrototypeDedicatedHostByGroupMap map[string]interface{}) vpcv1.DedicatedHostPrototypeDedicatedHostByGroup {
	dedicatedHostPrototypeDedicatedHostByGroup := vpcv1.DedicatedHostPrototypeDedicatedHostByGroup{}

	if dedicatedHostPrototypeDedicatedHostByGroupMap["instance_placement_enabled"] != nil {
		dedicatedHostPrototypeDedicatedHostByGroup.InstancePlacementEnabled = core.BoolPtr(dedicatedHostPrototypeDedicatedHostByGroupMap["instance_placement_enabled"].(bool))
	}
	if dedicatedHostPrototypeDedicatedHostByGroupMap["name"] != nil {
		dedicatedHostPrototypeDedicatedHostByGroup.Name = core.StringPtr(dedicatedHostPrototypeDedicatedHostByGroupMap["name"].(string))
	}
	if len(dedicatedHostPrototypeDedicatedHostByGroupMap["profile"].([]interface{})) != 0 {
		dedicatedHostPrototypeDedicatedHostByGroup.Profile = resourceIbmIsDedicatedHostMapToDedicatedHostProfileIdentity(dedicatedHostPrototypeDedicatedHostByGroupMap["profile"].([]interface{})[0].(map[string]interface{}))
	}
	// TODO: handle Profile of type DedicatedHostProfileIdentity -- not primitive type, not list
	if len(dedicatedHostPrototypeDedicatedHostByGroupMap["resource_group"].([]interface{})) != 0 {
		// TODO: handle ResourceGroup of type ResourceGroupIdentity -- not primitive type, not list
		dedicatedHostPrototypeDedicatedHostByGroup.ResourceGroup = resourceIbmIsDedicatedHostMapToResourceGroupIdentity(dedicatedHostPrototypeDedicatedHostByGroupMap["resource_group"].([]interface{})[0].(map[string]interface{}))
	}
	// TODO: handle Group of type DedicatedHostGroupIdentity -- not primitive type, not list
	if len(dedicatedHostPrototypeDedicatedHostByGroupMap["group"].([]interface{})) != 0 {
		// TODO: handle Group of type DedicatedHostGroupIdentity -- not primitive type, not list
		dedicatedHostPrototypeDedicatedHostByGroup.Group = resourceIbmIsDedicatedHostMapToDedicatedHostGroupIdentity(dedicatedHostPrototypeDedicatedHostByGroupMap["group"].([]interface{})[0].(map[string]interface{}))
	}

	return dedicatedHostPrototypeDedicatedHostByGroup
}

func resourceIbmIsDedicatedHostMapToDedicatedHostPrototypeDedicatedHostByZone(dedicatedHostPrototypeDedicatedHostByZoneMap map[string]interface{}) vpcv1.DedicatedHostPrototypeDedicatedHostByZone {
	dedicatedHostPrototypeDedicatedHostByZone := vpcv1.DedicatedHostPrototypeDedicatedHostByZone{}

	if dedicatedHostPrototypeDedicatedHostByZoneMap["instance_placement_enabled"] != nil {
		dedicatedHostPrototypeDedicatedHostByZone.InstancePlacementEnabled = core.BoolPtr(dedicatedHostPrototypeDedicatedHostByZoneMap["instance_placement_enabled"].(bool))
	}
	if dedicatedHostPrototypeDedicatedHostByZoneMap["name"].(string) != "" {
		dedicatedHostPrototypeDedicatedHostByZone.Name = core.StringPtr(dedicatedHostPrototypeDedicatedHostByZoneMap["name"].(string))
	}
	// TODO: handle Profile of type DedicatedHostProfileIdentity -- not primitive type, not list
	if len(dedicatedHostPrototypeDedicatedHostByZoneMap["profile"].([]interface{})) != 0 {
		dedicatedHostPrototypeDedicatedHostByZone.Profile = resourceIbmIsDedicatedHostMapToDedicatedHostProfileIdentity(dedicatedHostPrototypeDedicatedHostByZoneMap["profile"].([]interface{})[0].(map[string]interface{}))
	}
	if len(dedicatedHostPrototypeDedicatedHostByZoneMap["resource_group"].([]interface{})) != 0 {
		// TODO: handle ResourceGroup of type ResourceGroupIdentity -- not primitive type, not list
		dedicatedHostPrototypeDedicatedHostByZone.ResourceGroup = resourceIbmIsDedicatedHostMapToResourceGroupIdentity(dedicatedHostPrototypeDedicatedHostByZoneMap["resource_group"].([]interface{})[0].(map[string]interface{}))
	}
	if len(dedicatedHostPrototypeDedicatedHostByZoneMap["group"].([]interface{})) != 0 {
		// TODO: handle Group of type DedicatedHostGroupPrototypeDedicatedHostByZoneContext -- not primitive type, not list
		dedicatedHostPrototypeDedicatedHostByZone.Group = resourceIbmIsDedicatedHostMapToDedicatedHostGroupPrototypeDedicatedHostByZoneContext(dedicatedHostPrototypeDedicatedHostByZoneMap["group"].([]interface{})[0].(map[string]interface{}))
	}
	// TODO: handle Zone of type ZoneIdentity -- not primitive type, not list
	if len(dedicatedHostPrototypeDedicatedHostByZoneMap["zone"].([]interface{})) != 0 {
		dedicatedHostPrototypeDedicatedHostByZone.Zone = resourceIbmIsDedicatedHostMapToZoneIdentity(dedicatedHostPrototypeDedicatedHostByZoneMap["zone"].([]interface{})[0].(map[string]interface{}))

	}

	return dedicatedHostPrototypeDedicatedHostByZone
}

func resourceIbmIsDedicatedHostMapToDedicatedHostGroupPrototypeDedicatedHostByZoneContext(dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap map[string]interface{}) *vpcv1.DedicatedHostGroupPrototypeDedicatedHostByZoneContext {
	dedicatedHostGroupPrototypeDedicatedHostByZoneContext := vpcv1.DedicatedHostGroupPrototypeDedicatedHostByZoneContext{}

	if dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap["name"].(string) != "" {
		dedicatedHostGroupPrototypeDedicatedHostByZoneContext.Name = core.StringPtr(dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap["name"].(string))
	}
	if dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap["resource_group"] != nil {
		// TODO: handle ResourceGroup of type ResourceGroupIdentity -- not primitive type, not list
		dedicatedHostGroupPrototypeDedicatedHostByZoneContext.ResourceGroup = resourceIbmIsDedicatedHostMapToResourceGroupIdentity(dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap["resource_group"].([]interface{})[0].(map[string]interface{}))
	}

	return &dedicatedHostGroupPrototypeDedicatedHostByZoneContext
}

func resourceIbmIsDedicatedHostRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	getDedicatedHostOptions := &vpcv1.GetDedicatedHostOptions{}

	getDedicatedHostOptions.SetID(d.Id())

	dedicatedHost, response, err := vpcClient.GetDedicatedHostWithContext(context, getDedicatedHostOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetDedicatedHostWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if err = d.Set("available_memory", intValue(dedicatedHost.AvailableMemory)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting available_memory: %s", err))
	}
	availableVcpuMap := resourceIbmIsDedicatedHostVCPUToMap(*dedicatedHost.AvailableVcpu)
	if err = d.Set("available_vcpu", []map[string]interface{}{availableVcpuMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting available_vcpu: %s", err))
	}
	if err = d.Set("created_at", dedicatedHost.CreatedAt.String()); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("crn", dedicatedHost.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	groupMap := resourceIbmIsDedicatedHostDedicatedHostGroupReferenceToMap(*dedicatedHost.Group)
	if err = d.Set("group", groupMap); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting group: %s", err))
	}
	if err = d.Set("href", dedicatedHost.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	if err = d.Set("instance_placement_enabled", dedicatedHost.InstancePlacementEnabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instance_placement_enabled: %s", err))
	}
	instances := []map[string]interface{}{}
	for _, instancesItem := range dedicatedHost.Instances {
		instancesItemMap := resourceIbmIsDedicatedHostInstanceReferenceToMap(instancesItem)
		instances = append(instances, instancesItemMap)
	}
	if err = d.Set("instances", instances); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting instances: %s", err))
	}
	if err = d.Set("lifecycle_state", dedicatedHost.LifecycleState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting lifecycle_state: %s", err))
	}
	if err = d.Set("memory", intValue(dedicatedHost.Memory)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting memory: %s", err))
	}
	if err = d.Set("name", dedicatedHost.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	//profile := resourceIbmIsDedicatedHostDedicatedHostProfileReferenceToMap(*dedicatedHost.Profile)

	if err = d.Set("profile", *dedicatedHost.Profile.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting profile: %s", err))
	}
	if err = d.Set("provisionable", dedicatedHost.Provisionable); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting provisionable: %s", err))
	}
	//resourceGroupMap := resourceIbmIsDedicatedHostResourceGroupReferenceToMap(*dedicatedHost.ResourceGroup)
	if err = d.Set("resource_group", *dedicatedHost.ResourceGroup.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}
	if err = d.Set("resource_type", dedicatedHost.ResourceType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
	}
	if err = d.Set("socket_count", intValue(dedicatedHost.SocketCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting socket_count: %s", err))
	}
	if err = d.Set("state", dedicatedHost.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}
	supportedInstanceProfiles := []map[string]interface{}{}
	for _, supportedInstanceProfilesItem := range dedicatedHost.SupportedInstanceProfiles {
		supportedInstanceProfilesItemMap := resourceIbmIsDedicatedHostInstanceProfileReferenceToMap(supportedInstanceProfilesItem)
		supportedInstanceProfiles = append(supportedInstanceProfiles, supportedInstanceProfilesItemMap)
	}
	if err = d.Set("supported_instance_profiles", supportedInstanceProfiles); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting supported_instance_profiles: %s", err))
	}
	vcpuMap := resourceIbmIsDedicatedHostVCPUToMap(*dedicatedHost.Vcpu)
	if err = d.Set("vcpu", []map[string]interface{}{vcpuMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting vcpu: %s", err))
	}
	zoneMap := resourceIbmIsDedicatedHostZoneReferenceToMap(*dedicatedHost.Zone)
	if err = d.Set("zone", zoneMap); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting zone: %s", err))
	}

	return nil
}

func resourceIbmIsDedicatedHostDedicatedHostPrototypeToMap(dedicatedHostPrototype vpcv1.DedicatedHostPrototype) map[string]interface{} {
	dedicatedHostPrototypeMap := map[string]interface{}{}

	dedicatedHostPrototypeMap["instance_placement_enabled"] = dedicatedHostPrototype.InstancePlacementEnabled
	dedicatedHostPrototypeMap["name"] = dedicatedHostPrototype.Name
	ProfileMap := resourceIbmIsDedicatedHostDedicatedHostProfileIdentityToMap(*dedicatedHostPrototype.Profile.(*vpcv1.DedicatedHostProfileIdentity))
	dedicatedHostPrototypeMap["profile"] = []map[string]interface{}{ProfileMap}
	if dedicatedHostPrototype.ResourceGroup != nil {
		ResourceGroupMap := resourceIbmIsDedicatedHostResourceGroupIdentityToMap(*dedicatedHostPrototype.ResourceGroup.(*vpcv1.ResourceGroupIdentity))
		dedicatedHostPrototypeMap["resource_group"] = []map[string]interface{}{ResourceGroupMap}
	}
	if dedicatedHostPrototype.Group != nil {
		GroupMap := resourceIbmIsDedicatedHostDedicatedHostGroupIdentityToMap(*dedicatedHostPrototype.Group.(*vpcv1.DedicatedHostGroupIdentity))
		dedicatedHostPrototypeMap["group"] = []map[string]interface{}{GroupMap}
	}
	if dedicatedHostPrototype.Zone != nil {
		ZoneMap := resourceIbmIsDedicatedHostZoneIdentityToMap(*dedicatedHostPrototype.Zone.(*vpcv1.ZoneIdentity))
		dedicatedHostPrototypeMap["zone"] = []map[string]interface{}{ZoneMap}
	}

	return dedicatedHostPrototypeMap
}

func resourceIbmIsDedicatedHostDedicatedHostProfileIdentityToMap(dedicatedHostProfileIdentity vpcv1.DedicatedHostProfileIdentity) map[string]interface{} {
	dedicatedHostProfileIdentityMap := map[string]interface{}{}

	dedicatedHostProfileIdentityMap["name"] = dedicatedHostProfileIdentity.Name
	dedicatedHostProfileIdentityMap["href"] = dedicatedHostProfileIdentity.Href

	return dedicatedHostProfileIdentityMap
}

func resourceIbmIsDedicatedHostDedicatedHostProfileIdentityByNameToMap(dedicatedHostProfileIdentityByName vpcv1.DedicatedHostProfileIdentityByName) map[string]interface{} {
	dedicatedHostProfileIdentityByNameMap := map[string]interface{}{}

	dedicatedHostProfileIdentityByNameMap["name"] = dedicatedHostProfileIdentityByName.Name

	return dedicatedHostProfileIdentityByNameMap
}

func resourceIbmIsDedicatedHostDedicatedHostProfileIdentityByHrefToMap(dedicatedHostProfileIdentityByHref vpcv1.DedicatedHostProfileIdentityByHref) map[string]interface{} {
	dedicatedHostProfileIdentityByHrefMap := map[string]interface{}{}

	dedicatedHostProfileIdentityByHrefMap["href"] = dedicatedHostProfileIdentityByHref.Href

	return dedicatedHostProfileIdentityByHrefMap
}

func resourceIbmIsDedicatedHostResourceGroupIdentityToMap(resourceGroupIdentity vpcv1.ResourceGroupIdentity) map[string]interface{} {
	resourceGroupIdentityMap := map[string]interface{}{}

	resourceGroupIdentityMap["id"] = resourceGroupIdentity.ID

	return resourceGroupIdentityMap
}

func resourceIbmIsDedicatedHostResourceGroupIdentityByIDToMap(resourceGroupIdentityByID vpcv1.ResourceGroupIdentityByID) map[string]interface{} {
	resourceGroupIdentityByIDMap := map[string]interface{}{}

	resourceGroupIdentityByIDMap["id"] = resourceGroupIdentityByID.ID

	return resourceGroupIdentityByIDMap
}

func resourceIbmIsDedicatedHostDedicatedHostGroupIdentityToMap(dedicatedHostGroupIdentity vpcv1.DedicatedHostGroupIdentity) map[string]interface{} {
	dedicatedHostGroupIdentityMap := map[string]interface{}{}

	dedicatedHostGroupIdentityMap["id"] = dedicatedHostGroupIdentity.ID
	dedicatedHostGroupIdentityMap["crn"] = dedicatedHostGroupIdentity.CRN
	dedicatedHostGroupIdentityMap["href"] = dedicatedHostGroupIdentity.Href

	return dedicatedHostGroupIdentityMap
}

func resourceIbmIsDedicatedHostDedicatedHostGroupIdentityByIDToMap(dedicatedHostGroupIdentityByID vpcv1.DedicatedHostGroupIdentityByID) map[string]interface{} {
	dedicatedHostGroupIdentityByIDMap := map[string]interface{}{}

	dedicatedHostGroupIdentityByIDMap["id"] = dedicatedHostGroupIdentityByID.ID

	return dedicatedHostGroupIdentityByIDMap
}

func resourceIbmIsDedicatedHostDedicatedHostGroupIdentityByCRNToMap(dedicatedHostGroupIdentityByCRN vpcv1.DedicatedHostGroupIdentityByCRN) map[string]interface{} {
	dedicatedHostGroupIdentityByCRNMap := map[string]interface{}{}

	dedicatedHostGroupIdentityByCRNMap["crn"] = dedicatedHostGroupIdentityByCRN.CRN

	return dedicatedHostGroupIdentityByCRNMap
}

func resourceIbmIsDedicatedHostDedicatedHostGroupIdentityByHrefToMap(dedicatedHostGroupIdentityByHref vpcv1.DedicatedHostGroupIdentityByHref) map[string]interface{} {
	dedicatedHostGroupIdentityByHrefMap := map[string]interface{}{}

	dedicatedHostGroupIdentityByHrefMap["href"] = dedicatedHostGroupIdentityByHref.Href

	return dedicatedHostGroupIdentityByHrefMap
}

func resourceIbmIsDedicatedHostZoneIdentityToMap(zoneIdentity vpcv1.ZoneIdentity) map[string]interface{} {
	zoneIdentityMap := map[string]interface{}{}

	zoneIdentityMap["name"] = zoneIdentity.Name
	zoneIdentityMap["href"] = zoneIdentity.Href

	return zoneIdentityMap
}

func resourceIbmIsDedicatedHostZoneIdentityByNameToMap(zoneIdentityByName vpcv1.ZoneIdentityByName) map[string]interface{} {
	zoneIdentityByNameMap := map[string]interface{}{}

	zoneIdentityByNameMap["name"] = zoneIdentityByName.Name

	return zoneIdentityByNameMap
}

func resourceIbmIsDedicatedHostZoneIdentityByHrefToMap(zoneIdentityByHref vpcv1.ZoneIdentityByHref) map[string]interface{} {
	zoneIdentityByHrefMap := map[string]interface{}{}

	zoneIdentityByHrefMap["href"] = zoneIdentityByHref.Href

	return zoneIdentityByHrefMap
}

func resourceIbmIsDedicatedHostDedicatedHostPrototypeDedicatedHostByGroupToMap(dedicatedHostPrototypeDedicatedHostByGroup vpcv1.DedicatedHostPrototypeDedicatedHostByGroup) map[string]interface{} {
	dedicatedHostPrototypeDedicatedHostByGroupMap := map[string]interface{}{}

	dedicatedHostPrototypeDedicatedHostByGroupMap["instance_placement_enabled"] = dedicatedHostPrototypeDedicatedHostByGroup.InstancePlacementEnabled
	dedicatedHostPrototypeDedicatedHostByGroupMap["name"] = dedicatedHostPrototypeDedicatedHostByGroup.Name
	ProfileMap := resourceIbmIsDedicatedHostDedicatedHostProfileIdentityToMap(*dedicatedHostPrototypeDedicatedHostByGroup.Profile.(*vpcv1.DedicatedHostProfileIdentity))
	dedicatedHostPrototypeDedicatedHostByGroupMap["profile"] = []map[string]interface{}{ProfileMap}
	if dedicatedHostPrototypeDedicatedHostByGroup.ResourceGroup != nil {
		ResourceGroupMap := resourceIbmIsDedicatedHostResourceGroupIdentityToMap(*dedicatedHostPrototypeDedicatedHostByGroup.ResourceGroup.(*vpcv1.ResourceGroupIdentity))
		dedicatedHostPrototypeDedicatedHostByGroupMap["resource_group"] = []map[string]interface{}{ResourceGroupMap}
	}
	GroupMap := resourceIbmIsDedicatedHostDedicatedHostGroupIdentityToMap(*dedicatedHostPrototypeDedicatedHostByGroup.Group.(*vpcv1.DedicatedHostGroupIdentity))
	dedicatedHostPrototypeDedicatedHostByGroupMap["group"] = []map[string]interface{}{GroupMap}

	return dedicatedHostPrototypeDedicatedHostByGroupMap
}

func resourceIbmIsDedicatedHostDedicatedHostPrototypeDedicatedHostByZoneToMap(dedicatedHostPrototypeDedicatedHostByZone vpcv1.DedicatedHostPrototypeDedicatedHostByZone) map[string]interface{} {
	dedicatedHostPrototypeDedicatedHostByZoneMap := map[string]interface{}{}

	dedicatedHostPrototypeDedicatedHostByZoneMap["instance_placement_enabled"] = dedicatedHostPrototypeDedicatedHostByZone.InstancePlacementEnabled
	dedicatedHostPrototypeDedicatedHostByZoneMap["name"] = dedicatedHostPrototypeDedicatedHostByZone.Name
	ProfileMap := resourceIbmIsDedicatedHostDedicatedHostProfileIdentityToMap(*dedicatedHostPrototypeDedicatedHostByZone.Profile.(*vpcv1.DedicatedHostProfileIdentity))
	dedicatedHostPrototypeDedicatedHostByZoneMap["profile"] = []map[string]interface{}{ProfileMap}
	if dedicatedHostPrototypeDedicatedHostByZone.ResourceGroup != nil {
		ResourceGroupMap := resourceIbmIsDedicatedHostResourceGroupIdentityToMap(*dedicatedHostPrototypeDedicatedHostByZone.ResourceGroup.(*vpcv1.ResourceGroupIdentity))
		dedicatedHostPrototypeDedicatedHostByZoneMap["resource_group"] = []map[string]interface{}{ResourceGroupMap}
	}
	if dedicatedHostPrototypeDedicatedHostByZone.Group != nil {
		GroupMap := resourceIbmIsDedicatedHostDedicatedHostGroupPrototypeDedicatedHostByZoneContextToMap(*dedicatedHostPrototypeDedicatedHostByZone.Group)
		dedicatedHostPrototypeDedicatedHostByZoneMap["group"] = []map[string]interface{}{GroupMap}
	}
	ZoneMap := resourceIbmIsDedicatedHostZoneIdentityToMap(*dedicatedHostPrototypeDedicatedHostByZone.Zone.(*vpcv1.ZoneIdentity))
	dedicatedHostPrototypeDedicatedHostByZoneMap["zone"] = []map[string]interface{}{ZoneMap}

	return dedicatedHostPrototypeDedicatedHostByZoneMap
}

func resourceIbmIsDedicatedHostDedicatedHostGroupPrototypeDedicatedHostByZoneContextToMap(dedicatedHostGroupPrototypeDedicatedHostByZoneContext vpcv1.DedicatedHostGroupPrototypeDedicatedHostByZoneContext) map[string]interface{} {
	dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap := map[string]interface{}{}

	dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap["name"] = dedicatedHostGroupPrototypeDedicatedHostByZoneContext.Name
	if dedicatedHostGroupPrototypeDedicatedHostByZoneContext.ResourceGroup != nil {
		ResourceGroupMap := resourceIbmIsDedicatedHostResourceGroupIdentityToMap(*dedicatedHostGroupPrototypeDedicatedHostByZoneContext.ResourceGroup.(*vpcv1.ResourceGroupIdentity))
		dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap["resource_group"] = []map[string]interface{}{ResourceGroupMap}
	}

	return dedicatedHostGroupPrototypeDedicatedHostByZoneContextMap
}

func resourceIbmIsDedicatedHostVCPUToMap(vCPU vpcv1.Vcpu) map[string]interface{} {
	vCPUMap := map[string]interface{}{}

	vCPUMap["architecture"] = vCPU.Architecture
	vCPUMap["count"] = intValue(vCPU.Count)

	return vCPUMap
}

func resourceIbmIsDedicatedHostDedicatedHostToPrototypeMap(dedicatedHost vpcv1.DedicatedHost) []interface{} {
	dedicatedHostPrototypeMap := map[string]interface{}{}

	if dedicatedHost.InstancePlacementEnabled != nil {
		dedicatedHostPrototypeMap["instance_placement_enabled"] = dedicatedHost.InstancePlacementEnabled
	}
	if dedicatedHost.Name != nil {
		dedicatedHostPrototypeMap["name"] = dedicatedHost.Name
	}
	// TODO: handle Profile of type DedicatedHostProfileIdentity -- not primitive type, not list
	if dedicatedHost.Profile != nil {
		dedicatedHostPrototypeMap["profile"] = resourceIbmIsDedicatedHostDedicatedHostProfileReferenceToMap(*dedicatedHost.Profile)
	}
	if dedicatedHost.ResourceGroup != nil {
		// TODO: handle ResourceGroup of type ResourceGroupIdentity -- not primitive type, not list
		dedicatedHostPrototypeMap["resource_group"] = resourceIbmIsDedicatedHostResourceGroupReferenceToMap(*dedicatedHost.ResourceGroup)
	}
	if dedicatedHost.Group != nil {
		// TODO: handle Group of type DedicatedHostGroupIdentity -- not primitive type, not list
		dedicatedHostPrototypeMap["group"] = resourceIbmIsDedicatedHostDedicatedHostGroupReferenceToMap(*dedicatedHost.Group)
	}
	if dedicatedHost.Zone != nil {
		// TODO: handle Zone of type ZoneIdentity -- not primitive type, not list
		dedicatedHostPrototypeMap["zone"] = resourceIbmIsDedicatedHostZoneReferenceToMap(*dedicatedHost.Zone)
	}
	dedicatedHostPrototypelist := []interface{}{}
	dedicatedHostPrototypelist = append(dedicatedHostPrototypelist, dedicatedHostPrototypeMap)

	return dedicatedHostPrototypelist
}
func resourceIbmIsDedicatedHostDedicatedHostGroupReferenceToMap(dedicatedHostGroupReference vpcv1.DedicatedHostGroupReference) []interface{} {
	dedicatedHostGroupReferenceMap := map[string]interface{}{}

	dedicatedHostGroupReferenceMap["crn"] = dedicatedHostGroupReference.CRN
	if dedicatedHostGroupReference.Deleted != nil {
		DeletedMap := resourceIbmIsDedicatedHostDedicatedHostGroupReferenceDeletedToMap(*dedicatedHostGroupReference.Deleted)
		dedicatedHostGroupReferenceMap["deleted"] = []map[string]interface{}{DeletedMap}
	}
	dedicatedHostGroupReferenceMap["href"] = dedicatedHostGroupReference.Href
	dedicatedHostGroupReferenceMap["id"] = dedicatedHostGroupReference.ID
	dedicatedHostGroupReferenceMap["name"] = dedicatedHostGroupReference.Name
	dedicatedHostGroupReferenceMap["resource_type"] = dedicatedHostGroupReference.ResourceType

	dedicatedHostGroupReferenceint := []interface{}{}
	dedicatedHostGroupReferenceint = append(dedicatedHostGroupReferenceint, dedicatedHostGroupReferenceMap)

	return dedicatedHostGroupReferenceint
}

func resourceIbmIsDedicatedHostDedicatedHostGroupReferenceDeletedToMap(dedicatedHostGroupReferenceDeleted vpcv1.DedicatedHostGroupReferenceDeleted) map[string]interface{} {
	dedicatedHostGroupReferenceDeletedMap := map[string]interface{}{}

	dedicatedHostGroupReferenceDeletedMap["more_info"] = dedicatedHostGroupReferenceDeleted.MoreInfo

	return dedicatedHostGroupReferenceDeletedMap
}

func resourceIbmIsDedicatedHostInstanceReferenceToMap(instanceReference vpcv1.InstanceReference) map[string]interface{} {
	instanceReferenceMap := map[string]interface{}{}

	instanceReferenceMap["crn"] = instanceReference.CRN
	if instanceReference.Deleted != nil {
		DeletedMap := resourceIbmIsDedicatedHostInstanceReferenceDeletedToMap(*instanceReference.Deleted)
		instanceReferenceMap["deleted"] = []map[string]interface{}{DeletedMap}
	}
	instanceReferenceMap["href"] = instanceReference.Href
	instanceReferenceMap["id"] = instanceReference.ID
	instanceReferenceMap["name"] = instanceReference.Name

	return instanceReferenceMap
}

func resourceIbmIsDedicatedHostInstanceReferenceDeletedToMap(instanceReferenceDeleted vpcv1.InstanceReferenceDeleted) map[string]interface{} {
	instanceReferenceDeletedMap := map[string]interface{}{}

	instanceReferenceDeletedMap["more_info"] = instanceReferenceDeleted.MoreInfo

	return instanceReferenceDeletedMap
}

func resourceIbmIsDedicatedHostDedicatedHostProfileReferenceToMap(dedicatedHostProfileReference vpcv1.DedicatedHostProfileReference) []interface{} {
	prof := []interface{}{}

	profile := map[string]interface{}{}
	profile["name"] = dedicatedHostProfileReference.Name
	profile["href"] = dedicatedHostProfileReference.Href
	prof = append(prof, profile)

	return prof
}

func resourceIbmIsDedicatedHostResourceGroupReferenceToMap(resourceGroupReference vpcv1.ResourceGroupReference) []interface{} {

	resgrp := []interface{}{}

	resgroup := map[string]interface{}{}
	if resourceGroupReference.ID != nil && *resourceGroupReference.ID != "" {
		resgroup["id"] = resourceGroupReference.ID
	}

	if resourceGroupReference.Href != nil && *resourceGroupReference.Href != "" {
		resgroup["href"] = *resourceGroupReference.Href
	}

	if resourceGroupReference.Name != nil && *resourceGroupReference.Name != "" {
		resgroup["name"] = resourceGroupReference.Name
	}

	resgrp = append(resgrp, resgroup)

	return resgrp
}

func resourceIbmIsDedicatedHostInstanceProfileReferenceToMap(instanceProfileReference vpcv1.InstanceProfileReference) map[string]interface{} {
	instanceProfileReferenceMap := map[string]interface{}{}

	instanceProfileReferenceMap["href"] = instanceProfileReference.Href
	instanceProfileReferenceMap["name"] = instanceProfileReference.Name

	return instanceProfileReferenceMap
}

func resourceIbmIsDedicatedHostZoneReferenceToMap(zoneReference vpcv1.ZoneReference) []interface{} {
	zoneReferenceMap := map[string]interface{}{}

	zoneReferenceMap["href"] = zoneReference.Href
	zoneReferenceMap["name"] = zoneReference.Name

	zoneReferenceint := []interface{}{}
	zoneReferenceint = append(zoneReferenceint, zoneReferenceMap)
	return zoneReferenceint
}

func resourceIbmIsDedicatedHostUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	updateDedicatedHostOptions := &vpcv1.UpdateDedicatedHostOptions{}

	updateDedicatedHostOptions.SetID(d.Id())

	hasChange := false

	dedicatedHostPrototypemap := map[string]interface{}{}

	if d.HasChange("name") {

		dedicatedHostPrototypemap["name"] = d.Get("name").(interface{})
		hasChange = true
	}
	if d.HasChange("instance_placement_enabled") {

		dedicatedHostPrototypemap["instance_placement_enabled"] = d.Get("instance_placement_enabled").(interface{})
		hasChange = true
	}
	if d.HasChange("profile") {
		dedicatedHostPrototypemap["profile"] = d.Get("profile").(interface{})
		hasChange = true
	}
	if d.HasChange("resource_group") {
		dedicatedHostPrototypemap["resource_group"] = d.Get("resource_group").(interface{})
		hasChange = true
	}
	if d.HasChange("host_group") {
		dedicatedHostPrototypemap["group"] = d.Get("host_group").(interface{})
		hasChange = true
	}
	//log.Println("group final ******", dedicatedHostPrototypemap["group"].([]interface{})[0].(map[string]interface{})["id"].(string))

	if hasChange {
		updateDedicatedHostOptions.SetDedicatedHostPatch(dedicatedHostPrototypemap)
		_, response, err := vpcClient.UpdateDedicatedHostWithContext(context, updateDedicatedHostOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateDedicatedHostWithContext fails %s\n%s", err, response)
			return diag.FromErr(err)
		}
	}

	return resourceIbmIsDedicatedHostRead(context, d, meta)
}

func resourceIbmIsDedicatedHostDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}
	updateDedicatedHostOptions := &vpcv1.UpdateDedicatedHostOptions{}
	dedicatedHostPrototypeMap := map[string]interface{}{}
	dedicatedHostPrototypeMap["instance_placement_enabled"] = core.BoolPtr(false)
	updateDedicatedHostOptions.SetID(d.Id())
	updateDedicatedHostOptions.SetDedicatedHostPatch(dedicatedHostPrototypeMap)
	_, updateresponse, err := vpcClient.UpdateDedicatedHostWithContext(context, updateDedicatedHostOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateDedicatedHostWithContext failed %s\n%s", err, updateresponse)
		return diag.FromErr(err)
	}

	deleteDedicatedHostOptions := &vpcv1.DeleteDedicatedHostOptions{}

	deleteDedicatedHostOptions.SetID(d.Id())

	response, err := vpcClient.DeleteDedicatedHostWithContext(context, deleteDedicatedHostOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteDedicatedHostWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}
	_, err = isWaitForDedicatedHostDelete(vpcClient, d, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func isWaitForDedicatedHostDelete(instanceC *vpcv1.VpcV1, d *schema.ResourceData, id string) (interface{}, error) {

	stateConf := &resource.StateChangeConf{
		Pending: []string{isDedicatedHostDeleting, isDedicatedHostAvailable},
		Target:  []string{isDedicatedHostDeleteDone, ""},
		Refresh: func() (interface{}, string, error) {
			getdhoptions := &vpcv1.GetDedicatedHostOptions{
				ID: &id,
			}
			dedicatedhost, response, err := instanceC.GetDedicatedHost(getdhoptions)
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					return dedicatedhost, isDedicatedHostDeleteDone, nil
				}
				return nil, "", fmt.Errorf("Error Getting Dedicated Host: %s\n%s", err, response)
			}
			if *dedicatedhost.State == isDedicatedHostFailed {
				return dedicatedhost, *dedicatedhost.State, fmt.Errorf("The  Dedicated Host %s failed to delete: %v", d.Id(), err)
			}
			return dedicatedhost, isDedicatedHostDeleting, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
