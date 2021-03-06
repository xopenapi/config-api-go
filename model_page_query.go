/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// PageQuery struct for PageQuery
type PageQuery struct {
	// 查询条件，传键值对
	Filters map[string]interface{} `json:"filters,omitempty"`
	// 页码
	PageNo int32 `json:"pageNo,omitempty"`
	// 每页数据量
	PageSize int32 `json:"pageSize,omitempty"`
	// 排序
	Sort []SortSpec `json:"sort,omitempty"`
}
