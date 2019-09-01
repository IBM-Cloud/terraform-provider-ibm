package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	st "github.ibm.com/Bluemix/power-go-client/clients/instance"
	"log"
	"time"
)

const (
	PIKeyName = "name"
	PIKey     = "sshkey"
	PIKeyDate = "creationdate"
	PIKeyId   = "keyid"
)

func resourceIBMPIKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMPIKeyCreate,
		Read:   resourceIBMPIKeyRead,
		Update: resourceIBMPIKeyUpdate,
		Delete: resourceIBMPIKeyDelete,
		//Exists:   resourceIBMPIKeyExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			PIKeyId: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			PIKeyName: {
				Type:     schema.TypeString,
				Required: true,
			},

			PIKey: {
				Type:     schema.TypeString,
				Required: true,
			},
			PIKeyDate: {
				Type:     schema.TypeString,
				Optional: true,
			},

			"powerinstanceid": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIBMPIKeyCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}

	powerinstanceid := d.Get("powerinstanceid").(string)
	name := d.Get(PIKeyName).(string)
	sshkey := d.Get(PIKey).(string)
	client := st.NewIBMPIKeyClient(sess, powerinstanceid)

	sshResponse, _, err := client.Create(name, sshkey, powerinstanceid)

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

	return resourceIBMPIKeyRead(d, meta)
}

func resourceIBMPIKeyRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get("powerinstanceid").(string)
	sshkeyC := st.NewIBMPIKeyClient(sess, powerinstanceid)
	sshkeydata, err := sshkeyC.Get(d.Get("name").(string), powerinstanceid)

	if err != nil {
		return err
	}

	d.Set("name", sshkeydata.Name)
	d.Set("sshkey", sshkeydata.SSHKey)
	d.Set("creationdate", sshkeydata.CreationDate)

	return nil

}

func resourceIBMPIKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPIKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPIKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	id := d.Id()
	powerinstanceid := d.Get(IBMPIInstanceId).(string)
	client := st.NewIBMPIKeyClient(sess, powerinstanceid)

	key, err := client.Get(d.Id(), powerinstanceid)
	if err != nil {

		return false, err
	}
	return key.Name == &id, nil
}
