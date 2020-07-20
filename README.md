# empfetcher
This repo contains code for RESTful API Services using **Goa** framework


## Design file
The Service has been implemented using Goa framework which is built on top of Go and it uses built-in DSL that allows describing the design of the microservice

For more details: https://goa.design/learn/intro/

The below file contains the code for **design** of micro service
This is where our design of Micro-service starts, we need to mention all our methods/endpoints for API here

https://github.com/shivakumar5/empfetcher/blob/master/design/design.go

The Service has below endpoints:

* Add
* Update (YTBD)
* List
* Show
* Delete
* Restore
* Viewdeleted 
* Search (YTBD)

## Database
I have used **MSSQL SERVER** for database operations
The below file contains the code for SQL Operations:

https://github.com/shivakumar5/empfetcher/blob/master/services/mssql.go

## Business Logic
And, the below file contains the **Business Logic** where we call our downstream service(i.e. MSSQL in this case) 
https://github.com/shivakumar5/empfetcher/blob/master/empfetcher.go

## Main file
we will initliaze all the clients, endpoints and services in main.go, below is the link:
https://github.com/shivakumar5/empfetcher/blob/master/cmd/empfetcher/main.go

## Folder Structure
![Alt text](screenshots/Folder.png?raw=true "")

## Execute in Local

To execute it in Local, run '**go build cmd/empfecther/main.go cmd/empfecther/http.go**'
It will generate the binary file called '**empfetcher**', run the binary file like '**./empfetcher**' which will start running in default port **8080**
**Note:** If you want to change the default port, use flag options. execute below command, it will show the different flags available for service(taken care by goa based on our implementation)

**./empfecther --help**
``` 
Usage of ./empfetcher:
  -database string
        Name of the SQL database (default "employee")
  -debug
        Log request and response bodies
  -domain string
        Host domain name (overrides host domain specified in service design)
  -host string
        Server host (valid values: localhost) (default "localhost")
  -http-port string
        HTTP port (overrides host HTTP port specified in service design)
  -password string
        Password for SQL server (default "")
  -secure
        Use secure scheme (https or grpcs)
  -sql-server string
        SQL server (default "localhost")
  -username string
        Username for SQL server (default "sa")
 ```
 
 **./empfecther -http-port 8002**
 Service starts running like below:
 
 ```
[empfetcherapi] 07:29:12 HTTP "Add" mounted on POST /api/v1/employees
[empfetcherapi] 07:29:12 HTTP "Update" mounted on PUT /api/v1/employees
[empfetcherapi] 07:29:12 HTTP "List" mounted on GET /api/v1/employees
[empfetcherapi] 07:29:12 HTTP "Show" mounted on GET /api/v1/employees/{id}
[empfetcherapi] 07:29:12 HTTP "Delete" mounted on DELETE /api/v1/employees/{id}
[empfetcherapi] 07:29:12 HTTP "Restore" mounted on PUT /api/v1/employees/{id}
[empfetcherapi] 07:29:12 HTTP "Viewdeleted" mounted on GET /api/v1/employees/deactivated
[empfetcherapi] 07:29:12 HTTP "Search" mounted on GET /api/v1/employees/search/{name}
[empfetcherapi] 07:29:12 HTTP server listening on "localhost:8002"

```
 
 ## Request the service using Postman
 
 The server is running using binary file generated by *go build*.
 Now, we can call our service's different methods/endpoints in postman 
 
 **empfetcher** Service base URL: **http://localhost:8080/api/v1/employees/**
 
 * Add - Adds a new employee record
 ![Alt text](screenshots/Add.png?raw=true "")
 
 * Update - Updates existing employee record
 It is **Not yet implemented** due to time constrains. here, the multipart upload has to be implemented.



 
 * List - Lists all the employee records
 ![Alt text](screenshots/List.png?raw=true "")
 
 * Show- Shows the employee record based on id
  ![Alt text](screenshots/Show.png?raw=true "")
 
 * Delete - Deletes employee record - **softDelete** if **permdelete** flag is not passed.
![Alt text](screenshots/Delete.png?raw=true "")
 
 * Restore - Restores softDeleted employee record
![Alt text](screenshots/Restore.png?raw=true "")
 
 * Viewdeleted - Lists all the deactived employee records
![Alt text](screenshots/View.png?raw=true "")


 
 * Search - Not yet implemented due to time constraints
 
 
 
 
 
 
 
        
