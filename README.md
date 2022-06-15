# Todo List Frontend and Backend

Simple Todo List Page created with Javascript, HTML and CSS to make requests to Go API and store data in MySQL using Docker. Therefore it connects frontend to backend.

## Using Docker

Before get Docker started [Colima](https://smallsharpsoftwaretools.com/tutorials/use-colima-to-run-docker-containers-on-macos/) was used to prepare the environment to get the [Docker](https://www.docker.com/) up.

So after installing Colima and Docker and creating docker-compose.yml the terminal was opened:

`colima start`

After it successfully ran:

`docker-compose up`

Then we will have our Postgresql ready to accept connections. Furthermore it is possible to connect the database in a database tool as DBeaver to make queries and see data added in the table.

## Running Project

After Docker is up and the dbParams.env is correctly configured with host, port, user, password and database name we can run the project by:

`go run main.go`

Done that we will see the lines being added to the table by the terminal output. Then we can check if all the data is in the Postgresql table by making a query in a database tool such as DBeaver:

`SELECT * FROM todo_item_models`

## CRUD actions in Postman

Postman was used to make CRUD actions. For this we have to use the following URLS to each action:

> ACTION: POST /
> URL: http://localhost:3000/todo

> ACTION: GET /
> URL: http://localhost:3000/todo; 
> URL: http://localhost:3000/todo/{id}

> ACTION: PUT /
> URL: http://localhost:3000/todo/{id}

> ACTION: DELETE /
> URL: http://localhost:3000/todo/{id} 

Then we get the response with a successful message and the data requested.