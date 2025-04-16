// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMISReservationActivate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISReservationActivateCreate,
		ReadContext:   resourceIBMISReservationActivateRead,
		DeleteContext: resourceIBMISReservationActivateDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isReservation: &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique identifier for this reservation.",
			},
			isReservationAffinityPolicy: &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The affinity policy to use for this reservation",
			},
			isReservationCapacity: &schema.Schema{
				Type:        schema.TypeList,
				ForceNew:    true,
				Computed:    true,
				Description: "The capacity reservation configuration to use",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationCapacityTotal: &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
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
				Computed:    true,
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
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The maximum number of recent backups to keep. If unspecified, there will be no maximum.",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reservation name",
			},
			isReservationProfile: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The profile used for this reservation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isReservationProfileName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this virtual server instance profile.",
						},
						isReservationProfileResourceType: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type of the profile.",
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
							Computed:    true,
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
							Computed:    true,
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
				Computed:    true,
				Description: "The globally unique name for this zone.",
			},
		},
	}
}
func resourceIBMISReservationActivateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Get(isReservation).(string)
	activateReservationOptions := &vpcv1.ActivateReservationOptions{
		ID: core.StringPtr(id),
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = sess.ActivateReservationWithContext(context, activateReservationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ActivateReservationWithContext failed: %s", err.Error()), "ibm_is_reservation_activate", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	log.Printf("[INFO] Reservation activated: %s", id)
	d.SetId(id)

	return resourceIBMISReservationActivateRead(context, d, meta)
}

func resourceIBMISReservationActivateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()

	sess, err := vpcClient(meta)
	defer func() {

		log.Println("stacktrace from panic: \n", err, string(debug.Stack()))

	}()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "initialize-client")
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
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetReservationWithContext failed: %s", err.Error()), "ibm_is_reservation_activate", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(reservation.AffinityPolicy) {
		if err = d.Set("affinity_policy", reservation.AffinityPolicy); err != nil {
			err = fmt.Errorf("Error setting affinity_policy: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-affinity_policy").GetDiag()
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
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-capacity").GetDiag()
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
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-committed_use").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(reservation.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("crn", reservation.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-crn").GetDiag()
	}

	if err = d.Set("href", reservation.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-href").GetDiag()
	}

	if err = d.Set("lifecycle_state", reservation.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-lifecycle_state").GetDiag()
	}

	if !core.IsNil(reservation.Name) {
		if err = d.Set("name", reservation.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-name").GetDiag()
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
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-profile").GetDiag()
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
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-resource_group").GetDiag()
		}
	}

	if err = d.Set("resource_type", reservation.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-resource_type").GetDiag()
	}

	if err = d.Set("status", reservation.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-status").GetDiag()
	}

	if reservation.StatusReasons != nil {
		srLen := len(reservation.StatusReasons)
		srList := []vpcv1.ReservationStatusReason{}

		for i := 0; i < srLen; i++ {
			srList = append(srList, reservation.StatusReasons[i])
		}
		if err = d.Set("status_reasons", srList); err != nil {
			err = fmt.Errorf("Error setting status_reasons: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-status_reasons").GetDiag()
		}
	}

	if reservation.Zone != nil && reservation.Zone.Name != nil {
		if err = d.Set("zone", reservation.Zone.Name); err != nil {
			err = fmt.Errorf("Error setting zone: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_reservation_activate", "read", "set-zone").GetDiag()
		}
	}
	return nil
}

func resourceIBMISReservationActivateDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
