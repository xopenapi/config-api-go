/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// UpdateConfigTemplateRsp struct for UpdateConfigTemplateRsp
type UpdateConfigTemplateRsp struct {
	Code int64 `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data ConfigTemplate `json:"data,omitempty"`
}
