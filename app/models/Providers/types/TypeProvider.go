package types

import (
	"RBStask/app/models/entity"
	"RBStask/app/models/mappers"
	"database/sql"
	"fmt"

)


type TypeProvider struct {
	db            *sql.DB
	types         *mappers.TypeMapper
}

func (p *TypeProvider) Init() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=RBStask sslmode=disable")
	
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	} 

	p.db = db

	p.types = new(mappers.TypeMapper)
	err = p.types.Init(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %s", err)
	}
	return nil
}






func (p *TypeProvider) GetAllType() ([]entity.Type, error) {
	defer p.db.Close()
	return p.types.SelectAll()
}






