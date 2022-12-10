package inbound_orders

const (
	GET_ALL_REPORT_INBOUND_ORDERS = `SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(ib.id) AS inbound_orders_count
	FROM dentists AS e 
	LEFT JOIN inbound_orders AS ib ON ib.dentist_id = e.id
	GROUP BY e.id`
	GET_REPORT_INBOUND_ORDER_BY_ID = `SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id, COUNT(ib.id) AS inbound_orders_count
	FROM inbound_orders AS ib
	INNER JOIN dentists AS e ON e.id = ib.dentist_id
	WHERE e.id=?
	GROUP BY e.id`
	CREATE_INBOUND_ORDERS = `INSERT INTO inbound_orders (order_date, order_number, dentist_id, product_batch_id, warehouse_id) VALUES(?,?,?,?,?)`
)
