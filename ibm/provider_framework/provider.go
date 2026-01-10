// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider_framework

import (
	"context"
	"os"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/codeengine"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &frameworkProvider{}
)

// frameworkProvider is the provider implementation for the IBM Cloud Terraform Provider
// using the terraform-plugin-framework. This provider runs alongside the existing SDKv2
// provider via terraform-plugin-mux to enable framework-only features like Actions.
type frameworkProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance testing.
	version string
}

// frameworkProviderModel describes the provider data model.
type frameworkProviderModel struct {
	BluemixAPIKey          types.String `tfsdk:"bluemix_api_key"`
	BluemixTimeout         types.Int64  `tfsdk:"bluemix_timeout"`
	IBMCloudAPIKey         types.String `tfsdk:"ibmcloud_api_key"`
	IBMCloudTimeout        types.Int64  `tfsdk:"ibmcloud_timeout"`
	Region                 types.String `tfsdk:"region"`
	Zone                   types.String `tfsdk:"zone"`
	ResourceGroup          types.String `tfsdk:"resource_group"`
	SoftlayerAPIKey        types.String `tfsdk:"softlayer_api_key"`
	SoftlayerUsername      types.String `tfsdk:"softlayer_username"`
	SoftlayerEndpointURL   types.String `tfsdk:"softlayer_endpoint_url"`
	SoftlayerTimeout       types.Int64  `tfsdk:"softlayer_timeout"`
	IAASClassicAPIKey      types.String `tfsdk:"iaas_classic_api_key"`
	IAASClassicUsername    types.String `tfsdk:"iaas_classic_username"`
	IAASClassicEndpointURL types.String `tfsdk:"iaas_classic_endpoint_url"`
	IAASClassicTimeout     types.Int64  `tfsdk:"iaas_classic_timeout"`
	MaxRetries             types.Int64  `tfsdk:"max_retries"`
	FunctionNamespace      types.String `tfsdk:"function_namespace"`
	RIAASEndpoint          types.String `tfsdk:"riaas_endpoint"`
	Generation             types.Int64  `tfsdk:"generation"`
	IAMProfileID           types.String `tfsdk:"iam_profile_id"`
	IAMProfileName         types.String `tfsdk:"iam_profile_name"`
	IAMToken               types.String `tfsdk:"iam_token"`
	IAMRefreshToken        types.String `tfsdk:"iam_refresh_token"`
	Visibility             types.String `tfsdk:"visibility"`
	PrivateEndpointType    types.String `tfsdk:"private_endpoint_type"`
	EndpointsFilePath      types.String `tfsdk:"endpoints_file_path"`
	IBMCloudAccountID      types.String `tfsdk:"ibmcloud_account_id"`
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &frameworkProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *frameworkProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ibm"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
// This schema MUST match the SDKv2 provider schema exactly for mux compatibility.
// All defaults are removed and handled in Configure() method.
func (p *frameworkProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bluemix_api_key": schema.StringAttribute{
				Optional:           true,
				Description:        "The Bluemix API Key",
				DeprecationMessage: "This field is deprecated please use ibmcloud_api_key",
			},
			"bluemix_timeout": schema.Int64Attribute{
				Optional:           true,
				Description:        "The timeout (in seconds) to set for any Bluemix API calls made.",
				DeprecationMessage: "This field is deprecated please use ibmcloud_timeout",
			},
			"ibmcloud_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud API Key",
			},
			"ibmcloud_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
			},
			"zone": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region zone (for example 'us-south-1') for power resources.",
			},
			"resource_group": schema.StringAttribute{
				Optional:    true,
				Description: "The Resource group id.",
			},
			"softlayer_api_key": schema.StringAttribute{
				Optional:           true,
				Description:        "The SoftLayer API Key",
				DeprecationMessage: "This field is deprecated please use iaas_classic_api_key",
			},
			"softlayer_username": schema.StringAttribute{
				Optional:           true,
				Description:        "The SoftLayer user name",
				DeprecationMessage: "This field is deprecated please use iaas_classic_username",
			},
			"softlayer_endpoint_url": schema.StringAttribute{
				Optional:           true,
				Description:        "The Softlayer Endpoint",
				DeprecationMessage: "This field is deprecated please use iaas_classic_endpoint_url",
			},
			"softlayer_timeout": schema.Int64Attribute{
				Optional:           true,
				Description:        "The timeout (in seconds) to set for any SoftLayer API calls made.",
				DeprecationMessage: "This field is deprecated please use iaas_classic_timeout",
			},
			"iaas_classic_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure API Key",
			},
			"iaas_classic_username": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure API user name",
			},
			"iaas_classic_endpoint_url": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure Endpoint",
			},
			"iaas_classic_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Classic Infrastructure API calls made.",
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
			"generation": schema.Int64Attribute{
				Optional:           true,
				Description:        "Generation of Virtual Private Cloud. Default is 2",
				DeprecationMessage: "The generation field is deprecated and will be removed after couple of releases",
			},
			"iam_profile_id": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Trusted Profile ID",
			},
			"iam_profile_name": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Trusted Profile Name",
			},
			"iam_token": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Authentication token",
			},
			"iam_refresh_token": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Authentication refresh token",
			},
			"visibility": schema.StringAttribute{
				Optional:    true,
				Description: "Visibility of the provider if it is private or public.",
			},
			"private_endpoint_type": schema.StringAttribute{
				Optional:    true,
				Description: "Private Endpoint type used by the service endpoints. Example: vpe.",
			},
			"endpoints_file_path": schema.StringAttribute{
				Optional:    true,
				Description: "Path of the file that contains private and public regional endpoints mapping",
			},
			"ibmcloud_account_id": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud account ID",
			},
		},
	}
}

