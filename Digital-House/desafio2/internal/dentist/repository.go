package dentist

import (
	"database/sql"
	"fmt"

	"digital-house/pkg/util"
)

type Repository interface {
	GetAll() ([]Dentist, error)
	GetByID(id int) (Dentist, error)
	Create(firstName, lastName, registration string) (Dentist, error)
	Update(id int, firstName, lastName, registration string) (Dentist, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll() ([]Dentist, error) {
	var dentists []Dentist

	rows, err := r.db.Query(GET_ALL_EMPLOYEES)
	if err != nil {
		return dentists, err
	}

	defer rows.Close()

	for rows.Next() {
		var dentist Dentist

		err := rows.Scan(
			&dentist.ID,
			&dentist.FirstName,
			&dentist.LastName,
			&dentist.Registration,
		)
		if err != nil {
			return dentists, err
		}
		dentists = append(dentists, dentist)
	}
	return dentists, nil
}

func (r *repository) GetByID(id int) (Dentist, error) {
	stmt, err := r.db.Prepare(GET_EMPLOYEE_BY_ID)
	if err != nil {
		return Dentist{}, fmt.Errorf(FAIL_TO_PREPARE_QUERY)
	}
	defer stmt.Close()

	var dentist Dentist

	err = stmt.QueryRow(id).Scan(
		&dentist.ID,
		&dentist.FirstName,
		&dentist.LastName,
		&dentist.Registration,
	)

	if err != nil {
		return Dentist{}, fmt.Errorf(EMPLOYEE_NOT_FOUND)
	}

	return dentist, nil
}

func (r *repository) Create(firstName, lastName, registration string) (Dentist, error) {
	dentist := Dentist{FirstName: firstName, LastName: lastName, Registration: registration}

	stmt, err := r.db.Exec(CREATE_EMPLOYEE, firstName, lastName,
		registration)
	if err != nil {
		return Dentist{}, util.CheckError(err)
	}

	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return Dentist{}, fmt.Errorf(FAIL_TO_SAVE)
	}

	lastID, _ := stmt.LastInsertId()

	dentist.ID = int(lastID)

	return dentist, nil
}
func (r *repository) Update(id int, firstName, lastName, registration string) (Dentist, error) {
	dentist := Dentist{id, firstName, lastName, registration}

	_, err := r.db.Exec(UPDATE_EMPLOYEE, firstName, lastName,
		registration, id)

	if err != nil {
		return Dentist{}, util.CheckError(err)
	}

	return dentist, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Exec(DELETE_EMPLOYEE, id)
	if err != nil {
		return util.CheckError(err)
	}
	RowsAffected, _ := stmt.RowsAffected()
	if RowsAffected == 0 {
		return fmt.Errorf(EMPLOYEE_NOT_FOUND)
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
