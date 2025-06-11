# Flight Management System

A full-stack flight management system built with React, Go, and PostgreSQL.

## Features

- User authentication (Admin and Customer roles)
- Flight management (Add, Edit, Delete flights)
- Flight search and booking
- JWT-based authentication
- Responsive Material-UI design

## Prerequisites

- Go 1.16 or higher
- Node.js 14 or higher
- PostgreSQL 12 or higher
- npm or yarn

## Setup Instructions

### Database Setup

1. Create a PostgreSQL database named `flights_db`
2. Import the database schema:
   ```bash
   psql -d flights_db -f Flights_database/Flights_complete.sql
   ```

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install Go dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   Create a `.env` file in the backend directory with the following content:
   ```
   DATABASE_URL=postgres://username:password@localhost/flights_db?sslmode=disable
   JWT_SECRET=your-secret-key
   ```

4. Run the backend server:
   ```bash
   go run main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm start
   ```

## Usage

1. Open your browser and navigate to `http://localhost:3000`
2. Register a new account or login with existing credentials
3. Admin users can:
   - Add new flights
   - Edit existing flights
   - Delete flights
4. Regular users can:
   - Search for flights
   - Book flights
   - View their bookings

## API Endpoints

### Authentication
- POST `/api/auth/register` - Register a new user
- POST `/api/auth/login` - Login user

### Flights
- GET `/api/flights` - Get all flights
- POST `/api/admin/flights` - Add new flight (Admin only)
- PUT `/api/admin/flights/:id` - Update flight (Admin only)
- DELETE `/api/admin/flights/:id` - Delete flight (Admin only)

## Security

- JWT-based authentication
- Password hashing using bcrypt
- Role-based access control
- CORS enabled
- Input validation

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request 