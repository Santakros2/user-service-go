# Users Microservice

This is the **Users Microservice** for your microservices-based project.  
It is responsible for user management, authentication, and user profile handling.

Repository: https://github.com/Santakros2/user-service-go.git

---

## 🚀 Table of Contents

- [About](#about)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Service](#running-the-service)
- [API Endpoints](#api-endpoints)
- [Database](#database)
- [Testing](#testing)
- [Logging & Monitoring](#logging--monitoring)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

---

## 📌 About

The **Users Microservice** handles all user-related operations including:

- User registration
- User login & authentication
- User profile management
- Password hashing & validation
- JWT token issuance

It is designed to be a standalone, scalable microservice in your distributed architecture.

---

## ⚙️ Features

✔ RESTful API endpoints  
✔ Secure password hashing  
✔ JWT-based authentication  
✔ Middleware for validation & auth  
✔ Clean architecture with separation of concerns  
✔ Docker support  

---

## 🛠 Tech Stack

| Layer | Technology |
|-------|------------|
| Language | Golang |
| Web Framework | net/http, Gorilla Mux (or Echo/Fiber) |
| Database | PostgreSQL / MongoDB (replace as applicable) |
| ORM | GORM (or native driver) |
| Authentication | JWT |
| Testing | Go `testing` |
| Containerization | Docker & Docker Compose |

---

## 🧠 Architecture

