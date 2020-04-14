// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/grpc/proto/cluster/cluster.proto

// +kubebuilder:object:generate=true

package cluster

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Cluster struct {
	//Name contains cluster name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//Type contains kubernetes cluster installation type. ex: AWS, GCP
	Cloud string `protobuf:"bytes,2,opt,name=cloud,proto3" json:"cloud,omitempty"`
	//Config contains info to connect to the target cluster
	//This is same as config struct in https://github.com/kubernetes/client-go/blob/master/rest/config.go
	// but have to define it again here with whatever we need
	Config               *Config  `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3be2089bb94235a, []int{0}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *Cluster) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

// Config holds the common attributes that can be passed to a Kubernetes client on
// initialization.
// +optional
type Config struct {
	// Host must be a host string, a host:port pair, or a URL to the base of the apiserver.
	// If a URL is given then the (optional) Path of that URL represents a prefix that must
	// be appended to all request URIs used to access the apiserver. This allows a frontend
	// proxy to easily relocate all of the apiserver endpoints.
	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// password contains basic auth password
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Server requires Bearer authentication. This client will not attempt to use
	// refresh tokens for an OAuth2 flow.
	// TODO: demonstrate an OAuth2 compatible client.
	BearerToken string `protobuf:"bytes,4,opt,name=bearerToken,proto3" json:"bearerToken,omitempty"`
	// Secret containing a BearerToken.
	// If set, The last successfully read value takes precedence over BearerToken.
	// +optional
	BearerTokenSecret string `protobuf:"bytes,5,opt,name=bearerTokenSecret,proto3" json:"bearerTokenSecret,omitempty"`
	// TLSClientConfig contains settings to enable transport layer security
	// +optional
	TlsClientConfig      *TLSClientConfig `protobuf:"bytes,6,opt,name=tlsClientConfig,proto3" json:"tlsClientConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3be2089bb94235a, []int{1}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Config) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Config) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Config) GetBearerToken() string {
	if m != nil {
		return m.BearerToken
	}
	return ""
}

func (m *Config) GetBearerTokenSecret() string {
	if m != nil {
		return m.BearerTokenSecret
	}
	return ""
}

func (m *Config) GetTlsClientConfig() *TLSClientConfig {
	if m != nil {
		return m.TlsClientConfig
	}
	return nil
}

