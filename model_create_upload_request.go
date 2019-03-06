// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxvideosdk

type CreateUploadRequest struct {
	// Max time in seconds for the signed upload URL to be valid. If a successful upload has not occurred before the timeout limit, the direct upload is marked `timed_out`
	Timeout int32 `json:"timeout,omitempty"`
	// If the upload URL will be used in a browser, you must specify the origin in order for the signed URL to have the correct CORS headers.
	CorsOrigin string `json:"cors_origin,omitempty"`
	NewAssetSettings CreateAssetRequest `json:"new_asset_settings"`
}
