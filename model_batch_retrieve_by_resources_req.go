/*
 * 配置服务API
 *
 * 配置服务API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package config
// BatchRetrieveByResourcesReq struct for BatchRetrieveByResourcesReq
type BatchRetrieveByResourcesReq struct {
	// key
	Key string `json:"key,omitempty"`
	// 资源ID组
	ResourceIds []string `json:"resourceIds,omitempty"`
}
