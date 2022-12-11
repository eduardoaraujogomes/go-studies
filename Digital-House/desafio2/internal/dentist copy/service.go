package dentist

type Service interface {
	GetAll() ([]Dentist, error)
	GetByID(id int) (Dentist, error)
	Create(cardNumberID, firstName, lastName string, warehouseID int) (Dentist, error)
	Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Dentist, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]Dentist, error) {
	dentists, err := s.repository.GetAll()
	if err != nil {
		return []Dentist{}, err
	}
	return dentists, nil
}

func (s service) GetByID(id int) (Dentist, error) {
	dentist, err := s.repository.GetByID(id)
	if err != nil {
		return Dentist{}, err
	}
	return dentist, nil
}

func (s service) Create(cardNumberID, firstName, lastName string, warehouseID int) (Dentist, error) {
	dentist, err := s.repository.Create(cardNumberID, firstName, lastName, warehouseID)

	if err != nil {
		return Dentist{}, err
	}

	return dentist, nil
}

func (s service) Update(id int, cardNumberID, firstName, lastName string, warehouseID int) (Dentist, error) {
	dentist, err := s.GetByID(id)
	if err != nil {
		return Dentist{}, err
	}

	if cardNumberID != "" {
		dentist.CardNumberID = cardNumberID
	}
	if firstName != "" {
		dentist.FirstName = firstName
	}
	if lastName != "" {
		dentist.LastName = lastName
	}
	if warehouseID > 0 {
		dentist.WarehouseID = warehouseID
	}

	dentist, err = s.repository.Update(dentist.ID, dentist.CardNumberID, dentist.FirstName, dentist.LastName, dentist.WarehouseID)
	if err != nil {
		return Dentist{}, err
	}

	return dentist, nil
}

func (s service) Delete(id int) error {

	err := s.repository.Delete(id)

	if err != nil {
		return err
	}
	return err
}
