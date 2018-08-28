/*
 * voiceprintstorage
 *
 * api document
 *
 * API version: /cloudapi/v1beta
 * Contact: xiachengjia@speakin.mobi
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VoiceprinttopnVerifyRequest struct {
	AppName string `json:"appName"`
	Url string `json:"url"`
	SamplingRate string `json:"samplingRate"`
	Topn int64 `json:"topn"`
	Timestamp int64 `json:"timestamp"`
}