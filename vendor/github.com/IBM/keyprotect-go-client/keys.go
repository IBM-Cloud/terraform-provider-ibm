// Copyright 2019 IBM Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kp

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	ReturnMinimal        PreferReturn = 0
	ReturnRepresentation PreferReturn = 1

	keyType    = "application/vnd.ibm.kms.key+json"
	policyType = "application/vnd.ibm.kms.policy+json"
)

var (
	preferHeaders = []string{"return=minimal", "return=representation"}
)

// PreferReturn designates the value for the "Prefer" header.
type PreferReturn int

// Key represents a key as returned by the KP API.
type Key struct {
	ID                  string     `json:"id,omitempty"`
	Name                string     `json:"name,omitempty"`
	Description         string     `json:"description,omitempty"`
	Type                string     `json:"type,omitempty"`
	Tags                []string   `json:"Tags,omitempty"`
	AlgorithmType       string     `json:"algorithmType,omitempty"`
	CreatedBy           string     `json:"createdBy,omitempty"`
	CreationDate        *time.Time `json:"creationDate,omitempty"`
	LastUpdateDate      *time.Time `json:"lastUpdateDate,omitempty"`
	LastRotateDate      *time.Time `json:"lastRotateDate,omitempty"`
	Extractable         bool       `json:"extractable"`
	Expiration          *time.Time `json:"expirationDate,omitempty"`
	Payload             string     `json:"payload,omitempty"`
	State               int        `json:"state,omitempty"`
	EncryptionAlgorithm string     `json:"encryptionAlgorithm,omitempty"`
	CRN                 string     `json:"crn,omitempty"`
	EncryptedNonce      string     `json:"encryptedNonce,omitempty"`
	IV                  string     `json:"iv,omitempty"`
}

// KeysMetadata represents the metadata of a collection of keys.
type KeysMetadata struct {
	CollectionType string `json:"collectionType"`
	NumberOfKeys   int    `json:"collectionTotal"`
}

// Keys represents a collection of Keys.
type Keys struct {
	Metadata KeysMetadata `json:"metadata"`
	Keys     []Key        `json:"resources"`
}

// KeysActionRequest represents request parameters for a key action
// API call.
type KeysActionRequest struct {
	PlainText  string   `json:"plaintext,omitempty"`
	AAD        []string `json:"aad,omitempty"`
	CipherText string   `json:"ciphertext,omitempty"`
	Payload    string   `json:"payload,omitempty"`
}

// Create creates a new KP key.
func (c *Client) CreateKey(ctx context.Context, name string, expiration *time.Time, extractable bool) (*Key, error) {
	return c.CreateImportedKey(ctx, name, expiration, "", "", "", extractable)
}

// CreateImportedKey creates a new KP key from the given key material.
func (c *Client) CreateImportedKey(ctx context.Context, name string, expiration *time.Time, payload, encryptedNonce, iv string, extractable bool) (*Key, error) {
	key := Key{
		Name:        name,
		Type:        keyType,
		Extractable: extractable,
		Payload:     payload,
	}

	if payload != "" && encryptedNonce != "" && iv != "" {
		key.EncryptedNonce = encryptedNonce
		key.IV = iv
		key.EncryptionAlgorithm = importTokenEncAlgo
	}

	if expiration != nil {
		key.Expiration = expiration
	}

	keysRequest := Keys{
		Metadata: KeysMetadata{
			CollectionType: keyType,
			NumberOfKeys:   1,
		},
		Keys: []Key{key},
	}

	req, err := c.newRequest("POST", "keys", &keysRequest)
	if err != nil {
		return nil, err
	}

	keysResponse := Keys{}
	if _, err := c.do(ctx, req, &keysResponse); err != nil {
		return nil, err
	}

	return &keysResponse.Keys[0], nil
}

