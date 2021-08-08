// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

func resourceIbmCloudantDatabase() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmCloudantDatabaseCreate,
		ReadContext:   resourceIbmCloudantDatabaseRead,
		DeleteContext: resourceIbmCloudantDatabaseDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cloudant_guid": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cloudant Instance ID.",
			},
			"db": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Path parameter to specify the database name.",
			},
			"partitioned": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Query parameter to specify whether to enable database partitions when creating a database.",
			},
			"q": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config`.",
			},
		},
	}
}

func resourceIbmCloudantDatabaseCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceId := d.Get("cloudant_guid").(string)
	cUrl, err := getCloudantInstanceUrl(instanceId, meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	dbName := d.Get("db").(string)
	putDatabaseOptions := &cloudantv1.PutDatabaseOptions{}
	putDatabaseOptions.SetDb(dbName)
	if _, ok := d.GetOk("partitioned"); ok {
		putDatabaseOptions.SetPartitioned(d.Get("partitioned").(bool))
	}
	if _, ok := d.GetOk("q"); ok {
		putDatabaseOptions.SetQ(int64(d.Get("q").(int)))
	}

	_, response, err := cloudantClient.PutDatabaseWithContext(context, putDatabaseOptions)
	if err != nil {
		log.Printf("[DEBUG] PutDatabaseWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PutDatabaseWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", instanceId, dbName))

	return resourceIbmCloudantDatabaseRead(context, d, meta)
}

func resourceIbmCloudantDatabaseRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cUrl, err := getCloudantInstanceUrl(parts[0], meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	getDatabaseInformationOptions := &cloudantv1.GetDatabaseInformationOptions{}
	getDatabaseInformationOptions.SetDb(parts[1])

	databaseInformation, response, err := cloudantClient.GetDatabaseInformationWithContext(context, getDatabaseInformationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetDatabaseInformationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetDatabaseInformationWithContext failed %s\n%s", err, response))
	}

	d.Set("cloudant_guid", parts[0])

	if err = d.Set("db", *databaseInformation.DbName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting db: %s", err))
	}

	if _, ok := d.GetOk("partitioned"); ok {
		d.Set("partitioned", true)
	}

	if qData, ok := d.GetOk("q"); ok {
		d.Set("q", qData.(int))
	}

	return nil
}

func resourceIbmCloudantDatabaseDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudantClient, err := meta.(ClientSession).CloudantV1()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cUrl, err := getCloudantInstanceUrl(parts[0], meta)
	if err != nil {
		return diag.FromErr(err)
	}
	cloudantClient.Service.Options.URL = cUrl

	deleteDatabaseOptions := &cloudantv1.DeleteDatabaseOptions{}
	deleteDatabaseOptions.SetDb(parts[1])

	_, response, err := cloudantClient.DeleteDatabaseWithContext(context, deleteDatabaseOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteDatabaseWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteDatabaseWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func getCloudantInstanceUrl(instanceId string, meta interface{}) (string, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerV2API()
	if err != nil {
		return "", err
	}

	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: ptrToString(instanceId),
	}

	instance, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
	if err != nil {
		return "", fmt.Errorf("Error retrieving resource instance: %s with resp code: %s", err, resp)
	}

	if instance.Extensions != nil {
		instanceExtensionMap := Flatten(instance.Extensions)
		if instanceExtensionMap != nil {
			cloudantInstanceUrl := "https://" + instanceExtensionMap["endpoints.public"]
			cloudantInstanceUrl = envFallBack([]string{"IBMCLOUD_CLOUDANT_API_ENDPOINT"}, cloudantInstanceUrl)
			return cloudantInstanceUrl, nil
		}
	}

	return "", fmt.Errorf("Unable to get URL for cloudant instance")
}
