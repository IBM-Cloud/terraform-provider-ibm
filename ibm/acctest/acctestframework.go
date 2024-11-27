// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func init() {
	testlogger := os.Getenv("TF_LOG")
	if testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}
}

var (
	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can reattach.
	ProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error) = protoV6ProviderFactoriesInit(context.Background(), ProviderName)

	// testAccProviderConfigure ensures Provider is only configured once
)

// testAccProviderConfigure ensures Provider is only configured once
//
// The PreCheck(t) function is invoked for every test and this prevents
// extraneous reconfiguration to the same values each time. However, this does
// not prevent reconfiguration that may happen should the address of
// Provider be errantly reused in ProviderFactories.
var testAccProviderFrameworkConfigure sync.Once

func protoV6ProviderFactoriesInit(ctx context.Context, providerNames ...string) map[string]func() (tfprotov6.ProviderServer, error) {
	// Initialize logging
	log.Printf("[INFO] UJJK context is %v", vpc.BeautifyResponse(ctx))
	if testlogger := os.Getenv("TF_LOG"); testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}
	factories := make(map[string]func() (tfprotov6.ProviderServer, error), len(providerNames))

	for _, name := range providerNames {
		factories[name] = func() (tfprotov6.ProviderServer, error) {
			providerServerFactory, _, err := provider.ProtoV6ProviderServerFactory(ctx)

			if err != nil {
				return nil, err
			}

			return providerServerFactory(), nil
		}
	}

	return factories
}
