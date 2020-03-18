package controllers

import (
	"github.com/revel/revel"
	//	"log"
	//  "crypto/rand"
	"fmt"
	
)

type App struct {
	*revel.Controller
}


func (c *App) Index() revel.Result {
	
	key, err := c.Session.Get("KEY")
	if err != nil {
		fmt.Println(key)
		return c.Redirect((*App).Login)
	}
	return c.Render()
}


func (c *App) Login() revel.Result {
	return c.Render()
}


// func (c *App) SetCook() revel.Result {
// 	// c.Session["KEY"] = "KEY"
// 	fmt.Println("Ещё вывелся в setCook");
// 	c.Session["KEY"] = "124124124"
// 	// x:= randToken()
//     // new_cookie := http.Cookie{Name: "KEY", Value: x}
// 	// c.SetCookie(&new_cookie)

//   return c.Redirect((*App).Index)
// }


// func randToken() string {
// 	b := make([]byte, 8)
// 	rand.Read(b)
// 	return fmt.Sprintf("%x", b)
// }
