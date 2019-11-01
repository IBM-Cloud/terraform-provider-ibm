package lbaas

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.ibm.com/Bluemix/riaas-go-client/errors"
	riaaserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/client/l_baas"
	"github.ibm.com/Bluemix/riaas-go-client/riaas/models"
	"github.ibm.com/Bluemix/riaas-go-client/session"
)

// LoadBalancerClient ...
type LoadBalancerClient struct {
	session *session.Session
}

// NewLoadBalancerClient ...
func NewLoadBalancerClient(sess *session.Session) *LoadBalancerClient {
	return &LoadBalancerClient{
		sess,
	}
}

// List ...
func (f *LoadBalancerClient) List() ([]*models.LoadBalancer, error) {
	return f.ListWithFilter("", "")
}

// ListWithFilter ...
func (f *LoadBalancerClient) ListWithFilter(start, resourcegroupID string) ([]*models.LoadBalancer, error) {
	params := l_baas.NewGetLoadBalancersParamsWithTimeout(f.session.Timeout)

	if start != "" {
		params = params.WithStart(&start)
	}

	if resourcegroupID != "" {
		params = params.WithResourceGroupID(&resourcegroupID)
	}
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancers(params, session.Auth(f.session))

	if err != nil {
		return nil, riaaserrors.ToError(err)
	}

	return resp.Payload.LoadBalancers, nil
}

