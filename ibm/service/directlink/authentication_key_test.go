// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"testing"

	"github.com/IBM/networking-go-sdk/directlinkv1"
)

func TestNewAuthenticationKeyIdentity(t *testing.T) {
	t.Parallel()

	const crn = "crn:v1:bluemix:public:kms:us-south:a/account:instance:key:key-id"

	identity, ok := newAuthenticationKeyIdentity(crn).(*directlinkv1.AuthenticationKeyIdentity)
	if !ok {
		t.Fatal("newAuthenticationKeyIdentity returned an unexpected identity type")
	}
	if identity.Crn == nil || *identity.Crn != crn {
		t.Fatalf("newAuthenticationKeyIdentity CRN = %v, want %q", identity.Crn, crn)
	}
}

func TestAuthenticationKeyCRN(t *testing.T) {
	t.Parallel()

	const crn = "crn:v1:bluemix:public:kms:us-south:a/account:instance:key:key-id"

	tests := []struct {
		name              string
		authenticationKey directlinkv1.AuthenticationKeyReferenceIntf
		wantCRN           string
		wantOK            bool
	}{
		{
			name:              "nil reference",
			authenticationKey: nil,
		},
		{
			name:              "base reference",
			authenticationKey: &directlinkv1.AuthenticationKeyReference{Crn: stringPointer(crn)},
			wantCRN:           crn,
			wantOK:            true,
		},
		{
			name:              "key protect reference",
			authenticationKey: &directlinkv1.AuthenticationKeyReferenceKeyProtectAuthenticationKeyReference{Crn: stringPointer(crn)},
			wantCRN:           crn,
			wantOK:            true,
		},
		{
			name:              "HPCS reference",
			authenticationKey: &directlinkv1.AuthenticationKeyReferenceHpcsAuthenticationKeyReference{Crn: stringPointer(crn)},
			wantCRN:           crn,
			wantOK:            true,
		},
		{
			name:              "secrets manager reference",
			authenticationKey: &directlinkv1.AuthenticationKeyReferenceSecretsManagerAuthenticationKeyReference{Crn: stringPointer(crn)},
			wantCRN:           crn,
			wantOK:            true,
		},
		{
			name:              "reference without CRN",
			authenticationKey: &directlinkv1.AuthenticationKeyReference{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotCRN, gotOK := authenticationKeyCRN(test.authenticationKey)
			if gotCRN != test.wantCRN || gotOK != test.wantOK {
				t.Fatalf("authenticationKeyCRN() = (%q, %t), want (%q, %t)", gotCRN, gotOK, test.wantCRN, test.wantOK)
			}
		})
	}
}

func stringPointer(value string) *string {
	return &value
}
