package ibm

import (
	"fmt"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"time"
)

func resourceIBMPIImage() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPIImageCreate,
		Read:     resourceIBMPIImageRead,
		Update:   resourceIBMPIImageUpdate,
		Delete:   resourceIBMPIImageDelete,
		Exists:   resourceIBMPIImageExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			helpers.PIKeyId: {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},

			helpers.PIImageName: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PIInstanceImageName: {
				Type:     schema.TypeString,
				Required: true,
			},

			helpers.PICloudInstanceId: {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed Attribute

			"imageid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMPIImageCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		log.Printf("Failed to get the session")
		return err
	}

	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PIImageName).(string)
	imageid := d.Get(helpers.PIInstanceImageName).(string)

	client := st.NewIBMPIImageClient(sess, powerinstanceid)

	imageResponse, err := client.Create(name, imageid, powerinstanceid)

	if err != nil {
		return err
	}

	log.Printf("Printing the image post response %+v", &imageResponse)

	IBMPIImageID := imageResponse.ImageID
	log.Printf("the imageid from the post call is..%s", IBMPIImageID)

	d.SetId(*IBMPIImageID)

	log.Printf("the Image id from the post is %s", *IBMPIImageID)
	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}
	_, err = isWaitForIBMPIImageAvailable(client, d.Id(), d.Timeout(schema.TimeoutCreate), powerinstanceid)
	if err != nil {
		return err
	}

	if err != nil {
		log.Printf("[DEBUG]  err %s", isErrorToString(err))
		return err
	}

	return resourceIBMPIImageRead(d, meta)
}

func resourceIBMPIImageRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	imageC := st.NewIBMPIImageClient(sess, powerinstanceid)
	imagedata, err := imageC.Get(d.Get(helpers.PIImageName).(string), powerinstanceid)

	if err != nil {
		return err
	}

	imageid := *imagedata.ImageID
	d.SetId(imageid)

	return nil

}

func resourceIBMPIImageUpdate(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPIImageDelete(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMPIImageExists(d *schema.ResourceData, meta interface{}) (bool, error) {

	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return false, err
	}
	//id := d.Id()
	name := d.Get(helpers.PIImageName)
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIImageClient(sess, powerinstanceid)

	image, err := client.Get(d.Get(helpers.PIImageName).(string), powerinstanceid)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return image.Name == name, nil
}

func isWaitForIBMPIImageAvailable(client *st.IBMPIImageClient, id string, timeout time.Duration, powerinstanceid string) (interface{}, error) {
	log.Printf("Waiting for Power Image (%s) to be available.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PIImageQueStatus},
		Target:     []string{helpers.PIImageActiveStatus},
		Refresh:    isIBMPIImageRefreshFunc(client, id, powerinstanceid),
		Timeout:    timeout,
		Delay:      20 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isIBMPIImageRefreshFunc(client *st.IBMPIImageClient, id, powerinstanceid string) resource.StateRefreshFunc {

	log.Printf("Calling the isIBMPIImageRefreshFunc Refresh Function....")
	return func() (interface{}, string, error) {
		image, err := client.Get(id, powerinstanceid)
		if err != nil {
			return nil, "", err
		}

		if image.State == "active" {

			return image, helpers.PIImageActiveStatus, nil
		}

		return image, helpers.PIImageQueStatus, nil
	}
}
