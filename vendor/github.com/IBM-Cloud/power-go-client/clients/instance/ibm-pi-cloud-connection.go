package instance

import (
	"fmt"
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
	"time"
)

type IBMPICloudConnectionClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// IBMPICloudInstanceClient ...

func NewIBMPICloudConnectionClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPICloudConnectionClient {
	return &IBMPICloudConnectionClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

// Create a Cloud Connection

func (f *IBMPICloudConnectionClient) Create(pclouddef *p_cloud_cloud_connections.PcloudCloudconnectionsPostParams, powerinstanceid string) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPostParamsWithTimeout(postTimeOut).WithCloudInstanceID(powerinstanceid).WithBody(pclouddef.Body)
	postok, postcreated, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		log.Printf("failed to process the request..")
		return nil, nil
	}

	if postok != nil {
		log.Printf("Checking if the instance name is right ")
		log.Printf("Printing the CloudConnectionid %s", *postok.Payload.CloudConnectionID)
		return postok.Payload, nil
	}
	if postcreated != nil {
		log.Printf("Printing the CloudConnectionid %s", *postcreated.Payload.CloudConnectionID)
		return postcreated.Payload, nil
	}

	return nil, fmt.Errorf("No response Returned ")
}

/*
 gets a cloud connection s state information
*/
func (f *IBMPICloudConnectionClient) Get(pclouddef *p_cloud_cloud_connections.PcloudCloudconnectionsGetParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetParams().WithCloudInstanceID(pclouddef.CloudInstanceID).WithCloudConnectionID(pclouddef.CloudConnectionID)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGet(params, ibmpisession.NewAuth(f.session, pclouddef.CloudInstanceID))

	if err != nil {
		log.Printf("Failed to perform get information about the cloud connection object... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

/*
 gets a cloud connection s state information
*/
func (f *IBMPICloudConnectionClient) GetAll(powerinstanceid string, timeout time.Duration) (*models.CloudConnections, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsGetallParamsWithTimeout(timeout).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil {
		log.Printf("Failed to perform get information about the cloud connection object... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//  Update a cloud Connection

func (f *IBMPICloudConnectionClient) Update(updateparams *p_cloud_cloud_connections.PcloudCloudconnectionsPutParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsPutParams().WithCloudInstanceID(updateparams.CloudInstanceID).WithCloudConnectionID(updateparams.CloudConnectionID).WithBody(updateparams.Body)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsPut(params, ibmpisession.NewAuth(f.session, updateparams.CloudInstanceID))

	if err != nil {
		log.Printf("Failed to perform the update operations %v", err)
		return nil, errors.ToError(err)

	}
	return resp.Payload, nil
}

// Delete a Cloud Connection

func (f *IBMPICloudConnectionClient) Delete(pclouddef *p_cloud_cloud_connections.PcloudCloudconnectionsDeleteParams) (models.Object, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsDeleteParams().WithCloudInstanceID(pclouddef.CloudInstanceID).WithCloudConnectionID(pclouddef.CloudConnectionID)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsDelete(params, ibmpisession.NewAuth(f.session, pclouddef.CloudInstanceID))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Add a network to a cloud connection
func (f *IBMPICloudConnectionClient) AddNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksPutParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksPutParams().WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to add the network to the cloudconnection... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Deletes a network from a cloud connection

func (f *IBMPICloudConnectionClient) DeleteNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksDeleteParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksDeleteParams().WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksDelete(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the delete operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Update a network from a cloud connection

func (f *IBMPICloudConnectionClient) UpdateNetwork(pcloudnetworkdef *p_cloud_cloud_connections.PcloudCloudconnectionsNetworksPutParams) (*models.CloudConnection, error) {

	params := p_cloud_cloud_connections.NewPcloudCloudconnectionsNetworksPutParams().WithCloudInstanceID(pcloudnetworkdef.CloudInstanceID).WithCloudConnectionID(pcloudnetworkdef.CloudConnectionID).WithNetworkID(pcloudnetworkdef.NetworkID)
	resp, err := f.session.Power.PCloudCloudConnections.PcloudCloudconnectionsNetworksPut(params, ibmpisession.NewAuth(f.session, pcloudnetworkdef.CloudInstanceID))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the delete operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}
