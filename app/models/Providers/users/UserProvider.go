package users

import (
	"RBStask/app/models/entity"
	"RBStask/app/models/mappers"
	"database/sql"
	"fmt"
	//  "RBStask/app/controllers"


	// _ "github.com/lib/pq"
)


type UserProvider struct {
	db            *sql.DB
	users         *mappers.UserMapper
}

func (p *UserProvider) Init() error {
	connStr := "host=localhost port=5432 user=postgres password=123 dbname=RBStask sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		return fmt.Errorf("ошибка инициализации модели : %v", err)
	} 
	p.db = db
	p.users = new(mappers.UserMapper)
	err = p.users.Init(p.db)
	if err != nil {
		return fmt.Errorf("ошибка инициализации : %v", err)
	}
	return nil
}


func (p *UserProvider) Login(receivedUser *entity.User) bool {
	defer p.db.Close()

	err :=  p.users.Login(receivedUser)
	if err == true {
		fmt.Println("Получили то что хотели true: %s", err)
		return true
	}else {
		fmt.Println("Получили то что хотели false: %s", err)
		return false
	}

}

// func (p *UserProvider) SetCook() error {
// 	x := randToken()
//     new_cookie := &http.Cookie{Name: "fooddd", Value: "randToken()"}
// 	p.SetCookie(new_cookie)
	
// }


