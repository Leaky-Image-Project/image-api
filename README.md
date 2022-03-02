# chat-api


## Prerequisite 
You have to the go runtime version later than 1.2, so download the newest version
## Run Project
```
go run main.go
```

## Available API endpoints
### Auth
**login**: `http://localhost:3000/auth/login`

body
```json
{
    "username": "victim0",
    "password": "victim0Pass"
}
```
### Image

**upload**: `http://localhost:3000/image/upload`

body
```
    file (png, jpeg, jpg)
```

**download**: `http://localhost:3000/image/:id`

uri param
```
    :id
```


