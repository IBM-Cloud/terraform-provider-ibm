// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Used for retry logic on resource timeout.
func isResourceTimeoutError(err error) bool {
	timeoutErr, ok := err.(*resource.TimeoutError)
	return ok && timeoutErr.LastError == nil
}
