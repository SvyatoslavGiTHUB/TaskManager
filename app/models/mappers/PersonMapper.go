package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type PersonMapper struct {
	db *sql.DB
}

func (m *PersonMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

func (m *PersonMapper) Add(receivedPerson *entity.Person) error {

fmt.Println(receivedPerson)
sqlSelect := `
INSERT INTO public."persons"(name, surname, old, id_group, position) 
VALUES ($1, $2, $3, $4, $5);`
_, err := m.db.Exec(sqlSelect, receivedPerson.Name, receivedPerson.Surname, receivedPerson.Old, receivedPerson.IdGroup, receivedPerson.Position)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}






func (m *PersonMapper) GetPerson(GroupId int64) ([]entity.Person, error) {
	var (
		dbId            sql.NullInt64
		dbName          sql.NullString
		dbSurname		sql.NullString
		dbOld           sql.NullInt64 
		dbIdGroup       sql.NullInt64 
		dbPosition		sql.NullString


	)

	sqlSelect := `SELECT * FROM public.persons
	WHERE id_group = $1`
	lines, err := m.db.Query(sqlSelect, GroupId)
	if err != nil {
		fmt.Println("ERR")
		return nil, err
	}


	Persons := make([]entity.Person, 0)
	for lines.Next() {
err = lines.Scan(&dbId, &dbName, &dbSurname, &dbOld, &dbIdGroup, &dbPosition)
		if err != nil {
			return nil, err
		}

		currentPerson := entity.Person{
			Id:             dbId.Int64,
			Name:           dbName.String,
			Surname:        dbSurname.String,
		    Old:   	        dbOld.Int64,
			IdGroup :       dbIdGroup.Int64,
			Position:       dbPosition.String,

		}

		Persons = append(Persons, currentPerson)
	}
	return Persons, nil
}


func (m *PersonMapper) Edit(receivedPerson *entity.Person) error {
	fmt.Println(receivedPerson)
	
	sqlSelect := `
	UPDATE persons SET 
	name = $2,
	surname = $3, 
	old = $4, 
	id_group = $5, 
	position = $6
	WHERE id = $1`
	_, err := m.db.Exec(sqlSelect, 
		receivedPerson.Id,
		 receivedPerson.Name, 
		 receivedPerson.Surname,
		 receivedPerson.Old,
		 receivedPerson.IdGroup,
		 receivedPerson.Position)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



func (m *PersonMapper) Delete(PersonId *int64) error {
	sqlSelect := `DELETE FROM public.persons WHERE id = $1 `
	_, err := m.db.Exec(sqlSelect, PersonId)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil

}



