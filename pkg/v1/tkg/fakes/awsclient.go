// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"sigs.k8s.io/cluster-api-provider-aws/cmd/clusterawsadm/cloudformation/bootstrap"

	"github.com/vmware-tanzu/tanzu-framework/tkg/aws"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/web/server/models"
)

type AWSClient struct {
	CreateCloudFormationStackStub        func() error
	createCloudFormationStackMutex       sync.RWMutex
	createCloudFormationStackArgsForCall []struct {
	}
	createCloudFormationStackReturns struct {
		result1 error
	}
	createCloudFormationStackReturnsOnCall map[int]struct {
		result1 error
	}
	CreateCloudFormationStackWithTemplateStub        func(*bootstrap.Template) error
	createCloudFormationStackWithTemplateMutex       sync.RWMutex
	createCloudFormationStackWithTemplateArgsForCall []struct {
		arg1 *bootstrap.Template
	}
	createCloudFormationStackWithTemplateReturns struct {
		result1 error
	}
	createCloudFormationStackWithTemplateReturnsOnCall map[int]struct {
		result1 error
	}
	EncodeCredentialsStub        func() (string, error)
	encodeCredentialsMutex       sync.RWMutex
	encodeCredentialsArgsForCall []struct {
	}
	encodeCredentialsReturns struct {
		result1 string
		result2 error
	}
	encodeCredentialsReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GenerateBootstrapTemplateStub        func(aws.GenerateBootstrapTemplateInput) (*bootstrap.Template, error)
	generateBootstrapTemplateMutex       sync.RWMutex
	generateBootstrapTemplateArgsForCall []struct {
		arg1 aws.GenerateBootstrapTemplateInput
	}
	generateBootstrapTemplateReturns struct {
		result1 *bootstrap.Template
		result2 error
	}
	generateBootstrapTemplateReturnsOnCall map[int]struct {
		result1 *bootstrap.Template
		result2 error
	}
	GetSubnetGatewayAssociationsStub        func(string) (map[string]bool, error)
	getSubnetGatewayAssociationsMutex       sync.RWMutex
	getSubnetGatewayAssociationsArgsForCall []struct {
		arg1 string
	}
	getSubnetGatewayAssociationsReturns struct {
		result1 map[string]bool
		result2 error
	}
	getSubnetGatewayAssociationsReturnsOnCall map[int]struct {
		result1 map[string]bool
		result2 error
	}
	ListAvailabilityZonesStub        func() ([]*models.AWSAvailabilityZone, error)
	listAvailabilityZonesMutex       sync.RWMutex
	listAvailabilityZonesArgsForCall []struct {
	}
	listAvailabilityZonesReturns struct {
		result1 []*models.AWSAvailabilityZone
		result2 error
	}
	listAvailabilityZonesReturnsOnCall map[int]struct {
		result1 []*models.AWSAvailabilityZone
		result2 error
	}
	ListCloudFormationStacksStub        func() ([]string, error)
	listCloudFormationStacksMutex       sync.RWMutex
	listCloudFormationStacksArgsForCall []struct {
	}
	listCloudFormationStacksReturns struct {
		result1 []string
		result2 error
	}
	listCloudFormationStacksReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ListInstanceTypesStub        func(string) ([]string, error)
	listInstanceTypesMutex       sync.RWMutex
	listInstanceTypesArgsForCall []struct {
		arg1 string
	}
	listInstanceTypesReturns struct {
		result1 []string
		result2 error
	}
	listInstanceTypesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ListRegionsByUserStub        func() ([]string, error)
	listRegionsByUserMutex       sync.RWMutex
	listRegionsByUserArgsForCall []struct {
	}
	listRegionsByUserReturns struct {
		result1 []string
		result2 error
	}
	listRegionsByUserReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ListSubnetsStub        func(string) ([]*models.AWSSubnet, error)
	listSubnetsMutex       sync.RWMutex
	listSubnetsArgsForCall []struct {
		arg1 string
	}
	listSubnetsReturns struct {
		result1 []*models.AWSSubnet
		result2 error
	}
	listSubnetsReturnsOnCall map[int]struct {
		result1 []*models.AWSSubnet
		result2 error
	}
	ListVPCsStub        func() ([]*models.Vpc, error)
	listVPCsMutex       sync.RWMutex
	listVPCsArgsForCall []struct {
	}
	listVPCsReturns struct {
		result1 []*models.Vpc
		result2 error
	}
	listVPCsReturnsOnCall map[int]struct {
		result1 []*models.Vpc
		result2 error
	}
	VerifyAccountStub        func() error
	verifyAccountMutex       sync.RWMutex
	verifyAccountArgsForCall []struct {
	}
	verifyAccountReturns struct {
		result1 error
	}
	verifyAccountReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AWSClient) CreateCloudFormationStack() error {
	fake.createCloudFormationStackMutex.Lock()
	ret, specificReturn := fake.createCloudFormationStackReturnsOnCall[len(fake.createCloudFormationStackArgsForCall)]
	fake.createCloudFormationStackArgsForCall = append(fake.createCloudFormationStackArgsForCall, struct {
	}{})
	stub := fake.CreateCloudFormationStackStub
	fakeReturns := fake.createCloudFormationStackReturns
	fake.recordInvocation("CreateCloudFormationStack", []interface{}{})
	fake.createCloudFormationStackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *AWSClient) CreateCloudFormationStackCallCount() int {
	fake.createCloudFormationStackMutex.RLock()
	defer fake.createCloudFormationStackMutex.RUnlock()
	return len(fake.createCloudFormationStackArgsForCall)
}

