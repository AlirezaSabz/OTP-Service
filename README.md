# OTP Service (Golang)

A backend service written in Go that provides OTP-based login and registration with JWT authentication, PostgreSQL persistence, Redis caching, and basic user management.

---

##  Features
- OTP-based login & registration  
- Rate limiting
- Simple user management (list, search, pagination)  
- PostgreSQL database (user storage)  
- Redis for OTP & rate-limit caching  
- Docker & Docker Compose support
- JWT (just create jwt ,for future we will add VerifyJwt middleware )

---

##  How to Run

### 1. Run Locally
Make sure you have:
- Go 1.23+
- PostgreSQL
- Redis

Create .env file 

Run the app 
go run main.go

### 2. Run with Docker
Create .env file          
Run containers: docker-compose up --build


App → http://localhost:8080

Postgres → localhost:5432

Redis → localhost:6379


## API Endpoints

Note: For simplicity in this example, the phone number and OTP are sent as query parameters instead of in the request body

- POST /request-otp?phone=12345"
- POST /verify-otp?phone=12345&otp=123456"
- GET /user?phone=12345"
- GET /users?page=1&limit=5&search=123"

