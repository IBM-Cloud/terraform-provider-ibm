package compute

import (
	"github.com/go-openapi/strfmt"

	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/compute"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"

	"github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
	"github.ibm.com/Bluemix/riaas-go-client/utils"
)

// InstanceClient ...
type InstanceClient struct {
	session *session.Session
}

// NewInstanceClient ...
func NewInstanceClient(sess *session.Session) *InstanceClient {
	return &InstanceClient{
		sess,
	}
}

// ListWithFilter ...
func (f *InstanceClient) ListWithFilter(zone, vpcid, subnetid, resourcegroupID, start string) ([]*models.Instance, string, error) {
	params := compute.NewGetInstancesParamsWithTimeout(f.session.Timeout)

	if zone != "" {
		params = params.WithZoneName(&zone)
	}
	if vpcid != "" {
		params = params.WithVpcID(&vpcid)
	}
	if subnetid != "" {
		params = params.WithNetworkInterfacesSubnetID(&subnetid)
	}
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	if start != "" {
		params = params.WithStart(&start)
	}
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Compute.GetInstances(params, session.Auth(f.session))

	if err != nil {
		return nil, "", errors.ToError(err)
	}

	if resp.Payload.Instances != nil {
		return resp.Payload.Instances, utils.GetNext(resp.Payload.Next), nil
	}
	return resp.Payload.Instances, utils.GetNext(resp.Payload.Next), nil
}

// List ...
func (f *InstanceClient) List(start string) ([]*models.Instance, string, error) {
	return f.ListWithFilter("", "", "", "", start)
}

