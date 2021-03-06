// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/grpc/proto/namespace/template.proto

package namespace

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "k8s.io/api/core/v1"
	v11 "k8s.io/api/rbac/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NamespaceTemplate struct {
	//exportedParamName to be exported from this template
	//These params will be passed in namespace creation and values will be replaced
	//If you are using params in the template, make sure to to include in the resources with ${exportedParamName}
	//+optional
	ExportedParamName []string `protobuf:"bytes,1,rep,name=exportedParamName,proto3" json:"exportedParamName,omitempty"`
	//NamespaceResources consists of all the resources to be created in namespace including custom resources
	//+required
	NsResources          *NamespaceResources `protobuf:"bytes,2,opt,name=nsResources,proto3" json:"nsResources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NamespaceTemplate) Reset()         { *m = NamespaceTemplate{} }
func (m *NamespaceTemplate) String() string { return proto.CompactTextString(m) }
func (*NamespaceTemplate) ProtoMessage()    {}
func (*NamespaceTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f02f0592f6a6cd, []int{0}
}

func (m *NamespaceTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceTemplate.Unmarshal(m, b)
}
func (m *NamespaceTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceTemplate.Marshal(b, m, deterministic)
}
func (m *NamespaceTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceTemplate.Merge(m, src)
}
func (m *NamespaceTemplate) XXX_Size() int {
	return xxx_messageInfo_NamespaceTemplate.Size(m)
}
func (m *NamespaceTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceTemplate proto.InternalMessageInfo

func (m *NamespaceTemplate) GetExportedParamName() []string {
	if m != nil {
		return m.ExportedParamName
	}
	return nil
}

func (m *NamespaceTemplate) GetNsResources() *NamespaceResources {
	if m != nil {
		return m.NsResources
	}
	return nil
}

type NamespaceResources struct {
	//Namespace is mandatory
	// +required
	Namespace *v1.Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	//Resources to be created and must include at least namespace
	// +optional
	Resources            []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NamespaceResources) Reset()         { *m = NamespaceResources{} }
func (m *NamespaceResources) String() string { return proto.CompactTextString(m) }
func (*NamespaceResources) ProtoMessage()    {}
func (*NamespaceResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f02f0592f6a6cd, []int{1}
}

func (m *NamespaceResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceResources.Unmarshal(m, b)
}
func (m *NamespaceResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceResources.Marshal(b, m, deterministic)
}
func (m *NamespaceResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceResources.Merge(m, src)
}
func (m *NamespaceResources) XXX_Size() int {
	return xxx_messageInfo_NamespaceResources.Size(m)
}
func (m *NamespaceResources) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceResources.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceResources proto.InternalMessageInfo

func (m *NamespaceResources) GetNamespace() *v1.Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *NamespaceResources) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type Resource struct {
	//ServiceAccount to be created for this namespace.
	// Must include type=ServiceAccount and only service account will be read at
	// the server side and everything else will be ignored
	// +optional
	ServiceAccount *v1.ServiceAccount `protobuf:"bytes,1,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	//Role to be created for this namespace.
	// Must include type=Role and only Role will be read at
	// the server side and everything else will be ignored
	// +optional
	Role *v11.Role `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	//RoleBinding to bind the role and service account for this namespace.
	// Must include type=RoleBinding and only RoleBinding will be read at
	// the server side and everything else will be ignored
	// +optional
	RoleBinding *v11.RoleBinding `protobuf:"bytes,3,opt,name=roleBinding,proto3" json:"roleBinding,omitempty"`
	//ResourceQuota to be created for this namespace.
	// Must include type=ResourceQuota and only ResourceQuota will be read at
	// the server side and everything else will be ignored.
	// +optional
	ResourceQuota *v1.ResourceQuota `protobuf:"bytes,4,opt,name=resourceQuota,proto3" json:"resourceQuota,omitempty"`
	//CustomResource to be created for this namespace
	//Must include type=CustomResource and only CustomResource will be read at
	// the server side and everything else will be ignored
	// +optional
	CustomResource *CustomResource `protobuf:"bytes,5,opt,name=customResource,proto3" json:"customResource,omitempty"`
	// +required
	Name string `protobuf:"bytes,15,opt,name=name,proto3" json:"name,omitempty"`
	//Type represents which k8s resource is being included in the resource entry
	//Allowed values are
	// - ServiceAccount
	// - Role
	// - RoleBinding
	// - ResourceQuota
	// - CustomResource
	// +kubebuilder:validation:Enum=ServiceAccount;Role;RoleBinding;ResourceQuota;CustomResource
	Type string `protobuf:"bytes,16,opt,name=type,proto3" json:"type,omitempty"`
	//dependsOn is an optional field and can be used to delay the creation until the referenced resource got created
	//dependsOn should provide the name of the resource it dependent on
	// +optional
	DependsOn string `protobuf:"bytes,17,opt,name=dependsOn,proto3" json:"dependsOn,omitempty"`
	//createOnly param can be used to control whether resource to be created only once and do not overwrite in subsequent reconcile process
	// +optional
	// +kubebuilder:validation:Enum="true";"false"
	CreateOnly           string   `protobuf:"bytes,18,opt,name=createOnly,proto3" json:"createOnly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f02f0592f6a6cd, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetServiceAccount() *v1.ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func (m *Resource) GetRole() *v11.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *Resource) GetRoleBinding() *v11.RoleBinding {
	if m != nil {
		return m.RoleBinding
	}
	return nil
}

