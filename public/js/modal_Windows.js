//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по линии таска и открытие модального окна//////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////
function task_window_full(name, item, description, typeTask, time, priority, status){

if(status == "НОВЫЙ"){
  arrStatusMy = [{value: "Новый"}, {value: "В работе"}]
} else {
  arrStatusMy = [{value: "Пауза"}, {value: "В работе"}]
}

var idProject = $$("projTable").getSelectedItem().id;
    webix.ui({
        view:"window", width:500, id:"window", position:"center", move:true,
        head:"Задача", modal: true,
        body:{ rows: [
          { rows:[
 { view:"text",      id:"redactThemeTask",   value:name,  label:"Название"},
 { view:"textarea",  id:"redactDescriptionTask", value:description,  label:"Описание"},
 { view:"text",      id:"redactTypeTaskTask", value:typeTask,  label:"Тип", view:"select", options:arrTypes} ,
 { view:"text",      id:"redactTimeTask", value:time,  label:"Время"},
 { view:"text",      id:"redactPriorityTask", value:priority,  label:"Приоритет", view:"select", options:arrPrior},
 { view:"text",      id:"redactStatusTask", value:status,  label:"Статус", view:"select",options:arrStatusMy},
 { view:"text",      id:"redactidProjectTask", value: idProject,  label:"IDПроект"}
          ]},
          {cols:[
            {view:"button", id:"deleteTaskModal", value:"Delete", css:"webix_danger", inputWidth:70, click:deleteTaskModalWindow},
            {view:"button", id:"correctTaskModal", value:"correct", css:"webix_danger", inputWidth:70,  width: 70, click:correctTaskModalWindow},
            {view:"button", id:"exitTaskModal", value:"close", css:"webix_danger", inputWidth:70,  width: 70, click:exitTaskModalWindow},
    
          ]},
    
        ]
         }
      });
    
      webix.ui({
          rows:[
              {id:"frame", view:'iframe', src:'//docs.webix.com'}
          ]
      });
      $$("window").show();

      function deleteTaskModalWindow() { // кнопка delete - удаление в информация о задаче

        $$("taskTable").remove($$("taskTable").getSelectedId());
        $$("window").close();
        requestDEL_task(item);
    
      }
    
      function correctTaskModalWindow() { // кнопка correct - изменение в информация о задаче
        nameNew =  $$('redactThemeTask').getValue();
        descNew =  $$('redactDescriptionTask').getValue();
        typeNew =  $$('redactTypeTaskTask').getValue();
        timeNew =  $$('redactTimeTask').getValue();
        priorNew =  $$('redactPriorityTask').getValue();
        statusNew =  $$('redactStatusTask').getValue();
        idProject =  $$('redactidProjectTask').getValue();
        IdTask =  item.id;


         requestPOST_Task(idProject, IdTask, nameNew, descNew, typeNew, timeNew, priorNew, statusNew);
         $$("taskTable").clearAll()

         setTimeout(requestGET_task(idProject), 300);
            $$("window").close();
            $$("taskTable").refreshColumns();
         
      }


      function exitTaskModalWindow() { // кнопка close - выходи из информация о задаче
        $$("window").close();
      }
}
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по линии task и открытие модального окна///////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////



