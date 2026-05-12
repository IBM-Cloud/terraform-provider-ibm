package main

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider_framework"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
)

func main() {
	ctx := context.Background()
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)

	// Upgrade the SDKv2 provider to protocol version 6
	upgradedSdkProvider, err := tf5to6server.UpgradeServer(
		ctx,
		provider.Provider().GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create the framework provider server
	// New() returns a factory function, so we call it to get the provider
	frameworkProviderServer := providerserver.NewProtocol6(provider_framework.New(version.Version)())

	// Create mux server combining both providers
	muxServer, err := tf6muxserver.NewMuxServer(ctx,
		func() tfprotov6.ProviderServer {
			return upgradedSdkProvider
		},
		frameworkProviderServer,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the mux server
	err = tf6server.Serve(
		"registry.terraform.io/IBM-Cloud/ibm",
		muxServer.ProviderServer,
	)
	if err != nil {
		log.Fatal(err)
	}
}
