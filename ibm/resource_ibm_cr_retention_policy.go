// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
)

func resourceIBMCrRetentionPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCrRetentionPolicyCreate,
		ReadContext:   resourceIBMCrRetentionPolicyRead,
		UpdateContext: resourceIBMCrRetentionPolicyUpdate,
		DeleteContext: resourceIBMCrRetentionPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"namespace": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The namespace to which the retention policy is attached.",
			},
			"images_per_repo": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Determines how many images will be retained for each repository when the retention policy is executed. The value -1 denotes 'Unlimited' (all images are retained).",
			},
			"retain_untagged": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Determines if untagged images are retained when executing the retention policy. This is false by default meaning untagged images will be deleted when the policy is executed.",
			},
		},
	}
}

func resourceIBMCrRetentionPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	containerRegistryClient, err := meta.(ClientSession).ContainerRegistryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	setRetentionPolicyOptions := &containerregistryv1.SetRetentionPolicyOptions{}

	setRetentionPolicyOptions.SetNamespace(d.Get("namespace").(string))
	setRetentionPolicyOptions.SetImagesPerRepo(int64(d.Get("images_per_repo").(int)))
	if _, ok := d.GetOk("retain_untagged"); ok {
		setRetentionPolicyOptions.SetRetainUntagged(d.Get("retain_untagged").(bool))
	}

	response, err := containerRegistryClient.SetRetentionPolicyWithContext(context, setRetentionPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] SetRetentionPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(d.Get("namespace").(string))

	return resourceIBMCrRetentionPolicyRead(context, d, meta)
}

func resourceIBMCrRetentionPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	containerRegistryClient, err := meta.(ClientSession).ContainerRegistryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRetentionPolicyOptions := &containerregistryv1.GetRetentionPolicyOptions{}

	getRetentionPolicyOptions.SetNamespace(d.Id())

	retentionPolicy, response, err := containerRegistryClient.GetRetentionPolicyWithContext(context, getRetentionPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetRetentionPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	// A retention policy "does not exist" if `imagesPerRepo` is -1 `retainUntagged` is true
	if *retentionPolicy.ImagesPerRepo == -1 && *retentionPolicy.RetainUntagged == true {
		d.SetId("")
		return nil
	}

	if err = d.Set("namespace", retentionPolicy.Namespace); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting namespace: %s", err))
	}
	if err = d.Set("images_per_repo", intValue(retentionPolicy.ImagesPerRepo)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting images_per_repo: %s", err))
	}
	if err = d.Set("retain_untagged", retentionPolicy.RetainUntagged); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting retain_untagged: %s", err))
	}

	return nil
}

func resourceIBMCrRetentionPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	containerRegistryClient, err := meta.(ClientSession).ContainerRegistryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	setRetentionPolicyOptions := &containerregistryv1.SetRetentionPolicyOptions{}

	setRetentionPolicyOptions.SetNamespace(d.Id())

	hasChange := false

	if d.HasChange("namespace") {
		setRetentionPolicyOptions.SetNamespace(d.Get("namespace").(string))
		hasChange = true
	}
	if d.HasChange("images_per_repo") {
		setRetentionPolicyOptions.SetImagesPerRepo(int64(d.Get("images_per_repo").(int)))
		hasChange = true
	}
	if d.HasChange("retain_untagged") {
		setRetentionPolicyOptions.SetRetainUntagged(d.Get("retain_untagged").(bool))
		hasChange = true
	}

	if hasChange {
		response, err := containerRegistryClient.SetRetentionPolicyWithContext(context, setRetentionPolicyOptions)
		if err != nil {
			log.Printf("[DEBUG] SetRetentionPolicyWithContext failed %s\n%s", err, response)
			return diag.FromErr(err)
		}
	}

	return resourceIBMCrRetentionPolicyRead(context, d, meta)
}

func resourceIBMCrRetentionPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	containerRegistryClient, err := meta.(ClientSession).ContainerRegistryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	setRetentionPolicyOptions := &containerregistryv1.SetRetentionPolicyOptions{}

	setRetentionPolicyOptions.SetNamespace(d.Id())
	setRetentionPolicyOptions.SetImagesPerRepo(-1)
	setRetentionPolicyOptions.SetRetainUntagged(true)

	response, err := containerRegistryClient.SetRetentionPolicyWithContext(context, setRetentionPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] SetRetentionPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
