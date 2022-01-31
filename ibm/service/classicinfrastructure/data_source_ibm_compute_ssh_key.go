// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure

import (
	"fmt"
	"sort"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/services"
)

func DataSourceIBMComputeSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMComputeSSHKeyRead,

		Schema: map[string]*schema.Schema{
			"label": {
				Description: "The label associated with the ssh key",
				Type:        schema.TypeString,
				Required:    true,
			},

			"public_key": {
				Description: "The public ssh key",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"fingerprint": {
				Description: "A sequence of bytes to authenticate or lookup a longer ssh key",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"notes": {
				Description: "A small note about a ssh key to use at your discretion",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"most_recent": {
				Description: "If true and multiple entries are found, the most recently created key is used. " +
					"If false, an error is returned",
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dataSourceIBMComputeSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(conns.ClientSession).SoftLayerSession()
	service := services.GetAccountService(sess)

	label := d.Get("label").(string)
	mostRecent := d.Get("most_recent").(bool)

	keys, err := service.
		Filter(filter.Build(filter.Path("sshKeys.label").Eq(label))).
		Mask("id,label,key,fingerprint,notes,createDate").
		GetSshKeys()

	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving SSH key: %s", err)
	}
	if len(keys) == 0 {
		return fmt.Errorf("[ERROR] No ssh key found with name [%s]", label)
	}

	var key datatypes.Security_Ssh_Key
	if len(keys) > 1 {
		if mostRecent {
			key = mostRecentSSHKey(keys)
		} else {
			return fmt.Errorf(
				"More than one ssh key found with label matching [%s]. "+
					"Either set 'most_recent' to true in your "+
					"configuration to force the most recent ssh key "+
					"to be used, or ensure that the label is unique", label)
		}
	} else {
		key = keys[0]
	}

	d.SetId(fmt.Sprintf("%d", *key.Id))
	d.Set("label", label)
	d.Set("public_key", strings.TrimSpace(*key.Key))
	d.Set("fingerprint", key.Fingerprint)
	d.Set("notes", key.Notes)
	return nil
}

type sshKeys []datatypes.Security_Ssh_Key

func (k sshKeys) Len() int { return len(k) }

func (k sshKeys) Swap(i, j int) { k[i], k[j] = k[j], k[i] }

func (k sshKeys) Less(i, j int) bool {
	return k[i].CreateDate.Before(k[j].CreateDate.Time)
}

func mostRecentSSHKey(keys sshKeys) datatypes.Security_Ssh_Key {
	sortedKeys := keys
	sort.Sort(sortedKeys)
	return sortedKeys[len(sortedKeys)-1]
}
