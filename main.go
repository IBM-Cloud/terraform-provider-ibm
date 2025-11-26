package main

import (
	"context"
	"flag"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-mux/tf5to6server"
	"github.com/hashicorp/terraform-plugin-mux/tf6muxserver"
)

func main() {
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)
	var debugMode bool
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()

	upgradedSdkServer, err := tf5to6server.UpgradeServer(
		ctx,
		provider.New(version.Version)().GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	frameworkServer := providerserver.NewProtocol6(provider.NewFrameworkProvider(version.Version)())

	muxServer, err := tf6muxserver.NewMuxServer(
		ctx,
		func() tfprotov6.ProviderServer { return upgradedSdkServer },
		frameworkServer,
	)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf6server.ServeOpt

	if debugMode {
		serveOpts = append(serveOpts, tf6server.WithManagedDebug())
	}

	err = tf6server.Serve(
		"registry.terraform.io/IBM-Cloud/ibm",
		muxServer.ProviderServer,
		serveOpts...,
	)
	if err != nil {
		log.Fatal(err)
	}

}
