package ibm

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/vpn"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isIpSecName              = "name"
	isIpSecAuthenticationAlg = "authentication_algorithm"
	isIpSecEncryptionAlg     = "encryption_algorithm"
	isIpSecPFS               = "pfs"
	isIpSecKeyLifeTime       = "key_lifetime"
	isIPSecResourceGroup     = "resource_group"
	isIPSecEncapsulationMode = "encapsulation_mode"
	isIPSecTransformProtocol = "transform_protocol"
	isIPSecVPNConnections    = "vpn_connections"
	isIPSecVPNConnectionName = "name"
	isIPSecVPNConnectionId   = "id"
	isIPSecVPNConnectionHref = "href"
)

func resourceIBMISIPSecPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISIPSecPolicyCreate,
		Read:     resourceIBMISIPSecPolicyRead,
		Update:   resourceIBMISIPSecPolicyUpdate,
		Delete:   resourceIBMISIPSecPolicyDelete,
		Exists:   resourceIBMISIPSecPolicyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isIpSecName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isIpSecAuthenticationAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"md5", "sha1", "sha256"}),
			},

			isIpSecEncryptionAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"3des", "aes128", "aes256"}),
			},

			isIpSecPFS: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"disabled", "group_2", "group_5", "group_14"}),
			},

			isIPSecResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isIpSecKeyLifeTime: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      3600,
				ValidateFunc: validateKeyLifeTime,
			},

			isIPSecEncapsulationMode: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isIPSecTransformProtocol: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isIPSecVPNConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isIPSecVPNConnectionName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIPSecVPNConnectionId: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIPSecVPNConnectionHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMISIPSecPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Ip Sec create")
	name := d.Get(isIpSecName).(string)
	authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
	encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
	pfs := d.Get(isIpSecPFS).(string)
	resourceGrpId := d.Get(isIPSecResourceGroup).(string)
	keyLifetime := d.Get(isIpSecKeyLifeTime).(int)

	vpnC := vpn.NewVpnClient(sess)
	ipSec, err := vpnC.CreateIpsecPolicy(authenticationAlg, encryptionAlg, name, pfs, resourceGrpId, keyLifetime)
	if err != nil {
		log.Printf("[DEBUG] ipsec err %s", err)
		return err
	}

	d.SetId(ipSec.ID.String())
	log.Printf("[INFO] Ipsec : %s", ipSec.ID.String())
	return resourceIBMISIPSecPolicyRead(d, meta)
}

func resourceIBMISIPSecPolicyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpnC := vpn.NewVpnClient(sess)

	ipSec, err := vpnC.GetIpsecPolicy(d.Id())
	if err != nil {
		return err
	}

	d.Set(isIpSecName, ipSec.Name)
	d.Set(isIpSecAuthenticationAlg, ipSec.AuthenticationAlgorithm)
	d.Set(isIpSecEncryptionAlg, ipSec.EncryptionAlgorithm)
	if ipSec.ResourceGroup != nil {
		d.Set(isIPSecResourceGroup, ipSec.ResourceGroup.ID)
	} else {
		d.Set(isIPSecResourceGroup, nil)
	}
	d.Set(isIpSecPFS, ipSec.Pfs)
	if ipSec.KeyLifetime != 0 {
		d.Set(isIpSecKeyLifeTime, ipSec.KeyLifetime)
	}
	d.Set(isIPSecEncapsulationMode, ipSec.EncapsulationMode)
	d.Set(isIPSecTransformProtocol, ipSec.TransformProtocol)

	connList := make([]map[string]interface{}, 0)
	if ipSec.Connections != nil && len(ipSec.Connections) > 0 {
		for _, connection := range ipSec.Connections {
			conn := map[string]interface{}{}
			conn[isIPSecVPNConnectionName] = connection.Name
			conn[isIPSecVPNConnectionId] = connection.ID.String()
			conn[isIPSecVPNConnectionHref] = connection.Href
			connList = append(connList, conn)
		}
	}
	d.Set(isIPSecVPNConnections, connList)
	return nil
}

func resourceIBMISIPSecPolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpnC := vpn.NewVpnClient(sess)

	if d.HasChange(isIpSecName) || d.HasChange(isIpSecAuthenticationAlg) || d.HasChange(isIpSecEncryptionAlg) || d.HasChange(isIpSecPFS) || d.HasChange(isIpSecKeyLifeTime) {
		name := d.Get(isIpSecName).(string)
		authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
		encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
		pfs := d.Get(isIpSecPFS).(string)
		keyLifetime := d.Get(isIpSecKeyLifeTime).(int)
		_, err := vpnC.UpdateIpsecPolicy(d.Id(), authenticationAlg, encryptionAlg, name, pfs, keyLifetime)
		if err != nil {
			return err
		}
	}

	return resourceIBMISIPSecPolicyRead(d, meta)
}

func resourceIBMISIPSecPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpnC := vpn.NewVpnClient(sess)
	err = vpnC.DeleteIpsecPolicy(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISIPSecPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	vpnC := vpn.NewVpnClient(sess)

	_, err = vpnC.GetIpsecPolicy(d.Id())
	if err != nil {
		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
