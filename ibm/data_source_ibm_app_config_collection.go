// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func dataSourceIbmAppConfigCollection() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigCollectionRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"collection_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Collection Id of the collection.",
			},
			"expand": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to `true`, returns expanded view of the resource details.",
			},
			"includes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Include feature and property details in the response.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Collection name.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Collection description.",
			},
			"tags": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tags associated with the collection.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the collection.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last updated time of the collection data.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Collection URL.",
			},
			"features": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Features associated with the collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"feature_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Feature id.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Feature name.",
						},
					},
				},
			},
			"properties": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of properties associated with the collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property id.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Property name.",
						},
					},
				},
			},
			"features_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of features associated with the collection.",
			},
			"properties_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of features associated with the collection.",
			},
		},
	}
}

func dataSourceIbmAppConfigCollectionRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.GetCollectionOptions{}

	options.SetCollectionID(d.Get("collection_id").(string))
	if _, ok := d.GetOk("expand"); ok {
		options.SetExpand(d.Get("expand").(bool))
	}
	if _, ok := d.GetOk("includes"); ok {
		includes := []string{}
		for _, item := range d.Get("includes").([]interface{}) {
			includes = append(includes, item.(string))
		}
		options.SetInclude(includes)
	}
	result, response, err := appconfigClient.GetCollection(options)
	if err != nil {
		log.Printf("[DEBUG] GetCollection failed %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", guid, *result.CollectionID))
	if result.Name != nil {
		if err = d.Set("name", result.Name); err != nil {
			return fmt.Errorf("error setting name: %s", err)
		}
	}
	if result.Description != nil {
		if err = d.Set("description", result.Description); err != nil {
			return fmt.Errorf("error setting description: %s", err)
		}
	}
	if result.Tags != nil {
		if err = d.Set("tags", result.Tags); err != nil {
			return fmt.Errorf("error setting tags: %s", err)
		}

	}
	if result.CreatedTime != nil {
		if err = d.Set("created_time", result.CreatedTime.String()); err != nil {
			return fmt.Errorf("error setting created_time: %s", err)
		}
	}

	if result.UpdatedTime != nil {
		if err = d.Set("updated_time", result.UpdatedTime.String()); err != nil {
			return fmt.Errorf("error setting updated_time: %s", err)
		}
	}

	if result.Href != nil {
		if err = d.Set("href", result.Href); err != nil {
			return fmt.Errorf("error setting href: %s", err)
		}
	}
	if result.Features != nil {
		err = d.Set("features", dataSourceCollectionFlattenFeatures(result.Features))
		if err != nil {
			return fmt.Errorf("error setting features %s", err)
		}
	}

	if result.Properties != nil {
		err = d.Set("properties", dataSourceCollectionFlattenProperties(result.Properties))
		if err != nil {
			return fmt.Errorf("error setting properties %s", err)
		}
	}
	if err = d.Set("features_count", result.FeaturesCount); err != nil {
		return fmt.Errorf("error setting features_count: %s", err)
	}
	if err = d.Set("properties_count", result.PropertiesCount); err != nil {
		return fmt.Errorf("error setting properties_count: %s", err)
	}

	return nil
}

func dataSourceCollectionFlattenFeatures(result []appconfigurationv1.FeatureOutput) (features []map[string]interface{}) {
	for _, featuresItem := range result {
		features = append(features, dataSourceCollectionFeaturesToMap(featuresItem))
	}

	return features
}

func dataSourceCollectionFeaturesToMap(featuresItem appconfigurationv1.FeatureOutput) (featuresMap map[string]interface{}) {
	featuresMap = map[string]interface{}{}

	if featuresItem.FeatureID != nil {
		featuresMap["feature_id"] = featuresItem.FeatureID
	}
	if featuresItem.Name != nil {
		featuresMap["name"] = featuresItem.Name
	}

	return featuresMap
}

func dataSourceCollectionFlattenProperties(result []appconfigurationv1.PropertyOutput) (properties []map[string]interface{}) {
	for _, propertiesItem := range result {
		properties = append(properties, dataSourceCollectionPropertiesToMap(propertiesItem))
	}

	return properties
}

func dataSourceCollectionPropertiesToMap(propertiesItem appconfigurationv1.PropertyOutput) (propertiesMap map[string]interface{}) {
	propertiesMap = map[string]interface{}{}

	if propertiesItem.PropertyID != nil {
		propertiesMap["property_id"] = propertiesItem.PropertyID
	}
	if propertiesItem.Name != nil {
		propertiesMap["name"] = propertiesItem.Name
	}

	return propertiesMap
}
