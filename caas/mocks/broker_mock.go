// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/caas (interfaces: Broker)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	caas "github.com/juju/juju/caas"
	application "github.com/juju/juju/core/application"
	constraints "github.com/juju/juju/core/constraints"
	watcher "github.com/juju/juju/core/watcher"
	environs "github.com/juju/juju/environs"
	config "github.com/juju/juju/environs/config"
	context "github.com/juju/juju/environs/context"
	storage "github.com/juju/juju/storage"
	names "github.com/juju/names/v4"
	version "github.com/juju/version"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
)

// MockBroker is a mock of Broker interface
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// APIVersion mocks base method
func (m *MockBroker) APIVersion() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIVersion indicates an expected call of APIVersion
func (mr *MockBrokerMockRecorder) APIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockBroker)(nil).APIVersion))
}

// AdoptResources mocks base method
func (m *MockBroker) AdoptResources(arg0 context.ProviderCallContext, arg1 string, arg2 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdoptResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AdoptResources indicates an expected call of AdoptResources
func (mr *MockBrokerMockRecorder) AdoptResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdoptResources", reflect.TypeOf((*MockBroker)(nil).AdoptResources), arg0, arg1, arg2)
}

// AnnotateUnit mocks base method
func (m *MockBroker) AnnotateUnit(arg0 string, arg1 caas.DeploymentMode, arg2 string, arg3 names.UnitTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnotateUnit", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AnnotateUnit indicates an expected call of AnnotateUnit
func (mr *MockBrokerMockRecorder) AnnotateUnit(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnotateUnit", reflect.TypeOf((*MockBroker)(nil).AnnotateUnit), arg0, arg1, arg2, arg3)
}

// Application mocks base method
func (m *MockBroker) Application(arg0 string, arg1 caas.DeploymentType) caas.Application {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0, arg1)
	ret0, _ := ret[0].(caas.Application)
	return ret0
}

// Application indicates an expected call of Application
func (mr *MockBrokerMockRecorder) Application(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBroker)(nil).Application), arg0, arg1)
}

// Bootstrap mocks base method
func (m *MockBroker) Bootstrap(arg0 environs.BootstrapContext, arg1 context.ProviderCallContext, arg2 environs.BootstrapParams) (*environs.BootstrapResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*environs.BootstrapResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bootstrap indicates an expected call of Bootstrap
func (mr *MockBrokerMockRecorder) Bootstrap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockBroker)(nil).Bootstrap), arg0, arg1, arg2)
}

// CheckDefaultWorkloadStorage mocks base method
func (m *MockBroker) CheckDefaultWorkloadStorage(arg0 string, arg1 *caas.StorageProvisioner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDefaultWorkloadStorage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDefaultWorkloadStorage indicates an expected call of CheckDefaultWorkloadStorage
func (mr *MockBrokerMockRecorder) CheckDefaultWorkloadStorage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDefaultWorkloadStorage", reflect.TypeOf((*MockBroker)(nil).CheckDefaultWorkloadStorage), arg0, arg1)
}

// Config mocks base method
func (m *MockBroker) Config() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// Config indicates an expected call of Config
func (mr *MockBrokerMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockBroker)(nil).Config))
}

// ConstraintsValidator mocks base method
func (m *MockBroker) ConstraintsValidator(arg0 context.ProviderCallContext) (constraints.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstraintsValidator", arg0)
	ret0, _ := ret[0].(constraints.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConstraintsValidator indicates an expected call of ConstraintsValidator
func (mr *MockBrokerMockRecorder) ConstraintsValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstraintsValidator", reflect.TypeOf((*MockBroker)(nil).ConstraintsValidator), arg0)
}

// Create mocks base method
func (m *MockBroker) Create(arg0 context.ProviderCallContext, arg1 environs.CreateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockBrokerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBroker)(nil).Create), arg0, arg1)
}

