package controllers

import (
	"RBStask/app/models/Providers/users"
	"RBStask/app/models/entity"
	"RBStask/app/output"
    "encoding/json"
	"fmt"
    "log"
	"io/ioutil"
	"github.com/revel/revel"
)

type CUser struct {
	*revel.Controller
	provider *users.UserProvider

}


func (c *CUser) GetUser() revel.Result {
	c.provider = new(users.UserProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	receivedUser := entity.User{}
	
	readJson, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(readJson, &receivedUser)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	res := c.provider.Login(&receivedUser)
	if res == true {
		c.Session["KEY"] = "KEY"
		// fmt.Println(*App)
		// c.Redirect("")
		// 6e3e6da699ff08ad
		// return c.RenderJSON(output.Correct("ВЫ вошли"))
		// p.Redirect("/")
		return c.RenderJSON(output.Correct("true"))

	} else {
		return c.RenderJSON(output.Correct("Неверный логин или пароль!"))
	}


}



func (c *CUser) KillSession() revel.Result{
	c.Session.Del("KEY")
return c.RenderJSON(output.Correct("true"))


}