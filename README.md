## Basic Backend Structure in GO
This is a backend for Feeds App demonstrating a basic backend structure in go language. In this project the backend used is posgresql. So to run the backend you will need postgresql. To download postgresql go to [https://www.postgresql.org/download/] .

# How it works?
1. Make an .env file containing the port number and database url.
2. Either make your own database or follow the commands given below.
   ```
   go install github.com/pressly/goose
   ```
   After installing goose type the following commands.
   ```
   gooose postgres [Your Database URL] up
   ```
3. After creating the database run.
   ```
   go build && go_backend
   ```
4. This will run the project now you can check the endpoints given in the project.
   

