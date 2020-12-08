/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// ConfigTemplate struct for ConfigTemplate
type ConfigTemplate struct {
	// 配置ID
	Id string `json:"id,omitempty"`
	// 组ID
	GroupId string `json:"groupId,omitempty"`
	// 模版ID
	TemplateId string `json:"templateId,omitempty"`
	// 对应的字段名
	Key string `json:"key,omitempty"`
	// json格式的键值对数据
	Value map[string]interface{} `json:"value,omitempty"`
	// 创建时间，返回时间戳
	Ctime int64 `json:"ctime,omitempty"`
	// 更新时间，返回时间戳
	Mtime int64 `json:"mtime,omitempty"`
}
