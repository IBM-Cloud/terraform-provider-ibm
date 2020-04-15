package ibm

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateISName,
			},

			isIKEAuthenticationAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_ike_policy", isIKEAuthenticationAlg),
			},

			isIKEEncryptionAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_ike_policy", isIKEEncryptionAlg),
			},

			isIKEDhGroup: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: InvokeValidator("ibm_is_ike_policy", isIKEDhGroup),
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
				ValidateFunc: InvokeValidator("ibm_is_ike_policy", isIKEVERSION),
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

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},
	}
}

func resourceIBMISIKEValidator() *ResourceValidator {

	validateSchema := make([]ValidateSchema, 1)
	authentication_algorithm := "md5, sha1, sha256"
	encryption_algorithm := "3des, aes128, aes256"
	dh_group := "2, 5, 14"
	ike_version := "1, 2"

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isIKEAuthenticationAlg,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              authentication_algorithm})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isIKEEncryptionAlg,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              encryption_algorithm})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isIKEDhGroup,
			ValidateFunctionIdentifier: ValidateAllowedIntValue,
			Type:                       TypeInt,
			Required:                   true,
			AllowedValues:              dh_group})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 isIKEVERSION,
			ValidateFunctionIdentifier: ValidateAllowedIntValue,
			Type:                       TypeInt,
			Optional:                   true,
			AllowedValues:              ike_version})

	ibmISIKEResourceValidator := ResourceValidator{ResourceName: "ibm_is_ike_policy", Schema: validateSchema}
	return &ibmISIKEResourceValidator
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

	vpnC := vpn.NewVpnClient(sess)
	ike, err := vpnC.CreateIkePolicy(authenticationAlg, encryptionAlg, name, resourceGrpId, dhGroup, ikeVersion, keyLifetime)
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
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	if sess.Generation == 1 {
		d.Set(ResourceControllerURL, controller+"/vpc/network/ikepolicies")
	} else {
		d.Set(ResourceControllerURL, controller+"/vpc-ext/network/ikepolicies")
	}
	d.Set(ResourceName, ike.Name)
	if ike.ResourceGroup != nil {
		rsMangClient, err := meta.(ClientSession).ResourceManagementAPIv2()
		if err != nil {
			return err
		}
		grp, err := rsMangClient.ResourceGroup().Get(ike.ResourceGroup.ID.String())
		if err != nil {
			return err
		}
		d.Set(ResourceGroupName, grp.Name)
	}
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
				iserror.Payload.Errors[0].Code == "ike_policy_not_found" {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}
