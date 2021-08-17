// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

// An array of objects that each describe an input file to be used to create the asset. As a shortcut, `input` can also be a string URL for a file when only one input file is used. See `input[].url` for requirements.
type InputSettings struct {
	// The web address of the file that Mux should download and use. * For subtitles text tracks, the url is the location of subtitle/captions file. Mux supports [SubRip Text (SRT)](https://en.wikipedia.org/wiki/SubRip) and [Web Video Text Tracks](https://www.w3.org/TR/webvtt1/) format for ingesting Subtitles and Closed Captions. * For Watermarking or Overlay, the url is the location of the watermark image. * When creating clips from existing Mux assets, the url is defined with `mux://assets/{asset_id}` template where `asset_id` is the Asset Identifier for creating the clip from.
	Url             string                       `json:"url,omitempty"`
	OverlaySettings InputSettingsOverlaySettings `json:"overlay_settings,omitempty"`
	// The time offset in seconds from the beginning of the video indicating the clip's starting marker. The default value is 0 when not included. This parameter is only applicable for creating clips when `input.url` has `mux://assets/{asset_id}` format.
	StartTime float64 `json:"start_time,omitempty"`
	// The time offset in seconds from the beginning of the video, indicating the clip's ending marker. The default value is the duration of the video when not included. This parameter is only applicable for creating clips when `input.url` has `mux://assets/{asset_id}` format.
	EndTime float64 `json:"end_time,omitempty"`
	// This parameter is required for the `text` track type.
	Type string `json:"type,omitempty"`
	// Type of text track. This parameter only supports subtitles value. For more information on Subtitles / Closed Captions, [see this blog post](https://mux.com/blog/subtitles-captions-webvtt-hls-and-those-magic-flags/). This parameter is required for `text` track type.
	TextType string `json:"text_type,omitempty"`
	// The language code value must be a valid [BCP 47](https://tools.ietf.org/html/bcp47) specification compliant value. For example, en for English or en-US for the US version of English. This parameter is required for text type and subtitles text type track.
	LanguageCode string `json:"language_code,omitempty"`
	// The name of the track containing a human-readable description. This value must be unique across all text type and subtitles `text` type tracks. The hls manifest will associate a subtitle text track with this value. For example, the value should be \"English\" for subtitles text track with language_code as en. This optional parameter should be used only for `text` type and subtitles `text` type track. If this parameter is not included, Mux will auto-populate based on the `input[].language_code` value.
	Name string `json:"name,omitempty"`
	// Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH). This optional parameter should be used for `text` type and subtitles `text` type tracks.
	ClosedCaptions bool `json:"closed_captions,omitempty"`
	// This optional parameter should be used for `text` type and subtitles `text` type tracks.
	Passthrough string `json:"passthrough,omitempty"`
}
