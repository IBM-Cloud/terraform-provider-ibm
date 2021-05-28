// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func dataSourceIbmAppConfigCollections() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigCollectionsRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"sort": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sort the collection details based on the specified attribute.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.",
			},
			"tags": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "filter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.",
			},
			"includes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Include feature and property details in the response.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"expand": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to `true`, returns expanded view of the resource details.",
			},
			"features": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter collections by a list of comma separated features.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"properties": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter collections by a list of comma separated properties.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"collections": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of collections.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection name.",
						},
						"collection_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection Id.",
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
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of records.",
			},
			"next": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the next list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"first": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the first page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"previous": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the previous list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"last": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the last page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAppConfigCollectionsRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.ListCollectionsOptions{}

	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}
	if _, ok := d.GetOk("sort"); ok {
		options.SetTags(d.Get("sort").(string))
	}
	if _, ok := d.GetOk("includes"); ok {
		includes := []string{}
		for _, item := range d.Get("includes").([]interface{}) {
			includes = append(includes, item.(string))
		}
		options.SetInclude(includes)
	}
	if _, ok := d.GetOk("features"); ok {
		features := []string{}
		for _, item := range d.Get("features").([]interface{}) {
			features = append(features, item.(string))
		}
		options.SetFeatures(features)
	}
	if _, ok := d.GetOk("properties"); ok {
		properties := []string{}
		for _, item := range d.Get("properties").([]interface{}) {
			properties = append(properties, item.(string))
		}
		options.SetProperties(properties)
	}
	if _, ok := d.GetOk("expand"); ok {
		options.SetExpand(d.Get("expand").(bool))
	}
	var collectionList *appconfigurationv1.CollectionList
	var offset int64 = 0
	var limit int64 = 10
	finalList := []appconfigurationv1.Collection{}
	var isLimit bool
	if _, ok := d.GetOk("limit"); ok {
		isLimit = true
		limit = int64(d.Get("limit").(int))
	}
	options.SetLimit(limit)

	if _, ok := d.GetOk("offset"); ok {
		offset = int64(d.Get("offset").(int))
	}
	for {
		options.SetOffset(offset)

		result, response, err := appconfigClient.ListCollections(options)
		collectionList = result
		if err != nil {
			log.Printf("[DEBUG] ListCollections failed %s\n%s", err, response)
			return err
		}

		if isLimit {
			offset = 0
		} else {
			offset = dataSourceCollectionListGetNext(result.Next)
		}
		finalList = append(finalList, result.Collections...)
		if offset == 0 {
			break
		}
	}

	collectionList.Collections = finalList

	d.SetId(guid)

	if collectionList.Collections != nil {
		err = d.Set("collections", dataSourceCollectionListFlattenCollections(collectionList.Collections))
		if err != nil {
			return fmt.Errorf("error setting collections %s", err)
		}
	}
	if collectionList.TotalCount != nil {
		if err = d.Set("total_count", collectionList.TotalCount); err != nil {
			return fmt.Errorf("error setting total_count: %s", err)
		}
	}
	if collectionList.Limit != nil {
		if err = d.Set("limit", collectionList.Limit); err != nil {
			return fmt.Errorf("error setting limit: %s", err)
		}
	}
	if collectionList.Offset != nil {
		if err = d.Set("offset", collectionList.Offset); err != nil {
			return fmt.Errorf("error setting offset: %s", err)
		}
	}
	if collectionList.First != nil {
		err = d.Set("first", dataSourceEnvironmentListFlattenPagination(*collectionList.First))
		if err != nil {
			return fmt.Errorf("error setting first %s", err)
		}
	}

	if collectionList.Previous != nil {
		err = d.Set("previous", dataSourceEnvironmentListFlattenPagination(*collectionList.Previous))
		if err != nil {
			return fmt.Errorf("error setting previous %s", err)
		}
	}

	if collectionList.Last != nil {
		err = d.Set("last", dataSourceEnvironmentListFlattenPagination(*collectionList.Last))
		if err != nil {
			return fmt.Errorf("error setting last %s", err)
		}
	}
	if collectionList.Next != nil {
		err = d.Set("next", dataSourceEnvironmentListFlattenPagination(*collectionList.Next))
		if err != nil {
			return fmt.Errorf("error setting next %s", err)
		}
	}
	return nil
}

