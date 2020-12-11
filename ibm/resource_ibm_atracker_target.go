/**
 * (C) Copyright IBM Corp. 2020.
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
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func resourceIBMAtrackerTarget() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMAtrackerTargetCreate,
		Read:     resourceIBMAtrackerTargetRead,
		Update:   resourceIBMAtrackerTargetUpdate,
		Delete:   resourceIBMAtrackerTargetDelete,
		Exists:   resourceIBMAtrackerTargetExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the target. Must be 256 characters or less.",
			},
			"target_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the target.",
			},
			"cos_endpoint": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Property values for a Cloud Object Storage Endpoint.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The host name of this COS endpoint.",
						},
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of this COS instance.",
						},
						"bucket": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The bucket name under this COS instance.",
						},
						"api_key": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The IAM Api key that have writer access to this cos instance. This credential will be masked in the response.",
						},
					},
				},
			},
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The uuid of ATracker services in this region.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of this target type resource.",
			},
			"encrypt_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption key used to encrypt events before ATracker services buffer them on storage. This credential will be masked in the response.",
			},
		},
	}
}

func resourceIBMAtrackerTargetCreate(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	createTargetOptions := &atrackerv1.CreateTargetOptions{}

	createTargetOptions.SetName(d.Get("name").(string))
	createTargetOptions.SetTargetType(d.Get("target_type").(string))
	cosEndpoint := resourceIBMAtrackerTargetMapToCosEndpoint(d.Get("cos_endpoint.0").(map[string]interface{}))
	createTargetOptions.SetCosEndpoint(&cosEndpoint)

	target, response, err := atrackerClient.CreateTarget(createTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateTarget failed %s\n%s", err, response)
		return err
	}

	d.SetId(*target.ID)

	return resourceIBMAtrackerTargetRead(d, meta)
}

func resourceIBMAtrackerTargetMapToCosEndpoint(cosEndpointMap map[string]interface{}) atrackerv1.CosEndpoint {
	cosEndpoint := atrackerv1.CosEndpoint{}

	cosEndpoint.Endpoint = core.StringPtr(cosEndpointMap["endpoint"].(string))
	cosEndpoint.TargetCRN = core.StringPtr(cosEndpointMap["target_crn"].(string))
	cosEndpoint.Bucket = core.StringPtr(cosEndpointMap["bucket"].(string))
	cosEndpoint.APIKey = core.StringPtr(cosEndpointMap["api_key"].(string))

	return cosEndpoint
}

func resourceIBMAtrackerTargetRead(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	getTargetOptions := &atrackerv1.GetTargetOptions{}

	getTargetOptions.SetID(d.Id())

	target, response, err := atrackerClient.GetTarget(getTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] GetTarget failed %s\n%s", err, response)
		return err
	}

	d.Set("name", target.Name)
	d.Set("target_type", target.TargetType)
	cosEndpointMap := resourceIBMAtrackerTargetCosEndpointToMap(*target.CosEndpoint)
	// This line is a workaround for api_key, which comes back as "REDACTED" from the service.
	// This causes havok in the tests (in legacy testing framework) so we store the original
	// api_key value into the state.  This is the least bad solution I could come up with.
	cosEndpointMap["api_key"] = d.Get("cos_endpoint.0.api_key")
	d.Set("cos_endpoint", []map[string]interface{}{cosEndpointMap})
	d.Set("instance_id", target.InstanceID)
	d.Set("crn", target.CRN)
	d.Set("encrypt_key", target.EncryptKey)

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

func resourceIBMAtrackerTargetUpdate(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	replaceTargetOptions := &atrackerv1.ReplaceTargetOptions{}

	replaceTargetOptions.SetID(d.Id())
	replaceTargetOptions.SetName(d.Get("name").(string))
	replaceTargetOptions.SetTargetType(d.Get("target_type").(string))
	cosEndpoint := resourceIBMAtrackerTargetMapToCosEndpoint(d.Get("cos_endpoint.0").(map[string]interface{}))
	replaceTargetOptions.SetCosEndpoint(&cosEndpoint)

	_, response, err := atrackerClient.ReplaceTarget(replaceTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceTarget failed %s\n%s", err, response)
		return err
	}

	return resourceIBMAtrackerTargetRead(d, meta)
}

func resourceIBMAtrackerTargetDelete(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	deleteTargetOptions := &atrackerv1.DeleteTargetOptions{}

	deleteTargetOptions.SetID(d.Id())

	response, err := atrackerClient.DeleteTarget(deleteTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteTarget failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}

func resourceIBMAtrackerTargetExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return false, err
	}

	getTargetOptions := &atrackerv1.GetTargetOptions{}

	getTargetOptions.SetID(d.Id())

	target, response, err := atrackerClient.GetTarget(getTargetOptions)
	if err != nil {
		log.Printf("[DEBUG] GetTarget failed %s\n%s", err, response)
		return false, err
	}

	return *target.ID == d.Id(), nil
}
