/*
 * 声纹云api
 *
 * api document
 *
 * API version: /cloudapi/v1beta
 * Contact: xiachengjia@speakin.mobi
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VoiceprintClipsResponse struct {
	// 人声
	HumanSlice []VoiceprintClipsSlice `json:"humanSlice,omitempty"`
	// 播报
	BroadcastSlice []VoiceprintClipsSlice `json:"broadcastSlice,omitempty"`
	// 音乐
	MusicSlice []VoiceprintClipsSlice `json:"musicSlice,omitempty"`
	// 忙音
	BusySignalSlice []VoiceprintClipsSlice `json:"busySignalSlice,omitempty"`
	// 铃声
	RingingSlice []VoiceprintClipsSlice `json:"ringingSlice,omitempty"`
	// 静音
	SilentSlice []VoiceprintClipsSlice `json:"silentSlice,omitempty"`
}