// DeleteOperator mocks base method
func (m *MockBroker) DeleteOperator(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOperator", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOperator indicates an expected call of DeleteOperator
func (mr *MockBrokerMockRecorder) DeleteOperator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOperator", reflect.TypeOf((*MockBroker)(nil).DeleteOperator), arg0)
}

// DeleteService mocks base method
func (m *MockBroker) DeleteService(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockBrokerMockRecorder) DeleteService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockBroker)(nil).DeleteService), arg0)
}

// Destroy mocks base method
func (m *MockBroker) Destroy(arg0 context.ProviderCallContext) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockBrokerMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockBroker)(nil).Destroy), arg0)
}

// DestroyController mocks base method
func (m *MockBroker) DestroyController(arg0 context.ProviderCallContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DestroyController", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DestroyController indicates an expected call of DestroyController
func (mr *MockBrokerMockRecorder) DestroyController(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DestroyController", reflect.TypeOf((*MockBroker)(nil).DestroyController), arg0, arg1)
}

// EnsureModelOperator mocks base method
func (m *MockBroker) EnsureModelOperator(arg0, arg1 string, arg2 *caas.ModelOperatorConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureModelOperator", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureModelOperator indicates an expected call of EnsureModelOperator
func (mr *MockBrokerMockRecorder) EnsureModelOperator(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureModelOperator", reflect.TypeOf((*MockBroker)(nil).EnsureModelOperator), arg0, arg1, arg2)
}

// EnsureOperator mocks base method
func (m *MockBroker) EnsureOperator(arg0, arg1 string, arg2 *caas.OperatorConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureOperator", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureOperator indicates an expected call of EnsureOperator
func (mr *MockBrokerMockRecorder) EnsureOperator(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureOperator", reflect.TypeOf((*MockBroker)(nil).EnsureOperator), arg0, arg1, arg2)
}

// EnsureService mocks base method
func (m *MockBroker) EnsureService(arg0 string, arg1 caas.StatusCallbackFunc, arg2 *caas.ServiceParams, arg3 int, arg4 application.ConfigAttributes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureService", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureService indicates an expected call of EnsureService
func (mr *MockBrokerMockRecorder) EnsureService(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureService", reflect.TypeOf((*MockBroker)(nil).EnsureService), arg0, arg1, arg2, arg3, arg4)
}

// EnsureStorageProvisioner mocks base method
func (m *MockBroker) EnsureStorageProvisioner(arg0 caas.StorageProvisioner) (*caas.StorageProvisioner, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureStorageProvisioner", arg0)
	ret0, _ := ret[0].(*caas.StorageProvisioner)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// EnsureStorageProvisioner indicates an expected call of EnsureStorageProvisioner
func (mr *MockBrokerMockRecorder) EnsureStorageProvisioner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureStorageProvisioner", reflect.TypeOf((*MockBroker)(nil).EnsureStorageProvisioner), arg0)
}

// ExposeService mocks base method
func (m *MockBroker) ExposeService(arg0 string, arg1 map[string]string, arg2 application.ConfigAttributes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExposeService", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExposeService indicates an expected call of ExposeService
func (mr *MockBrokerMockRecorder) ExposeService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExposeService", reflect.TypeOf((*MockBroker)(nil).ExposeService), arg0, arg1, arg2)
}

// GetClusterMetadata mocks base method
func (m *MockBroker) GetClusterMetadata(arg0 string) (*caas.ClusterMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterMetadata", arg0)
	ret0, _ := ret[0].(*caas.ClusterMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterMetadata indicates an expected call of GetClusterMetadata
func (mr *MockBrokerMockRecorder) GetClusterMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMetadata", reflect.TypeOf((*MockBroker)(nil).GetClusterMetadata), arg0)
}

// GetCurrentNamespace mocks base method
func (m *MockBroker) GetCurrentNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCurrentNamespace indicates an expected call of GetCurrentNamespace
func (mr *MockBrokerMockRecorder) GetCurrentNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentNamespace", reflect.TypeOf((*MockBroker)(nil).GetCurrentNamespace))
}

// GetNamespace mocks base method
func (m *MockBroker) GetNamespace(arg0 string) (*v1.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0)
	ret0, _ := ret[0].(*v1.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockBrokerMockRecorder) GetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockBroker)(nil).GetNamespace), arg0)
}

