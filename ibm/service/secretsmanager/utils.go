package secretsmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ArbitrarySecretType          = "arbitrary"
	UsernamePasswordSecretType   = "username_password"
	CustomCredentialsSecretType  = "custom_credentials"
	IAMCredentialsSecretType     = "iam_credentials"
	ServiceCredentialsSecretType = "service_credentials"
	KvSecretType                 = "kv"
	ImportedCertSecretType       = "imported_cert"
	PublicCertSecretType         = "public_cert"
	PrivateCertSecretType        = "private_cert"

	ArbitrarySecretResourceName          = "ibm_sm_arbitrary_secret"
	UsernamePasswordSecretResourceName   = "ibm_sm_username_password_secret"
	IAMCredentialsSecretResourceName     = "ibm_sm_iam_credentials_secret"
	ServiceCredentialsSecretResourceName = "ibm_sm_service_credentials_secret"
	KvSecretResourceName                 = "ibm_sm_kv_secret"
	ImportedCertSecretResourceName       = "ibm_sm_imported_certificate"
	PublicCertSecretResourceName         = "ibm_sm_public_certificate"
	PrivateCertSecretResourceName        = "ibm_sm_private_certificate"
	CustomCredentialsSecretResourceName  = "ibm_sm_custom_credentials_secret"

	EnRegistrationResourceName                           = "ibm_sm_en_registration"
	IAMCredentialsConfigResourceName                     = "ibm_sm_iam_credentials_configuration"
	CustomCredentialsConfigResourceName                  = "ibm_sm_custom_credentials_configuration"
	ConfigurationsResourceName                           = "ibm_sm_configurations"
	PrivateCertConfigIntermediateCAResourceName          = "ibm_sm_private_certificate_configuration_intermediate_ca"
	PrivateCertConfigRootCAResourceName                  = "ibm_sm_private_certificate_configuration_root_ca"
	PrivateCertConfigTemplateResourceName                = "ibm_sm_private_certificate_configuration_template"
	PrivateCertConfigActionSetSigned                     = "ibm_sm_private_certificate_configuration_action_set_signed"
	PrivateCertConfigActionSignCsr                       = "ibm_sm_private_certificate_configuration_action_sign_csr"
	PublicCertConfigCALetsEncryptResourceName            = "ibm_sm_public_certificate_configuration_ca_lets_encrypt"
	PublicCertConfigDnsCISResourceName                   = "ibm_sm_public_certificate_configuration_dns_cis"
	PublicCertConfigDnsClassicInfrastructureResourceName = "ibm_sm_public_certificate_configuration_dns_classic_infrastructure"
	PublicCertConfigActionValidateManualDNSResourceName  = "ibm_sm_public_certificate_action_validate_manual_dns"

	SecretGroupResourceName  = "ibm_sm_secret_group"
	SecretGroupsResourceName = "ibm_sm_secret_groups"
	SecretsResourceName      = "ibm_sm_secrets"
)

func getRegion(originalClient *secretsmanagerv2.SecretsManagerV2, d *schema.ResourceData) string {
	_, ok := d.GetOk("region")
	if ok {
		return d.Get("region").(string)
	} else {
		// extract region from base URL (provider config)
		// base url is like that : "https://<private.>secrets-manager.<region>.<rest of domain>"
		baseUrl := originalClient.Service.GetServiceURL()
		u := strings.Replace(baseUrl, "private.", "", 1)
		return strings.Split(u, ".")[1]
	}
}

// Clone the base secrets manager client and set the API endpoint per the instance
func getEndpointType(originalClient *secretsmanagerv2.SecretsManagerV2, d *schema.ResourceData) string {
	_, ok := d.GetOk("endpoint_type")
	if ok {
		return d.Get("endpoint_type").(string)
	} else {
		baseUrl := originalClient.Service.GetServiceURL()

		if strings.Contains(baseUrl, "private.") {
			return "private"
		} else {
			return "public"
		}
	}
}

// Get the Secrets Manager session and the endpoints file from the provider's configuration
func getSecretsManagerSession(clientSession conns.ClientSession) (*secretsmanagerv2.SecretsManagerV2, string, error) {
	secretsManagerClient, err := clientSession.SecretsManagerV2()
	if err != nil {
		return secretsManagerClient, "", err
	}

	bmxsession, err := clientSession.BluemixSession()
	if err != nil {
		return secretsManagerClient, "", err
	}

	return secretsManagerClient, bmxsession.Config.EndpointsFile, nil
}

