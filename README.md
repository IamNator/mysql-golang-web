# mysql-stuff

#### Contributors
https://github.com/Tee-py - FrontEnd        
natverior1@gmail.com - BackEnd

https://mysql-golang-app.herokuapp.com/

#### PhoneBook
A simple web interface. A phone book with the Backend built entirely on golang, mySql database used.
 
#### Database Operations


#### /api/update - POST
Request POST application/json for adding contacts   
	{    
		"fname": "Oluwa,    
		"lname": "Tobi",   
		"phone_number": "+2349043495346",   
		"id": ""   
	}    
	
##### /api/fetch - GET
Response application/json returned from fetching contacts     
[  
 	{    
 		"fname": "Nator,    
 		"lname": "Verinumbe",   
 		"phone_number": "+2349045057268",   
 		"id": "2du-43-432-34ddf-43f-fqe"   
 	}   
]       
	
##### /api/delete - DELETE
Request Delete application/json for delete contacts        
	{       
		"id": "242-43-432-34ddf-43f-fq4"    
	}     
 
### Sessions
SessionID is the cookie value. A cookie is created when the user login or registers


##### /api/register - POST
Request POST application/json for Registering User     
	{      
	    "username": "Oluwatobi"   
	    "password": "Password"   
	}   


##### /api/login - POST
json for Logging in User     
    {   
	    "username": "Nator"    
	    "password": "29PassWerd"    
    }    
    Please properly enclose the parameters in double quotes.

A user can only alter data when he logs into his account,
 and he/she can only alter/view data he/she stored.


