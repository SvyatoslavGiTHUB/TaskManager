package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type RemTaskMapper struct {
	db *sql.DB
}

func (m *RemTaskMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

func (m *RemTaskMapper) RemoveTasks(projectId *entity.Project) error {
	sqlSelect := `DELETE FROM public.tasks WHERE 
	id_project = $1;`
	_, err := m.db.Exec(sqlSelect, projectId.Id)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



