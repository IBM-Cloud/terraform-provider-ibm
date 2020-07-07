package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateISName,
				Description:  "IPSEC name",
			},

			isIpSecAuthenticationAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"md5", "sha1", "sha256"}),
				Description:  "Authentication alorothm",
			},

			isIpSecEncryptionAlg: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"triple_des", "aes128", "aes256"}),
				Description:  "Encryption algorithm",
			},

			isIpSecPFS: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue([]string{"disabled", "group_2", "group_5", "group_14"}),
				Description:  "PFS info",
			},

			isIPSecResourceGroup: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "Resource group info",
			},

			isIpSecKeyLifeTime: {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      3600,
				ValidateFunc: validateKeyLifeTime,
				Description:  "IPSEC key lifetime",
			},

			isIPSecEncapsulationMode: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPSEC encapsulation mode",
			},

			isIPSecTransformProtocol: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPSEC transform protocol",
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

func resourceIBMISIPSecPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Ip Sec create")
	name := d.Get(isIpSecName).(string)
	authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
	encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
	pfs := d.Get(isIpSecPFS).(string)

	if userDetails.generation == 1 {
		err := classicIpsecpCreate(d, meta, authenticationAlg, encryptionAlg, name, pfs)
		if err != nil {
			return err
		}
	} else {
		err := ipsecpCreate(d, meta, authenticationAlg, encryptionAlg, name, pfs)
		if err != nil {
			return err
		}
	}
	return resourceIBMISIPSecPolicyRead(d, meta)
}

func classicIpsecpCreate(d *schema.ResourceData, meta interface{}, authenticationAlg, encryptionAlg, name, pfs string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcclassicv1.CreateIpsecPolicyOptions{
		AuthenticationAlgorithm: &authenticationAlg,
		EncryptionAlgorithm:     &encryptionAlg,
		Pfs:                     &pfs,
		Name:                    &name,
	}

	if keylt, ok := d.GetOk(isIpSecKeyLifeTime); ok {
		keyLifetime := int64(keylt.(int))
		options.KeyLifetime = &keyLifetime
	} else {
		keyLifetime := int64(3600)
		options.KeyLifetime = &keyLifetime
	}

	if rgrp, ok := d.GetOk(isIPSecResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcclassicv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	ipSec, response, err := sess.CreateIpsecPolicy(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] ipSec policy err %s\n%s", err, response)
	}
	d.SetId(*ipSec.ID)
	log.Printf("[INFO] ipSec policy : %s", *ipSec.ID)
	return nil
}

func ipsecpCreate(d *schema.ResourceData, meta interface{}, authenticationAlg, encryptionAlg, name, pfs string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	options := &vpcv1.CreateIpsecPolicyOptions{
		AuthenticationAlgorithm: &authenticationAlg,
		EncryptionAlgorithm:     &encryptionAlg,
		Pfs:                     &pfs,
		Name:                    &name,
	}

	if keylt, ok := d.GetOk(isIpSecKeyLifeTime); ok {
		keyLifetime := int64(keylt.(int))
		options.KeyLifetime = &keyLifetime
	} else {
		keyLifetime := int64(3600)
		options.KeyLifetime = &keyLifetime
	}

	if rgrp, ok := d.GetOk(isIPSecResourceGroup); ok {
		rg := rgrp.(string)
		options.ResourceGroup = &vpcv1.ResourceGroupIdentity{
			ID: &rg,
		}
	}
	ipSec, response, err := sess.CreateIpsecPolicy(options)
	if err != nil {
		return fmt.Errorf("[DEBUG] ipSec policy err %s\n%s", err, response)
	}
	d.SetId(*ipSec.ID)
	log.Printf("[INFO] ipSec policy : %s", *ipSec.ID)
	return nil
}

