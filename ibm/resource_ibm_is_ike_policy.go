package ibm

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/vpn"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isIKEName              = "name"
	isIKEAuthenticationAlg = "authentication_algorithm"
	isIKEEncryptionAlg     = "encryption_algorithm"
	isIKEDhGroup           = "dh_group"
	isIKEVERSION           = "ike_version"
	isIKEKeyLifeTime       = "key_lifetime"
	isIKEResourceGroup     = "resource_group"
	isIKENegotiationMode   = "negotiation_mode"
	isIKEVPNConnections    = "vpn_connections"
	isIKEVPNConnectionName = "name"
	isIKEVPNConnectionId   = "id"
	isIKEVPNConnectionHref = "href"
	isIKEHref              = "href"
)

func resourceIBMISIKEPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISIKEPolicyCreate,
		Read:     resourceIBMISIKEPolicyRead,
		Update:   resourceIBMISIKEPolicyUpdate,
		Delete:   resourceIBMISIKEPolicyDelete,
		Exists:   resourceIBMISIKEPolicyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			isIKEName: {
				Type:     schema.TypeString,
				Required: true,
			},

			isIKEAuthenticationAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"md5", "sha1", "sha256"}),
			},

			isIKEEncryptionAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"3des", "aes128", "aes256"}),
			},

			isIKEDhGroup: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validateAllowedIntValue([]int{2, 5, 14}),
			},

			isIKEResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isIKEKeyLifeTime: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      28800,
				ValidateFunc: validateKeyLifeTime,
			},

			isIKEVERSION: {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateAllowedIntValue([]int{1, 2}),
			},

			isIKENegotiationMode: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isIKEHref: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isIKEVPNConnections: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						isIKEVPNConnectionName: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIKEVPNConnectionId: {
							Type:     schema.TypeString,
							Computed: true,
						},
						isIKEVPNConnectionHref: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIBMISIKEPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] IKE Policy create")
	name := d.Get(isIKEName).(string)
	authenticationAlg := d.Get(isIKEAuthenticationAlg).(string)
	encryptionAlg := d.Get(isIKEEncryptionAlg).(string)
	dhGroup := d.Get(isIKEDhGroup).(int)
	ikeVersion := d.Get(isIKEVERSION).(int)
	resourceGrpId := d.Get(isIKEResourceGroup).(string)
	keyLifetime := d.Get(isIKEKeyLifeTime).(int)
	//Send an empty array for tags as we are not able to read the tags back
	tags := []string{}

	vpnC := vpn.NewVpnClient(sess)
	ike, err := vpnC.CreateIkePolicy(authenticationAlg, encryptionAlg, name, resourceGrpId, tags, dhGroup, ikeVersion, keyLifetime)
	if err != nil {
		log.Printf("[DEBUG] ike policy err %s", err)
		return err
	}

	d.SetId(ike.ID.String())
	log.Printf("[INFO] IKE : %s", ike.ID.String())
	return resourceIBMISIKEPolicyRead(d, meta)
}

func resourceIBMISIKEPolicyRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpnC := vpn.NewVpnClient(sess)

	ike, err := vpnC.GetIkePolicy(d.Id())
	if err != nil {
		return err
	}

	d.Set(isIKEName, ike.Name)
	d.Set(isIKEAuthenticationAlg, ike.AuthenticationAlgorithm)
	d.Set(isIKEEncryptionAlg, ike.EncryptionAlgorithm)
	if ike.ResourceGroup != nil {
		d.Set(isIKEResourceGroup, ike.ResourceGroup.ID)
	} else {
		d.Set(isIKEResourceGroup, nil)
	}
	if ike.KeyLifetime != 0 {
		d.Set(isIKEKeyLifeTime, ike.KeyLifetime)
	}
	d.Set(isIKEHref, ike.Href)
	d.Set(isIKENegotiationMode, ike.NegotiationMode)
	d.Set(isIKEVERSION, ike.IkeVersion)
	d.Set(isIKEDhGroup, ike.DhGroup)
	connList := make([]map[string]interface{}, 0)
	if ike.Connections != nil && len(ike.Connections) > 0 {
		for _, connection := range ike.Connections {
			conn := map[string]interface{}{}
			conn[isIKEVPNConnectionName] = connection.Name
			conn[isIKEVPNConnectionId] = connection.ID.String()
			conn[isIKEVPNConnectionHref] = connection.Href
			connList = append(connList, conn)
		}
	}
	d.Set(isIKEVPNConnections, connList)
	return nil
}

func resourceIBMISIKEPolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpnC := vpn.NewVpnClient(sess)

	if d.HasChange(isIKEName) || d.HasChange(isIKEAuthenticationAlg) || d.HasChange(isIKEEncryptionAlg) || d.HasChange(isIKEDhGroup) || d.HasChange(isIKEVERSION) || d.HasChange(isIKEKeyLifeTime) {
		name := d.Get(isIKEName).(string)
		authenticationAlg := d.Get(isIKEAuthenticationAlg).(string)
		encryptionAlg := d.Get(isIKEEncryptionAlg).(string)
		keyLifetime := d.Get(isIKEKeyLifeTime).(int)
		dhGroup := d.Get(isIKEDhGroup).(int)
		ikeVersion := d.Get(isIKEVERSION).(int)
		_, err := vpnC.UpdateIkePolicy(d.Id(), authenticationAlg, encryptionAlg, name, dhGroup, ikeVersion, keyLifetime)
		if err != nil {
			return err
		}
	}

	return resourceIBMISIKEPolicyRead(d, meta)
}

func resourceIBMISIKEPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpnC := vpn.NewVpnClient(sess)
	err = vpnC.DeleteIkePolicy(d.Id())
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISIKEPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	vpnC := vpn.NewVpnClient(sess)

	_, err = vpnC.GetIkePolicy(d.Id())
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