// GetService mocks base method
func (m *MockBroker) GetService(arg0 string, arg1 caas.DeploymentMode, arg2 bool) (*caas.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*caas.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockBrokerMockRecorder) GetService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockBroker)(nil).GetService), arg0, arg1, arg2)
}

// ModelOperator mocks base method
func (m *MockBroker) ModelOperator() (*caas.ModelOperatorConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOperator")
	ret0, _ := ret[0].(*caas.ModelOperatorConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOperator indicates an expected call of ModelOperator
func (mr *MockBrokerMockRecorder) ModelOperator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOperator", reflect.TypeOf((*MockBroker)(nil).ModelOperator))
}

// ModelOperatorExists mocks base method
func (m *MockBroker) ModelOperatorExists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOperatorExists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOperatorExists indicates an expected call of ModelOperatorExists
func (mr *MockBrokerMockRecorder) ModelOperatorExists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOperatorExists", reflect.TypeOf((*MockBroker)(nil).ModelOperatorExists))
}

// Namespaces mocks base method
func (m *MockBroker) Namespaces() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespaces")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Namespaces indicates an expected call of Namespaces
func (mr *MockBrokerMockRecorder) Namespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespaces", reflect.TypeOf((*MockBroker)(nil).Namespaces))
}

// Operator mocks base method
func (m *MockBroker) Operator(arg0 string) (*caas.Operator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Operator", arg0)
	ret0, _ := ret[0].(*caas.Operator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Operator indicates an expected call of Operator
func (mr *MockBrokerMockRecorder) Operator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Operator", reflect.TypeOf((*MockBroker)(nil).Operator), arg0)
}

// OperatorExists mocks base method
func (m *MockBroker) OperatorExists(arg0 string) (caas.DeploymentState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperatorExists", arg0)
	ret0, _ := ret[0].(caas.DeploymentState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperatorExists indicates an expected call of OperatorExists
func (mr *MockBrokerMockRecorder) OperatorExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperatorExists", reflect.TypeOf((*MockBroker)(nil).OperatorExists), arg0)
}

// PrecheckInstance mocks base method
func (m *MockBroker) PrecheckInstance(arg0 context.ProviderCallContext, arg1 environs.PrecheckInstanceParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrecheckInstance", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrecheckInstance indicates an expected call of PrecheckInstance
func (mr *MockBrokerMockRecorder) PrecheckInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrecheckInstance", reflect.TypeOf((*MockBroker)(nil).PrecheckInstance), arg0, arg1)
}

// PrepareForBootstrap mocks base method
func (m *MockBroker) PrepareForBootstrap(arg0 environs.BootstrapContext, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForBootstrap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForBootstrap indicates an expected call of PrepareForBootstrap
func (mr *MockBrokerMockRecorder) PrepareForBootstrap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForBootstrap", reflect.TypeOf((*MockBroker)(nil).PrepareForBootstrap), arg0, arg1)
}

// Provider mocks base method
func (m *MockBroker) Provider() caas.ContainerEnvironProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(caas.ContainerEnvironProvider)
	return ret0
}

// Provider indicates an expected call of Provider
func (mr *MockBrokerMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockBroker)(nil).Provider))
}

// SetConfig mocks base method
func (m *MockBroker) SetConfig(arg0 *config.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig
func (mr *MockBrokerMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockBroker)(nil).SetConfig), arg0)
}

