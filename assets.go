package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func templates_layout_gold() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x5c, 0x91,
		0x4d, 0x6e, 0xdc, 0x30, 0x0c, 0x85, 0xf7, 0x73, 0x0a, 0xc2, 0xeb, 0xd8,
		0x6a, 0xd1, 0x2e, 0xba, 0xd1, 0x3a, 0x17, 0xc8, 0x05, 0x38, 0xd2, 0xb3,
		0xad, 0x8e, 0x24, 0x1a, 0x12, 0xe7, 0x0f, 0x41, 0xee, 0x5e, 0xcb, 0x4e,
		0x13, 0xcc, 0x6c, 0x24, 0xfa, 0xc1, 0xfc, 0xc8, 0xf7, 0xe4, 0xc5, 0xe9,
		0x7d, 0x01, 0xcd, 0x9a, 0xe2, 0xa1, 0x1d, 0x14, 0x39, 0x4f, 0x16, 0xf9,
		0x40, 0x34, 0x83, 0xfd, 0x7a, 0x11, 0x25, 0x28, 0x93, 0x9b, 0xb9, 0x54,
		0xa8, 0x3d, 0xeb, 0xd8, 0xff, 0xf9, 0x96, 0x33, 0x27, 0xd8, 0x4b, 0xc0,
		0x75, 0x91, 0xa2, 0xe4, 0x24, 0x2b, 0xb2, 0xda, 0x6b, 0xf0, 0x3a, 0x5b,
		0x8f, 0x4b, 0x70, 0xe8, 0xb7, 0x8f, 0x97, 0x73, 0x45, 0xe9, 0xab, 0xe3,
		0xc8, 0xc7, 0x08, 0x9b, 0xe5, 0x25, 0xf1, 0x2d, 0xa4, 0x73, 0xda, 0x34,
		0xd8, 0x9f, 0x4f, 0x48, 0x8f, 0xea, 0x4a, 0x58, 0x34, 0x48, 0xfe, 0xa2,
		0x76, 0xaf, 0x12, 0x3d, 0xf5, 0xf4, 0x86, 0xb4, 0x44, 0x56, 0x10, 0xf2,
		0x14, 0x32, 0x68, 0x94, 0x42, 0xaf, 0xd2, 0x3d, 0x11, 0x4e, 0xb8, 0x5f,
		0xa5, 0xf8, 0xfa, 0xdd, 0x3e, 0x09, 0x4d, 0xd2, 0xfc, 0xb5, 0xcb, 0x93,
		0x3e, 0x61, 0x36, 0xfb, 0x35, 0x86, 0x44, 0x7f, 0xd9, 0x63, 0xc7, 0x69,
		0xd0, 0x08, 0x7a, 0x7f, 0x1f, 0xde, 0x5a, 0xf1, 0xf1, 0xb1, 0x89, 0x31,
		0xe4, 0x13, 0x15, 0x44, 0xdb, 0xd5, 0x79, 0x75, 0xed, 0xce, 0x4a, 0x61,
		0x1d, 0xd2, 0x51, 0x8b, 0xd2, 0x86, 0xc4, 0x13, 0xcc, 0xad, 0x6f, 0x12,
		0xcd, 0x05, 0xa3, 0xed, 0x42, 0x9a, 0xcc, 0xc8, 0x97, 0xa6, 0x0c, 0xeb,
		0xd1, 0x3d, 0x52, 0xaa, 0xde, 0x23, 0xea, 0x0c, 0xe8, 0x0e, 0x50, 0xdc,
		0xd4, 0xb8, 0x5a, 0xf7, 0xe6, 0xb5, 0x30, 0xdb, 0x1f, 0x43, 0x93, 0x12,
		0x7c, 0x60, 0xbb, 0x46, 0x83, 0xed, 0x89, 0x3e, 0x29, 0xfb, 0x14, 0x63,
		0x32, 0xd4, 0x67, 0x1e, 0x8e, 0x22, 0x5a, 0xb5, 0xf0, 0xe2, 0x7c, 0x1e,
		0x9c, 0x24, 0x33, 0xae, 0x09, 0xf4, 0x7c, 0x45, 0x95, 0x04, 0xf3, 0x7b,
		0xf8, 0x31, 0xfc, 0x6a, 0xfc, 0x07, 0xb9, 0xd1, 0xbb, 0x4f, 0x53, 0x5f,
		0xfb, 0xb4, 0x45, 0x8f, 0xe2, 0xef, 0xdb, 0xa4, 0x63, 0x14, 0x77, 0xfa,
		0x1f, 0xe6, 0xe1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xca, 0x95, 0x82,
		0x39, 0x39, 0x02, 0x00, 0x00,
		},
		"templates/layout.gold",
	)
}

func templates_top_gold() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x54,
		0x4d, 0x73, 0xda, 0x30, 0x10, 0xbd, 0xf3, 0x2b, 0x5e, 0xe9, 0x21, 0xed,
		0x01, 0x33, 0xb9, 0x26, 0xc3, 0xb5, 0xed, 0xa1, 0xbd, 0x31, 0xd3, 0xe9,
		0x71, 0xb1, 0xd7, 0xb6, 0x1a, 0x59, 0xd2, 0x48, 0x0b, 0x09, 0xe3, 0xf0,
		0xdf, 0xbb, 0xc6, 0x86, 0x40, 0x4a, 0x06, 0xc3, 0x30, 0x63, 0x69, 0xf5,
		0xde, 0xee, 0xdb, 0x0f, 0x89, 0x5f, 0x84, 0x5d, 0x91, 0x60, 0x69, 0xeb,
		0xd7, 0x32, 0x99, 0xac, 0xac, 0xcf, 0x9f, 0x90, 0x7b, 0xa7, 0x66, 0x99,
		0x00, 0x35, 0x53, 0xc1, 0x31, 0x13, 0x23, 0x96, 0x75, 0x0b, 0x64, 0xbd,
		0x65, 0x66, 0x9c, 0x42, 0xe2, 0xde, 0xa4, 0xa8, 0x7b, 0x7c, 0xf7, 0xb6,
		0x18, 0x76, 0x01, 0x4b, 0x6e, 0x82, 0x25, 0x61, 0xb0, 0xab, 0x8c, 0x63,
		0x94, 0x3e, 0x2a, 0x40, 0x8f, 0x13, 0xe7, 0x62, 0xbc, 0xcb, 0x12, 0x29,
		0x60, 0x70, 0xd8, 0xaf, 0x3b, 0x87, 0x47, 0x7f, 0x21, 0xf2, 0xb0, 0x82,
		0x6a, 0x29, 0xde, 0x36, 0xea, 0x21, 0x90, 0xcb, 0x84, 0x2a, 0x14, 0x3e,
		0x97, 0x6d, 0x38, 0x3d, 0x7a, 0x45, 0xdb, 0x4e, 0x51, 0x4b, 0x63, 0xa7,
		0xbb, 0xdd, 0x89, 0x7d, 0x15, 0x2f, 0xf1, 0x3b, 0xdc, 0x7f, 0x64, 0x4b,
		0xae, 0x5a, 0xb0, 0x1b, 0xc3, 0xef, 0xe0, 0xfb, 0xf2, 0x8c, 0x06, 0x03,
		0xfb, 0x32, 0x9e, 0xe3, 0xfb, 0xb8, 0x6d, 0x9b, 0x2d, 0xbb, 0xb3, 0xdd,
		0x6e, 0xb4, 0xb7, 0x95, 0x2f, 0xb6, 0x37, 0x84, 0xae, 0xef, 0x2f, 0xc5,
		0xed, 0xba, 0x86, 0xd9, 0x07, 0xfd, 0xba, 0xe6, 0x9d, 0x2d, 0x37, 0x33,
		0x53, 0x1c, 0x22, 0x7c, 0xee, 0xa6, 0x86, 0x94, 0x1e, 0xcf, 0x89, 0x6f,
		0xd8, 0xdc, 0x52, 0x4a, 0x1d, 0x3c, 0x7b, 0x8e, 0x14, 0xc2, 0x7b, 0xe0,
		0x59, 0x84, 0xd7, 0xe1, 0xab, 0xe8, 0xb6, 0x35, 0x25, 0x24, 0xae, 0x6f,
		0xa9, 0x4e, 0xff, 0x0b, 0x78, 0x9f, 0xf4, 0x1f, 0xbf, 0x46, 0x4e, 0x0e,
		0xeb, 0xc4, 0xd0, 0x0f, 0xbf, 0xe8, 0xa8, 0xa5, 0xa4, 0x23, 0x09, 0x5f,
		0x6a, 0xca, 0x10, 0xbd, 0x10, 0x73, 0x39, 0x94, 0x23, 0x50, 0xfe, 0x44,
		0x15, 0x67, 0xa3, 0x44, 0xea, 0x3d, 0xba, 0x51, 0x60, 0xc8, 0x30, 0xa2,
		0x02, 0xe8, 0xbb, 0x64, 0x12, 0x08, 0x72, 0xb1, 0x51, 0x57, 0xf5, 0x01,
		0xcb, 0x5a, 0xf9, 0xc9, 0x28, 0xdb, 0x94, 0x86, 0x13, 0x7e, 0x2c, 0x7f,
		0xfd, 0xec, 0xae, 0x96, 0x71, 0x15, 0x8c, 0xc3, 0x33, 0xaf, 0xa0, 0x1d,
		0x41, 0xc1, 0x1b, 0xb6, 0x3e, 0x34, 0x7a, 0xf9, 0x3f, 0x74, 0x7a, 0xc8,
		0x63, 0x96, 0xf2, 0x68, 0x82, 0x1c, 0xd2, 0xf9, 0x4b, 0x1b, 0xea, 0x2d,
		0x0f, 0x23, 0x92, 0xda, 0x50, 0x44, 0x93, 0x2a, 0x2c, 0x70, 0xf7, 0x9b,
		0x6d, 0xee, 0x1b, 0x86, 0xf8, 0x7d, 0xa2, 0x9f, 0xee, 0x1e, 0xaf, 0x91,
		0xc9, 0x72, 0x94, 0x2f, 0x4a, 0xff, 0xfa, 0x78, 0xf2, 0xa6, 0x54, 0x46,
		0xea, 0xf5, 0xaa, 0x7f, 0x53, 0xfa, 0xf5, 0xf9, 0x9b, 0x82, 0x6f, 0x5a,
		0xac, 0xc6, 0x47, 0xd6, 0x2c, 0x75, 0x4e, 0x6d, 0x1a, 0x0e, 0x8e, 0x60,
		0x9d, 0xdf, 0x63, 0x30, 0x42, 0x1d, 0xb9, 0x5c, 0x4c, 0x6b, 0x91, 0x90,
		0x1e, 0xe6, 0xf3, 0x1e, 0x93, 0xa9, 0xd0, 0xf9, 0xd6, 0x27, 0x9d, 0x98,
		0x79, 0xa5, 0x5a, 0xa7, 0x27, 0xe2, 0x4c, 0x56, 0x92, 0xfe, 0x67, 0x83,
		0x8c, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xca, 0xc8, 0xbf, 0x59,
		0x05, 0x00, 0x00,
		},
		"templates/top.gold",
	)
}

