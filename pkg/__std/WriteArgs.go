// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package __std

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type WriteArgs struct {
	_tab flatbuffers.Table
}

func GetRootAsWriteArgs(buf []byte, offset flatbuffers.UOffsetT) *WriteArgs {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WriteArgs{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *WriteArgs) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WriteArgs) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WriteArgs) Path() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WriteArgs) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *WriteArgs) Format() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *WriteArgs) MutateFormat(n int8) bool {
	return rcv._tab.MutateInt8Slot(8, n)
}

func (rcv *WriteArgs) Indent() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *WriteArgs) MutateIndent(n int8) bool {
	return rcv._tab.MutateInt8Slot(10, n)
}

func WriteArgsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func WriteArgsAddPath(builder *flatbuffers.Builder, path flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(path), 0)
}
func WriteArgsAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func WriteArgsAddFormat(builder *flatbuffers.Builder, format int8) {
	builder.PrependInt8Slot(2, format, 0)
}
func WriteArgsAddIndent(builder *flatbuffers.Builder, indent int8) {
	builder.PrependInt8Slot(3, indent, 0)
}
func WriteArgsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
