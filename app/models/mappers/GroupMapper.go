package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type GroupMapper struct {
	db *sql.DB
}

func (m *GroupMapper) Connect(db *sql.DB) error {
	m.db = db
	return nil
}




func (m *GroupMapper) Add(currentGroup *entity.Group) error {
	fmt.Println(currentGroup)
	sqlSelect := `
	INSERT INTO public."groups"(name) 
	VALUES ($1);`
	_, err := m.db.Exec(sqlSelect, currentGroup.Name)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



func (m *GroupMapper) Delete(groupId *entity.Group) error {
	sqlSelect := `
	DELETE FROM public.groups WHERE id = $1;
	`
	_, err := m.db.Exec(sqlSelect, groupId.Id)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}

func (m *GroupMapper) GetAll() ([]entity.Group, error) {
	var (
		dbId        sql.NullInt64
		dbName      sql.NullString
	)

	sqlSelect := `SELECT * FROM groups;`

	rows, err := m.db.Query(sqlSelect)
	if err != nil {
		fmt.Println("ERR")
		return nil, err
	}


	Groups := make([]entity.Group, 0)
	for rows.Next() {
		err = rows.Scan(&dbId, &dbName)
		if err != nil {
			return nil, err
		}

		currentGroup := entity.Group{
			Id:             dbId.Int64,
			Name:           dbName.String,
		}
		Groups = append(Groups, currentGroup)
	}
	return Groups, nil
}



func (m *GroupMapper) Edit(receivedGroup *entity.Group) error {
	fmt.Println(receivedGroup)
	sqlSelect := `
	UPDATE groups SET 
	name = $2
	WHERE id = $1`
	_, err := m.db.Exec(sqlSelect, receivedGroup.Id, receivedGroup.Name,)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}