func public_css_style_css() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x55,
		0x6b, 0x6e, 0xeb, 0x28, 0x14, 0xfe, 0xcf, 0x2a, 0x90, 0x46, 0x55, 0x67,
		0x24, 0x63, 0xc5, 0x69, 0xd3, 0xa4, 0xee, 0x9f, 0xae, 0x61, 0x76, 0x80,
		0x01, 0xdb, 0xa8, 0x18, 0x2c, 0x7c, 0xdc, 0xd7, 0x28, 0x7b, 0x9f, 0x83,
		0x4d, 0xd2, 0xda, 0x49, 0xea, 0xdc, 0x2b, 0x5d, 0x45, 0x08, 0x03, 0xdf,
		0x79, 0x7d, 0xe7, 0x91, 0x67, 0x51, 0x73, 0xdf, 0x29, 0xa0, 0xb7, 0x3d,
		0x94, 0x6c, 0x77, 0xfb, 0x44, 0x0a, 0x27, 0x3f, 0x12, 0x52, 0x67, 0xb8,
		0xd6, 0xb8, 0xee, 0x70, 0xdd, 0xe3, 0xda, 0xe0, 0x7a, 0x48, 0x08, 0x4f,
		0x88, 0x6e, 0xaa, 0x84, 0x94, 0xce, 0x37, 0xf8, 0x69, 0xdb, 0x1e, 0xf0,
		0x41, 0x71, 0xa9, 0x7c, 0xb8, 0x74, 0x10, 0x76, 0x69, 0x70, 0xe1, 0xbd,
		0x94, 0x09, 0x31, 0x3a, 0x21, 0x3d, 0x9e, 0x3b, 0x65, 0x94, 0x80, 0xb0,
		0x0b, 0xd0, 0xce, 0xe2, 0xa3, 0x7e, 0x4d, 0xc8, 0xab, 0x96, 0xca, 0x25,
		0xa4, 0x4d, 0x48, 0xd1, 0x03, 0x84, 0x6b, 0x50, 0xef, 0xc0, 0xbd, 0x42,
		0x3b, 0xad, 0x57, 0x09, 0x11, 0x4e, 0x2a, 0xfa, 0x1f, 0xa1, 0xb4, 0xe1,
		0xbe, 0xd2, 0x36, 0xa7, 0xab, 0x27, 0x3c, 0xb4, 0x5c, 0x4a, 0x6d, 0xab,
		0xe1, 0xb4, 0x27, 0x35, 0x34, 0x66, 0xc0, 0x94, 0xce, 0x02, 0x2b, 0x79,
		0xa3, 0xcd, 0x47, 0x4e, 0xff, 0x75, 0x85, 0x03, 0x97, 0xd0, 0x8e, 0xdb,
		0x8e, 0x75, 0xca, 0xeb, 0xf2, 0xe9, 0x00, 0xe9, 0xf4, 0xa7, 0xca, 0xe9,
		0xc3, 0xe6, 0x26, 0xdc, 0x18, 0x6d, 0x15, 0xab, 0x95, 0xae, 0x6a, 0xc8,
		0x69, 0x96, 0x6e, 0xc2, 0xdd, 0xf1, 0xb8, 0x5a, 0xdd, 0x04, 0x0b, 0x81,
		0x93, 0xd1, 0x0b, 0x6d, 0xd9, 0xf4, 0xf1, 0x14, 0x3c, 0xb2, 0x91, 0x82,
		0x06, 0x33, 0xba, 0x5e, 0x70, 0xf1, 0x52, 0x79, 0xd7, 0x5b, 0x99, 0xd3,
		0xbf, 0xca, 0xb2, 0x2c, 0x14, 0x0f, 0x72, 0xc2, 0x19, 0xe7, 0xf1, 0x46,
		0xed, 0x36, 0x9b, 0xbb, 0xec, 0x44, 0x32, 0x1d, 0x4f, 0x4c, 0x5b, 0x8b,
		0xa4, 0x0e, 0x8a, 0x8e, 0x61, 0x67, 0xf7, 0x5e, 0x35, 0x31, 0xf8, 0x1f,
		0x64, 0xea, 0x6c, 0x10, 0x0b, 0x94, 0x32, 0x6e, 0x74, 0x85, 0xf4, 0x09,
		0x15, 0x5e, 0x66, 0x4c, 0x6c, 0x51, 0xdb, 0x92, 0xae, 0xf6, 0x3a, 0x55,
		0x77, 0x83, 0xaa, 0x78, 0xf5, 0x16, 0x99, 0x29, 0x9c, 0x91, 0x41, 0x7f,
		0xcc, 0x7d, 0xda, 0xf1, 0xa6, 0x3d, 0xc7, 0xcd, 0xba, 0x08, 0xbf, 0xef,
		0xdc, 0x20, 0x5b, 0x93, 0x7c, 0x7f, 0xc5, 0x3d, 0xd3, 0x15, 0xf7, 0xe0,
		0x6d, 0x24, 0xeb, 0x58, 0x30, 0x94, 0xf7, 0xe0, 0x26, 0x5a, 0xb2, 0xe8,
		0xe4, 0x9b, 0x96, 0x50, 0xe7, 0xf4, 0x71, 0xcc, 0xe3, 0xab, 0xf2, 0xa0,
		0x05, 0x37, 0x87, 0xf8, 0xc0, 0xb5, 0x8b, 0x86, 0xb0, 0x46, 0x7f, 0xa3,
		0xf2, 0xb0, 0xca, 0xa2, 0x07, 0xd3, 0xe2, 0xcb, 0x46, 0x47, 0x1c, 0x7a,
		0x52, 0x1a, 0xf7, 0xc6, 0xde, 0x73, 0xda, 0x09, 0xef, 0x8c, 0xf9, 0xee,
		0x47, 0xa5, 0xa1, 0xee, 0x8b, 0x53, 0xf2, 0x36, 0xab, 0xc7, 0xad, 0xd8,
		0xfe, 0x1a, 0x79, 0x51, 0x57, 0xdc, 0xff, 0x20, 0x79, 0x67, 0x0d, 0xb5,
		0x5f, 0xd4, 0x2d, 0x16, 0xcf, 0xf9, 0xd2, 0xdb, 0x93, 0x14, 0x78, 0x35,
		0xa8, 0x39, 0x06, 0x5d, 0x0c, 0x11, 0xa6, 0x38, 0x6e, 0x1a, 0xa6, 0x71,
		0x00, 0x05, 0x04, 0x43, 0x1a, 0x75, 0x0b, 0x13, 0xe0, 0x7a, 0x2d, 0x07,
		0x82, 0x0e, 0x58, 0x61, 0x78, 0xd7, 0x4d, 0x10, 0x5b, 0x51, 0x8a, 0xd5,
		0xa8, 0xed, 0xe0, 0xb7, 0x70, 0xf6, 0xba, 0x3e, 0xd8, 0xc5, 0x96, 0x9a,
		0x48, 0xf2, 0xa9, 0xa7, 0x63, 0x7a, 0x30, 0xd0, 0xe2, 0x45, 0x03, 0x03,
		0x8f, 0xe5, 0xa2, 0x03, 0x5b, 0xac, 0xf5, 0xae, 0x45, 0x3e, 0xb1, 0x94,
		0x06, 0xec, 0x05, 0x94, 0xec, 0x3d, 0x0f, 0x1f, 0x98, 0xa5, 0x74, 0xdd,
		0x5d, 0x00, 0x81, 0xc6, 0x91, 0x55, 0xb1, 0xb2, 0xb7, 0x62, 0xc4, 0x2a,
		0xde, 0x85, 0xe2, 0x1d, 0xe0, 0x8d, 0xfb, 0x5c, 0x32, 0x3b, 0x83, 0x9c,
		0xb1, 0x39, 0x43, 0xfc, 0x68, 0xd0, 0x2d, 0x99, 0x73, 0x0b, 0xc6, 0xdc,
		0x75, 0xa6, 0xe6, 0xbc, 0xe7, 0x75, 0xe8, 0xa9, 0x19, 0xfb, 0x43, 0x66,
		0x9f, 0x1b, 0x25, 0x35, 0x0f, 0x6d, 0xa6, 0x14, 0x02, 0xad, 0xa4, 0x7f,
		0x87, 0x19, 0x1f, 0x8b, 0x7b, 0xfb, 0xb0, 0x6b, 0xdf, 0xff, 0x19, 0xe4,
		0x8e, 0x7f, 0x32, 0x93, 0x34, 0xc7, 0xea, 0xdf, 0xe3, 0x5a, 0x1c, 0x48,
		0xc7, 0x96, 0xd9, 0x5e, 0x29, 0x74, 0x18, 0x2e, 0xb3, 0xd9, 0xb1, 0x8e,
		0x5d, 0x72, 0x61, 0x7a, 0x7c, 0x57, 0x7b, 0xb1, 0xbf, 0x4f, 0x7d, 0xd9,
		0x93, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x06, 0x07, 0xbe, 0x84, 0x0b,
		0x08, 0x00, 0x00,
		},
		"public/css/style.css",
	)
}

