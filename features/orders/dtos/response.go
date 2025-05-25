package dtos

type ResOrders struct {
	ID         uint                `json:"id"`
	UserID     uint                `json:"user_id"`
	ServiceID  uint                `json:"service_id"`
	PackageID  uint                `json:"package_id"`
	IsPickup   bool                `json:"is_pickup"`
	Weight     float64             `json:"weight,omitempty"`
	TotalPrice float64             `json:"total_price"`
	Status     string              `json:"status"`
	Address    string              `json:"address"`
	Notes      string              `json:"notes"`
	OrderItems []OrderItemResponse `json:"order_items,omitempty"`
	CreatedAt  string              `json:"created_at"`
}

type OrderItemResponse struct {
	ID                  uint    `json:"id"`
	IndividualPackageID uint    `json:"individual_package_id"`
	Qty                 int     `json:"qty"`
	SubTotal            float64 `json:"sub_total"`
}
