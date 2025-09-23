package main

import (
	"log"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/provider"
	"github.com/Mavrickk3/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
