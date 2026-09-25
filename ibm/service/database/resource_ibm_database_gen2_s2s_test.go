// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

// s2sAuthErr is the rejection the provisioning API returns while the service-to-service
// authorization for the encryption key is not yet visible to the database service.
var s2sAuthErr = errors.New("Missing or mis-configured S2S Authorization Policy. The service could not " +
	"configure access to the encryption key specified as crn:v1:bluemix:public:kms:us-south:a/acct:inst:key:kid")

// withInstantS2SRetrySleep removes the real backoff so the retry loop runs at test speed.
func withInstantS2SRetrySleep(t *testing.T) {
	t.Helper()

	original := s2sAuthRetrySleep
	s2sAuthRetrySleep = func(time.Duration) {}
	t.Cleanup(func() { s2sAuthRetrySleep = original })
}

func TestIsS2SAuthorizationError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		response *core.DetailedResponse
		expected bool
	}{
		{
			name:     "nil error is never retryable",
			err:      nil,
			expected: false,
		},
		{
			name:     "message on the error",
			err:      s2sAuthErr,
			expected: true,
		},
		{
			name:     "message only in the response body",
			err:      errors.New("Bad Request"),
			response: &core.DetailedResponse{Result: map[string]interface{}{"description": s2sAuthErr.Error()}},
			expected: true,
		},
		{
			name:     "encryption key marker alone is enough",
			err:      errors.New("The service could not configure access to the encryption key specified"),
			expected: true,
		},
		{
			name:     "unrelated failure is not retried",
			err:      errors.New("no deployment found for service plan standard-gen2"),
			expected: false,
		},
		{
			name:     "unrelated failure with a response body is not retried",
			err:      errors.New("Quota exceeded"),
			response: &core.DetailedResponse{Result: map[string]interface{}{"description": "resource quota exceeded"}},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isS2SAuthorizationError(tc.err, tc.response); got != tc.expected {
				t.Fatalf("isS2SAuthorizationError() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestCreateInstanceWithS2SRetry_succeedsOnceAuthorizationPropagates(t *testing.T) {
	withInstantS2SRetrySleep(t)

	calls := 0
	create := func(*rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
		calls++
		if calls < 3 {
			return nil, &core.DetailedResponse{StatusCode: 400}, s2sAuthErr
		}
		return &rc.ResourceInstance{ID: core.StringPtr("crn:instance")}, &core.DetailedResponse{StatusCode: 201}, nil
	}

	instance, _, err := createInstanceWithS2SRetry(context.Background(), create, &rc.CreateResourceInstanceOptions{})
	if err != nil {
		t.Fatalf("expected creation to succeed after the policy propagated, got: %v", err)
	}
	if calls != 3 {
		t.Fatalf("expected 3 attempts, got %d", calls)
	}
	if instance == nil || *instance.ID != "crn:instance" {
		t.Fatalf("expected the instance from the successful attempt, got %#v", instance)
	}
}

func TestCreateInstanceWithS2SRetry_succeedsOnFirstAttempt(t *testing.T) {
	withInstantS2SRetrySleep(t)

	calls := 0
	create := func(*rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
		calls++
		return &rc.ResourceInstance{ID: core.StringPtr("crn:instance")}, &core.DetailedResponse{StatusCode: 201}, nil
	}

	if _, _, err := createInstanceWithS2SRetry(context.Background(), create, &rc.CreateResourceInstanceOptions{}); err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected a single attempt when creation succeeds, got %d", calls)
	}
}

func TestCreateInstanceWithS2SRetry_doesNotRetryUnrelatedErrors(t *testing.T) {
	withInstantS2SRetrySleep(t)

	calls := 0
	otherErr := errors.New("no deployment found for service plan standard-gen2")
	create := func(*rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
		calls++
		return nil, &core.DetailedResponse{StatusCode: 400}, otherErr
	}

	_, _, err := createInstanceWithS2SRetry(context.Background(), create, &rc.CreateResourceInstanceOptions{})
	if !errors.Is(err, otherErr) {
		t.Fatalf("expected the original error to be returned unchanged, got: %v", err)
	}
	if calls != 1 {
		t.Fatalf("expected no retry for an unrelated error, got %d attempts", calls)
	}
}

func TestCreateInstanceWithS2SRetry_givesUpWithGuidanceWhenPolicyIsMissing(t *testing.T) {
	withInstantS2SRetrySleep(t)

	calls := 0
	create := func(*rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
		calls++
		return nil, &core.DetailedResponse{StatusCode: 400}, s2sAuthErr
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(45*time.Second))
	defer cancel()

	_, _, err := createInstanceWithS2SRetry(ctx, create, &rc.CreateResourceInstanceOptions{})
	if err == nil {
		t.Fatal("expected an error when the authorization never propagates")
	}
	if !errors.Is(err, s2sAuthErr) {
		t.Fatalf("expected the API rejection to be wrapped, got: %v", err)
	}
	requireErrContains(t, err, "Authorization Delegator")
	requireErrContains(t, err, "depends_on")
	requireErrContains(t, err, fmt.Sprintf("after %d attempt(s)", calls))
	if calls < 2 {
		t.Fatalf("expected at least one retry before giving up, got %d attempts", calls)
	}
}

func TestCreateInstanceWithS2SRetry_stopsWhenContextIsCancelled(t *testing.T) {
	withInstantS2SRetrySleep(t)

	ctx, cancel := context.WithCancel(context.Background())
	calls := 0
	create := func(*rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
		calls++
		cancel()
		return nil, &core.DetailedResponse{StatusCode: 400}, s2sAuthErr
	}
	defer cancel()

	_, _, err := createInstanceWithS2SRetry(ctx, create, &rc.CreateResourceInstanceOptions{})
	if err == nil {
		t.Fatal("expected an error after cancellation")
	}
	if calls != 1 {
		t.Fatalf("expected the loop to stop after cancellation, got %d attempts", calls)
	}
}
