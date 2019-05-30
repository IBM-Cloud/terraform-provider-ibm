package vpn

import (
	"github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/riaas/rias-api/riaas/client/v_p_naa_s"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

// vpn ...
type VpnClient struct {
	session *session.Session
}

// NewVpnClient ...
func NewVpnClient(sess *session.Session) *VpnClient {
	return &VpnClient{
		sess,
	}
}

// GetIkePolicies ...
func (f *VpnClient) ListIkePolicies(limit int32, start, tag string) (*models.IKEPolicyCollection, error) {
	params := v_p_naa_s.NewGetIkePoliciesParams()
	if start != "" {
		params = params.WithStart(&start)
	}
	if limit != 0 {
		params = params.WithLimit(&limit)
	}
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetIkePolicies(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateIkePolicy ...
func (f *VpnClient) CreateIkePolicy(authenticationAlgorithm, encryptionAlgorithm, name, resourceGrpId string, tags []string, dhGroup, ikeVersion, keyLifetime int) (*models.IKEPolicy, error) {
	var body = models.IKEPolicyTemplate{}
	body.Name = name
	body.AuthenticationAlgorithm = authenticationAlgorithm
	body.DhGroup = int64(dhGroup)
	body.EncryptionAlgorithm = encryptionAlgorithm
	body.IkeVersion = int64(ikeVersion)
	if resourceGrpId != "" {
		rgref := models.IKEPolicyTemplateResourceGroup{
			ID: strfmt.UUID(resourceGrpId),
		}
		body.ResourceGroup = &rgref
	}
	if keyLifetime != 0 {
		body.KeyLifetime = int64(keyLifetime)
	}
	if len(tags) != 0 {
		body.Tags = tags
	}
	params := v_p_naa_s.NewPostIkePoliciesParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PostIkePolicies(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteIkePolicy ...
func (f *VpnClient) DeleteIkePolicy(ikePolicyId string) error {
	params := v_p_naa_s.NewDeleteIkePoliciesIDParams().WithID(ikePolicyId)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.VPNaaS.DeleteIkePoliciesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// GetIkePolicy ...
func (f *VpnClient) GetIkePolicy(id string) (*models.IKEPolicy, error) {
	params := v_p_naa_s.NewGetIkePoliciesIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetIkePoliciesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateIkePolicy ...
func (f *VpnClient) UpdateIkePolicy(id, authenticationAlgorithm, encryptionAlgorithm, name string, dhGroup, ikeVersion, keyLifetime int) (*models.IKEPolicy, error) {
	var body = models.IKEPolicyPatch{}
	if name != "" {
		body.Name = name
	}
	if authenticationAlgorithm != "" {
		body.AuthenticationAlgorithm = authenticationAlgorithm
	}
	if dhGroup != 0 {
		body.DhGroup = int64(dhGroup)
	}
	if encryptionAlgorithm != "" {
		body.EncryptionAlgorithm = encryptionAlgorithm
	}
	if ikeVersion != 0 {
		body.IkeVersion = int64(ikeVersion)
	}
	if keyLifetime != 0 {
		body.KeyLifetime = int64(keyLifetime)
	}
	params := v_p_naa_s.NewPatchIkePoliciesIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PatchIkePoliciesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetIkePoliciesConnections ...
func (f *VpnClient) GetIkePoliciesConnections(id string) (*models.VPNGatewayConnectionCollection, error) {
	params := v_p_naa_s.NewGetIkePoliciesIDConnectionsParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetIkePoliciesIDConnections(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// ListIpsecPolicies ...
func (f *VpnClient) ListIpsecPolicies(limit int32, start string) (*models.IpsecPolicyCollection, error) {
	params := v_p_naa_s.NewGetIpsecPoliciesParams()
	if start != "" {
		params = params.WithStart(&start)
	}
	if limit != 0 {
		params = params.WithLimit(&limit)
	}
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetIpsecPolicies(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateIpsecPolicy ...
func (f *VpnClient) CreateIpsecPolicy(authenticationAlgorithm, encryptionAlgorithm, name, pfs, resourceGrpId string, keyLifetime int) (*models.IpsecPolicy, error) {
	var body = models.IpsecPolicyTemplate{}
	body.Name = name
	body.AuthenticationAlgorithm = authenticationAlgorithm
	body.EncryptionAlgorithm = encryptionAlgorithm
	body.Pfs = pfs
	if resourceGrpId != "" {
		rgref := models.IpsecPolicyTemplateResourceGroup{
			ID: strfmt.UUID(resourceGrpId),
		}
		body.ResourceGroup = &rgref
	}
	if keyLifetime != 0 {
		body.KeyLifetime = int64(keyLifetime)
	}
	params := v_p_naa_s.NewPostIpsecPoliciesParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PostIpsecPolicies(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteIpsecPolicy ...
func (f *VpnClient) DeleteIpsecPolicy(id string) error {
	params := v_p_naa_s.NewDeleteIpsecPoliciesIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.VPNaaS.DeleteIpsecPoliciesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// GetIpsecPolicy ...
func (f *VpnClient) GetIpsecPolicy(id string) (*models.IpsecPolicy, error) {
	params := v_p_naa_s.NewGetIpsecPoliciesIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetIpsecPoliciesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateIpsecPolicy ...
func (f *VpnClient) UpdateIpsecPolicy(id, authenticationAlgorithm, encryptionAlgorithm, name, pfs string, keyLifetime int) (*models.IpsecPolicy, error) {
	var body = models.IpsecPolicyPatch{}
	if name != "" {
		body.Name = name
	}
	if authenticationAlgorithm != "" {
		body.AuthenticationAlgorithm = authenticationAlgorithm
	}
	if pfs != "" {
		body.Pfs = pfs
	}
	if encryptionAlgorithm != "" {
		body.EncryptionAlgorithm = encryptionAlgorithm
	}
	if keyLifetime != 0 {
		body.KeyLifetime = int64(keyLifetime)
	}
	params := v_p_naa_s.NewPatchIpsecPoliciesIDParams().WithID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PatchIpsecPoliciesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetIpsecPoliciesConnections ...
func (f *VpnClient) GetIpsecPoliciesConnections(id string) (*models.VPNGatewayConnectionCollection, error) {
	params := v_p_naa_s.NewGetIpsecPoliciesIDConnectionsParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetIpsecPoliciesIDConnections(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// List ...
func (f *VpnClient) List(limit int32, resourceGrpId, start string) (*models.VPNGatewayCollection, error) {
	params := v_p_naa_s.NewGetVpnGatewaysParams()
	if start != "" {
		params = params.WithStart(&start)
	}
	if resourceGrpId != "" {
		params = params.WithResourceGroupID(&resourceGrpId)
	}
	if limit != 0 {
		params = params.WithLimit(&limit)
	}
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetVpnGateways(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *VpnClient) Create(name, subnetId, resourceGrpId string) (*models.VPNGateway, error) {
	var body = models.VPNGatewayTemplate{}
	body.Name = name
	if resourceGrpId != "" {
		rgref := models.VPNGatewayTemplateResourceGroup{
			ID: strfmt.UUID(resourceGrpId),
		}
		body.ResourceGroup = &rgref
	}

	if subnetId != "" {

		subnetref := models.VPNGatewayTemplateSubnet{
			ID: strfmt.UUID(subnetId),
		}
		body.Subnet = &subnetref
	}

	params := v_p_naa_s.NewPostVpnGatewaysParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PostVpnGateways(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *VpnClient) Delete(id string) error {
	params := v_p_naa_s.NewDeleteVpnGatewaysIDParams().WithID(id)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.VPNaaS.DeleteVpnGatewaysID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Get ...
func (f *VpnClient) Get(id string) (*models.VPNGateway, error) {
	params := v_p_naa_s.NewGetVpnGatewaysIDParams().WithID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetVpnGatewaysID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Update ...
func (f *VpnClient) Update(id, name string) (*models.VPNGateway, error) {
	var body = models.VPNGatewayPatch{}
	if name != "" {
		body.Name = name
	}

	params := v_p_naa_s.NewPatchVpnGatewaysIDParams().WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PatchVpnGatewaysID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetConnections ...
func (f *VpnClient) GetConnections(id string) (*models.VPNGatewayConnectionCollection, error) {
	params := v_p_naa_s.NewGetVpnGatewaysVpnGatewayIDConnectionsParams().WithVpnGatewayID(id)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetVpnGatewaysVpnGatewayIDConnections(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

//TODO
// Create ...
func (f *VpnClient) CreateConnections(id, name, peerAddress, psk string, peerCidrs, localCidrs []string, adminStateUp bool,
	deadPeerDetection *models.VPNGatewayConnectionDPD, ikePolicy *models.IKEPolicyIdentity, ipsecPolicy *models.IpsecPolicyIdentity) (*models.VPNGatewayConnection, error) {
	var body = models.VPNGatewayConnectionTemplate{}
	body.Name = name
	body.PeerAddress = peerAddress
	body.Psk = psk
	body.PeerCidrs = peerCidrs
	body.LocalCidrs = localCidrs
	if !adminStateUp {
		body.AdminStateUp = &adminStateUp
	}

	if deadPeerDetection != nil {
		body.DeadPeerDetection = deadPeerDetection
	}
	if ikePolicy != nil {
		body.IkePolicy = ikePolicy
	}
	if ipsecPolicy != nil {
		body.IpsecPolicy = ipsecPolicy
	}
	params := v_p_naa_s.NewPostVpnGatewaysVpnGatewayIDConnectionsParams().WithVpnGatewayID(id).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PostVpnGatewaysVpnGatewayIDConnections(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteConnection ...
func (f *VpnClient) DeleteConnection(vpnGatewayId, conenctionId string) error {
	params := v_p_naa_s.NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDParams().WithID(conenctionId).WithVpnGatewayID(vpnGatewayId)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.VPNaaS.DeleteVpnGatewaysVpnGatewayIDConnectionsID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// GetConnection ...
func (f *VpnClient) GetConnection(vpnGatewayId, conenctionId string) (*models.VPNGatewayConnection, error) {
	params := v_p_naa_s.NewGetVpnGatewaysVpnGatewayIDConnectionsIDParams().WithID(conenctionId).WithVpnGatewayID(vpnGatewayId)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetVpnGatewaysVpnGatewayIDConnectionsID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateConnection ...
func (f *VpnClient) UpdateConnection(id, vpnGatewayId, name, peerAddress, psk string, adminStateUp bool,
	deadPeerDetection *models.VPNGatewayConnectionDPD, ikePolicy *models.IKEPolicyIdentity, ipsecPolicy *models.IpsecPolicyIdentity) (*models.VPNGatewayConnection, error) {
	var body = models.VPNGatewayConnectionPatch{}
	if name != "" {
		body.Name = name
	}
	if peerAddress != "" {
		body.PeerAddress = peerAddress
	}
	if psk != "" {
		body.Psk = psk
	}
	body.AdminStateUp = &adminStateUp

	if deadPeerDetection != nil {
		body.DeadPeerDetection = deadPeerDetection
	}
	if ikePolicy != nil {
		body.IkePolicy = ikePolicy
	}
	if ipsecPolicy != nil {
		body.IpsecPolicy = ipsecPolicy
	}
	params := v_p_naa_s.NewPatchVpnGatewaysVpnGatewayIDConnectionsIDParams().WithID(id).WithVpnGatewayID(vpnGatewayId).WithBody(&body)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.PatchVpnGatewaysVpnGatewayIDConnectionsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// ListIpsecPolicies ...
func (f *VpnClient) ListLocalDirs(id, vpnGatewayId string) (*models.VPNGatewayConnectionLocalCIDRs, error) {
	params := v_p_naa_s.NewGetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsParams().WithID(id).WithVpnGatewayID(vpnGatewayId)
	params.Version = "2019-03-26"
	resp, err := f.session.Riaas.VPNaaS.GetVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrs(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteLocalCidr ...
func (f *VpnClient) DeleteLocalCidr(vpnGatewayId, conenctionId, prefixAddress, prefixLength string) error {
	params := v_p_naa_s.NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLengthParams().WithID(conenctionId).WithVpnGatewayID(vpnGatewayId).WithPrefixAddress(prefixAddress).WithPrefixLength(prefixLength)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.VPNaaS.DeleteVpnGatewaysVpnGatewayIDConnectionsIDLocalCidrsPrefixAddressPrefixLength(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// DeletePeerCidr ...
func (f *VpnClient) DeletePeerCidr(vpnGatewayId, conenctionId, prefixAddress, prefixLength string) error {
	params := v_p_naa_s.NewDeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLengthParams().WithID(conenctionId).WithVpnGatewayID(vpnGatewayId).WithPrefixAddress(prefixAddress).WithPrefixLength(prefixLength)
	params.Version = "2019-03-26"
	_, err := f.session.Riaas.VPNaaS.DeleteVpnGatewaysVpnGatewayIDConnectionsIDPeerCidrsPrefixAddressPrefixLength(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}
