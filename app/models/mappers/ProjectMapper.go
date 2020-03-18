package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type ProjectMapper struct {
	db *sql.DB
}

func (m *ProjectMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

func (m *ProjectMapper) Add(receivedProject *entity.Project) error {

	fmt.Println(receivedProject)

	sqlSelect := `
	INSERT INTO public."projects"(name, id_group) 
	VALUES ($1, $2);`

	_, err := m.db.Exec(sqlSelect, receivedProject.Name, receivedProject.IdGroup)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}


func (m *ProjectMapper) Edit(receivedProject *entity.Project) error {
	fmt.Println(receivedProject)
	sqlSelect := `
	UPDATE projects SET 
	name = $2, 
	id_group = $3  
	
	WHERE id = $1`

	_, err := m.db.Exec(sqlSelect, receivedProject.Id, receivedProject.Name, receivedProject.IdGroup)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



func (m *ProjectMapper) DeleteProjectId(projectId *entity.Project) error {
	sqlSelect := `DELETE FROM public.projects WHERE id = $1;`
	_, err := m.db.Exec(sqlSelect, projectId.Id)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}

func (m *ProjectMapper) SelectAll() ([]entity.Project, error) {
	var (
		dbId        sql.NullInt64
		dbName      sql.NullString
		dbIdGroup   sql.NullInt64
	)

	sqlSelect := `SELECT * FROM projects;`

	rows, err := m.db.Query(sqlSelect)
	if err != nil {
		fmt.Println("ERR")
		return nil, err
	}


	Projects := make([]entity.Project, 0)
	for rows.Next() {
		err = rows.Scan(&dbId, &dbName, &dbIdGroup)
		if err != nil {
			return nil, err
		}

		currentProject := entity.Project{
			Id:             dbId.Int64,
			Name:           dbName.String,
			IdGroup:      dbIdGroup.Int64,
		}

		Projects = append(Projects, currentProject)
	}
	return Projects, nil
}


