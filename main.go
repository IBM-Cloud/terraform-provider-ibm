package main

import (
	"context"
	"flag"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the IBM terraform provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: ibm.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/IBM-Cloud/ibm", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
