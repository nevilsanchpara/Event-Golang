
# Event App - Golang

Developed an event management application using **Golang** and **PostgreSQL**, deployed on Render and **Railway**. The app features user authentication with **JWT tokens**. To ensure proper access control, I implemented middleware to verify tokens for authorization. Users can create, update, and delete events while also being able to fetch all events or retrieve a single event by its ID. For database management, I utilized **Supabase** to handle the PostgreSQL cloud URL, allowing users to register for or cancel their participation in events seamlessly. This project showcases my ability to integrate various technologies to deliver a functional and user-friendly application.










## Authors

- [@nevilsanchpara](https://www.github.com/nevilsanchpara)


## Demo

**Deployed URL:**
- [Railway Deployment](https://nevil-golang-event.up.railway.app/events)
- [Render Deployment](https://event-golang.onrender.com/events)

**Swagger:**
- [Railway Swagger](https://nevil-golang-event.up.railway.app/swagger/index.html)
- [Render Swagger](https://event-golang.onrender.com/swagger/index.html)


**Postman Collection:**

https://www.postman.com/universal-eclipse-106264/nevil-golang-event-app/overview
## Deployment

To deploy this project run

Create Account on Railway & connect with github.Then,choose last commit to deploy. Make sure, Project is working fine on localhost as well.


## Installation

Install Eventlify with Git

```bash
go version
go mod tidy
go run main.go
```
    
## Screenshots












## API Reference

### Auth Endpoints
#### Auth - Signup

```http
  POST /signup
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Email |
| `password` | `string` | **Required** Password |

#### Auth - Login

```http
  POST /login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Email |
| `password` | `string` | **Required** Password |

### Event Endpoints

#### Create Event

```http
  POST /events [Authorization -Token Needed]
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of Event |
| `description`      | `string` | **Required**. Description of Event |
| `location`      | `string` | **Required**. Place of Event |
| `dateTime`      | `DateTime` | Timing of Event |


#### Fetch Single Event
```http
  GET /events/${eventId}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|   -      |  -  | - |


#### Fetch All events
```http
  GET /events
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|   -      |  -  | - |


#### Update Event

```http
  PUT /event/${eventId} [Authorization -Token Needed]
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of Event |
| `description`      | `string` | **Required**. Description of Event |
| `location`      | `string` | **Required**. Place of Event |

#### Delete Event

```http
  DELETE /events/${eventId} [Authorization -Token Needed]
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|   -      |  -  | - |

### Registering Event Endpoints

#### Register For specific Event

```http
  POST /events/${eventId}/register [Authorization -Token Needed]
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|   -      |  -  | - |

#### Cancel Registeration For Event


```http
  DELETE /events/${eventId}/register [Authorization -Token Needed]
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
|   -      |  -  | - |




