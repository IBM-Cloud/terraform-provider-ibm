// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func ResourceIBMIamIdp() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamIdpCreate,
		ReadContext:   resourceIBMIamIdpRead,
		UpdateContext: resourceIBMIamIdpUpdate,
		DeleteContext: resourceIBMIamIdpDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Account where the IdP resides in.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Speaking name of the Identity Provider.",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Type of the IDP. Valid values: saml, appid, ldap.",
			},
			"active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Defines if the IDP is active (enabled) for all accounts. Default during creation is true.",
			},
			"share_scope": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of targets which can consume the IdP.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the account or enterprise.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of share scope. Valid values: account, enterprise.",
						},
					},
				},
			},
			"properties": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Properties of the IDP (stored plain-text). Required for SAML type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Identity Provider configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"xml_import": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Flag indicating if IdP should be imported from metadata.xml.",
									},
									"entity_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "SAML IDP entity ID. Required for SAML when xml_import is false.",
									},
									"redirect_binding_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Redirect binding URL. Required for SAML when xml_import is false.",
									},
									"want_request_signed": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if IDP wants requests to be signed.",
									},
									"logout_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "SAML IDP logout URL (optional).",
									},
								},
							},
						},
						"sp": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Service Provider configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"want_assertion_signed": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if SP wants assertions to be signed.",
									},
									"want_response_signed": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if SP wants responses to be signed.",
									},
									"encrypt_response": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates if responses should be encrypted.",
									},
									"idp_initiated_login_enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enables IDP-initiated login.",
									},
									"logout_url_enabled_when_available": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enables logout URL when available.",
									},
									"idp_initiated_urls": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "URLs for IDP-initiated login.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"secrets": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Sensitive:   true,
				Description: "Secrets of the IDP (stored encrypted). Required for SAML type (may be empty to auto-generate SP certificates).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Identity Provider secrets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"xml_import": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Flag indicating if secrets should be imported from metadata.xml.",
									},
								},
							},
						},
						"sp": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Service Provider secrets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
			// Computed fields
			"idp_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the IDP.",
			},
			"entity_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the IDP. Used for updates to avoid stale writes.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the IDP was created.",
			},
			"modified_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the IDP was last modified.",
			},
		},
	}
}

func resourceIBMIamIdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createOpts := iamIdentityClient.NewCreateIdpOptions(
		d.Get("account_id").(string),
		d.Get("name").(string),
		d.Get("type").(string),
	)

	if v, ok := d.GetOkExists("active"); ok {
		createOpts.SetActive(v.(bool))
	}
	if v, ok := d.GetOk("share_scope"); ok {
		createOpts.SetShareScope(expandShareScope(v.([]interface{})))
	}
	if v, ok := d.GetOk("properties"); ok {
		props := expandIdpCreateProperties(v.([]interface{}))
		if props != nil {
			createOpts.SetProperties(props)
		}
	}
	// secrets block is required by the API for SAML IDPs; always send it (may be empty).
	secrets := expandIdpCreateSecrets(d.Get("secrets").([]interface{}))
	createOpts.SetSecrets(secrets)

	idp, _, err := iamIdentityClient.CreateIdpWithContext(ctx, createOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateIdpWithContext failed: %s", err.Error()), "ibm_iam_idp", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*idp.IdpID)
	return resourceIBMIamIdpRead(ctx, d, meta)
}

func resourceIBMIamIdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getOpts := iamIdentityClient.NewGetIdpOptions(d.Id())
	idp, response, err := iamIdentityClient.GetIdpWithContext(ctx, getOpts)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIdpWithContext failed: %s", err.Error()), "ibm_iam_idp", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if idp.IdpID != nil {
		d.Set("idp_id", idp.IdpID)
	}
	if idp.AccountID != nil {
		d.Set("account_id", idp.AccountID)
	}
	if idp.Name != nil {
		d.Set("name", idp.Name)
	}
	if idp.Type != nil {
		d.Set("type", idp.Type)
	}
	if idp.Active != nil {
		d.Set("active", idp.Active)
	}
	if idp.EntityTag != nil {
		d.Set("entity_tag", idp.EntityTag)
	}
	if idp.CreatedAt != nil {
		d.Set("created_at", flex.DateTimeToString(idp.CreatedAt))
	}
	if idp.ModifiedAt != nil {
		d.Set("modified_at", flex.DateTimeToString(idp.ModifiedAt))
	}
	if idp.ShareScope != nil {
		d.Set("share_scope", flattenShareScope(idp.ShareScope))
	}
	return nil
}

func resourceIBMIamIdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	entityTag := d.Get("entity_tag").(string)
	updateOpts := iamIdentityClient.NewUpdateIdpOptions(d.Id(), entityTag)

	if d.HasChange("name") {
		updateOpts.SetName(d.Get("name").(string))
	}
	if d.HasChange("active") {
		updateOpts.SetActive(d.Get("active").(bool))
	}
	if d.HasChange("share_scope") {
		updateOpts.SetShareScope(expandShareScope(d.Get("share_scope").([]interface{})))
	}
	if d.HasChange("properties") {
		props := expandIdpUpdateProperties(d.Get("properties").([]interface{}))
		if props != nil {
			updateOpts.SetProperties(props)
		}
	}
	if d.HasChange("secrets") {
		secrets := expandIdpUpdateSecrets(d.Get("secrets").([]interface{}))
		updateOpts.SetSecrets(secrets)
	}

	_, _, err = iamIdentityClient.UpdateIdpWithContext(ctx, updateOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateIdpWithContext failed: %s", err.Error()), "ibm_iam_idp", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamIdpRead(ctx, d, meta)
}

func resourceIBMIamIdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteOpts := iamIdentityClient.NewDeleteIdpOptions(d.Id())
	_, err = iamIdentityClient.DeleteIdpWithContext(ctx, deleteOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteIdpWithContext failed: %s", err.Error()), "ibm_iam_idp", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}

// FlattenShareScopeForTest is an exported wrapper used by unit tests.
func FlattenShareScopeForTest(scopes []iamidentityv1.ShareScope) []map[string]interface{} {
	return flattenShareScope(scopes)
}

// ExpandShareScopeForTest is an exported wrapper used by unit tests.
func ExpandShareScopeForTest(l []interface{}) []iamidentityv1.ShareScope {
	return expandShareScope(l)
}

// expandShareScope converts a Terraform list to []iamidentityv1.ShareScope.
func expandShareScope(l []interface{}) []iamidentityv1.ShareScope {
	result := make([]iamidentityv1.ShareScope, 0, len(l))
	for _, item := range l {
		m := item.(map[string]interface{})
		s := iamidentityv1.ShareScope{}
		if v, ok := m["id"].(string); ok && v != "" {
			s.ID = core.StringPtr(v)
		}
		if v, ok := m["type"].(string); ok && v != "" {
			s.Type = core.StringPtr(v)
		}
		result = append(result, s)
	}
	return result
}

// flattenShareScope converts []iamidentityv1.ShareScope to a Terraform-compatible list.
func flattenShareScope(scopes []iamidentityv1.ShareScope) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(scopes))
	for _, s := range scopes {
		m := map[string]interface{}{}
		if s.ID != nil {
			m["id"] = *s.ID
		}
		if s.Type != nil {
			m["type"] = *s.Type
		}
		result = append(result, m)
	}
	return result
}

// expandIdpCreateProperties converts properties list to *iamidentityv1.CreateIdpRequestProperties.
// The API requires both idp and sp sub-objects; sp defaults to an empty struct if not supplied.
func expandIdpCreateProperties(l []interface{}) *iamidentityv1.CreateIdpRequestProperties {
	if len(l) == 0 || l[0] == nil {
		return nil
	}
	m := l[0].(map[string]interface{})
	props := &iamidentityv1.CreateIdpRequestProperties{
		Sp: &iamidentityv1.CreateIdpRequestPropertiesSp{},
	}

	if idpList, ok := m["idp"].([]interface{}); ok && len(idpList) > 0 && idpList[0] != nil {
		idpMap := idpList[0].(map[string]interface{})
		idpProps := &iamidentityv1.CreateIdpRequestPropertiesIdp{}
		if v, ok := idpMap["xml_import"].(bool); ok {
			idpProps.XMLImport = core.BoolPtr(v)
		}
		if v, ok := idpMap["entity_id"].(string); ok && v != "" {
			idpProps.EntityID = core.StringPtr(v)
		}
		if v, ok := idpMap["redirect_binding_url"].(string); ok && v != "" {
			idpProps.RedirectBindingURL = core.StringPtr(v)
		}
		if v, ok := idpMap["want_request_signed"].(bool); ok {
			idpProps.WantRequestSigned = core.BoolPtr(v)
		}
		if v, ok := idpMap["logout_url"].(string); ok && v != "" {
			idpProps.LogoutURL = core.StringPtr(v)
		}
		props.Idp = idpProps
	}

	if spList, ok := m["sp"].([]interface{}); ok && len(spList) > 0 && spList[0] != nil {
		spMap := spList[0].(map[string]interface{})
		spProps := &iamidentityv1.CreateIdpRequestPropertiesSp{}
		if v, ok := spMap["want_assertion_signed"].(bool); ok {
			spProps.WantAssertionSigned = core.BoolPtr(v)
		}
		if v, ok := spMap["want_response_signed"].(bool); ok {
			spProps.WantResponseSigned = core.BoolPtr(v)
		}
		if v, ok := spMap["encrypt_response"].(bool); ok {
			spProps.EncryptResponse = core.BoolPtr(v)
		}
		if v, ok := spMap["idp_initiated_login_enabled"].(bool); ok {
			spProps.IdpInitiatedLoginEnabled = core.BoolPtr(v)
		}
		if v, ok := spMap["logout_url_enabled_when_available"].(bool); ok {
			spProps.LogoutURLEnabledWhenAvailable = core.BoolPtr(v)
		}
		if v, ok := spMap["idp_initiated_urls"].([]interface{}); ok {
			urls := make([]string, 0, len(v))
			for _, u := range v {
				if s, ok := u.(string); ok {
					urls = append(urls, s)
				}
			}
			spProps.IdpInitiatedUrls = urls
		}
		props.Sp = spProps
	}

	return props
}

