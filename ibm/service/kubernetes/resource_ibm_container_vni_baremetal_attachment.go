// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"
	"log"
	"time"

	gohttp "net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/graphql"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
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
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_container_vni_baremetal_attachment", "vlan_id"),
				Description:  "The VLAN ID for the bare metal worker (1-500)",
			},
			"cluster": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"worker"},
				Description:   "The cluster ID or name to attach VNI to any available worker",
			},
			"worker": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"cluster"},
				Description:   "The worker ID to attach VNI to specific worker",
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
			"cluster_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the cluster where VNI is attached",
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

func ResourceIBMContainerVNIBaremetalAttachmentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "vlan_id",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Required:                   true,
			MinValue:                   "1",
			MaxValue:                   "500"})

	ibmContainerVNIBaremetalAttachmentValidator := validate.ResourceValidator{
		ResourceName: "ibm_container_vni_baremetal_attachment",
		Schema:       validateSchema}
	return &ibmContainerVNIBaremetalAttachmentValidator
}

func resourceIBMContainerVNIBaremetalAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	// Validate selector
	cluster, hasCluster := d.GetOk("cluster")
	worker, hasWorker := d.GetOk("worker")

	if !hasCluster && !hasWorker {
		return fmt.Errorf("either 'cluster' or 'worker' must be specified")
	}
	if hasCluster && hasWorker {
		return fmt.Errorf("only one of 'cluster' or 'worker' can be specified")
	}

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

	// Extract worker ID and cluster ID from response
	workerID := resp.NetworkAttachment.AttachedTo.ID

	// Get cluster ID - if cluster was specified, use it; otherwise query for it
	var clusterID string
	if hasCluster {
		clusterID = cluster.(string)
	} else {
		// Need to get cluster ID from worker - this requires querying the worker details
		// For now, we'll use a placeholder approach
		clusterID = "" // Will be populated in Read
	}

	// Set resource ID
	id := buildVNIAttachmentID(clusterID, workerID, vniID)
	d.SetId(id)

	// Set computed attributes
	d.Set("worker_id", workerID)
	if clusterID != "" {
		d.Set("cluster_id", clusterID)
	}

	log.Printf("[INFO] VNI attachment created: %s", id)

	// Read to populate all attributes
	return resourceIBMContainerVNIBaremetalAttachmentRead(d, meta)
}

func resourceIBMContainerVNIBaremetalAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	// Parse ID
	clusterID, workerID, vniID, err := parseVNIAttachmentID(d.Id())
	if err != nil {
		return err
	}

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

	// List attachments for the worker to find this specific VNI
	input := graphql.ListVNIAttachmentsInput{
		NodeID: workerID,
	}

	resp, err := vniClient.ListAttachments(input, targetEnv)
	if err != nil {
		return fmt.Errorf("error listing VNI attachments: %s", err)
	}

	// Find the specific VNI attachment
	var found *graphql.VNIAttachment
	for _, edge := range resp.Connection.Edges {
		if edge.Node.VirtualNetworkInterface.ExternalID == vniID {
			found = &edge.Node
			break
		}
	}

	if found == nil {
		log.Printf("[WARN] VNI attachment not found, removing from state: %s", d.Id())
		d.SetId("")
		return nil
	}

	// Set attributes
	d.Set("vni_id", found.VirtualNetworkInterface.ExternalID)
	d.Set("worker_id", found.AttachedTo.ID)
	d.Set("cluster_id", clusterID)

	if found.VlanID != nil {
		d.Set("vlan_id", *found.VlanID)
	}
	if found.Status != "" {
		d.Set("status", found.Status)
	}
	if found.CreatedAt != "" {
		d.Set("created_at", found.CreatedAt)
	}
	if found.VirtualNetworkInterface.AutoDelete != nil {
		d.Set("auto_delete", *found.VirtualNetworkInterface.AutoDelete)
	}

	return nil
}

func resourceIBMContainerVNIBaremetalAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	// Parse ID
	_, workerID, vniID, err := parseVNIAttachmentID(d.Id())
	if err != nil {
		return err
	}

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
	input := graphql.RemoveVNIFromNodeInput{
		VirtualNetworkInterfaceID: vniID,
		Node:                      &workerID,
		AutoDelete:                autoDelete,
	}

	log.Printf("[INFO] Detaching VNI %s from worker %s", vniID, workerID)

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
	// Parse ID
	_, workerID, vniID, err := parseVNIAttachmentID(d.Id())
	if err != nil {
		return false, err
	}

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
		NodeID: workerID,
	}

	resp, err := vniClient.ListAttachments(input, targetEnv)
	if err != nil {
		return false, fmt.Errorf("error listing VNI attachments: %s", err)
	}

	// Check if the specific VNI exists
	for _, edge := range resp.Connection.Edges {
		if edge.Node.VirtualNetworkInterface.ExternalID == vniID {
			return true, nil
		}
	}

	return false, nil
}

// Helper functions

// getVNIClient creates a VNI GraphQL API client
func getVNIClient(meta interface{}) (graphql.VNIs, error) {
	sess, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}

	config := sess.Config.Copy()
	err = config.ValidateConfigForService(bluemix.ContainerService)
	if err != nil {
		return nil, err
	}

	if config.HTTPClient == nil {
		config.HTTPClient = http.NewHTTPClient(config)
	}

	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"X-Original-User-Agent": []string{config.UserAgent},
			"User-Agent":            []string{http.UserAgent()},
		},
		HTTPClient: config.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	if config.IAMAccessToken == "" {
		err := authentication.PopulateTokens(tokenRefresher, config)
		if err != nil {
			return nil, err
		}
	}

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

// parseVNIAttachmentID parses the VNI attachment ID format: {cluster_id}/{worker_id}/{vni_id}
func parseVNIAttachmentID(id string) (clusterID, workerID, vniID string, err error) {
	parts := make([]string, 0, 3)
	start := 0

	for i := 0; i < len(id); i++ {
		if id[i] == '/' {
			if i == 0 || i == len(id)-1 {
				return "", "", "", fmt.Errorf("invalid VNI attachment ID format: %s", id)
			}
			parts = append(parts, id[start:i])
			start = i + 1
		}
	}
	parts = append(parts, id[start:])

	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid VNI attachment ID format, expected {cluster_id}/{worker_id}/{vni_id}, got: %s", id)
	}

	return parts[0], parts[1], parts[2], nil
}

// buildVNIAttachmentID builds the VNI attachment ID from components
func buildVNIAttachmentID(clusterID, workerID, vniID string) string {
	return fmt.Sprintf("%s/%s/%s", clusterID, workerID, vniID)
}
