// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package provider

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
)

// Ensure the implementation satisfies the provider.Provider interface.
var _ provider.Provider = &IbmCloudProvider{}

type IbmCloudProvider struct {
	version       string
	clientSession interface{}
}

func NewFrameworkProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &IbmCloudProvider{
			version: version,
		}
	}
}

func (p *IbmCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ibm"
	resp.Version = p.version
}

func (p *IbmCloudProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bluemix_api_key": schema.StringAttribute{
				Optional:           true,
				Description:        "The Bluemix API Key",
				DeprecationMessage: "This field is deprecated please use ibmcloud_api_key",
			},
			"ibmcloud_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud API Key",
			},
			"ibmcloud_account_id": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud account ID",
			},
			"iam_token": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Authentication token",
			},
			"iam_refresh_token": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Authentication refresh token",
			},
			"iam_profile_id": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Trusted Profile ID",
			},
			"iam_profile_name": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Trusted Profile Name",
			},
			"softlayer_username": schema.StringAttribute{
				Optional:           true,
				Description:        "The SoftLayer user name",
				DeprecationMessage: "This field is deprecated please use iaas_classic_username",
			},
			"iaas_classic_username": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure API user name",
			},
			"softlayer_api_key": schema.StringAttribute{
				Optional:           true,
				Description:        "The SoftLayer API Key",
				DeprecationMessage: "This field is deprecated please use iaas_classic_api_key",
			},
			"iaas_classic_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure API Key",
			},
			"softlayer_endpoint_url": schema.StringAttribute{
				Optional:           true,
				Description:        "The Softlayer Endpoint",
				DeprecationMessage: "This field is deprecated please use iaas_classic_endpoint_url",
			},
			"iaas_classic_endpoint_url": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure Endpoint",
			},
			"softlayer_timeout": schema.Int64Attribute{
				Optional:           true,
				Description:        "The timeout (in seconds) to set for any SoftLayer API calls made.",
				DeprecationMessage: "This field is deprecated please use iaas_classic_timeout",
			},
			"generation": schema.Int64Attribute{
				Optional:           true,
				Description:        "Generation of Virtual Private Cloud. Default is 2",
				DeprecationMessage: "The generation field is deprecated and will be removed after couple of releases",
			},
			"iaas_classic_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Classic Infrastructure API calls made.",
			},
			"bluemix_timeout": schema.Int64Attribute{
				Optional:           true,
				Description:        "The timeout (in seconds) to set for any Bluemix API calls made.",
				DeprecationMessage: "This field is deprecated please use ibmcloud_timeout",
			},
			"ibmcloud_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
			},
			"visibility": schema.StringAttribute{
				Optional:    true,
				Description: "Visibility of the provider if it is private or public.",
				Validators: []validator.String{
					stringvalidator.OneOf("public", "private", "public-and-private"),
				},
			},
			"endpoints_file_path": schema.StringAttribute{
				Optional:    true,
				Description: "Path of the file that contains private and public regional endpoints mapping",
			},
			"private_endpoint_type": schema.StringAttribute{
				Optional:    true,
				Description: "Private Endpoint type used by the service endpoints. Example: vpe.",
				Validators: []validator.String{
					stringvalidator.OneOf("vpe"),
				},
			},
			"resource_group": schema.StringAttribute{
				Optional:    true,
				Description: "The Resource group id.",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
			},
			"zone": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region zone (for example 'us-south-1') for power resources.",
			},
			"max_retries": schema.Int64Attribute{
				Optional:    true,
				Description: "The retry count to set for API calls.",
			},
			"function_namespace": schema.StringAttribute{
				Optional:           true,
				Description:        "The IBM Cloud Function namespace",
				DeprecationMessage: "This field will be deprecated soon",
			},
			"riaas_endpoint": schema.StringAttribute{
				Optional:           true,
				Description:        "The next generation infrastructure service endpoint url.",
				DeprecationMessage: "This field is deprecated use generation",
			},
		},
	}
}

