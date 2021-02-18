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

func dataSourceIbmCmVersionInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmCmVersionInstanceRead,

		Schema: map[string]*schema.Schema{
			"instance_identifier": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Version Instance identifier.",
			},
			"rid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "provisioned instance ID (part of the CRN).",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "url reference to this object.",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "platform CRN for this instance.",
			},
			"label": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the label for this instance.",
			},
			"catalog_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Catalog ID this instance was created from.",
			},
			"offering_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Offering ID this instance was created from.",
			},
			"kind_format": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the format this instance has (helm, operator, ova...).",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version this instance was installed from (not version id).",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster ID.",
			},
			"cluster_region": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster region (e.g., us-south).",
			},
			"cluster_namespaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of target namespaces to install into.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cluster_all_namespaces": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "designate to install into all namespaces.",
			},
		},
	}
}

func dataSourceIbmCmVersionInstanceRead(d *schema.ResourceData, meta interface{}) error {
	catalogManagementClient, err := meta.(ClientSession).CatalogManagementV1()
	if err != nil {
		return err
	}

	getVersionInstanceOptions := &catalogmanagementv1.GetVersionInstanceOptions{}

	getVersionInstanceOptions.SetInstanceIdentifier(d.Get("instance_identifier").(string))

	versionInstance, response, err := catalogManagementClient.GetVersionInstance(getVersionInstanceOptions)
	if err != nil {
		log.Printf("[DEBUG] GetVersionInstance failed %s\n%s", err, response)
		return err
	}

	d.SetId(*versionInstance.ID)
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
	if err = d.Set("cluster_namespaces", versionInstance.ClusterNamespaces); err != nil {
		return fmt.Errorf("Error setting cluster_namespaces: %s", err)
	}
	if err = d.Set("cluster_all_namespaces", versionInstance.ClusterAllNamespaces); err != nil {
		return fmt.Errorf("Error setting cluster_all_namespaces: %s", err)
	}

	return nil
}