func resourceIBMISIPSecPolicyRead(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicIpsecpGet(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := ipsecpGet(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicIpsecpGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getIpsecPolicyOptions := &vpcclassicv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	ipSec, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting IPSEC Policy(%s): %s\n%s", id, err, response)
	}

	d.Set(isIpSecName, *ipSec.Name)
	d.Set(isIpSecAuthenticationAlg, *ipSec.AuthenticationAlgorithm)
	d.Set(isIpSecEncryptionAlg, *ipSec.EncryptionAlgorithm)
	if ipSec.ResourceGroup != nil {
		d.Set(isIPSecResourceGroup, *ipSec.ResourceGroup.ID)
	} else {
		d.Set(isIPSecResourceGroup, nil)
	}
	d.Set(isIpSecPFS, *ipSec.Pfs)
	if ipSec.KeyLifetime != nil {
		d.Set(isIpSecKeyLifeTime, *ipSec.KeyLifetime)
	}
	d.Set(isIPSecEncapsulationMode, *ipSec.EncapsulationMode)
	d.Set(isIPSecTransformProtocol, *ipSec.TransformProtocol)

	connList := make([]map[string]interface{}, 0)
	if ipSec.Connections != nil && len(ipSec.Connections) > 0 {
		for _, connection := range ipSec.Connections {
			conn := map[string]interface{}{}
			conn[isIPSecVPNConnectionName] = *connection.Name
			conn[isIPSecVPNConnectionId] = *connection.ID
			conn[isIPSecVPNConnectionHref] = *connection.Href
			connList = append(connList, conn)
		}
	}
	d.Set(isIPSecVPNConnections, connList)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc/network/ipsecpolicies")
	d.Set(ResourceName, *ipSec.Name)
	// d.Set(ResourceCRN, *ipSec.Crn)
	if ipSec.ResourceGroup != nil {
		rsMangClient, err := meta.(ClientSession).ResourceManagementAPIv2()
		if err != nil {
			return err
		}
		grp, err := rsMangClient.ResourceGroup().Get(*ipSec.ResourceGroup.ID)
		if err != nil {
			return err
		}
		d.Set(ResourceGroupName, grp.Name)
	}
	return nil
}

func ipsecpGet(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	ipSec, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting IPSEC Policy(%s): %s\n%s", id, err, response)
	}
	d.Set(isIpSecName, *ipSec.Name)
	d.Set(isIpSecAuthenticationAlg, *ipSec.AuthenticationAlgorithm)
	d.Set(isIpSecEncryptionAlg, *ipSec.EncryptionAlgorithm)
	if ipSec.ResourceGroup != nil {
		d.Set(isIPSecResourceGroup, *ipSec.ResourceGroup.ID)
		d.Set(ResourceGroupName, *ipSec.ResourceGroup.Name)
	} else {
		d.Set(isIPSecResourceGroup, nil)
	}
	d.Set(isIpSecPFS, *ipSec.Pfs)
	if ipSec.KeyLifetime != nil {
		d.Set(isIpSecKeyLifeTime, *ipSec.KeyLifetime)
	}
	d.Set(isIPSecEncapsulationMode, *ipSec.EncapsulationMode)
	d.Set(isIPSecTransformProtocol, *ipSec.TransformProtocol)

	connList := make([]map[string]interface{}, 0)
	if ipSec.Connections != nil && len(ipSec.Connections) > 0 {
		for _, connection := range ipSec.Connections {
			conn := map[string]interface{}{}
			conn[isIPSecVPNConnectionName] = *connection.Name
			conn[isIPSecVPNConnectionId] = *connection.ID
			conn[isIPSecVPNConnectionHref] = *connection.Href
			connList = append(connList, conn)
		}
	}
	d.Set(isIPSecVPNConnections, connList)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/vpc-ext/network/ipsecpolicies")
	d.Set(ResourceName, *ipSec.Name)
	// d.Set(ResourceCRN, *ipSec.Crn)
	return nil
}

func resourceIBMISIPSecPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicIpsecpUpdate(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := ipsecpUpdate(d, meta, id)
		if err != nil {
			return err
		}
	}
	return resourceIBMISIPSecPolicyRead(d, meta)
}

func classicIpsecpUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}

	options := &vpcclassicv1.UpdateIpsecPolicyOptions{
		ID: &id,
	}
	if d.HasChange(isIpSecName) || d.HasChange(isIpSecAuthenticationAlg) || d.HasChange(isIpSecEncryptionAlg) || d.HasChange(isIpSecPFS) || d.HasChange(isIpSecKeyLifeTime) {
		name := d.Get(isIpSecName).(string)
		authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
		encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
		pfs := d.Get(isIpSecPFS).(string)
		keyLifetime := int64(d.Get(isIpSecKeyLifeTime).(int))

		options.Name = &name
		options.AuthenticationAlgorithm = &authenticationAlg
		options.EncryptionAlgorithm = &encryptionAlg
		options.Pfs = &pfs
		options.KeyLifetime = &keyLifetime

		_, response, err := sess.UpdateIpsecPolicy(options)
		if err != nil {
			return fmt.Errorf("Error on update of IPSEC Policy(%s): %s\n%s", id, err, response)
		}
	}
	return nil
}

func ipsecpUpdate(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	options := &vpcv1.UpdateIpsecPolicyOptions{
		ID: &id,
	}
	if d.HasChange(isIpSecName) || d.HasChange(isIpSecAuthenticationAlg) || d.HasChange(isIpSecEncryptionAlg) || d.HasChange(isIpSecPFS) || d.HasChange(isIpSecKeyLifeTime) {
		name := d.Get(isIpSecName).(string)
		authenticationAlg := d.Get(isIpSecAuthenticationAlg).(string)
		encryptionAlg := d.Get(isIpSecEncryptionAlg).(string)
		pfs := d.Get(isIpSecPFS).(string)
		keyLifetime := int64(d.Get(isIpSecKeyLifeTime).(int))

		options.Name = &name
		options.AuthenticationAlgorithm = &authenticationAlg
		options.EncryptionAlgorithm = &encryptionAlg
		options.Pfs = &pfs
		options.KeyLifetime = &keyLifetime

		_, response, err := sess.UpdateIpsecPolicy(options)
		if err != nil {
			return fmt.Errorf("Error on update of IPSEC Policy(%s): %s\n%s", id, err, response)
		}
	}
	return nil
}

func resourceIBMISIPSecPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	id := d.Id()
	if userDetails.generation == 1 {
		err := classicIpsecpDelete(d, meta, id)
		if err != nil {
			return err
		}
	} else {
		err := ipsecpDelete(d, meta, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func classicIpsecpDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return err
	}
	getIpsecPolicyOptions := &vpcclassicv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting IPSEC Policy(%s): %s\n%s", id, err, response)
	}

	deleteIpsecPolicyOptions := &vpcclassicv1.DeleteIpsecPolicyOptions{
		ID: &id,
	}
	response, err = sess.DeleteIpsecPolicy(deleteIpsecPolicyOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting IPSEC Policy(%s): %s\n%s", id, err, response)
	}
	d.SetId("")
	return nil
}

func ipsecpDelete(d *schema.ResourceData, meta interface{}, id string) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}
	getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error getting IPSEC Policy(%s): %s\n%s", id, err, response)
	}
	deleteIpsecPolicyOptions := &vpcv1.DeleteIpsecPolicyOptions{
		ID: &id,
	}
	response, err = sess.DeleteIpsecPolicy(deleteIpsecPolicyOptions)
	if err != nil {
		return fmt.Errorf("Error Deleting IPSEC Policy(%s): %s\n%s", id, err, response)
	}
	d.SetId("")
	return nil
}

func resourceIBMISIPSecPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	id := d.Id()

	if userDetails.generation == 1 {
		exists, err := classicIpsecpExists(d, meta, id)
		return exists, err
	} else {
		exists, err := ipsecpExists(d, meta, id)
		return exists, err
	}
}

func classicIpsecpExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := classicVpcClient(meta)
	if err != nil {
		return false, err
	}
	getIpsecPolicyOptions := &vpcclassicv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting IPSEC Policy(%s): %s\n%s", id, err, response)
	}
	return true, nil
}

func ipsecpExists(d *schema.ResourceData, meta interface{}, id string) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}
	getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
		ID: &id,
	}
	_, response, err := sess.GetIpsecPolicy(getIpsecPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error getting IPSEC Policy(%s): %s\n%s", id, err, response)
	}
	return true, nil
}
