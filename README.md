# Cerebro: API Gateway for Microservices

Cerebro is a Golang project focused on developing an API gateway, connected to microservices using gRPC connections. It acts as a central entry point for client requests, providing functionalities such as routing, authentication, and load balancing to the underlying microservices.

## Features

- API routing and forwarding
- Authentication and authorization
- Load balancing and service discovery
- Logging and monitoring

## Frameworks & Libraries

The project relies on the following frameworks and libraries:

- [Fiber v3](https://github.com/gofiber/fiber/v3) v3.0.0-beta.2
- [Logrus](https://github.com/sirupsen/logrus) v1.9.3
- [Viper](https://github.com/spf13/viper) v1.18.2
- [gRPC](https://google.golang.org/grpc) v1.63.2
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf) v1.33.0


## Usage

### Building Docker Images

To build both Cerebro and CustomerService for test there already is a docker compose yaml file.

1️⃣ Clone the Cerebro repository:

   ```bash
   git clone https://github.com/HouseCham/cerebro.git
   ```

2️⃣ Navigate to the project directory:
    ```bash
    cd cerebro
    ```

3️⃣ Be sure to create a .env file at the root of the repo:
    
```ini
# 📌 Cerebro (API Gateway)
CEREBRO_NAME=cerebro
CEREBRO_PORT=3000
CEREBRO_HOST=localhost
JWT_SECRET=my-random-secret

# 📌 Customer Service
CUSTOMER_SERVICE_NAME=customer-service
CUSTOMER_SERVICE_HOST=localhost
CUSTOMER_SERVICE_PORT=3001

# 📌 Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=my-db-password
DB_NAME=cerebroDB
```

4️⃣ **Run the Services with Docker Compose**
   ```bash
   docker compose up -d
   ```
   - The `-d` flag runs the containers in the background.
   - To see logs in real-time, remove `-d`:
     ```bash
     docker compose up
     ```

---

## 🛠 **Testing the gRPC Web API**

Once the services are running, you can use **Postman**, **cURL**, or any API testing tool to interact with the system.

### 📌 **Create a New Customer**
- **Endpoint:** `POST http://localhost:3000/api/v1/customer`
- **Request Body:**
  ```json
  {
      "id": 0,
      "first_name": "John",
      "second_name": "",
      "last_name_p": "Doe",
      "last_name_m": "Doe2",
      "phone_number": "0000000000",
      "email": "john.doe@example.com",
      "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
      "data": 1,
      "hasError": false,
      "errorMessage": "",
      "statusCode": 200
  }
  ```
  - The `data` field represents the newly created customer's ID in the database.

---

### 📌 **Fetch Customer Information**
- **Endpoint:** `GET http://localhost:3000/api/v1/customer/{userId}`
- **Example Response:**
  ```json
  {
      "data": {
          "id": 1,
          "firstName": "John",
          "secondName": "",
          "lastNameP": "Doe",
          "lastNameM": "Doe2",
          "phoneNumber": "0000000000",
          "email": "john.doe@example.com",
          "passwordHash": "",
          "HashedPassword": ""
      },
      "hasError": false,
      "errorMessage": "",
      "statusCode": 200
  }
  ```

---

## 🛑 **Stopping and Cleaning Up**
To stop the services, run:
```bash
docker compose down
```
To remove all containers, networks, and volumes:
```bash
docker compose down -v
```