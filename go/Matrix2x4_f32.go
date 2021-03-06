// automatically generated by the FlatBuffers compiler, do not modify

package flatmath

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// 2x4 matrix of floats
type Matrix2x4_f32 struct {
	_tab flatbuffers.Struct
}

func (rcv *Matrix2x4_f32) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Matrix2x4_f32) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Matrix2x4_f32) _0(obj *Vector4_f32) *Vector4_f32 {
	if obj == nil {
		obj = new(Vector4_f32)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+0)
	return obj
}
func (rcv *Matrix2x4_f32) _1(obj *Vector4_f32) *Vector4_f32 {
	if obj == nil {
		obj = new(Vector4_f32)
	}
	obj.Init(rcv._tab.Bytes, rcv._tab.Pos+16)
	return obj
}

func CreateMatrix2x4_f32(builder *flatbuffers.Builder, _0_x float32, _0_y float32, _0_z float32, _0_w float32, _1_x float32, _1_y float32, _1_z float32, _1_w float32) flatbuffers.UOffsetT {
	builder.Prep(4, 32)
	builder.Prep(4, 16)
	builder.PrependFloat32(_1_w)
	builder.PrependFloat32(_1_z)
	builder.PrependFloat32(_1_y)
	builder.PrependFloat32(_1_x)
	builder.Prep(4, 16)
	builder.PrependFloat32(_0_w)
	builder.PrependFloat32(_0_z)
	builder.PrependFloat32(_0_y)
	builder.PrependFloat32(_0_x)
	return builder.Offset()
}
