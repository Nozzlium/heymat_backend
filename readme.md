# Heymat Backend

Backend application for Heymat (URL coming soon) written in Go. For the frontend side, visit: https://github.com/Nozzlium/heymat_frontend

Heymat is an app where users can set budget goals and record expenses to help keep them stay within the set goals.

# Dependencies

- Go-Fiber: github.com/gofiber/fiber
- golang-jwt: github.com/golang-jwt/jwt
- golang-migrate: github.com/golang-migrate/migrate
- godotenv: github.com/joho/godotenv
- testify: github.com/stretchr/testify
- crypto: golang.org/x/crypto
- pq: github.com/lib/pq

# Running the app

Install dependencies with:

    go get

To run the app, run this command. Make sure you have your database setup:

    go run . --migrate-up

# REST API

Base URL: `http://86.38.203.29:4040` (URL coming soon)

## Registering a user

### Request

`POST /api/register`

body:

    {
    	"username": string,
    	"email": string,
    	"password": string
    }

| Key        | Type     | Description                 |
| ---------- | -------- | --------------------------- |
| `username` | `string` | **Required.** Your username |
| `email`    | `string` | **Required.** Your email    |
| `password` | `string` | **Required.** Your password |

example:

    curl -H 'content-type: application/json' -d '{ "email": "youremail@email.com", "username": "yourusername", "password": "yourpassword" }' http://86.38.203.29:4040/api/register

### Response

    {
        "code": int,
        "status": string,
        "data": {
    	    "id": number,
    	    "username": string,
    	    "email": string,
    	    "isEmailConfirmed": bool
    	}
    }

## Logging into an account

### Request

`POST /api/login`

body:

    {
        "identity": string,
        "password": string
    }

example:

    curl -H 'content-type: application/json' -d '{ "identity": "yourusername", "password": "yourpassword" }' http://86.38.203.29:4040/api/login

| Key        | Type     | Description                          |
| ---------- | -------- | ------------------------------------ |
| `identity` | `string` | **Required.** Your username or email |
| `password` | `string` | **Required.** Your password          |

### Response

    {
        "code": int,
        "status": string,
        "data": {
    	    "token": string
    	}
    }

## Creating a budgeting plan

### Request

`POST /api/budget`

body:

    {
        "title": string,
        "amount": number,
        "private": bool
    }

example:

    curl -H 'content-type: application/json' -H 'Authorization: Bearer $TOKEN' -d { "title": "yourtitle", "amount": 15000000, "private": true } http://86.38.203.29:4040/api/budget

| Key       | Type     | Description                                                                            |
| --------- | -------- | -------------------------------------------------------------------------------------- |
| `title`   | `string` | **Required.** A name to identify the plan                                              |
| `amount`  | `number` | **Required.** The amount the user is planning to stay within                           |
| `private` | `bool`   | Determines if the plan is only visible to the user who created it, defaults to `false` |
| `$TOKEN`  | `string` | **Required.** Token string generated when a user login                                 |

### Response

    {
        "code": number,
        "status": string,
        "data": {
    	    "id": number,
    	    "title": string,
    	    "user": {
    		    "id": number,
    		    "username": string
    	    },
    	    "amount": number,
    	    "amountString": string,
    	    "expense": number,
    	    "expenseString": string,
    	    "balance": number,
    	    "balanceString": string,
    	    "private": bool,
    	    "createdAt": timestamp,
    	    "createdAtString": string,
    	    "updatedAt": timestamp,
    	    "updatedAtString": string,
    	    "hasBeenEdited": bool
        }
    }

## Get a plan by its id

### Request

`GET /api/budget/:id`

example:

    curl --get -H 'Authorization: Bearer $TOKEN' http://86.38.203.29:4040/api/budget

| Key      | Type     | Description                                                                                                                                                                               |
| -------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `:id`    | `number` | **Required.** An id for a particular plan                                                                                                                                                 |
| `$TOKEN` | `number` | Token string acquired when a user login. If not provided and the plan is private or if the key provided is not from the user who created the plan, `Error 400 not found` will be returned |

### Response

    {
        "code": number,
        "status": string,
        "data": {
    	    "id": number,
    	    "title": string,
    	    "user": {
    		    "id": number,
    		    "username": string
    	    },
    	    "amount": number,
    	    "amountString": string,
    	    "expense": number,
    	    "expenseString": string,
    	    "balance": number,
    	    "balanceString": string,
    	    "private": bool,
    	    "createdAt": timestamp,
    	    "createdAtString": string,
    	    "updatedAt": timestamp,
    	    "updatedAtString": string,
    	    "hasBeenEdited": bool
        }
    }

## Get a list of plans

`GET /api/budget`

example:

    curl --get -H 'Authorization: Bearer $TOKEN' http://86.38.203.29:4040/api/budget?keyword=$KEYWORD&pageNo=$PAGE_NO&pageSize=$PAGE_SIZE

| Key          | Type     | Description                                                  |
| ------------ | -------- | ------------------------------------------------------------ |
| `$TOKEN`     | `number` | **Required.** Token string acquired when a user login.       |
| `$KEYWORD`   | `string` | To filter the plans by their titles                          |
| `$PAGE_NO`   | `number` | Pagination. Defaults to `1`                                  |
| `$PAGE_SIZE` | `number` | Determines how many records shown per page, defaults to `10` |

### Response

    {
        "code": number,
        "status": string,
        "data": {
    	    "pageNo": number,
    	    "pageSize": number,
    	    "recordCount": number,
    	    "budgetPlans": [
    		    {
    			    "id": number,
    			    "title": string,
    			    "user": {
    				    "id": number,
    				    "username": string
    			    },
    			    "amount": number,
    			    "amountString": string,
    			    "expense": number,
    			    "expenseString": string,
    			    "balance": number,
    			    "balanceString": string,
    			    "private": bool,
    			    "createdAt": timestamp,
    			    "createdAtString": string,
    			    "updatedAt": timestamp,
    			    "updatedAtString": string,
    			    "hasBeenEdited": bool
    		    }
    	    ]
        }
    }

**More features under construction!**
_check back soon_

# Error

## Generic error response

    {
        "code": number,
        "status": "error",
        "error": string
    }