// Get ...
func (f *InstanceClient) Get(id string) (*models.Instance, error) {
	params := compute.NewGetInstancesIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstancesID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetInitParms ...
func (f *InstanceClient) GetInitParms(id string) (*models.InstanceInitialization, error) {
	params := compute.NewGetInstancesInstanceIDInitializationParamsWithTimeout(f.session.Timeout).WithInstanceID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDInitialization(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Create ...
func (f *InstanceClient) Create(instancedef *models.PostInstancesParamsBody) (*models.Instance, error) {
	params := compute.NewPostInstancesParamsWithTimeout(f.session.Timeout).WithBody(instancedef)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PostInstances(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateEasy ...
func (f *InstanceClient) CreateEasy(name string) (*models.Instance, error) {

	var body = models.PostInstancesParamsBody{
		Name: name,
	}
	return f.Create(&body)
}

// Update ...
func (f *InstanceClient) Update(id, name, profileName string) (*models.Instance, error) {
	var body = models.PatchInstancesIDParamsBody{}
	if name != "" {
		body.Name = name
	}
	if profileName != "" {
		var profile = models.PatchInstancesIDParamsBodyProfile{
			Name: profileName,
		}
		body.Profile = &profile
	}

	params := compute.NewPatchInstancesIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PatchInstancesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *InstanceClient) Delete(id string) error {
	params := compute.NewDeleteInstancesIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Compute.DeleteInstancesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// CreateAction ...
func (f *InstanceClient) CreateAction(instanceid, actiontype string) (*models.InstanceAction, error) {
	body := models.PostInstancesInstanceIDActionsParamsBody{
		Type: actiontype,
	}
	params := compute.NewPostInstancesInstanceIDActionsParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PostInstancesInstanceIDActions(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// ListProfiles ...
func (f *InstanceClient) ListProfiles(start string) ([]*models.InstanceProfile, string, error) {
	params := compute.NewGetInstanceProfilesParamsWithTimeout(f.session.Timeout)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstanceProfiles(params, session.Auth(f.session))
	if err != nil {
		return nil, "", errors.ToError(err)
	}
	return resp.Payload.Profiles, utils.GetNext(resp.Payload.Next), nil
}

// GetProfile ...
func (f *InstanceClient) GetProfile(profileName string) (*models.InstanceProfile, error) {
	params := compute.NewGetInstanceProfilesNameParamsWithTimeout(f.session.Timeout).WithName(profileName)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstanceProfilesName(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// ListInterfacesWithFilter ...
func (f *InstanceClient) ListInterfacesWithFilter(instanceid, resourcegroupID string) ([]*models.InstanceNetworkInterface, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfaces(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.NetworkInterfaces, nil
}

// ListInterfaces ...
func (f *InstanceClient) ListInterfaces(instanceid string) ([]*models.InstanceNetworkInterface, error) {
	return f.ListInterfacesWithFilter(instanceid, "")
}

// GetInterface ...
func (f *InstanceClient) GetInterface(instanceid, interfaceid string) (*models.InstanceNetworkInterface, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesIDParamsWithTimeout(f.session.Timeout).
		WithInstanceID(instanceid).
		WithID(interfaceid)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AddInterface ...
func (f *InstanceClient) AddInterface(instanceid, name, subnetID string, portSpeed int, v4address, v6address string,
	secondaryAddresses, securityGroupIDs []string) (*models.InstanceNetworkInterface, error) {

	body := models.PostInstancesInstanceIDNetworkInterfacesParamsBody{}
	body.Name = name
	body.PortSpeed = int64(portSpeed)
	if v6address != "" {
		body.PrimaryIPV6Address = v6address
	}
	if v4address != "" {
		body.PrimaryIPV4Address = v4address
	}

	if len(secondaryAddresses) != 0 {
		body.SecondaryAddresses = secondaryAddresses
	}

	if len(securityGroupIDs) != 0 {
		sgs := make([]*models.PostInstancesInstanceIDNetworkInterfacesParamsBodySecurityGroupsItems, len(securityGroupIDs))
		for i, sgid := range securityGroupIDs {
			sgref := models.PostInstancesInstanceIDNetworkInterfacesParamsBodySecurityGroupsItems{
				ID: strfmt.UUID(sgid),
			}
			sgs[i] = &sgref
		}
		body.SecurityGroups = sgs
	}

	subnetref := models.PostInstancesInstanceIDNetworkInterfacesParamsBodySubnet{
		ID: strfmt.UUID(subnetID),
	}
	body.Subnet = &subnetref

	params := compute.NewPostInstancesInstanceIDNetworkInterfacesParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PostInstancesInstanceIDNetworkInterfaces(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteInterface ...
func (f *InstanceClient) DeleteInterface(instanceid, interfaceid string) error {
	params := compute.NewDeleteInstancesInstanceIDNetworkInterfacesIDParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithID(interfaceid)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// UpdateInterface ...
func (f *InstanceClient) UpdateInterface(instanceid, interfaceid, name string, portSpeed int) (*models.InstanceNetworkInterface, error) {

	body := models.PatchInstancesInstanceIDNetworkInterfacesIDParamsBody{}
	if name != "" {
		body.Name = name
	}
	if portSpeed != 0 {
		body.PortSpeed = int64(portSpeed)
	}
	params := compute.NewPatchInstancesInstanceIDNetworkInterfacesIDParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithID(interfaceid).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PatchInstancesInstanceIDNetworkInterfacesID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// ListInterfaceFloatingIPs ...
func (f *InstanceClient) ListInterfaceFloatingIPs(instanceid, interfaceid string) ([]*models.FloatingIP, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsParamsWithTimeout(f.session.Timeout).
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIps(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.FloatingIps, nil
}

// GetInterfaceFloatingIP ...
func (f *InstanceClient) GetInterfaceFloatingIP(instanceid, interfaceid, address string) (*models.FloatingIP, error) {
	params := compute.NewGetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithTimeout(f.session.Timeout).
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AddInterfaceFloatingIP ...
func (f *InstanceClient) AddInterfaceFloatingIP(instanceid, interfaceid, address string) (*models.FloatingIP, error) {
	params := compute.NewPutInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithTimeout(f.session.Timeout).
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PutInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// RemoveInterfaceFloatingIP ...
func (f *InstanceClient) RemoveInterfaceFloatingIP(instanceid, interfaceid, address string) error {
	params := compute.NewDeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddressParamsWithTimeout(f.session.Timeout).
		WithInstanceID(instanceid).WithNetworkInterfaceID(interfaceid).WithAddress(address)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDNetworkInterfacesNetworkInterfaceIDFloatingIpsAddress(
		params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// ListVolAttachmentsWithFilter ...
func (f *InstanceClient) ListVolAttachmentsWithFilter(instanceid, resourcegroupID string) ([]*models.InstanceVolumeAttachment, error) {
	params := compute.NewGetInstancesInstanceIDVolumeAttachmentsParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid)
	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation

	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDVolumeAttachments(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload.VolumeAttachments, nil
}

// ListVolAttachments ...
func (f *InstanceClient) ListVolAttachments(instanceid string) ([]*models.InstanceVolumeAttachment, error) {
	return f.ListVolAttachmentsWithFilter(instanceid, "")
}

// GetVolAttachment ...
func (f *InstanceClient) GetVolAttachment(instanceid, volAttachID string) (*models.InstanceVolumeAttachment, error) {
	params := compute.NewGetInstancesInstanceIDVolumeAttachmentsIDParamsWithTimeout(f.session.Timeout).
		WithInstanceID(instanceid).
		WithID(volAttachID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.GetInstancesInstanceIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// AttachVolume ...
func (f *InstanceClient) AttachVolume(instanceid, volumeID, name string, resourcegroupID string) (*models.InstanceVolumeAttachment, error) {

	body := models.PostInstancesInstanceIDVolumeAttachmentsParamsBody{}
	if name != "" {
		body.Name = name
	}
	body.Volume = &models.PostInstancesInstanceIDVolumeAttachmentsParamsBodyVolume{
		ID: strfmt.UUID(volumeID),
	}

	params := compute.NewPostInstancesInstanceIDVolumeAttachmentsParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PostInstancesInstanceIDVolumeAttachments(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// DeleteVolAttachment ...
func (f *InstanceClient) DeleteVolAttachment(instanceid, volAttachID string) error {
	params := compute.NewDeleteInstancesInstanceIDVolumeAttachmentsIDParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithID(volAttachID)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.Compute.DeleteInstancesInstanceIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// UpdateVolAttachment ...
func (f *InstanceClient) UpdateVolAttachment(instanceid, volAttachID, name string) (*models.InstanceVolumeAttachment, error) {

	body := models.PatchInstancesInstanceIDVolumeAttachmentsIDParamsBody{}
	if name != "" {
		body.Name = name
	}

	params := compute.NewPatchInstancesInstanceIDVolumeAttachmentsIDParamsWithTimeout(f.session.Timeout).WithInstanceID(instanceid).WithID(volAttachID).WithBody(&body)
	params.Version = "2019-07-02"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.Compute.PatchInstancesInstanceIDVolumeAttachmentsID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
