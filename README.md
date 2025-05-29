# 📂 File Sharing Backend (Go + PostgreSQL + Redis + AWS S3)

A secure and efficient backend for file sharing, supporting **user authentication, file upload/download, search, Redis caching, PostgreSQL**, and **automated file deletion**.

---

## 🚀 Features

✅ User Authentication (JWT)  
✅ File Upload & Storage (AWS S3)  
✅ File Search (Name, Type, Upload Date, User)  
✅ Redis Caching for Performance  
✅ PostgreSQL for File Metadata  
✅ Background Job for Expired File Deletion  

---

## ⚙️ Setup Instructions

### 1️⃣ Clone the Repository
```sh
https://github.com/ujjawalpatidar4/File-Sharing-Backend.git
cd File-Sharing-Backend
```

### 2️⃣ Configure Environment Variables
Create a .env file and add the following:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
```

### 3️⃣ Install Dependencies
```
go mod tidy
```

### 4️⃣ Run Migrations
```
go run migrations/migrate.go
```

### 5️⃣ Start the Server
```
go run cmd/main.go
```
The server will start on: http://localhost:8080

## 📌 API Endpoints

### 🔐 Authentication
| Method | Endpoint         | Description          |
|--------|-----------------|----------------------|
| POST   | `/auth/register` | Register a new user |
| POST   | `/auth/login`    | Login & get JWT     |

### 📂 File Management
| Method  | Endpoint                      | Description                     |
|---------|--------------------------------|---------------------------------|
| POST    | `/files/upload`               | Upload a file                   |
| GET     | `/files/list`                 | List all uploaded files         |
| GET     | `/files/download/:filename`   | Download a specific file        |
| DELETE  | `/files/delete/:filename`     | Delete a specific file          |
| GET     | `/files/search`               | Search files by name/type/date  |
| DELETE  | `/files/delete-expired`       | Delete expired files            |

## 🎉 Screenshots

### 🟢 1. User Registration
- **Endpoint:** `POST /auth/register`
- **Description:** Registers a new user.
- **Screenshot:**
  
  ![Register API](screenshots/Screenshot(40).png)

---

### 🟢 2. User Login
- **Endpoint:** `POST /auth/login`
- **Description:** Logs in an existing user and returns a JWT token.
- **Screenshot:**
  
  ![Login API](screenshots/Screenshot(41).png)

---

### 🟢 3. Upload File
- **Endpoint:** `POST /files/upload`
- **Description:** Uploads a file and saves metadata.
- **Screenshot:**
  
  ![Upload API](screenshots/Screenshot(42).png)

---

### 🟢 4. List Files
- **Endpoint:** `GET /files/list`
- **Description:** Retrieves a list of uploaded files.
- **Screenshot:**
  
  ![List Files API](screenshots/Screenshot(43).png)

---

### 🟢 5. Search Files
- **Endpoint:** `GET /files/search`
- **Description:** Searches for files based on criteria.
- **Screenshot:**
  
  ![Search Files API](screenshots/Screenshot(44).png)

---

### 🟢 6. Delete File
- **Endpoint:** `DELETE /files/delete/:filename`
- **Description:** Deletes a specific file.
- **Screenshot:**
  
  ![Delete File API](screenshots/Screenshot(45).png)

---

### 🟢 7. Download File
- **Endpoint:** `GET /files/download/:filename`
- **Description:** Downloads a specific file.
- **Screenshot:**
  
  ![Download File API](screenshots/Screenshot(46).png)

---

### 🟢 8. Delete Expired Files (Background Job)
- **Endpoint:** `DELETE /files/delete-expired`
- **Description:** Deletes all expired files and removes metadata from the database.
- **Screenshot:**
  
  ![Delete Expired API](screenshots/Screenshot(47).png)

---