// CreateRootKey creates a new, non-extractable key resource without
// key material.
func (c *Client) CreateRootKey(ctx context.Context, name string, expiration *time.Time) (*Key, error) {
	return c.CreateKey(ctx, name, expiration, false)
}

// CreateStandardKey creates a new, extractable key resource without
// key material.
func (c *Client) CreateStandardKey(ctx context.Context, name string, expiration *time.Time) (*Key, error) {
	return c.CreateKey(ctx, name, expiration, true)
}

// CreateImportedRootKey creates a new, non-extractable key resource
// with the given key material.
func (c *Client) CreateImportedRootKey(ctx context.Context, name string, expiration *time.Time, payload, encryptedNonce, iv string) (*Key, error) {
	return c.CreateImportedKey(ctx, name, expiration, payload, encryptedNonce, iv, false)
}

// CreateStandardKey creates a new, extractable key resource with the
// given key material.
func (c *Client) CreateImportedStandardKey(ctx context.Context, name string, expiration *time.Time, payload string) (*Key, error) {
	return c.CreateImportedKey(ctx, name, expiration, payload, "", "", true)
}

// GetKeys retrieves a collection of keys that can be paged through.
func (c *Client) GetKeys(ctx context.Context, limit int, offset int) (*Keys, error) {
	if limit == 0 {
		limit = 2000
	}

	req, err := c.newRequest("GET", "keys", nil)
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Set("limit", strconv.Itoa(limit))
	v.Set("offset", strconv.Itoa(offset))
	req.URL.RawQuery = v.Encode()

	keys := Keys{}
	_, err = c.do(ctx, req, &keys)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}

