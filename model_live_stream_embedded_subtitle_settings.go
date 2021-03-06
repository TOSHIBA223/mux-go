// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type LiveStreamEmbeddedSubtitleSettings struct {
	// A name for this live stream closed caption track.
	Name string `json:"name,omitempty"`
	// Arbitrary user-supplied metadata set for the live stream closed caption track. Max 255 characters.
	Passthrough string `json:"passthrough,omitempty"`
	// The language of the closed caption stream. Value must be BCP 47 compliant.
	LanguageCode string `json:"language_code,omitempty"`
	// CEA-608 caption channel to read data from.
	LanguageChannel string `json:"language_channel,omitempty"`
}
