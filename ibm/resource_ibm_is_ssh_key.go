package ibm

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	isKeyName          = "name"
	isKeyPublicKey     = "public_key"
	isKeyType          = "type"
	isKeyFingerprint   = "fingerprint"
	isKeyLength        = "length"
	isKeyTags          = "tags"
	isKeyResourceGroup = "resource_group"
)

func resourceIBMISSSHKey() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISSSHKeyCreate,
		Read:     resourceIBMISSSHKeyRead,
		Update:   resourceIBMISSSHKeyUpdate,
		Delete:   resourceIBMISSSHKeyDelete,
		Exists:   resourceIBMISSSHKeyExists,
		Importer: &schema.ResourceImporter{},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{
			isKeyName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validateISName,
				Description:  "SSH Key name",
			},

			isKeyPublicKey: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "SSH Public key data",
			},

			isKeyType: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key type",
			},

			isKeyFingerprint: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SSH key Fingerprint info",
			},

			isKeyLength: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "SSH key Length",
			},
			isKeyTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "List of tags for SSH key",
			},

			isKeyResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group ID",
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

func resourceIBMISSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Key create")
	name := d.Get(isKeyName).(string)
	publickey := d.Get(isKeyPublicKey).(string)

	if userDetails.generation == 1 {
		err := classicKeyCreate(d, meta, name, publickey)
		if err != nil {
			return err
		}
	} else {
		err := keyCreate(d, meta, name, publickey)
		if err != nil {
			return err
		}
	}
	return resourceIBMISSSHKeyRead(d, meta)
}

