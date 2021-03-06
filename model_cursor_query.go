/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// CursorQuery struct for CursorQuery
type CursorQuery struct {
	// 查询条件，传键值对
	Filters map[string]interface{} `json:"filters,omitempty"`
	// 游标值
	Cursor string `json:"cursor,omitempty"`
	// 排序 $ref:\"#/components/schemas/SortSpec\"
	CursorSort *interface{} `json:"cursorSort,omitempty"`
	// 大小
	Size int64 `json:"size,omitempty"`
	// 查询方向，0表示游标前，1表示游标后
	Direction int32 `json:"direction,omitempty"`
}
