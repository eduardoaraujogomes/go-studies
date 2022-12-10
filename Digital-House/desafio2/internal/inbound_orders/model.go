package inbound_orders

type InboundOrders struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	DentistID      int    `json:"dentist_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WarehouseID    int    `json:"warehouse_id"`
}

type InboundOrdersCreate struct {
	OrderNumber    string `json:"order_number" binding:"required"`
	DentistID      int    `json:"dentist_id" binding:"required"`
	ProductBatchID int    `json:"product_batch_id" binding:"required"`
	WarehouseID    int    `json:"warehouse_id" binding:"required"`
}

type InboundOrdersResponse struct {
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	DentistID      int    `json:"dentist_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WarehouseID    int    `json:"warehouse_id"`
}

type ReportInboundOrders struct {
	ID                 int    `json:"id"`
	CardNumberID       string `json:"card_number_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	WarehouseID        int    `json:"warehouse_id"`
	InboundOrdersCount int    `json:"inbound_orders_count"`
}
