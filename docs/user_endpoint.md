## User Endpoint

### Add New User
Endpoint

    POST: http://localhost:8024/api/v1/user

Request

    {
    	"name": "Muflih Kholidin",
    	"email": "muflic.24@gmail.com",
    	"phone": "08123456789",
    	"role": "administrator",
    	"password": "secret"
    }
    
Response
    
    {
        "data": {
            "id": "0dc21127-5d19-4a85-8809-4a9cc14b07d5",
            "email": "muflic.24@gmail.com",
            "name": "Muflih Kholidin",
            "phone": "08123456789",
            "role": "administrator",
            "last_login": "2019-03-09T15:45:45Z",
            "created_at": "2019-03-09T22:45:44Z",
            "updated_at": "2019-03-09T22:45:44Z"
        },
        "code": 200
    }

### Edit User
Endpoint

    PUT: http://localhost:8024/api/v1/user/0dc21127-5d19-4a85-8809-4a9cc14b07d5

Request

    {
    	"name": "Muhammad Muflih Kholidin",
    	"email": "muflic.24@gmail.com",
    	"phone": "08123456789",
    	"role": "administrator",
    	"password": "secret"
    }
    
Response
    
    {
        "data": {
            "id": "0dc21127-5d19-4a85-8809-4a9cc14b07d5",
            "email": "muflic.24@gmail.com",
            "name": "Muhammad Muflih Kholidin",
            "phone": "08123456789",
            "role": "administrator",
            "last_login": "2019-03-09T15:45:45Z",
            "created_at": "2019-03-09T22:45:44Z",
            "updated_at": "2019-03-09T22:45:44Z"
        },
        "code": 200
    }
### Find User
Endpoint

    GET: http://localhost:8024/api/v1/user/0dc21127-5d19-4a85-8809-4a9cc14b07d5

Response
    
    {
        "data": {
            "id": "0dc21127-5d19-4a85-8809-4a9cc14b07d5",
            "email": "muflic.24@gmail.com",
            "name": "Muhammad Muflih Kholidin",
            "phone": "08123456789",
            "role": "administrator",
            "last_login": "2019-03-09T15:45:45Z",
            "created_at": "2019-03-09T22:45:44Z",
            "updated_at": "2019-03-09T22:45:44Z"
        },
        "code": 200
    }
### List User
Endpoint

    GET: http://localhost:8024/api/v1/user

Response
    
    {
        "data": [
            {
                "id": "0dc21127-5d19-4a85-8809-4a9cc14b07d5",
                "email": "muflic.24@gmail.com",
                "name": "Muhammad Muflih Kholidin",
                "phone": "08123456789",
                "role": "administrator",
                "last_login": "2019-03-09 15:45:45",
                "created_at": "2019-03-09 22:45:44",
                "updated_at": "2019-03-09 22:52:40"
            }
        ],
        "total": 1,
        "page": 1,
        "size": 1000,
        "code": 200
    }
