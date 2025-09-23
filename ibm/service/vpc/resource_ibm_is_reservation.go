// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isReservation                            = "reservation"
	isReservationName                        = "name"
	isReservationAffinityPolicy              = "affinity_policy"
	isReservationCapacity                    = "capacity"
	isReservationCapacityTotal               = "total"
	isReservationCommittedUse                = "committed_use"
	isReservationComittedUseExpirationPolicy = "expiration_policy"
	isReservationComittedUseTerm             = "term"
	isReservationProfile                     = "profile"
	isReservationProfileName                 = "name"
	isReservationProfileResourceType         = "resource_type"
	isReservationResourceGroup               = "resource_group"
	isReservationZone                        = "zone"

	isReservationCapacityAllocated        = "allocated"
	isReservationCapacityAvailable        = "available"
	isReservationCapacityStatus           = "status"
	isReservationCapacityUsed             = "used"
	isReservationCommittedUseExpirationAt = "expiration_at"
	isReservationCreatedAt                = "created_at"
	isReservationCrn                      = "crn"
	isReservationHref                     = "href"
	isReservationId                       = "id"
	isReservationLifecycleState           = "lifecycle_state"
	isReservationProfileHref              = "href"
	isReservationResourceGroupHref        = "href"
	isReservationResourceGroupId          = "id"
	isReservationResourceGroupName        = "name"
	isReservationResourceType             = "resource_type"
	isReservationStatusReasons            = "status_reasons"
	isReservationStatusReasonCode         = "code"
	isReservationStatusReasonMessage      = "message"
	isReservationStatusReasonMoreInfo     = "more_info"
	isReservationZoneHref                 = "href"
	isReservationZoneName                 = "name"
	isReservationStatus                   = "status"
)

func ResourceIBMISReservation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISReservationCreate,
		ReadContext:   resourceIBMISReservationRead,
		UpdateContext: resourceIBMISReservationUpdate,
		DeleteContext: resourceIBMISReservationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isReservationAffinityPolicy: &schema.Schema{
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_reservation", isReservationAffinityPolicy),
				Description:  "The affinity policy to use for this reservation",
			},
			isReservationCapacity: &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The capacity reservation configuration to use",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationCapacityTotal: &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The total amount to use for this capacity reservation.",
						},
						isReservationCapacityAllocated: &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount allocated to this capacity reservation.",
						},
						isReservationCapacityAvailable: &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of this capacity reservation available for new attachments.",
						},
						isReservationCapacityUsed: &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of this capacity reservation used by existing attachments.",
						},
						isReservationCapacityStatus: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the capacity reservation.",
						},
					},
				},
			},
			isReservationCommittedUse: &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The committed use configuration to use for this reservation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationComittedUseExpirationPolicy: &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The maximum number of days to keep each backup after creation.",
						},
						isReservationComittedUseTerm: &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_reservation", isReservationComittedUseTerm),
							Description:  "The maximum number of recent backups to keep. If unspecified, there will be no maximum.",
						},
						isReservationCommittedUseExpirationAt: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expiration date and time for this committed use reservation.",
						},
					},
				},
			},
			isReservationCreatedAt: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the reservation was created.",
			},
			isReservationCrn: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this reservation.",
			},
			isReservationHref: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this reservation.",
			},
			isReservationId: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this reservation.",
			},
			isReservationLifecycleState: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of this reservation.",
			},
			isReservationName: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_reservation", isReservationName),
				Description:  "Reservation name",
			},
			isReservationProfile: &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The profile to use for this reservation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationProfileName: &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_reservation", isReservationProfileName),
							Description:  "The globally unique name for this virtual server instance profile.",
						},
						isReservationProfileResourceType: &schema.Schema{
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validate.InvokeValidator("ibm_is_reservation", isReservationProfileResourceType),
							Description:  "The resource type of the profile.",
						},
						isReservationProfileHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this virtual server instance profile.",
						},
					},
				},
			},
			isReservationResourceGroup: &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The committed use configuration to use for this reservation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationResourceGroupHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this resource group.",
						},
						isReservationResourceGroupId: &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The unique identifier for this resource group",
						},
						isReservationResourceGroupName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this resource group.",
						},
					},
				},
			},
			isReservationResourceType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			isReservationStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the reservation.",
			},
			isReservationStatusReasons: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The committed use configuration to use for this reservation",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationStatusReasonCode: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: " snake case string succinctly identifying the status reason.",
						},
						isReservationStatusReasonMessage: &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "An explanation of the status reason",
						},
						isReservationStatusReasonMoreInfo: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Link to documentation about this status reason.",
						},
					},
				},
			},
			isReservationZone: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The globally unique name for this zone.",
			},
		},
	}
}

func ResourceIBMISReservationValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	affinityPolicy := "automatic, restricted"
	term := "one_year, three_year"
	resourceType := "bare_metal_server_profile, instance_profile"

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservationName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservationAffinityPolicy,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              affinityPolicy})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservationComittedUseTerm,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              term})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservationProfileResourceType,
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              resourceType})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservationProfileName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isReservationZone,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63})
	ibmISVPCResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_reservation", Schema: validateSchema}
	return &ibmISVPCResourceValidator
}

func resourceIBMISReservationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	createReservationOptions := &vpcv1.CreateReservationOptions{}
	if _, ok := d.GetOk(isReservationCapacity); ok {
		resCapacity := d.Get(isReservationCapacity + ".0").(map[string]interface{})
		reservationCapacityPrototype := &vpcv1.ReservationCapacityPrototype{}

		if resCapacity[isReservationCapacityTotal] != nil {
			reservationCapacityPrototype.Total = core.Int64Ptr(int64(resCapacity[isReservationCapacityTotal].(int)))
		}
		createReservationOptions.Capacity = reservationCapacityPrototype
	}

	if _, ok := d.GetOk(isReservationCommittedUse); ok {
		resCommittedUse := d.Get(isReservationCommittedUse + ".0").(map[string]interface{})
		reservationCommittedUsePrototype := &vpcv1.ReservationCommittedUsePrototype{}

		if resCommittedUse[isReservationComittedUseTerm] != nil && resCommittedUse[isReservationComittedUseTerm].(string) != "" {
			reservationCommittedUsePrototype.Term = core.StringPtr(resCommittedUse[isReservationComittedUseTerm].(string))
		}
		if resCommittedUse[isReservationComittedUseExpirationPolicy] != nil && resCommittedUse[isReservationComittedUseExpirationPolicy].(string) != "" {
			reservationCommittedUsePrototype.ExpirationPolicy = core.StringPtr(resCommittedUse[isReservationComittedUseExpirationPolicy].(string))
		}
		createReservationOptions.CommittedUse = reservationCommittedUsePrototype
	}

	if _, ok := d.GetOk(isReservationProfile); ok {
		resProfile := d.Get(isReservationProfile + ".0").(map[string]interface{})
		reservationProfilePrototype := &vpcv1.ReservationProfilePrototype{}

		if resProfile[isReservationProfileName] != nil && resProfile[isReservationProfileName].(string) != "" {
			reservationProfilePrototype.Name = core.StringPtr(resProfile[isReservationProfileName].(string))
		}
		if resProfile[isReservationProfileResourceType] != nil && resProfile[isReservationProfileResourceType].(string) != "" {
			reservationProfilePrototype.ResourceType = core.StringPtr(resProfile[isReservationProfileResourceType].(string))
		}
		createReservationOptions.Profile = reservationProfilePrototype
	}

	if _, ok := d.GetOk(isReservationResourceGroup); ok {
		resGroup := d.Get(isReservationResourceGroup + ".0").(map[string]interface{})
		if resGroup[isReservationResourceGroupId] != nil && resGroup[isReservationResourceGroupId].(string) != "" {
			createReservationOptions.ResourceGroup = &vpcv1.ResourceGroupIdentity{
				ID: core.StringPtr(resGroup[isReservationResourceGroupId].(string)),
			}
		}
	}

	if zone, ok := d.GetOk(isReservationZone); ok {
		if zone.(string) != "" {
			createReservationOptions.Zone = &vpcv1.ZoneIdentity{Name: core.StringPtr(zone.(string))}
		}
	}

	if name, ok := d.GetOk(isReservationName); ok {
		if name.(string) != "" {
			createReservationOptions.Name = core.StringPtr(name.(string))
		}
	}

	if affPol, ok := d.GetOk(isReservationAffinityPolicy); ok {
		if affPol.(string) != "" {
			createReservationOptions.AffinityPolicy = core.StringPtr(affPol.(string))
		}
	}
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	reservation, _, err := sess.CreateReservationWithContext(context, createReservationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateReservationWithContext failed: %s", err.Error()), "ibm_is_reservation", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*reservation.ID)
	log.Printf("[INFO] Reservation : %s", *reservation.ID)

	return resourceIBMISReservationRead(context, d, meta)
}

func resourceIBMISReservationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	getReservationOptions := &vpcv1.GetReservationOptions{
		ID: &id,
	}
	reservation, response, err := sess.GetReservationWithContext(context, getReservationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetReservationWithContext failed: %s", err.Error()), "ibm_is_reservation", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(reservation.AffinityPolicy) {
		if err = d.Set("affinity_policy", reservation.AffinityPolicy); err != nil {
			err = fmt.Errorf("Error setting affinity_policy: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-affinity_policy").GetDiag()
		}
	}

	if reservation.Capacity != nil {
		capacityMap := []map[string]interface{}{}
		finalList := map[string]interface{}{}

		if reservation.Capacity.Allocated != nil {
			finalList[isReservationCapacityAllocated] = flex.IntValue(reservation.Capacity.Allocated)
		}
		if reservation.Capacity.Available != nil {
			finalList[isReservationCapacityAvailable] = flex.IntValue(reservation.Capacity.Available)
		}
		if reservation.Capacity.Total != nil {
			finalList[isReservationCapacityTotal] = flex.IntValue(reservation.Capacity.Total)
		}
		if reservation.Capacity.Used != nil {
			finalList[isReservationCapacityUsed] = flex.IntValue(reservation.Capacity.Used)
		}
		if reservation.Capacity.Status != nil {
			finalList[isReservationCapacityStatus] = reservation.Capacity.Status
		}
		capacityMap = append(capacityMap, finalList)
		if err = d.Set("capacity", capacityMap); err != nil {
			err = fmt.Errorf("Error setting capacity: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-capacity").GetDiag()
		}
	}

	if reservation.CommittedUse != nil {
		committedUseMap := []map[string]interface{}{}
		finalList := map[string]interface{}{}

		if reservation.CommittedUse.ExpirationAt != nil {
			finalList[isReservationCommittedUseExpirationAt] = flex.DateTimeToString(reservation.CommittedUse.ExpirationAt)
		}
		if reservation.CommittedUse.ExpirationPolicy != nil {
			finalList[isReservationComittedUseExpirationPolicy] = *reservation.CommittedUse.ExpirationPolicy
		}
		if reservation.CommittedUse.Term != nil {
			finalList[isReservationComittedUseTerm] = *reservation.CommittedUse.Term
		}
		committedUseMap = append(committedUseMap, finalList)
		if err = d.Set("committed_use", committedUseMap); err != nil {
			err = fmt.Errorf("Error setting committed_use: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-committed_use").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(reservation.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("crn", reservation.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-crn").GetDiag()
	}

	if err = d.Set("href", reservation.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-href").GetDiag()
	}

	if err = d.Set("lifecycle_state", reservation.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-lifecycle_state").GetDiag()
	}

	if !core.IsNil(reservation.Name) {
		if err = d.Set("name", reservation.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-name").GetDiag()
		}
	}

	if reservation.Profile != nil {
		profileMap := []map[string]interface{}{}
		finalList := map[string]interface{}{}

		profileItem := reservation.Profile.(*vpcv1.ReservationProfile)

		if profileItem.Href != nil {
			finalList[isReservationProfileHref] = profileItem.Href
		}
		if profileItem.Name != nil {
			finalList[isReservationProfileName] = profileItem.Name
		}
		if profileItem.ResourceType != nil {
			finalList[isReservationProfileResourceType] = profileItem.ResourceType
		}
		profileMap = append(profileMap, finalList)
		if err = d.Set("profile", profileMap); err != nil {
			err = fmt.Errorf("Error setting profile: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-profile").GetDiag()
		}
	}

	if reservation.ResourceGroup != nil {
		rgMap := []map[string]interface{}{}
		finalList := map[string]interface{}{}

		if reservation.ResourceGroup.Href != nil {
			finalList[isReservationResourceGroupHref] = reservation.ResourceGroup.Href
		}
		if reservation.ResourceGroup.ID != nil {
			finalList[isReservationResourceGroupId] = reservation.ResourceGroup.ID
		}
		if reservation.ResourceGroup.Name != nil {
			finalList[isReservationResourceGroupName] = reservation.ResourceGroup.Name
		}
		rgMap = append(rgMap, finalList)
		if err = d.Set("resource_group", rgMap); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-resource_group").GetDiag()
		}
	}

	if err = d.Set("resource_type", reservation.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-resource_type").GetDiag()
	}

	if err = d.Set("status", reservation.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-status").GetDiag()
	}

	if reservation.StatusReasons != nil {
		srLen := len(reservation.StatusReasons)
		srList := []vpcv1.ReservationStatusReason{}

		for i := 0; i < srLen; i++ {
			srList = append(srList, reservation.StatusReasons[i])
		}
		if err = d.Set("status_reasons", srList); err != nil {
			err = fmt.Errorf("Error setting status_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-status_reasons").GetDiag()
		}
	}

	if reservation.Zone != nil && reservation.Zone.Name != nil {
		if err = d.Set("zone", reservation.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "read", "set-zone").GetDiag()
		}
	}
	return nil
}

func resourceIBMISReservationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	hasChanged := false
	name := ""
	affPol := ""

	reservationPatchModel := &vpcv1.ReservationPatch{}
	if d.HasChange(isReservationName) {
		name = d.Get(isReservationName).(string)
		if name != "" {
			reservationPatchModel.Name = &name
			hasChanged = true
		}
	}
	if d.HasChange(isReservationAffinityPolicy) {
		affPol = d.Get(isReservationAffinityPolicy).(string)
		if affPol != "" {
			reservationPatchModel.AffinityPolicy = &affPol
			hasChanged = true
		}
	}
	if d.HasChange(isReservationCapacity) {
		capacityIntf := d.Get(isReservationCapacity)
		capacityMap := capacityIntf.([]interface{})[0].(map[string]interface{})
		if d.HasChange(isReservationCapacity + ".0." + isReservationCapacityTotal) {
			if totalIntf, ok := capacityMap[isReservationCapacityTotal]; ok {
				reservationPatchModel.Capacity = &vpcv1.ReservationCapacityPatch{
					Total: core.Int64Ptr(int64(totalIntf.(int))),
				}
				hasChanged = true
			}
		}
	}
	if d.HasChange(isReservationCommittedUse) {
		committedUseIntf := d.Get(isReservationCommittedUse)
		committedUseMap := committedUseIntf.([]interface{})[0].(map[string]interface{})
		cuPatch := &vpcv1.ReservationCommittedUsePatch{}
		if d.HasChange(isReservationCommittedUse + ".0." + isReservationComittedUseExpirationPolicy) {
			if expPolIntf, ok := committedUseMap[isReservationComittedUseExpirationPolicy]; ok {
				if expPolIntf.(string) != "" {
					cuPatch.ExpirationPolicy = core.StringPtr(string(expPolIntf.(string)))
				}
				hasChanged = true
			}
		}
		if d.HasChange(isReservationCommittedUse + ".0." + isReservationComittedUseTerm) {
			if termIntf, ok := committedUseMap[isReservationComittedUseTerm]; ok {
				cuPatch.Term = core.StringPtr(string(termIntf.(string)))
			}
			hasChanged = true
		}
		reservationPatchModel.CommittedUse = cuPatch
	}
	if d.HasChange(isReservationProfile) {
		profileIntf := d.Get(isReservationProfile)
		profileMap := profileIntf.([]interface{})[0].(map[string]interface{})
		profPatch := &vpcv1.ReservationProfilePatch{}
		if d.HasChange(isReservationProfile + ".0." + isReservationProfileName) {
			if profNameIntf, ok := profileMap[isReservationProfileName]; ok {
				if profNameIntf.(string) != "" {
					profPatch.Name = core.StringPtr(string(profNameIntf.(string)))
				}
				hasChanged = true
			}
		}
		if d.HasChange(isReservationProfile + ".0." + isReservationProfileResourceType) {
			if resTypeIntf, ok := profileMap[isReservationProfileResourceType]; ok {
				if resTypeIntf.(string) != "" {
					profPatch.ResourceType = core.StringPtr(string(resTypeIntf.(string)))
				}
				hasChanged = true
			}
		}
		reservationPatchModel.Profile = profPatch
	}
	if hasChanged {
		reservationPatch, err := reservationPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("reservationPatchModel.AsPatch failed: %s", err.Error()), "ibm_is_reservation", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		updateReservationOptions := &vpcv1.UpdateReservationOptions{}
		updateReservationOptions.ReservationPatch = reservationPatch
		updateReservationOptions.ID = core.StringPtr(d.Id())
		_, _, err = sess.UpdateReservationWithContext(context, updateReservationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateReservationWithContext failed: %s", err.Error()), "ibm_is_reservation", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMISReservationRead(context, d, meta)
}

func resourceIBMISReservationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteReservationOptions := &vpcv1.DeleteReservationOptions{
		ID: &id,
	}
	_, _, err = sess.DeleteReservationWithContext(context, deleteReservationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteReservationWithContext failed: %s", err.Error()), "ibm_is_reservation", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}
