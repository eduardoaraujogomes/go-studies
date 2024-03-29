package inbound_orders

import (
	"database/sql"
	"fmt"

	"digital-house/pkg/util"
)

type Repository interface {
	GetAll() ([]ReportInboundOrders, error)
	GetByID(id int) (ReportInboundOrders, error)
	Create(createdAt string, inboundOrders InboundOrdersCreate) (InboundOrdersResponse, error)
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll() ([]ReportInboundOrders, error) {
	var reportInboundOrders []ReportInboundOrders

	rows, err := r.db.Query(GET_ALL_REPORT_INBOUND_ORDERS)
	if err != nil {
		return reportInboundOrders, err
	}
	defer rows.Close()

	for rows.Next() {
		var reportInboundOrder ReportInboundOrders

		err := rows.Scan(
			&reportInboundOrder.ID,
			&reportInboundOrder.CardNumberID,
			&reportInboundOrder.FirstName,
			&reportInboundOrder.LastName,
			&reportInboundOrder.WarehouseID,
			&reportInboundOrder.InboundOrdersCount,
		)
		if err != nil {
			return reportInboundOrders, err
		}
		reportInboundOrders = append(reportInboundOrders, reportInboundOrder)
	}
	return reportInboundOrders, nil
}

func (r *repository) GetByID(id int) (ReportInboundOrders, error) {
	row := r.db.QueryRow(GET_REPORT_INBOUND_ORDER_BY_ID, id)
	var reportInboundOrder ReportInboundOrders

	if err := row.Scan(
		&reportInboundOrder.ID,
		&reportInboundOrder.CardNumberID,
		&reportInboundOrder.FirstName,
		&reportInboundOrder.LastName,
		&reportInboundOrder.WarehouseID,
		&reportInboundOrder.InboundOrdersCount,
	); err != nil {
		return ReportInboundOrders{}, fmt.Errorf(EMPLOYEE_NOT_FOUND)
	}

	return reportInboundOrder, nil
}

func (r *repository) Create(createdAt string, object InboundOrdersCreate) (InboundOrdersResponse, error) {
	inboundOrder := InboundOrdersResponse{OrderDate: createdAt, OrderNumber: object.OrderNumber, DentistID: object.DentistID, ProductBatchID: object.ProductBatchID, WarehouseID: object.WarehouseID}

	stmt, err := r.db.Exec(CREATE_INBOUND_ORDERS,
		&inboundOrder.OrderDate,
		&inboundOrder.OrderNumber,
		&inboundOrder.DentistID,
		&inboundOrder.ProductBatchID,
		&inboundOrder.WarehouseID,
	)

	if err != nil {
		return InboundOrdersResponse{}, util.CheckError(err)
	}

	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return InboundOrdersResponse{}, fmt.Errorf(FAIL_TO_SAVE)
	}
	return inboundOrder, nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
