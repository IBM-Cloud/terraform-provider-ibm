package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
)

func dataSourceIBMISSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMISSSHKeyRead,

		Schema: map[string]*schema.Schema{
			isKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			isKeyType: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isKeyFingerprint: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isKeyLength: {
				Type:     schema.TypeInt,
				Computed: true,
			},
			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},
			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func dataSourceIBMISSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	keyC := compute.NewKeyClient(sess)

	name := d.Get(isKeyName).(string)

	keys, _, err := keyC.List("")
	if err != nil {
		return err
	}

	for _, key := range keys {
		if key.Name == name {
			d.SetId(key.ID.String())
			d.Set(isKeyName, key.Name)
			d.Set(isKeyType, key.Type)
			d.Set(isKeyFingerprint, key.Fingerprint)
			d.Set(isKeyLength, key.Length)
			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			if sess.Generation == 1 {
				d.Set(ResourceControllerURL, controller+"/vpc/compute/sshKeys")
			} else {
				d.Set(ResourceControllerURL, controller+"/vpc-ext/compute/sshKeys")
			}
			d.Set(ResourceName, key.Name)
			d.Set(ResourceCRN, key.Crn)
			if key.ResourceGroup != nil {
				d.Set(ResourceGroupName, key.ResourceGroup.Name)
			}
			return nil
		}
	}
	return fmt.Errorf("No SSH Key found with name %s", name)
}
