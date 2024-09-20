# Golang Backend GOLANG - POSTGRESQL ON NEON.TECH , AZURE , GORM , FIBER , DOCKER ,GODOTENV AND COMMAND TO RUN AIR.

## Description

This is a simple Golang backend application that connects to a PostgreSQL database on Neon.Tech. The application uses GORM as the ORM and Fiber as the web framework. The application is containerized using Docker and uses godotenv to manage environment variables. The application is also configured to use Air for live reloading.

## Installation

1. Clone the repository
2. Create a `.env` file in the root directory and add the following environment variables:

```
DB_USER=your_neon_tech_db_username
DB_PASSWORD=your_neon_tech_db_password
DB_NAME=your_neon_tech_db_name
DB_HOST=your_neon_tech_db_host
DB_PORT=your_neon_tech_db_port
```

3. Run the following command to start the application:

```bash
air
```

## Docker

```
docker-compose up
```

4. The application should now be running on `http://localhost:3000`

## API Endpoints

1. GET `/api/users` - Get all users
2. GET `/api/users/:id` - Get a user by ID
3. POST `/api/users` - Create a new user
4. PUT `/api/users/:id` - Update a user by ID
5. DELETE `/api/users/:id` - Delete a user by ID
6. POST `/api/users/login` - Login a user
7. POST `/api/users/register` - Register a new user

## Azure

```
url : https://golang-backend-cxhte0gkang7e8by.eastus-01.azurewebsites.net/
```
