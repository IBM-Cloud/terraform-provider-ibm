// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBaasDownloadAgent() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBaasDownloadAgentCreate,
		ReadContext:   resourceIbmBaasDownloadAgentRead,
		DeleteContext: resourceIbmBaasDownloadAgentDelete,
		UpdateContext: resourceIbmBaasDownloadAgentUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.",
			},
			"file_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"linux_params": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "Linux agent parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"package_type": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies the type of installer.",
						},
					},
				},
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_baas_download_agent", "platform"),
				Description: "Specifies the platform for which agent needs to be downloaded.",
			},
		},
	}
}

func ResourceIbmBaasDownloadAgentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "platform",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "kLinux, kWindows",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_baas_download_agent", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBaasDownloadAgentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_download_agent", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	downloadAgentOptions := &backuprecoveryv1.DownloadAgentOptions{}

	downloadAgentOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	downloadAgentOptions.SetPlatform(d.Get("platform").(string))
	if _, ok := d.GetOk("linux_params"); ok {
		linuxParamsModel, err := ResourceIbmBaasDownloadAgentMapToLinuxAgentParams(d.Get("linux_params.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_download_agent", "read", "parse-linux_params").GetDiag()
		}
		downloadAgentOptions.SetLinuxParams(linuxParamsModel)
	}

	typeString, _, err := backupRecoveryClient.DownloadAgentWithContext(context, downloadAgentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DownloadAgentWithContext failed: %s", err.Error()), "ibm_baas_download_agent", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmBaasAgentDownloadID(d))

	err = saveToFile(typeString, d.Get("file_path").(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_baas_download_agent", "read", "parse-linux_params").GetDiag()
	}

	return resourceIbmBaasDownloadAgentRead(context, d, meta)
}

func saveToFile(response io.ReadCloser, filePath string) error {
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response)
	if err != nil {
		return err
	}

	err = response.Close()
	if err != nil {
		return err
	}

	return nil
}

func resourceIbmBaasAgentDownloadID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmBaasDownloadAgentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIbmBaasDownloadAgentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "Delete operation is not supported for this resource. The resource will be removed from the terraform file but will continue to exist in the backend.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmBaasDownloadAgentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Resource Update Will Only Affect Terraform State",
		Detail:   "Update operation for this resource not supported and will only affect the terraform statefile. No changes will be made to actual backend resource. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func ResourceIbmBaasDownloadAgentMapToLinuxAgentParams(modelMap map[string]interface{}) (*backuprecoveryv1.LinuxAgentParams, error) {
	model := &backuprecoveryv1.LinuxAgentParams{}
	model.PackageType = core.StringPtr(modelMap["package_type"].(string))
	return model, nil
}

func ResourceIbmBaasDownloadAgentLinuxAgentParamsToMap(model *backuprecoveryv1.LinuxAgentParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["package_type"] = *model.PackageType
	return modelMap, nil
}
