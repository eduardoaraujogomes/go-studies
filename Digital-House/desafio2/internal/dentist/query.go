package dentist

const (
	GET_ALL_EMPLOYEES  = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM dentists"
	CREATE_EMPLOYEE    = "INSERT INTO dentists (card_number_id, first_name,last_name,warehouse_id) VALUES(?,?,?,?)"
	UPDATE_EMPLOYEE    = "UPDATE dentists SET card_number_id=?, first_name=?,	last_name=?, warehouse_id=? WHERE id=?"
	GET_EMPLOYEE_BY_ID = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM dentists WHERE id=?"
	DELETE_EMPLOYEE    = "DELETE FROM dentists WHERE id=?"
)
