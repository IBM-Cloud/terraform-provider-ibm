package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"time"
)

const (
	PowerSSHKeyName = "name"
	PowerSSHKey     = "sshkey"
	PowerSSHKeyDate = "creationdate"
	PowerSSHKeyId   = "keyid"
)

func resourceIBMPowerSSHKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPowerSSHKeyCreate,
		Read:   resourceIBMPowerSSHKeyRead,
		Update: resourceIBMPowerSSHKeyUpdate,
		Delete: resourceIBMPowerSSHKeyDelete,
		//Exists:   resourceIBMPowerSSHKeyExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			PowerSSHKeyId: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			PowerSSHKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},

			PowerSSHKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			PowerSSHKeyDate: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIBMPowerSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	name := d.Get(PowerSSHKeyName).(string)
	sshkey := d.Get(PowerSSHKey).(string)
	//createdate := d.Get(PowerSSHKeyDate).(strfmt.DateTime)

	client := st.NewPowerSSHKeyClient(sess)

	sshResponse, _, err := client.Create(name, sshkey)

	if err != nil {
		return err
	}

	log.Printf("Printing the sshkey %+v", &sshResponse)

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	return resourceIBMPowerSSHKeyRead(d, meta)
}

func resourceIBMPowerSSHKeyRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	sshkeyC := st.NewPowerSSHKeyClient(sess)
	sshkeydata, err := sshkeyC.Get(d.Get("name").(string))

	if err != nil {
		return err
	}

	d.Set("name", sshkeydata.Name)
	d.Set("sshkey", sshkeydata.SSHKey)
	d.Set("creationdate", sshkeydata.CreationDate)

	return nil

}

func resourceIBMPowerSSHKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPowerSSHKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}
