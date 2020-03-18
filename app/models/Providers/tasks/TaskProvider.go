package tasks

import (
	"RBStask/app/models/entity"
	"RBStask/app/models/mappers"
	"database/sql"
	"fmt"

)


type TaskProvider struct {
	db            *sql.DB
	tasks         *mappers.TaskMapper

}


func (p *TaskProvider) Init() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=RBStask sslmode=disable")
	
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	} 
	

	p.db = db

	p.tasks = new(mappers.TaskMapper)
	err = p.tasks.Init(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	}
	return nil
}


func (p *TaskProvider) Delete(receivedTask *entity.Task) error {
	defer p.db.Close()
	
	err := p.tasks.DeleteProjectId(receivedTask)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}

	return nil
}


func (p *TaskProvider) GetTasks(ProjectId int64) ([]entity.Task, error) {
		defer p.db.Close()
		return p.tasks.SelectAll(ProjectId)
	}
	
	
func (p *TaskProvider) Add(receivedTask *entity.Task) error {
	defer p.db.Close()

	err := p.tasks.Add(receivedTask)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}


func (p *TaskProvider) Edit(receivedTask *entity.Task) error {
	defer p.db.Close()

	err := p.tasks.Edit(receivedTask)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



