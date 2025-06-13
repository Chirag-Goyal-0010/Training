import React, { useEffect } from 'react';
import {
  Box,
  Container,
  Typography,
  Grid,
  Card,
  CardContent,
  CircularProgress,
  Alert,
} from '@mui/material';
import { format } from 'date-fns';
import { useBookings } from '../../context/BookingsContext';

const Dashboard = () => {
  const { bookings, loadingBookings, bookingsError, fetchUserBookings } = useBookings();

  useEffect(() => {
    fetchUserBookings();
  }, [fetchUserBookings]);

  if (loadingBookings) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4, textAlign: 'center' }}>
        <CircularProgress />
        <Typography sx={{ mt: 2 }}>Loading your bookings...</Typography>
      </Container>
    );
  }

  if (bookingsError) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Alert severity="error">Error: {bookingsError}</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" gutterBottom>
        My Bookings
      </Typography>

      {bookings.length === 0 ? (
        <Typography variant="h6" align="center" sx={{ mt: 4 }}>
          You have no bookings yet.
        </Typography>
      ) : (
        <Grid container spacing={3}>
          {bookings.map((booking) => (
            <Grid item xs={12} md={6} lg={4} key={booking.ID}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    Flight {booking.Flight.FlightCode} ({booking.Flight.Origin} → {booking.Flight.Destination})
                  </Typography>
                  <Typography color="textSecondary">
                    Departure: {format(new Date(booking.Flight.DepartureTime), 'PPP p')}
                  </Typography>
                  <Typography color="textSecondary">
                    Arrival: {format(new Date(booking.Flight.ArrivalTime), 'PPP p')}
                  </Typography>
                  <Typography color="textSecondary">
                    Seat: {booking.Seat.SeatNumber || booking.Seat.ID} | Status: {booking.Status}
                  </Typography>
                  <Typography variant="h6" color="primary" sx={{ mt: 2 }}>
                    Price: ${booking.Flight.Price}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          ))}
        </Grid>
      )}
    </Container>
  );
};

export default Dashboard; 