// expandIdpUpdateProperties converts properties list to *iamidentityv1.UpdateIdpRequestProperties.
// The API requires sp to be present; defaults to an empty struct if not supplied.
func expandIdpUpdateProperties(l []interface{}) *iamidentityv1.UpdateIdpRequestProperties {
	if len(l) == 0 || l[0] == nil {
		return nil
	}
	m := l[0].(map[string]interface{})
	props := &iamidentityv1.UpdateIdpRequestProperties{
		Sp: &iamidentityv1.UpdateIdpRequestPropertiesSp{},
	}

	if idpList, ok := m["idp"].([]interface{}); ok && len(idpList) > 0 && idpList[0] != nil {
		idpMap := idpList[0].(map[string]interface{})
		idpProps := &iamidentityv1.UpdateIdpRequestPropertiesIdp{}
		if v, ok := idpMap["entity_id"].(string); ok && v != "" {
			idpProps.EntityID = core.StringPtr(v)
		}
		if v, ok := idpMap["redirect_binding_url"].(string); ok && v != "" {
			idpProps.RedirectBindingURL = core.StringPtr(v)
		}
		if v, ok := idpMap["want_request_signed"].(bool); ok {
			idpProps.WantRequestSigned = core.BoolPtr(v)
		}
		if v, ok := idpMap["logout_url"].(string); ok && v != "" {
			idpProps.LogoutURL = core.StringPtr(v)
		}
		props.Idp = idpProps
	}

	if spList, ok := m["sp"].([]interface{}); ok && len(spList) > 0 && spList[0] != nil {
		spMap := spList[0].(map[string]interface{})
		spProps := &iamidentityv1.UpdateIdpRequestPropertiesSp{}
		if v, ok := spMap["want_assertion_signed"].(bool); ok {
			spProps.WantAssertionSigned = core.BoolPtr(v)
		}
		if v, ok := spMap["want_response_signed"].(bool); ok {
			spProps.WantResponseSigned = core.BoolPtr(v)
		}
		if v, ok := spMap["encrypt_response"].(bool); ok {
			spProps.EncryptResponse = core.BoolPtr(v)
		}
		if v, ok := spMap["idp_initiated_login_enabled"].(bool); ok {
			spProps.IdpInitiatedLoginEnabled = core.BoolPtr(v)
		}
		if v, ok := spMap["logout_url_enabled_when_available"].(bool); ok {
			spProps.LogoutURLEnabledWhenAvailable = core.BoolPtr(v)
		}
		if v, ok := spMap["idp_initiated_urls"].([]interface{}); ok {
			urls := make([]string, 0, len(v))
			for _, u := range v {
				if s, ok := u.(string); ok {
					urls = append(urls, s)
				}
			}
			spProps.IdpInitiatedUrls = urls
		}
		props.Sp = spProps
	}

	return props
}

// expandIdpCreateSecrets converts the secrets list to *iamidentityv1.CreateIdpRequestSecrets.
// Always returns a non-nil pointer (may be empty structs) because the API requires the secrets
// field for SAML IDPs even when no certificates are supplied.
func expandIdpCreateSecrets(l []interface{}) *iamidentityv1.CreateIdpRequestSecrets {
	secrets := &iamidentityv1.CreateIdpRequestSecrets{
		Idp: &iamidentityv1.CreateIdpRequestSecretsIdp{},
		Sp:  &iamidentityv1.CreateIdpRequestSecretsSp{},
	}
	if len(l) == 0 || l[0] == nil {
		return secrets
	}
	m := l[0].(map[string]interface{})

	if idpList, ok := m["idp"].([]interface{}); ok && len(idpList) > 0 && idpList[0] != nil {
		idpMap := idpList[0].(map[string]interface{})
		idpSecrets := &iamidentityv1.CreateIdpRequestSecretsIdp{}
		if v, ok := idpMap["xml_import"].(bool); ok {
			idpSecrets.XMLImport = core.BoolPtr(v)
		}
		secrets.Idp = idpSecrets
	}
	return secrets
}

// expandIdpUpdateSecrets converts the secrets list to *iamidentityv1.UpdateIdpRequestSecrets.
func expandIdpUpdateSecrets(l []interface{}) *iamidentityv1.UpdateIdpRequestSecrets {
	secrets := &iamidentityv1.UpdateIdpRequestSecrets{
		Idp: &iamidentityv1.UpdateIdpRequestSecretsIdp{},
		Sp:  &iamidentityv1.UpdateIdpRequestSecretsSp{},
	}
	if len(l) == 0 || l[0] == nil {
		return secrets
	}
	return secrets
}
