/**
 * (C) Copyright IBM Corp. 2021.
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
	"fmt"
	"log"

	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIbmCmVersionInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmCmVersionInstanceCreate,
		Read:     resourceIbmCmVersionInstanceRead,
		Update:   resourceIbmCmVersionInstanceUpdate,
		Delete:   resourceIbmCmVersionInstanceDelete,
		Exists:   resourceIbmCmVersionInstanceExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"x_auth_refresh_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IAM Refresh token.",
			},
			"rid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "provisioned instance ID (part of the CRN).",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "url reference to this object.",
			},
			"crn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "platform CRN for this instance.",
			},
			"label": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the label for this instance.",
			},
			"catalog_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Catalog ID this instance was created from.",
			},
			"offering_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Offering ID this instance was created from.",
			},
			"kind_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "the format this instance has (helm, operator, ova...).",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version this instance was installed from (not version id).",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster ID.",
			},
			"cluster_region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster region (e.g., us-south).",
			},
			"cluster_namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of target namespaces to install into.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cluster_all_namespaces": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "designate to install into all namespaces.",
			},
		},
	}
}

func resourceIbmCmVersionInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	createVersionInstanceOptions := &catalogmanagementv1.CreateVersionInstanceOptions{}

	createVersionInstanceOptions.SetXAuthRefreshToken(d.Get("x_auth_refresh_token").(string))
	if _, ok := d.GetOk("rid"); ok {
		createVersionInstanceOptions.SetID(d.Get("rid").(string))
	}
	if _, ok := d.GetOk("url"); ok {
		createVersionInstanceOptions.SetURL(d.Get("url").(string))
	}
	if _, ok := d.GetOk("crn"); ok {
		createVersionInstanceOptions.SetCRN(d.Get("crn").(string))
	}
	if _, ok := d.GetOk("label"); ok {
		createVersionInstanceOptions.SetLabel(d.Get("label").(string))
	}
	if _, ok := d.GetOk("catalog_id"); ok {
		createVersionInstanceOptions.SetCatalogID(d.Get("catalog_id").(string))
	}
	if _, ok := d.GetOk("offering_id"); ok {
		createVersionInstanceOptions.SetOfferingID(d.Get("offering_id").(string))
	}
	if _, ok := d.GetOk("kind_format"); ok {
		createVersionInstanceOptions.SetKindFormat(d.Get("kind_format").(string))
	}
	if _, ok := d.GetOk("version"); ok {
		createVersionInstanceOptions.SetVersion(d.Get("version").(string))
	}
	if _, ok := d.GetOk("cluster_id"); ok {
		createVersionInstanceOptions.SetClusterID(d.Get("cluster_id").(string))
	}
	if _, ok := d.GetOk("cluster_region"); ok {
		createVersionInstanceOptions.SetClusterRegion(d.Get("cluster_region").(string))
	}
	if ns, ok := d.GetOk("cluster_namespaces"); ok {
		list := expandStringList(ns.([]interface{}))
		createVersionInstanceOptions.SetClusterNamespaces(list)
	}
	if _, ok := d.GetOk("cluster_all_namespaces"); ok {
		createVersionInstanceOptions.SetClusterAllNamespaces(d.Get("cluster_all_namespaces").(bool))
	}

	versionInstance, response, err := catalogManagementClient.CreateVersionInstance(createVersionInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateVersionInstance failed %s\n%s", err, response)
		return err
	}

	d.SetId(*versionInstance.ID)

	// Env for CM URL
	// Standard out vs debug logs?
	log.Printf("LOG2 Service version instance of type %q was created on cluster %q", *createVersionInstanceOptions.KindFormat, *createVersionInstanceOptions.ClusterID)

	return resourceIbmCmVersionInstanceRead(d, meta)
}

func resourceIbmCmVersionInstanceRead(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	getVersionInstanceOptions := &catalogmanagementv1.GetVersionInstanceOptions{}

	getVersionInstanceOptions.SetInstanceIdentifier(d.Id())

	versionInstance, response, err := catalogManagementClient.GetVersionInstance(getVersionInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetVersionInstance failed %s\n%s", err, response)
		return err
	}
	if err = d.Set("rid", versionInstance.ID); err != nil {
		return fmt.Errorf("Error setting id: %s", err)
	}
	if err = d.Set("url", versionInstance.URL); err != nil {
		return fmt.Errorf("Error setting url: %s", err)
	}
	if err = d.Set("crn", versionInstance.CRN); err != nil {
		return fmt.Errorf("Error setting crn: %s", err)
	}
	if err = d.Set("label", versionInstance.Label); err != nil {
		return fmt.Errorf("Error setting label: %s", err)
	}
	if err = d.Set("catalog_id", versionInstance.CatalogID); err != nil {
		return fmt.Errorf("Error setting catalog_id: %s", err)
	}
	if err = d.Set("offering_id", versionInstance.OfferingID); err != nil {
		return fmt.Errorf("Error setting offering_id: %s", err)
	}
	if err = d.Set("kind_format", versionInstance.KindFormat); err != nil {
		return fmt.Errorf("Error setting kind_format: %s", err)
	}
	if err = d.Set("version", versionInstance.Version); err != nil {
		return fmt.Errorf("Error setting version: %s", err)
	}
	if err = d.Set("cluster_id", versionInstance.ClusterID); err != nil {
		return fmt.Errorf("Error setting cluster_id: %s", err)
	}
	if err = d.Set("cluster_region", versionInstance.ClusterRegion); err != nil {
		return fmt.Errorf("Error setting cluster_region: %s", err)
	}
	if versionInstance.ClusterNamespaces != nil {
		if err = d.Set("cluster_namespaces", versionInstance.ClusterNamespaces); err != nil {
			return fmt.Errorf("Error setting cluster_namespaces: %s", err)
		}
	}
	if err = d.Set("cluster_all_namespaces", versionInstance.ClusterAllNamespaces); err != nil {
		return fmt.Errorf("Error setting cluster_all_namespaces: %s", err)
	}

	return nil
}

func resourceIbmCmVersionInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	putVersionInstanceOptions := &catalogmanagementv1.PutVersionInstanceOptions{}

	putVersionInstanceOptions.SetInstanceIdentifier(d.Id())
	putVersionInstanceOptions.SetXAuthRefreshToken(d.Get("x_auth_refresh_token").(string))
	if _, ok := d.GetOk("rid"); ok {
		putVersionInstanceOptions.SetID(d.Get("rid").(string))
	}
	if _, ok := d.GetOk("url"); ok {
		putVersionInstanceOptions.SetURL(d.Get("url").(string))
	}
	if _, ok := d.GetOk("crn"); ok {
		putVersionInstanceOptions.SetCRN(d.Get("crn").(string))
	}
	if _, ok := d.GetOk("label"); ok {
		putVersionInstanceOptions.SetLabel(d.Get("label").(string))
	}
	if _, ok := d.GetOk("catalog_id"); ok {
		putVersionInstanceOptions.SetCatalogID(d.Get("catalog_id").(string))
	}
	if _, ok := d.GetOk("offering_id"); ok {
		putVersionInstanceOptions.SetOfferingID(d.Get("offering_id").(string))
	}
	if _, ok := d.GetOk("kind_format"); ok {
		putVersionInstanceOptions.SetKindFormat(d.Get("kind_format").(string))
	}
	if _, ok := d.GetOk("version"); ok {
		putVersionInstanceOptions.SetVersion(d.Get("version").(string))
	}
	if _, ok := d.GetOk("cluster_id"); ok {
		putVersionInstanceOptions.SetClusterID(d.Get("cluster_id").(string))
	}
	if _, ok := d.GetOk("cluster_region"); ok {
		putVersionInstanceOptions.SetClusterRegion(d.Get("cluster_region").(string))
	}
	if ns, ok := d.GetOk("cluster_namespaces"); ok {
		list := expandStringList(ns.([]interface{}))
		putVersionInstanceOptions.SetClusterNamespaces(list)
	}
	if _, ok := d.GetOk("cluster_all_namespaces"); ok {
		putVersionInstanceOptions.SetClusterAllNamespaces(d.Get("cluster_all_namespaces").(bool))
	}

	_, response, err := catalogManagementClient.PutVersionInstance(putVersionInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] PutVersionInstance failed %s\n%s", err, response)
		return err
	}

	return resourceIbmCmVersionInstanceRead(d, meta)
}

func resourceIbmCmVersionInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	deleteVersionInstanceOptions := &catalogmanagementv1.DeleteVersionInstanceOptions{}

	deleteVersionInstanceOptions.SetInstanceIdentifier(d.Id())

	response, err := catalogManagementClient.DeleteVersionInstance(deleteVersionInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteVersionInstance failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}

func resourceIbmCmVersionInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return false, err
	}

	getVersionInstanceOptions := &catalogmanagementv1.GetVersionInstanceOptions{}

	getVersionInstanceOptions.SetInstanceIdentifier(d.Id())

	versionInstance, response, err := catalogManagementClient.GetVersionInstance(getVersionInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetVersionInstance failed %s\n%s", err, response)
		return false, err
	}

	return *versionInstance.ID == d.Id(), nil
}
