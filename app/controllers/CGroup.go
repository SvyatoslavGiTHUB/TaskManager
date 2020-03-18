package controllers

import (
	"RBStask/app/models/Providers/groups"
	"RBStask/app/models/entity"
	"RBStask/app/output"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/revel/revel"
)

// отладку !!
// перевод до конца !
// ВЫВОД ПАРОЛЯ ЮЗЕРА!!!!

type CGroup struct {
	*revel.Controller
	provider *groups.GroupProvider
}

func (c *CGroup) Add() revel.Result {
	c.provider = new(groups.GroupProvider)
	err := c.provider.Connect()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	currentGroup := entity.Group{}

	// c.Params.BindJSON(&currentGroup)
	// fmt.Println(currentGroup);

	result, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(result, &currentGroup)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	err = c.provider.Add(&currentGroup)
	if err != nil {
		fmt.Printf("Ошибка Add Provider: %s", err)
		return nil
	}

	return nil
}

func (c *CGroup) Delete() revel.Result {
	c.provider = new(groups.GroupProvider)
	err := c.provider.Connect()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}


	currentGroup := entity.Group{}

	result, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(result, &currentGroup)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	err = c.provider.Delete(&currentGroup)
	if err != nil {
		fmt.Printf("Ошибка Delete Provider: %s", err)
		return nil
	}


	return nil
}


func (c *CGroup) GetAll() revel.Result {
	c.provider = new(groups.GroupProvider)
	err := c.provider.Connect()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	
	groups, err := c.provider.GetAll()

	if err != nil {
	log.Fatalf("Ошибка вывода GetAll Projects: %s", err)
	}
	return c.RenderJSON(output.Correct(groups))
}




func (c *CGroup) Edit() revel.Result {
	c.provider = new(groups.GroupProvider)
	err := c.provider.Connect()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	receivedGroup := entity.Group{}

	currentItem, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(currentItem, &receivedGroup)
	fmt.Println(receivedGroup)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}
	
	err = c.provider.Edit(&receivedGroup)
	if err != nil {
		fmt.Printf("Ошибка Edit Provider: %s", err)
		return nil
	}

	return nil
}