func (m *Resource) GetResourceQuota() *v1.ResourceQuota {
	if m != nil {
		return m.ResourceQuota
	}
	return nil
}

func (m *Resource) GetCustomResource() *CustomResource {
	if m != nil {
		return m.CustomResource
	}
	return nil
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetDependsOn() string {
	if m != nil {
		return m.DependsOn
	}
	return ""
}

func (m *Resource) GetCreateOnly() string {
	if m != nil {
		return m.CreateOnly
	}
	return ""
}

type CustomResource struct {
	//GroupVersionKind should be used to provide the specific GVK for this custom resource
	GVK *GroupVersionKind `protobuf:"bytes,1,opt,name=GVK,proto3" json:"GVK,omitempty"`
	//manifest should be used to provide the custom resource manifest and must be in JSON
	Manifest             string   `protobuf:"bytes,2,opt,name=manifest,proto3" json:"manifest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomResource) Reset()         { *m = CustomResource{} }
func (m *CustomResource) String() string { return proto.CompactTextString(m) }
func (*CustomResource) ProtoMessage()    {}
func (*CustomResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f02f0592f6a6cd, []int{3}
}

func (m *CustomResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomResource.Unmarshal(m, b)
}
func (m *CustomResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomResource.Marshal(b, m, deterministic)
}
func (m *CustomResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomResource.Merge(m, src)
}
func (m *CustomResource) XXX_Size() int {
	return xxx_messageInfo_CustomResource.Size(m)
}
func (m *CustomResource) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomResource.DiscardUnknown(m)
}

var xxx_messageInfo_CustomResource proto.InternalMessageInfo

func (m *CustomResource) GetGVK() *GroupVersionKind {
	if m != nil {
		return m.GVK
	}
	return nil
}

func (m *CustomResource) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

//GroupVersionKind can be used to provide GVK of a custom resource
type GroupVersionKind struct {
	//group -custom resource group
	// +required
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	//version - custom resource version
	// +required
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	//kind - custom resource kind
	// +required
	Kind                 string   `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupVersionKind) Reset()         { *m = GroupVersionKind{} }
func (m *GroupVersionKind) String() string { return proto.CompactTextString(m) }
func (*GroupVersionKind) ProtoMessage()    {}
func (*GroupVersionKind) Descriptor() ([]byte, []int) {
	return fileDescriptor_32f02f0592f6a6cd, []int{4}
}

func (m *GroupVersionKind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupVersionKind.Unmarshal(m, b)
}
func (m *GroupVersionKind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupVersionKind.Marshal(b, m, deterministic)
}
func (m *GroupVersionKind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupVersionKind.Merge(m, src)
}
func (m *GroupVersionKind) XXX_Size() int {
	return xxx_messageInfo_GroupVersionKind.Size(m)
}
func (m *GroupVersionKind) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupVersionKind.DiscardUnknown(m)
}

