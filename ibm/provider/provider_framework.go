package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

var (
	// Self-check to make sure the implementation satisfies the expected interface.
	_ provider.Provider = &IbmCloudProvider{}
)

type IbmCloudProvider struct {
	version       string
	clientSession any
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
			"ibmcloud_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud API Key",
			},
			"resource_group": schema.StringAttribute{
				Optional:    true,
				Description: "The Resource group id.",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
			},
			"ibmcloud_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
			},
			"max_retries": schema.Int64Attribute{
				Optional:    true,
				Description: "The retry count to set for API calls.",
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
	config.IBMCloudTimeout = GetInt64FromEnv(config.IBMCloudTimeout, []string{"IC_TIMEOUT", "IBMCLOUD_TIMEOUT"})
	config.Region = GetStringFromEnv(config.Region, []string{"IC_REGION", "IBMCLOUD_REGION", "BM_REGION", "BLUEMIX_REGION"})
	config.ResourceGroup = GetStringFromEnv(config.ResourceGroup, []string{"IC_RESOURCE_GROUP", "IBMCLOUD_RESOURCE_GROUP", "BM_RESOURCE_GROUP", "BLUEMIX_RESOURCE_GROUP"})
	config.MaxRetries = GetInt64FromEnv(config.MaxRetries, []string{"MAX_RETRIES"})

	resourceGroup := config.ResourceGroup.ValueString()
	region := config.Region.ValueString()
	retryCount := int(config.MaxRetries.ValueInt64())
	bluemixAPIKey := config.IBMCloudAPIKey.ValueString()
	bluemixTimeout := config.IBMCloudTimeout.ValueInt64()

	if bluemixAPIKey == "" {
		err := fmt.Errorf("IAM apikey not configured (IC_API_KEY, IBMCLOUD_API_KEY)")
		resp.Diagnostics.AddError("configuration error", err.Error())
		tflog.Debug(ctx, fmt.Sprintf("Error: %s\n", err.Error()))
		return
	}

	providerConfig := conns.Config{
		BluemixAPIKey:  bluemixAPIKey,
		Region:         region,
		ResourceGroup:  resourceGroup,
		BluemixTimeout: time.Duration(bluemixTimeout) * time.Second,
		RetryCount:     retryCount,
		RetryDelay:     RetryAPIDelay,
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
	resp.ResourceData = clientSession
	resp.DataSourceData = clientSession

	// Store provider's metadata so that acceptance tests can access it.
	TestProviderClientSession = clientSession
}

func (p *IbmCloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add new datasources here
	}
}

func (p *IbmCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		// Add new resources here
	}
}

type ProviderFrameworkModel struct {
	IBMCloudAPIKey  types.String `tfsdk:"ibmcloud_api_key"`
	IBMCloudTimeout types.Int64  `tfsdk:"ibmcloud_timeout"`
	ResourceGroup   types.String `tfsdk:"resource_group"`
	Region          types.String `tfsdk:"region"`
	MaxRetries      types.Int64  `tfsdk:"max_retries"`
}

// TestAccProviderClientSession is used by acceptance tests. It is initialized during
// the initialization of the IBM provider.
var TestProviderClientSession any
