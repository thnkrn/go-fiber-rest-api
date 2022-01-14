# GO Restful API service with Fiber framework and Gorm SQLite

## Template Structure

- [Fiber](https://github.com/gofiber) is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.
- [GOFiber JWT](https://github.com/gofiber/jw) JWT returns a JSON Web Token (JWT) auth middleware. For valid token, it sets the user in Ctx.Locals and calls next handler. For invalid token, it returns "401 - Unauthorized" error. For missing token, it returns "400 - Bad Request" error.
- [GORM](https://gorm.io/index.html) with [SQLite](https://gorm.io/docs/connecting_to_the_database.html#SQLite)The fantastic ORM library for Golang aims to be developer friendly.
- [Fiber Logger](https://docs.gofiber.io/api/middleware/logger)

## Available Endpoint

In the project directory, you can run:

### `POST /login`

For generating a JWT

### `GET /users`

For getting all of users data

### `GET /users/:id`

For getting specific user data

### `POST /user`

For creating new user data

### `DELETE /user/:id`

For removing existing user data

### `PUT /user/:id`

For updating existing user data
