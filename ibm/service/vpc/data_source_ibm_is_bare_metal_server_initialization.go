// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/ScaleFT/sshkeys"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	isBareMetalServerImageName                = "image_name"
	isBareMetalServerUserAccountUserName      = "username"
	isBareMetalServerUserAccountEncryptionKey = "encryption_key"
	isBareMetalServerUserAccountEncPwd        = "encrypted_password"
	isBareMetalServerUserAccountPassword      = "password"
	isBareMetalServerPEM                      = "private_key"
	isBareMetalServerPassphrase               = "passphrase"
	isBareMetalServerUserAccountResourceType  = "resource_type"
)

func DataSourceIBMIsBareMetalServerInitialization() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMISBareMetalServerInitializationRead,

		Schema: map[string]*schema.Schema{
			isBareMetalServerID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The bare metal server identifier",
			},

			isBareMetalServerPEM: {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Bare Metal Server Private Key file",
			},

			isBareMetalServerPassphrase: {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Passphrase for Bare Metal Server Private Key file",
			},

			isBareMetalServerDefaultTrustedProfile: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_link": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If set to true, the system will create a link to the specified target trusted profile during server creation. Regardless of whether a link is created by the system or manually using the IAM Identity service, it will be automatically deleted when the server is deleted.",
						},
						"target": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The default IAM trusted profile to use for this bare metal server",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique identifier for this trusted profile",
									},
									"crn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN for this trusted profile",
									},
								},
							},
						},
					},
				},
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
						isBareMetalServerUserAccountEncPwd: {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The password at initialization, encrypted using encryption_key, and returned base64-encoded",
						},
						isBareMetalServerUserAccountPassword: {
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The password at initialization, encrypted using encryption_key, and returned base64-encoded",
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

func dataSourceIBMISBareMetalServerInitializationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	bareMetalServerID := d.Get(isBareMetalServerID).(string)
	sess, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	options := &vpcv1.GetBareMetalServerInitializationOptions{
		ID: &bareMetalServerID,
	}

	initialization, _, err := sess.GetBareMetalServerInitializationWithContext(context, options)
	if err != nil || initialization == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerInitializationWithContext failed: %s", err.Error()), "(Data) ibm_is_bare_metal_server_initialization", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(bareMetalServerID)
	if initialization.Image != nil {

		if err = d.Set(isBareMetalServerImage, initialization.Image.ID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting image: %s", err), "(Data) ibm_is_bare_metal_server_initialization", "read", "set-image").GetDiag()
		}

		if err = d.Set(isBareMetalServerImageName, initialization.Image.Name); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting image_name: %s", err), "(Data) ibm_is_bare_metal_server_initialization", "read", "set-image_name").GetDiag()
		}
	}

	if initialization.DefaultTrustedProfile != nil {
		defaultTrustedProfileList := make([]map[string]interface{}, 0)
		defaultTrustedProfileMap := map[string]interface{}{}

		targetMap := map[string]interface{}{}
		targetMap["id"] = *initialization.DefaultTrustedProfile.Target.ID
		targetMap["crn"] = *initialization.DefaultTrustedProfile.Target.CRN

		defaultTrustedProfileMap["auto_link"] = *initialization.DefaultTrustedProfile.AutoLink
		defaultTrustedProfileMap["target"] = []map[string]interface{}{targetMap}

		defaultTrustedProfileList = append(defaultTrustedProfileList, defaultTrustedProfileMap)
		if err = d.Set(isBareMetalServerDefaultTrustedProfile, defaultTrustedProfileList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting default_trusted_profile: %s", err), "(Data) ibm_is_bare_metal_server_initialization", "read", "set-default_trusted_profile").GetDiag()
		}
	}

	var keys []string
	keys = make([]string, 0)
	if initialization.Keys != nil {
		for _, key := range initialization.Keys {
			keys = append(keys, *key.ID)
		}
	}

	if err = d.Set(isBareMetalServerKeys, flex.NewStringSet(schema.HashString, keys)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting keys: %s", err), "(Data) ibm_is_bare_metal_server_initialization", "read", "set-keys").GetDiag()
	}
	accList := make([]map[string]interface{}, 0)
	if initialization.UserAccounts != nil {

		for _, accIntf := range initialization.UserAccounts {
			acc := accIntf.(*vpcv1.BareMetalServerInitializationUserAccount)
			currAccount := map[string]interface{}{
				isBareMetalServerUserAccountUserName: *acc.Username,
			}
			currAccount[isBareMetalServerUserAccountResourceType] = *acc.ResourceType
			currAccount[isBareMetalServerUserAccountEncryptionKey] = *acc.EncryptionKey.CRN
			encPassword := base64.StdEncoding.EncodeToString(*acc.EncryptedPassword)
			currAccount[isBareMetalServerUserAccountEncPwd] = encPassword

			var rsaKey *rsa.PrivateKey
			if privatekey, ok := d.GetOk(isBareMetalServerPEM); ok {
				keyFlag := privatekey.(string)
				keybytes := []byte(keyFlag)

				if keyFlag != "" {
					block, err := pem.Decode(keybytes)
					if block == nil {
						err := fmt.Errorf("[ERROR] Failed to load the private key from the given key contents. Instead of the key file path, please make sure the private key is pem format (%v)", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "decode-pem").GetDiag()

					}
					isEncrypted := false
					if block.Type == "OPENSSH PRIVATE KEY" {
						var err error
						isEncrypted, err = isOpenSSHPrivKeyEncrypted(block.Bytes)
						if err != nil {
							err = fmt.Errorf("[ERROR] Failed to check if the provided open ssh key is encrypted or not %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "check-encryption").GetDiag()
						}
					} else {
						isEncrypted = x509.IsEncryptedPEMBlock(block)
					}
					passphrase := ""
					var privateKey interface{}
					if isEncrypted {
						if pass, ok := d.GetOk(isBareMetalServerPassphrase); ok {
							passphrase = pass.(string)
						} else {
							err := fmt.Errorf("[ERROR] Mandatory field 'passphrase' not provided")
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "missing-passphrase").GetDiag()
						}
						var err error
						privateKey, err = sshkeys.ParseEncryptedRawPrivateKey(keybytes, []byte(passphrase))
						if err != nil {
							err := fmt.Errorf("[ERROR] Fail to decrypting the private key: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "decrypt-private_key").GetDiag()
						}
					} else {
						var err error
						privateKey, err = sshkeys.ParseEncryptedRawPrivateKey(keybytes, nil)
						if err != nil {
							err := fmt.Errorf("[ERROR] Fail to decrypting the private key: %s", err)
							return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "fail-decrypt-private_key").GetDiag()
						}
					}
					var ok bool
					rsaKey, ok = privateKey.(*rsa.PrivateKey)
					if !ok {
						err := fmt.Errorf("[ERROR] Failed to convert to RSA private key")
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "convert-rsa").GetDiag()
					}
				}
			}

			if acc.EncryptedPassword != nil {
				ciphertext := *acc.EncryptedPassword
				password := base64.StdEncoding.EncodeToString(ciphertext)
				if rsaKey != nil {
					rng := rand.Reader
					clearPassword, err := rsa.DecryptPKCS1v15(rng, rsaKey, ciphertext)
					if err != nil {
						err := fmt.Errorf("[ERROR] Can not decrypt the password with the given key, %s", err)
						return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_bare_metal_server_initialization", "read", "decrypt-password").GetDiag()
					}
					password = string(clearPassword)
				}
				currAccount[isBareMetalServerUserAccountPassword] = password
			}
			accList = append(accList, currAccount)
		}

		if err = d.Set(isBareMetalServerUserAccounts, accList); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user_accounts: %s", err), "(Data) ibm_is_bare_metal_server_initialization", "read", "set-user_accounts").GetDiag()
		}
	}
	return nil
}