//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по линии person и открытие модального окна///////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////
function group_window_full(item, name, surname, yearsold, position, group){


  var idGroup = $$("projGroupTable").getSelectedItem().id;
  
  
 webix.ui({
    view:"window", width:500, id:"windowPerson", position:"center", move:true,
    head:"Сотрудник", modal: true,
    body:{ rows: [
      { rows:[
        { view:"text",      id:"redactNamePerson",   value:name,  label:"name"},
        { view:"text",      id:"redactsurnamePerson", value:surname,  label:"surname"},
        { view:"text",      id:"redactyearsoldPerson", value:yearsold,  label:"yearsold"},
        { view:"text",      id:"redactpositionPerson", value:position,  label:"Роль" , view:"select", options:[{value:"Ответственный"},{value:"Тим лид"}, {value:"Разработчик"}]},
 { view:"text",      id:"redactgroupPerson", value:idGroup,  label:"group" }, // вот
      ]},
      { cols:[
        {view:"button", value:"Delete", css:"webix_danger", inputWidth:70, click:deletePersonModalWindow},
        {view:"button", value:"correct", css:"webix_danger", inputWidth:70,  width: 70, click:correctPersonModalWindow},
        {view:"button", value:"close", css:"webix_danger", inputWidth:70,  width: 70, click:exitPersonModalWindow},
  
      ]},
  
    ]
     }
  });
  
  webix.ui({
      rows:[
          {id:"frame", view:'iframe', src:'//docs.webix.com'}
      ]
  });
  $$("windowPerson").show();
  
  
  
  
  function deletePersonModalWindow() { // кнопка delete - удаление в информация о задаче
    var item = $$("projGroupPersonTable").getSelectedItem();
    
   requestDEL_person(item);
   $$("projGroupPersonTable").clearAll()
   var idGroup = $$("projGroupTable").getSelectedItem().id;
   setTimeout(requestGET_person(idGroup), 300);

    $$("windowPerson").close();
  
  }
  
  function correctPersonModalWindow() { // кнопка correct - изменение person

    var id = item.id;
    var idgroup = $$("projGroupTable").getSelectedItem().id;
    var nameNew = $$("redactNamePerson").getValue();
    var SurnameNew = $$("redactsurnamePerson").getValue();
    var OldNew = $$("redactyearsoldPerson").getValue();
    var posNew = $$("redactpositionPerson").getValue();
    console.log(idgroup);

   requestPOST_Person(id, nameNew, SurnameNew, OldNew, posNew, idgroup);
   $$("projGroupPersonTable").clearAll()

   setTimeout(requestGET_person (idgroup), 300);

    $$("windowPerson").close();
  }
  
  function exitPersonModalWindow() { // кнопка close - выходи из информация о person
    $$("windowPerson").close();
  
  }


  
}


//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по линии проектной группы и открытие модального окна///////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////



//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по add task открытие модального окна///////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////

function add_task_window_full(){
  webix.ui({
        view:"window", width:500, id:"windowCreateTask", position:"center", move:true,
        head:"Создать задачу", modal: true,
        body:{ rows: [
          { rows:[
            { view:"text",    id:"themeTask", label:"name"},
            { view:"textarea",id:"descriptionTask", label:"description"},
            { view:"text",    id:"typeTaskTask",  label:"Тип", view:"select", options:arrTypes},
            { view:"text",    id:"timeTask", label:"time"},
            { view:"text",    id:"priorityTask",  label:"priority", view:"select", options:arrStatus},
          ]},
          {cols:[
            {view:"button", id:"createTaskModal", value:"create", css:"webix_danger", click:createCreateNewTaskWindow},
            {view:"button", id:"exitCreateTaskModal", value:"close", css:"webix_danger", click:closeCreateNewTaskWindow },
  
          ]},
  
        ]
         }
      });
      webix.ui({
          rows:[
              {id:"frame", view:'iframe', src:'//docs.webix.com'}
          ]
      });
      $$("windowCreateTask").show();
  
      function closeCreateNewTaskWindow() { // кнопка close - выходи из создание новой задачи
        
        $$("windowCreateTask").close();
      }
  


      function createCreateNewTaskWindow() { // кнопка create - создание новой задачи
        var name = $$('themeTask').getValue(); // nameP : информация о названии проекта
        var descT = $$('descriptionTask').getValue(); // descriptionTask : информация о описании Задачи
        var typeT = $$('typeTaskTask').getValue(); // typeTaskTask : информация о типе Задачи
        var timeT = $$('timeTask').getValue(); // timeTask : информация о времени Задачи
        var priorityT = $$('priorityTask').getValue(); // priorityTask : информация о приоритете Задачи
        var statusT = "НОВЫЙ" // statusTask : информация о статусе Задачи
        var IdProject = $$("projTable").getSelectedItem().id;



var arrElementAdd = {
  Time:        parseInt(timeT), 
  Name:        name, 
  IdProject:   parseInt(IdProject), 
  Description: descT, 
  TypeTask:    typeT, 
  Priority:    priorityT,
  Status:      statusT}
        requestPUT_task(arrElementAdd);

    $$("taskTable").clearAll()
    var idProject = $$("projTable").getSelectedItem().id;

    setTimeout(requestGET_task(idProject), 300);
        $$("windowCreateTask").close();



      }



}
//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по add task открытие модального окна////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////


//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по edit group открытие модального окна/////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////

