package ibm

import (
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/schema"
	//"fmt"
	"github.com/hashicorp/terraform/helper/validation"
	"github.ibm.com/Bluemix/power-go-client/clients/instance"
)

func dataSourceIBMPowerSSHKey() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMPowerSSHKeysRead,
		Schema: map[string]*schema.Schema{

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "SSHKey Name to be used for pvminstances",
				ValidateFunc: validation.NoZeroValues,
			},

			"creationdate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sshkey": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIBMPowerSSHKeysRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).PowerSession()
	if err != nil {
		return err
	}

	sshkeyC := instance.NewPowerSSHKeyClient(sess)
	sshkeydata, err := sshkeyC.Get(d.Get("name").(string))

	if err != nil {
		return err
	}

	var clientgenU, _ = uuid.GenerateUUID()
	d.SetId(clientgenU)
	d.Set("name", sshkeydata.Name)
	d.Set("sshkey", sshkeydata.SSHKey)
	d.Set("creationdate", sshkeydata.CreationDate)

	return nil
	//return fmt.Errorf("No Image found with name %s", imagedata.)

}
