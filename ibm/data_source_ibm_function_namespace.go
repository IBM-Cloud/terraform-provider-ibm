/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/namespace-go-sdk/ibmcloudfunctionsnamespaceapiv1"
)

func dataSourceIBMFunctionNamespace() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMFunctionNamespaceRead,
		Schema: map[string]*schema.Schema{
			funcNamespaceName: {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of namespace.",
				ValidateFunc: InvokeValidator("ibm_function_namespace", funcNamespaceName),
			},
			funcNamespaceDesc: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Namespace Description.",
			},
			funcNamespaceResGrpId: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource Group ID.",
			},
			funcNamespaceLoc: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Namespace Location.",
			},
		},
	}
}

func dataSourceIBMFunctionNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	nsClient, err := meta.(ClientSession).IAMNamespaceAPI()
	if err != nil {
		return err
	}

	namespaceOptions := &ibmcloudfunctionsnamespaceapiv1.GetNamespacesOptions{}
	nsList, _, err := nsClient.GetNamespaces(namespaceOptions)
	if err != nil {
		return nil
	}

	name := d.Get("name").(string)
	for _, n := range nsList.Namespaces {
		if n.Name != nil && *n.Name == name {
			getOptions := &ibmcloudfunctionsnamespaceapiv1.GetNamespaceOptions{
				ID: n.ID,
			}

			instance, response, err := nsClient.GetNamespace(getOptions)
			if err != nil {
				if response != nil && response.StatusCode == 404 {
					d.SetId("")
					return nil
				}
			}

			if instance.ID != nil {
				d.SetId(*instance.ID)
			}

			if instance.Name != nil {
				d.Set(funcNamespaceName, *instance.Name)
			}

			if instance.ResourceGroupID != nil {
				d.Set(funcNamespaceResGrpId, *instance.ResourceGroupID)
			}

			if instance.Location != nil {
				d.Set(funcNamespaceLoc, *instance.Location)
			}

			if instance.Description != nil {
				d.Set(funcNamespaceDesc, *instance.Description)
			}

			return nil
		}
	}

	return fmt.Errorf("No cloud function namespace found with name [%s]", name)
}
