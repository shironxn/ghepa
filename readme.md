
# GOLANG HEXAGONAL EVENT PLANNING APP
![Logo](https://mir-s3-cdn-cf.behance.net/project_modules/hd/e7d2bd61228185.5a67a07360e75.gif)
GHEPA (Golang Hexagonal Event Planning App) is an event planning application built using the Go (Golang) programming language with a hexagonal architecture approach. This application is designed to assist users in planning, creating, and managing their events easily and efficiently.

## Structure Project

```bash
├── cmd
│   └── main.go
├── config
│   └── config.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── adapter
│   │   ├── handler
│   │   │   ├── comment.go
│   │   │   ├── event.go
│   │   │   └── user.go
│   │   └── repository
│   │       ├── comment.go
│   │       ├── event.go
│   │       └── user.go
│   ├── core
│   │   ├── domain
│   │   │   ├── claims.go
│   │   │   ├── comment.go
│   │   │   ├── event.go
│   │   │   ├── participant.go
│   │   │   ├── response.go
│   │   │   └── user.go
│   │   ├── port
│   │   │   ├── comment.go
│   │   │   ├── event.go
│   │   │   ├── response.go
│   │   │   └── user.go
│   │   └── service
│   │       ├── comment.go
│   │       ├── event.go
│   │       └── user.go
│   ├── middleware
│   │   └── auth.go
│   ├── route
│   │   └── route.go
│   └── util
│       ├── bcrypt.go
│       ├── jwt.go
│       ├── response.go
│       └── validate.go
└── README.md
```
## Installation

1. Clone this repository
    ```bash
    git clone <repository URL>
    ```

2. Navigate to the project directory
    ```bash
    cd ghepa
    ```

3. Copy or rename `.env.example` to `.env`

4. Install dependencies
    ```bash
    go mod tidy
    ```
## 1.1 Usage With Golang

1. First, navigate to the cmd directory:
    ```bash
    cd cmd
    ```


2. Next, run the main.go file:
    ```go
    go run main.go
    ```

## 1.2 Usage With Docker
    not done yet

## API Reference

### Auth

| Method | Url                      | Description        |
| :----- | :----------------------- | :----------------- |
| `POST` | `/api/v1/auth/login`     | Login account.     |
| `POST` | `/api/v1/auth/register`  | Create an account. |

### User

| Method   | Url                  | Description                |
| :------- | :------------------- | :------------------------- |
| `GET`    | `/api/v1/user`       | Get all user data.        |
| `GET`    | `/api/v1/user/{id}`  | Get user data by ID.      |
| `PUT`    | `/api/v1/user/{id}`  | Update user data by ID.   |
| `DELETE` | `/api/v1/user/{id}`  | Delete user data by ID.   |

### Event

| Method   | Url                   | Description                   |
| :------- | :-------------------- | :---------------------------- |
| `POST`   | `/api/v1/event`       | Create a new event.           |
| `GET`    | `/api/v1/event`       | Get all events.                |
| `GET`    | `/api/v1/event/{id}`  | Get event details by ID.      |
| `PUT`    | `/api/v1/event/{id}`  | Update event details by ID.   |
| `DELETE` | `/api/v1/event/{id}`  | Delete event by ID.           |
| `POST`   | `/api/v1/event/{id}/join` | Join an event by ID.       |

### Comment

| Method   | Url                     | Description                  |
| :------- | :---------------------- | :--------------------------- |
| `POST`   | `/api/v1/comment/{id}`  | Create a comment for an event. |
| `GET`    | `/api/v1/comment`       | Get all comments.            |

[![forthebadge](https://forthebadge.com/images/featured/featured-built-with-love.svg)](https://forthebadge.com)
