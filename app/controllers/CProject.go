package controllers

import (
	// "RBStask/app/test"
	"RBStask/app/models/Providers/projects"
	"RBStask/app/models/entity"
	"RBStask/app/output"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"

	"github.com/revel/revel"
)

type CProject struct {
	*revel.Controller
	provider *projects.ProjectProvider
}


func (c *CProject) Add() revel.Result {
	c.provider = new(projects.ProjectProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	 receivedProject := entity.Project{}

	// c.Params.BindJSON(&receivedProject)


	getbody, err := ioutil.ReadAll(c.Request.GetBody())
	
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}
	err = json.Unmarshal(getbody, &receivedProject)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	err = c.provider.Add(&receivedProject)
	if err != nil {
		fmt.Printf("Ошибка Add Provider: %s", err)
		return nil
	}

	return nil
}



func (c *CProject) Edit() revel.Result {
	c.provider = new(projects.ProjectProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	var receivedProject entity.Project

	currentItem, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}


	err = json.Unmarshal(currentItem, &receivedProject)

	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}
	
	err = c.provider.Edit(&receivedProject)
	if err != nil {
		fmt.Printf("Ошибка Edit Provider: %s", err)
		return nil
	}

	return nil
}

func (c *CProject) Delete() revel.Result {
	c.provider = new(projects.ProjectProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	var receivedProject entity.Project


	getbody, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(getbody, &receivedProject)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	err = c.provider.Delete(&receivedProject)
	if err != nil {
		fmt.Printf("Ошибка Add Provider: %s", err)
		return nil
	}

	return nil
}


func (c *CProject) GetAll() revel.Result {
	c.provider = new(projects.ProjectProvider)
	err := c.provider.Init()
	if err != nil {
		fmt.Println(err)
	}
	projects, err := c.provider.List()
	if err != nil {
	log.Fatalf("Ошибка вывода List Provider: %s", err)

	}
	return c.RenderJSON(output.Correct(projects))
}








