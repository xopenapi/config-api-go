/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// CursorConfigsResponseData struct for CursorConfigsResponseData
type CursorConfigsResponseData struct {
	Extra CursorExtra `json:"extra,omitempty"`
	Items []Config `json:"items,omitempty"`
}
