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

type VoiceprintDeleteRequest struct {
	AppName string `json:"appName"`
	UnionID string `json:"unionID"`
	Timestamp int64 `json:"timestamp"`
}