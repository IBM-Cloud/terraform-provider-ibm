// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceIBMPIPlacementGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIPlacementGroupCreate,
		ReadContext:   resourceIBMPIPlacementGroupRead,
		UpdateContext: resourceIBMPIPlacementGroupUpdate,
		DeleteContext: resourceIBMPIPlacementGroupDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourcePowerUserTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			// Arguments
			Arg_CloudInstanceID: {
				Description:  "The GUID of the service instance associated with an account.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_PlacementGroupName: {
				Description:  "The name of the placement group.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.NoZeroValues,
			},
			Arg_PlacementGroupPolicy: {
				Description:  "The value of the group's affinity policy. Valid values are 'affinity' and 'anti-affinity'.",
				ForceNew:     true,
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{Affinity, AntiAffinity}),
			},
			Arg_UserTags: {
				Computed:    true,
				Description: "List of user tags attached to the resource.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Set:         schema.HashString,
				Type:        schema.TypeSet,
			},

			// Attributes
			Attr_CRN: {
				Computed:    true,
				Description: "The CRN of the resource.",
				Type:        schema.TypeString,
			},
			Attr_Members: {
				Computed:    true,
				Description: "The list of server instances IDs that are members of the placement group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Type:        schema.TypeSet,
			},
			Attr_PlacementGroupID: {
				Computed:    true,
				Description: "The placement group ID.",
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceIBMPIPlacementGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	name := d.Get(Arg_PlacementGroupName).(string)
	policy := d.Get(Arg_PlacementGroupPolicy).(string)
	client := instance.NewIBMPIPlacementGroupClient(ctx, sess, cloudInstanceID)
	body := &models.PlacementGroupCreate{
		Name:   &name,
		Policy: &policy,
	}

	if tags, ok := d.GetOk(Arg_UserTags); ok {
		body.UserTags = flex.FlattenSet(tags.(*schema.Set))
	}

	response, err := client.Create(body)
	if err != nil || response == nil {
		return diag.FromErr(fmt.Errorf("error creating the shared processor pool: %s", err))
	}

	log.Printf("Printing the placement group %+v", &response)
	if response.Crn != "" {
		oldList, newList := d.GetChange(Arg_UserTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, string(response.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on update of pi spp placement group (%s) pi_user_tags during creation: %s", *response.ID, err)
		}
	}

	d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *response.ID))
	return resourceIBMPIPlacementGroupRead(ctx, d, meta)
}

func resourceIBMPIPlacementGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	client := instance.NewIBMPIPlacementGroupClient(ctx, sess, cloudInstanceID)

	response, err := client.Get(parts[1])
	if err != nil {
		log.Printf("[DEBUG]  err %s", err)
		return diag.FromErr(err)
	}

	d.Set(Arg_PlacementGroupName, response.Name)
	d.Set(Arg_PlacementGroupPolicy, response.Policy)
	d.Set(Attr_Members, response.Members)
	d.Set(Attr_PlacementGroupID, response.ID)
	if response.Crn != "" {
		d.Set(Attr_CRN, response.Crn)
		tags, err := flex.GetGlobalTagsUsingCRN(meta, string(response.Crn), "", UserTagType)
		if err != nil {
			log.Printf("Error on get of ibm pi placement group (%s) pi_user_tags: %s", *response.ID, err)
		}
		d.Set(Arg_UserTags, tags)
	}

	return nil
}

func resourceIBMPIPlacementGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_, pgID, err := splitID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	if d.HasChange(Arg_UserTags) {
		if crn, ok := d.GetOk(Attr_CRN); ok {
			oldList, newList := d.GetChange(Arg_UserTags)
			err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, crn.(string), "", UserTagType)
			if err != nil {
				log.Printf("Error on update of pi placement group (%s) pi_user_tags: %s", pgID, err)
			}
		}
	}

	return resourceIBMPIPlacementGroupRead(ctx, d, meta)
}

func resourceIBMPIPlacementGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	cloudInstanceID := parts[0]
	client := instance.NewIBMPIPlacementGroupClient(ctx, sess, cloudInstanceID)
	err = client.Delete(parts[1])
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = isWaitForPIPlacementGroupDeleted(ctx, client, parts[1], d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func isWaitForPIPlacementGroupDeleted(ctx context.Context, client *instance.IBMPIPlacementGroupClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for placement group (%s) to be deleted.", id)

	stateConf := &retry.StateChangeConf{
		Pending:    []string{State_Deleting},
		Target:     []string{State_NotFound},
		Refresh:    isPIPlacementGroupDeleteRefreshFunc(client, id),
		Delay:      10 * time.Second,
		MinTimeout: 30 * time.Second,
		Timeout:    timeout,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isPIPlacementGroupDeleteRefreshFunc(client *instance.IBMPIPlacementGroupClient, id string) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pg, err := client.Get(id)
		if err != nil && strings.Contains(err.Error(), NotFound) {
			log.Printf("The power placement group does not exist")
			return pg, State_NotFound, nil
		}
		return pg, State_Deleting, nil
	}
}