// GetKey retrieves a key by ID.
func (c *Client) GetKey(ctx context.Context, id string) (*Key, error) {
	keys := Keys{}

	req, err := c.newRequest("GET", fmt.Sprintf("keys/%s", id), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.do(ctx, req, &keys)
	if err != nil {
		return nil, err
	}

	return &keys.Keys[0], nil
}

// Delete deletes a key resource by specifying the ID of the key.
func (c *Client) DeleteKey(ctx context.Context, id string, prefer PreferReturn) (*Key, error) {

	req, err := c.newRequest("DELETE", fmt.Sprintf("keys/%s", id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Prefer", preferHeaders[prefer])

	keys := Keys{}
	_, err = c.do(ctx, req, &keys)
	if err != nil {
		return nil, err
	}

	if len(keys.Keys) > 0 {
		return &keys.Keys[0], nil
	}

	return nil, nil
}

// Wrap calls the wrap action with the given plain text.
func (c *Client) Wrap(ctx context.Context, id string, plainText []byte, additionalAuthData *[]string) ([]byte, error) {
	pt, _, err := c.wrap(ctx, id, plainText, additionalAuthData)
	return pt, err
}

// WrapCreateDEK calls the wrap action without plain text.
func (c *Client) WrapCreateDEK(ctx context.Context, id string, additionalAuthData *[]string) ([]byte, []byte, error) {
	return c.wrap(ctx, id, nil, additionalAuthData)
}

func (c *Client) wrap(ctx context.Context, id string, plainText []byte, additionalAuthData *[]string) ([]byte, []byte, error) {
	keysActionReq := &KeysActionRequest{}

	if plainText != nil {
		_, err := base64.StdEncoding.DecodeString(string(plainText))
		if err != nil {
			return nil, nil, err
		}
		keysActionReq.PlainText = string(plainText)
	}

	if additionalAuthData != nil {
		keysActionReq.AAD = *additionalAuthData
	}

	keysAction, err := c.doKeysAction(ctx, id, "wrap", keysActionReq)
	if err != nil {
		return nil, nil, err
	}

	pt := []byte(keysAction.PlainText)
	ct := []byte(keysAction.CipherText)

	return pt, ct, nil
}

// Unwrap is deprecated since it returns only plaintext and doesn't know how to handle rotation.
func (c *Client) Unwrap(ctx context.Context, id string, cipherText []byte, additionalAuthData *[]string) ([]byte, error) {
	plainText, _, err := c.UnwrapV2(ctx, id, cipherText, additionalAuthData)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// Unwrap with rotation support.
func (c *Client) UnwrapV2(ctx context.Context, id string, cipherText []byte, additionalAuthData *[]string) ([]byte, []byte, error) {

	keysAction := &KeysActionRequest{
		CipherText: string(cipherText),
	}

	if additionalAuthData != nil {
		keysAction.AAD = *additionalAuthData
	}

	respAction, err := c.doKeysAction(ctx, id, "unwrap", keysAction)
	if err != nil {
		return nil, nil, err
	}

	plainText := []byte(respAction.PlainText)
	rewrapped := []byte(respAction.CipherText)

	return plainText, rewrapped, nil
}

// Rotate rotates a CRK.
func (c *Client) Rotate(ctx context.Context, id, payload string) error {

	actionReq := &KeysActionRequest{
		Payload: payload,
	}

	_, err := c.doKeysAction(ctx, id, "rotate", actionReq)
	if err != nil {
		return err
	}

	return nil
}

// Policy represents a policy as returned by the KP API.
type Policy struct {
	Type      string     `json:"type,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	CreatedAt *time.Time `json:"creationDate,omitempty"`
	CRN       string     `json:"crn,omitempty"`
	UpdatedAt *time.Time `json:"lastUpdateDate,omitempty"`
	UpdatedBy string     `json:"updatedBy,omitempty"`
	Rotation  struct {
		Interval int `json:"interval_month,omitempty"`
	} `json:"rotation,omitempty"`
}

// PoliciesMetadata represents the metadata of a collection of keys.
type PoliciesMetadata struct {
	CollectionType   string `json:"collectionType"`
	NumberOfPolicies int    `json:"collectionTotal"`
}

// Policies represents a collection of Policies.
type Policies struct {
	Metadata PoliciesMetadata `json:"metadata"`
	Policies []Policy         `json:"resources"`
}

// GetPolicy retrieves a policy by Key ID.
func (c *Client) GetPolicy(ctx context.Context, id string) (*Policy, error) {
	policyresponse := Policies{}

	req, err := c.newRequest("GET", fmt.Sprintf("keys/%s/policy", id), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.do(ctx, req, &policyresponse)
	if err != nil {
		return nil, err
	}

	return &policyresponse.Policies[0], nil
}

// SetPolicy updates a policy resource by specifying the ID of the key and the rotation interval needed.
func (c *Client) SetPolicy(ctx context.Context, id string, prefer PreferReturn, rotationInterval int) (*Policy, error) {

	policy := Policy{
		Type: policyType,
	}
	policy.Rotation.Interval = rotationInterval

	policyRequest := Policies{
		Metadata: PoliciesMetadata{
			CollectionType:   keyType,
			NumberOfPolicies: 1,
		},
		Policies: []Policy{policy},
	}

	policyresponse := Policies{}

	req, err := c.newRequest("PUT", fmt.Sprintf("keys/%s/policy", id), &policyRequest)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Prefer", preferHeaders[prefer])

	_, err = c.do(ctx, req, &policyresponse)
	if err != nil {
		return nil, err
	}

	return &policyresponse.Policies[0], nil
}

// doKeysAction calls the KP Client to perform an action on a key.
func (c *Client) doKeysAction(ctx context.Context, id string, action string, keysActionReq *KeysActionRequest) (*KeysActionRequest, error) {
	keyActionRsp := KeysActionRequest{}

	v := url.Values{}
	v.Set("action", action)

	req, err := c.newRequest("POST", fmt.Sprintf("keys/%s", id), keysActionReq)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = v.Encode()

	_, err = c.do(ctx, req, &keyActionRsp)
	if err != nil {
		return nil, err
	}
	return &keyActionRsp, nil
}
