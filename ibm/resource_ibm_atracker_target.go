// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/atrackerv1"
)

func resourceIBMAtrackerTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAtrackerTargetCreate,
		ReadContext:   resourceIBMAtrackerTargetRead,
		DeleteContext: resourceIBMAtrackerTargetDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the target resource.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the target resource.",
			},
			"target_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the target.",
			},
			"encrypt_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption key that is used to encrypt events before Activity Tracking services buffer them on storage. This credential is masked in the response.",
			},
			"cos_endpoint": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Property values for a Cloud Object Storage Endpoint.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The host name of the Cloud Object Storage endpoint.",
						},
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of the Cloud Object Storage instance.",
						},
						"bucket": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The bucket name under the Cloud Object Storage instance.",
						},
						"api_key": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							Description: "The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response.",
						},
					},
				},
			},
			"cos_write_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status of the write attempt with the provided cos_endpoint parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The status such as failed or success.",
						},
						"last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The timestamp of the failure.",
						},
						"reason_for_last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Detailed description of the cause of the failure.",
						},
					},
				},
			},
			"created": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the target creation time.",
			},
			"updated": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the target last updated time.",
			},
		},
	}
}

func resourceIBMAtrackerTargetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	validateTargetOptions := &atrackerv1.ValidateTargetOptions{}

	target, response, err := atrackerClient.ValidateTargetWithContext(context, validateTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] ValidateTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ValidateTargetWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *validateTargetOptions.ID, *target.ID))

	return resourceIBMAtrackerTargetRead(context, d, meta)
}

func resourceIBMAtrackerTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getTargetOptions := &atrackerv1.GetTargetOptions{}

	parts, err := sepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getTargetOptions.SetID(parts[0])
	getTargetOptions.SetID(parts[1])

	target, response, err := atrackerClient.GetTargetWithContext(context, getTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetTargetWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", target.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("crn", target.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("target_type", target.TargetType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting target_type: %s", err))
	}
	if err = d.Set("encrypt_key", target.EncryptKey); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting encrypt_key: %s", err))
	}
	if target.CosEndpoint != nil {
		cosEndpointMap := resourceIBMAtrackerTargetCosEndpointToMap(*target.CosEndpoint)
		if err = d.Set("cos_endpoint", []map[string]interface{}{cosEndpointMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cos_endpoint: %s", err))
		}
	}
	if target.CosWriteStatus != nil {
		cosWriteStatusMap := resourceIBMAtrackerTargetCosWriteStatusToMap(*target.CosWriteStatus)
		if err = d.Set("cos_write_status", []map[string]interface{}{cosWriteStatusMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cos_write_status: %s", err))
		}
	}
	if err = d.Set("created", dateTimeToString(target.Created)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created: %s", err))
	}
	if err = d.Set("updated", dateTimeToString(target.Updated)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated: %s", err))
	}

	return nil
}

func resourceIBMAtrackerTargetCosEndpointToMap(cosEndpoint atrackerv1.CosEndpoint) map[string]interface{} {
	cosEndpointMap := map[string]interface{}{}

	cosEndpointMap["endpoint"] = cosEndpoint.Endpoint
	cosEndpointMap["target_crn"] = cosEndpoint.TargetCRN
	cosEndpointMap["bucket"] = cosEndpoint.Bucket
	cosEndpointMap["api_key"] = cosEndpoint.APIKey

	return cosEndpointMap
}

func resourceIBMAtrackerTargetCosWriteStatusToMap(cosWriteStatus atrackerv1.CosWriteStatus) map[string]interface{} {
	cosWriteStatusMap := map[string]interface{}{}

	if cosWriteStatus.Status != nil {
		cosWriteStatusMap["status"] = cosWriteStatus.Status
	}
	if cosWriteStatus.LastFailure != nil {
		cosWriteStatusMap["last_failure"] = cosWriteStatus.LastFailure.String()
	}
	if cosWriteStatus.ReasonForLastFailure != nil {
		cosWriteStatusMap["reason_for_last_failure"] = cosWriteStatus.ReasonForLastFailure
	}

	return cosWriteStatusMap
}

func resourceIBMAtrackerTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTargetOptions := &atrackerv1.DeleteTargetOptions{}

	parts, err := sepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteTargetOptions.SetID(parts[0])
	deleteTargetOptions.SetID(parts[1])

	_, response, err := atrackerClient.DeleteTargetWithContext(context, deleteTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTargetWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteTargetWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
