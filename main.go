package main

import (
	"context"
	"flag"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)

	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		opts := &plugin.ServeOpts{ProviderFunc: provider.Provider}
		// TODO: update this string with the full name of your provider as used in your configs
		err := plugin.Debug(context.Background(), "ibm.com/IBM-Cloud/ibm", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
