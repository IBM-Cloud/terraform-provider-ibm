// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/mtlsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	cisMtlsAppName       = "app_name"
	cisMtlsDuration      = "duration"
	cisMtlsURL           = "url"
	cisMtlsRuleCertVal   = "rule_cert"
	cisMtlsRuleCommonVal = "rule_common"
	cisMtlsPolicyName    = "policy_name"
	cisMtlsPolicyAction  = "policy_action"
	cisMtlsAppCreatedAt  = "app_created_at"
	cisMtlsAppUpdatedAt  = "app_updated_at"
	cisMtlsPolCreatedAt  = "pol_created_at"
	cisMtlsPolUpdatedAt  = "pol_updated_at"
)

func ResourceIBMCISMtlsApp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCISMtlsAppCreate,
		ReadContext:   resourceIBMCISMtlsAppRead,
		UpdateContext: resourceIBMCISMtlsAppUpdate,
		DeleteContext: resourceIBMCISMtlsAppDelete,
		Importer:      &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisMtlsHostName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host associated with",
			},
			cisMtlsAppName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "App Name",
			},
			cisMtlsDuration: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "24h",
				Description: "Duration for app validatidity",
			},
			cisMtlsURL: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "App URL",
			},
			cisMtlsPolicyName: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "MTLS-Policy",
				Description: "Policy Name",
			},
			cisMtlsPolicyAction: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "non_identity",
				Description: "Policy Action",
			},
			cisMtlsAppCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Created At",
			},
			cisMtlsAppUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Created At",
			},
			cisMtlsPolCreatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Created At",
			},
			cisMtlsPolUpdatedAt: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Created At",
			},
		},
	}
}
func resourceIBMCISMtlsAppCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess))
	}
	crn := d.Get(cisID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(d.Get(cisDomainID).(string))

	sess.Crn = core.StringPtr(crn)
	OptionsApp := sess.NewCreateAccessApplicationOptions(zoneID)

	if app_val, ok := d.GetOk(cisMtlsAppName); ok {
		OptionsApp.SetName(app_val.(string))
	}

	if host_val, ok := d.GetOk(cisMtlsHostName); ok {
		OptionsApp.SetDomain(host_val.(string))
	}

	if dur_val, ok := d.GetOk(cisMtlsDuration); ok {
		OptionsApp.SetSessionDuration(dur_val.(string))
	}

	resultApp, responseApp, operationErrApp := sess.CreateAccessApplication(OptionsApp)

	if operationErrApp != nil || resultApp == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error creating access application  %v %v %v", operationErrApp, resultApp, responseApp))
	}

	d.SetId(flex.ConvertCisToTfThreeVar(*resultApp.Result.ID, zoneID, crn))

	//save appID
	appId := *resultApp.Result.ID

	// Create an access policy
	policyRuleModel := &mtlsv1.PolicyRulePolicyCertRule{
		Certificate: map[string]interface{}{"certifcate": "CA root certificate"},
	}
	policyCnModel := &mtlsv1.PolicyCnRuleCommonName{
		CommonName: core.StringPtr("Access CA"),
	}
	policyModel := &mtlsv1.PolicyRulePolicyCnRule{
		CommonName: policyCnModel,
	}
	optionsPolicy := sess.NewCreateAccessPolicyOptions(zoneID, appId)

	// get policy name and action/decsion
	if policy_val, ok := d.GetOk(cisMtlsPolicyName); ok {
		optionsPolicy.SetName(policy_val.(string))
	}
	if action_val, ok := d.GetOk(cisMtlsPolicyAction); ok {
		optionsPolicy.SetDecision(action_val.(string))
	}

	optionsPolicy.SetInclude([]mtlsv1.PolicyRuleIntf{policyModel, policyRuleModel})
	resultPolicy, responsePolicy, operationErrPolicy := sess.CreateAccessPolicy(optionsPolicy)

	if operationErrPolicy != nil || resultPolicy == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error creating app policy  %v", responsePolicy))
	}

	d.SetId(flex.ConvertCisToTfThreeVar(*resultApp.Result.ID, zoneID, *resultPolicy.Result.ID))
	return resourceIBMCISMtlsAppRead(context, d, meta)

}
func resourceIBMCISMtlsAppRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess))
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)
	appID, zoneID, policyID, _ := flex.ConvertTfToCisThreeVar(d.Id())
	getAppOptions := sess.NewGetAccessApplicationOptions(zoneID, appID)
	getAppResult, getAppResp, getAppErr := sess.GetAccessApplication(getAppOptions)

	if getAppErr != nil || getAppResult == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting app deatil  %v", getAppResp))
	}

	getPolicyOptions := sess.NewGetAccessPolicyOptions(zoneID, appID, policyID)
	getPolicyResult, getPolicyResp, getPolicyErr := sess.GetAccessPolicy(getPolicyOptions)

	if getPolicyErr != nil || getPolicyResult == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting Policy  detail  %v", getPolicyResp))
	}

	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisMtlsAppCreatedAt, *getAppResult.Result.CreatedAt)
	d.Set(cisMtlsAppUpdatedAt, *getAppResult.Result.UpdatedAt)
	d.Set(cisMtlsPolCreatedAt, *getPolicyResult.Result.CreatedAt)
	d.Set(cisMtlsPolUpdatedAt, *getPolicyResult.Result.CreatedAt)

	return nil
}
func resourceIBMCISMtlsAppUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess))
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)
	appID, zoneID, _, _ := flex.ConvertTfToCisThreeVar(d.Id())

	if d.HasChange(cisMtlsAppName) ||
		d.HasChange(cisMtlsURL) ||
		d.HasChange(cisMtlsPolicyName) || d.HasChange(cisMtlsPolicyAction) ||
		d.HasChange(cisMtlsDuration) {

		updateOptionApp := sess.NewUpdateAccessApplicationOptions(zoneID, appID)

		if app_name, ok := d.GetOk(cisMtlsAppName); ok {
			updateOptionApp.SetName(app_name.(string))
		}
		if app_url, ok := d.GetOk(cisMtlsURL); ok {
			updateOptionApp.SetName(app_url.(string))
		}

		if duration_val, ok := d.GetOk(cisMtlsDuration); ok {
			updateOptionApp.SetName(duration_val.(string))
		}
		updateResultApp, updateRespApp, updateErrApp := sess.UpdateAccessApplication(updateOptionApp)
		if updateErrApp != nil {
			if updateRespApp != nil {
				d.SetId("")
				return nil
			}
			return diag.FromErr(fmt.Errorf("[ERROR] Error while updating the applicatoin values %v", updateResultApp))
		}

		optionsPolicy := sess.NewCreateAccessPolicyOptions(zoneID, appID)
		if policy_name, ok := d.GetOk(cisMtlsPolicyName); ok {
			optionsPolicy.SetName(policy_name.(string))
		}
		if action_name, ok := d.GetOk(cisMtlsPolicyAction); ok {
			optionsPolicy.SetName(action_name.(string))
		}

		resultPolicy, responsePolicy, operationErrPolicy := sess.CreateAccessPolicy(optionsPolicy)

		if operationErrPolicy != nil {
			if responsePolicy != nil {
				d.SetId("")
				return nil
			}
			return diag.FromErr(fmt.Errorf("[ERROR] Error while updating the applicatoin values %v", resultPolicy))
		}

	}

	return resourceIBMCISMtlsAppRead(context, d, meta)
}
func resourceIBMCISMtlsAppDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).CisMtlsSession()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error while getting the CisMtlsSession() %s %v", err, sess))
	}

	crn := d.Get(cisID).(string)
	zoneID := d.Get(cisDomainID).(string)
	sess.Crn = core.StringPtr(crn)
	listAccOpt := sess.NewListAccessApplicationsOptions(zoneID)
	listAccResult, listAccResp, listAccErr := sess.ListAccessApplications(listAccOpt)
	if listAccErr != nil {
		if listAccResp != nil && listAccResp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error While getting application detail for deletion"))
	}
	// Delete an access applications
	for _, appId := range listAccResult.Result {
		// List access policy
		listOptPolicy := sess.NewListAccessPoliciesOptions(zoneID, *appId.ID)
		listResultPolicy, listRespPolicy, listErrPolicy := sess.ListAccessPolicies(listOptPolicy)
		if listErrPolicy != nil {
			if listRespPolicy != nil && listRespPolicy.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return diag.FromErr(fmt.Errorf("[ERROR] Error While getting policy detail for deletion"))
		}
		// Delete access policy
		for _, policyId := range listResultPolicy.Result {
			delOptPolicy := sess.NewDeleteAccessPolicyOptions(zoneID, *appId.ID, *policyId.ID)
			_, delRespPolicy, delErrPolicy := sess.DeleteAccessPolicy(delOptPolicy)
			if delErrPolicy != nil {
				if delRespPolicy != nil && delRespPolicy.StatusCode == 404 {
					d.SetId("")
					return nil
				}
				return diag.FromErr(fmt.Errorf("[ERROR] Error While deleting the policy"))
			}

		}
		delAccOpt := sess.NewDeleteAccessApplicationOptions(zoneID, *appId.ID)
		_, delAccResp, delAccErr := sess.DeleteAccessApplication(delAccOpt)
		if delAccErr != nil {
			if delAccResp != nil && delAccResp.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			return diag.FromErr(fmt.Errorf("[ERROR] Error While deleting the app"))
		}

	}

	return nil

}
