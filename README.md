# Indian Flights Booking System

A full-stack web application for domestic flight bookings in India, built with Go (backend), React (frontend), and PostgreSQL (database).

## Features

- User Authentication (Customer & Admin)
- Flight Management (Admin)
- Flight Booking System (Customer)
- Domestic Indian Routes
- Real-time Seat Availability
- Booking Management

## Tech Stack

### Backend
- Go
- Gin Web Framework
- GORM (ORM)
- JWT Authentication
- PostgreSQL

### Frontend
- React
- Material-UI
- Axios
- React Router
- Date-fns

## Prerequisites

1. Go (1.21 or later)
2. Node.js (14.x or later)
3. PostgreSQL (12 or later)
4. npm or yarn

## Setup Instructions

### Database Setup

1. Create a PostgreSQL database:
```sql
CREATE DATABASE flights_booking;
```

### Backend Setup

1. Navigate to the project root directory:
```bash
cd flights-booking
```

2. Install Go dependencies:
```bash
go mod download
```

3. Set up environment variables:
- Copy `.env.example` to `.env`
- Update the database connection string and other configurations

4. Run the backend server:
```bash
go run .
```

The server will start at `http://localhost:8080`

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

The frontend will be available at `http://localhost:3000`

## API Endpoints

### Authentication
- POST `/api/register` - User registration
- POST `/api/login` - User login

### Flights
- GET `/api/flights` - List all flights
- POST `/api/admin/flights` - Add new flight (Admin only)
- DELETE `/api/admin/flights/:id` - Delete flight (Admin only)

### Bookings
- POST `/api/bookings` - Create new booking
- GET `/api/bookings` - Get user's bookings

## Default Admin Account

To create an admin account, manually update the `is_admin` field in the database for a registered user:

```sql
UPDATE users SET is_admin = true WHERE email = 'admin@example.com';
```

## Security Considerations

1. Change the JWT secret key in production
2. Use HTTPS in production
3. Implement rate limiting
4. Sanitize user inputs
5. Regular security updates

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request