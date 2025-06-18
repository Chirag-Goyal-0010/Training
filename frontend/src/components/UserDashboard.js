import React, { useState, useEffect } from 'react';
import axios from 'axios';
import {
  Container,
  Paper,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Alert,
  Box,
  Chip,
} from '@mui/material';
import { format } from 'date-fns';

function UserDashboard() {
  const [bookings, setBookings] = useState([]);
  const [error, setError] = useState('');

  useEffect(() => {
    fetchBookings();
  }, []);

  const fetchBookings = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await axios.get('http://localhost:8080/api/bookings', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setBookings(response.data.data);
    } catch (error) {
      setError('Failed to fetch bookings');
    }
  };

  const getStatusColor = (status) => {
    switch (status.toLowerCase()) {
      case 'confirmed':
        return 'success';
      case 'cancelled':
        return 'error';
      default:
        return 'default';
    }
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        My Bookings
      </Typography>

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}

      {bookings.length === 0 ? (
        <Box sx={{ mt: 4, textAlign: 'center' }}>
          <Typography variant="h6" color="text.secondary">
            You haven't made any bookings yet.
          </Typography>
        </Box>
      ) : (
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Booking ID</TableCell>
                <TableCell>Flight Number</TableCell>
                <TableCell>Route</TableCell>
                <TableCell>Travel Date</TableCell>
                <TableCell>Seats</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Booking Date</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bookings.map((booking) => {
                console.log("Booking with Flight Data:", booking);
                return (
                <TableRow key={booking.ID}>
                  <TableCell>#{booking.ID}</TableCell>
                  <TableCell>{booking.Flight?.ID}</TableCell>
                  <TableCell>
                    {booking.Flight?.origin} → {booking.Flight?.destination}
                  </TableCell>
                  <TableCell>
                    {booking.Flight?.departure_time ? format(new Date(booking.Flight.departure_time), 'dd/MM/yyyy HH:mm') : 'N/A'}
                  </TableCell>
                  <TableCell>{booking.seats}</TableCell>
                  <TableCell>
                    <Chip
                      label={booking.status}
                      color={getStatusColor(booking.status)}
                      size="small"
                    />
                  </TableCell>
                  <TableCell>
                    {format(new Date(booking.booking_date), 'dd/MM/yyyy HH:mm')}
                  </TableCell>
                </TableRow>
              )})}
            </TableBody>
          </Table>
        </TableContainer>
      )}
    </Container>
  );
}

export default UserDashboard;