import React, { useEffect } from 'react';
import { useBookings } from '../../context/BookingsContext.tsx';
import { useAuth } from '../../context/AuthContext.tsx';
import { Grid, Paper, Typography, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { IBooking, IFlight } from '../../api';
import { useFlights } from '../../context/FlightsContext.tsx';

const Dashboard: React.FC = () => {
  const { user } = useAuth();
  const { bookings, loadingBookings, bookingsError, fetchUserBookings } = useBookings();
  const { flights, fetchFlights } = useFlights();
  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      fetchUserBookings(user.id);
    }
  }, [user, fetchUserBookings]);

  useEffect(() => {
    if (bookings.length > 0 && flights.length === 0) {
      fetchFlights();
    }
  }, [bookings, flights, fetchFlights]);

  const handleBookFlight = () => {
    navigate('/book-flight');
  };

  if (loadingBookings) {
    return <Typography>Loading your bookings...</Typography>;
  }

  if (bookingsError) {
    return <Typography color="error">{bookingsError}</Typography>;
  }

  const getFlightDetails = (flightId: number): IFlight | undefined => {
    return flights.find(flight => flight.id === flightId);
  };

  return (
    <Grid container spacing={3}>
      <Grid item xs={12}>
        <Paper elevation={3} sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom>
            Welcome, {user?.username}!
          </Typography>
          <Button
            variant="contained"
            color="primary"
            onClick={handleBookFlight}
            sx={{ mb: 3 }}
          >
            Book a New Flight
          </Button>

          <Typography variant="h6" gutterBottom>
            Your Bookings
          </Typography>

          {bookings.length === 0 ? (
            <Typography>You have no bookings yet.</Typography>
          ) : (
            <Grid container spacing={2}>
              {bookings.map((booking: IBooking) => {
                const flight = getFlightDetails(booking.flightId);
                return (
                  <Grid item xs={12} key={booking.id}>
                    <Paper elevation={2} sx={{ p: 2 }}>
                      <Typography variant="subtitle1">
                        Flight: {flight?.flightCode || 'N/A'}
                      </Typography>
                      <Typography>
                        Route: {flight?.origin || 'N/A'} → {flight?.destination || 'N/A'}
                      </Typography>
                      <Typography>
                        Departure: {flight ? new Date(flight.departureTime).toLocaleString() : 'N/A'}
                      </Typography>
                      <Typography>
                        Arrival: {flight ? new Date(flight.arrivalTime).toLocaleString() : 'N/A'}
                      </Typography>
                      <Typography>
                        Seat: {booking.seatNumber}
                      </Typography>
                      <Typography>
                        Status: {booking.bookingStatus}
                      </Typography>
                    </Paper>
                  </Grid>
                );
              })}
            </Grid>
          )}
        </Paper>
      </Grid>
    </Grid>
  );
};

export default Dashboard; 