function edit_group_window_full(name, item){

  webix.ui({
    view:"window", width:500, id:"windowEditGroup", position:"center", move:true,
    head:"Группа", modal: true,
    body:{ rows: [
      { rows:[
        { view:"text",      id:"redactnameGroup", value:name,  label:"Group"},
      ]},
      {cols:[
        {view:"button", value:"Delete", css:"webix_danger", inputWidth:70, click:deleteEditGroupModalWindow},
        {view:"button", value:"correct", css:"webix_danger", inputWidth:70,  width: 70, click:correctEditGroupModalWindow},
        {view:"button", value:"close", css:"webix_danger", inputWidth:70,  width: 70, click:exitEditGroupModalWindow},
  
      ]},
  
    ]
     }
  });
  
  webix.ui({
      rows:[
          {id:"frame", view:'iframe', src:'//docs.webix.com'}
      ]
  });
  $$("windowEditGroup").show();
  
  
  
  
  function deleteEditGroupModalWindow() { // кнопка delete - удаление группы
      let jsonProj = JSON.stringify({Id:item.id});
      const xhr = new XMLHttpRequest();
      console.log(jsonProj)
      xhr.open("DELETE", '/group:/idgroup');
      xhr.send(jsonProj);
      $$("projGroupTable").clearAll() 
      setTimeout(requestGET_groups, 300);
    $$("windowEditGroup").close();

  }
  function correctEditGroupModalWindow() { 
    var idGroup = item.id;
    var NameGroup = $$("redactnameGroup").getValue();
    requestPOST_Group(idGroup, NameGroup);
    $$("projGroupTable").clearAll() ;
    setTimeout(requestGET_groups, 300);
    $$("windowEditGroup").close();
  }
  
  
  function exitEditGroupModalWindow() { 
    $$("windowEditGroup").close();
  }



}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по edit group открытие модального окна/////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////





//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по add project открытие модального   ////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////

function add_project_window_full(){
    var modalFormCreateNewProj = webix.ui({
        view:"window", width:500, id:"windowProj", position:"center", move:true,
        head:"Создание проекта", modal: true,
        body:{ rows: [
          { rows:[
            
            { view:"text", id:"nameProj",  label:"Name"},
            { view:"text", id:"groupProj",  label:"Proj group", view:"select", name:"login", options:smallProjectsGroup }, 
          ]},
          {cols:[
            {view:"button", id:"createProjModal", value:"create", css:"webix_danger", click:createwindowProj},
            {view:"button", id:"exitProjkModal", value:"close", css:"webix_danger", click:closewindowProj},
    
          ]},
        ]
         }
      });
    
      webix.ui({
        rows:[
            {id:"frame", view:'iframe', src:'//docs.webix.com'}
        ]
      });
    
    
      $$("windowProj").show();
    
    
    
    
    
    
      function closewindowProj() { // кнопка close - выход из CREATE PROJECT
        $$("windowProj").close();
    
      }
      function createwindowProj() { // кнопка create - выход из CREATE PROJECT
    
       var nameP = $$('nameProj').getValue(); // nameP : информация о названии проекта
       var groupP = $$('groupProj').getValue(); // group : информация о группе проекта
       

             $$("windowProj").close();

            let jsonProj = JSON.stringify({Name:nameP, IdGroup:parseInt(groupP)});
            const xhr = new XMLHttpRequest();
            xhr.open("PUT", '/project');
            xhr.send(jsonProj);

             $$("projTable").clearAll() 

           requestGET_project();
      }
}


//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по add project открытие модального окна/////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////



//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по edit project открытие модального окна/////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////


function edit_project_window_full(item, idProject, name, group){

    webix.ui({
        view:"window", width:500, id:"windowEditProj", position:"center", move:true,
        head:"Проект", modal: true,
        body:{ rows: [
          { rows:[
            { view:"text",      id:"redactnameProj", value:name,  label:"Название"},
            { view:"text",      id:"redactgroupProj", value:group, view:"select", label:"Группа", options: smallProjectsGroup
          },
          ]},
          {cols:[
            {view:"button", value:"Delete", css:"webix_danger", inputWidth:70, click:deleteEditModalWindow},
            {view:"button", value:"correct", css:"webix_danger", inputWidth:70,  width: 70, click:correctEditModalWindow},
            {view:"button", value:"close", css:"webix_danger", inputWidth:70,  width: 70, click:exitEditModalWindow},
      
          ]},
      
        ]
         }
      });
      
      webix.ui({
          rows:[
              {id:"frame", view:'iframe', src:'//docs.webix.com'}
          ]
      });
      $$("windowEditProj").show();
      
      
      
      
      function deleteEditModalWindow() { // кнопка delete - удаление проекта
          let jsonProj = JSON.stringify({Id:item.id});
          const xhr = new XMLHttpRequest();
          console.log(jsonProj)
          xhr.open("DELETE", '/project:/idproject');
          xhr.send(jsonProj);
          $$("projTable").clearAll() 
          requestGET_project();

        $$("windowEditProj").close();
      }
      
    
      function correctEditModalWindow() { // кнопка correct - изменение person
        var nameProject = $$("redactnameProj").getValue();
        var idgroup = $$("redactgroupProj").getValue();
      
        requestPOST_Project(idProject, nameProject, idgroup);

        $$("projTable").clearAll() 
        setTimeout(requestGET_project, 300);
        

        $$("windowEditProj").close();

       

      }
      
      function exitEditModalWindow() { // кнопка close - выходи из информация о person
        $$("windowEditProj").close();
      }
      
}


