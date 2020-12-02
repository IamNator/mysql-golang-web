# mysql-stuff

###  Contributors
https://github.com/Tee-py - FrontEnd        
natverior1@gmail.com - BackEnd

https://mysql-golang-app.herokuapp.com/

## PhoneBook
 A simple web interface. The Backend is built entirely on golang, mySql database used.
 
### Database Operations

##### /api/update - POST
json for adding contacts   
	{    
		"fname":<First Name>,    
		"lname":<Last Name>,   
		"phone_number":<Phone Number>,   
		"id":<id>   
	}    
	
##### /api/fetch - GET
json for fetching contacts   

	
##### /api/delete - DELETE
json for delete contacts        
	{       
		"id":<id>    
	}     
 
### Sessions
SessionID is the cookie value. A cookie is created when the user login or registers


##### /api/register - POST
json for Registering User     
	{      
	    "username":<username>   
	    "password":<password>   
	}   


##### /api/login - POST
json for Logging in User     
    {   
	    "username":<username>    
	    "password":<password>    
    }    
