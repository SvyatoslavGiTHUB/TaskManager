package groups

import (
	"RBStask/app/models/entity"
	"RBStask/app/models/mappers"
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

type GroupProvider struct {
	db            *sql.DB
	groups         *mappers.GroupMapper
	remgroups      *mappers.RemGroupMapper

}


func (p *GroupProvider) Connect() error {
	
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=RBStask sslmode=disable")
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	} 

	p.db = db
	
	p.groups = new(mappers.GroupMapper)
	err = p.groups.Connect(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	}


	p.remgroups = new(mappers.RemGroupMapper)
	err = p.remgroups.Connect(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	}
	
	return nil
}


func (p *GroupProvider) Delete(currentGroup *entity.Group) error {
	defer p.db.Close()

	err := p.groups.Delete(currentGroup)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}

	
	Err :=p.remgroups.RemoveGroupId(*currentGroup)
	if Err != nil {
		fmt.Printf("ERR: %s", Err)
	}

	return nil
}


func (p *GroupProvider) GetAll() ([]entity.Group, error) {
	defer p.db.Close()
	return p.groups.GetAll()
}


func (p *GroupProvider) Add(currentGroup *entity.Group) error {
	defer p.db.Close()

	err := p.groups.Add(currentGroup)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}



func (p *GroupProvider) Edit(receivedGroup *entity.Group) error {
	defer p.db.Close()

	err := p.groups.Edit(receivedGroup)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}

