/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// PageConfigsResponse struct for PageConfigsResponse
type PageConfigsResponse struct {
	Code int64 `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data PageConfigsResponseData `json:"data,omitempty"`
}