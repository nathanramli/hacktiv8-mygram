# mygram

## directory structure 
```
   .
    ├── common                   # Response JSON data
    ├── cmd                      # Include main.go
    ├── config                   # Configure Database Connection 
    ├── httpserver               # Include controllers, repositories, services and router.go
    ├── .env.example             # copying secret variable to connect db from .env
    ├── .gitignore               # hiding .env
    ├── go.mod                 
    ├── go.sum                   
    ├── LICENSE
    └── 
```
### run hacktiv8-mygram
```
go run cmd/main.go
```

## list endpoint

### Register : Registration user account
```
POST https://hacktiv8-mygram-production.up.railway.app/v1/users/register
```

### Login : Login user account
```
POST https://hacktiv8-mygram-production.up.railway.app/v1/users/login
```
### UpdateUser : Edit and Update User account
```
PUT https://hacktiv8-mygram-production.up.railway.app/v1/users
```

### DeleteUser : Delete User account from list
```
DELETE https://hacktiv8-mygram-production.up.railway.app/v1/users
```
### GetPhotos : Display all photos list  
```
GET https://hacktiv8-mygram-production.up.railway.app/v1/photos
```

### CreatePhoto : Create/Post Photo from specific user
```
POST https://hacktiv8-mygram-production.up.railway.app/v1/photos
```

### UpdatePhoto : Edit and Update Photo 
```
PUT https://hacktiv8-mygram-production.up.railway.app/v1/photos/:photoId
```

### DeletePhoto : Delete Photo from list
```
DELETE https://hacktiv8-mygram-production.up.railway.app/v1/photos/:photoId
```
### GetComments : Display all comment list  
```
GET https://hacktiv8-mygram-production.up.railway.app/v1/comments
```

### CreateComment : Create/Post comment from specific user
```
POST https://hacktiv8-mygram-production.up.railway.app/v1/comments
```

### UpdateComment : Edit and Update Comment
```
PUT https://hacktiv8-mygram-production.up.railway.app/v1/comments/:commentId
```

### DeleteComment : Delete Comment from list
```
DELETE https://hacktiv8-mygram-production.up.railway.app/v1/comments/:commentId
```
### GetSocialMedias : Display all social medias list from specific users  
```
GET https://hacktiv8-mygram-production.up.railway.app/v1/socialmedias
```

### CreateSocialMedia : Create/Post socialmedia from specific user
```
POST https://hacktiv8-mygram-production.up.railway.app/v1/socialmedias
```

### UpdateSocialMedia : Edit and Update social media
```
PUT https://hacktiv8-mygram-production.up.railway.app/v1/socialmedias/:socialmediaId
```

### DeleteSocialMedia : Delete social media from list
```
DELETE https://hacktiv8-mygram-production.up.railway.app/v1/socialmedias/:socialmediaId
```

### Jobdesk member

- MAULA IZZA AZIZI (GLNG-KS04-020) : 
   -  All Social Medias EndPoint
   -  UpdateComment


- HEZKYA NATANAEL RAMLI (GLNG-KS04-008) : 
   -  Initialize Project
   -  All Users Endpoint
   -  DeleteComment
   -  Mocking
   -  Unit Test
   -  Deployment

- MUHAMAD RESTU FADILLAH (GLNG-KS04-002) : 
   -  All Photos EndPoint
   -  GetComments
   -  CreateComment
   -  Postman Collection

