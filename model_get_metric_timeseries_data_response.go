// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxvideosdk

type GetMetricTimeseriesDataResponse struct {
	Data [][]string `json:"data,omitempty"`
	TotalRowCount int32 `json:"total_row_count,omitempty"`
	Timeframe []string `json:"timeframe,omitempty"`
}
