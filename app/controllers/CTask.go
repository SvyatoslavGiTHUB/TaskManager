package controllers

import (
	"RBStask/app/models/Providers/tasks"
	 "RBStask/app/models/entity"
	"RBStask/app/output"
	 "encoding/json"
	"fmt"
	"log"
	"strconv"
	"io/ioutil"

	"github.com/revel/revel"
)


type CTask struct {
	*revel.Controller
	provider *tasks.TaskProvider
}


func (c *CTask) GetAll() revel.Result {
	c.provider = new(tasks.TaskProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	getrequestTask := c.Params.Get("tasks")
	ProjectId, err := strconv.ParseInt(getrequestTask, 10, 64)
	if err != nil {
		fmt.Printf("Ошибка чтения ParseInt: %s", err)
	}

	tasks, err := c.provider.GetTasks(ProjectId);
	if err != nil {
		log.Fatalf("Ошибка передачи GetTasks в proveder: %s", err)
	}
	return c.RenderJSON(output.Correct(tasks))
}




func (c *CTask) Add() revel.Result {
	c.provider = new(tasks.TaskProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}
	receivedTask := entity.Task{}
	getBody, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Ошибка чтения Body: %s", err)
		return nil
	}

	err = json.Unmarshal(getBody, &receivedTask)
	fmt.Println(receivedTask)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}
	err = c.provider.Add(&receivedTask)
	if err != nil {
		fmt.Printf("Ошибка Add Provider: %s", err)
		return nil
	}
	return nil
}


func (c *CTask) Delete() revel.Result {
	c.provider = new(tasks.TaskProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	receivedTask := entity.Task{}

	getBody, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Err: %s", err)
		return nil
	}

	err = json.Unmarshal(getBody, &receivedTask)
	if err != nil {
		log.Fatalf("Сбой JSON формата: %s", err)
		return nil
	}

	err = c.provider.Delete(&receivedTask)
	if err != nil {
		fmt.Printf("Ошибка Delete Provider: %s", err)
		return nil
	}

	return nil
}





func (c *CTask) Edit() revel.Result {
	c.provider = new(tasks.TaskProvider)
	err := c.provider.Init()
	if err != nil {
		log.Fatalf("Сбой подключения к базе данных: %s", err)
	}

	receivedTask := entity.Task{}

	currentItem, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		fmt.Printf("Err: %s", err)
		return nil
	}


	err = json.Unmarshal(currentItem, &receivedTask)
	fmt.Println(receivedTask)
	if err != nil {
		log.Fatalf("Сбой демаршалинга JSON: %s", err)
		return nil
	}
	
	err = c.provider.Edit(&receivedTask)
	if err != nil {
		fmt.Printf("Err: %s", err)
		return nil
	}

	return nil
}











