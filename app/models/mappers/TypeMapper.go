package mappers

import (
	"RBStask/app/models/entity"
	"database/sql"
	"fmt"
)

type TypeMapper struct {
	db *sql.DB
}

func (m *TypeMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}



func (m *TypeMapper) SelectAll() ([]entity.Type, error) {
	var (
		dbId        sql.NullInt64
		dbName      sql.NullString
	)

	sqlSelect := `SELECT * FROM types;`

	rows, err := m.db.Query(sqlSelect)
	if err != nil {
		fmt.Println("ERR")
		return nil, err
	}

	Types := make([]entity.Type, 0)
	for rows.Next() {
		err = rows.Scan(&dbId, &dbName)
		if err != nil {
			return nil, err
		}

		currentType := entity.Type{
			Id:             dbId.Int64,
			Name:           dbName.String,
		}

		Types = append(Types, currentType)
	}
	return Types, nil
}

