
///GET запрос / project ///
function requestGET_project(){
    let url = '/project';
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url);
    xhr.onload = function(callback){
      let res = JSON.parse(xhr.response);
      var fullprojects = res.Data; 
      if (res.Result == 0){
        console.log("Загрузка проектов выполнена");

        fullprojects.find(function(currentItem){ 
          var arrElementAdd = {id:currentItem.Id, name:currentItem.Name, id_group:currentItem.IdGroup}
        
            $$('projTable').add(arrElementAdd);
        })  
      } else console.log("Ошибка загрузки данных должностей " + res.ErrorText);
  
    };
    xhr.onerror = () => {
  
    }
    xhr.send(); 
  }



  //////////////////////////////////////////////////////////

  function requestGET_groups(){
    let url = '/group';
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url);
    xhr.onload = function(){
      let res = JSON.parse(xhr.response);
      var fullprojects = res.Data; 
      smallProjectsGroup.splice(0, smallProjectsGroup.length);
      smallProjectsGroup.push({value:"Без группы"});
      if (res.Result == 0){
        $$('projGroupTable').clearAll();
        
        console.log("Загрузка групп выполнена");

        fullprojects.find(function(currentItem){ 
          var arrElementAdd = {id:currentItem.Id, name:currentItem.Name, value:currentItem.Name }
          
          smallProjectsGroup.push(arrElementAdd);
            $$('projGroupTable').add(arrElementAdd);
        })  
      } else console.log("Ошибка загрузки данных " + res.ErrorText);
      return fullprojects;
    };
    xhr.onerror = () => {
  
    }
    xhr.send(); 
 

  }



  /////////////////////////////////////////////////////////

  function requestGET_task(IdProject){
    
    let url = '/tasks' + '/' + IdProject;
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url);
    xhr.onload = function(callback){
      let res = JSON.parse(xhr.response);
      var fulltasks = res.Data; 
      if (res.Result == 0){

        console.log("Загрузка тасков выполнена");
        fulltasks.find(function(currentItem){ 
          var arrElementAdd = {id:currentItem.Id, name:currentItem.Name, desc:currentItem.Description, time:currentItem.Time, priority:currentItem.Priority, status:currentItem.Status, typeTask:currentItem.TypeTask}
          $$('taskTable').add(arrElementAdd);
        })
        // smallTask.find(function(currentItem){
        // })


      } else console.log("Ошибка загрузки тасков " + res.ErrorText);
  
    };
    xhr.onerror = () => {
 
    }
    xhr.send(); 

  }


  //////////////////////////////////////////////////////////
  function requestGET_type(){
    let url = '/type';
    const xhr = new XMLHttpRequest();
    xhr.open('GET', url);
    xhr.onload = function(){
      let res = JSON.parse(xhr.response);
      var fulltype = res.Data; 
      if (res.Result == 0){
        console.log("Загрузка типов выполнена");
        fulltype.find(function(currentItem){ 
      var     currentType = {value: currentItem.Name}
          arrTypes.push(currentType)  
        }) 
      } else console.log("Ошибка загрузки типов " + res.ErrorText);
  
    };
    xhr.onerror = () => {
  
    }
    xhr.send(); 

  }

  //////////////////////////////////////////////////////////


  function requestPUT_task(objectAdd){
    console.log(objectAdd);
    let jsontask = JSON.stringify(objectAdd);
    const xhr = new XMLHttpRequest();
    xhr.open("PUT", '/project:/idproject:/task');
    xhr.send(jsontask);
    smallTask.push(objectAdd); 
    $$('taskTable').add(objectAdd);
  }


  //////////////////////////////////////////////////////////

  function requestDEL_task(item){
  
    let jsonProj = JSON.stringify({Id:item.id});
    const xhr = new XMLHttpRequest();
    console.log(jsonProj)
    xhr.open("DELETE", '/project:/idproject:/task:/idtask');
    xhr.send(jsonProj);
  }

/////////////////////////////////////////////////////////////////////////////////////del person



  function requestDEL_person(item){
    let url = '/group' + '/' + item.id;
    const xhr = new XMLHttpRequest();
    xhr.open("DELETE", url);
    xhr.send(null);  
  }