// Create ...
func (f *LoadBalancerClient) Create(lbaasdef *l_baas.PostLoadBalancersParams) (*models.LoadBalancer, error) {
	params := l_baas.NewPostLoadBalancersParamsWithTimeout(f.session.Timeout).WithBody(lbaasdef.Body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PostLoadBalancers(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Delete ...
func (f *LoadBalancerClient) Delete(id string) error {
	params := l_baas.NewDeleteLoadBalancersIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.LBaas.DeleteLoadBalancersID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Get ...
func (f *LoadBalancerClient) Get(id string) (*models.LoadBalancer, error) {
	params := l_baas.NewGetLoadBalancersIDParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// Update ...
func (f *LoadBalancerClient) Update(id, name string) (*models.LoadBalancer, error) {
	var body = l_baas.PatchLoadBalancersIDParams{}
	if name != "" {
		body.Body = &models.LoadBalancerTemplatePatch{
			Name: name,
		}
	}

	params := l_baas.NewPatchLoadBalancersIDParamsWithTimeout(f.session.Timeout).WithID(id).WithBody(body.Body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PatchLoadBalancersID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetListeners ...
func (f *LoadBalancerClient) GetListeners(id string) (*models.ListenerCollection, error) {
	params := l_baas.NewGetLoadBalancersIDListenersParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDListeners(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreateListeners ...
func (f *LoadBalancerClient) CreateListeners(lbaasListners *l_baas.PostLoadBalancersIDListenersParams) (*models.Listener, error) {

	params := l_baas.NewPostLoadBalancersIDListenersParamsWithTimeout(f.session.Timeout).WithBody(lbaasListners.Body).WithID(lbaasListners.ID)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PostLoadBalancersIDListeners(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetPools ...
func (f *LoadBalancerClient) GetPools(id string) (*models.PoolCollection, error) {
	params := l_baas.NewGetLoadBalancersIDPoolsParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDPools(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreatePool ...
func (f *LoadBalancerClient) CreatePool(lbaasPool *l_baas.PostLoadBalancersIDPoolsParams) (*models.Pool, error) {

	params := l_baas.NewPostLoadBalancersIDPoolsParamsWithTimeout(f.session.Timeout).WithBody(lbaasPool.Body).WithID(lbaasPool.ID)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PostLoadBalancersIDPools(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetStatistics ...
func (f *LoadBalancerClient) GetStatistics(id string) (*models.LoadBalancerStatistics, error) {
	params := l_baas.NewGetLoadBalancersIDStatisticsParamsWithTimeout(f.session.Timeout).WithID(id)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDStatistics(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteListener ...
func (f *LoadBalancerClient) DeleteListener(lbaasId, listenerId string) error {
	params := l_baas.NewDeleteLoadBalancersIDListenersListenerIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithListenerID(listenerId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.LBaas.DeleteLoadBalancersIDListenersListenerID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// GetListener ...
func (f *LoadBalancerClient) GetListener(lbaasId, listenerId string) (*models.Listener, error) {
	params := l_baas.NewGetLoadBalancersIDListenersListenerIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithListenerID(listenerId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDListenersListenerID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdateListener ...
func (f *LoadBalancerClient) UpdateListener(lbaasId, listenerId, crn, protocol, poolId string, port, connectionLimit int) (*models.Listener, error) {
	body := models.ListenerTemplatePatch{}
	if crn != "" {
		body.CertificateInstance = &models.ListenerTemplatePatchCertificateInstance{
			Crn: crn,
		}
	}
	if poolId != "" {
		body.DefaultPool = &models.ListenerTemplatePatchDefaultPool{
			ID: strfmt.UUID(poolId),
		}
	}

	if connectionLimit > 0 {
		body.ConnectionLimit = int64(connectionLimit)
	}

	if port > 0 {
		body.Port = int64(port)
	}

	if protocol != "" {
		body.Protocol = protocol
	}

	params := l_baas.NewPatchLoadBalancersIDListenersListenerIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithListenerID(listenerId).WithBody(&body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PatchLoadBalancersIDListenersListenerID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeleteListener ...
func (f *LoadBalancerClient) DeletePool(lbaasId, poolId string) error {
	params := l_baas.NewDeleteLoadBalancersIDPoolsPoolIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.LBaas.DeleteLoadBalancersIDPoolsPoolID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// GetPool ...
func (f *LoadBalancerClient) GetPool(lbaasId, poolId string) (*models.Pool, error) {
	params := l_baas.NewGetLoadBalancersIDPoolsPoolIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDPoolsPoolID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdatePool ...
func (f *LoadBalancerClient) UpdatePool(lbaasId, poolId, algorithm, name, protocol string, hmTemplate models.HealthMonitorTemplate, sessionTemplate models.SessionPersistenceTemplate) (*models.Pool, error) {

	var body = models.PoolTemplatePatch{
		Algorithm: algorithm,
		Name:      name,
		Protocol:  protocol,
	}
	if sessionTemplate.Type != "" {
		body.SessionPersistence = &sessionTemplate
	}
	if hmTemplate.Type != "" {
		body.HealthMonitor = &hmTemplate
	}
	params := l_baas.NewPatchLoadBalancersIDPoolsPoolIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId).WithBody(&body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PatchLoadBalancersIDPoolsPoolID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// GetPoolMembers ...
func (f *LoadBalancerClient) GetPoolMembers(lbaasId, poolId string) (*models.MemberCollection, error) {
	params := l_baas.NewGetLoadBalancersIDPoolsPoolIDMembersParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDPoolsPoolIDMembers(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// CreatePoolMember ...
func (f *LoadBalancerClient) CreatePoolMember(lbaasId, poolId, address string, port, weight int) (*models.Member, error) {
	var memTemplate = models.MemberTemplateTarget{
		Address: address,
	}

	var body = models.MemberTemplate{
		Port:   int64(port),
		Weight: int64(weight),
		Target: &memTemplate,
	}
	params := l_baas.NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId).WithBody(&body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PostLoadBalancersIDPoolsPoolIDMembers(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// UpdatePoolMember ...
func (f *LoadBalancerClient) UpdatePoolMember(lbaasId, poolId, memberId, address string, port, weight int) (*models.Member, error) {
	var memTemplate = models.MemberTemplateCommonTarget{
		Address: address,
	}

	var body = models.MemberTemplatePatch{}
	body.Port = int64(port)
	body.Weight = int64(weight)
	body.Target = &memTemplate

	params := l_baas.NewPatchLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId).WithMemberID(memberId).WithBody(body)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.PatchLoadBalancersIDPoolsPoolIDMembersMemberID(params, session.Auth(f.session))
	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}

// DeletePoolMember ...
func (f *LoadBalancerClient) DeletePoolMember(lbaasId, poolId, memberId string) error {
	params := l_baas.NewDeleteLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId).WithMemberID(memberId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	_, err := f.session.Riaas.LBaas.DeleteLoadBalancersIDPoolsPoolIDMembersMemberID(params, session.Auth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// GetPoolMembers ...
func (f *LoadBalancerClient) GetPoolMember(lbaasId, poolId, memberId string) (*models.Member, error) {
	params := l_baas.NewGetLoadBalancersIDPoolsPoolIDMembersMemberIDParamsWithTimeout(f.session.Timeout).WithID(lbaasId).WithPoolID(poolId).WithMemberID(memberId)
	params.Version = "2019-08-27"
	params.Generation = f.session.Generation
	resp, err := f.session.Riaas.LBaas.GetLoadBalancersIDPoolsPoolIDMembersMemberID(params, session.Auth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}
