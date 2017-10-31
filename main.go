package main

import (
	"log"

	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ibm/ibm"
	"github.com/terraform-providers/terraform-provider-ibm/version"
)

func main() {
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ibm.Provider,
	})
}
