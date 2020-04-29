package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	name := d.Get(isKeyName).(string)
	if userDetails.generation == 1 {
		err := classicKeyGet(d, meta, name)
		if err != nil {
			return err
		}
	} else {
		err := keyGet(d, meta, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicKeyGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	listKeysOptions := &vpcclassicv1.ListKeysOptions{}
	keys, _, err := sess.ListKeys(listKeysOptions)
	if err != nil {
		return err
	}
	for _, key := range keys.Keys {
		if *key.Name == name {
			d.SetId(*key.ID)
			d.Set("name", *key.Name)
			d.Set(isKeyType, *key.Type)
			d.Set(isKeyFingerprint, *key.Fingerprint)
			d.Set(isKeyLength, *key.Length)
			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			d.Set(ResourceControllerURL, controller+"/vpc/compute/sshKeys")
			d.Set(ResourceName, *key.Name)
			d.Set(ResourceCRN, *key.Crn)
			if key.ResourceGroup != nil {
				d.Set(ResourceGroupName, *key.ResourceGroup.ID)
			}
			return nil
		}
	}
	return fmt.Errorf("No SSH Key found with name %s", name)
}

func keyGet(d *schema.ResourceData, meta interface{}, name string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	listKeysOptions := &vpcv1.ListKeysOptions{}
	keys, _, err := sess.ListKeys(listKeysOptions)
	if err != nil {
		return err
	}
	for _, key := range keys.Keys {
		if *key.Name == name {
			d.SetId(*key.ID)
			d.Set("name", *key.Name)
			d.Set(isKeyType, *key.Type)
			d.Set(isKeyFingerprint, *key.Fingerprint)
			d.Set(isKeyLength, *key.Length)
			controller, err := getBaseController(meta)
			if err != nil {
				return err
			}
			d.Set(ResourceControllerURL, controller+"/vpc/compute/sshKeys")
			d.Set(ResourceName, *key.Name)
			d.Set(ResourceCRN, *key.Crn)
			if key.ResourceGroup != nil {
				d.Set(ResourceGroupName, *key.ResourceGroup.ID)
			}
			return nil
		}
	}
	return fmt.Errorf("No SSH Key found with name %s", name)
}
