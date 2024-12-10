// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package provider

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ProtoV6ProviderServerFactory returns a muxed terraform-plugin-go protocol v6 provider factory function.
// This factory function is suitable for use with the terraform-plugin-go Serve function.
// The primary (Plugin SDK) provider server is also returned (useful for testing).
func ProtoV6ProviderServerFactory(ctx context.Context) (func() tfprotov6.ProviderServer, *schema.Provider, error) {
	primary := New(version.Version)()
	upgradedSdkServer, err := tf5to6server.UpgradeServer(
		ctx,
		primary.GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	frameworkServer := providerserver.NewProtocol6(NewFrameworkProvider(version.Version)())

	muxServer, err := tf6muxserver.NewMuxServer(
		ctx,
		func() tfprotov6.ProviderServer { return upgradedSdkServer },
		frameworkServer,
	)
	if err != nil {
		log.Fatal(err)
	}
	return muxServer.ProviderServer, primary, nil
}
