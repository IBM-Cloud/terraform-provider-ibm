// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerImageName                = "image_name"
	isBareMetalServerUserAccountUserName      = "username"
	isBareMetalServerUserAccountEncryptionKey = "encryption_key"
	isBareMetalServerUserAccountResourceType  = "resource_type"
)

func dataSourceBareMetalServerInitialization() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISBareMetalServerInitializationRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},
			isBareMetalServerImage: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the image the bare metal server was provisioned from",
			},
			isBareMetalServerImageName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user-defined or system-provided name for the image the bare metal server was provisioned from",
			},

			isBareMetalServerUserAccounts: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The user accounts that are created at initialization. There can be multiple account types distinguished by the resource_type attribute.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isBareMetalServerUserAccountUserName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username for the account created at initialization",
						},
						isBareMetalServerUserAccountEncryptionKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN for the encryption key",
						},
						isBareMetalServerUserAccountResourceType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource referenced : [ host_user_account ]",
						},
					},
				},
			},

			isBareMetalServerKeys: {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "SSH key Ids for the bare metal server",
			},
		},
	}
}

func dataSourceIBMISBareMetalServerInitializationRead(d *schema.ResourceData, meta interface{}) error {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	err := bmsGetInitializationById(d, meta, bareMetalServerID)
	if err != nil {
		return err
	}
	return nil
}

func bmsGetInitializationById(d *schema.ResourceData, meta interface{}, bareMetalServerID string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetBareMetalServerInitializationOptions{
		ID: &bareMetalServerID,
	}

	initialization, response, err := sess.GetBareMetalServerInitialization(options)
	if err != nil || initialization == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting Bare Metal Server (%s) initialization : %s\n%s", bareMetalServerID, err, response)
	}
	d.SetId(bareMetalServerID)
	if initialization.Image != nil {
		d.Set(isBareMetalServerImage, initialization.Image.ID)
		d.Set(isBareMetalServerImageName, initialization.Image.Name)
	}
	var keys []string
	keys = make([]string, 0)
	if initialization.Keys != nil {
		for _, key := range initialization.Keys {
			keys = append(keys, *key.ID)
		}
	}
	d.Set(isBareMetalServerKeys, newStringSet(schema.HashString, keys))
	accList := make([]map[string]interface{}, 0)
	if initialization.UserAccounts != nil {

		for _, accIntf := range initialization.UserAccounts {
			acc := accIntf.(*vpcv1.BareMetalServerInitializationUserAccountsItem)
			currAccount := map[string]interface{}{
				isBareMetalServerUserAccountUserName: *acc.Username,
			}
			currAccount[isBareMetalServerUserAccountResourceType] = *acc.ResourceType
			currAccount[isBareMetalServerUserAccountEncryptionKey] = *acc.EncryptionKey.CRN

			accList = append(accList, currAccount)
		}
		d.Set(isBareMetalServerUserAccounts, accList)
	}

	return nil
}