func public_css_style_styl() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xb4, 0x54,
		0x6d, 0x6e, 0xa3, 0x30, 0x10, 0xfd, 0xcf, 0x29, 0x46, 0xaa, 0xba, 0xed,
		0x4a, 0x18, 0x25, 0x69, 0xd3, 0xa4, 0xfd, 0xd5, 0x33, 0xec, 0x0d, 0x0c,
		0x1e, 0xc0, 0xaa, 0xb1, 0x91, 0x19, 0x9a, 0xa6, 0xa7, 0xdf, 0x31, 0x90,
		0x00, 0x49, 0x9a, 0x76, 0x7f, 0xac, 0x22, 0x0b, 0x7b, 0x3e, 0xf2, 0xe6,
		0xf9, 0x3d, 0x78, 0xcd, 0x4a, 0xe9, 0x1b, 0x24, 0xb8, 0x6b, 0x29, 0x17,
		0xdb, 0xbb, 0x28, 0xca, 0x9d, 0x25, 0x91, 0xcb, 0x4a, 0x9b, 0xfd, 0xfd,
		0xef, 0x08, 0x60, 0x72, 0x7e, 0x81, 0x3f, 0x2e, 0x75, 0xe4, 0x62, 0x68,
		0xa4, 0x6d, 0x44, 0x83, 0x5e, 0xe7, 0x51, 0xe4, 0xa5, 0xd2, 0x6d, 0xd3,
		0xd5, 0x8a, 0x1d, 0xa6, 0x6f, 0x9a, 0x44, 0xea, 0xbc, 0x42, 0x2f, 0xfa,
		0x0c, 0x48, 0x5f, 0xb4, 0x15, 0x5a, 0x6a, 0x42, 0x45, 0xd5, 0x5c, 0xcb,
		0xba, 0x2b, 0xc9, 0xaf, 0x32, 0x11, 0x79, 0x9e, 0x46, 0x93, 0x76, 0x76,
		0x36, 0xc4, 0x18, 0x16, 0xb5, 0x77, 0x35, 0x7a, 0xda, 0xcf, 0xc1, 0xce,
		0xeb, 0x54, 0xeb, 0x65, 0xd8, 0x40, 0xb2, 0xfa, 0xa2, 0x82, 0x74, 0xa5,
		0x6d, 0x21, 0xf2, 0xd6, 0x66, 0x5d, 0x21, 0xca, 0x06, 0x85, 0xb6, 0x1d,
		0x33, 0xf7, 0xf9, 0x3d, 0xe6, 0x49, 0xd1, 0x29, 0xe0, 0x49, 0xfa, 0x0a,
		0x9a, 0xfb, 0x1e, 0xcb, 0x5d, 0x43, 0x72, 0x3f, 0xc1, 0x89, 0x52, 0xa7,
		0xf6, 0x31, 0x94, 0x4b, 0x5e, 0x2b, 0x5e, 0x0f, 0xbc, 0x1e, 0x79, 0xad,
		0x79, 0x3d, 0xc5, 0x20, 0x63, 0xd0, 0x55, 0x11, 0xb3, 0x45, 0x7c, 0xc5,
		0x5b, 0x5b, 0xb7, 0xc4, 0x09, 0x94, 0x2c, 0x53, 0x08, 0x3a, 0x0a, 0x4f,
		0x65, 0x78, 0x71, 0x5c, 0xa9, 0x18, 0x8c, 0x8e, 0xa1, 0xe5, 0x73, 0x83,
		0x06, 0x33, 0x0a, 0xcf, 0x0e, 0x8f, 0x93, 0xfa, 0x3d, 0x86, 0x77, 0xad,
		0x90, 0xbd, 0x55, 0xc7, 0x90, 0xb6, 0x44, 0x21, 0x4c, 0xf8, 0x41, 0xd2,
		0x23, 0xe3, 0xd4, 0x1e, 0x63, 0xc8, 0x9c, 0x42, 0x9e, 0xbd, 0x62, 0x96,
		0xda, 0xc2, 0x82, 0xb7, 0xb5, 0x54, 0x8a, 0x27, 0xe7, 0x7d, 0x49, 0x95,
		0x99, 0xbb, 0x75, 0x74, 0x6f, 0xa3, 0x3f, 0x11, 0x9e, 0xd6, 0xb7, 0x7c,
		0x36, 0xda, 0xa2, 0x28, 0x51, 0x17, 0x25, 0xc1, 0x32, 0x59, 0x73, 0xe4,
		0x70, 0x58, 0x2c, 0x6e, 0x3b, 0xba, 0x01, 0x40, 0x5b, 0x31, 0x0d, 0xcf,
		0x8b, 0x7a, 0x82, 0x09, 0x69, 0x32, 0x61, 0x9a, 0x54, 0x66, 0x6f, 0x85,
		0x77, 0xad, 0x55, 0x70, 0x93, 0xe7, 0x79, 0x8a, 0x92, 0x83, 0x99, 0x33,
		0xce, 0xc3, 0x0d, 0x6e, 0xd7, 0xeb, 0x87, 0x25, 0x9f, 0x93, 0xbe, 0x89,
		0xef, 0xd4, 0xf2, 0xa5, 0x70, 0x60, 0x1c, 0x7d, 0xf9, 0xe8, 0xb1, 0xea,
		0xc8, 0x30, 0xcc, 0xb2, 0x7b, 0x40, 0x47, 0x5c, 0x48, 0xa3, 0x0b, 0x0b,
		0x19, 0x1e, 0x5b, 0xa6, 0x74, 0x36, 0xdc, 0xd5, 0xff, 0xcf, 0x8f, 0x5b,
		0x1e, 0x0e, 0x2d, 0x43, 0x70, 0xd7, 0x93, 0x4a, 0x9d, 0x51, 0xd1, 0xa0,
		0x44, 0xd2, 0xc8, 0xaa, 0x3e, 0xa3, 0xb5, 0x4a, 0xc3, 0xef, 0x48, 0x6b,
		0x57, 0x6a, 0xc2, 0xc9, 0xe5, 0x1f, 0x09, 0x0c, 0xdd, 0x81, 0xe5, 0x00,
		0x7f, 0x90, 0x0a, 0x64, 0x4b, 0x6e, 0x4e, 0xfb, 0x30, 0xcc, 0x4e, 0x2b,
		0x2a, 0xe1, 0xb9, 0xbb, 0x66, 0x80, 0x77, 0xb6, 0xb1, 0xce, 0xa4, 0x19,
		0x98, 0x90, 0xeb, 0xe9, 0xb1, 0xfc, 0xd3, 0xc9, 0x27, 0x02, 0xcf, 0x29,
		0xb2, 0xa4, 0x23, 0xc9, 0x99, 0xd6, 0xcb, 0x01, 0x00, 0xc0, 0x31, 0x46,
		0x6e, 0xdc, 0x4e, 0x7c, 0x40, 0x93, 0x79, 0x67, 0xcc, 0x91, 0x7c, 0xa1,
		0xa9, 0x6c, 0xd3, 0x13, 0xf2, 0xeb, 0xc5, 0xf3, 0x26, 0xdb, 0xfc, 0x80,
		0x7c, 0xdf, 0xfd, 0x3f, 0xc8, 0xff, 0x93, 0x92, 0x5f, 0xb9, 0x21, 0x21,
		0x59, 0x8c, 0xce, 0x64, 0xa3, 0xa6, 0x8b, 0x45, 0x94, 0xf0, 0x8b, 0x58,
		0x09, 0xcd, 0xaf, 0x66, 0x48, 0x0b, 0xbe, 0x0f, 0x5d, 0xd3, 0x58, 0xb5,
		0x5a, 0x29, 0xb6, 0xf4, 0x50, 0x95, 0x19, 0xd9, 0x34, 0x63, 0x6e, 0x93,
		0xe5, 0x59, 0xf8, 0x87, 0x03, 0xed, 0xcc, 0x85, 0xef, 0xd2, 0x25, 0x17,
		0x8e, 0x73, 0x6f, 0xfb, 0xb9, 0x65, 0x37, 0xe4, 0xfc, 0x3e, 0xb9, 0xf5,
		0xf8, 0x3d, 0xea, 0x53, 0x5d, 0xf0, 0xd7, 0x4b, 0x19, 0xf4, 0x1a, 0x58,
		0xf5, 0x2d, 0x7b, 0x34, 0x2c, 0x5f, 0xf4, 0x5a, 0xa1, 0xd2, 0x32, 0x68,
		0x88, 0x68, 0x41, 0xb2, 0x56, 0xf7, 0xe1, 0xc5, 0xed, 0xae, 0xf4, 0x05,
		0x36, 0x4f, 0xdb, 0xfa, 0x23, 0x18, 0x64, 0xf8, 0x30, 0x4c, 0xc7, 0xe8,
		0xef, 0xfb, 0xcc, 0xf2, 0x97, 0x0c, 0x7c, 0x90, 0x68, 0x73, 0xb4, 0xcf,
		0x68, 0xc5, 0xb9, 0xf1, 0x56, 0xa3, 0x26, 0x17, 0xad, 0x77, 0x66, 0xb2,
		0x4b, 0x96, 0x99, 0xc2, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x59,
		0x4b, 0x72, 0x9b, 0x07, 0x00, 0x00,
		},
		"public/css/style.styl",
	)
}