func (p *IbmCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config ProviderFrameworkModel
	if diags := req.Config.Get(ctx, &config); diags.HasError() {
		resp.Diagnostics = diags
		return
	}
	config.IBMCloudAPIKey = GetStringFromEnv(config.IBMCloudAPIKey, []string{"IC_API_KEY", "IBMCLOUD_API_KEY"})
	config.IBMCloudAccountID = GetStringFromEnv(config.IBMCloudAccountID, []string{"IC_ACCOUNT_ID", "IBMCLOUD_ACCOUNT_ID"})
	config.IBMCloudTimeout = GetInt64FromEnv(config.IBMCloudTimeout, []string{"IC_TIMEOUT", "IBMCLOUD_TIMEOUT"})
	config.Region = GetStringFromEnv(config.Region, []string{"IC_REGION", "IBMCLOUD_REGION", "BM_REGION", "BLUEMIX_REGION"})
	config.Zone = GetStringFromEnv(config.Zone, []string{"IC_ZONE", "IBMCLOUD_ZONE"})
	config.ResourceGroup = GetStringFromEnv(config.ResourceGroup, []string{"IC_RESOURCE_GROUP", "IBMCLOUD_RESOURCE_GROUP", "BM_RESOURCE_GROUP", "BLUEMIX_RESOURCE_GROUP"})
	config.IAASClassicAPIKey = GetStringFromEnv(config.IAASClassicAPIKey, []string{"IAAS_CLASSIC_API_KEY"})
	config.IAASClassicUsername = GetStringFromEnv(config.IAASClassicUsername, []string{"IAAS_CLASSIC_USERNAME"})
	config.IAASClassicEndpointURL = GetStringFromEnv(config.IAASClassicEndpointURL, []string{"IAAS_CLASSIC_ENDPOINT_URL"})
	config.IAASClassicTimeout = GetInt64FromEnv(config.IAASClassicTimeout, []string{"IAAS_CLASSIC_TIMEOUT"})
	config.MaxRetries = GetInt64FromEnv(config.MaxRetries, []string{"MAX_RETRIES"})
	config.IAMProfileID = GetStringFromEnv(config.IAMProfileID, []string{"IC_IAM_PROFILE_ID", "IBMCLOUD_IAM_PROFILE_ID"})
	config.IAMProfileName = GetStringFromEnv(config.IAMProfileName, []string{"IC_IAM_PROFILE_NAME", "IBMCLOUD_IAM_PROFILE_NAME"})
	config.IAMToken = GetStringFromEnv(config.IAMToken, []string{"IC_IAM_TOKEN", "IBMCLOUD_IAM_TOKEN"})
	config.IAMRefreshToken = GetStringFromEnv(config.IAMRefreshToken, []string{"IC_IAM_REFRESH_TOKEN", "IBMCLOUD_IAM_REFRESH_TOKEN"})
	config.Visibility = GetStringFromEnv(config.Visibility, []string{"IC_VISIBILITY", "IBMCLOUD_VISIBILITY"})
	config.EndpointsFilePath = GetStringFromEnv(config.EndpointsFilePath, []string{"IC_ENDPOINTS_FILE_PATH", "IBMCLOUD_ENDPOINTS_FILE_PATH"})
	config.PrivateEndpointType = GetStringFromEnv(config.PrivateEndpointType, []string{"IC_PRIVATE_ENDPOINT_TYPE", "IBMCLOUD_PRIVATE_ENDPOINT_TYPE"})

	bluemixAPIKey := config.BluemixAPIKey.ValueString()
	if bluemixAPIKey == "" {
		bluemixAPIKey = config.IBMCloudAPIKey.ValueString()
	}

	iamToken := config.IAMToken.ValueString()
	iamRefreshToken := config.IAMRefreshToken.ValueString()
	iamTrustedProfileId := config.IAMProfileID.ValueString()
	iamTrustedProfileName := config.IAMProfileName.ValueString()
	accountID := config.IBMCloudAccountID.ValueString()

	softlayerUsername := config.SoftLayerUsername.ValueString()
	if softlayerUsername == "" {
		softlayerUsername = config.IAASClassicUsername.ValueString()
	}

	softlayerAPIKey := config.SoftLayerAPIKey.ValueString()
	if softlayerAPIKey == "" {
		softlayerAPIKey = config.IAASClassicAPIKey.ValueString()
	}

	softlayerEndpointUrl := config.SoftLayerEndpointURL.ValueString()
	if softlayerEndpointUrl == "" {
		softlayerEndpointUrl = config.IAASClassicEndpointURL.ValueString()
	}

	softlayerTimeout := int(config.SoftLayerTimeout.ValueInt64())
	if softlayerTimeout == 0 {
		softlayerTimeout = int(config.IAASClassicTimeout.ValueInt64())
	}

	bluemixTimeout := int(config.BluemixTimeout.ValueInt64())
	if bluemixTimeout == 0 {
		bluemixTimeout = int(config.IBMCloudTimeout.ValueInt64())
	}

	visibility := config.Visibility.ValueString()
	endpointsFilePath := config.EndpointsFilePath.ValueString()
	privateEndpointType := config.PrivateEndpointType.ValueString()
	resourceGroup := config.ResourceGroup.ValueString()
	region := config.Region.ValueString()
	zone := config.Zone.ValueString()
	retryCount := int(config.MaxRetries.ValueInt64())
	functionNamespace := config.FunctionNamespace.ValueString()
	riaasEndpoint := config.RIAASEndpoint.ValueString()

	if functionNamespace == "" {
		functionNamespace = os.Getenv("FUNCTION_NAMESPACE")
	}
	if functionNamespace != "" {
		os.Setenv("FUNCTION_NAMESPACE", functionNamespace)
	}

	providerConfig := conns.Config{
		BluemixAPIKey:         bluemixAPIKey,
		Region:                region,
		ResourceGroup:         resourceGroup,
		BluemixTimeout:        time.Duration(bluemixTimeout) * time.Second,
		SoftLayerTimeout:      time.Duration(softlayerTimeout) * time.Second,
		SoftLayerUserName:     softlayerUsername,
		SoftLayerAPIKey:       softlayerAPIKey,
		RetryCount:            retryCount,
		SoftLayerEndpointURL:  softlayerEndpointUrl,
		RetryDelay:            conns.RetryAPIDelay,
		FunctionNameSpace:     functionNamespace,
		RiaasEndPoint:         riaasEndpoint,
		IAMToken:              iamToken,
		IAMRefreshToken:       iamRefreshToken,
		Zone:                  zone,
		Visibility:            visibility,
		EndpointsFile:         endpointsFilePath,
		PrivateEndpointType:   privateEndpointType,
		IAMTrustedProfileID:   iamTrustedProfileId,
		IAMTrustedProfileName: iamTrustedProfileName,
		Account:               accountID,
	}

	clientSession, err := providerConfig.ClientSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client session",
			fmt.Sprintf("Unable to create client session: %v", err),
		)
		return
	}
	p.clientSession = clientSession
}

