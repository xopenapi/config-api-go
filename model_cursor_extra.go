/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// CursorExtra struct for CursorExtra
type CursorExtra struct {
	// 大小
	Size int64 `json:"size,omitempty"`
	// 查询方向，0表示游标前，1表示游标后
	Direction int32 `json:"direction,omitempty"`
	// 资源ID
	ResourceId string `json:"resourceId,omitempty"`
}