func (fake *AWSClient) CreateCloudFormationStackCalls(stub func() error) {
	fake.createCloudFormationStackMutex.Lock()
	defer fake.createCloudFormationStackMutex.Unlock()
	fake.CreateCloudFormationStackStub = stub
}

func (fake *AWSClient) CreateCloudFormationStackReturns(result1 error) {
	fake.createCloudFormationStackMutex.Lock()
	defer fake.createCloudFormationStackMutex.Unlock()
	fake.CreateCloudFormationStackStub = nil
	fake.createCloudFormationStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *AWSClient) CreateCloudFormationStackReturnsOnCall(i int, result1 error) {
	fake.createCloudFormationStackMutex.Lock()
	defer fake.createCloudFormationStackMutex.Unlock()
	fake.CreateCloudFormationStackStub = nil
	if fake.createCloudFormationStackReturnsOnCall == nil {
		fake.createCloudFormationStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createCloudFormationStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *AWSClient) CreateCloudFormationStackWithTemplate(arg1 *bootstrap.Template) error {
	fake.createCloudFormationStackWithTemplateMutex.Lock()
	ret, specificReturn := fake.createCloudFormationStackWithTemplateReturnsOnCall[len(fake.createCloudFormationStackWithTemplateArgsForCall)]
	fake.createCloudFormationStackWithTemplateArgsForCall = append(fake.createCloudFormationStackWithTemplateArgsForCall, struct {
		arg1 *bootstrap.Template
	}{arg1})
	stub := fake.CreateCloudFormationStackWithTemplateStub
	fakeReturns := fake.createCloudFormationStackWithTemplateReturns
	fake.recordInvocation("CreateCloudFormationStackWithTemplate", []interface{}{arg1})
	fake.createCloudFormationStackWithTemplateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *AWSClient) CreateCloudFormationStackWithTemplateCallCount() int {
	fake.createCloudFormationStackWithTemplateMutex.RLock()
	defer fake.createCloudFormationStackWithTemplateMutex.RUnlock()
	return len(fake.createCloudFormationStackWithTemplateArgsForCall)
}

func (fake *AWSClient) CreateCloudFormationStackWithTemplateCalls(stub func(*bootstrap.Template) error) {
	fake.createCloudFormationStackWithTemplateMutex.Lock()
	defer fake.createCloudFormationStackWithTemplateMutex.Unlock()
	fake.CreateCloudFormationStackWithTemplateStub = stub
}

func (fake *AWSClient) CreateCloudFormationStackWithTemplateArgsForCall(i int) *bootstrap.Template {
	fake.createCloudFormationStackWithTemplateMutex.RLock()
	defer fake.createCloudFormationStackWithTemplateMutex.RUnlock()
	argsForCall := fake.createCloudFormationStackWithTemplateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AWSClient) CreateCloudFormationStackWithTemplateReturns(result1 error) {
	fake.createCloudFormationStackWithTemplateMutex.Lock()
	defer fake.createCloudFormationStackWithTemplateMutex.Unlock()
	fake.CreateCloudFormationStackWithTemplateStub = nil
	fake.createCloudFormationStackWithTemplateReturns = struct {
		result1 error
	}{result1}
}

func (fake *AWSClient) CreateCloudFormationStackWithTemplateReturnsOnCall(i int, result1 error) {
	fake.createCloudFormationStackWithTemplateMutex.Lock()
	defer fake.createCloudFormationStackWithTemplateMutex.Unlock()
	fake.CreateCloudFormationStackWithTemplateStub = nil
	if fake.createCloudFormationStackWithTemplateReturnsOnCall == nil {
		fake.createCloudFormationStackWithTemplateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createCloudFormationStackWithTemplateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *AWSClient) EncodeCredentials() (string, error) {
	fake.encodeCredentialsMutex.Lock()
	ret, specificReturn := fake.encodeCredentialsReturnsOnCall[len(fake.encodeCredentialsArgsForCall)]
	fake.encodeCredentialsArgsForCall = append(fake.encodeCredentialsArgsForCall, struct {
	}{})
	stub := fake.EncodeCredentialsStub
	fakeReturns := fake.encodeCredentialsReturns
	fake.recordInvocation("EncodeCredentials", []interface{}{})
	fake.encodeCredentialsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) EncodeCredentialsCallCount() int {
	fake.encodeCredentialsMutex.RLock()
	defer fake.encodeCredentialsMutex.RUnlock()
	return len(fake.encodeCredentialsArgsForCall)
}

func (fake *AWSClient) EncodeCredentialsCalls(stub func() (string, error)) {
	fake.encodeCredentialsMutex.Lock()
	defer fake.encodeCredentialsMutex.Unlock()
	fake.EncodeCredentialsStub = stub
}

func (fake *AWSClient) EncodeCredentialsReturns(result1 string, result2 error) {
	fake.encodeCredentialsMutex.Lock()
	defer fake.encodeCredentialsMutex.Unlock()
	fake.EncodeCredentialsStub = nil
	fake.encodeCredentialsReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) EncodeCredentialsReturnsOnCall(i int, result1 string, result2 error) {
	fake.encodeCredentialsMutex.Lock()
	defer fake.encodeCredentialsMutex.Unlock()
	fake.EncodeCredentialsStub = nil
	if fake.encodeCredentialsReturnsOnCall == nil {
		fake.encodeCredentialsReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.encodeCredentialsReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) GenerateBootstrapTemplate(arg1 aws.GenerateBootstrapTemplateInput) (*bootstrap.Template, error) {
	fake.generateBootstrapTemplateMutex.Lock()
	ret, specificReturn := fake.generateBootstrapTemplateReturnsOnCall[len(fake.generateBootstrapTemplateArgsForCall)]
	fake.generateBootstrapTemplateArgsForCall = append(fake.generateBootstrapTemplateArgsForCall, struct {
		arg1 aws.GenerateBootstrapTemplateInput
	}{arg1})
	stub := fake.GenerateBootstrapTemplateStub
	fakeReturns := fake.generateBootstrapTemplateReturns
	fake.recordInvocation("GenerateBootstrapTemplate", []interface{}{arg1})
	fake.generateBootstrapTemplateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) GenerateBootstrapTemplateCallCount() int {
	fake.generateBootstrapTemplateMutex.RLock()
	defer fake.generateBootstrapTemplateMutex.RUnlock()
	return len(fake.generateBootstrapTemplateArgsForCall)
}

func (fake *AWSClient) GenerateBootstrapTemplateCalls(stub func(aws.GenerateBootstrapTemplateInput) (*bootstrap.Template, error)) {
	fake.generateBootstrapTemplateMutex.Lock()
	defer fake.generateBootstrapTemplateMutex.Unlock()
	fake.GenerateBootstrapTemplateStub = stub
}

func (fake *AWSClient) GenerateBootstrapTemplateArgsForCall(i int) aws.GenerateBootstrapTemplateInput {
	fake.generateBootstrapTemplateMutex.RLock()
	defer fake.generateBootstrapTemplateMutex.RUnlock()
	argsForCall := fake.generateBootstrapTemplateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AWSClient) GenerateBootstrapTemplateReturns(result1 *bootstrap.Template, result2 error) {
	fake.generateBootstrapTemplateMutex.Lock()
	defer fake.generateBootstrapTemplateMutex.Unlock()
	fake.GenerateBootstrapTemplateStub = nil
	fake.generateBootstrapTemplateReturns = struct {
		result1 *bootstrap.Template
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) GenerateBootstrapTemplateReturnsOnCall(i int, result1 *bootstrap.Template, result2 error) {
	fake.generateBootstrapTemplateMutex.Lock()
	defer fake.generateBootstrapTemplateMutex.Unlock()
	fake.GenerateBootstrapTemplateStub = nil
	if fake.generateBootstrapTemplateReturnsOnCall == nil {
		fake.generateBootstrapTemplateReturnsOnCall = make(map[int]struct {
			result1 *bootstrap.Template
			result2 error
		})
	}
	fake.generateBootstrapTemplateReturnsOnCall[i] = struct {
		result1 *bootstrap.Template
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) GetSubnetGatewayAssociations(arg1 string) (map[string]bool, error) {
	fake.getSubnetGatewayAssociationsMutex.Lock()
	ret, specificReturn := fake.getSubnetGatewayAssociationsReturnsOnCall[len(fake.getSubnetGatewayAssociationsArgsForCall)]
	fake.getSubnetGatewayAssociationsArgsForCall = append(fake.getSubnetGatewayAssociationsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetSubnetGatewayAssociationsStub
	fakeReturns := fake.getSubnetGatewayAssociationsReturns
	fake.recordInvocation("GetSubnetGatewayAssociations", []interface{}{arg1})
	fake.getSubnetGatewayAssociationsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) GetSubnetGatewayAssociationsCallCount() int {
	fake.getSubnetGatewayAssociationsMutex.RLock()
	defer fake.getSubnetGatewayAssociationsMutex.RUnlock()
	return len(fake.getSubnetGatewayAssociationsArgsForCall)
}

func (fake *AWSClient) GetSubnetGatewayAssociationsCalls(stub func(string) (map[string]bool, error)) {
	fake.getSubnetGatewayAssociationsMutex.Lock()
	defer fake.getSubnetGatewayAssociationsMutex.Unlock()
	fake.GetSubnetGatewayAssociationsStub = stub
}

func (fake *AWSClient) GetSubnetGatewayAssociationsArgsForCall(i int) string {
	fake.getSubnetGatewayAssociationsMutex.RLock()
	defer fake.getSubnetGatewayAssociationsMutex.RUnlock()
	argsForCall := fake.getSubnetGatewayAssociationsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AWSClient) GetSubnetGatewayAssociationsReturns(result1 map[string]bool, result2 error) {
	fake.getSubnetGatewayAssociationsMutex.Lock()
	defer fake.getSubnetGatewayAssociationsMutex.Unlock()
	fake.GetSubnetGatewayAssociationsStub = nil
	fake.getSubnetGatewayAssociationsReturns = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) GetSubnetGatewayAssociationsReturnsOnCall(i int, result1 map[string]bool, result2 error) {
	fake.getSubnetGatewayAssociationsMutex.Lock()
	defer fake.getSubnetGatewayAssociationsMutex.Unlock()
	fake.GetSubnetGatewayAssociationsStub = nil
	if fake.getSubnetGatewayAssociationsReturnsOnCall == nil {
		fake.getSubnetGatewayAssociationsReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
			result2 error
		})
	}
	fake.getSubnetGatewayAssociationsReturnsOnCall[i] = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListAvailabilityZones() ([]*models.AWSAvailabilityZone, error) {
	fake.listAvailabilityZonesMutex.Lock()
	ret, specificReturn := fake.listAvailabilityZonesReturnsOnCall[len(fake.listAvailabilityZonesArgsForCall)]
	fake.listAvailabilityZonesArgsForCall = append(fake.listAvailabilityZonesArgsForCall, struct {
	}{})
	stub := fake.ListAvailabilityZonesStub
	fakeReturns := fake.listAvailabilityZonesReturns
	fake.recordInvocation("ListAvailabilityZones", []interface{}{})
	fake.listAvailabilityZonesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) ListAvailabilityZonesCallCount() int {
	fake.listAvailabilityZonesMutex.RLock()
	defer fake.listAvailabilityZonesMutex.RUnlock()
	return len(fake.listAvailabilityZonesArgsForCall)
}

func (fake *AWSClient) ListAvailabilityZonesCalls(stub func() ([]*models.AWSAvailabilityZone, error)) {
	fake.listAvailabilityZonesMutex.Lock()
	defer fake.listAvailabilityZonesMutex.Unlock()
	fake.ListAvailabilityZonesStub = stub
}

func (fake *AWSClient) ListAvailabilityZonesReturns(result1 []*models.AWSAvailabilityZone, result2 error) {
	fake.listAvailabilityZonesMutex.Lock()
	defer fake.listAvailabilityZonesMutex.Unlock()
	fake.ListAvailabilityZonesStub = nil
	fake.listAvailabilityZonesReturns = struct {
		result1 []*models.AWSAvailabilityZone
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListAvailabilityZonesReturnsOnCall(i int, result1 []*models.AWSAvailabilityZone, result2 error) {
	fake.listAvailabilityZonesMutex.Lock()
	defer fake.listAvailabilityZonesMutex.Unlock()
	fake.ListAvailabilityZonesStub = nil
	if fake.listAvailabilityZonesReturnsOnCall == nil {
		fake.listAvailabilityZonesReturnsOnCall = make(map[int]struct {
			result1 []*models.AWSAvailabilityZone
			result2 error
		})
	}
	fake.listAvailabilityZonesReturnsOnCall[i] = struct {
		result1 []*models.AWSAvailabilityZone
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListCloudFormationStacks() ([]string, error) {
	fake.listCloudFormationStacksMutex.Lock()
	ret, specificReturn := fake.listCloudFormationStacksReturnsOnCall[len(fake.listCloudFormationStacksArgsForCall)]
	fake.listCloudFormationStacksArgsForCall = append(fake.listCloudFormationStacksArgsForCall, struct {
	}{})
	stub := fake.ListCloudFormationStacksStub
	fakeReturns := fake.listCloudFormationStacksReturns
	fake.recordInvocation("ListCloudFormationStacks", []interface{}{})
	fake.listCloudFormationStacksMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) ListCloudFormationStacksCallCount() int {
	fake.listCloudFormationStacksMutex.RLock()
	defer fake.listCloudFormationStacksMutex.RUnlock()
	return len(fake.listCloudFormationStacksArgsForCall)
}

func (fake *AWSClient) ListCloudFormationStacksCalls(stub func() ([]string, error)) {
	fake.listCloudFormationStacksMutex.Lock()
	defer fake.listCloudFormationStacksMutex.Unlock()
	fake.ListCloudFormationStacksStub = stub
}

func (fake *AWSClient) ListCloudFormationStacksReturns(result1 []string, result2 error) {
	fake.listCloudFormationStacksMutex.Lock()
	defer fake.listCloudFormationStacksMutex.Unlock()
	fake.ListCloudFormationStacksStub = nil
	fake.listCloudFormationStacksReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListCloudFormationStacksReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listCloudFormationStacksMutex.Lock()
	defer fake.listCloudFormationStacksMutex.Unlock()
	fake.ListCloudFormationStacksStub = nil
	if fake.listCloudFormationStacksReturnsOnCall == nil {
		fake.listCloudFormationStacksReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listCloudFormationStacksReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListInstanceTypes(arg1 string) ([]string, error) {
	fake.listInstanceTypesMutex.Lock()
	ret, specificReturn := fake.listInstanceTypesReturnsOnCall[len(fake.listInstanceTypesArgsForCall)]
	fake.listInstanceTypesArgsForCall = append(fake.listInstanceTypesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListInstanceTypesStub
	fakeReturns := fake.listInstanceTypesReturns
	fake.recordInvocation("ListInstanceTypes", []interface{}{arg1})
	fake.listInstanceTypesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) ListInstanceTypesCallCount() int {
	fake.listInstanceTypesMutex.RLock()
	defer fake.listInstanceTypesMutex.RUnlock()
	return len(fake.listInstanceTypesArgsForCall)
}

func (fake *AWSClient) ListInstanceTypesCalls(stub func(string) ([]string, error)) {
	fake.listInstanceTypesMutex.Lock()
	defer fake.listInstanceTypesMutex.Unlock()
	fake.ListInstanceTypesStub = stub
}

func (fake *AWSClient) ListInstanceTypesArgsForCall(i int) string {
	fake.listInstanceTypesMutex.RLock()
	defer fake.listInstanceTypesMutex.RUnlock()
	argsForCall := fake.listInstanceTypesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AWSClient) ListInstanceTypesReturns(result1 []string, result2 error) {
	fake.listInstanceTypesMutex.Lock()
	defer fake.listInstanceTypesMutex.Unlock()
	fake.ListInstanceTypesStub = nil
	fake.listInstanceTypesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListInstanceTypesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listInstanceTypesMutex.Lock()
	defer fake.listInstanceTypesMutex.Unlock()
	fake.ListInstanceTypesStub = nil
	if fake.listInstanceTypesReturnsOnCall == nil {
		fake.listInstanceTypesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listInstanceTypesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListRegionsByUser() ([]string, error) {
	fake.listRegionsByUserMutex.Lock()
	ret, specificReturn := fake.listRegionsByUserReturnsOnCall[len(fake.listRegionsByUserArgsForCall)]
	fake.listRegionsByUserArgsForCall = append(fake.listRegionsByUserArgsForCall, struct {
	}{})
	stub := fake.ListRegionsByUserStub
	fakeReturns := fake.listRegionsByUserReturns
	fake.recordInvocation("ListRegionsByUser", []interface{}{})
	fake.listRegionsByUserMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) ListRegionsByUserCallCount() int {
	fake.listRegionsByUserMutex.RLock()
	defer fake.listRegionsByUserMutex.RUnlock()
	return len(fake.listRegionsByUserArgsForCall)
}

func (fake *AWSClient) ListRegionsByUserCalls(stub func() ([]string, error)) {
	fake.listRegionsByUserMutex.Lock()
	defer fake.listRegionsByUserMutex.Unlock()
	fake.ListRegionsByUserStub = stub
}

func (fake *AWSClient) ListRegionsByUserReturns(result1 []string, result2 error) {
	fake.listRegionsByUserMutex.Lock()
	defer fake.listRegionsByUserMutex.Unlock()
	fake.ListRegionsByUserStub = nil
	fake.listRegionsByUserReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListRegionsByUserReturnsOnCall(i int, result1 []string, result2 error) {
	fake.listRegionsByUserMutex.Lock()
	defer fake.listRegionsByUserMutex.Unlock()
	fake.ListRegionsByUserStub = nil
	if fake.listRegionsByUserReturnsOnCall == nil {
		fake.listRegionsByUserReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listRegionsByUserReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListSubnets(arg1 string) ([]*models.AWSSubnet, error) {
	fake.listSubnetsMutex.Lock()
	ret, specificReturn := fake.listSubnetsReturnsOnCall[len(fake.listSubnetsArgsForCall)]
	fake.listSubnetsArgsForCall = append(fake.listSubnetsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ListSubnetsStub
	fakeReturns := fake.listSubnetsReturns
	fake.recordInvocation("ListSubnets", []interface{}{arg1})
	fake.listSubnetsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) ListSubnetsCallCount() int {
	fake.listSubnetsMutex.RLock()
	defer fake.listSubnetsMutex.RUnlock()
	return len(fake.listSubnetsArgsForCall)
}

func (fake *AWSClient) ListSubnetsCalls(stub func(string) ([]*models.AWSSubnet, error)) {
	fake.listSubnetsMutex.Lock()
	defer fake.listSubnetsMutex.Unlock()
	fake.ListSubnetsStub = stub
}

func (fake *AWSClient) ListSubnetsArgsForCall(i int) string {
	fake.listSubnetsMutex.RLock()
	defer fake.listSubnetsMutex.RUnlock()
	argsForCall := fake.listSubnetsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AWSClient) ListSubnetsReturns(result1 []*models.AWSSubnet, result2 error) {
	fake.listSubnetsMutex.Lock()
	defer fake.listSubnetsMutex.Unlock()
	fake.ListSubnetsStub = nil
	fake.listSubnetsReturns = struct {
		result1 []*models.AWSSubnet
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListSubnetsReturnsOnCall(i int, result1 []*models.AWSSubnet, result2 error) {
	fake.listSubnetsMutex.Lock()
	defer fake.listSubnetsMutex.Unlock()
	fake.ListSubnetsStub = nil
	if fake.listSubnetsReturnsOnCall == nil {
		fake.listSubnetsReturnsOnCall = make(map[int]struct {
			result1 []*models.AWSSubnet
			result2 error
		})
	}
	fake.listSubnetsReturnsOnCall[i] = struct {
		result1 []*models.AWSSubnet
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListVPCs() ([]*models.Vpc, error) {
	fake.listVPCsMutex.Lock()
	ret, specificReturn := fake.listVPCsReturnsOnCall[len(fake.listVPCsArgsForCall)]
	fake.listVPCsArgsForCall = append(fake.listVPCsArgsForCall, struct {
	}{})
	stub := fake.ListVPCsStub
	fakeReturns := fake.listVPCsReturns
	fake.recordInvocation("ListVPCs", []interface{}{})
	fake.listVPCsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AWSClient) ListVPCsCallCount() int {
	fake.listVPCsMutex.RLock()
	defer fake.listVPCsMutex.RUnlock()
	return len(fake.listVPCsArgsForCall)
}

func (fake *AWSClient) ListVPCsCalls(stub func() ([]*models.Vpc, error)) {
	fake.listVPCsMutex.Lock()
	defer fake.listVPCsMutex.Unlock()
	fake.ListVPCsStub = stub
}

func (fake *AWSClient) ListVPCsReturns(result1 []*models.Vpc, result2 error) {
	fake.listVPCsMutex.Lock()
	defer fake.listVPCsMutex.Unlock()
	fake.ListVPCsStub = nil
	fake.listVPCsReturns = struct {
		result1 []*models.Vpc
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) ListVPCsReturnsOnCall(i int, result1 []*models.Vpc, result2 error) {
	fake.listVPCsMutex.Lock()
	defer fake.listVPCsMutex.Unlock()
	fake.ListVPCsStub = nil
	if fake.listVPCsReturnsOnCall == nil {
		fake.listVPCsReturnsOnCall = make(map[int]struct {
			result1 []*models.Vpc
			result2 error
		})
	}
	fake.listVPCsReturnsOnCall[i] = struct {
		result1 []*models.Vpc
		result2 error
	}{result1, result2}
}

func (fake *AWSClient) VerifyAccount() error {
	fake.verifyAccountMutex.Lock()
	ret, specificReturn := fake.verifyAccountReturnsOnCall[len(fake.verifyAccountArgsForCall)]
	fake.verifyAccountArgsForCall = append(fake.verifyAccountArgsForCall, struct {
	}{})
	stub := fake.VerifyAccountStub
	fakeReturns := fake.verifyAccountReturns
	fake.recordInvocation("VerifyAccount", []interface{}{})
	fake.verifyAccountMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *AWSClient) VerifyAccountCallCount() int {
	fake.verifyAccountMutex.RLock()
	defer fake.verifyAccountMutex.RUnlock()
	return len(fake.verifyAccountArgsForCall)
}

func (fake *AWSClient) VerifyAccountCalls(stub func() error) {
	fake.verifyAccountMutex.Lock()
	defer fake.verifyAccountMutex.Unlock()
	fake.VerifyAccountStub = stub
}

func (fake *AWSClient) VerifyAccountReturns(result1 error) {
	fake.verifyAccountMutex.Lock()
	defer fake.verifyAccountMutex.Unlock()
	fake.VerifyAccountStub = nil
	fake.verifyAccountReturns = struct {
		result1 error
	}{result1}
}

func (fake *AWSClient) VerifyAccountReturnsOnCall(i int, result1 error) {
	fake.verifyAccountMutex.Lock()
	defer fake.verifyAccountMutex.Unlock()
	fake.VerifyAccountStub = nil
	if fake.verifyAccountReturnsOnCall == nil {
		fake.verifyAccountReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyAccountReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *AWSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCloudFormationStackMutex.RLock()
	defer fake.createCloudFormationStackMutex.RUnlock()
	fake.createCloudFormationStackWithTemplateMutex.RLock()
	defer fake.createCloudFormationStackWithTemplateMutex.RUnlock()
	fake.encodeCredentialsMutex.RLock()
	defer fake.encodeCredentialsMutex.RUnlock()
	fake.generateBootstrapTemplateMutex.RLock()
	defer fake.generateBootstrapTemplateMutex.RUnlock()
	fake.getSubnetGatewayAssociationsMutex.RLock()
	defer fake.getSubnetGatewayAssociationsMutex.RUnlock()
	fake.listAvailabilityZonesMutex.RLock()
	defer fake.listAvailabilityZonesMutex.RUnlock()
	fake.listCloudFormationStacksMutex.RLock()
	defer fake.listCloudFormationStacksMutex.RUnlock()
	fake.listInstanceTypesMutex.RLock()
	defer fake.listInstanceTypesMutex.RUnlock()
	fake.listRegionsByUserMutex.RLock()
	defer fake.listRegionsByUserMutex.RUnlock()
	fake.listSubnetsMutex.RLock()
	defer fake.listSubnetsMutex.RUnlock()
	fake.listVPCsMutex.RLock()
	defer fake.listVPCsMutex.RUnlock()
	fake.verifyAccountMutex.RLock()
	defer fake.verifyAccountMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AWSClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ aws.Client = new(AWSClient)
