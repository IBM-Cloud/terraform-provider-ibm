// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"
	"log"
	gohttp "net/http"
	"time"

	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/graphql"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMContainerVNIBaremetalAttachment() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerVNIBaremetalAttachmentCreate,
		Read:     resourceIBMContainerVNIBaremetalAttachmentRead,
		Delete:   resourceIBMContainerVNIBaremetalAttachmentDelete,
		Exists:   resourceIBMContainerVNIBaremetalAttachmentExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"vni_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the VNI to attach",
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value < 1 || value > 500 {
						errors = append(errors, fmt.Errorf("%q must be between 1 and 500, got: %d", k, value))
					}
					return
				},
				Description: "The VLAN ID for the bare metal worker (1-500)",
			},
			"cluster": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"cluster", "worker"},
				Description:  "The cluster ID or name to attach VNI to any available worker",
			},
			"worker": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"cluster", "worker"},
				Description:  "The worker ID to attach VNI to specific worker",
			},
			"auto_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether to delete the VNI when the attachment is destroyed",
			},
			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the resource group",
			},
			"worker_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the worker where VNI is attached",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the attachment",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp when the attachment was created",
			},
		},
	}
}

func resourceIBMContainerVNIBaremetalAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	cluster, hasCluster := d.GetOk("cluster")
	worker, _ := d.GetOk("worker")

	// Get VNI client
	vniClient, err := getVNIClient(meta)
	if err != nil {
		return fmt.Errorf("error creating VNI client: %s", err)
	}

	// Get target header
	targetEnv, err := getVNITargetHeader(d, meta)
	if err != nil {
		return fmt.Errorf("error creating target header: %s", err)
	}

	// Prepare input
	vniID := d.Get("vni_id").(string)
	vlanID := d.Get("vlan_id").(int)
	autoDelete := d.Get("auto_delete").(bool)

	input := graphql.AddVNIToBareMetalNodeInput{
		VirtualNetworkInterfaceID: vniID,
		VlanID:                    vlanID,
		AutoDelete:                autoDelete,
	}

	if hasCluster {
		clusterStr := cluster.(string)
		input.Cluster = &clusterStr
	} else {
		workerStr := worker.(string)
		input.Node = &workerStr
	}

	log.Printf("[INFO] Attaching VNI %s with VLAN ID %d", vniID, vlanID)

	// Attach VNI
	resp, err := vniClient.AttachToBareMetalNode(input, targetEnv)
	if err != nil {
		return fmt.Errorf("error attaching VNI: %s", err)
	}

	if resp == nil {
		return fmt.Errorf("error: received nil response from VNI attach operation")
	}

	// Extract worker ID from response
	workerID := resp.NetworkAttachment.AttachedTo.ID

	// Set resource ID (vni_id is unique and avoids issues with worker ID format)
	d.SetId(vniID)

	// Set computed attributes
	d.Set("worker_id", workerID)

	log.Printf("[INFO] VNI attachment created: %s", vniID)

	// Read to populate all attributes
	return resourceIBMContainerVNIBaremetalAttachmentRead(d, meta)
}

func resourceIBMContainerVNIBaremetalAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	vniID := d.Id()

	// Get VNI client
	vniClient, err := getVNIClient(meta)
	if err != nil {
		return fmt.Errorf("error creating VNI client: %s", err)
	}

	// Get target header
	targetEnv, err := getVNITargetHeader(d, meta)
	if err != nil {
		return fmt.Errorf("error creating target header: %s", err)
	}

	// workerID is stored in state from the create/previous read
	workerID, _ := d.GetOk("worker_id")

	// List attachments for the worker to find this specific VNI
	input := graphql.ListVNIAttachmentsInput{
		NodeID: workerID.(string),
	}

	resp, err := vniClient.ListAttachments(input, targetEnv)
	if err != nil {
		return fmt.Errorf("error listing VNI attachments: %s", err)
	}

	// Find the specific VNI attachment
	var attachment *graphql.VNIAttachment
	for _, edge := range resp.Connection.Edges {
		if edge.Node.VirtualNetworkInterface.ExternalID == vniID {
			attachment = &edge.Node
			break
		}
	}

	if attachment == nil {
		log.Printf("[WARN] VNI attachment not found, removing from state: %s", d.Id())
		d.SetId("")
		return nil
	}

	// Set attributes
	d.Set("vni_id", attachment.VirtualNetworkInterface.ExternalID)
	d.Set("worker_id", attachment.AttachedTo.ID)

	if attachment.VlanID != nil && *attachment.VlanID > 0 {
		d.Set("vlan_id", *attachment.VlanID)
	}
	if attachment.Status != "" {
		d.Set("status", attachment.Status)
	}
	if attachment.CreatedAt != "" {
		d.Set("created_at", attachment.CreatedAt)
	}
	if attachment.VirtualNetworkInterface.AutoDelete != nil {
		d.Set("auto_delete", *attachment.VirtualNetworkInterface.AutoDelete)
	}

	return nil
}

func resourceIBMContainerVNIBaremetalAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	vniID := d.Id()
	workerID, _ := d.GetOk("worker_id")

	// Get VNI client
	vniClient, err := getVNIClient(meta)
	if err != nil {
		return fmt.Errorf("error creating VNI client: %s", err)
	}

	// Get target header
	targetEnv, err := getVNITargetHeader(d, meta)
	if err != nil {
		return fmt.Errorf("error creating target header: %s", err)
	}

	// Prepare input
	autoDelete := d.Get("auto_delete").(bool)
	workerIDStr := workerID.(string)
	input := graphql.RemoveVNIFromNodeInput{
		VirtualNetworkInterfaceID: vniID,
		Node:                      &workerIDStr,
		AutoDelete:                autoDelete,
	}

	log.Printf("[INFO] Detaching VNI %s from worker %s", vniID, workerIDStr)

	// Detach VNI
	_, err = vniClient.DetachFromNode(input, targetEnv)
	if err != nil {
		return fmt.Errorf("error detaching VNI: %s", err)
	}

	log.Printf("[INFO] VNI attachment deleted: %s", d.Id())

	d.SetId("")
	return nil
}

func resourceIBMContainerVNIBaremetalAttachmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	vniID := d.Id()
	workerID, _ := d.GetOk("worker_id")

	// Get VNI client
	vniClient, err := getVNIClient(meta)
	if err != nil {
		return false, fmt.Errorf("error creating VNI client: %s", err)
	}

	// Get target header
	targetEnv, err := getVNITargetHeader(d, meta)
	if err != nil {
		return false, fmt.Errorf("error creating target header: %s", err)
	}

	// List attachments for the worker
	input := graphql.ListVNIAttachmentsInput{
		NodeID: workerID.(string),
	}

	resp, err := vniClient.ListAttachments(input, targetEnv)
	if err != nil {
		return false, fmt.Errorf("error listing VNI attachments: %s", err)
	}

	// Check if the specific VNI attachment exists
	for _, edge := range resp.Connection.Edges {
		if edge.Node.VirtualNetworkInterface.ExternalID == vniID {
			return true, nil
		}
	}

	return false, nil
}

// getVNIClient creates a VNI GraphQL client
func getVNIClient(meta interface{}) (graphql.VNIs, error) {
	sess, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}

	config := sess.Config.Copy()
	config.HTTPClient = http.NewHTTPClient(config)

	// Set up authentication
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent":            []string{http.UserAgent()},
			"X-Original-User-Agent": []string{config.UserAgent},
		},
	})
	if err != nil {
		return nil, err
	}

	// Set GraphQL endpoint
	if config.Endpoint == nil {
		ep, err := config.EndpointLocator.ContainerEndpoint()
		if err != nil {
			return nil, err
		}
		config.Endpoint = &ep
	}

	c := client.New(config, bluemix.ContainerService, tokenRefresher)
	return graphql.NewVNIAPI(c), nil
}

// getVNITargetHeader creates a target header for VNI operations
func getVNITargetHeader(d *schema.ResourceData, meta interface{}) (containerv1.ClusterTargetHeader, error) {
	_, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return containerv1.ClusterTargetHeader{}, err
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return containerv1.ClusterTargetHeader{}, err
	}

	accountID := userDetails.UserAccount

	targetEnv := containerv1.ClusterTargetHeader{
		AccountID: accountID,
	}

	resourceGroup := ""
	if v, ok := d.GetOk("resource_group_id"); ok {
		resourceGroup = v.(string)
		targetEnv.ResourceGroup = resourceGroup
	}

	return targetEnv, nil
}
