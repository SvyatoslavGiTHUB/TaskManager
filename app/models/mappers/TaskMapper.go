package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type TaskMapper struct {
	db *sql.DB
}

func (m *TaskMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}



func (m *TaskMapper) Add(receivedTask *entity.Task) error {
	fmt.Println(receivedTask)
	sqlSelect := `
	INSERT INTO public."tasks"(name, description, prioritu, status, type, id_project, time)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

_, err := m.db.Exec(sqlSelect, receivedTask.Name, receivedTask.Description, receivedTask.Priority, receivedTask.Status, receivedTask.TypeTask, receivedTask.IdProject, receivedTask.Time)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}




func (m *TaskMapper) DeleteProjectId(taskId *entity.Task) error {
	sqlSelect := `DELETE FROM public.tasks WHERE id = $1 `
	_, err := m.db.Exec(sqlSelect, taskId.Id)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}


func (m *TaskMapper) Edit(receivedTask *entity.Task) error {
	fmt.Println(receivedTask)
	sqlSelect := `
	UPDATE tasks SET 
	name = $2,
	id_project = $3, 
	description = $4, 
	status = $5, 
	type = $6,
	prioritu = $7,
	time = $8

	WHERE id = $1`
	_, err := m.db.Exec(sqlSelect, receivedTask.Id, receivedTask.Name, receivedTask.IdProject, receivedTask.Description, receivedTask.Status, receivedTask.TypeTask, receivedTask.Priority, receivedTask.Time)

	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}


func (m *TaskMapper) SelectAll(ProjectId int64) ([]entity.Task, error) {
	var (
		dbId            sql.NullInt64
		dbName          sql.NullString
		dbIdProject		sql.NullInt64
		dbDescription   sql.NullString
		dbTime          sql.NullInt64 
		dbPriority	    sql.NullString
		dbStatus        sql.NullString
		dbTypeTask      sql.NullString
	)
	sqlSelect := `SELECT * FROM public.tasks
	WHERE id_project = $1`
	lines, err := m.db.Query(sqlSelect, ProjectId)
	if err != nil {
		fmt.Println("ERR")
		return nil, err
	}
	Tasks := make([]entity.Task, 0)
	for lines.Next() {
		 err = lines.Scan(&dbId, &dbName, &dbIdProject, &dbDescription, &dbTime, &dbPriority, &dbStatus, &dbTypeTask)
		// err = lines.Scan(&dbId, &dbName, &dbDescription, &dbPriority, &dbStatus, &dbTypeTask, &dbIdProject, &dbTime)
		if err != nil {
			return nil, err
		}

		currentTask := entity.Task{
			Id:             dbId.Int64,
			Name:           dbName.String,
			IdProject:      dbIdProject.Int64,
			Description:   	dbDescription.String,
			Time :        	dbTime.Int64,
			Priority:		dbPriority.String,
			Status:      	dbStatus.String, 
			TypeTask:   	dbTypeTask.String,
		}

		Tasks = append(Tasks, currentTask)
	}
	return Tasks, nil
}


