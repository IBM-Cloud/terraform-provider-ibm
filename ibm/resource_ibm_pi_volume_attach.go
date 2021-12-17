// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"log"
	"time"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/helpers"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMPIVolumeAttach() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPIVolumeAttachCreate,
		ReadContext:   resourceIBMPIVolumeAttachRead,
		UpdateContext: resourceIBMPIVolumeAttachUpdate,
		DeleteContext: resourceIBMPIVolumeAttachDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			"volumeattachid": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Optional:    true,
				Description: "Volume attachment ID",
			},

			helpers.PICloudInstanceId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: " Cloud Instance ID - This is the service_instance_id.",
			},

			helpers.PIVolumeAttachName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the volume to attach. Note these  volumes should have been created",
			},

			helpers.PIInstanceName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI Instance name",
			},

			helpers.PIVolumeAttachStatus: {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			helpers.PIVolumeShareable: {
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceIBMPIVolumeAttachCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get(helpers.PIVolumeAttachName).(string)
	servername := d.Get(helpers.PIInstanceName).(string)
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)

	client := st.NewIBMPIVolumeClient(ctx, sess, powerinstanceid)

	volinfo, err := client.Get(name)
	if err != nil {
		return diag.FromErr(err)
	}
	//log.Print("The volume info is %s", volinfo)

	if volinfo.State == "available" || *volinfo.Shareable {
		log.Printf(" In the current state the volume can be attached to the instance ")
	}

	if volinfo.State == "in-use" && *volinfo.Shareable {

		log.Printf("Volume State /Status is  permitted and hence attaching the volume to the instance")
	}

	if volinfo.State == helpers.PIVolumeAllowableAttachStatus && !*volinfo.Shareable {
		return diag.Errorf("the volume cannot be attached in the current state. The volume must be in the *available* state. No other states are permissible")
	}

	err = client.Attach(servername, name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*volinfo.VolumeID)
	if err != nil {
		log.Printf("[DEBUG]  err %s", err)
		return diag.FromErr(err)
	}

	_, err = isWaitForIBMPIVolumeAttachAvailable(ctx, client, d.Id(), powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}
	//return nil
	return resourceIBMPIVolumeAttachRead(ctx, d, meta)
}

func resourceIBMPIVolumeAttachRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	servername := d.Get(helpers.PIInstanceName).(string)

	client := st.NewIBMPIVolumeClient(ctx, sess, powerinstanceid)

	vol, err := client.CheckVolumeAttach(servername, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	//d.SetId(vol.ID.String())
	d.Set(helpers.PIVolumeAttachName, vol.Name)
	d.Set(helpers.PIVolumeSize, vol.Size)
	d.Set(helpers.PIVolumeShareable, vol.Shareable)
	return nil
}

func resourceIBMPIVolumeAttachUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	client := st.NewIBMPIVolumeClient(ctx, sess, powerinstanceid)

	name := ""
	if d.HasChange(helpers.PIVolumeAttachName) {
		name = d.Get(helpers.PIVolumeAttachName).(string)
	}

	size := float64(d.Get(helpers.PIVolumeSize).(float64))
	shareable := bool(d.Get(helpers.PIVolumeShareable).(bool))

	body := &models.UpdateVolume{
		Name:      &name,
		Size:      size,
		Shareable: &shareable,
	}
	volrequest, err := client.UpdateVolume(d.Id(), body)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = isWaitForIBMPIVolumeAttachAvailable(ctx, client, *volrequest.VolumeID, powerinstanceid, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMPIVolumeRead(ctx, d, meta)
}

func resourceIBMPIVolumeAttachDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, _ := meta.(ClientSession).IBMPISession()
	powerinstanceid := d.Get(helpers.PICloudInstanceId).(string)
	name := d.Get(helpers.PIVolumeAttachName).(string)
	servername := d.Get(helpers.PIInstanceName).(string)
	client := st.NewIBMPIVolumeClient(ctx, sess, powerinstanceid)

	log.Printf("the id of the volume to detach is%s ", d.Id())
	err := client.Detach(servername, name)
	if err != nil {
		return diag.FromErr(err)
	}

	// wait for power volume states to be back as available. if it's attached it will be in-use
	d.SetId("")
	return nil
}

func isWaitForIBMPIVolumeAttachAvailable(ctx context.Context, client *st.IBMPIVolumeClient, id, powerinstanceid string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for Volume (%s) to be available for attachment", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", helpers.PIVolumeProvisioning},
		Target:     []string{helpers.PIVolumeAllowableAttachStatus},
		Refresh:    isIBMPIVolumeAttachRefreshFunc(client, id, powerinstanceid),
		Delay:      10 * time.Second,
		MinTimeout: 2 * time.Minute,
		Timeout:    10 * time.Minute,
	}

	return stateConf.WaitForStateContext(ctx)
}

func isIBMPIVolumeAttachRefreshFunc(client *st.IBMPIVolumeClient, id, powerinstanceid string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		vol, err := client.Get(id)
		if err != nil {
			return nil, "", err
		}

		if vol.State == "in-use" {
			return vol, helpers.PIVolumeAllowableAttachStatus, nil
		}

		return vol, helpers.PIVolumeProvisioning, nil
	}
}