// TLSClientConfig contains settings to enable transport layer security
type TLSClientConfig struct {
	// Server should be accessed without verifying the TLS certificate. For testing only.
	InSecure bool `protobuf:"varint,1,opt,name=inSecure,proto3" json:"inSecure,omitempty"`
	// ServerName is passed to the server for SNI and is used in the client to check server
	// ceritificates against. If ServerName is empty, the hostname used to contact the
	// server is used.
	ServerName string `protobuf:"bytes,2,opt,name=serverName,proto3" json:"serverName,omitempty"`
	// CertData holds PEM-encoded bytes (typically read from a client certificate file).
	// CertData takes precedence over CertFile
	CertData []byte `protobuf:"bytes,3,opt,name=certData,proto3" json:"certData,omitempty"`
	// KeyData holds PEM-encoded bytes (typically read from a client certificate key file).
	// KeyData takes precedence over KeyFile
	KeyData []byte `protobuf:"bytes,4,opt,name=keyData,proto3" json:"keyData,omitempty"`
	// CAData holds PEM-encoded bytes (typically read from a root certificates bundle).
	// CAData takes precedence over CAFile
	CaData []byte `protobuf:"bytes,5,opt,name=caData,proto3" json:"caData,omitempty"`
	// NextProtos is a list of supported application level protocols, in order of preference.
	// Used to populate tls.Config.NextProtos.
	// To indicate to the server http/1.1 is preferred over http/2, set to ["http/1.1", "h2"] (though the server is free to ignore that preference).
	// To use only http/1.1, set to ["http/1.1"].
	NextProtos           []string `protobuf:"bytes,6,rep,name=nextProtos,proto3" json:"nextProtos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TLSClientConfig) Reset()         { *m = TLSClientConfig{} }
func (m *TLSClientConfig) String() string { return proto.CompactTextString(m) }
func (*TLSClientConfig) ProtoMessage()    {}
func (*TLSClientConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3be2089bb94235a, []int{2}
}

func (m *TLSClientConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSClientConfig.Unmarshal(m, b)
}
func (m *TLSClientConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSClientConfig.Marshal(b, m, deterministic)
}
func (m *TLSClientConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSClientConfig.Merge(m, src)
}
func (m *TLSClientConfig) XXX_Size() int {
	return xxx_messageInfo_TLSClientConfig.Size(m)
}
func (m *TLSClientConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSClientConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TLSClientConfig proto.InternalMessageInfo

func (m *TLSClientConfig) GetInSecure() bool {
	if m != nil {
		return m.InSecure
	}
	return false
}

func (m *TLSClientConfig) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *TLSClientConfig) GetCertData() []byte {
	if m != nil {
		return m.CertData
	}
	return nil
}

func (m *TLSClientConfig) GetKeyData() []byte {
	if m != nil {
		return m.KeyData
	}
	return nil
}

func (m *TLSClientConfig) GetCaData() []byte {
	if m != nil {
		return m.CaData
	}
	return nil
}

func (m *TLSClientConfig) GetNextProtos() []string {
	if m != nil {
		return m.NextProtos
	}
	return nil
}

func init() {
	proto.RegisterType((*Cluster)(nil), "cluster.Cluster")
	proto.RegisterType((*Config)(nil), "cluster.Config")
	proto.RegisterType((*TLSClientConfig)(nil), "cluster.TLSClientConfig")
}

func init() {
	proto.RegisterFile("pkg/grpc/proto/cluster/cluster.proto", fileDescriptor_a3be2089bb94235a)
}

var fileDescriptor_a3be2089bb94235a = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0xc9, 0x55, 0xa3, 0x1e, 0x2f, 0xc8, 0x1d, 0x2e, 0x97, 0xe1, 0x2e, 0x4a, 0x90, 0x42,
	0x5d, 0x14, 0x03, 0x95, 0xbe, 0x80, 0x76, 0x59, 0x4a, 0x89, 0xae, 0x4a, 0x37, 0xe3, 0x78, 0x1a,
	0xd3, 0xc4, 0x4c, 0x98, 0x99, 0xf4, 0xcf, 0xb3, 0xf5, 0x85, 0xfa, 0x18, 0x65, 0x4e, 0x4c, 0x0c,
	0xb6, 0x2b, 0xcf, 0xef, 0xfb, 0x66, 0xce, 0xe7, 0x97, 0x04, 0xce, 0x8b, 0x34, 0x0e, 0x63, 0x5d,
	0xc8, 0xb0, 0xd0, 0xca, 0xaa, 0x50, 0x66, 0xa5, 0xb1, 0xa8, 0xeb, 0xdf, 0x19, 0xa9, 0xac, 0x7f,
	0xc0, 0xc9, 0x23, 0xf4, 0x97, 0xd5, 0xc8, 0x18, 0x74, 0x73, 0xb1, 0x47, 0xee, 0x05, 0xde, 0x74,
	0x18, 0xd1, 0xcc, 0xfe, 0x42, 0x4f, 0x66, 0xaa, 0xdc, 0xf2, 0x5f, 0x24, 0x56, 0xc0, 0x2e, 0xc0,
	0x97, 0x2a, 0x7f, 0x4a, 0x62, 0xde, 0x09, 0xbc, 0xe9, 0xe8, 0x6a, 0x3c, 0xab, 0xb7, 0x2f, 0x49,
	0x8e, 0x0e, 0xf6, 0xe4, 0xd3, 0x03, 0xbf, 0x92, 0xdc, 0xf6, 0x9d, 0x32, 0xb6, 0xde, 0xee, 0x66,
	0xf6, 0x1f, 0x06, 0xa5, 0x41, 0x4d, 0xa9, 0x55, 0x40, 0xc3, 0xce, 0x2b, 0x84, 0x31, 0xaf, 0x4a,
	0x6f, 0x29, 0x65, 0x18, 0x35, 0xcc, 0x02, 0x18, 0x6d, 0x50, 0x68, 0xd4, 0x6b, 0x95, 0x62, 0xce,
	0xbb, 0x64, 0xb7, 0x25, 0x76, 0x09, 0x7f, 0x5a, 0xb8, 0x42, 0xa9, 0xd1, 0xf2, 0x1e, 0x9d, 0xfb,
	0x6e, 0xb0, 0x05, 0x8c, 0x6d, 0x66, 0x96, 0x59, 0x82, 0xb9, 0xad, 0xfe, 0x2e, 0xf7, 0xa9, 0x18,
	0x6f, 0x8a, 0xad, 0x6f, 0x57, 0x6d, 0x3f, 0x3a, 0xbd, 0x30, 0xf9, 0xf0, 0x60, 0x7c, 0x72, 0xc8,
	0x75, 0x48, 0x5c, 0x46, 0xa9, 0xab, 0xa7, 0x3a, 0x88, 0x1a, 0x66, 0x67, 0x00, 0x06, 0xf5, 0x0b,
	0xea, 0xbb, 0x63, 0xfb, 0x96, 0xe2, 0xee, 0x4a, 0xd4, 0xf6, 0x46, 0x58, 0x41, 0xfd, 0x7f, 0x47,
	0x0d, 0x33, 0x0e, 0xfd, 0x14, 0xdf, 0xc9, 0xea, 0x92, 0x55, 0x23, 0xfb, 0x07, 0xbe, 0x14, 0x64,
	0xf4, 0xc8, 0x38, 0x90, 0x4b, 0xcb, 0xf1, 0xcd, 0xde, 0xbb, 0x97, 0x6f, 0xb8, 0x1f, 0x74, 0x5c,
	0xda, 0x51, 0x59, 0x5c, 0x3f, 0xcc, 0xe3, 0xc4, 0xee, 0xca, 0xcd, 0x4c, 0xaa, 0x7d, 0x98, 0x62,
	0x92, 0xaa, 0x42, 0xab, 0xe7, 0x70, 0x2f, 0x72, 0x11, 0xa3, 0x0e, 0x7f, 0xfe, 0xa8, 0x36, 0x3e,
	0xe1, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x10, 0x38, 0x6c, 0xe4, 0x75, 0x02, 0x00, 0x00,
}
