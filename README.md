# Go Events REST API

A simple **RESTful API built with Go and Gin** to manage events.
This project demonstrates how to build a backend service using:

* Go
* Gin Web Framework
* SQL database
* REST API architecture

The API allows you to **Create, Read, Update, Patch, and Delete events**.

---

# Features

* Create new events
* Fetch all events
* Fetch a single event
* Update an event (PUT)
* Partially update an event (PATCH)
* Delete an event
* JSON request validation using Gin

---

# Tech Stack

* **Go**
* **Gin Framework**
* **SQL Database**
* **REST API**

---

# Project Structure

```
go-rest/
│
├── main.go
├── db/
│   └── db.go
│
├── models/
│   └── event.go
│
├── routes/
│   ├── routes.go
│   └── events.go
│
├── go.mod
└── README.md
```

---

# Installation

### 1 Clone the repository

```bash
git clone https://github.com/your-username/go-rest.git
cd go-rest
```

---

### 2 Install dependencies

```bash
go mod tidy
```

---

### 3 Run the server

```bash
go run main.go
```

Server will start at

```
http://localhost:8080
```

---

# API Endpoints

| Method | Endpoint            | Description            |
| ------ | ------------------- | ---------------------- |
| GET    | `/events`           | Get all events         |
| GET    | `/events/:id`       | Get event by ID        |
| POST   | `/events`           | Create event           |
| PUT    | `/events/:id`       | Update event           |
| PATCH  | `/patch/events/:id` | Partially update event |
| DELETE | `/events/:id`       | Delete event           |

---

# Request Examples

## Create Event

`POST /events`

```json
{
  "name": "Flutter Conference",
  "description": "Flutter developers meetup",
  "location": "Kochi",
  "dateTime": "2026-03-03T15:30:00Z"
}
```

Response

```json
{
  "message": "Success",
  "status_code": 201,
  "data": {
    "id": 1,
    "name": "Flutter Conference",
    "description": "Flutter developers meetup",
    "location": "Kochi",
    "dateTime": "2026-03-03T15:30:00Z"
  }
}
```

---

# Get All Events

`GET /events`

Example Response

```json
{
  "message": "Success",
  "status_code": 200,
  "data": [
    {
      "id": 1,
      "name": "Flutter Meetup",
      "description": "Developers event",
      "location": "Kochi",
      "dateTime": "2026-03-03T15:30:00Z"
    }
  ]
}
```

---

# Get Single Event

`GET /events/:id`

Example

```
GET /events/1
```

---

# Update Event

`PUT /events/:id`

Request

```json
{
  "name": "Updated Event",
  "description": "Updated description",
  "location": "Trivandrum",
  "dateTime": "2026-03-05T10:00:00Z"
}
```

---

# Patch Event

`PATCH /patch/events/:id`

Request

```json
{
  "location": "Bangalore"
}
```

Only the fields provided will be updated.

---

# Delete Event

`DELETE /events/:id`

Example

```
DELETE /events/1
```

Response

```json
{
  "message": "Event Deleted successfully"
}
```

---