/////////////////////////////////////////////////////////////////////////////////////get person



  function requestGET_person(idgroup){
      let url = '/group' + '/' + idgroup;
      const xhr = new XMLHttpRequest();
      xhr.open('GET', url);
      xhr.onload = function(){
        let res = JSON.parse(xhr.response);
        var fullperson = res.Data; 
        if (res.Result == 0){
          console.log("Загрузка сотрудников выполнена");
          fullperson.find(function(currentItem){ 
 var arrElementAdd = {id:currentItem.Id, name:currentItem.Name, position:currentItem.Position, surname:currentItem.Surname, yearsold:currentItem.Old}
            $$('projGroupPersonTable').add(arrElementAdd);
            console.log(arrElementAdd);
          })
          // smallTask.find(function(currentItem){
          // })
        } else console.log("Ошибка загрузки тасков " + res.ErrorText);
    
      };
      xhr.onerror = () => {
   
      }
      xhr.send(); 
  
    
  }



/////////////////////////////////////////////////////////////////////////////////////создание person
  function requestPUT_person(nameP, yearsP, positionP, surnameP, id){
    let jsonperson = JSON.stringify({
      Name:nameP, 
      Surname:surnameP, 
      Old:parseInt(yearsP), 
      IdGroup: id, 
      Position:positionP});
    const xhr = new XMLHttpRequest();
    xhr.open("PUT", '/person');
    xhr.onload = function(){
      let res = JSON.parse(xhr.response);
      webix.message(
        "Логин:" + nameP + "\n"+
        "Пароль:" + res.Data );
    }
    xhr.send(jsonperson);

    $$('projGroupPersonTable').add({name:nameP, yearsold:yearsP , position:positionP,  surname:surnameP, id_group:id});
  }
////////////////////////////////////////////////////////////////////////////////////////////////////edit project

function requestPOST_Project(idProject, nameProject, idgroup){

  console.log(idgroup);

  let jsonProjEdit = JSON.stringify({Id:idProject, Name: nameProject, IdGroup:parseInt(idgroup)});
  const xhr = new XMLHttpRequest();
  console.log(jsonProjEdit)
  xhr.open("POST", '/project:/idproject');
  xhr.send(jsonProjEdit);
}


function requestPUT_Group(namePG){
  let jsonProj = JSON.stringify({Name:namePG});
  const xhr = new XMLHttpRequest();
  console.log(jsonProj)
  xhr.open("PUT", '/group');
  xhr.send(jsonProj);
}


function requestPOST_Person(id, nameNew, SurnameNew, OldNew, posNew, idgroup){
  let jsonProjEdit = JSON.stringify({
    Id:           parseInt(id),
    Name:         nameNew, 
    Surname:      SurnameNew,
    Old:          parseInt(OldNew),
    IdGroup:      parseInt(idgroup),
    Position:     posNew,
    });

  const xhr = new XMLHttpRequest();
  console.log(jsonProjEdit)
  xhr.open("POST", '/group:/idgroup:/person:/idperson');
  xhr.send(jsonProjEdit);
}





function requestPOST_Task(idProject, IdTask, nameNew, descNew, typeNew, timeNew, priorNew, statusNew){

  let jsonProjEdit = JSON.stringify({
    Id:           parseInt(IdTask), 
    IdProject:    parseInt(idProject),
    Name:         nameNew, 
    Description:  descNew, 
    Time:         parseInt(timeNew),
    Priority:     priorNew,
    Status:       statusNew,
    TypeTask:     typeNew,
  });
  const xhr = new XMLHttpRequest();
  console.log(jsonProjEdit)

  xhr.open("POST", '/task/:idtask');

  xhr.send(jsonProjEdit);
  
}







function requestPOST_Group(idGroup, NameGroup){

  let jsonProjEdit = JSON.stringify({
    Id:           parseInt(idGroup),
    Name:         NameGroup, 
  });
  const xhr = new XMLHttpRequest();
  console.log(jsonProjEdit)
  xhr.open("POST", '/group:/idgroup');

  xhr.send(jsonProjEdit);
  
}









function requestPOST_USER(login, Password){
  let jsontask = JSON.stringify({Login: login , Password: Password });
  const xhr = new XMLHttpRequest();
  xhr.open("PUT", '/user');
  xhr.onload = function(callback){
    let res = JSON.parse(xhr.response);
    var resul = res.Data;
    if(resul == "true"){
      window.location.replace("/")
    }else{
      webix.message(resul);
    } 

    
  }
  xhr.onerror = () => {
  }

  xhr.send(jsontask);
}




function delsession(){
const xhr = new XMLHttpRequest();
xhr.open("DELETE", '/delete');
xhr.onload = function(callback){
  let res = JSON.parse(xhr.response);
  var resul = res.Data;
  if(resul == "true"){
    window.location.replace("/")
  }
}
xhr.send();
}



