package endpoints

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/helpers"
)

//EndpointLocator ...
type EndpointLocator interface {
	AccountManagementEndpoint() (string, error)
	CertificateManagerEndpoint() (string, error)
	CFAPIEndpoint() (string, error)
	ContainerEndpoint() (string, error)
	ContainerRegistryEndpoint() (string, error)
	CisEndpoint() (string, error)
	GlobalSearchEndpoint() (string, error)
	GlobalTaggingEndpoint() (string, error)
	IAMEndpoint() (string, error)
	IAMPAPEndpoint() (string, error)
	ICDEndpoint() (string, error)
	MCCPAPIEndpoint() (string, error)
	ResourceManagementEndpoint() (string, error)
	ResourceControllerEndpoint() (string, error)
	ResourceCatalogEndpoint() (string, error)
	UAAEndpoint() (string, error)
	CseEndpoint() (string, error)
	SchematicsEndpoint() (string, error)
	UserManagementEndpoint() (string, error)
	HpcsEndpoint() (string, error)
}

const (
	//ErrCodeServiceEndpoint ...
	ErrCodeServiceEndpoint = "ServiceEndpointDoesnotExist"
)

var regionToEndpoint = map[string]map[string]string{
	"account": {
		"global": "https://accounts.cloud.ibm.com",
	},
	"certificate-manager": {
		"us-south": "https://us-south.certificate-manager.cloud.ibm.com",
		"us-east":  "https://us-east.certificate-manager.cloud.ibm.com",
		"eu-gb":    "https://eu-gb.certificate-manager.cloud.ibm.com",
		"au-syd":   "https://au-syd.certificate-manager.cloud.ibm.com",
		"eu-de":    "https://eu-de.certificate-manager.cloud.ibm.com",
		"jp-tok":   "https://jp-tok.certificate-manager.cloud.ibm.com",
	},
	"cf": {
		"us-south": "https://api.ng.bluemix.net",
		"us-east":  "https://api.us-east.bluemix.net",
		"eu-gb":    "https://api.eu-gb.bluemix.net",
		"au-syd":   "https://api.au-syd.bluemix.net",
		"eu-de":    "https://api.eu-de.bluemix.net",
		"jp-tok":   "https://api.jp-tok.bluemix.net",
	},
	"cr": {
		"us-south": "https://registry.ng.bluemix.net",
		"us-east":  "https://registry.ng.bluemix.net",
		"eu-de":    "https://registry.eu-de.bluemix.net",
		"au-syd":   "https://registry.au-syd.bluemix.net",
		"eu-gb":    "https://registry.eu-gb.bluemix.net",
		"jp-tok":   "https://registry.jp-tok.bluemix.net",
	},
	"cs": {
		"global": "https://containers.cloud.ibm.com/global",
	},
	"cis": {
		"global": "https://api.cis.cloud.ibm.com",
	},
	"global-search": {
		"global": "https://api.global-search-tagging.cloud.ibm.com",
	},
	"global-tagging": {
		"global": "https://tags.global-search-tagging.cloud.ibm.com",
	},
	"iam": {
		"global": "https://iam.cloud.ibm.com",
	},
	"iampap": {
		"global": "https://iam.cloud.ibm.com",
	},
	"icd": {
		"us-south": "https://api.us-south.databases.cloud.ibm.com",
		"us-east":  "https://api.us-east.databases.cloud.ibm.com",
		"eu-de":    "https://api.eu-de.databases.cloud.ibm.com",
		"eu-fr2":   "https://api.eu-fr2.databases.cloud.ibm.com",
		"eu-gb":    "https://api.eu-gb.databases.cloud.ibm.com",
		"au-syd":   "https://api.au-syd.databases.cloud.ibm.com",
		"jp-tok":   "https://api.jp-tok.databases.cloud.ibm.com",
		"osl01":    "https://api.osl01.databases.cloud.ibm.com",
		"seo01":    "https://api.seo01.databases.cloud.ibm.com",
		"che01":    "https://api.che01.databases.cloud.ibm.com",
	},
	"mccp": {
		"us-south": "https://mccp.us-south.cf.cloud.ibm.com",
		"us-east":  "https://mccp.us-east.cf.cloud.ibm.com",
		"eu-gb":    "https://mccp.eu-gb.cf.cloud.ibm.com",
		"au-syd":   "https://mccp.au-syd.cf.cloud.ibm.com",
		"eu-de":    "https://mccp.eu-de.cf.cloud.ibm.com",
	},
	"resource-manager": {
		"global": "https://resource-controller.cloud.ibm.com",
	},
	"resource-catalog": {
		"global": "https://globalcatalog.cloud.ibm.com",
	},
	"resource-controller": {
		"global": "https://resource-controller.cloud.ibm.com",
	},
	"uaa": {
		"us-south": "https://iam.cloud.ibm.com/cloudfoundry/login/us-south",
		"us-east":  "https://iam.cloud.ibm.com/cloudfoundry/login/us-east",
		"eu-gb":    "https://iam.cloud.ibm.com/cloudfoundry/login/uk-south",
		"au-syd":   "https://iam.cloud.ibm.com/cloudfoundry/login/ap-south",
		"eu-de":    "https://iam.cloud.ibm.com/cloudfoundry/login/eu-central",
	},
	"cse": {
		"global": "https://api.serviceendpoint.cloud.ibm.com",
	},
	"schematics": {
		"us-south": "https://us.schematics.cloud.ibm.com",
		"eu-gb":    "https://eu-gb.schematics.cloud.ibm.com",
		"eu-de":    "https://eu-de.schematics.cloud.ibm.com",
	},
	"usermanagement": {
		"global": "https://user-management.cloud.ibm.com",
	},
	"hpcs": {
		"us-south": "https://us-south.broker.hs-crypto.cloud.ibm.com/crypto_v2/",
		"us-east":  "https://us-east.broker.hs-crypto.cloud.ibm.com/crypto_v2/",
		"au-syd":   "https://au-syd.broker.hs-crypto.cloud.ibm.com/crypto_v2/",
		"eu-de":    "https://eu-de.broker.hs-crypto.cloud.ibm.com/crypto_v2/",
	},
}

