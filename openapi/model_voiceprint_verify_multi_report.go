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

type VoiceprintVerifyMultiReport struct {
	// 唯一id
	UnionID string `json:"unionID,omitempty"`
	// 分数
	Score float64 `json:"score,omitempty"`
}
