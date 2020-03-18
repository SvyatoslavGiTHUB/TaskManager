package projects

import (
	"RBStask/app/models/entity"
	"RBStask/app/models/mappers"
	"database/sql"
	"fmt"

)

type ProjectProvider struct {
	db            *sql.DB
	projects         *mappers.ProjectMapper
	remtasks         *mappers.RemTaskMapper
}

func (p *ProjectProvider) Init() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=RBStask sslmode=disable")
	
	if err != nil {
		return fmt.Errorf("ERR: %v", err)
	} 

	p.db = db

	p.projects = new(mappers.ProjectMapper)

	err = p.projects.Init(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %v", err)
	}

	p.remtasks = new(mappers.RemTaskMapper)

	err = p.remtasks.Init(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %v", err)
	}

	return nil


}


func (p *ProjectProvider) Delete(receivedProject *entity.Project) error {
	defer p.db.Close()

	err := p.projects.DeleteProjectId(receivedProject)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}

	
	Err :=p.remtasks.RemoveTasks(receivedProject)
	if Err != nil {
		fmt.Printf("ERR: %s", Err)
	}
	
	return nil
}


func (p *ProjectProvider) List() ([]entity.Project, error) {
	defer p.db.Close()
	return p.projects.SelectAll()
}



func (p *ProjectProvider) Add(receivedProject *entity.Project) error {
	defer p.db.Close()

	err := p.projects.Add(receivedProject)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}


func (p *ProjectProvider) Edit(receivedProject *entity.Project) error {
	defer p.db.Close()

	err := p.projects.Edit(receivedProject)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}

