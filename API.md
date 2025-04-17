# API

### Login Endpoint

```http
POST http://localhost:8001/login
```


This endpoint allows users to log in with their credentials.

#### Request Body

- `login` (string) - The user account login.
    
- `password` (string) - The password for the user's account.
    

#### Response

The response is in JSON format with the following schema:

``` json
{
    "type": "object",
    "properties": {
        "login": {
            "type": "string"
        },
        "uuid": {
            "type": "string"
        }
    }
}

 ```

- `login` (string) - The username or email of the logged-in user.
    
- `uuid` (string) - The unique identifier for the user session.

---


### Register Endpoint

```http
POST http://localhost:8001/register
```

This endpoint is used to register a new user with their credentials.

#### Request Body

- `login` (string) - The username or email of the user.
    
- `password` (string) - The password for the user's account.
    

#### Response

The response is in JSON format with the following schema:

``` json
{
    "type": "object",
    "properties": {
        "login": {
            "type": "string"
        },
        "uuid": {
            "type": "string"
        }
    }
}

 ```

- `login` (string) - The username or email of the logged-in user.
    
- `uuid` (string) - The unique identifier for the user session.
