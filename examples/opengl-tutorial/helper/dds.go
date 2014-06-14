// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package helper

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/go-gl/gl"
)

type ddsHeader struct {
	Size, Flags   uint32
	Height, Width uint32 // @header[8],header[12]
	LinearSize    uint32 // @header[16]
	Depth         uint32
	MipMapCount   uint32 // @header[24]
	_             [11]uint32
	PixelFormat   struct {
		Size, Flags                            uint32
		FourCC                                 uint32
		RGBBitCount                            uint32
		RBitMask, GBitMask, BBitMask, ABitMask uint32
	}
	Caps, Caps2, Caps3, Caps4 uint32
	_                         uint32
}

const (
	fourCC_DXT1 uint32 = 0x31545844 // Equivalent to "DXT1" in ASCII
	fourCC_DXT3        = 0x33545844 // Equivalent to "DXT3" in ASCII
	fourCC_DXT5        = 0x35545844 // Equivalent to "DXT5" in ASCII
)

func TextureFromDDS(fname string) (gl.Texture, error) {
	file, err := os.Open(fname)
	if err != nil {
		return gl.Texture(0), fmt.Errorf("Cannot open DDS file: %v", err)
	}
	defer file.Close()

	var filecode [4]byte
	binary.Read(file, binary.LittleEndian, &filecode) // DDS is always Little Endian encoded
	if string(filecode[:]) != "DDS " {
		return gl.Texture(0), fmt.Errorf("File code is not DDS, instead got %v.", string(filecode[:]))
	}

	header := ddsHeader{}
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return gl.Texture(0), fmt.Errorf("Couldn't read DDS header: %v", err)
	}

	stat, err := file.Stat()
	if err != nil {
		return gl.Texture(0), fmt.Errorf("Couldn't get size of file: %v", err)
	}

	// All of the file after the header and file identifier "DDS "
	// is images. 4 byte identifier + 124 byte header = 128 bytes
	// to take off the file size.
	bufSize := stat.Size() - 128

	buffer := make([]byte, bufSize)
	err = binary.Read(file, binary.LittleEndian, buffer)
	if err != nil {
		return gl.Texture(0), fmt.Errorf("Couldn't read all mipmaps into buffer: %v", err)
	}

	//var components uint32
	var blockSize uint32
	var format gl.GLenum
	switch header.PixelFormat.FourCC {
	case fourCC_DXT1:
		format = gl.COMPRESSED_RGBA_S3TC_DXT1_EXT
		//components = 3
		blockSize = 8
	case fourCC_DXT3:
		format = gl.COMPRESSED_RGBA_S3TC_DXT3_EXT
		//components = 4
		blockSize = 16
	case fourCC_DXT5:
		format = gl.COMPRESSED_RGBA_S3TC_DXT5_EXT
		//components = 4
		blockSize = 16
	default:
		return gl.Texture(0), fmt.Errorf("Invalid four CC in DDS header. Got: %x; Expected %x; %x; or %x", header.PixelFormat.FourCC, fourCC_DXT1, fourCC_DXT3, fourCC_DXT5)
	}

	offset := 0

	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	for level, width, height := 0, header.Width, header.Height; level < int(header.MipMapCount) && (width > 0 || height > 0); level, width, height = level+1, width/2, height/2 {
		size := int(((width + 3) / 4) * ((height + 3) / 4) * blockSize)
		gl.CompressedTexImage2D(gl.TEXTURE_2D, level, format, int(width), int(height), 0, size, &buffer[offset])

		offset += size
	}

	return tex, nil
}
