package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type RemGroupMapper struct {
	db *sql.DB
}

func (m *RemGroupMapper) Connect(db *sql.DB) error {
	m.db = db
	return nil
}



func (m *RemGroupMapper) RemoveGroupId(groupId entity.Group) error {
	sqlSelect := `DELETE FROM public.persons WHERE id_group = $1;`
	_, err := m.db.Exec(sqlSelect, groupId.Id)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



