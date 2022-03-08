# image-api

## Prerequisite 
You have to the go runtime version later than 1.12, so download the newest version
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
    image_data: multipart (png, jpeg, jpg)
```

**download**: `http://localhost:3000/image/:id`

uri param
```
    :id
```

## Limitations
We would cover this in the report
* we jwt token does not support "logout", we could achieves this by creating a blacklist, but it is out of this research scope
* login credentials are plain texts hardcoded in the file, still, just for simplicity we choose to hardcode them.

