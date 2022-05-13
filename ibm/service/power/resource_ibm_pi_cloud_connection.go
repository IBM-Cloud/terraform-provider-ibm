// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

func ResourceIBMPICloudConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPICloudConnectionCreate,
		ReadContext:   resourceIBMPICloudConnectionRead,
		UpdateContext: resourceIBMPICloudConnectionUpdate,
		DeleteContext: resourceIBMPICloudConnectionDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			// Required Attributes
			Arg_CloudInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PI cloud instance ID",
			},
			Arg_CloudConnectionName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the cloud connection",
			},
			Arg_CloudConnectionSpeed: {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validate.ValidateAllowedIntValues([]int{50, 100, 200, 500, 1000, 2000, 5000, 10000}),
				Description:  "Speed of the cloud connection (speed in megabits per second)",
			},

			// Optional Attributes
			Arg_CloudConnectionRouting: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable global routing for this cloud connection",
			},
			Arg_CloudConnectionMetered: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable metered for this cloud connection",
			},
			Arg_CloudConnectionNetworks: {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Set of Networks to attach to this cloud connection",
			},
			Arg_CloudConnectionClassic: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable classic endpoint destination",
			},
			Arg_CloudConnectionGreCIDR: {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{Arg_CloudConnectionClassic, Arg_CloudConnectionGreDest},
				Description:  "GRE network in CIDR notation",
			},
			Arg_CloudConnectionGreDest: {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{Arg_CloudConnectionClassic, Arg_CloudConnectionGreCIDR},
				Description:  "GRE destination IP address",
			},
			Arg_CloudConnectionVPC: {
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      false,
				RequiredWith: []string{Arg_CloudConnectionVPCCrns},
				Description:  "Enable VPC for this cloud connection",
			},
			Arg_CloudConnectionVPCCrns: {
				Type:         schema.TypeSet,
				Optional:     true,
				Elem:         &schema.Schema{Type: schema.TypeString},
				RequiredWith: []string{Arg_CloudConnectionVPC},
				Description:  "Set of VPCs to attach to this cloud connection",
			},

			//Computed Attributes
			Attr_CloudConnectionID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud connection ID",
			},
			Attr_CloudConnectionStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link status",
			},
			Attr_CloudConnectionIbmIP: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IBM IP address",
			},
			Attr_CloudConnectionUserIP: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User IP address",
			},
			Attr_CloudConnectionPort: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Port",
			},
			Attr_CloudConnectionSourceGreAddr: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "GRE auto-assigned source IP address",
			},
		},
	}
}

func resourceIBMPICloudConnectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := d.Get(Arg_CloudInstanceID).(string)
	name := d.Get(Arg_CloudConnectionName).(string)
	speed := int64(d.Get(Arg_CloudConnectionSpeed).(int))

	body := &models.CloudConnectionCreate{
		Name:  &name,
		Speed: &speed,
	}
	if v, ok := d.GetOk(Arg_CloudConnectionRouting); ok {
		body.GlobalRouting = v.(bool)
	}
	if v, ok := d.GetOk(Arg_CloudConnectionMetered); ok {
		body.Metered = v.(bool)
	}
	// networks
	if v, ok := d.GetOk(Arg_CloudConnectionNetworks); ok && v.(*schema.Set).Len() > 0 {
		body.Subnets = flex.ExpandStringList(v.(*schema.Set).List())
	}
	// classic
	if v, ok := d.GetOk(Arg_CloudConnectionClassic); ok {
		classicEnabled := v.(bool)
		classic := &models.CloudConnectionEndpointClassicUpdate{
			Enabled: classicEnabled,
		}
		if v, ok := d.GetOk(Arg_CloudConnectionGreCIDR); ok {
			greCIDR := v.(string)
			classic.Gre.Cidr = &greCIDR
		}
		if v, ok := d.GetOk(Arg_CloudConnectionGreDest); ok {
			greDest := v.(string)
			classic.Gre.DestIPAddress = &greDest
		}
		body.Classic = classic
	}

	// VPC
	if v, ok := d.GetOk(Arg_CloudConnectionVPC); ok {
		vpcEnabled := v.(bool)
		vpc := &models.CloudConnectionEndpointVPC{
			Enabled: vpcEnabled,
		}
		if v, ok := d.GetOk(Arg_CloudConnectionVPCCrns); ok && v.(*schema.Set).Len() > 0 {
			vpcIds := flex.ExpandStringList(v.(*schema.Set).List())
			vpcs := make([]*models.CloudConnectionVPC, len(vpcIds))
			for i, vpcId := range vpcIds {
				vpcs[i] = &models.CloudConnectionVPC{
					VpcID: &vpcId,
				}
			}
			vpc.Vpcs = vpcs
		}
		body.Vpc = vpc
	}

	client := st.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)
	cloudConnection, cloudConnectionJob, err := client.Create(body)
	if err != nil {
		log.Printf("[DEBUG] create cloud connection failed %v", err)
		return diag.FromErr(err)
	}

	if cloudConnection != nil {
		d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *cloudConnection.CloudConnectionID))
	} else if cloudConnectionJob != nil {
		d.SetId(fmt.Sprintf("%s/%s", cloudInstanceID, *cloudConnectionJob.CloudConnectionID))

		jobID := *cloudConnectionJob.JobRef.ID

		client := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)
		_, err = waitForIBMPIJobCompleted(ctx, client, jobID, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceIBMPICloudConnectionRead(ctx, d, meta)
}

func resourceIBMPICloudConnectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	Attr_CloudConnectionID := parts[1]

	ccName := d.Get(Arg_CloudConnectionName).(string)
	ccSpeed := int64(d.Get(Arg_CloudConnectionSpeed).(int))

	client := st.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)
	jobClient := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)

	if d.HasChangesExcept(Arg_CloudConnectionNetworks) {

		body := &models.CloudConnectionUpdate{
			Name:  &ccName,
			Speed: &ccSpeed,
		}
		if v, ok := d.GetOk(Arg_CloudConnectionRouting); ok {
			globalRouting := v.(bool)
			body.GlobalRouting = &globalRouting
		}
		if v, ok := d.GetOk(Arg_CloudConnectionMetered); ok {
			metered := v.(bool)
			body.Metered = &metered
		}
		// classic
		if v, ok := d.GetOk(Arg_CloudConnectionClassic); ok {
			classicEnabled := v.(bool)
			classic := &models.CloudConnectionEndpointClassicUpdate{
				Enabled: classicEnabled,
			}
			if v, ok := d.GetOk(Arg_CloudConnectionGreCIDR); ok {
				greCIDR := v.(string)
				classic.Gre.Cidr = &greCIDR
			}
			if v, ok := d.GetOk(Arg_CloudConnectionGreDest); ok {
				greDest := v.(string)
				classic.Gre.DestIPAddress = &greDest
			}
			body.Classic = classic
		} else {
			// need to disable classic if not provided
			classic := &models.CloudConnectionEndpointClassicUpdate{
				Enabled: false,
			}
			body.Classic = classic
		}
		// vpc
		if v, ok := d.GetOk(Arg_CloudConnectionVPC); ok {
			vpcEnabled := v.(bool)
			vpc := &models.CloudConnectionEndpointVPC{
				Enabled: vpcEnabled,
			}
			if v, ok := d.GetOk(Arg_CloudConnectionVPCCrns); ok && v.(*schema.Set).Len() > 0 {
				vpcIds := flex.ExpandStringList(v.(*schema.Set).List())
				vpcs := make([]*models.CloudConnectionVPC, len(vpcIds))
				for i, vpcId := range vpcIds {
					vpcs[i] = &models.CloudConnectionVPC{
						VpcID: &vpcId,
					}
				}
				vpc.Vpcs = vpcs
			}
			body.Vpc = vpc
		} else {
			// need to disable VPC if not provided
			vpc := &models.CloudConnectionEndpointVPC{
				Enabled: false,
			}
			body.Vpc = vpc
		}

		_, cloudConnectionJob, err := client.Update(Attr_CloudConnectionID, body)
		if err != nil {
			return diag.FromErr(err)
		}
		if cloudConnectionJob != nil {
			_, err = waitForIBMPIJobCompleted(ctx, jobClient, *cloudConnectionJob.ID, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}
	if d.HasChange(Arg_CloudConnectionNetworks) {
		oldRaw, newRaw := d.GetChange(Arg_CloudConnectionNetworks)
		old := oldRaw.(*schema.Set)
		new := newRaw.(*schema.Set)

		toAdd := new.Difference(old)
		toRemove := old.Difference(new)

		// call network add api for each toAdd
		for _, n := range flex.ExpandStringList(toAdd.List()) {
			_, jobReference, err := client.AddNetwork(Attr_CloudConnectionID, n)
			if err != nil {
				return diag.FromErr(err)
			}
			if jobReference != nil {
				_, err = waitForIBMPIJobCompleted(ctx, jobClient, *jobReference.ID, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return diag.FromErr(err)
				}
			}
		}

		// call network delete api for each toRemove
		for _, n := range flex.ExpandStringList(toRemove.List()) {
			_, jobReference, err := client.DeleteNetwork(Attr_CloudConnectionID, n)
			if err != nil {
				return diag.FromErr(err)
			}
			if jobReference != nil {
				_, err = waitForIBMPIJobCompleted(ctx, jobClient, *jobReference.ID, d.Timeout(schema.TimeoutUpdate))
				if err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	return resourceIBMPICloudConnectionRead(ctx, d, meta)
}

func resourceIBMPICloudConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	Attr_CloudConnectionID := parts[1]

	client := st.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)
	cloudConnection, err := client.Get(Attr_CloudConnectionID)
	if err != nil {
		uErr := errors.Unwrap(err)
		switch uErr.(type) {
		case *p_cloud_cloud_connections.PcloudCloudconnectionsGetNotFound:
			log.Printf("[DEBUG] cloud connection does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] get cloud connection failed %v", err)
		return diag.FromErr(err)
	}

	d.Set(Attr_CloudConnectionID, cloudConnection.CloudConnectionID)
	d.Set(Arg_CloudConnectionName, cloudConnection.Name)
	d.Set(Arg_CloudConnectionRouting, cloudConnection.GlobalRouting)
	d.Set(Arg_CloudConnectionMetered, cloudConnection.Metered)
	d.Set(Attr_CloudConnectionIbmIP, cloudConnection.IbmIPAddress)
	d.Set(Attr_CloudConnectionUserIP, cloudConnection.UserIPAddress)
	d.Set(Attr_CloudConnectionStatus, cloudConnection.LinkStatus)
	d.Set(Attr_CloudConnectionPort, cloudConnection.Port)
	d.Set(Arg_CloudConnectionSpeed, cloudConnection.Speed)
	d.Set(Arg_CloudInstanceID, cloudInstanceID)
	if cloudConnection.Networks != nil {
		networks := make([]string, 0)
		for _, ccNetwork := range cloudConnection.Networks {
			if ccNetwork != nil {
				networks = append(networks, *ccNetwork.NetworkID)
			}
		}
		d.Set(Arg_CloudConnectionNetworks, networks)
	}
	if cloudConnection.Classic != nil {
		d.Set(Arg_CloudConnectionClassic, cloudConnection.Classic.Enabled)
		if cloudConnection.Classic.Gre != nil {
			d.Set(Arg_CloudConnectionGreDest, cloudConnection.Classic.Gre.DestIPAddress)
			d.Set(Attr_CloudConnectionSourceGreAddr, cloudConnection.Classic.Gre.SourceIPAddress)
		}
	}
	if cloudConnection.Vpc != nil {
		d.Set(Arg_CloudConnectionVPC, cloudConnection.Vpc.Enabled)
		if cloudConnection.Vpc.Vpcs != nil && len(cloudConnection.Vpc.Vpcs) > 0 {
			vpcCRNs := make([]string, len(cloudConnection.Vpc.Vpcs))
			for i, vpc := range cloudConnection.Vpc.Vpcs {
				vpcCRNs[i] = *vpc.VpcID
			}
			d.Set(Arg_CloudConnectionVPCCrns, vpcCRNs)
		}
	}

	return nil
}
func resourceIBMPICloudConnectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(conns.ClientSession).IBMPISession()
	if err != nil {
		return diag.FromErr(err)
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	cloudInstanceID := parts[0]
	Attr_CloudConnectionID := parts[1]

	client := st.NewIBMPICloudConnectionClient(ctx, sess, cloudInstanceID)
	_, err = client.Get(Attr_CloudConnectionID)
	if err != nil {
		uErr := errors.Unwrap(err)
		switch uErr.(type) {
		case *p_cloud_cloud_connections.PcloudCloudconnectionsGetNotFound:
			log.Printf("[DEBUG] cloud connection does not exist %v", err)
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] get cloud connection failed %v", err)
		return diag.FromErr(err)
	}
	log.Printf("[INFO] Found cloud connection with id %s", Attr_CloudConnectionID)

	deleteJob, err := client.Delete(Attr_CloudConnectionID)
	if err != nil {
		log.Printf("[DEBUG] delete cloud connection failed %v", err)
		return diag.FromErr(err)
	}
	if deleteJob != nil {
		jobID := *deleteJob.ID

		client := st.NewIBMPIJobClient(ctx, sess, cloudInstanceID)
		_, err = waitForIBMPIJobCompleted(ctx, client, jobID, d.Timeout(schema.TimeoutDelete))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
