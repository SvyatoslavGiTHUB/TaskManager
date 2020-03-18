package mappers

import (
	 "RBStask/app/models/entity"
	"database/sql"
	 "fmt"
	//   "log"
)

type UserMapper struct {
	db *sql.DB
}

func (m *UserMapper) Init(db *sql.DB) error {
	m.db = db
	return nil
}

func (m *UserMapper) Login(receivedUser *entity.User) (bool)  {
	var (
		
			dbLogin             sql.NullString
			dbPassword          sql.NullString
			dbId                sql.NullInt64 
	)

	sqlSelect := `SELECT * FROM public.user
	WHERE 
	login = $1 and
	password = $2`
	lines, err := m.db.Query(sqlSelect, receivedUser.Login, receivedUser.Password)
	Tasks := make([]entity.User, 0)
	for lines.Next() {
		err = lines.Scan(&dbId, &dbLogin, &dbPassword)
		currentTask := entity.User{
		Login:              dbLogin.String,
		Password:           dbPassword.String,
		}

		fmt.Println(err)
		 Tasks = append(Tasks, currentTask)
		 return true
	}
	
	return false 
}