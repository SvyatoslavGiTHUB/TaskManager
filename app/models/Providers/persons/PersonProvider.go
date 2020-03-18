package persons

import (
	"RBStask/app/models/entity"
	"RBStask/app/models/mappers"
	"database/sql"
	"fmt"
	"log"
	"crypto/rand"
	_ "github.com/lib/pq"
)


type PersonProvider struct {
	db            *sql.DB
	persons         *mappers.PersonMapper
	registUser      *mappers.RegistUserMapper

}




func (p *PersonProvider) Init() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=RBStask sslmode=disable")
	if err != nil {
		return fmt.Errorf("ERR: %v", err)
	} 
	p.db = db

	p.persons = new(mappers.PersonMapper)
	err = p.persons.Init(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %v", err)
	}

	p.registUser = new(mappers.RegistUserMapper)
	err = p.registUser.Connect(p.db)
	if err != nil {
		return fmt.Errorf("ERR: %v", err)
	}
	return nil


}



func (p *PersonProvider) Add(receivedPerson *entity.Person) (string, error)  {
	defer p.db.Close()

	err := p.persons.Add(receivedPerson)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return  "Ошибка в err := p.persons.Add(receivedPerson )%s", err
	}
	 passwordRAND := randPassword() 

	err = p.registUser.Regist(passwordRAND, receivedPerson)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return "Ошибка в err = p.registUser.Regist(passwordRAND, receivedPerson)%s", err
	}

	return passwordRAND, err
}

	func randPassword() string {
		b := make([]byte, 8)
		rand.Read(b)
		return fmt.Sprintf("%x", b)
	}


func (p *PersonProvider) Edit(receivedPerson *entity.Person) error {
	defer p.db.Close()
	
	err := p.persons.Edit(receivedPerson)
	if err != nil {
		fmt.Printf("ERR: %s", err)
		return err
	}
	return nil
}

func (p *PersonProvider) Delete(PersonId *int64) error {
	defer p.db.Close()

	err := p.persons.Delete(PersonId)
	if err != nil {
		log.Fatalf("ERR: %s", err)
		return err
	}
	return nil
}


func (p *PersonProvider) GetPerson(GroupId int64) ([]entity.Person, error) {
	defer p.db.Close()
	return p.persons.GetPerson(GroupId)
}


