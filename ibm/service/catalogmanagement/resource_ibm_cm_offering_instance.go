// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package catalogmanagement

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
)

func ResourceIBMCmOfferingInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCmOfferingInstanceCreate,
		ReadContext:   resourceIBMCmOfferingInstanceRead,
		UpdateContext: resourceIBMCmOfferingInstanceUpdate,
		DeleteContext: resourceIBMCmOfferingInstanceDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"x_auth_refresh_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IAM Refresh token.",
			},
			"rev": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloudant revision.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "url reference to this object.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "platform CRN for this instance.",
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the label for this instance.",
			},
			"catalog_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Catalog ID this instance was created from.",
			},
			"offering_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Offering ID this instance was created from.",
			},
			"kind_format": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the format this instance has (helm, operator, ova...).",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version this instance was installed from (semver - not version id).",
			},
			"version_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version id this instance was installed from (version id - not semver).",
			},
			"cluster_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster ID.",
			},
			"cluster_region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster region (e.g., us-south).",
			},
			"cluster_namespaces": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of target namespaces to install into.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cluster_all_namespaces": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "designate to install into all namespaces.",
			},
			"schematics_workspace_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the schematics workspace, for offering instances provisioned through schematics.",
			},
			"install_plan": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster.",
			},
			"channel": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Channel to pin the operator subscription to.",
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "date and time create.",
			},
			"updated": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "date and time updated.",
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of metadata values for this offering instance.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"resource_group_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the resource group to provision the offering instance into.",
			},
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "String location of OfferingInstance deployment.",
			},
			"disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if Resource Controller has disabled this instance.",
			},
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account this instance is owned by.",
			},
			"last_operation": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "the last operation performed and status.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"operation": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "last operation performed.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "state after the last operation performed.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "additional information about the last operation.",
						},
						"transaction_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "transaction id from the last operation.",
						},
						"updated": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date and time last updated.",
						},
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Error code from the last operation, if applicable.",
						},
					},
				},
			},
			"kind_target": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The target kind for the installed software version.",
			},
			"sha": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The digest value of the installed software version.",
			},
		},
	}
}

func resourceIBMCmOfferingInstanceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createOfferingInstanceOptions := &catalogmanagementv1.CreateOfferingInstanceOptions{}

	createOfferingInstanceOptions.SetXAuthRefreshToken(d.Get("x_auth_refresh_token").(string))
	if _, ok := d.GetOk("rev"); ok {
		createOfferingInstanceOptions.SetRev(d.Get("rev").(string))
	}
	if _, ok := d.GetOk("url"); ok {
		createOfferingInstanceOptions.SetURL(d.Get("url").(string))
	}
	if _, ok := d.GetOk("crn"); ok {
		createOfferingInstanceOptions.SetCRN(d.Get("crn").(string))
	}
	if _, ok := d.GetOk("label"); ok {
		createOfferingInstanceOptions.SetLabel(d.Get("label").(string))
	}
	if _, ok := d.GetOk("catalog_id"); ok {
		createOfferingInstanceOptions.SetCatalogID(d.Get("catalog_id").(string))
	}
	if _, ok := d.GetOk("offering_id"); ok {
		createOfferingInstanceOptions.SetOfferingID(d.Get("offering_id").(string))
	}
	if _, ok := d.GetOk("kind_format"); ok {
		createOfferingInstanceOptions.SetKindFormat(d.Get("kind_format").(string))
	}
	if _, ok := d.GetOk("version"); ok {
		createOfferingInstanceOptions.SetVersion(d.Get("version").(string))
	}
	if _, ok := d.GetOk("version_id"); ok {
		createOfferingInstanceOptions.SetVersionID(d.Get("version_id").(string))
	}
	if _, ok := d.GetOk("cluster_id"); ok {
		createOfferingInstanceOptions.SetClusterID(d.Get("cluster_id").(string))
	}
	if _, ok := d.GetOk("cluster_region"); ok {
		createOfferingInstanceOptions.SetClusterRegion(d.Get("cluster_region").(string))
	}
	if _, ok := d.GetOk("cluster_namespaces"); ok {
		createOfferingInstanceOptions.SetClusterNamespaces(d.Get("cluster_namespaces").([]string))
	}
	if _, ok := d.GetOk("cluster_all_namespaces"); ok {
		createOfferingInstanceOptions.SetClusterAllNamespaces(d.Get("cluster_all_namespaces").(bool))
	}
	if _, ok := d.GetOk("schematics_workspace_id"); ok {
		createOfferingInstanceOptions.SetSchematicsWorkspaceID(d.Get("schematics_workspace_id").(string))
	}
	if _, ok := d.GetOk("install_plan"); ok {
		createOfferingInstanceOptions.SetInstallPlan(d.Get("install_plan").(string))
	}
	if _, ok := d.GetOk("channel"); ok {
		createOfferingInstanceOptions.SetChannel(d.Get("channel").(string))
	}
	if _, ok := d.GetOk("created"); ok {
		fmtDateTimeCreated, err := core.ParseDateTime(d.Get("created").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOfferingInstanceOptions.SetCreated(&fmtDateTimeCreated)
	}
	if _, ok := d.GetOk("updated"); ok {
		fmtDateTimeUpdated, err := core.ParseDateTime(d.Get("updated").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		createOfferingInstanceOptions.SetUpdated(&fmtDateTimeUpdated)
	}
	if _, ok := d.GetOk("metadata"); ok {
		// TODO: Add code to handle map container: Metadata
	}
	if _, ok := d.GetOk("resource_group_id"); ok {
		createOfferingInstanceOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		createOfferingInstanceOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("disabled"); ok {
		createOfferingInstanceOptions.SetDisabled(d.Get("disabled").(bool))
	}
	if _, ok := d.GetOk("account"); ok {
		createOfferingInstanceOptions.SetAccount(d.Get("account").(string))
	}
	if _, ok := d.GetOk("last_operation"); ok {
		lastOperationModel, err := resourceIBMCmOfferingInstanceMapToOfferingInstanceLastOperation(d.Get("last_operation.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createOfferingInstanceOptions.SetLastOperation(lastOperationModel)
	}
	if _, ok := d.GetOk("kind_target"); ok {
		createOfferingInstanceOptions.SetKindTarget(d.Get("kind_target").(string))
	}
	if _, ok := d.GetOk("sha"); ok {
		createOfferingInstanceOptions.SetSha(d.Get("sha").(string))
	}

	offeringInstance, response, err := catalogManagementClient.CreateOfferingInstanceWithContext(context, createOfferingInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateOfferingInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateOfferingInstanceWithContext failed %s\n%s", err, response))
	}

	d.SetId(*offeringInstance.ID)

	return resourceIBMCmOfferingInstanceRead(context, d, meta)
}

func resourceIBMCmOfferingInstanceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getOfferingInstanceOptions := &catalogmanagementv1.GetOfferingInstanceOptions{}

	getOfferingInstanceOptions.SetInstanceIdentifier(d.Id())

	offeringInstance, response, err := catalogManagementClient.GetOfferingInstanceWithContext(context, getOfferingInstanceOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetOfferingInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetOfferingInstanceWithContext failed %s\n%s", err, response))
	}

	// if err = d.Set("x_auth_refresh_token", getOfferingInstanceOptions.XAuthRefreshToken); err != nil {
	// 	return diag.FromErr(fmt.Errorf("Error setting x_auth_refresh_token: %s", err))
	// }
	if err = d.Set("rev", offeringInstance.Rev); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rev: %s", err))
	}
	if err = d.Set("url", offeringInstance.URL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting url: %s", err))
	}
	if err = d.Set("crn", offeringInstance.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("label", offeringInstance.Label); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting label: %s", err))
	}
	if err = d.Set("catalog_id", offeringInstance.CatalogID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting catalog_id: %s", err))
	}
	if err = d.Set("offering_id", offeringInstance.OfferingID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting offering_id: %s", err))
	}
	if err = d.Set("kind_format", offeringInstance.KindFormat); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting kind_format: %s", err))
	}
	if err = d.Set("version", offeringInstance.Version); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}
	if err = d.Set("version_id", offeringInstance.VersionID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version_id: %s", err))
	}
	if err = d.Set("cluster_id", offeringInstance.ClusterID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cluster_id: %s", err))
	}
	if err = d.Set("cluster_region", offeringInstance.ClusterRegion); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cluster_region: %s", err))
	}
	if offeringInstance.ClusterNamespaces != nil {
		if err = d.Set("cluster_namespaces", offeringInstance.ClusterNamespaces); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cluster_namespaces: %s", err))
		}
	}
	if err = d.Set("cluster_all_namespaces", offeringInstance.ClusterAllNamespaces); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cluster_all_namespaces: %s", err))
	}
	if err = d.Set("schematics_workspace_id", offeringInstance.SchematicsWorkspaceID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting schematics_workspace_id: %s", err))
	}
	if err = d.Set("install_plan", offeringInstance.InstallPlan); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting install_plan: %s", err))
	}
	if err = d.Set("channel", offeringInstance.Channel); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting channel: %s", err))
	}
	if err = d.Set("created", flex.DateTimeToString(offeringInstance.Created)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if err = d.Set("updated", flex.DateTimeToString(offeringInstance.Updated)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}
	if offeringInstance.Metadata != nil {
		// TODO: handle Metadata of type TypeMap -- not primitive type, not list
	}
	if err = d.Set("resource_group_id", offeringInstance.ResourceGroupID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group_id: %s", err))
	}
	if err = d.Set("location", offeringInstance.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if err = d.Set("disabled", offeringInstance.Disabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting disabled: %s", err))
	}
	if err = d.Set("account", offeringInstance.Account); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account: %s", err))
	}
	if offeringInstance.LastOperation != nil {
		lastOperationMap, err := resourceIBMCmOfferingInstanceOfferingInstanceLastOperationToMap(offeringInstance.LastOperation)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("last_operation", []map[string]interface{}{lastOperationMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_operation: %s", err))
		}
	}
	if err = d.Set("kind_target", offeringInstance.KindTarget); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting kind_target: %s", err))
	}
	if err = d.Set("sha", offeringInstance.Sha); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting sha: %s", err))
	}

	return nil
}

func resourceIBMCmOfferingInstanceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	putOfferingInstanceOptions := &catalogmanagementv1.PutOfferingInstanceOptions{}

	putOfferingInstanceOptions.SetInstanceIdentifier(d.Id())
	putOfferingInstanceOptions.SetXAuthRefreshToken(d.Get("x_auth_refresh_token").(string))
	if _, ok := d.GetOk("rev"); ok {
		putOfferingInstanceOptions.SetRev(d.Get("rev").(string))
	}
	if _, ok := d.GetOk("url"); ok {
		putOfferingInstanceOptions.SetURL(d.Get("url").(string))
	}
	if _, ok := d.GetOk("crn"); ok {
		putOfferingInstanceOptions.SetCRN(d.Get("crn").(string))
	}
	if _, ok := d.GetOk("label"); ok {
		putOfferingInstanceOptions.SetLabel(d.Get("label").(string))
	}
	if _, ok := d.GetOk("catalog_id"); ok {
		putOfferingInstanceOptions.SetCatalogID(d.Get("catalog_id").(string))
	}
	if _, ok := d.GetOk("offering_id"); ok {
		putOfferingInstanceOptions.SetOfferingID(d.Get("offering_id").(string))
	}
	if _, ok := d.GetOk("kind_format"); ok {
		putOfferingInstanceOptions.SetKindFormat(d.Get("kind_format").(string))
	}
	if _, ok := d.GetOk("version"); ok {
		putOfferingInstanceOptions.SetVersion(d.Get("version").(string))
	}
	if _, ok := d.GetOk("version_id"); ok {
		putOfferingInstanceOptions.SetVersionID(d.Get("version_id").(string))
	}
	if _, ok := d.GetOk("cluster_id"); ok {
		putOfferingInstanceOptions.SetClusterID(d.Get("cluster_id").(string))
	}
	if _, ok := d.GetOk("cluster_region"); ok {
		putOfferingInstanceOptions.SetClusterRegion(d.Get("cluster_region").(string))
	}
	if _, ok := d.GetOk("cluster_namespaces"); ok {
		putOfferingInstanceOptions.SetClusterNamespaces(d.Get("cluster_namespaces").([]string))
	}
	if _, ok := d.GetOk("cluster_all_namespaces"); ok {
		putOfferingInstanceOptions.SetClusterAllNamespaces(d.Get("cluster_all_namespaces").(bool))
	}
	if _, ok := d.GetOk("schematics_workspace_id"); ok {
		putOfferingInstanceOptions.SetSchematicsWorkspaceID(d.Get("schematics_workspace_id").(string))
	}
	if _, ok := d.GetOk("install_plan"); ok {
		putOfferingInstanceOptions.SetInstallPlan(d.Get("install_plan").(string))
	}
	if _, ok := d.GetOk("channel"); ok {
		putOfferingInstanceOptions.SetChannel(d.Get("channel").(string))
	}
	if _, ok := d.GetOk("created"); ok {
		fmtDateTimeCreated, err := core.ParseDateTime(d.Get("created").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		putOfferingInstanceOptions.SetCreated(&fmtDateTimeCreated)
	}
	if _, ok := d.GetOk("updated"); ok {
		fmtDateTimeUpdated, err := core.ParseDateTime(d.Get("updated").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		putOfferingInstanceOptions.SetUpdated(&fmtDateTimeUpdated)
	}
	if _, ok := d.GetOk("metadata"); ok {
		// TODO: Non-primitive types that are not models or lists
	}
	if _, ok := d.GetOk("resource_group_id"); ok {
		putOfferingInstanceOptions.SetResourceGroupID(d.Get("resource_group_id").(string))
	}
	if _, ok := d.GetOk("location"); ok {
		putOfferingInstanceOptions.SetLocation(d.Get("location").(string))
	}
	if _, ok := d.GetOk("disabled"); ok {
		putOfferingInstanceOptions.SetDisabled(d.Get("disabled").(bool))
	}
	if _, ok := d.GetOk("account"); ok {
		putOfferingInstanceOptions.SetAccount(d.Get("account").(string))
	}
	if _, ok := d.GetOk("last_operation"); ok {
		lastOperation, err := resourceIBMCmOfferingInstanceMapToOfferingInstanceLastOperation(d.Get("last_operation.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		putOfferingInstanceOptions.SetLastOperation(lastOperation)
	}
	if _, ok := d.GetOk("kind_target"); ok {
		putOfferingInstanceOptions.SetKindTarget(d.Get("kind_target").(string))
	}
	if _, ok := d.GetOk("sha"); ok {
		putOfferingInstanceOptions.SetSha(d.Get("sha").(string))
	}

	_, response, err := catalogManagementClient.PutOfferingInstanceWithContext(context, putOfferingInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] PutOfferingInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PutOfferingInstanceWithContext failed %s\n%s", err, response))
	}

	return resourceIBMCmOfferingInstanceRead(context, d, meta)
}

func resourceIBMCmOfferingInstanceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	catalogManagementClient, err := meta.(conns.ClientSession).CatalogManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteOfferingInstanceOptions := &catalogmanagementv1.DeleteOfferingInstanceOptions{}

	deleteOfferingInstanceOptions.SetInstanceIdentifier(d.Id())

	response, err := catalogManagementClient.DeleteOfferingInstanceWithContext(context, deleteOfferingInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteOfferingInstanceWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteOfferingInstanceWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIBMCmOfferingInstanceMapToOfferingInstanceLastOperation(modelMap map[string]interface{}) (*catalogmanagementv1.OfferingInstanceLastOperation, error) {
	model := &catalogmanagementv1.OfferingInstanceLastOperation{}
	if modelMap["operation"] != nil && modelMap["operation"].(string) != "" {
		model.Operation = core.StringPtr(modelMap["operation"].(string))
	}
	if modelMap["state"] != nil && modelMap["state"].(string) != "" {
		model.State = core.StringPtr(modelMap["state"].(string))
	}
	if modelMap["message"] != nil && modelMap["message"].(string) != "" {
		model.Message = core.StringPtr(modelMap["message"].(string))
	}
	if modelMap["transaction_id"] != nil && modelMap["transaction_id"].(string) != "" {
		model.TransactionID = core.StringPtr(modelMap["transaction_id"].(string))
	}
	if modelMap["updated"] != nil {

	}
	if modelMap["code"] != nil && modelMap["code"].(string) != "" {
		model.Code = core.StringPtr(modelMap["code"].(string))
	}
	return model, nil
}

func resourceIBMCmOfferingInstanceOfferingInstanceLastOperationToMap(model *catalogmanagementv1.OfferingInstanceLastOperation) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Operation != nil {
		modelMap["operation"] = model.Operation
	}
	if model.State != nil {
		modelMap["state"] = model.State
	}
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.TransactionID != nil {
		modelMap["transaction_id"] = model.TransactionID
	}
	if model.Updated != nil {
		modelMap["updated"] = model.Updated.String()
	}
	if model.Code != nil {
		modelMap["code"] = model.Code
	}
	return modelMap, nil
}
