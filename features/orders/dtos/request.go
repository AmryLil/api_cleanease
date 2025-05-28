package dtos

type InputOrders struct {
	UserID     uint                   `json:"user_id" `
	ServiceID  uint                   `json:"service_id" binding:"required"`
	PackageID  uint                   `json:"package_id" binding:"required"`
	IsPickup   bool                   `json:"is_pickup"`
	Weight     float64                `json:"weight"` // Optional jika bukan per kg
	Address    string                 `json:"address"`
	Notes      string                 `json:"notes"`
	OrderItems []OrderItemCreateInput `json:"order_items"`
}

type OrderItemCreateInput struct {
	IndividualPackageID uint `json:"individual_package_id" binding:"required"`
	Qty                 int  `json:"qty" binding:"required,min=1"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