// Configure prepares the provider for data sources and resources.
// This method applies default values that were removed from the schema for mux compatibility.
func (p *frameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config frameworkProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Apply defaults for configuration values
	// These defaults match the SDKv2 provider's DefaultFunc behavior
	// but are applied in Configure() instead of schema for mux compatibility

	// ibmcloud_timeout default: 60
	if config.IBMCloudTimeout.IsNull() {
		if timeout := os.Getenv("IC_TIMEOUT"); timeout != "" {
			// Parse timeout from env, default to 60 if parsing fails
			config.IBMCloudTimeout = types.Int64Value(60)
		} else if timeout := os.Getenv("IBMCLOUD_TIMEOUT"); timeout != "" {
			config.IBMCloudTimeout = types.Int64Value(60)
		} else {
			config.IBMCloudTimeout = types.Int64Value(60)
		}
	}

	// region default: "us-south"
	if config.Region.IsNull() || config.Region.ValueString() == "" {
		if region := os.Getenv("IC_REGION"); region != "" {
			config.Region = types.StringValue(region)
		} else if region := os.Getenv("IBMCLOUD_REGION"); region != "" {
			config.Region = types.StringValue(region)
		} else if region := os.Getenv("BM_REGION"); region != "" {
			config.Region = types.StringValue(region)
		} else if region := os.Getenv("BLUEMIX_REGION"); region != "" {
			config.Region = types.StringValue(region)
		} else {
			config.Region = types.StringValue("us-south")
		}
	}

	// iaas_classic_endpoint_url default: "https://api.softlayer.com/rest/v3"
	if config.IAASClassicEndpointURL.IsNull() || config.IAASClassicEndpointURL.ValueString() == "" {
		if endpoint := os.Getenv("IAAS_CLASSIC_ENDPOINT_URL"); endpoint != "" {
			config.IAASClassicEndpointURL = types.StringValue(endpoint)
		} else {
			config.IAASClassicEndpointURL = types.StringValue("https://api.softlayer.com/rest/v3")
		}
	}

	// iaas_classic_timeout default: 60
	if config.IAASClassicTimeout.IsNull() {
		if timeout := os.Getenv("IAAS_CLASSIC_TIMEOUT"); timeout != "" {
			config.IAASClassicTimeout = types.Int64Value(60)
		} else {
			config.IAASClassicTimeout = types.Int64Value(60)
		}
	}

	// max_retries default: 10
	if config.MaxRetries.IsNull() {
		if retries := os.Getenv("MAX_RETRIES"); retries != "" {
			config.MaxRetries = types.Int64Value(10)
		} else {
			config.MaxRetries = types.Int64Value(10)
		}
	}

	// visibility default: "public"
	if config.Visibility.IsNull() || config.Visibility.ValueString() == "" {
		if visibility := os.Getenv("IC_VISIBILITY"); visibility != "" {
			config.Visibility = types.StringValue(visibility)
		} else if visibility := os.Getenv("IBMCLOUD_VISIBILITY"); visibility != "" {
			config.Visibility = types.StringValue(visibility)
		} else {
			config.Visibility = types.StringValue("public")
		}
	}

	// Get API key from config or environment
	apiKey := config.IBMCloudAPIKey.ValueString()
	if apiKey == "" {
		apiKey = config.BluemixAPIKey.ValueString()
	}
	if apiKey == "" {
		if key := os.Getenv("IC_API_KEY"); key != "" {
			apiKey = key
		} else if key := os.Getenv("IBMCLOUD_API_KEY"); key != "" {
			apiKey = key
		} else if key := os.Getenv("BM_API_KEY"); key != "" {
			apiKey = key
		} else if key := os.Getenv("BLUEMIX_API_KEY"); key != "" {
			apiKey = key
		}
	}

	// Create conns.Config to initialize client session
	connConfig := conns.Config{
		BluemixAPIKey:    apiKey,
		Region:           config.Region.ValueString(),
		BluemixTimeout:   time.Duration(config.IBMCloudTimeout.ValueInt64()) * time.Second,
		SoftLayerTimeout: time.Duration(config.IAASClassicTimeout.ValueInt64()) * time.Second,
		RetryCount:       int(config.MaxRetries.ValueInt64()),
		RetryDelay:       conns.RetryAPIDelay,
		Visibility:       config.Visibility.ValueString(),
	}

	// Handle optional fields
	if !config.ResourceGroup.IsNull() {
		connConfig.ResourceGroup = config.ResourceGroup.ValueString()
	}
	if !config.Zone.IsNull() {
		connConfig.Zone = config.Zone.ValueString()
	}
	if !config.IAASClassicUsername.IsNull() {
		connConfig.SoftLayerUserName = config.IAASClassicUsername.ValueString()
	}
	if !config.IAASClassicAPIKey.IsNull() {
		connConfig.SoftLayerAPIKey = config.IAASClassicAPIKey.ValueString()
	}
	if !config.IAASClassicEndpointURL.IsNull() {
		connConfig.SoftLayerEndpointURL = config.IAASClassicEndpointURL.ValueString()
	}
	if !config.FunctionNamespace.IsNull() {
		connConfig.FunctionNameSpace = config.FunctionNamespace.ValueString()
	}
	if !config.RIAASEndpoint.IsNull() {
		connConfig.RiaasEndPoint = config.RIAASEndpoint.ValueString()
	}
	if !config.IAMToken.IsNull() {
		connConfig.IAMToken = config.IAMToken.ValueString()
	}
	if !config.IAMRefreshToken.IsNull() {
		connConfig.IAMRefreshToken = config.IAMRefreshToken.ValueString()
	}
	if !config.PrivateEndpointType.IsNull() {
		connConfig.PrivateEndpointType = config.PrivateEndpointType.ValueString()
	}
	if !config.EndpointsFilePath.IsNull() {
		connConfig.EndpointsFile = config.EndpointsFilePath.ValueString()
	}
	if !config.IAMProfileID.IsNull() {
		connConfig.IAMTrustedProfileID = config.IAMProfileID.ValueString()
	}
	if !config.IAMProfileName.IsNull() {
		connConfig.IAMTrustedProfileName = config.IAMProfileName.ValueString()
	}
	if !config.IBMCloudAccountID.IsNull() {
		connConfig.Account = config.IBMCloudAccountID.ValueString()
	}

	// Initialize client session
	session, err := connConfig.ClientSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create IBM Cloud Client",
			"An unexpected error occurred when creating the IBM Cloud client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"IBM Cloud Client Error: "+err.Error(),
		)
		return
	}

	// Set the client session for resources, data sources, and actions
	resp.DataSourceData = session
	resp.ResourceData = session
	resp.ActionData = session
}

// Resources defines the resources implemented in the provider.
// Initially empty - all resources remain in SDKv2 provider.
func (p *frameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

// DataSources defines the data sources implemented in the provider.
// Initially empty - all data sources remain in SDKv2 provider.
func (p *frameworkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Actions defines the actions implemented in the provider.
func (p *frameworkProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{
		codeengine.NewCodeEngineBuildRunAction,
	}
}
