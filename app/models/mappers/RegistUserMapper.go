package mappers

import (
	 "RBStask/app/models/entity"
	"database/sql"
	 "fmt"
)

type RegistUserMapper struct {
	db *sql.DB
}

func (m *RegistUserMapper) Connect(db *sql.DB) error {
	m.db = db
	return nil
}


func (m *RegistUserMapper) Regist(passwordRAND string, receivedPerson *entity.Person) error {

	sqlSelect := `

	INSERT INTO public."user"(login, password) 
	VALUES ($1, $2);
	`
	_, err := m.db.Exec(sqlSelect, receivedPerson.Name, passwordRAND)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}
