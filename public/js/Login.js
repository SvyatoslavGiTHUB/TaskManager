
document.addEventListener("DOMContentLoaded", function (event) {
webix.ui({
    view:"form", 
    id:"log_form",
    width:300,
    elements:[
        { view:"text", label:"Логин", id:"Login", name:"email"},
        { view:"text", type:"password", id:"Password", label:"Пароль", name:"password"},
        { margin:5, cols:[
            { view:"button", value:"Войти" , css:"webix_primary", click:LoginRequest},
        ]}
    ]
});


function LoginRequest(){
    login =  $$('Login').getValue();
    Password =  $$('Password').getValue();

    requestPOST_USER(login, Password);

}




})