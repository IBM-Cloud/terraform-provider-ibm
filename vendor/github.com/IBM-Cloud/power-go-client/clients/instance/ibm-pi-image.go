package instance

import (
	"github.com/IBM-Cloud/power-go-client/errors"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"log"
)

type IBMPIImageClient struct {
	session         *ibmpisession.IBMPISession
	powerinstanceid string
}

// NewPowerImageClient ...
func NewIBMPIImageClient(sess *ibmpisession.IBMPISession, powerinstanceid string) *IBMPIImageClient {
	return &IBMPIImageClient{
		session:         sess,
		powerinstanceid: powerinstanceid,
	}
}

func (f *IBMPIImageClient) Get(id, powerinstanceid string) (*models.Image, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParamsWithTimeout(postTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params, ibmpisession.NewAuth(f.session, powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

func (f *IBMPIImageClient) GetAll(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudCloudinstancesImagesGetallParamsWithTimeout(getTimeOut).WithCloudInstanceID(powerinstanceid)
	resp, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil

}

//Post the stock image

func (f *IBMPIImageClient) Create(name, imageid string, powerinstanceid string) (*models.Image, error) {

	var source = "root-project"
	//createDate := strfmt.DateTime(time.Now())
	var body = models.CreateImage{
		ImageName: name,
		ImageID:   imageid,
		Source:    &source,
	}
	params := p_cloud_images.NewPcloudCloudinstancesImagesPostParamsWithTimeout(postTimeOut).WithCloudInstanceID(powerinstanceid).WithBody(&body)
	resp, err, _ := f.session.Power.PCloudImages.PcloudCloudinstancesImagesPost(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err.Payload.State == "queued" {
		log.Printf("Post is successful %s", *err.Payload.ImageID)

	}

	if resp != nil {
		log.Printf("Failed to initiate the copy job ")
	}

	return err.Payload, nil

}

// Delete ...
func (f *IBMPIImageClient) Delete(id string, powerinstanceid string) error {
	params := p_cloud_images.NewPcloudCloudinstancesImagesDeleteParamsWithTimeout(deleteTimeOut).WithCloudInstanceID(powerinstanceid).WithImageID(id)
	_, err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesDelete(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

//Get stock images

func (f *IBMPIImageClient) GetStockImages(powerinstanceid string) (*models.Images, error) {

	params := p_cloud_images.NewPcloudImagesGetallParams()
	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, f.powerinstanceid))

	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Get SAP Images

func (f *IBMPIImageClient) GetSAPImages(powerinstanceid string, sapimage bool) (*models.Images, error) {
	params := p_cloud_images.NewPcloudImagesGetallParams()
	log.Printf("the value of sap image query is set to %t", sapimage)
	params.Sap = &sapimage
	resp, err := f.session.Power.PCloudImages.PcloudImagesGetall(params, ibmpisession.NewAuth(f.session, powerinstanceid))
	if err != nil || resp.Payload == nil {
		log.Printf("Failed to perform the operation... %v", err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

// Get a single SAP Image
