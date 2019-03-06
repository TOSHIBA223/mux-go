// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxvideosdk

type AssetStaticRenditionsFiles struct {
	Name string `json:"name,omitempty"`
	// Extension of the static rendition file
	Ext string `json:"ext,omitempty"`
	// The height of the static rendition's file in pixels
	Height int32 `json:"height,omitempty"`
	// The width of the static rendition's file in pixels
	Width int32 `json:"width,omitempty"`
	// The bitrate in bits per second
	Bitrate int64 `json:"bitrate,omitempty"`
	Filesize string `json:"filesize,omitempty"`
}