// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package bele

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/q191201771/naza/pkg/assert"
)

func TestBEUint16(t *testing.T) {
	vector := []struct {
		input  []byte
		output uint16
	}{
		{input: []byte{0, 0}, output: 0},
		{input: []byte{0, 1}, output: 1},
		{input: []byte{0, 255}, output: 255},
		{input: []byte{1, 0}, output: 256},
		{input: []byte{255, 0}, output: 255 * 256},
		{input: []byte{12, 34}, output: 12*256 + 34},
	}

	for i := 0; i < len(vector); i++ {
		assert.Equal(t, vector[i].output, BEUint16(vector[i].input))
	}
}

func TestBEUint24(t *testing.T) {
	vector := []struct {
		input  []byte
		output uint32
	}{
		{input: []byte{0, 0, 0}, output: 0},
		{input: []byte{0, 0, 1}, output: 1},
		{input: []byte{0, 1, 0}, output: 256},
		{input: []byte{1, 0, 0}, output: 1 * 256 * 256},
		{input: []byte{12, 34, 56}, output: 12*256*256 + 34*256 + 56},
	}

	for i := 0; i < len(vector); i++ {
		assert.Equal(t, vector[i].output, BEUint24(vector[i].input))
	}
}

func TestBEUint32(t *testing.T) {
	vector := []struct {
		input  []byte
		output uint32
	}{
		{input: []byte{0, 0, 0, 0}, output: 0},
		{input: []byte{0, 0, 1, 0}, output: 1 * 256},
		{input: []byte{0, 1, 0, 0}, output: 1 * 256 * 256},
		{input: []byte{1, 0, 0, 0}, output: 1 * 256 * 256 * 256},
		{input: []byte{12, 34, 56, 78}, output: 12*256*256*256 + 34*256*256 + 56*256 + 78},
	}

	for i := 0; i < len(vector); i++ {
		assert.Equal(t, vector[i].output, BEUint32(vector[i].input))
	}
}

func TestBEFloat64(t *testing.T) {
	vector := []int{
		1,
		0xFF,
		0xFFFF,
		0xFFFFFF,
	}
	for i := 0; i < len(vector); i++ {
		b := &bytes.Buffer{}
		err := binary.Write(b, binary.BigEndian, float64(vector[i]))
		assert.Equal(t, nil, err)
		assert.Equal(t, vector[i], int(BEFloat64(b.Bytes())))
	}
}

func TestLEUint32(t *testing.T) {
	vector := []struct {
		input  []byte
		output uint32
	}{
		{input: []byte{0, 0, 0, 0}, output: 0},
		{input: []byte{0, 0, 1, 0}, output: 1 * 256 * 256},
		{input: []byte{0, 1, 0, 0}, output: 1 * 256},
		{input: []byte{1, 0, 0, 0}, output: 1},
		{input: []byte{12, 34, 56, 78}, output: 12 + 34*256 + 56*256*256 + 78*256*256*256},
	}

	for i := 0; i < len(vector); i++ {
		assert.Equal(t, vector[i].output, LEUint32(vector[i].input))
	}
}

func TestBEPutUint24(t *testing.T) {
	vector := []struct {
		input  uint32
		output []byte
	}{
		{input: 0, output: []byte{0, 0, 0}},
		{input: 1, output: []byte{0, 0, 1}},
		{input: 256, output: []byte{0, 1, 0}},
		{input: 1 * 256 * 256, output: []byte{1, 0, 0}},
		{input: 12*256*256 + 34*256 + 56, output: []byte{12, 34, 56}},
	}

	out := make([]byte, 3)
	for i := 0; i < len(vector); i++ {
		BEPutUint24(out, vector[i].input)
		assert.Equal(t, vector[i].output, out)
	}
}