var xxx_messageInfo_GroupVersionKind proto.InternalMessageInfo

func (m *GroupVersionKind) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *GroupVersionKind) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GroupVersionKind) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*NamespaceTemplate)(nil), "namespace.NamespaceTemplate")
	proto.RegisterType((*NamespaceResources)(nil), "namespace.NamespaceResources")
	proto.RegisterType((*Resource)(nil), "namespace.Resource")
	proto.RegisterType((*CustomResource)(nil), "namespace.CustomResource")
	proto.RegisterType((*GroupVersionKind)(nil), "namespace.GroupVersionKind")
}

func init() {
	proto.RegisterFile("pkg/grpc/proto/namespace/template.proto", fileDescriptor_32f02f0592f6a6cd)
}

var fileDescriptor_32f02f0592f6a6cd = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4f, 0x8f, 0xd3, 0x3c,
	0x10, 0xc6, 0x95, 0xb7, 0xbb, 0x2f, 0x9b, 0xa9, 0x28, 0x5b, 0xc3, 0xc1, 0x2c, 0xff, 0x4a, 0x2e,
	0xf4, 0xb0, 0x24, 0xea, 0x22, 0x04, 0x12, 0x07, 0xd4, 0xe5, 0x50, 0x89, 0x95, 0x58, 0x30, 0xa8,
	0x07, 0x38, 0xb9, 0xc9, 0x10, 0x4c, 0x1b, 0xdb, 0x72, 0x9c, 0x8a, 0xbd, 0x22, 0xbe, 0x14, 0xdf,
	0x0e, 0xd9, 0x4d, 0xda, 0xb4, 0xdb, 0x3d, 0xd5, 0x9e, 0xf9, 0x3d, 0x8f, 0x9f, 0xc9, 0x48, 0x85,
	0x67, 0x7a, 0x9e, 0x27, 0xb9, 0xd1, 0x69, 0xa2, 0x8d, 0xb2, 0x2a, 0x91, 0xbc, 0xc0, 0x52, 0xf3,
	0x14, 0x13, 0x8b, 0x85, 0x5e, 0x70, 0x8b, 0xb1, 0x6f, 0x90, 0x70, 0xdd, 0x39, 0x89, 0xe6, 0xaf,
	0xcb, 0x58, 0xa8, 0x84, 0x6b, 0x91, 0xa4, 0xca, 0x60, 0xb2, 0x1c, 0x25, 0x39, 0x4a, 0x34, 0xdc,
	0x62, 0xb6, 0xc2, 0xb7, 0x18, 0x33, 0xe3, 0xe9, 0x1e, 0x26, 0xfa, 0x1d, 0x40, 0xff, 0x43, 0xe3,
	0xfa, 0xa5, 0x7e, 0x8e, 0x9c, 0x42, 0x1f, 0x7f, 0x69, 0x65, 0x2c, 0x66, 0x1f, 0xb9, 0xe1, 0x85,
	0x23, 0x68, 0x30, 0xe8, 0x0c, 0x43, 0x76, 0xbd, 0x41, 0xde, 0x42, 0x57, 0x96, 0x0c, 0x4b, 0x55,
	0x99, 0x14, 0x4b, 0xfa, 0xdf, 0x20, 0x18, 0x76, 0xcf, 0x1e, 0xc5, 0xeb, 0xb0, 0xf1, 0xfa, 0x81,
	0x35, 0xc4, 0xda, 0x8a, 0xe8, 0x4f, 0x00, 0xe4, 0x3a, 0x43, 0xde, 0xc0, 0x66, 0x60, 0x1a, 0xd4,
	0xae, 0xab, 0x99, 0x62, 0xae, 0x45, 0xec, 0xe6, 0x8e, 0x97, 0xa3, 0x96, 0xfd, 0x86, 0x27, 0x23,
	0x08, 0x4d, 0x2b, 0x52, 0x67, 0xd8, 0x3d, 0xbb, 0xdb, 0x8a, 0xd4, 0xbc, 0xc2, 0x36, 0x54, 0xf4,
	0xb7, 0x03, 0x47, 0x4d, 0x9d, 0xbc, 0x87, 0x5e, 0x89, 0x66, 0x29, 0x52, 0x1c, 0xa7, 0xa9, 0xaa,
	0xa4, 0xad, 0x13, 0x44, 0xfb, 0x12, 0x7c, 0xde, 0x22, 0xd9, 0x8e, 0x92, 0x9c, 0xc2, 0x81, 0x51,
	0x0b, 0xac, 0xbf, 0x0c, 0x6d, 0x3b, 0xb8, 0xbd, 0x38, 0x07, 0xa6, 0x16, 0xc8, 0x3c, 0x45, 0xc6,
	0xd0, 0x75, 0xbf, 0xe7, 0x42, 0x66, 0x42, 0xe6, 0xb4, 0xe3, 0x45, 0x4f, 0x6e, 0x12, 0xd5, 0x18,
	0x6b, 0x6b, 0xc8, 0x04, 0x6e, 0x37, 0x63, 0x7d, 0xaa, 0x94, 0xe5, 0xf4, 0xc0, 0x9b, 0x3c, 0xdd,
	0x97, 0x9d, 0xb5, 0x41, 0xb6, 0xad, 0x23, 0x63, 0xe8, 0xa5, 0x55, 0x69, 0x55, 0xd1, 0x50, 0xf4,
	0xd0, 0x3b, 0xdd, 0x6f, 0x7d, 0xca, 0x77, 0x5b, 0x00, 0xdb, 0x11, 0x10, 0x02, 0x07, 0x8e, 0xa5,
	0x77, 0x06, 0xc1, 0x30, 0x64, 0xfe, 0xec, 0x6a, 0xf6, 0x4a, 0x23, 0x3d, 0x5e, 0xd5, 0xdc, 0x99,
	0x3c, 0x84, 0x30, 0x43, 0x8d, 0x32, 0x2b, 0x2f, 0x25, 0xed, 0xfb, 0xc6, 0xa6, 0x40, 0x1e, 0x03,
	0xa4, 0x06, 0xb9, 0xc5, 0x4b, 0xb9, 0xb8, 0xa2, 0xc4, 0xb7, 0x5b, 0x95, 0xe8, 0x1b, 0xf4, 0xb6,
	0x73, 0x90, 0xe7, 0xd0, 0x99, 0x4c, 0x2f, 0xea, 0xad, 0x3d, 0x68, 0xe5, 0x9d, 0x18, 0x55, 0xe9,
	0x29, 0x9a, 0x52, 0x28, 0x79, 0x21, 0x64, 0xc6, 0x1c, 0x47, 0x4e, 0xe0, 0xa8, 0xe0, 0x52, 0x7c,
	0xc7, 0xd2, 0xfa, 0x3d, 0x85, 0x6c, 0x7d, 0x8f, 0xa6, 0x70, 0xbc, 0x2b, 0x22, 0xf7, 0xe0, 0x30,
	0x77, 0x35, 0xff, 0x40, 0xc8, 0x56, 0x17, 0x42, 0xe1, 0xd6, 0x72, 0x05, 0xd5, 0x26, 0xcd, 0xd5,
	0x8d, 0x3c, 0x17, 0x32, 0xf3, 0xeb, 0x0c, 0x99, 0x3f, 0x9f, 0xbf, 0xfa, 0xfa, 0x32, 0x17, 0xf6,
	0x47, 0x35, 0x8b, 0x53, 0x55, 0x24, 0x73, 0x14, 0x73, 0xa5, 0x8d, 0xfa, 0x99, 0x14, 0x5c, 0xf2,
	0x1c, 0x4d, 0x72, 0xd3, 0xff, 0xc2, 0xec, 0x7f, 0x5f, 0x78, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x6b, 0x12, 0x53, 0xb7, 0x3a, 0x04, 0x00, 0x00,
}
