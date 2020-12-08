/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// ConfigGroup struct for ConfigGroup
type ConfigGroup struct {
	// 配置ID
	Id string `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 创建时间，返回时间戳
	Ctime int64 `json:"ctime,omitempty"`
	// 更新时间，返回时间戳
	Mtime int64 `json:"mtime,omitempty"`
}