// StorageProvider mocks base method
func (m *MockBroker) StorageProvider(arg0 storage.ProviderType) (storage.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProvider", arg0)
	ret0, _ := ret[0].(storage.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProvider indicates an expected call of StorageProvider
func (mr *MockBrokerMockRecorder) StorageProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProvider", reflect.TypeOf((*MockBroker)(nil).StorageProvider), arg0)
}

// StorageProviderTypes mocks base method
func (m *MockBroker) StorageProviderTypes() ([]storage.ProviderType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageProviderTypes")
	ret0, _ := ret[0].([]storage.ProviderType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageProviderTypes indicates an expected call of StorageProviderTypes
func (mr *MockBrokerMockRecorder) StorageProviderTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageProviderTypes", reflect.TypeOf((*MockBroker)(nil).StorageProviderTypes))
}

// UnexposeService mocks base method
func (m *MockBroker) UnexposeService(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnexposeService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnexposeService indicates an expected call of UnexposeService
func (mr *MockBrokerMockRecorder) UnexposeService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnexposeService", reflect.TypeOf((*MockBroker)(nil).UnexposeService), arg0)
}

// Units mocks base method
func (m *MockBroker) Units(arg0 string, arg1 caas.DeploymentMode) ([]caas.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units", arg0, arg1)
	ret0, _ := ret[0].([]caas.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockBrokerMockRecorder) Units(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockBroker)(nil).Units), arg0, arg1)
}

// Upgrade mocks base method
func (m *MockBroker) Upgrade(arg0 string, arg1 version.Number) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade
func (mr *MockBrokerMockRecorder) Upgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockBroker)(nil).Upgrade), arg0, arg1)
}

// ValidateStorageClass mocks base method
func (m *MockBroker) ValidateStorageClass(arg0 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStorageClass", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStorageClass indicates an expected call of ValidateStorageClass
func (mr *MockBrokerMockRecorder) ValidateStorageClass(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStorageClass", reflect.TypeOf((*MockBroker)(nil).ValidateStorageClass), arg0)
}

// Version mocks base method
func (m *MockBroker) Version() (*version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockBrokerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockBroker)(nil).Version))
}

// WatchContainerStart mocks base method
func (m *MockBroker) WatchContainerStart(arg0, arg1 string) (watcher.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchContainerStart", arg0, arg1)
	ret0, _ := ret[0].(watcher.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchContainerStart indicates an expected call of WatchContainerStart
func (mr *MockBrokerMockRecorder) WatchContainerStart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchContainerStart", reflect.TypeOf((*MockBroker)(nil).WatchContainerStart), arg0, arg1)
}

// WatchNamespace mocks base method
func (m *MockBroker) WatchNamespace() (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchNamespace")
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchNamespace indicates an expected call of WatchNamespace
func (mr *MockBrokerMockRecorder) WatchNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchNamespace", reflect.TypeOf((*MockBroker)(nil).WatchNamespace))
}

// WatchOperator mocks base method
func (m *MockBroker) WatchOperator(arg0 string) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchOperator", arg0)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchOperator indicates an expected call of WatchOperator
func (mr *MockBrokerMockRecorder) WatchOperator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOperator", reflect.TypeOf((*MockBroker)(nil).WatchOperator), arg0)
}

// WatchService mocks base method
func (m *MockBroker) WatchService(arg0 string, arg1 caas.DeploymentMode) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchService", arg0, arg1)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchService indicates an expected call of WatchService
func (mr *MockBrokerMockRecorder) WatchService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchService", reflect.TypeOf((*MockBroker)(nil).WatchService), arg0, arg1)
}

// WatchUnits mocks base method
func (m *MockBroker) WatchUnits(arg0 string, arg1 caas.DeploymentMode) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUnits", arg0, arg1)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUnits indicates an expected call of WatchUnits
func (mr *MockBrokerMockRecorder) WatchUnits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUnits", reflect.TypeOf((*MockBroker)(nil).WatchUnits), arg0, arg1)
}
