package controllers

import (
	"RBStask/app/models/Providers/types"
	"RBStask/app/output"
	"log"
	"github.com/revel/revel"
)

type CType struct {
	*revel.Controller
	provider *types.TypeProvider
}


func (c *CType) GetAll() revel.Result {
	c.provider = new(types.TypeProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)

	}


	types, err := c.provider.GetAllType()
	if err != nil {
	log.Fatalf("Ошибка вывода GetAll Projects: %s", err)
	}
	return c.RenderJSON(output.Correct(types))
}

