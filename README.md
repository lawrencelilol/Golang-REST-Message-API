# Golang REST Message API
A simple API that allows users to get and post messages. It uses https://reqres.in/api/users as a third party API for users authentication (with user's first name as password)

# Build and Run

make build
chmod +x ./server
./server

# Endpoints

/admin/user - GET
/admin/user - POST

# Testing

```
go run main.go

```

```
 # Check the user needs basic authentication to access the api
 
 curl localhost:8080/admin/user
 
 # Check the friendly warning if you don't have a message saved yet
 
 curl --user "janet.weaver@reqres.in:Janet" http://localhost:8080/admin/user
 
 # Use the username:janet.weaver@reqres.in with password Janet to post a message
 
 curl -X POST http://localhost:8080/admin/user -H "Content-Type: application/json" -d'{"message": "I am Emma !"}' --user "janet.weaver@reqres.in:Janet"

# Retrieve the message you just set

curl --user "janet.weaver@reqres.in:Janet" http://localhost:8080/admin/user

# Stop the program

"control" + "c"

```

# Save Message

The message will be saved in the same folder as the main.go file and named after the user's ID







 