func init() {
	//TODO populate the endpoints which can be retrieved from given endpoints dynamically
	//Example - UAA can be found from the CF endpoint
}

type endpointLocator struct {
	region string
}

//NewEndpointLocator ...
func NewEndpointLocator(region string) EndpointLocator {
	return &endpointLocator{region: region}
}

func (e *endpointLocator) AccountManagementEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["account"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Account Management endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) CertificateManagerEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["certificate-manager"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CERTIFICATE_MANAGER_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Certificate Manager Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) CFAPIEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cf"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CF_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Cloud Foundry endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) ContainerEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cs"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Container Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) SchematicsEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["schematics"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_SCHEMATICS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Schematics Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) ContainerRegistryEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cr"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CR_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Container Registry Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) CisEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cis"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CIS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Cis Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) GlobalSearchEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["global-search"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_GS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Global Search Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) GlobalTaggingEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["global-tagging"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_GT_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Global Tagging Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) IAMEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["iam"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("IAM endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) IAMPAPEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["iampap"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_IAMPAP_API_ENDPOINT"}, ep), nil

	}
	return "", fmt.Errorf("IAMPAP endpoint doesn't exist for region: %q", e.region)
}

func (e *endpointLocator) ICDEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["icd"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_ICD_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("ICD Service endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) MCCPAPIEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["mccp"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_MCCP_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("MCCP API endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) ResourceManagementEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["resource-manager"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Resource Management endpoint doesn't exist"))
}

func (e *endpointLocator) ResourceControllerEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["resource-controller"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Resource Controller endpoint doesn't exist"))
}

func (e *endpointLocator) ResourceCatalogEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["resource-catalog"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_RESOURCE_CATALOG_API_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("Resource Catalog endpoint doesn't exist"))
}

func (e *endpointLocator) UAAEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["uaa"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_UAA_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("UAA endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) CseEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["cse"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_CSE_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("CSE endpoint doesn't exist for region: %q", e.region))
}

func (e *endpointLocator) UserManagementEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["usermanagement"]["global"]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_USER_MANAGEMENT_ENDPOINT"}, ep), nil

	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("User Management endpoint doesn't exist"))
}

func (e *endpointLocator) HpcsEndpoint() (string, error) {
	if ep, ok := regionToEndpoint["hpcs"][e.region]; ok {
		//As the current list of regionToEndpoint above is not exhaustive we allow to read endpoints from the env
		return helpers.EnvFallBack([]string{"IBMCLOUD_HPCS_API_ENDPOINT"}, ep), nil
	}
	return "", bmxerror.New(ErrCodeServiceEndpoint, fmt.Sprintf("HPCS Service endpoint doesn't exist for region: %q", e.region))
}
