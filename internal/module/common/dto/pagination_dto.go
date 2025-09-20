// internal/dto/pagination_dto.go
package dto

// PaginationRequest 分页请求参数（适用于所有需要分页的列表接口）
type PaginationRequest struct {
	Page     int `json:"page" binding:"required,min=1"`              // 页码（从1开始）
	PageSize int `json:"page_size" binding:"required,min=1,max=100"` // 每页条数（限制最大值防止查询过大）
}

// PaginationResponse 分页响应结果（用于统一返回分页数据）
type PaginationResponse struct {
	Total    int64       `json:"total"`     // 总记录数
	Page     int         `json:"page"`      // 当前页码
	PageSize int         `json:"page_size"` // 每页条数
	Data     interface{} `json:"data"`      // 分页数据（具体业务数据）
}
