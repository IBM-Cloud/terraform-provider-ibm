package secretsmanager

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func ResourceIbmSmPrivateCertificateConfigurationActionRotateIntermediate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateCreateOrUpdate,
		ReadContext:   resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateRead,
		UpdateContext: resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateCreateOrUpdate,
		DeleteContext: resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name that uniquely identifies a configuration",
			},
		},
	}
}

func resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateCreateOrUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", PrivateCertConfigActionRotateIntermediate, "create/update")
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerClient, d)
	instanceId := d.Get("instance_id").(string)
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d))

	createConfigurationActionOptions := &secretsmanagerv2.CreateConfigurationActionOptions{}

	configurationActionPrototypeModel, err := resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateMapToConfigurationActionPrototype(d)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", PrivateCertConfigActionRotateIntermediate, "create/update")
		return tfErr.GetDiag()
	}
	createConfigurationActionOptions.SetConfigActionPrototype(configurationActionPrototypeModel)
	createConfigurationActionOptions.SetName(d.Get("name").(string))

	_, response, err := secretsManagerClient.CreateConfigurationActionWithContext(context, createConfigurationActionOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateConfigurationActionWithContext failed %s\n%s", err, response)
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateConfigurationActionWithContext failed: %s\n%s", err.Error(), response), PrivateCertConfigActionRotateIntermediate, "create/update")
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/rotate", region, instanceId, d.Get("name").(string)))

	return nil
}

func resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateMapToConfigurationActionPrototype(d *schema.ResourceData) (secretsmanagerv2.ConfigurationActionPrototypeIntf, error) {
	model := &secretsmanagerv2.PrivateCertificateConfigurationActionRotatePrototype{
		ActionType: core.StringPtr("private_cert_configuration_action_rotate_intermediate"),
	}
	return model, nil
}

func resourceIbmSmPrivateCertificateConfigurationActionRotateIntermediateDataToMap(model secretsmanagerv2.PrivateCertificateConfigurationCACertificate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	if model.Certificate != nil {
		modelMap["certificate"] = model.Certificate
	}
	if model.IssuingCa != nil {
		modelMap["issuing_ca"] = model.IssuingCa
	}
	if model.CaChain != nil {
		modelMap["ca_chain"] = model.CaChain
	}
	if model.Expiration != nil {
		modelMap["expiration"] = flex.IntValue(model.Expiration)
	}
	return modelMap, nil
}