func (p *IbmCloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add your data sources here
		func() datasource.DataSource {
			return vpc.NewIsSshKeyDataSource(p.clientSession)
		},
	}
}

func (p *IbmCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource {
			return vpc.NewSSHKeyResource(p.clientSession)
		},
	}

}

type ProviderFrameworkModel struct {
	BluemixAPIKey          types.String `tfsdk:"bluemix_api_key"`
	IBMCloudAPIKey         types.String `tfsdk:"ibmcloud_api_key"`
	IBMCloudAccountID      types.String `tfsdk:"ibmcloud_account_id"`
	IAMToken               types.String `tfsdk:"iam_token"`
	IAMRefreshToken        types.String `tfsdk:"iam_refresh_token"`
	IAMProfileID           types.String `tfsdk:"iam_profile_id"`
	IAMProfileName         types.String `tfsdk:"iam_profile_name"`
	SoftLayerUsername      types.String `tfsdk:"softlayer_username"`
	IAASClassicUsername    types.String `tfsdk:"iaas_classic_username"`
	SoftLayerAPIKey        types.String `tfsdk:"softlayer_api_key"`
	IAASClassicAPIKey      types.String `tfsdk:"iaas_classic_api_key"`
	SoftLayerEndpointURL   types.String `tfsdk:"softlayer_endpoint_url"`
	IAASClassicEndpointURL types.String `tfsdk:"iaas_classic_endpoint_url"`
	SoftLayerTimeout       types.Int64  `tfsdk:"softlayer_timeout"`
	IAASClassicTimeout     types.Int64  `tfsdk:"iaas_classic_timeout"`
	BluemixTimeout         types.Int64  `tfsdk:"bluemix_timeout"`
	IBMCloudTimeout        types.Int64  `tfsdk:"ibmcloud_timeout"`
	Visibility             types.String `tfsdk:"visibility"`
	EndpointsFilePath      types.String `tfsdk:"endpoints_file_path"`
	PrivateEndpointType    types.String `tfsdk:"private_endpoint_type"`
	ResourceGroup          types.String `tfsdk:"resource_group"`
	Region                 types.String `tfsdk:"region"`
	Zone                   types.String `tfsdk:"zone"`
	MaxRetries             types.Int64  `tfsdk:"max_retries"`
	Generation             types.Int64  `tfsdk:"generation"`
	FunctionNamespace      types.String `tfsdk:"function_namespace"`
	RIAASEndpoint          types.String `tfsdk:"riaas_endpoint"`
}
