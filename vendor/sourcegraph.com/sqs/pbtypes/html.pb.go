// Code generated by protoc-gen-gogo.
// source: html.proto
// DO NOT EDIT!

package pbtypes

import proto "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// HTML is a type which marshals into {__html: "html code here"} to designate
// that this value is sanitized HTML code, see
// https://facebook.github.io/react/tips/dangerously-set-inner-html.html
type HTML struct {
	HTML string `protobuf:"bytes,1,opt,proto3" json:"__html"`
}

func (m *HTML) Reset()         { *m = HTML{} }
func (m *HTML) String() string { return proto.CompactTextString(m) }
func (*HTML) ProtoMessage()    {}

func init() {
}