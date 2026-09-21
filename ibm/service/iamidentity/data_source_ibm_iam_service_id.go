// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func DataSourceIBMIamServiceID() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamServiceIDRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "Name of the serviceID",
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator("ibm_iam_service_id",
					"name"),
			},
			"service_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bound_to": &schema.Schema{
							Description: "bound to of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
							Deprecated:  "bound_to attribute in service_ids list has been deprecated",
						},
						"crn": &schema.Schema{
							Description: "CRN of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": &schema.Schema{
							Description: "description of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: "Version of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"locked": &schema.Schema{
							Description: "lock state of the serviceID",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"iam_id": &schema.Schema{
							Description: "The IAM ID of the serviceID",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMIamServiceIDValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "iam",
			CloudDataRange:             []string{"service:service_id", "resolved_to:name"},
			Required:                   true})

	iBMIAMServiceIDValidator := validate.ResourceValidator{ResourceName: "ibm_iam_service_id", Schema: validateSchema}
	return &iBMIAMServiceIDValidator
}

func dataSourceIBMIamServiceIDRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_service_id", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	name := d.Get("name").(string)
	start := ""
	allrecs := []iamidentityv1.ServiceID{}
	var pg int64 = 100
	for {
		listServiceIDOptions := iamidentityv1.ListServiceIdsOptions{
			AccountID: &userDetails.UserAccount,
			Pagesize:  &pg,
			Name:      &name,
		}
		if start != "" {
			listServiceIDOptions.Pagetoken = &start
		}

		serviceIDs, _, err := iamIdentityClient.ListServiceIds(&listServiceIDOptions)
		if err != nil {
			err = fmt.Errorf("Error listing Service Ids: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "list-service_ids").GetDiag()
		}
		start = flex.GetNextIAM(serviceIDs.Next)
		allrecs = append(allrecs, serviceIDs.Serviceids...)
		if start == "" {
			break
		}
	}
	if len(allrecs) == 0 {
		err = fmt.Errorf("No serviceID found with name [%s]", name)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_service_id", "read", "set-service_ids").GetDiag()
	}

	serviceIDListMap := make([]map[string]interface{}, 0, len(allrecs))
	for _, serviceID := range allrecs {
		l := map[string]interface{}{
			"id":          serviceID.ID,
			"version":     serviceID.EntityTag,
			"description": serviceID.Description,
			"crn":         serviceID.CRN,
			"locked":      serviceID.Locked,
			"iam_id":      serviceID.IamID,
			// "bound_to":    serviceID.BoundTo,
		}
		serviceIDListMap = append(serviceIDListMap, l)
	}
	d.SetId(name)
	d.Set("service_ids", serviceIDListMap)
	return nil
}