func TestBEPutUint32(t *testing.T) {
	vector := []struct {
		input  uint32
		output []byte
	}{
		{input: 0, output: []byte{0, 0, 0, 0}},
		{input: 1 * 256, output: []byte{0, 0, 1, 0}},
		{input: 1 * 256 * 256, output: []byte{0, 1, 0, 0}},
		{input: 1 * 256 * 256 * 256, output: []byte{1, 0, 0, 0}},
		{input: 12*256*256*256 + 34*256*256 + 56*256 + 78, output: []byte{12, 34, 56, 78}},
	}

	out := make([]byte, 4)
	for i := 0; i < len(vector); i++ {
		BEPutUint32(out, vector[i].input)
		assert.Equal(t, vector[i].output, out)
	}
}

func TestLEPutUint32(t *testing.T) {
	vector := []struct {
		input  uint32
		output []byte
	}{
		{input: 0, output: []byte{0, 0, 0, 0}},
		{input: 1 * 256 * 256, output: []byte{0, 0, 1, 0}},
		{input: 1 * 256, output: []byte{0, 1, 0, 0}},
		{input: 1, output: []byte{1, 0, 0, 0}},
		{input: 78*256*256*256 + 56*256*256 + 34*256 + 12, output: []byte{12, 34, 56, 78}},
	}

	out := make([]byte, 4)
	for i := 0; i < len(vector); i++ {
		LEPutUint32(out, vector[i].input)
		assert.Equal(t, vector[i].output, out)
	}
}

func TestWriteBEUint24(t *testing.T) {
	vector := []struct {
		input  uint32
		output []byte
	}{
		{input: 0, output: []byte{0, 0, 0}},
		{input: 1, output: []byte{0, 0, 1}},
		{input: 256, output: []byte{0, 1, 0}},
		{input: 1 * 256 * 256, output: []byte{1, 0, 0}},
		{input: 12*256*256 + 34*256 + 56, output: []byte{12, 34, 56}},
	}

	for i := 0; i < len(vector); i++ {
		out := &bytes.Buffer{}
		err := WriteBEUint24(out, vector[i].input)
		assert.Equal(t, nil, err)
		assert.Equal(t, vector[i].output, out.Bytes())
	}
}

func TestWriteBE(t *testing.T) {
	vector := []struct {
		input  interface{}
		output []byte
	}{
		{input: uint32(1), output: []byte{0, 0, 0, 1}},
		{input: uint64(1), output: []byte{0, 0, 0, 0, 0, 0, 0, 1}},
	}
	for i := 0; i < len(vector); i++ {
		out := &bytes.Buffer{}
		err := WriteBE(out, vector[i].input)
		assert.Equal(t, nil, err)
		assert.Equal(t, vector[i].output, out.Bytes())
	}
}

func TestWriteLE(t *testing.T) {
	vector := []struct {
		input  interface{}
		output []byte
	}{
		{input: uint32(1), output: []byte{1, 0, 0, 0}},
		{input: uint64(1), output: []byte{1, 0, 0, 0, 0, 0, 0, 0}},
	}
	for i := 0; i < len(vector); i++ {
		out := &bytes.Buffer{}
		err := WriteLE(out, vector[i].input)
		assert.Equal(t, nil, err)
		assert.Equal(t, vector[i].output, out.Bytes())
	}
}

func BenchmarkBEFloat64(b *testing.B) {
	buf := &bytes.Buffer{}
	_ = binary.Write(buf, binary.BigEndian, float64(123.4))
	for i := 0; i < b.N; i++ {
		BEFloat64(buf.Bytes())
	}
}

func BenchmarkBEPutUint24(b *testing.B) {
	out := make([]byte, 3)
	for i := 0; i < b.N; i++ {
		BEPutUint24(out, uint32(i))
	}
}

func BenchmarkBEUint24(b *testing.B) {
	buf := []byte{1, 2, 3}
	for i := 0; i < b.N; i++ {
		BEUint24(buf)
	}
}

func BenchmarkWriteBE(b *testing.B) {
	out := &bytes.Buffer{}
	in := uint64(123)
	for i := 0; i < b.N; i++ {
		_ = WriteBE(out, in)
	}
}