//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по edit project открытие модального окна/////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////




//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по add person открытие модального окна////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////

function add_person_window_full(id){

    webix.ui({
		view:"window", width:500, id:"windowCreateNewProjectsGroupPerson", position:"center", move:true,
		head:"Добавить сотрудника", modal: true,
		body:{ rows: [
			{ rows:[
				{ view:"text",    id:"namePerson", label:"Имя"},
				{ view:"text",		id:"surnamePerson", label:"Фамилия"},
				{ view:"text",    id:"yearsPerson", label:"Лет"},
				{ view:"text",    id:"positionPerson", label:"Роль", view:"select", options:[{value:"Ответственный"},{value:"Тим лид"}, {value:"Разработчик"}]},
			]},
			{cols:[
				{view:"button", id:"createTaskModal", value:"Создать", css:"webix_danger", click:creatCreateProjectsGroupPersonWindow},
				{view:"button", id:"exitCreateTaskModal", value:"Закрыть", css:"webix_danger", click:closeCreateProjectsGroupPersonWindow },

			]},

		]
		 }
	});
	webix.ui({
			rows:[
					{id:"frame", view:'iframe', src:'//docs.webix.com'}
			]
	});
	$$("windowCreateNewProjectsGroupPerson").show();


	function closeCreateProjectsGroupPersonWindow() { // кнопка close - выходи из создание нового person
		$$("windowCreateNewProjectsGroupPerson").close();

	}

	function creatCreateProjectsGroupPersonWindow() { // кнопка create - создание нового person

		var nameP = $$('namePerson').getValue(); // namePerson : Имя создание нового person
		var surnameP = $$('surnamePerson').getValue(); // surnamePerson : Фамиллия создание нового person
		var yearsP = $$('yearsPerson').getValue(); // yearsPerson : Лет создание нового person
		var positionP = $$('positionPerson').getValue(); // positionPerson : Должность создание нового person
        $$("windowCreateNewProjectsGroupPerson").close();
    var idGroup =  $$("projGroupTable").getSelectedItem().id;

         requestPUT_person(nameP, yearsP, positionP, surnameP, id);
         $$("projGroupPersonTable").clearAll() 
         setTimeout(requestGET_person(idGroup), 300);


	}



}



//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по add person открытие модального окна////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////



//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Начало)Клик по add group открытие модального окна////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////

function add_group_window_full(){

    webix.ui({
        view:"window", width:500, id:"windowCreateNewProjectsGroup", position:"center", move:true,
        head:"Создание группы", modal: true,
        body:{ rows: [
            { rows:[
                { view:"text", id:"nameProjectsGroup",  label:"Название"},
            ]},
            {cols:[
                {view:"button", id:"createProjModal", value:"create", css:"webix_danger", click:createwindowProjGroup},
                {view:"button", id:"exitProjkModal", value:"close", css:"webix_danger", click:closewindowProjGroup},
    
            ]},
        ]
         }
    });
    webix.ui({
        rows:[
                {id:"frame", view:'iframe', src:'//docs.webix.com'}
        ]
    });
    
    
    $$("windowCreateNewProjectsGroup").show(); 
    function closewindowProjGroup() { // кнопка close - выход из CREATE PROJECT GROUP
      $$("windowCreateNewProjectsGroup").close();
    }



    function createwindowProjGroup() { // кнопка create 
    
     var namePG = $$('nameProjectsGroup').getValue(); // namePG : информация о названии проекта
            // var arrElementAdd = {name:namePG, value:namePG}
            // $$('projGroupTable').add(arrElementAdd);

            requestPUT_Group(namePG);
            setTimeout(requestGET_groups, 300);

        $$("windowCreateNewProjectsGroup").close();
      
    
    
    
    
    }
}


//////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////(Конец)Клик по add group открытие модального окна////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////


