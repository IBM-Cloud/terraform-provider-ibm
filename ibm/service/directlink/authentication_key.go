// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import "github.com/IBM/networking-go-sdk/directlinkv1"

func newAuthenticationKeyIdentity(crn string) directlinkv1.AuthenticationKeyIdentityIntf {
	return &directlinkv1.AuthenticationKeyIdentity{Crn: &crn}
}

func authenticationKeyCRN(authenticationKey directlinkv1.AuthenticationKeyReferenceIntf) (string, bool) {
	var crn *string

	switch key := authenticationKey.(type) {
	case *directlinkv1.AuthenticationKeyReference:
		crn = key.Crn
	case *directlinkv1.AuthenticationKeyReferenceHpcsAuthenticationKeyReference:
		crn = key.Crn
	case *directlinkv1.AuthenticationKeyReferenceKeyProtectAuthenticationKeyReference:
		crn = key.Crn
	case *directlinkv1.AuthenticationKeyReferenceSecretsManagerAuthenticationKeyReference:
		crn = key.Crn
	}

	if crn == nil {
		return "", false
	}
	return *crn, true
}
