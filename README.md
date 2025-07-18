# MindenAirport

A full-stack airport management system built with Go (backend) and React + Electron (frontend).

## 📋 Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Environment Setup](#environment-setup)
- [Backend Setup](#backend-setup)
- [Frontend Setup](#frontend-setup)
- [Running the Application](#running-the-application)
- [Docker Setup](#docker-setup)
- [Database Setup](#database-setup)
- [Development](#development)
- [Building for Production](#building-for-production)

## 🚀 Prerequisites

Before running this project, make sure you have the following installed:

### Backend Requirements
- **Go 1.23.4** or later - [Download Go](https://golang.org/dl/)
- **Oracle Database** - The application uses Oracle DB with godror driver

### Frontend Requirements
- **Node.js 18+** - [Download Node.js](https://nodejs.org/)
- **Bun** - Package manager (recommended) - [Install Bun](https://bun.sh/)
  ```bash
  # Install Bun (Windows)
  powershell -c "irm bun.sh/install.ps1 | iex"
  ```

### Docker Requirements (Optional)
- **Docker Desktop** - [Download Docker Desktop](https://www.docker.com/products/docker-desktop/)
- **Docker Compose** - Usually included with Docker Desktop

## 📁 Project Structure

```
MindenAirport/
├── backend/          # Go backend API
├── frontend/         # React + Electron frontend
├── docs/            # Documentation files
├── imgs/            # Images and diagrams
└── scripts/         # Database scripts
```

## ⚙️ Environment Setup

### Backend Environment

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Copy the environment template:
   ```bash
   cp .env.template .env
   ```

3. Edit `.env` file with your Oracle database credentials:
   ```bash
   CONNECTIONSTRING="your_username/your_password@ORCL"
   JWT_SECRET="your_jwt_secret_key"
   ```

### Frontend Environment

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

## 🔧 Backend Setup

1. **Navigate to backend directory:**
   ```bash
   cd backend
   ```

2. **Install Go dependencies:**
   ```bash
   go mod download
   ```

3. **Verify installation:**
   ```bash
   go mod verify
   ```

4. **Run the backend server:**
   ```bash
   go run main.go
   ```

The backend server will start on `http://localhost:8080`

### Backend Dependencies

- **Gin** - HTTP web framework
- **CORS** - Cross-Origin Resource Sharing
- **Oracle Driver** - Database connectivity
- **JWT** - Authentication
- **Crypto** - Password hashing
- **UUID** - Unique identifier generation

## 🎨 Frontend Setup

1. **Navigate to frontend directory:**
   ```bash
   cd frontend
   ```

2. **Install dependencies using Bun:**
   ```bash
   bun install
   ```

   Or using npm:
   ```bash
   npm install
   ```

3. **Start the development server:**
   ```bash
   bun run dev
   ```

   Or using npm:
   ```bash
   npm run dev
   ```

### Frontend Technologies

- **React 19** - UI framework
- **TypeScript** - Type safety
- **Tailwind CSS** - Styling
- **Electron** - Desktop application
- **Vite** - Build tool
- **React Router** - Navigation

## 🚀 Running the Application

### Development Mode

1. **Start the backend:**
   ```bash
   cd backend
   go run main.go
   ```

2. **Start the frontend (in a new terminal):**
   ```bash
   cd frontend
   bun run dev
   ```

3. **Access the application:**
   - The Electron app will launch automatically
   - Backend API: `http://localhost:8080`
   - Frontend dev server: `http://localhost:3000`

### Available Scripts (Frontend)

```bash
# Development
bun run dev              # Start Electron development server

# Building
bun run build:web        # Build for web
bun run build:electron   # Build Electron app

# Linting
bun run lint            # Run ESLint

# Preview
bun run preview         # Preview production build
```

### Available Commands (Backend)

```bash
# Development
go run main.go          # Start development server
go build               # Build binary

# Dependencies
go mod download        # Download dependencies
go mod tidy           # Clean up dependencies
```

## 🐳 Docker Setup

The project includes Docker configurations for deployment. The backend uses the devcontainer setup with Oracle Instant Client and connects to an external Oracle database via tnsnames.ora configuration.

### Prerequisites for Docker Setup

- **Docker Desktop** - [Download Docker Desktop](https://www.docker.com/products/docker-desktop/)
- **Database credentials** - Username and password for the Oracle database

### Quick Start with Docker

1. **Clone the repository and navigate to the project root:**
   ```bash
   git clone <repository-url>
   cd MindenAirport
   ```

2. **Configure your database credentials:**
   ```bash
   # Edit enviorment variables in docker compose file with your Oracle database credentials:
   # CONNECTIONSTRING="your_username/your_password@ORCL"
   # JWT_SECRET="your_jwt_secret_here"
   ```

3. **Run the setup script**

   **Or run manually:**
   ```bash
   docker-compose up -d
   ```

   This will start:
   - Backend API on port `8080`
   - Frontend web app on port `3000`

4. **Access the application:**
   - Frontend: `http://localhost:3000`
   - Backend API: `http://localhost:8080`


### Docker Services

| Service | Container Name | Port | Description |
|---------|---------------|------|-------------|
| `backend` | `minden_backend` | 8080 | Go API Server with Oracle Client |
| `frontend` | `minden_frontend` | 3000 | React Web App |


### Environment Variables for Docker

The Docker setup uses the following environment variables:

**Backend:**
- `CONNECTIONSTRING=your_username/your_password@ORCL`
- `JWT_SECRET=your_jwt_secret_here_change_in_production`

### Docker Troubleshooting

**Common Docker Issues:**

1. **Port conflicts:**
   ```bash
   # Check what's using the ports
   netstat -tulpn | grep :8080
   netstat -tulpn | grep :3000
   netstat -tulpn | grep :1521
   ```

## 📦 Building for Production

### Backend Production Build

```bash
cd backend
go build -o mindenairport.exe main.go
```

### Frontend Production Build

```bash
cd frontend

# For web deployment
bun run build:web

# For Electron desktop app
bun run build:electron
```

## 🔧 Troubleshooting

### Common Issues

1. **Oracle Driver Issues:**
   - Ensure Oracle Instant Client is installed
   - Verify `CONNECTIONSTRING` format in `.env`

2. **Port Conflicts:**
   - Backend runs on port 8080
   - Frontend dev server on port 5173
   - Docker frontend on port 3000
   - Check if ports are available

3. **Environment Variables:**
   - Ensure enviroment variables are properly configured
   - Verify Oracle database credentials are correct

4. **Dependencies:**
   - Run `go mod download` for backend
   - Run `bun install` for frontend

5. **Docker Issues:**
   - Ensure Docker Desktop is running
   - Use `docker-compose logs [service-name]` to check logs
   - Try rebuilding with `--no-cache` flag if builds fail
   - If frontend build fails with lockfile errors, the dependencies will be resolved automatically

6. **Database Connection:**
   - Ensure to in HSBI network via Cisco VPN
