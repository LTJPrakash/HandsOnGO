# GOlang use as a backend

## Overview

This project developed with **Go (Golang)** for the backend . The backend provides RESTful APIs for user authentication, CRUD operations for products, and Google OAuth2 sign-in. The project also integrates **Swagger** for API documentation.

## Features

- **User Authentication**: Standard sign-in and sign-up functionality.
- **Google Sign-In**: Google OAuth2 login support.
- **CRUD Operations for Products**: Ability to create, read, update, and delete products.
- **Swagger Integration**: Automatically generated API documentation.
- **MongoDB**: MongoDB is used for data storage (users and products).

---

## Setup Instructions

### 1. Install Dependencies

```bash
go mod tidy
```
### 2. Set Environment Variables

```bash
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_REDIRECT_URL=http://localhost:5000/auth/google/callback
MONGO_URI=mongodb://localhost:27017/ecommerce
```

### 3. Run the Go Backend

```bash
go run main.go
```

### 4. Swagger Generation:

```bash
swag init
```

### 5. Access Swagger API Documentation

```bash
http://localhost:5000/swagger/index.html
```

## Project Structure

```bash
backend/
├── controller/
│   └── auth.go                # Handles user authentication logic
│   └── product.go             # Handles product CRUD operations
├── model/
│   └── user.go                # DB structure for user
│   └── product.go             # DB structure for product
├── router/
│   └── router.go              # Defines API routes
├── utils/
│   └── auth.go                # Utility functions for authentication
│   └── responseManager.go     # Utility functions for success/error responses
│   └── config.go              # Google credentials setup
├── main.go                    # Main entry point for the Go backend
├── .env                       # Environment variables (not committed to Git)
└── README.md                  # Documentation for the project
