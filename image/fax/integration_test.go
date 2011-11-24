// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fax

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"testing"
)

const (
	inputFile = "testdata/red.tiff"
	outputFile = "testdata/red.png"
	width = 501
	height = 713
)

func blob() []byte {
	data, error := ioutil.ReadFile(inputFile)
	if error != nil {
		panic(error)
	}
	return data[8:23800] // strip TIFF
}

func TestFullRun(t *testing.T) {
	b := bytes.NewBuffer(blob())
	result, err := DecodeG4(b, width, height)
	if err != nil {
		t.Fatal(err)
	}

	b.Reset()
	err = png.Encode(b, result)
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(outputFile, b.Bytes(), 0666)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkFullRun(b *testing.B) {
        b.StopTimer()
	data := blob()
        b.SetBytes(int64(len(data)))
        for rounds := b.N; rounds != 0; rounds-- {
		source := bytes.NewBuffer(data)
		b.StartTimer()
		_, err := DecodeG4(source, width, height)
		b.StopTimer()
		if err != nil {
			panic(err)
		}
        }
}
