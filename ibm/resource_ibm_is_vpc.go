package ibm

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.ibm.com/Bluemix/riaas-go-client/clients/network"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
)

const (
	isVPCDefaultNetworkACL     = "default_network_acl"
	isVPCIsDefault             = "is_default"
	isVPCIDefaultSecurityGroup = "default_security_group"
	isVPCName                  = "name"
	isVPCResourceGroup         = "resource_group"
	isVPCStatus                = "status"
	isVPCDeleting              = "deleting"
	isVPCDeleted               = "done"
	isVPCTags                  = "tags"
	isVPCClassicAccess         = "classic_access"
)

func resourceIBMISVPC() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVPCCreate,
		Read:     resourceIBMISVPCRead,
		Update:   resourceIBMISVPCUpdate,
		Delete:   resourceIBMISVPCDelete,
		Exists:   resourceIBMISVPCExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			isVPCDefaultNetworkACL: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
				Computed: true,
			},

			isVPCIsDefault: {
				Type:     schema.TypeBool,
				ForceNew: true,
				Default:  false,
				Optional: true,
				Removed:  "This field is removed use classic_access",
			},

			isVPCClassicAccess: {
				Type:     schema.TypeBool,
				ForceNew: true,
				Default:  false,
				Optional: true,
			},

			isVPCName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			isVPCResourceGroup: {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			isVPCStatus: {
				Type:     schema.TypeString,
				Computed: true,
			},

			isVPCIDefaultSecurityGroup: {
				Type:     schema.TypeString,
				Computed: true,
			},
			isVPCTags: {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      resourceIBMVPCHash,
			},
		},
	}
}

func resourceIBMISVPCCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] VPC create")
	name := d.Get(isVPCName).(string)
	isClassic := d.Get(isVPCClassicAccess).(bool)
	nwacl := d.Get(isVPCDefaultNetworkACL).(string)
	var rg string

	if grp, ok := d.GetOk(isVPCResourceGroup); ok {
		rg = grp.(string)
	}

	vpcC := network.NewVPCClient(sess)
	vpc, err := vpcC.Create(name, isClassic, nwacl, rg)
	if err != nil {
		log.Printf("[DEBUG] VPC err %s", isErrorToString(err))
		return err
	}

	d.SetId(vpc.ID.String())
	log.Printf("[INFO] VPC : %s", vpc.ID.String())

	if _, ok := d.GetOk(isVPCTags); ok {
		oldList, newList := d.GetChange(isVPCTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, vpc.Crn)
		if err != nil {
			return fmt.Errorf(
				"Error on create of resource vpc (%s) tags: %s", d.Id(), err)
		}
	}
	return resourceIBMISVPCRead(d, meta)
}

func resourceIBMISVPCRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcC := network.NewVPCClient(sess)

	vpc, err := vpcC.Get(d.Id())
	if err != nil {
		return err
	}

	d.Set("id", vpc.ID.String())
	d.Set(isVPCName, vpc.Name)
	d.Set(isVPCClassicAccess, vpc.ClassicAccess)
	d.Set(isVPCStatus, vpc.Status)
	if vpc.DefaultNetworkACL != nil {
		log.Printf("[DEBUG] vpc default network acl is not null :%s", vpc.DefaultNetworkACL.ID)
		d.Set(isVPCDefaultNetworkACL, vpc.DefaultNetworkACL.ID)
	} else {
		log.Printf("[DEBUG] vpc default network acl is  null")
		d.Set(isVPCDefaultNetworkACL, nil)
	}
	if vpc.DefaultSecurityGroup != nil {
		d.Set(isVPCIDefaultSecurityGroup, vpc.DefaultSecurityGroup.ID)
	} else {
		d.Set(isVPCIDefaultSecurityGroup, nil)
	}
	tags, err := GetTagsUsingCRN(meta, vpc.Crn)
	if err != nil {
		return fmt.Errorf(
			"Error on get of resource vpc (%s) tags: %s", d.Id(), err)
	}
	d.Set(isVPCTags, tags)
	d.Set(isVPCResourceGroup, vpc.ResourceGroup.ID)
	return nil
}

func resourceIBMISVPCUpdate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcC := network.NewVPCClient(sess)

	vpc, err := vpcC.Get(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange(isVPCName) {
		name := d.Get(isVPCName).(string)
		_, err := vpcC.Update(d.Id(), name)
		if err != nil {
			return err
		}
	}

	if d.HasChange(isVPCTags) {
		oldList, newList := d.GetChange(isVPCTags)
		err = UpdateTagsUsingCRN(oldList, newList, meta, vpc.Crn)
		if err != nil {
			return fmt.Errorf(
				"Error on create of resource vpc (%s) tags: %s", d.Id(), err)
		}

	}

	return resourceIBMISVPCRead(d, meta)
}

func resourceIBMISVPCDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}
	vpcC := network.NewVPCClient(sess)
	err = vpcC.Delete(d.Id())
	if err != nil {
		return err
	}

	_, err = isWaitForVPCDeleted(vpcC, d.Id(), d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func isWaitForVPCDeleted(vpc *network.VPCClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for VPC (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isVPCDeleting},
		Target:     []string{},
		Refresh:    isVPCDeleteRefreshFunc(vpc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isVPCDeleteRefreshFunc(vpc *network.VPCClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] delete function here")
		VPC, err := vpc.Get(id)
		if err == nil {
			return VPC, isVPCDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "not_found" {
				log.Printf("[DEBUG] returning deleted")
				return nil, isVPCDeleted, nil
			}
		}
		log.Printf("[DEBUG] returning x")
		return nil, isVPCDeleting, err
	}
}

func resourceIBMISVPCExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return false, err
	}
	vpcC := network.NewVPCClient(sess)

	_, err = vpcC.Get(d.Id())
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

func isErrorToString(err error) string {
	iserror, ok := err.(iserrors.RiaasError)
	if ok {
		log.Printf("[DEBUG] Hit Riaas Error")
		retmsg := ""

		for _, e := range iserror.Payload.Errors {
			retmsg = retmsg + "\n" + e.Message + "\n" + e.Code + "\n" + e.MoreInfo + "\n" + e.Target.Name + "\n" + e.Target.Type
		}
		return retmsg
	}
	return err.Error()
}

func resourceIBMVPCHash(v interface{}) int {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s",
		strings.ToLower(v.(string))))
	return hashcode.String(buf.String())
}
