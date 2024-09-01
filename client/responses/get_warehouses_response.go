package responses

type GetWarehousesResponse = Response[Warehouse]

type Warehouse struct {
	WarehouseID   string `json:"warehouse_id"`
	WarehouseKey  string `json:"warehouse_key"`
	WarehouseName string `json:"warehouse_name"`
}