func classicKeyCreate(d *schema.ResourceData, meta interface{}, name, publickey string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.CreateKeyOptions{
		PublicKey: &publickey,
		Name:      &name,
	}

	if rgrp, ok := d.GetOk(isKeyResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcclassicv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	key, response, err := sess.CreateKey(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create SSH Key %s\n%s", err, response)
	}
	d.SetId(*key.ID)
	log.Printf("[INFO] Key : %s", *key.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isKeyTags); ok || v != "" {
		oldList, newList := d.GetChange(isKeyTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *key.Crn)
		if err != nil {
			return fmt.Errorf(
				"Error on create of vpc SSH Key (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func keyCreate(d *schema.ResourceData, meta interface{}, name, publickey string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateKeyOptions{
		PublicKey: &publickey,
		Name:      &name,
	}

	if rgrp, ok := d.GetOk(isKeyResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}

	key, response, err := sess.CreateKey(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] Create SSH Key %s\n%s", err, response)
	}
	d.SetId(*key.ID)
	log.Printf("[INFO] Key : %s", *key.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isKeyTags); ok || v != "" {
		oldList, newList := d.GetChange(isKeyTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *key.Crn)
		if err != nil {
			return fmt.Errorf(
				"Error on create of vpc SSH Key (%s) tags: %s", d.Id(), err)
		}
	}
	return nil
}

func resourceIBMISSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	id := d.Id()
	if userDetails.generation == 1 {
		err := classicKeyGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := keyGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicKeyGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.GetKeyOptions{
		ID: &id,
	}
	key, response, err := sess.GetKey(options)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting SSH Key (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isKeyName, *key.Name)
	d.Set(isKeyPublicKey, *key.PublicKey)
	d.Set(isKeyType, *key.Type)
	d.Set(isKeyFingerprint, *key.Fingerprint)
	d.Set(isKeyLength, *key.Length)
	tags, err := GetTagsUsingCRN(meta, *key.Crn)
	if err != nil {
		return fmt.Errorf(
			"Error on get of vpc SSH Key (%s) tags: %s", d.Id(), err)
	}
	d.Set(isKeyTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/compute/sshKeys")
	d.Set(ResourceName, *key.Name)
	d.Set(ResourceCRN, *key.Crn)
	if key.ResourceGroup != nil {
		d.Set(ResourceGroupName, *key.ResourceGroup.ID)
		d.Set(isKeyResourceGroup, *key.ResourceGroup.ID)
	}
	return nil
}

func keyGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetKeyOptions{
		ID: &id,
	}
	key, response, err := sess.GetKey(options)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting SSH Key (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		d.SetId("")
		return nil
	}
	d.Set(isKeyName, *key.Name)
	d.Set(isKeyPublicKey, *key.PublicKey)
	d.Set(isKeyType, *key.Type)
	d.Set(isKeyFingerprint, *key.Fingerprint)
	d.Set(isKeyLength, *key.Length)
	tags, err := GetTagsUsingCRN(meta, *key.Crn)
	if err != nil {
		return fmt.Errorf(
			"Error on get of vpc SSH Key (%s) tags: %s", d.Id(), err)
	}
	d.Set(isKeyTags, tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc-ext/compute/sshKeys")
	d.Set(ResourceName, *key.Name)
	d.Set(ResourceCRN, *key.Crn)
	if key.ResourceGroup != nil {
		d.Set(ResourceGroupName, *key.ResourceGroup.Name)
		d.Set(isKeyResourceGroup, *key.ResourceGroup.ID)
	}
	return nil
}

func resourceIBMISSSHKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	id := d.Id()
	name := ""
	hasChanged := false

	if d.HasChange(isKeyName) {
		name = d.Get(isKeyName).(string)
		hasChanged = true
	}

	if userDetails.generation == 1 {
		err := classicKeyUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	} else {
		err := keyUpdate(d, meta, id, name, hasChanged)
		if err != nil {
			return err
		}
	}
	return resourceIBMISSSHKeyRead(d, meta)
}

func classicKeyUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isKeyTags) {
		options := &vpcclassicv1.GetKeyOptions{
			ID: &id,
		}
		key, response, err := sess.GetKey(options)
		if err != nil {
			return fmt.Errorf("Error getting SSH Key : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isKeyTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *key.Crn)
		if err != nil {
			return fmt.Errorf(
				"Error on update of resource vpc SSH Key (%s) tags: %s", id, err)
		}
	}
	if hasChanged {
		options := &vpcclassicv1.UpdateKeyOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateKey(options)
		if err != nil {
			return fmt.Errorf("Error updating vpc SSH Key: %s\n%s", err, response)
		}
	}
	return nil
}

func keyUpdate(d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	if d.HasChange(isKeyTags) {
		options := &vpcv1.GetKeyOptions{
			ID: &id,
		}
		key, response, err := sess.GetKey(options)
		if err != nil {
			return fmt.Errorf("Error getting SSH Key : %s\n%s", err, response)
		}
		oldList, newList := d.GetChange(isKeyTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, *key.Crn)
		if err != nil {
			return fmt.Errorf(
				"Error on update of resource vpc SSH Key (%s) tags: %s", id, err)
		}
	}
	if hasChanged {
		options := &vpcv1.UpdateKeyOptions{
			ID:   &id,
			Name: &name,
		}
		_, response, err := sess.UpdateKey(options)
		if err != nil {
			return fmt.Errorf("Error updating vpc SSH Key: %s\n%s", err, response)
		}
	}
	return nil
}

func resourceIBMISSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicKeyDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := keyDelete(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicKeyDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	getKeyOptions := &vpcclassicv1.GetKeyOptions{
		ID: &id,
	}
	_, response, err := sess.GetKey(getKeyOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting SSH Key (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	options := &vpcclassicv1.DeleteKeyOptions{
		ID: &id,
	}
	response, err = sess.DeleteKey(options)
	if err != nil {
		return fmt.Errorf("Error Deleting SSH Key : %s\n%s", err, response)
	}
	d.SetId("")
	return nil
}

func keyDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	getKeyOptions := &vpcv1.GetKeyOptions{
		ID: &id,
	}
	_, response, err := sess.GetKey(getKeyOptions)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error Getting SSH Key (%s): %s\n%s", id, err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}

	options := &vpcv1.DeleteKeyOptions{
		ID: &id,
	}
	response, err = sess.DeleteKey(options)
	if err != nil {
		return fmt.Errorf("Error Deleting SSH Key : %s\n%s", err, response)
	}
	d.SetId("")
	return nil
}

func resourceIBMISSSHKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()

	if userDetails.generation == 1 {
		err := classicKeyExists(d, meta, id)
		if err != nil {
			return false, err
		}
	} else {
		err := keyExists(d, meta, id)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func classicKeyExists(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.GetKeyOptions{
		ID: &id,
	}
	_, response, err := sess.GetKey(options)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting SSH Key: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}

func keyExists(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.GetKeyOptions{
		ID: &id,
	}
	_, response, err := sess.GetKey(options)
	if err != nil && response.StatusCode != 404 {
		return fmt.Errorf("Error getting SSH Key: %s\n%s", err, response)
	}
	if response.StatusCode == 404 {
		return nil
	}
	return nil
}