// Clone the base secrets manager client and set the API endpoint per the instance
func getClientWithInstanceEndpoint(originalClient *secretsmanagerv2.SecretsManagerV2, instanceId string, region string,
	endpointType string, endpointsFile string) *secretsmanagerv2.SecretsManagerV2 {
	// build the api endpoint
	domain := "appdomain.cloud"

	// Check if we're running in the staging environment based on the configuration of the IAM API endpoint
	iamUrl := os.Getenv("IBMCLOUD_IAM_API_ENDPOINT")
	if iamUrl == "" {
		iamUrl = conns.FileFallBack(endpointsFile, endpointType, "IBMCLOUD_IAM_API_ENDPOINT", region, "https://iam.cloud.ibm.com")
	}

	if strings.Contains(iamUrl, "test") {
		domain = "test.appdomain.cloud"
	}
	var endpoint string
	if endpointType == "private" {
		endpoint = fmt.Sprintf("https://%s.private.%s.secrets-manager.%s", instanceId, region, domain)
	} else {
		endpoint = fmt.Sprintf("https://%s.%s.secrets-manager.%s", instanceId, region, domain)
	}

	// clone the client and set endpoint
	newClient := &secretsmanagerv2.SecretsManagerV2{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(endpoint)
	return newClient
}

// Add the fields needed for building the instance endpoint to the given schema
func AddInstanceFields(resource *schema.Resource) *schema.Resource {
	resource.Schema["instance_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "The ID of the Secrets Manager instance.",
	}
	resource.Schema["region"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: "The region of the Secrets Manager instance.",
	}
	resource.Schema["endpoint_type"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "public or private.",
	}

	return resource
}

func StringIsIntBetween(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		vs, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		v, err := strconv.Atoi(vs)

		if err != nil {
			errors = append(errors, fmt.Errorf("expected %s to represent an integer", k))
			return warnings, errors
		}

		if v < min || v > max {
			errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d), got %d", k, min, max, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

func StringIsIntWithMinimum(min int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		vs, ok := i.(string)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
			return warnings, errors
		}

		v, err := strconv.Atoi(vs)

		if err != nil {
			errors = append(errors, fmt.Errorf("expected %s to represent an integer", k))
			return warnings, errors
		}

		if v < min {
			errors = append(errors, fmt.Errorf("expected %s to be st least %d, got %d", k, min, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

func DateTimeToRFC3339(dt *strfmt.DateTime) (s string) {
	if dt != nil {
		s = time.Time(*dt).Format(time.RFC3339)
	}
	return
}

func getSecretByIdOrByName(context context.Context, d *schema.ResourceData, meta interface{}, secretType string, dataSourceName string) (secretsmanagerv2.SecretIntf, string, string, diag.Diagnostics) {

	secretsManagerClient, endpointsFile, err := getSecretsManagerSession(meta.(conns.ClientSession))
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "", fmt.Sprintf("(Data) %s", dataSourceName), "read")
		return nil, "", "", tfErr.GetDiag()
	}
	region := getRegion(secretsManagerClient, d)
	instanceId := d.Get("instance_id").(string)
	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, instanceId, region, getEndpointType(secretsManagerClient, d), endpointsFile)

	secretId := d.Get("secret_id").(string)
	secretName := d.Get("name").(string)
	groupName := d.Get("secret_group_name").(string)

	log.Printf("[DEBUG] getSecretByIdOrByName %q %q %q %q\n", secretId, secretName, groupName, secretType)

	var secretIntf secretsmanagerv2.SecretIntf
	var response *core.DetailedResponse

	if secretId != "" {
		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}
		getSecretOptions.SetID(secretId)

		secretIntf, response, err = secretsManagerClient.GetSecretWithContext(context, getSecretOptions)
		if err != nil {
			log.Printf("[DEBUG] GetSecretWithContext failed %s\n%s", err, response)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecretWithContext failed %s\n%s", err, response), fmt.Sprintf("(Data) %s", dataSourceName), "read")
			return nil, "", "", tfErr.GetDiag()
		}
		return secretIntf, region, instanceId, nil
	}

	if secretName != "" && groupName != "" {
		// Locate secret by name
		getSecretByNameOptions := &secretsmanagerv2.GetSecretByNameTypeOptions{}

		getSecretByNameOptions.SetName(secretName)
		getSecretByNameOptions.SetSecretType(secretType)
		getSecretByNameOptions.SetSecretGroupName(groupName)

		secretIntf, response, err = secretsManagerClient.GetSecretByNameTypeWithContext(context, getSecretByNameOptions)
		if err != nil {
			log.Printf("[DEBUG] GetSecretByNameTypeWithContext failed %s\n%s", err, response)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSecretByNameTypeWithContext failed %s\n%s", err, response), fmt.Sprintf("(Data) %s", dataSourceName), "read")
			return nil, "", "", tfErr.GetDiag()
		}
		return secretIntf, region, instanceId, nil
	}

	tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Missing required arguments. Please make sure that either \"secret_id\" or \"name\" and \"secret_group_name\" are provided\n"), fmt.Sprintf("(Data) %s", dataSourceName), "read")
	return nil, "", "", tfErr.GetDiag()
}

func secretVersionMetadataAsPatchFunction(secretVersionMetadataPatch *secretsmanagerv2.SecretVersionMetadataPatch) (_patch map[string]interface{}, err error) {
	jsonData, err := json.Marshal(struct {
		VersionCustomMetadata map[string]interface{} `json:"version_custom_metadata"`
	}{VersionCustomMetadata: secretVersionMetadataPatch.VersionCustomMetadata})
	if err == nil {
		err = json.Unmarshal(jsonData, &_patch)
	}
	return
}
