/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// CursorConfigTemplatesRsp struct for CursorConfigTemplatesRsp
type CursorConfigTemplatesRsp struct {
	Code int64 `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data CursorConfigTemplatesRspData `json:"data,omitempty"`
}