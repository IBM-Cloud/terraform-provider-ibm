// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMIsBareMetalServerInitialization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMISBareMetalServerInitializationCreate,
		ReadContext:   resourceIBMISBareMetalServerInitializationRead,
		DeleteContext: resourceIBMISBareMetalServerInitializationDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{

			isBareMetalServerID: {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Bare metal server identifier",
			},
			isBareMetalServerImage: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The image to be used when provisioning the bare metal server.",
			},
			isBareMetalServerKeys: {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "SSH key Ids for the bare metal server",
			},
			"user_data": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Bare metal server user data to replace initialization",
			},
		},
	}
}

func resourceIBMISBareMetalServerInitializationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var bareMetalServerId, userdata, image string
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}
	if userdataOk, ok := d.GetOk("user_data"); ok {
		userdata = userdataOk.(string)
	}
	if imageOk, ok := d.GetOk("image"); ok {
		image = imageOk.(string)
	}

	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_initialization", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	stopServerIfStartingForInitialization := false
	options := &vpcv1.GetBareMetalServerInitializationOptions{
		ID: &bareMetalServerId,
	}
	stopServerIfStartingForInitialization, err = resourceStopServerIfRunning(bareMetalServerId, "hard", d, context, sess, stopServerIfStartingForInitialization)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceStopServerIfRunning failed: %s", err.Error()), "ibm_is_bare_metal_server_initialization", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	init, _, err := sess.GetBareMetalServerInitializationWithContext(context, options)
	if err != nil || init == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerInitializationWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_initialization", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	d.SetId(bareMetalServerId)

	initializationReplaceOptions := &vpcv1.ReplaceBareMetalServerInitializationOptions{
		ID: &bareMetalServerId,
		Image: &vpcv1.ImageIdentityByID{
			ID: &image,
		},
		UserData: &userdata,
	}
	keySet := d.Get(isBareMetalServerKeys).(*schema.Set)
	if keySet.Len() != 0 {
		keyobjs := make([]vpcv1.KeyIdentityIntf, keySet.Len())
		for i, key := range keySet.List() {
			keystr := key.(string)
			keyobjs[i] = &vpcv1.KeyIdentity{
				ID: &keystr,
			}
		}
		initializationReplaceOptions.Keys = keyobjs
	}
	initInitializationReplace, _, err := sess.ReplaceBareMetalServerInitializationWithContext(context, initializationReplaceOptions)
	if err != nil || initInitializationReplace == nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceBareMetalServerInitializationWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_initialization", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = isWaitForBareMetalServerInitializationStopped(sess, bareMetalServerId, d.Timeout(schema.TimeoutUpdate), d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("isWaitForBareMetalServerInitializationStopped failed: %s", err.Error()), "ibm_is_bare_metal_server_initialization", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if stopServerIfStartingForInitialization {
		_, err = resourceStartServerIfStopped(bareMetalServerId, "hard", d, context, sess, stopServerIfStartingForInitialization)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("resourceStartServerIfStopped failed: %s", err.Error()), "ibm_is_bare_metal_server_initialization", "create")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	diagErr := BareMetalServerInitializationGet(context, d, sess, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}

func BareMetalServerInitializationGet(context context.Context, d *schema.ResourceData, sess *vpcv1.VpcV1, bareMetalServerId string) diag.Diagnostics {

	options := &vpcv1.GetBareMetalServerInitializationOptions{
		ID: &bareMetalServerId,
	}
	init, response, err := sess.GetBareMetalServerInitializationWithContext(context, options)
	if err != nil || init == nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBareMetalServerInitializationWithContext failed: %s", err.Error()), "ibm_is_bare_metal_server_initialization", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set(isBareMetalServerID, bareMetalServerId)
	return nil
}

func resourceIBMISBareMetalServerInitializationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var bareMetalServerId string
	if bmsId, ok := d.GetOk(isBareMetalServerID); ok {
		bareMetalServerId = bmsId.(string)
	}
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_bare_metal_server_initialization", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	diagErr := BareMetalServerInitializationGet(context, d, sess, bareMetalServerId)
	if diagErr != nil {
		return diagErr
	}
	return nil
}
func resourceIBMISBareMetalServerInitializationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")

	return nil
}

func isWaitForBareMetalServerInitializationStopped(client *vpcv1.VpcV1, id string, timeout time.Duration, d *schema.ResourceData) (interface{}, error) {
	log.Printf("Waiting for Bare Metal Server (%s) to be stopped for reload success.", id)
	communicator := make(chan interface{})
	stateConf := &resource.StateChangeConf{
		Pending:    []string{isBareMetalServerStatusPending, isBareMetalServerActionStatusStarting, "reinitializing"},
		Target:     []string{isBareMetalServerStatusRunning, isBareMetalServerStatusFailed, "stopped"},
		Refresh:    isBareMetalServerInitializationRefreshFunc(client, id, d, communicator),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	return stateConf.WaitForState()
}

func isBareMetalServerInitializationRefreshFunc(client *vpcv1.VpcV1, id string, d *schema.ResourceData, communicator chan interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		bmsgetoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &id,
		}
		bms, response, err := client.GetBareMetalServer(bmsgetoptions)
		if err != nil {
			return nil, "", fmt.Errorf("[ERROR] Error getting Bare Metal Server: %s\n%s", err, response)
		}

		select {
		case data := <-communicator:
			return nil, "", data.(error)
		default:
			fmt.Println("no message sent")
		}

		if *bms.Status == "running" || *bms.Status == "failed" {
			// let know the isRestartStartAction() to stop
			close(communicator)
			if *bms.Status == "failed" {
				bmsStatusReason := bms.StatusReasons

				out, err := json.MarshalIndent(bmsStatusReason, "", "    ")
				if err != nil {
					return bms, *bms.Status, fmt.Errorf("[ERROR] The Bare Metal Server (%s) went into failed state during the operation \n [WARNING] Running terraform apply again will remove the tainted bare metal server and attempt to create the bare metal server again replacing the previous configuration", *bms.ID)
				}
				return bms, *bms.Status, fmt.Errorf("[ERROR] Bare Metal Server (%s) went into failed state during the operation \n (%+v) \n [WARNING] Running terraform apply again will remove the tainted Bare Metal Server and attempt to create the Bare Metal Server again replacing the previous configuration", *bms.ID, string(out))
			}
			return bms, *bms.Status, nil

		}
		return bms, *bms.Status, nil
	}
}
