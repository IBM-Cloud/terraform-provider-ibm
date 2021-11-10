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

	// "github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

const (
	isKeys                 = "keys"
	isKeyCreatedAt         = "created_at"
	isKeyCRN               = "crn"
	isKeysHref             = "href"
	isKeyId                = "id"
	isKeyResourceGroupHref = "href"
	isKeyResourceGroupId   = "id"
	isKeyResourceGroupName = "name"
	isKeysLimit            = "limit"
	isKeysTotalCount       = "total_count"
)

func dataSourceIBMIsSshKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsSshKeysRead,

		Schema: map[string]*schema.Schema{
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A link to the first page of resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			isKeys: &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of keys.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isKeyCreatedAt: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the key was created.",
						},
						isKeyCRN: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for this key.",
						},
						isKeyFingerprint: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The fingerprint for this key.  The value is returned base64-encoded and prefixed with the hash algorithm (always `SHA256`).",
						},
						isKeysHref: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this key.",
						},
						isKeyId: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this key.",
						},
						isKeyLength: &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The length of this key (in bits).",
						},
						isKeyName: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique user-defined name for this key. If unspecified, the name will be a hyphenated list of randomly-selected words.",
						},
						isKeyPublicKey: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public SSH key, consisting of two space-separated fields: the algorithm name, and the base64-encoded key.",
						},
						isKeyResourceGroup: &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resource group for this key.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									isKeyResourceGroupHref: &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL for this resource group.",
									},
									isKeyResourceGroupId: &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this resource group.",
									},
									isKeyResourceGroupName: &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The user-defined name for this resource group.",
									},
								},
							},
						},
						isKeyType: &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crypto-system used by this key.",
						},
					},
				},
			},
			isKeysLimit: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of resources that can be returned by the request.",
			},
			"next": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A link to the next page of resources. This property is present for all pagesexcept the last page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for a page of resources.",
						},
					},
				},
			},
			isKeysTotalCount: &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources across all pages.",
			},
		},
	}
}

func dataSourceIBMIsSshKeysRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	listKeysOptions := &vpcv1.ListKeysOptions{}

	keyCollection, response, err := vpcClient.ListKeysWithContext(context, listKeysOptions)
	if err != nil {
		log.Printf("[DEBUG] ListKeysWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListKeysWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIBMIsSshKeysID(d))

	if keyCollection.First != nil {
		err = d.Set("first", dataSourceKeyCollectionFlattenFirst(*keyCollection.First))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting first %s", err))
		}
	}

	if keyCollection.Keys != nil {
		err = d.Set(isKeys, dataSourceKeyCollectionFlattenKeys(keyCollection.Keys))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting keys %s", err))
		}
	}
	if err = d.Set(isKeysLimit, intValue(keyCollection.Limit)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting limit: %s", err))
	}

	if keyCollection.Next != nil {
		err = d.Set("next", dataSourceKeyCollectionFlattenNext(*keyCollection.Next))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting next %s", err))
		}
	}
	if err = d.Set(isKeysTotalCount, intValue(keyCollection.TotalCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting total_count: %s", err))
	}

	return nil
}

// dataSourceIBMIsSshKeysID returns a reasonable ID for the list.
func dataSourceIBMIsSshKeysID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceKeyCollectionFlattenFirst(result vpcv1.KeyCollectionFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceKeyCollectionFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceKeyCollectionFirstToMap(firstItem vpcv1.KeyCollectionFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}

func dataSourceKeyCollectionFlattenKeys(result []vpcv1.Key) (keys []map[string]interface{}) {
	for _, keysItem := range result {
		keys = append(keys, dataSourceKeyCollectionKeysToMap(keysItem))
	}

	return keys
}

func dataSourceKeyCollectionKeysToMap(keysItem vpcv1.Key) (keysMap map[string]interface{}) {
	keysMap = map[string]interface{}{}

	if keysItem.CreatedAt != nil {
		keysMap[isKeyCreatedAt] = keysItem.CreatedAt.String()
	}
	if keysItem.CRN != nil {
		keysMap[isKeyCRN] = keysItem.CRN
	}
	if keysItem.Fingerprint != nil {
		keysMap[isKeyFingerprint] = keysItem.Fingerprint
	}
	if keysItem.Href != nil {
		keysMap[isKeysHref] = keysItem.Href
	}
	if keysItem.ID != nil {
		keysMap[isKeyId] = keysItem.ID
	}
	if keysItem.Length != nil {
		keysMap[isKeyLength] = keysItem.Length
	}
	if keysItem.Name != nil {
		keysMap[isKeyName] = keysItem.Name
	}
	if keysItem.PublicKey != nil {
		keysMap[isKeyPublicKey] = keysItem.PublicKey
	}
	if keysItem.ResourceGroup != nil {
		resourceGroupList := []map[string]interface{}{}
		resourceGroupMap := dataSourceKeyCollectionKeysResourceGroupToMap(*keysItem.ResourceGroup)
		resourceGroupList = append(resourceGroupList, resourceGroupMap)
		keysMap[isKeyResourceGroup] = resourceGroupList
	}
	if keysItem.Type != nil {
		keysMap[isKeyType] = keysItem.Type
	}

	return keysMap
}

func dataSourceKeyCollectionKeysResourceGroupToMap(resourceGroupItem vpcv1.ResourceGroupReference) (resourceGroupMap map[string]interface{}) {
	resourceGroupMap = map[string]interface{}{}

	if resourceGroupItem.Href != nil {
		resourceGroupMap[isKeyResourceGroupHref] = resourceGroupItem.Href
	}
	if resourceGroupItem.ID != nil {
		resourceGroupMap[isKeyResourceGroupId] = resourceGroupItem.ID
	}
	if resourceGroupItem.Name != nil {
		resourceGroupMap[isKeyResourceGroupName] = resourceGroupItem.Name
	}

	return resourceGroupMap
}

func dataSourceKeyCollectionFlattenNext(result vpcv1.KeyCollectionNext) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceKeyCollectionNextToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceKeyCollectionNextToMap(nextItem vpcv1.KeyCollectionNext) (nextMap map[string]interface{}) {
	nextMap = map[string]interface{}{}

	if nextItem.Href != nil {
		nextMap["href"] = nextItem.Href
	}

	return nextMap
}
