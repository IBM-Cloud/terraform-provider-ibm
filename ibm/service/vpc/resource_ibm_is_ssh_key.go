// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"golang.org/x/crypto/ssh"
)

const (
	isKeyName          = "name"
	IsKeyCRN           = "crn"
	isKeyPublicKey     = "public_key"
	isKeyType          = "type"
	isKeyFingerprint   = "fingerprint"
	isKeyLength        = "length"
	isKeyTags          = "tags"
	isKeyResourceGroup = "resource_group"
	isKeyAccessTags    = "access_tags"
	isKeyUserTagType   = "user"
	isKeyAccessTagType = "access"
)

func ResourceIBMISSSHKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISSSHKeyCreate,
		ReadContext:   resourceIBMISSSHKeyRead,
		UpdateContext: resourceIBMISSSHKeyUpdate,
		DeleteContext: resourceIBMISSSHKeyDelete,
		Exists:        resourceIBMISSSHKeyExists,
		Importer:      &schema.ResourceImporter{},

		CustomizeDiff: customdiff.All(
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceTagsCustomizeDiff(diff)
				},
			),
			customdiff.Sequence(
				func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
					return flex.ResourceValidateAccessTags(diff, v)
				}),
		),

		Schema: map[string]*schema.Schema{
			isKeyName: {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
				ValidateFunc: validate.InvokeValidator("ibm_is_security_group", isKeyName),
				Description:  "SSH Key name",
			},

			isKeyPublicKey: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressPublicKeyDiff,
				Description:      "SSH Public key data",
			},

			isKeyType: {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"ed25519", "rsa"}, false),
				Description:  "Key type",
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
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_ssh_key", "tags")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of tags for SSH key",
			},

			isKeyResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group ID",
			},
			// missing schema
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this key.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the key was created.",
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance",
			},

			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			IsKeyCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},

			isKeyAccessTags: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator("ibm_is_ssh_key", "access_tag")},
				Set:         flex.ResourceIBMVPCHash,
				Description: "List of access management tags for SSH key",
			},
		},
	}
}

func ResourceIBMISSHKeyValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 isKeyName,
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`,
			MinValueLength:             1,
			MaxValueLength:             63})

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128})
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "accesstag",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128})

	ibmISSSHKeyResourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_ssh_key", Schema: validateSchema}
	return &ibmISSSHKeyResourceValidator
}

func resourceIBMISSSHKeyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	log.Printf("[DEBUG] Key create")
	name := d.Get(isKeyName).(string)
	publickey := d.Get(isKeyPublicKey).(string)

	diag := keyCreate(context, d, meta, name, publickey)
	if diag != nil {
		return diag
	}
	return resourceIBMISSSHKeyRead(context, d, meta)
}

func keyCreate(context context.Context, d *schema.ResourceData, meta interface{}, name, publickey string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
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

	if keytype, ok := d.GetOk(isKeyType); ok {
		kt := keytype.(string)
		options.Type = &kt
	}

	key, _, err := sess.CreateKeyWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(*key.ID)
	log.Printf("[INFO] Key : %s", *key.ID)

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk(isKeyTags); ok || v != "" {
		oldList, newList := d.GetChange(isKeyTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *key.CRN, "", isKeyUserTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc SSH Key (%s) tags: %s", d.Id(), err)
		}
	}

	if _, ok := d.GetOk(isKeyAccessTags); ok {
		oldList, newList := d.GetChange(isKeyAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *key.CRN, "", isKeyAccessTagType)
		if err != nil {
			log.Printf(
				"Error on create of vpc SSH Key (%s) access tags: %s", d.Id(), err)
		}
	}
	return nil
}

func resourceIBMISSSHKeyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()

	diag := keyGet(context, d, meta, id)
	if diag != nil {
		return diag
	}
	return nil
}

func keyGet(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetKeyOptions{
		ID: &id,
	}
	key, response, err := sess.GetKeyWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(key.Name) {
		if err = d.Set("name", key.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-name").GetDiag()
		}
		if err = d.Set(flex.ResourceName, key.Name); err != nil {
			err = fmt.Errorf("Error setting resource_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-resource_name").GetDiag()
		}
	}
	if err = d.Set("public_key", key.PublicKey); err != nil {
		err = fmt.Errorf("Error setting public_key: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-public_key").GetDiag()
	}
	if !core.IsNil(key.Type) {
		if err = d.Set("type", key.Type); err != nil {
			err = fmt.Errorf("Error setting type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-type").GetDiag()
		}
	}
	if err = d.Set("fingerprint", key.Fingerprint); err != nil {
		err = fmt.Errorf("Error setting fingerprint: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-fingerprint").GetDiag()
	}
	if err = d.Set("length", flex.IntValue(key.Length)); err != nil {
		err = fmt.Errorf("Error setting length: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-length").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(key.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", key.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-href").GetDiag()
	}
	tags, err := flex.GetGlobalTagsUsingCRN(meta, *key.CRN, "", isKeyUserTagType)
	if err != nil {
		log.Printf(
			"Error on get of vpc SSH Key (%s) tags: %s", d.Id(), err)
	}
	d.Set(isKeyTags, tags)
	accesstags, err := flex.GetGlobalTagsUsingCRN(meta, *key.CRN, "", isKeyAccessTagType)
	if err != nil {
		log.Printf(
			"Error on get of vpc SSH Key (%s) access tags: %s", d.Id(), err)
	}
	d.Set(isKeyAccessTags, accesstags)
	controller, err := flex.GetBaseController(meta)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-resource_controller_url").GetDiag()
	}
	d.Set(flex.ResourceControllerURL, controller+"/vpc-ext/compute/sshKeys")
	if err = d.Set(flex.ResourceCRN, key.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-resource_crn").GetDiag()
	}
	if err = d.Set("crn", key.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-crn").GetDiag()
	}
	if key.ResourceGroup != nil {
		if err = d.Set("resource_group", key.ResourceGroup.ID); err != nil {
			err = fmt.Errorf("Error setting resource_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-resource_group").GetDiag()
		}
		if err = d.Set(flex.ResourceGroupName, key.ResourceGroup.Name); err != nil {
			err = fmt.Errorf("Error setting resource_group_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "read", "set-resource_group_name").GetDiag()
		}
	}
	return nil
}

func resourceIBMISSSHKeyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Id()
	name := ""
	hasChanged := false

	if d.HasChange(isKeyName) {
		name = d.Get(isKeyName).(string)
		hasChanged = true
	}

	diag := keyUpdate(context, d, meta, id, name, hasChanged)
	if diag != nil {
		return diag
	}
	return resourceIBMISSSHKeyRead(context, d, meta)
}

func keyUpdate(context context.Context, d *schema.ResourceData, meta interface{}, id, name string, hasChanged bool) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if d.HasChange(isKeyTags) {
		options := &vpcv1.GetKeyOptions{
			ID: &id,
		}
		key, _, err := sess.GetKeyWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isKeyTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *key.CRN, "", isKeyUserTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc SSH Key (%s) tags: %s", id, err)
		}
	}
	if d.HasChange(isKeyAccessTags) {
		options := &vpcv1.GetKeyOptions{
			ID: &id,
		}
		key, _, err := sess.GetKeyWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		oldList, newList := d.GetChange(isKeyAccessTags)
		err = flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *key.CRN, "", isKeyAccessTagType)
		if err != nil {
			log.Printf(
				"Error on update of resource vpc SSH Key (%s) access tags: %s", id, err)
		}
	}
	if hasChanged {
		options := &vpcv1.UpdateKeyOptions{
			ID: &id,
		}
		keyPatchModel := &vpcv1.KeyPatch{
			Name: &name,
		}
		keyPatch, err := keyPatchModel.AsPatch()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AsPatch failed: %s", err.Error()), "ibm_is_ssh_key", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		options.KeyPatch = keyPatch
		_, _, err = sess.UpdateKeyWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return nil
}

func resourceIBMISSSHKeyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Id()

	diag := keyDelete(context, d, meta, id)
	if diag != nil {
		return diag
	}
	return nil
}

func keyDelete(context context.Context, d *schema.ResourceData, meta interface{}, id string) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_ssh_key", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getKeyOptions := &vpcv1.GetKeyOptions{
		ID: &id,
	}
	_, response, err := sess.GetKey(getKeyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &vpcv1.DeleteKeyOptions{
		ID: &id,
	}
	response, err = sess.DeleteKeyWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteKeyWithContext failed: %s", err.Error()), "ibm_is_ssh_key", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId("")
	return nil
}

func resourceIBMISSSHKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id := d.Id()

	exists, err := keyExists(d, meta, id)
	return exists, err
}

func keyExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	options := &vpcv1.GetKeyOptions{
		ID: &id,
	}
	_, response, err := sess.GetKey(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("[ERROR] Error getting SSH Key: %s\n%s", err, response)
	}
	return true, nil
}

// to suppress any change shown when keys are same
func suppressPublicKeyDiff(k, old, new string, d *schema.ResourceData) bool {
	// if there are extra spaces or new lines, suppress that change
	if strings.Compare(strings.TrimSpace(old), strings.TrimSpace(new)) != 0 {
		// if old is empty
		if old != "" {
			//create a new piblickey object from the string
			usePK, error := parseKey(new)
			if error != nil {
				return false
			}
			// returns the key in byte format with an extra added new line at the end
			newkey := strings.TrimRight(string(ssh.MarshalAuthorizedKey(usePK)), "\n")
			// check if both keys are same, if yes suppress the change
			return strings.TrimSpace(strings.TrimPrefix(newkey, old)) == ""
		} else {
			return strings.TrimSpace(strings.TrimPrefix(new, old)) == ""
		}
	} else {
		return true
	}
}

// takes a string and returns public key object
func parseKey(s string) (ssh.PublicKey, error) {
	keyBytes := []byte(s)

	// Accepts formats of PublicKey:
	// - <base64 key>
	// - ssh-rsa/ssh-ed25519 <base64 key>
	// - ssh-rsa/ssh-ed25519 <base64 key> <comment>
	// if PublicKey provides other than just base64 key, then first part must be "ssh-rsa" or "ssh-ed25519"
	if subStrs := strings.Split(s, " "); len(subStrs) > 1 && subStrs[0] != "ssh-rsa" && subStrs[0] != "ssh-ed25519" {
		return nil, errors.New("not an RSA key OR ED25519 key")
	}

	pk, _, _, _, e := ssh.ParseAuthorizedKey(keyBytes)
	if e == nil {
		return pk, nil
	}

	decodedKey := make([]byte, base64.StdEncoding.DecodedLen(len(keyBytes)))
	n, e := base64.StdEncoding.Decode(decodedKey, keyBytes)
	if e != nil {
		return nil, e
	}
	decodedKey = decodedKey[:n]

	pk, e = ssh.ParsePublicKey(decodedKey)
	if e == nil {
		return pk, nil
	}
	return nil, e
}
