package controllers

import (
	"RBStask/app/models/Providers/persons"
	"RBStask/app/models/entity"
	"RBStask/app/output"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"io/ioutil"

	"github.com/revel/revel"
)

type CPerson struct {
	*revel.Controller
	provider *persons.PersonProvider
}

func (c *CPerson) Delete() revel.Result {
	c.provider = new(persons.PersonProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	delrequestPerson := c.Params.Get("person")
	PersonId, err := strconv.ParseInt(delrequestPerson, 10, 64)
	if err != nil {
		fmt.Printf("Ошибка чтения ParseInt: %s", err)
	}
	c.provider.Delete(&PersonId);
	if err != nil {
		log.Fatalf("Ошибка передачи в provederpersons: %s", err)
	}
	return nil
}

func (c *CPerson) Add() revel.Result {
	c.provider = new(persons.PersonProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	receivedPerson := entity.Person{}
	
	getdann, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}
	err = json.Unmarshal(getdann, &receivedPerson)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}
	var password string  
	password, err = c.provider.Add(&receivedPerson) 
	if err != nil {
		fmt.Printf("Ошибка Add Provider: %s", err)
		return nil
	}

	return c.RenderJSON(output.Correct(password))
}



func (c *CPerson) GetAll() revel.Result {
	c.provider = new(persons.PersonProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	getrequestPerson := c.Params.Get("persons")
	GroupId, err := strconv.ParseInt(getrequestPerson, 10, 64)
	if err != nil {
		fmt.Printf("Ошибка чтения ParseInt: %s", err)
	}
	persons, err := c.provider.GetPerson(GroupId);
	if err != nil {
		log.Fatalf("Ошибка передачи в provederPerson: %s", err)
	}
	return c.RenderJSON(output.Correct(persons))
}



func (c *CPerson) Edit() revel.Result {
	c.provider = new(persons.PersonProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	receivedPerson := entity.Person{}
	currentItem, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(currentItem, &receivedPerson)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	err = c.provider.Edit(&receivedPerson)
	if err != nil {
		log.Fatalf("Ошибка передачи в Edit proveder: %s", err)
		return nil
	}

	return nil
}