func dataSourceCollectionListFlattenCollections(result []appconfigurationv1.Collection) (collections []map[string]interface{}) {
	for _, collectionsItem := range result {
		collections = append(collections, dataSourceCollectionListCollectionsToMap(collectionsItem))
	}

	return collections
}

func dataSourceCollectionListCollectionsToMap(collectionsItem appconfigurationv1.Collection) (collectionsMap map[string]interface{}) {
	collectionsMap = map[string]interface{}{}

	if collectionsItem.Name != nil {
		collectionsMap["name"] = collectionsItem.Name
	}
	if collectionsItem.CollectionID != nil {
		collectionsMap["collection_id"] = collectionsItem.CollectionID
	}
	if collectionsItem.Description != nil {
		collectionsMap["description"] = collectionsItem.Description
	}
	if collectionsItem.Tags != nil {
		collectionsMap["tags"] = collectionsItem.Tags
	}
	if collectionsItem.CreatedTime != nil {
		collectionsMap["created_time"] = collectionsItem.CreatedTime.String()
	}
	if collectionsItem.UpdatedTime != nil {
		collectionsMap["updated_time"] = collectionsItem.UpdatedTime.String()
	}
	if collectionsItem.Href != nil {
		collectionsMap["href"] = collectionsItem.Href
	}
	if collectionsItem.Features != nil {
		featuresList := []map[string]interface{}{}
		for _, featuresItem := range collectionsItem.Features {
			featuresList = append(featuresList, dataSourceCollectionListCollectionsFeaturesToMap(featuresItem))
		}
		collectionsMap["features"] = featuresList
	}
	if collectionsItem.Properties != nil {
		propertiesList := []map[string]interface{}{}
		for _, propertiesItem := range collectionsItem.Properties {
			propertiesList = append(propertiesList, dataSourceCollectionListCollectionsPropertiesToMap(propertiesItem))
		}
		collectionsMap["properties"] = propertiesList
	}
	if collectionsItem.FeaturesCount != nil {
		collectionsMap["features_count"] = collectionsItem.FeaturesCount
	}
	if collectionsItem.PropertiesCount != nil {
		collectionsMap["properties_count"] = collectionsItem.PropertiesCount
	}

	return collectionsMap
}

func dataSourceCollectionListCollectionsFeaturesToMap(featuresItem appconfigurationv1.FeatureOutput) (featuresMap map[string]interface{}) {
	featuresMap = map[string]interface{}{}

	if featuresItem.FeatureID != nil {
		featuresMap["feature_id"] = featuresItem.FeatureID
	}
	if featuresItem.Name != nil {
		featuresMap["name"] = featuresItem.Name
	}

	return featuresMap
}

func dataSourceCollectionListCollectionsPropertiesToMap(propertiesItem appconfigurationv1.PropertyOutput) (propertiesMap map[string]interface{}) {
	propertiesMap = map[string]interface{}{}

	if propertiesItem.PropertyID != nil {
		propertiesMap["property_id"] = propertiesItem.PropertyID
	}
	if propertiesItem.Name != nil {
		propertiesMap["name"] = propertiesItem.Name
	}

	return propertiesMap
}

func dataSourceCollectionListGetNext(next interface{}) int64 {
	if reflect.ValueOf(next).IsNil() {
		return 0
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return 0
	}

	q := u.Query()
	var page string

	if q.Get("start") != "" {
		page = q.Get("start")
	} else if q.Get("offset") != "" {
		page = q.Get("offset")
	}

	convertedVal, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return convertedVal
}
