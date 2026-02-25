// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMIAMTrustedProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamTrustedProfileCreate,
		ReadContext:   resourceIBMIamTrustedProfileRead,
		UpdateContext: resourceIBMIamTrustedProfileUpdate,
		DeleteContext: resourceIBMIamTrustedProfileDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute), // enables HCL `timeouts { read = "…" }`
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the trusted profile. The name is checked for uniqueness. Therefore trusted profiles with the same names can not exist in the same account.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The optional description of the trusted profile. The 'description' property is only available if a description was provided during a create of a trusted profile.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the account that this trusted profile belong to.",
			},
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the unique identifier of the trusted profile. Example:'Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c'.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the trusted profile details object. You need to specify this value when updating the trusted profile to avoid stale updates.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::profile:Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c'.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set contains a date time string of the creation date in ISO format.",
			},
			"modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set contains a date time string of the last modification date in ISO format.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The iam_id of this trusted profile.",
			},
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the IAM template that was used to create an enterprise-managed trusted profile in your account. When returned, this indicates that the trusted profile is created from and managed by a template in the root enterprise account.",
			},
			"assignment_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the assignment that was used to create an enterprise-managed trusted profile in your account. When returned, this indicates that the trusted profile is created from and managed by a template in the root enterprise account.",
			},
			"ims_account_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IMS acount ID of the trusted profile.",
			},
			"ims_user_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IMS user ID of the trusted profile.",
			},
			"history": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "History of the trusted profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timestamp": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the action was triggered.",
						},
						"iam_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IAM ID of the identity which triggered the action.",
						},
						"iam_id_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account of the identity which triggered the action.",
						},
						"action": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Action of the history entry.",
						},
						"params": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Params of the history entry.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message which summarizes the executed action.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMIamTrustedProfileCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createProfileOptions := &iamidentityv1.CreateProfileOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := userDetails.UserAccount

	createProfileOptions.SetName(d.Get("name").(string))
	createProfileOptions.SetAccountID(accountID)
	if _, ok := d.GetOk("description"); ok {
		createProfileOptions.SetDescription(d.Get("description").(string))
	}

	trustedProfile, _, err := iamIdentityClient.CreateProfileWithContext(context, createProfileOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateProfileWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*trustedProfile.ID)

	return resourceIBMIamTrustedProfileRead(context, d, meta)
}

func resourceIBMIamTrustedProfileRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	log.Printf("[UJJK][READ][assert-timeout] create=%s read=%s update=%s delete=%s",
		d.Timeout(schema.TimeoutCreate), d.Timeout(schema.TimeoutRead),
		d.Timeout(schema.TimeoutUpdate), d.Timeout(schema.TimeoutDelete),
	)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "initialize-client")
		log.Printf("[UJJK][READ][init-error] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getProfileOptions := &iamidentityv1.GetProfileOptions{}
	getProfileOptions.SetProfileID(d.Id())

	var trustedProfile *iamidentityv1.TrustedProfile
	var response *core.DetailedResponse

	// ceiling from HCL timeouts (or default)
	timeout := d.Timeout(schema.TimeoutRead)
	if timeout <= 0 {
		timeout = 5 * time.Minute
	}
	start := time.Now()

	// retry ONLY after Create (eventual consistency)
	awaiting := d.IsNewResource()
	log.Printf("[UJJK][READ][start] id=%s timeout=%s awaiting=%t", d.Id(), timeout, awaiting)

	// bounded backoff (used only if awaiting==true)
	initialDelay := 2 * time.Second
	maxDelay := 60 * time.Second
	retryCount, attempt := 0, 1

	for {
		elapsed := time.Since(start)
		if elapsed > timeout {
			log.Printf("[UJJK][READ][timeout] id=%s elapsed=%s awaiting=%t -> clear state", d.Id(), elapsed, awaiting)
			d.SetId("")
			return nil
		}

		log.Printf("[UJJK][READ][attempt] id=%s attempt=%d elapsed=%s retryCount=%d awaiting=%t", d.Id(), attempt, elapsed, retryCount, awaiting)

		trustedProfile, response, err = iamIdentityClient.GetProfileWithContext(context, getProfileOptions)

		var status any = "nil"
		if response != nil {
			status = response.StatusCode
		}

		// success → proceed to mapping
		if err == nil && trustedProfile != nil {
			name := "<nil>"
			if !core.IsNil(trustedProfile.Name) {
				name = *trustedProfile.Name
			}
			etag := "<nil>"
			if !core.IsNil(trustedProfile.EntityTag) {
				etag = *trustedProfile.EntityTag
			}
			crn := "<nil>"
			if !core.IsNil(trustedProfile.CRN) {
				crn = *trustedProfile.CRN
			}
			log.Printf("[UJJK][READ][success] id=%s status=%v name=%s etag=%s crn=%s", d.Id(), status, name, etag, crn)
			break
		}

		// log failure
		if err != nil {
			log.Printf("[UJJK][READ][error] id=%s status=%v err=%v", d.Id(), status, err)
		} else {
			log.Printf("[UJJK][READ][nil-profile] id=%s status=%v", d.Id(), status)
		}

		// 404 handling: retry only if awaiting visibility after create; otherwise clear immediately
		if response != nil && response.StatusCode == 404 {
			if awaiting {
				delay := initialDelay * time.Duration(1<<uint(retryCount))
				if delay > maxDelay {
					delay = maxDelay
				}
				remaining := timeout - elapsed
				log.Printf("[UJJK][READ][404-await][retry] id=%s attempt=%d delay=%s remaining=%s",
					d.Id(), attempt, delay, remaining)
				time.Sleep(delay)
				retryCount++
				attempt++
				continue
			}
			log.Printf("[UJJK][READ][404-steady][clear] id=%s elapsed=%s", d.Id(), elapsed)
			d.SetId("")
			return nil
		}

		// non-404 error: non-retryable in Read
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProfileWithContext failed: %v", err), "ibm_iam_trusted_profile", "read")
		log.Printf("[UJJK][READ][non-retryable] id=%s status=%v msg=%s", d.Id(), status, tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", trustedProfile.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-name").GetDiag()
	}
	if !core.IsNil(trustedProfile.Description) {
		if err = d.Set("description", trustedProfile.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("account_id", trustedProfile.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-account_id").GetDiag()
	}
	if err = d.Set("profile_id", trustedProfile.ID); err != nil {
		err = fmt.Errorf("Error setting id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-id").GetDiag()
	}
	if err = d.Set("entity_tag", trustedProfile.EntityTag); err != nil {
		err = fmt.Errorf("Error setting entity_tag: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-entity_tag").GetDiag()
	}
	if err = d.Set("crn", trustedProfile.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(trustedProfile.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(trustedProfile.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-created_at").GetDiag()
		}
	}
	if !core.IsNil(trustedProfile.ModifiedAt) {
		if err = d.Set("modified_at", flex.DateTimeToString(trustedProfile.ModifiedAt)); err != nil {
			err = fmt.Errorf("Error setting modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-modified_at").GetDiag()
		}
	}
	if err = d.Set("iam_id", trustedProfile.IamID); err != nil {
		err = fmt.Errorf("Error setting iam_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-iam_id").GetDiag()
	}
	if !core.IsNil(trustedProfile.TemplateID) {
		if err = d.Set("template_id", trustedProfile.TemplateID); err != nil {
			err = fmt.Errorf("Error setting template_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-template_id").GetDiag()
		}
	}
	if !core.IsNil(trustedProfile.AssignmentID) {
		if err = d.Set("assignment_id", trustedProfile.AssignmentID); err != nil {
			err = fmt.Errorf("Error setting assignment_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-assignment_id").GetDiag()
		}
	}
	if !core.IsNil(trustedProfile.ImsAccountID) {
		if err = d.Set("ims_account_id", flex.IntValue(trustedProfile.ImsAccountID)); err != nil {
			err = fmt.Errorf("Error setting ims_account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-ims_account_id").GetDiag()
		}
	}
	if !core.IsNil(trustedProfile.ImsUserID) {
		if err = d.Set("ims_user_id", flex.IntValue(trustedProfile.ImsUserID)); err != nil {
			err = fmt.Errorf("Error setting ims_user_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-ims_user_id").GetDiag()
		}
	}
	history := []map[string]interface{}{}
	if !core.IsNil(trustedProfile.History) {
		for _, historyItem := range trustedProfile.History {
			historyItemMap, err := EnityHistoryRecordToMap(&historyItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "history-to-map").GetDiag()
			}
			history = append(history, historyItemMap)
		}
	}
	if err = d.Set("history", history); err != nil {
		err = fmt.Errorf("Error setting history: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "read", "set-history").GetDiag()
	}

	log.Printf("[UJJK][READ][done] id=%s elapsed=%s", d.Id(), time.Since(start))
	return nil
}

func resourceIBMIamTrustedProfileUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateProfileOptions := &iamidentityv1.UpdateProfileOptions{}

	updateProfileOptions.SetIfMatch("*")
	updateProfileOptions.SetProfileID(d.Id())
	if _, ok := d.GetOk("name"); ok {
		updateProfileOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		updateProfileOptions.SetDescription(d.Get("description").(string))
	}

	_, _, err = iamIdentityClient.UpdateProfileWithContext(context, updateProfileOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateProfileWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamTrustedProfileRead(context, d, meta)
}

func resourceIBMIamTrustedProfileDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_trusted_profile", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteProfileOptions := &iamidentityv1.DeleteProfileOptions{}
	deleteProfileOptions.SetProfileID(d.Id())

	_, err = iamIdentityClient.DeleteProfileWithContext(context, deleteProfileOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteProfileWithContext failed: %s", err.Error()), "ibm_iam_trusted_profile", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}
