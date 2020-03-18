document.addEventListener("DOMContentLoaded", function (event) {
  //GET запросы//
  requestGET_project();
  requestGET_type(); 
  requestGET_groups();



  ///////////////////////////////////////////////////////////////////
    var projectsTable = {
      view: "datatable",
      css: "webix_header_border",
      id: "projTable",

      columns: [
        {id: "name", css: "align", fillspace: true,},
      ],
      scroll: true,
      select: true,
      data: smallProjects,
    }
    
  // /////////////////////////////////////////////////////////////////
   var task = {
    view: "datatable",
    css: "webix_header_border",
    id: "taskTable",
    columns: [
      {id: "name", css: "align", width:100, css: "rank", adjust: "data"},
      {id: "desc", css: "align", fillspace: true},
      {id: "typeTask", css: "align",  width:200},
      {id: "time", css: "align", width:50},
      {id: "priority", css: "align", width:75},
      {id: "status", css: "align", width:75}

    ],
    scroll: false,
    select: true,
   }
  // /////////////////////////////////////////////////////////////////
	var projectsGroupTable = {
		view: "datatable",
		css: "webix_header_border",
		id: "projGroupTable",
		columns: [
      {id: "name", css: "align", fillspace: true},
      
		],
		scroll: false,
		select: true,
		data: smallProjectsGroup,
	}

	// /////////////////////////////////////////////////////////////////
	var projectsGroupPersonTable = {
		view: "datatable",
		css: "webix_header_border",
		id: "projGroupPersonTable",
		columns: [
			{id: "name", css: "align", fillspace: true},
			{id: "surname", css: "align", fillspace: true},
			{id: "yearsold", css: "align", fillspace: true},
			{id: "position", css: "align", fillspace: true},

		],
		scroll: false,
		select: true,
		data: smallProjectsGroupPerson,
	}

  // ///////////////////////////////////////////////////////////////// Основная часть ///////////////////////////////////////////////////////

	webix.ui({
    
    view: "tabview",
	  cells: [  
      {
				width: 150,
				header: "<span class='webix_icon fa-tasks'></span>Проекты",
				body: {
					rows: [
            
					{     cols: [
								{
									rows: [
										{type: "header", template: "Проекты", width:150},
										projectsTable, {cols:[
                      {view:"button", id:"createProj", value:"Создать", css:"webix_primary"},{view:"button", id:"editProj", value:"Изменить", css:"webix_primary"}
                    ]},
									]
								},
								{
									rows: [
										{view: "toolbar", elements:[
											{view: "label", label: "Задачи"}, {view:"button", id:"createNewTask", value:"Добавить задачу", css:"webix_primary", inputWidth:200, width: 200 }
										]},
										task
									]

								}
							]}
					]
				}
      },

      {
				width: 150,
	      header: "<span class='webix_icon fa-users'></span>Группы",
	      body: {
					rows: [
					{     cols: [
								{
									rows: [
										{type: "header", template: "Группы", width:150},
                    projectsGroupTable, {cols: [
                      {view:"button", id:"createGroupProject", value:"Создать", css:"webix_primary"},{view:"button", id:"editGroup", value:"Изменить", css:"webix_primary"}

                    ]},
									]
								},
								{
									rows: [
										{view: "toolbar", elements:[
											{view: "label", label: "Персонал"}, {view:"button", id:"createNewProjectsGroupPerson", value:"Добавить сотрудника", css:"webix_primary", inputWidth:200, width: 200 }
										]},
										projectsGroupPersonTable
									]

								}
							]}
					]
	      }
      },
	  ]
  });
  

  ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
  ///////////////////////////// клик по линии task                      /////////////////////////////////////////
  ///////////////////////////////////////////////////////////////////////////////////////////////////////////////


  $$("taskTable").attachEvent("onItemClick", function(clicked_json, mous_ev, div){
    
    var item = $$("taskTable").getSelectedItem();
    var name = item.name;
    var description = item.desc;
    var typeTask = item.typeTask;
    var time = item.time;
    var priority = item.priority;
    var status = item.status;
    var peoples = item.peoples;
    task_window_full(name, item, description, typeTask, time, priority, status);

})


  ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
  /////////////////////////////          клик по кнопке add task        /////////////////////////////////////////
  ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
  $$("createNewTask").attachEvent("onItemClick", function(clicked_json, mous_ev, div){ //  клик по кнопке add task
    var id = $$("projTable").getSelectedItem(); // id для того чтобы присвоить это же число ip 
    add_task_window_full();
    
  });
  /////////////////////////////////////////////////////////////////////////////////////////////////////////////////
  ///////////////////////////////////////клик по кнопке add projct       //////////////////////////////////////////
  /////////////////////////////////////////////////////////////////////////////////////////////////////////////////


  $$("createProj").attachEvent("onItemClick", function(clicked_json, mous_ev, div){ //  клик по кнопке add project

    add_project_window_full();

  })
  ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
  /////////////////////////////////////////клик по проекту и создание таблицы с задачами/////////////////////////
  ///////////////////////////////////////////////////////////////////////////////////////////////////////////////


  $$("projTable").attachEvent("onItemClick", function(clicked_json, mous_ev, div){

    // var item = $$("projTable").getSelectedItem();


    $$("taskTable").clearAll()
    var idProject = $$("projTable").getSelectedItem().id;
    requestGET_task(idProject);

//     smallTask.find(function(currentItem){
//       if(item.id == currentItem.ip){
// var arrElementAdd = {id:currentItem.id, ip:currentItem.ip, theme: currentItem.theme, desc:currentItem.desc, typeTask:currentItem.typeTask, peoples:currentItem.peoples, time: currentItem.time, priority: currentItem.priority , status: currentItem.status}
//         $$('taskTable').add(arrElementAdd);
//       }
//     })
	
  })




////////////////////////////////////////////////////////////////////////////////////////////////////////////управление проектными группами//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////



//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////Кнопка add person//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

$$("createNewProjectsGroupPerson").attachEvent("onItemClick", function(clicked_json, mous_ev, div){ //  клик по кнопке add person
  var id = $$("projGroupTable").getSelectedItem().id;
  add_person_window_full(id);

});




//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////Кнопка add group///////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

$$("createGroupProject").attachEvent("onItemClick", function(clicked_json, mous_ev, div){ //  клик по кнопке add project
  add_group_window_full();
})

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////Кнопка edit group///////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


$$("editGroup").attachEvent("onItemClick", function(clicked_json, mous_ev, div){ //  клик по кнопке add task
  
  var item = $$("projGroupTable").getSelectedItem();
  var name = item.name;

  edit_group_window_full(name, item);

});


///////////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////клик по проектной группе и создание таблицы с person//////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
$$("projGroupTable").attachEvent("onItemClick", function(clicked_json, mous_ev, div){
	var item = $$("projGroupTable").getSelectedItem();

  var idgroup = $$("projGroupTable").getSelectedItem().id;
  requestGET_person(idgroup);
  $$("projGroupPersonTable").clearAll()


	// smallProjectsGroupPerson.find(function(currentItem){
	// 	 if(item.id == currentItem.ig){
	// 		var arrElementAdd = {id:currentItem.id, ig:currentItem.ig, name:currentItem.name, yearsold:currentItem.yearsold, surname:currentItem.surname, position:currentItem.position}
	// 		$$('projGroupPersonTable').add(arrElementAdd);
	// 	  }
  // })
})

///////////////////////////////////////////клик по линии person group ///////////////////////////////////////////////
//!!//
$$("projGroupPersonTable").attachEvent("onItemClick", function(clicked_json, mous_ev, div){

  var item = $$("projGroupPersonTable").getSelectedItem();
   var name = item.name;
   var surname = item.surname;
   var yearsold = item.yearsold;
   var position = item.position;
   var group = item.ig;

   group_window_full(item, name, surname, yearsold, position, group);
})


///////////////////////////////////////////////////////////клик по кнопке edit ///////////////////////////////////////
$$("editProj").attachEvent("onItemClick", function(clicked_json, mous_ev, div){
  
   var item = $$("projTable").getSelectedItem();
   var idProject = item.id 
   var name = item.name;
   var group = item.id_group;

   console.log(group);

   edit_project_window_full(item, idProject, name, group);


})


var exit = document.getElementById("ButtonExitDelSession");


exit.addEventListener("click", function(){
  
  delsession();

})










});
