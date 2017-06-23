package main

import (
	"log"

	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ibm/ibm"
)

func main() {
	log.Println("IBM Cloud Provider version", Version, VersionPrerelease, GitCommit)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ibm.Provider,
	})
}