func public_img_favicon_ico() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xec, 0x5b,
		0x09, 0x74, 0x94, 0xd5, 0xf5, 0xff, 0xfe, 0xe7, 0x6f, 0x7b, 0x3c, 0xa7,
		0x2e, 0xb8, 0xd5, 0xba, 0x61, 0x54, 0x90, 0x04, 0x12, 0x88, 0xa8, 0x20,
		0xc8, 0x1e, 0x5c, 0xaa, 0xb8, 0x61, 0xa9, 0xb6, 0x9e, 0x23, 0xee, 0x75,
		0x6b, 0x5d, 0xf1, 0x54, 0x5c, 0xb1, 0x6a, 0x2b, 0x8a, 0x0b, 0x55, 0x5b,
		0x37, 0x14, 0xb1, 0x88, 0x55, 0x14, 0xb4, 0xee, 0x0b, 0x82, 0x8a, 0x3b,
		0x02, 0x2a, 0x8a, 0x38, 0x7b, 0x26, 0x99, 0x2d, 0xb3, 0x64, 0x32, 0x49,
		0x66, 0x26, 0x33, 0xb7, 0xf7, 0x77, 0xdf, 0xf7, 0xc2, 0x97, 0xc9, 0x44,
		0x33, 0x33, 0x99, 0x50, 0x3c, 0xe4, 0x9c, 0x7b, 0x66, 0xe6, 0xfb, 0xde,
		0x7b, 0x77, 0x79, 0x77, 0x7f, 0x2f, 0x86, 0xf1, 0x7f, 0xc6, 0xff, 0x1b,
		0x03, 0x06, 0xe0, 0xb3, 0xc2, 0xb8, 0x70, 0x3b, 0xc3, 0x18, 0x65, 0x18,
		0x46, 0x45, 0x85, 0xfa, 0xbd, 0x84, 0x9f, 0x2f, 0xe0, 0x67, 0xb5, 0xb5,
		0xe6, 0xef, 0xc1, 0x86, 0x71, 0xf4, 0x1e, 0x86, 0x51, 0xc9, 0x63, 0x06,
		0x60, 0x9c, 0xa1, 0x9e, 0xf7, 0xf4, 0x17, 0xec, 0xa0, 0x87, 0x02, 0x1d,
		0xf4, 0x48, 0x20, 0x95, 0x61, 0xa0, 0x2e, 0xe0, 0x6b, 0xcf, 0x28, 0xc8,
		0x79, 0x27, 0x73, 0x52, 0xb4, 0x30, 0x98, 0xa6, 0x3b, 0x92, 0x44, 0xd4,
		0xc1, 0x90, 0x36, 0x3f, 0x01, 0x29, 0xf3, 0x77, 0x86, 0x21, 0x6b, 0x79,
		0x96, 0x32, 0xbf, 0x63, 0x0e, 0xfe, 0x02, 0x69, 0x8a, 0x33, 0x24, 0x03,
		0xe9, 0x6c, 0x9a, 0x3f, 0x05, 0xfc, 0x29, 0x4a, 0x47, 0x79, 0x7a, 0x33,
		0x03, 0xe3, 0x49, 0xfb, 0x93, 0x94, 0x0e, 0xf3, 0xf7, 0x38, 0x43, 0x53,
		0x96, 0xc7, 0xf0, 0x7b, 0x3f, 0xcf, 0xe1, 0xef, 0xc4, 0x74, 0xdb, 0x98,
		0x86, 0x0c, 0x8f, 0xc3, 0x5a, 0xc4, 0x73, 0x89, 0xe7, 0xd2, 0xd7, 0x0d,
		0x69, 0xba, 0xf3, 0x91, 0x10, 0x4d, 0x3f, 0xcf, 0x4b, 0xc7, 0x9e, 0xe1,
		0xa1, 0x3f, 0xdd, 0xe4, 0xa3, 0x67, 0xdf, 0x6a, 0x26, 0x77, 0x3c, 0x4b,
		0x32, 0x8f, 0xe7, 0xf0, 0xfa, 0xc4, 0xf3, 0x3c, 0x01, 0x73, 0xbe, 0x2f,
		0x99, 0xa5, 0x18, 0x3f, 0xfb, 0x74, 0x53, 0x3b, 0x55, 0xd7, 0xb9, 0xc8,
		0x38, 0xd0, 0x4e, 0xb3, 0xe7, 0xfa, 0xe9, 0x96, 0xf9, 0x01, 0x1a, 0x7f,
		0x8a, 0x9b, 0x0c, 0xe3, 0x7b, 0xba, 0xf0, 0xda, 0x46, 0xaa, 0x4f, 0x64,
		0x29, 0x94, 0xa1, 0x0c, 0xc6, 0x32, 0x2d, 0x32, 0x1f, 0xb8, 0xb1, 0x06,
		0xe0, 0xfc, 0x59, 0x0d, 0x64, 0xec, 0x67, 0xa7, 0xd5, 0x1b, 0xda, 0x28,
		0x61, 0xca, 0x21, 0xc2, 0x30, 0xf3, 0xf2, 0x06, 0x1a, 0x3a, 0xd9, 0x49,
		0xb6, 0x70, 0x07, 0x31, 0x3f, 0x99, 0xa8, 0x65, 0xbe, 0xa6, 0x7b, 0xbd,
		0x27, 0x45, 0xc6, 0x41, 0x76, 0xba, 0xf6, 0x0e, 0x3f, 0xb5, 0x90, 0xa2,
		0xe9, 0x89, 0x65, 0x51, 0xba, 0xe7, 0xf1, 0x26, 0x1a, 0x77, 0xb2, 0x9b,
		0x86, 0x1f, 0xe5, 0xca, 0x3f, 0x9f, 0x05, 0x0a, 0x7a, 0x80, 0xd3, 0x30,
		0x6c, 0x34, 0xff, 0xc9, 0x26, 0x02, 0x7f, 0x0d, 0x6d, 0x59, 0xba, 0xe6,
		0x76, 0x3f, 0x1d, 0x7a, 0xac, 0x8b, 0x4e, 0x39, 0xb7, 0x9e, 0x2e, 0x9a,
		0xdd, 0x48, 0x8e, 0x68, 0x86, 0xc2, 0xd9, 0xae, 0xf4, 0xfb, 0xcc, 0xf9,
		0x6b, 0xec, 0x49, 0x32, 0x76, 0xb7, 0xd3, 0xed, 0x0f, 0x06, 0x15, 0xfe,
		0xf6, 0xac, 0xac, 0x01, 0xfe, 0x98, 0x67, 0x2d, 0xbb, 0x4e, 0xf9, 0x59,
		0xe9, 0xc7, 0x7b, 0x8c, 0x3d, 0xf1, 0xac, 0x7a, 0xaa, 0x3d, 0xda, 0x45,
		0xdf, 0xf1, 0x8b, 0x56, 0x1e, 0xd3, 0xce, 0x60, 0x8f, 0x74, 0xd0, 0x0d,
		0xf3, 0x02, 0xb4, 0xf8, 0x95, 0x98, 0x8c, 0xc3, 0x1c, 0x2b, 0x7e, 0x2d,
		0x7f, 0xac, 0xf9, 0xe6, 0xa7, 0x09, 0x32, 0x7e, 0x61, 0xa3, 0x31, 0x27,
		0xb8, 0xe8, 0xd1, 0xe7, 0x22, 0xf4, 0xd4, 0x7f, 0xa2, 0x74, 0xd2, 0xd9,
		0xf5, 0xcc, 0xd7, 0x46, 0x7a, 0xe4, 0xd9, 0x88, 0xc8, 0x89, 0xf1, 0x75,
		0xc1, 0x1f, 0xb0, 0xec, 0x3f, 0x64, 0xbd, 0x72, 0x7d, 0x2b, 0x9d, 0x73,
		0x65, 0x03, 0x1d, 0x3c, 0xd1, 0x49, 0x07, 0x1e, 0xe9, 0xa4, 0x13, 0x66,
		0x7a, 0xe8, 0xdf, 0x6f, 0x34, 0x0b, 0xfd, 0xc0, 0x95, 0x3b, 0x5f, 0xeb,
		0x8f, 0x06, 0xd0, 0x06, 0x3a, 0x21, 0x2b, 0xc8, 0x5b, 0xcb, 0xc7, 0x32,
		0xce, 0x3a, 0x3f, 0xad, 0x9f, 0x69, 0xe0, 0xfd, 0xc0, 0x9a, 0x19, 0xec,
		0x13, 0xd3, 0x93, 0x81, 0xbe, 0x00, 0xa7, 0xe0, 0x52, 0x63, 0xa0, 0xe3,
		0x58, 0xcf, 0x05, 0x9a, 0xa2, 0xa6, 0x8e, 0xe4, 0x42, 0xd8, 0x02, 0xd6,
		0xe7, 0x98, 0xd3, 0xa6, 0xec, 0xa7, 0x91, 0x79, 0x0f, 0xb2, 0x3d, 0x34,
		0xf1, 0x67, 0x21, 0x10, 0xe4, 0xb9, 0x09, 0x86, 0x75, 0x3d, 0x3a, 0x86,
		0x5e, 0xfe, 0x55, 0x9a, 0x3e, 0x66, 0xd2, 0x8f, 0xf9, 0x99, 0x34, 0x4d,
		0x0e, 0xa5, 0xa9, 0xae, 0x3f, 0x01, 0x38, 0x59, 0x36, 0x47, 0x85, 0x88,
		0x46, 0x5b, 0xe5, 0x1c, 0xed, 0x27, 0x00, 0xce, 0x94, 0xda, 0xa7, 0x4d,
		0xc1, 0x3c, 0x7a, 0x56, 0x10, 0x14, 0x38, 0xd7, 0xc4, 0x95, 0x81, 0x1d,
		0x33, 0xee, 0xaf, 0x2d, 0x7a, 0x5a, 0x30, 0xc0, 0xae, 0x34, 0x14, 0x38,
		0x17, 0x7e, 0x5b, 0xe3, 0x2f, 0x88, 0x7f, 0xf8, 0x34, 0xf8, 0x04, 0x8d,
		0xd3, 0xea, 0xe3, 0xf1, 0x1c, 0x36, 0xd9, 0x8b, 0x75, 0xac, 0xfc, 0xf7,
		0x0a, 0xbf, 0xe6, 0x13, 0x7a, 0x02, 0xda, 0x63, 0xe6, 0x3e, 0x62, 0x6e,
		0x13, 0x7f, 0x6f, 0x36, 0x9f, 0x47, 0x48, 0xd3, 0xd2, 0x77, 0xf8, 0xb1,
		0x16, 0x70, 0x01, 0xe7, 0xb7, 0xac, 0xb8, 0xcf, 0xaf, 0x88, 0xd3, 0xdc,
		0x87, 0x43, 0x34, 0x8b, 0x7d, 0xf8, 0x1f, 0x6f, 0xf4, 0xd1, 0x35, 0x7f,
		0xf5, 0xd3, 0xbc, 0x05, 0x4d, 0xb4, 0x7c, 0x65, 0x9c, 0xbe, 0xf1, 0xa5,
		0x85, 0x86, 0x18, 0xfd, 0xe0, 0xde, 0xf4, 0x1a, 0x3f, 0x64, 0x8a, 0xf5,
		0x10, 0xdf, 0x1e, 0x5c, 0x1c, 0x66, 0xff, 0xcb, 0x31, 0xaf, 0xd2, 0xc1,
		0x60, 0x27, 0x63, 0xb0, 0x09, 0x55, 0x0c, 0x43, 0x14, 0x1c, 0xc1, 0xef,
		0x6f, 0x99, 0x1f, 0xa4, 0xcf, 0x6d, 0x49, 0x32, 0x7d, 0x59, 0x3e, 0x1a,
		0x7a, 0x85, 0x1f, 0x7c, 0x03, 0x37, 0x7c, 0xed, 0xb9, 0x57, 0x37, 0x48,
		0x0c, 0x3b, 0x68, 0x9c, 0x93, 0xf6, 0x3a, 0xdc, 0x21, 0x71, 0x78, 0xcf,
		0xc3, 0x1c, 0xb4, 0xff, 0x58, 0x27, 0x19, 0x43, 0xf9, 0xf7, 0x20, 0xbb,
		0xbc, 0x3b, 0x00, 0xbf, 0xf7, 0xe4, 0x77, 0x87, 0x3b, 0x69, 0xc1, 0xf3,
		0xd1, 0xa2, 0xe5, 0x0f, 0x9a, 0x35, 0xed, 0xc8, 0x13, 0x10, 0x7b, 0x8e,
		0x98, 0xe6, 0xa2, 0x9f, 0xd5, 0x38, 0x68, 0xda, 0xcc, 0x7a, 0x5a, 0xf4,
		0x52, 0x8c, 0x56, 0x7d, 0xd9, 0x46, 0xab, 0xbf, 0x69, 0xa7, 0x57, 0x56,
		0xb7, 0x70, 0x6c, 0x0c, 0xd1, 0x21, 0xc7, 0xb8, 0x68, 0xf7, 0x91, 0x0e,
		0x1a, 0xcd, 0xe3, 0xb6, 0xab, 0x76, 0x48, 0x5c, 0xfe, 0x8e, 0x17, 0x89,
		0x50, 0x37, 0x7d, 0xf8, 0x51, 0xfc, 0x18, 0x0f, 0x9d, 0x7a, 0xfe, 0x9d,
		0x38, 0x19, 0x3b, 0xdb, 0x68, 0xd4, 0x71, 0x9c, 0x7b, 0x1c, 0x6c, 0xa7,
		0x4b, 0xae, 0xf7, 0x49, 0xfc, 0x49, 0x98, 0x3a, 0x17, 0x35, 0xf7, 0x1a,
		0x38, 0xd6, 0xb9, 0x52, 0x74, 0xd6, 0x15, 0x0d, 0xb4, 0x1b, 0xd3, 0x00,
		0x39, 0x4c, 0x3d, 0xcd, 0x53, 0x14, 0x7e, 0x1d, 0xb7, 0x1b, 0x39, 0xce,
		0x63, 0xbd, 0xed, 0x87, 0x3b, 0x24, 0x5e, 0x4e, 0x3c, 0xd5, 0x2d, 0xeb,
		0xc5, 0xcd, 0x1c, 0x02, 0xba, 0x81, 0x71, 0x61, 0xd3, 0x16, 0x90, 0xd7,
		0x3d, 0xf4, 0x4c, 0x84, 0x8c, 0xbd, 0xec, 0x32, 0x7e, 0xd2, 0x6f, 0x3c,
		0xec, 0xe0, 0x0a, 0xc7, 0xaf, 0xe3, 0xea, 0x67, 0xac, 0x43, 0x07, 0xf0,
		0x3a, 0xd5, 0x53, 0x78, 0x4f, 0xf7, 0xb1, 0xd3, 0xdf, 0x17, 0x85, 0x85,
		0x67, 0xe0, 0xd5, 0xb6, 0x0f, 0x7a, 0xd6, 0xba, 0x92, 0x02, 0x1b, 0x1a,
		0xd3, 0x34, 0xe7, 0xbe, 0xa0, 0xe8, 0xe1, 0xc0, 0x23, 0x1c, 0x8c, 0xdf,
		0x5d, 0x14, 0xff, 0xc8, 0x6f, 0xc0, 0xe3, 0xf2, 0x55, 0x71, 0xc9, 0xbb,
		0x86, 0x4f, 0x65, 0xfc, 0x55, 0x0e, 0x7a, 0xe3, 0x93, 0x84, 0xe0, 0x87,
		0x5c, 0x30, 0x0e, 0x7c, 0x63, 0xdf, 0x0d, 0xde, 0xeb, 0x31, 0x27, 0xba,
		0x69, 0x24, 0xe7, 0x53, 0x23, 0x38, 0xc7, 0xc1, 0x27, 0xe4, 0x7f, 0xf4,
		0xef, 0x98, 0xff, 0x50, 0x31, 0xfc, 0x2b, 0xfc, 0x0b, 0x97, 0x47, 0xc9,
		0xd8, 0xdf, 0x2e, 0xf9, 0x21, 0xf4, 0xfc, 0xa3, 0x8d, 0xed, 0x22, 0x17,
		0x8d, 0x1f, 0xeb, 0x22, 0x27, 0x34, 0x8c, 0x4d, 0xa2, 0x1b, 0xc6, 0x9e,
		0x36, 0x19, 0x2f, 0xf6, 0xb8, 0xbd, 0x8d, 0x86, 0x4d, 0x71, 0xd1, 0xc6,
		0x40, 0xf1, 0xf8, 0x1f, 0x5b, 0xaa, 0xf0, 0x57, 0x4e, 0x72, 0x4a, 0xce,
		0x04, 0x9b, 0x8e, 0x99, 0xf9, 0x1f, 0xc6, 0x61, 0xcf, 0x91, 0x87, 0x9e,
		0x37, 0xab, 0x91, 0xae, 0xba, 0xd5, 0x4f, 0x97, 0xcf, 0xf1, 0xd1, 0x65,
		0x37, 0xfb, 0xe8, 0x0a, 0xfe, 0xbc, 0xf8, 0xba, 0x46, 0xba, 0x76, 0x6e,
		0x80, 0x73, 0xbc, 0x8c, 0xc8, 0x29, 0xc7, 0x07, 0xf4, 0x4a, 0xfe, 0xcf,
		0xbc, 0x1e, 0x23, 0xa3, 0xc2, 0xce, 0x7c, 0x38, 0xc5, 0x9e, 0xdf, 0xff,
		0xba, 0xad, 0xcb, 0xfe, 0x0b, 0x74, 0x6c, 0xce, 0xb1, 0x03, 0x1d, 0x9b,
		0x7d, 0xaf, 0xbf, 0x04, 0xff, 0xab, 0x73, 0xce, 0x15, 0x5f, 0xb4, 0x92,
		0xc1, 0xf6, 0x2e, 0xfb, 0xcf, 0x72, 0x58, 0xfa, 0x4e, 0xb3, 0xd0, 0xa5,
		0xf9, 0xd7, 0x3a, 0x68, 0x05, 0xbc, 0x4f, 0x58, 0x62, 0x41, 0x31, 0xf8,
		0xc1, 0x03, 0x64, 0xf6, 0x7d, 0x53, 0x07, 0xd5, 0xfd, 0xd6, 0xa3, 0x7c,
		0x1a, 0xef, 0x29, 0xfc, 0x10, 0x64, 0xce, 0x2a, 0x25, 0x3a, 0xa0, 0xfd,
		0xbb, 0xd6, 0x07, 0x7c, 0x87, 0x5f, 0x5a, 0xc2, 0x72, 0xfb, 0xf0, 0xdb,
		0x76, 0xd1, 0xbd, 0x52, 0xfc, 0x1f, 0x7c, 0xcb, 0xbc, 0xc7, 0x9a, 0xc8,
		0xd8, 0xcd, 0x46, 0x87, 0xc3, 0xff, 0xb0, 0x5d, 0x3d, 0xfd, 0x6a, 0x4c,
		0x6a, 0x00, 0x1d, 0xff, 0xb4, 0xfd, 0x23, 0x9f, 0x7f, 0xed, 0x23, 0xce,
		0xd1, 0xf7, 0x30, 0x63, 0x01, 0xfb, 0xe0, 0xbb, 0x1e, 0x6d, 0xea, 0x8c,
		0x43, 0x85, 0xe2, 0xd7, 0xbe, 0x7f, 0xa3, 0xbf, 0x43, 0x6c, 0x6b, 0xc7,
		0x11, 0x0e, 0x1a, 0xc1, 0xfe, 0x74, 0xc0, 0x21, 0x0e, 0xba, 0x8b, 0x69,
		0x5a, 0xcb, 0xbe, 0xce, 0x1d, 0xcf, 0x88, 0x1f, 0x82, 0x9c, 0x96, 0xbc,
		0xde, 0x2c, 0xf5, 0xd9, 0xc1, 0x13, 0x9c, 0x52, 0xd7, 0x20, 0x2e, 0xad,
		0x58, 0xdb, 0xda, 0x5d, 0x5f, 0x0a, 0x8c, 0x3f, 0x98, 0xff, 0xce, 0x9a,
		0x56, 0xda, 0x7b, 0x34, 0xe2, 0x8c, 0x9d, 0x0e, 0x61, 0xfb, 0x46, 0xbc,
		0xa9, 0x99, 0xea, 0xa2, 0x19, 0x17, 0xd4, 0xd3, 0x99, 0x97, 0x79, 0x69,
		0x0a, 0xef, 0x91, 0x31, 0xcc, 0x41, 0x55, 0x6c, 0x27, 0xa0, 0x11, 0x71,
		0x6a, 0x2e, 0xd7, 0x99, 0xe1, 0x22, 0xf7, 0x3f, 0x77, 0x1f, 0x40, 0xc3,
		0x07, 0xac, 0xfb, 0x27, 0x9f, 0xc3, 0xf5, 0xce, 0xaf, 0xec, 0x2a, 0xde,
		0x55, 0xdb, 0xc5, 0x37, 0x19, 0x03, 0xed, 0x12, 0x8b, 0x77, 0x61, 0xb9,
		0xc8, 0x77, 0xe6, 0xfb, 0x8e, 0x87, 0x42, 0x2a, 0x67, 0xc8, 0x94, 0x1e,
		0xff, 0x85, 0x06, 0xd3, 0x1f, 0x43, 0xde, 0x90, 0x33, 0x6c, 0x7c, 0xda,
		0x99, 0xf5, 0x34, 0x9e, 0x6b, 0xcf, 0x23, 0x4f, 0x72, 0xd3, 0xe4, 0x19,
		0x1e, 0x3a, 0xfd, 0x62, 0x2f, 0xdd, 0x7a, 0x7f, 0x90, 0x56, 0xad, 0x6f,
		0x13, 0xbe, 0x43, 0x3d, 0xe3, 0x2e, 0x18, 0xbf, 0xb6, 0x49, 0xe8, 0x5b,
		0x9c, 0x36, 0xd7, 0x60, 0xc8, 0x75, 0xe0, 0xf3, 0xe1, 0xe3, 0x91, 0x9f,
		0x44, 0xcd, 0x98, 0xa8, 0xe5, 0x56, 0xac, 0xfd, 0xff, 0xd0, 0x5e, 0x68,
		0x5d, 0xca, 0xad, 0xbd, 0x34, 0xbf, 0x79, 0x74, 0xad, 0xcf, 0xf0, 0x77,
		0xa7, 0x65, 0x33, 0x14, 0x98, 0x83, 0x97, 0x8c, 0xbf, 0x44, 0xf8, 0x9f,
		0xc1, 0x1f, 0xcc, 0x53, 0xa7, 0x97, 0x13, 0x82, 0x0a, 0xd2, 0xa6, 0x9e,
		0x7e, 0xa9, 0xf3, 0xfa, 0xfe, 0xaa, 0x3d, 0xa5, 0xfe, 0x24, 0xb3, 0xfe,
		0x4c, 0x49, 0x7a, 0xe6, 0x62, 0xf0, 0xf6, 0x17, 0x30, 0xcf, 0xf8, 0xf4,
		0xb0, 0x0c, 0x22, 0xfc, 0x7d, 0x45, 0xa9, 0x7d, 0x86, 0x52, 0xff, 0xd0,
		0xa7, 0xa8, 0x65, 0x38, 0xcb, 0xf8, 0xe1, 0x3e, 0x85, 0xbc, 0xde, 0x8a,
		0x81, 0x65, 0xfd, 0x39, 0xc3, 0xba, 0xad, 0x0c, 0xbe, 0x30, 0x75, 0x74,
		0x11, 0x7c, 0x6c, 0x8b, 0x99, 0x57, 0x6d, 0x2d, 0x00, 0x5b, 0xcb, 0x28,
		0x3b, 0xff, 0x02, 0xbd, 0xe9, 0x40, 0x4a, 0x7c, 0x4e, 0xbf, 0xf8, 0x99,
		0x3e, 0x02, 0xed, 0xa7, 0x3e, 0xc4, 0xf7, 0x2d, 0xe0, 0xa7, 0xfb, 0xc4,
		0xcf, 0x73, 0x5d, 0xb0, 0x8d, 0xfe, 0x9f, 0x18, 0xfd, 0xd6, 0x7e, 0x9f,
		0x3f, 0xb9, 0xb9, 0xd7, 0xe5, 0x4b, 0x5a, 0x9e, 0x17, 0xde, 0x97, 0x2c,
		0x2b, 0xfd, 0x9a, 0x66, 0xac, 0x13, 0x31, 0x73, 0x55, 0xdd, 0xc3, 0xd1,
		0xfd, 0xe7, 0x18, 0x75, 0xed, 0x2f, 0xea, 0x39, 0x45, 0xf2, 0xd2, 0x67,
		0xf4, 0x5b, 0xf3, 0x6f, 0xc8, 0x76, 0x8d, 0x23, 0x49, 0xcb, 0x56, 0xb6,
		0xd0, 0x3f, 0x97, 0x44, 0xe8, 0xf6, 0x7f, 0x84, 0xe8, 0xc6, 0x7b, 0x82,
		0x34, 0x67, 0x7e, 0x90, 0xee, 0x7b, 0x32, 0x4c, 0x4f, 0xbf, 0xd6, 0x4c,
		0xef, 0x7d, 0xd9, 0x26, 0xb5, 0xbf, 0xce, 0xc9, 0xf5, 0x59, 0x4e, 0x7f,
		0xd3, 0xaf, 0xf5, 0x02, 0x74, 0xa3, 0x66, 0x58, 0xf8, 0x62, 0x8c, 0x4e,
		0xbf, 0xc4, 0x4b, 0xfb, 0x8d, 0x51, 0x35, 0x3d, 0xfa, 0x4a, 0xe8, 0x31,
		0xa0, 0x9f, 0x28, 0xb0, 0x1f, 0xc3, 0xde, 0x76, 0xe9, 0xe9, 0x4c, 0x98,
		0xee, 0xa6, 0xeb, 0xe7, 0x05, 0xe8, 0x8d, 0x4f, 0x13, 0x5c, 0x63, 0xab,
		0x35, 0x14, 0x1f, 0xfd, 0x43, 0xbf, 0xee, 0x41, 0x42, 0x86, 0x2f, 0xbd,
		0xdf, 0x22, 0x75, 0x16, 0x68, 0x1b, 0x50, 0xeb, 0xa0, 0xca, 0x89, 0x4e,
		0xe9, 0x4b, 0xa0, 0x1e, 0x1d, 0x79, 0x8c, 0x02, 0xf4, 0x1f, 0x6b, 0xb9,
		0xfe, 0xc4, 0xf3, 0x61, 0x93, 0x9d, 0xd2, 0x03, 0x43, 0x5f, 0x14, 0xbd,
		0xa1, 0x19, 0x17, 0x7a, 0xa5, 0x6f, 0x8c, 0x75, 0xa3, 0xf9, 0x6b, 0xf2,
		0x3e, 0xa5, 0x5f, 0xd3, 0x0e, 0x9d, 0xb9, 0x77, 0x61, 0x58, 0xe4, 0xbc,
		0xef, 0x68, 0x55, 0xc3, 0xa3, 0x9f, 0x55, 0x53, 0xe7, 0x94, 0xbe, 0xce,
		0xde, 0xa3, 0x1c, 0xb4, 0x5d, 0x8d, 0x43, 0x6a, 0xeb, 0x9d, 0xb8, 0xc6,
		0xaf, 0x18, 0xab, 0xde, 0x55, 0x33, 0xa0, 0x9e, 0xae, 0x36, 0xbf, 0xa3,
		0xa7, 0x89, 0xbd, 0x41, 0xaf, 0x69, 0x9d, 0x3b, 0xd5, 0x69, 0x1b, 0xe5,
		0xa4, 0x1f, 0x72, 0xba, 0xfb, 0xf1, 0x26, 0x32, 0x76, 0xb4, 0x89, 0x4c,
		0x41, 0xf7, 0x90, 0x09, 0x66, 0x3f, 0xb2, 0x4a, 0xe9, 0x0b, 0xce, 0x80,
		0x7f, 0x7f, 0x69, 0x03, 0x9d, 0x73, 0x75, 0x23, 0x4d, 0x3f, 0xdf, 0x2b,
		0x3d, 0x5f, 0xd1, 0xa9, 0x4a, 0xd5, 0xb7, 0x03, 0x0c, 0x9e, 0xa0, 0x3e,
		0x71, 0xfe, 0x69, 0xec, 0x6b, 0xa7, 0xba, 0xd3, 0x3c, 0xa2, 0x8b, 0xbd,
		0xe0, 0xa1, 0x28, 0xfa, 0x75, 0xff, 0x0a, 0xf6, 0x89, 0x3e, 0xcc, 0x08,
		0xa6, 0x7d, 0xc8, 0x44, 0x05, 0xd0, 0x1b, 0xe8, 0xc4, 0x45, 0xb3, 0x7d,
		0xf4, 0xe6, 0x67, 0x09, 0xb2, 0x85, 0x33, 0xa2, 0xdb, 0x3e, 0xa6, 0xc3,
		0xdb, 0x9a, 0x95, 0x33, 0xf2, 0x65, 0x2b, 0xe3, 0x74, 0xf1, 0x75, 0x3e,
		0xe9, 0x33, 0xa0, 0xaf, 0x80, 0x7d, 0x00, 0x0f, 0xd2, 0x07, 0xe4, 0xcf,
		0x5a, 0xd6, 0x33, 0x9c, 0x43, 0xe4, 0xe9, 0xeb, 0x95, 0x4c, 0x3f, 0xd6,
		0x83, 0x5c, 0x40, 0xd7, 0x78, 0xb6, 0xbf, 0x7d, 0x58, 0x3f, 0x20, 0x3b,
		0xe0, 0x05, 0x1d, 0xe8, 0xc7, 0x3d, 0xfa, 0x5c, 0xb4, 0xf3, 0xcc, 0x45,
		0x9f, 0xc1, 0x00, 0x50, 0x83, 0xeb, 0xf3, 0x19, 0x3c, 0x7f, 0x7b, 0x4d,
		0xab, 0x9c, 0x6f, 0xef, 0x33, 0x5a, 0xf1, 0x00, 0xfa, 0x61, 0x13, 0xa3,
		0xa6, 0xb9, 0xa5, 0x2f, 0x5a, 0x0e, 0xfa, 0x75, 0xcf, 0x68, 0xc1, 0x0b,
		0x51, 0x32, 0x76, 0xb1, 0xc9, 0x9e, 0x0f, 0x1e, 0xaf, 0x74, 0x1a, 0xfa,
		0x72, 0xff, 0xbf, 0xc2, 0xd2, 0x1f, 0x94, 0xb1, 0x96, 0x1e, 0x9e, 0x15,
		0x54, 0xfc, 0xca, 0x4a, 0x8d, 0xb6, 0x7c, 0x55, 0x8b, 0xe8, 0x0c, 0xe6,
		0x83, 0xfe, 0xfd, 0xc7, 0x38, 0xe8, 0xb0, 0xe3, 0xca, 0x43, 0xbf, 0xb6,
		0x59, 0xe8, 0xc1, 0xa9, 0x17, 0x78, 0x69, 0xe7, 0x5a, 0x25, 0x7b, 0xd8,
		0x29, 0xfa, 0x6f, 0x33, 0xfe, 0xe0, 0xa5, 0xc6, 0xa4, 0x92, 0x73, 0xae,
		0x0f, 0x14, 0xda, 0x2d, 0x71, 0x18, 0x7d, 0x39, 0xd8, 0xcf, 0x0b, 0xec,
		0x73, 0xe0, 0x5f, 0xab, 0xfb, 0x81, 0x7e, 0xd0, 0x04, 0x9c, 0x1f, 0x7f,
		0xd7, 0x4e, 0x3b, 0xd4, 0xaa, 0x3d, 0x87, 0xde, 0xd4, 0xa0, 0x7f, 0xcb,
		0x7e, 0x73, 0xf1, 0xab, 0x31, 0xd1, 0x8d, 0xc6, 0xf6, 0xee, 0x71, 0x48,
		0x9f, 0x03, 0x6b, 0x00, 0x6d, 0x49, 0xf1, 0xbb, 0x09, 0xf1, 0x3b, 0xfd,
		0x21, 0x7f, 0x6d, 0xb7, 0x88, 0x9f, 0xc6, 0x2f, 0xed, 0xe2, 0x2b, 0x61,
		0xb3, 0xd0, 0x9f, 0x81, 0x1c, 0xaf, 0xd6, 0x3a, 0x53, 0xca, 0x77, 0x27,
		0xbb, 0xfb, 0xa9, 0x97, 0x57, 0xb7, 0x48, 0x5f, 0xee, 0x86, 0xbb, 0x03,
		0x74, 0xdd, 0x5d, 0x01, 0x9a, 0x7d, 0x67, 0x80, 0x6e, 0xba, 0x37, 0x48,
		0x33, 0xaf, 0x68, 0x10, 0xfd, 0x87, 0xdd, 0xf7, 0x87, 0xfe, 0x20, 0x6f,
		0xc1, 0xd9, 0x24, 0xce, 0x80, 0x86, 0x4f, 0x55, 0xbe, 0x1e, 0xbe, 0xfb,
		0xf8, 0x33, 0xeb, 0xe5, 0x9e, 0x4d, 0x6e, 0x5f, 0x50, 0x9f, 0x35, 0x3c,
		0xb8, 0x38, 0x22, 0xe7, 0x18, 0x3f, 0xaf, 0x51, 0x67, 0x7b, 0xd2, 0xdb,
		0x64, 0xbd, 0xd9, 0xc9, 0x8c, 0x75, 0x90, 0x81, 0xb6, 0xdf, 0x72, 0xd2,
		0x8f, 0x35, 0xff, 0xfc, 0x37, 0xbf, 0xe4, 0x06, 0xd0, 0x1b, 0xe8, 0x3e,
		0x62, 0xd3, 0xd9, 0x57, 0x36, 0x76, 0x39, 0x13, 0xcc, 0xdd, 0x33, 0x9c,
		0x35, 0xc2, 0x5f, 0x8e, 0x3d, 0xd1, 0x25, 0xfd, 0x67, 0xc0, 0x61, 0xbf,
		0x56, 0x31, 0xb9, 0xa6, 0xce, 0x25, 0xb2, 0xa8, 0x35, 0xf7, 0x73, 0xec,
		0x49, 0xe5, 0xf3, 0x3f, 0x18, 0x73, 0xc9, 0x0d, 0x3e, 0x89, 0xf9, 0xb0,
		0x39, 0xf1, 0x3b, 0xec, 0xef, 0x2f, 0xbf, 0xd9, 0x67, 0xbd, 0xa7, 0xd3,
		0x4d, 0xfe, 0x0f, 0x2c, 0x0e, 0xab, 0x73, 0xa4, 0x2a, 0x33, 0x07, 0x1a,
		0x98, 0x07, 0x0e, 0x60, 0xd8, 0xd5, 0x46, 0x03, 0x46, 0x3a, 0xd9, 0xff,
		0x97, 0x87, 0x7e, 0xe8, 0x07, 0xf4, 0x18, 0x34, 0x77, 0xd2, 0xcf, 0xba,
		0x80, 0xb3, 0xa5, 0x7c, 0xf4, 0x6b, 0xfd, 0x7f, 0xf5, 0xc3, 0x84, 0x9c,
		0x41, 0xe1, 0xfc, 0x0f, 0xb9, 0x68, 0x3e, 0x80, 0x3d, 0x20, 0x9f, 0xfb,
		0xcb, 0xfd, 0x21, 0xb2, 0x97, 0x21, 0xfe, 0xea, 0xfb, 0x01, 0xb8, 0xd3,
		0x04, 0x19, 0x0e, 0xb7, 0xea, 0xcf, 0x55, 0xf9, 0xf5, 0xa7, 0x27, 0xff,
		0xf3, 0x63, 0xd0, 0xcb, 0x5c, 0xa0, 0x60, 0xff, 0x09, 0x5d, 0x9e, 0xbf,
		0x28, 0x2c, 0x79, 0x83, 0xb6, 0xdf, 0x3d, 0x0e, 0x75, 0xd0, 0x31, 0x67,
		0x78, 0xc8, 0xd5, 0xdc, 0xb3, 0xcc, 0xf2, 0xc5, 0x31, 0x89, 0x65, 0xec,
		0x6b, 0xe1, 0x6f, 0x75, 0x0d, 0x53, 0xe0, 0x7d, 0x91, 0xa2, 0xe8, 0x7f,
		0xf1, 0xbd, 0x16, 0xf1, 0xf7, 0x90, 0x3f, 0xf2, 0x35, 0xd8, 0xdc, 0xae,
		0xec, 0x83, 0x3e, 0xd9, 0x64, 0x9e, 0x5f, 0xf6, 0x32, 0x7f, 0xd7, 0xba,
		0x65, 0xbd, 0xeb, 0xe1, 0x2f, 0xac, 0xae, 0x2c, 0x4a, 0x7f, 0xbe, 0xf2,
		0xa6, 0xa9, 0x6a, 0xb2, 0x4b, 0xee, 0x33, 0x0c, 0x31, 0xf3, 0x7c, 0xf0,
		0xf3, 0x00, 0xe7, 0x0e, 0x3d, 0xc5, 0xaf, 0x7c, 0x6b, 0xc1, 0x96, 0x56,
		0xae, 0x6f, 0xa3, 0x77, 0xd7, 0xb6, 0xca, 0x19, 0x09, 0xf2, 0x3c, 0x5d,
		0x6b, 0x06, 0xca, 0x98, 0xff, 0x87, 0xb2, 0xa6, 0x0d, 0x0f, 0x56, 0x71,
		0xb3, 0xca, 0x12, 0x77, 0x70, 0x66, 0x03, 0x79, 0xea, 0xfb, 0x83, 0xf9,
		0x00, 0xfc, 0x85, 0xcd, 0x7a, 0xad, 0xe6, 0x28, 0x95, 0x33, 0x8f, 0x3a,
		0xde, 0x4d, 0x67, 0x5c, 0xea, 0xa5, 0xab, 0x6f, 0xf3, 0xd3, 0x6b, 0x1f,
		0x27, 0x7a, 0xe3, 0x7b, 0x8a, 0xa2, 0x5f, 0xdf, 0x3d, 0x78, 0xf3, 0xb3,
		0x56, 0xc1, 0x8b, 0x18, 0x8c, 0xd8, 0x03, 0x5f, 0x84, 0x33, 0xbe, 0x73,
		0x67, 0xa9, 0xbb, 0x9e, 0x71, 0xea, 0x9a, 0xab, 0x75, 0x82, 0x49, 0x3b,
		0xf4, 0xf0, 0xb6, 0x07, 0x42, 0xe2, 0x4f, 0x11, 0x03, 0x2a, 0xc6, 0x98,
		0xf5, 0x98, 0xf1, 0x3d, 0x3d, 0xf5, 0x72, 0xcc, 0x3c, 0x97, 0x2c, 0x8f,
		0xfc, 0x01, 0x41, 0xf8, 0xd1, 0x39, 0x7e, 0x89, 0xa1, 0xc8, 0x41, 0x07,
		0xe9, 0x1c, 0x94, 0x7d, 0x11, 0xce, 0xea, 0x3e, 0xd8, 0xd0, 0x2e, 0xfb,
		0x94, 0xdb, 0xb3, 0x04, 0x5d, 0xb0, 0x73, 0xd4, 0xf4, 0xb8, 0xf3, 0xa0,
		0x73, 0x6f, 0xa9, 0x63, 0x78, 0x8d, 0x71, 0x5c, 0xef, 0xe0, 0xbc, 0xb5,
		0x9c, 0xf5, 0x97, 0x8e, 0xc3, 0x38, 0xdb, 0x83, 0xce, 0xe0, 0x4e, 0x0d,
		0xf6, 0x61, 0x90, 0xb9, 0x0f, 0xa8, 0x49, 0x76, 0x66, 0xc0, 0xbd, 0xdb,
		0x87, 0x9f, 0x8d, 0x88, 0xbd, 0xbf, 0xc5, 0xb5, 0xcc, 0xb2, 0x77, 0xe3,
		0x72, 0xe6, 0x8c, 0x33, 0x73, 0xe8, 0xde, 0x30, 0x93, 0x76, 0xd4, 0x2e,
		0xb8, 0x93, 0x81, 0x3b, 0x11, 0xb8, 0xd3, 0xd0, 0x4b, 0xd9, 0x97, 0x24,
		0x7f, 0xad, 0x47, 0xb8, 0xb3, 0x31, 0x64, 0x12, 0xe3, 0x1e, 0xa6, 0xf6,
		0x01, 0xb9, 0x0c, 0x64, 0x29, 0xb5, 0xcc, 0x50, 0x33, 0xd6, 0x22, 0xae,
		0x0e, 0x32, 0xfb, 0x10, 0x9c, 0x77, 0xa0, 0xe6, 0x01, 0x9f, 0xba, 0xe6,
		0x81, 0xfe, 0x20, 0xee, 0xe2, 0xae, 0x97, 0x8e, 0x21, 0xe5, 0xf0, 0x3f,
		0xdd, 0x79, 0x50, 0xb9, 0x0d, 0x7a, 0x3d, 0x27, 0x73, 0x1d, 0x85, 0x3b,
		0xc0, 0xb8, 0x8b, 0x00, 0x7f, 0x84, 0x9c, 0x06, 0x7b, 0x32, 0xc2, 0xec,
		0x37, 0x40, 0xb7, 0x10, 0x2f, 0xf0, 0x5b, 0xd7, 0xef, 0x78, 0x0e, 0xbb,
		0x37, 0x76, 0xb0, 0x89, 0x3f, 0xd0, 0xf7, 0xac, 0xcb, 0xe5, 0x3f, 0x7b,
		0xe2, 0x01, 0xfb, 0xe0, 0x69, 0x51, 0xf7, 0xd7, 0x44, 0x96, 0xfb, 0xaa,
		0x9e, 0x08, 0xf2, 0x52, 0xf4, 0x1b, 0x06, 0x8d, 0x53, 0xba, 0x0d, 0x7f,
		0x0b, 0x7a, 0xd1, 0x5f, 0x91, 0xfd, 0xe0, 0x9c, 0x07, 0x77, 0x0f, 0x9e,
		0x58, 0x1e, 0xed, 0xd4, 0xc9, 0x02, 0x7a, 0x3f, 0x7d, 0x42, 0xbf, 0xe6,
		0x01, 0x72, 0xc3, 0x5e, 0x7c, 0xcb, 0x8b, 0xa1, 0x8e, 0xc1, 0x9d, 0xbf,
		0xe9, 0xe7, 0xd5, 0x8b, 0x5f, 0x1c, 0x3a, 0x45, 0xe5, 0x95, 0xe8, 0x3d,
		0x80, 0xde, 0xd3, 0x2e, 0xf2, 0xca, 0xdd, 0x9b, 0xa5, 0x6f, 0xc7, 0xc5,
		0x56, 0x75, 0xdf, 0xaa, 0x88, 0x1e, 0x62, 0x9f, 0xf5, 0x0f, 0xb5, 0xaf,
		0xd4, 0xf5, 0x39, 0x3e, 0x51, 0x67, 0x22, 0x2e, 0x6d, 0x68, 0x48, 0xd3,
		0x57, 0xf5, 0x29, 0x89, 0x0d, 0xb8, 0xb3, 0xa2, 0xe3, 0x54, 0x9c, 0xb4,
		0xcc, 0xb7, 0x7c, 0xff, 0x33, 0x97, 0x0f, 0x1d, 0x5f, 0x73, 0xef, 0x01,
		0x84, 0x2d, 0xb2, 0x2e, 0x81, 0xee, 0xb2, 0xd1, 0x9f, 0xcb, 0x4b, 0x27,
		0x24, 0x4b, 0xba, 0xcf, 0xbb, 0x45, 0xe8, 0xef, 0x07, 0xd8, 0x46, 0xff,
		0x36, 0xfa, 0xb7, 0xd1, 0xff, 0x13, 0xa0, 0x3f, 0x28, 0xff, 0x4f, 0xa6,
		0x9e, 0x6d, 0x45, 0xd0, 0x79, 0xfe, 0xbe, 0x25, 0xee, 0x09, 0x95, 0x0a,
		0x88, 0x85, 0x69, 0x45, 0xff, 0xe7, 0x81, 0x54, 0x16, 0x2d, 0x2f, 0xdb,
		0x56, 0x05, 0x29, 0xda, 0x14, 0x50, 0xff, 0x4b, 0xb5, 0xb4, 0x98, 0xbb,
		0x45, 0xdb, 0xfe, 0xfa, 0xee, 0xef, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
		0xc5, 0xe3, 0x5b, 0x7c, 0xee, 0x3a, 0x00, 0x00,
		},
		"public/img/favicon.ico",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"templates/layout.gold": templates_layout_gold,
	"templates/top.gold": templates_top_gold,
	"public/css/style.css": public_css_style_css,
	"public/css/style.styl": public_css_style_styl,
	"public/img/favicon.ico": public_img_favicon_ico,

}
