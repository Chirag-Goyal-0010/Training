import React, { useState } from 'react';
import {
  Box,
  Paper,
  Typography,
  Grid,
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Snackbar,
  Alert,
  CircularProgress,
} from '@mui/material';
import { styled } from '@mui/material/styles';
import { useBookings } from '../../context/BookingsContext';
import { showSuccessMessage, showErrorMessage } from '../../utils/notifications';

const Seat = styled(Box)(({ theme, selected, occupied }) => ({
  width: 40,
  height: 40,
  margin: 4,
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  cursor: occupied ? 'not-allowed' : 'pointer',
  backgroundColor: selected
    ? theme.palette.primary.main
    : occupied
    ? theme.palette.grey[400]
    : theme.palette.grey[200],
  color: selected ? theme.palette.primary.contrastText : theme.palette.text.primary,
  borderRadius: 4,
  '&:hover': {
    backgroundColor: occupied
      ? theme.palette.grey[400]
      : selected
      ? theme.palette.primary.dark
      : theme.palette.grey[300],
  },
}));

const Booking = ({ flight, onClose }) => {
  const { createBooking, loadingBookings, bookingsError } = useBookings();

  const [selectedSeat, setSelectedSeat] = useState(null);
  const [passengerInfo, setPassengerInfo] = useState({
    name: '',
    email: '',
    phone: '',
  });
  const [showConfirmation, setShowConfirmation] = useState(false);

  const seats = Array.from({ length: 60 }, (_, i) => ({
    id: i + 1,
    occupied: Math.random() > 0.7,
  }));

  const handleSeatClick = (seatId) => {
    if (seats[seatId - 1].occupied) return;
    
    setSelectedSeat(selectedSeat === seatId ? null : seatId);
  };

  const handlePassengerInfoChange = (e) => {
    setPassengerInfo({
      ...passengerInfo,
      [e.target.name]: e.target.value,
    });
  };

  const handleBooking = async () => {
    setShowConfirmation(false);

    if (selectedSeat === null) {
      showErrorMessage("Please select a seat.");
      return;
    }

    const bookingData = {
      flight_id: flight.ID,
      seat_id: selectedSeat,
    };

    const result = await createBooking(bookingData);

    if (result.success) {
      onClose();
    } else {
    }
  };

  return (
    <Dialog open={true} onClose={onClose} maxWidth="md" fullWidth>
      <DialogTitle>Book Flight</DialogTitle>
      <DialogContent>
        {loadingBookings && (
          <Box sx={{ display: 'flex', justifyContent: 'center', my: 2 }}>
            <CircularProgress />
          </Box>
        )}
        {bookingsError && (
          <Alert severity="error" sx={{ my: 2 }}>{bookingsError}</Alert>
        )}
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <Typography variant="h6" gutterBottom>
              Flight Details
            </Typography>
            <Typography>
              From: {flight?.Origin} To: {flight?.Destination}
            </Typography>
            <Typography>
              Date: {new Date(flight?.DepartureTime).toLocaleDateString()} Time: {new Date(flight?.DepartureTime).toLocaleTimeString()}
            </Typography>
            <Typography>
              Flight Code: {flight?.FlightCode} Price: ${flight?.Price}
            </Typography>
          </Grid>

          <Grid item xs={12}>
            <Typography variant="h6" gutterBottom>
              Select Seat
            </Typography>
            <Box sx={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center' }}>
              {seats.map((seat) => (
                <Seat
                  key={seat.id}
                  selected={selectedSeat === seat.id}
                  occupied={seat.occupied}
                  onClick={() => handleSeatClick(seat.id)}
                >
                  {seat.id}
                </Seat>
              ))}
            </Box>
          </Grid>

          <Grid item xs={12}>
            <Typography variant="h6" gutterBottom>
              Passenger Information
            </Typography>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField
                  fullWidth
                  label="Full Name"
                  name="name"
                  value={passengerInfo.name}
                  onChange={handlePassengerInfoChange}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Email"
                  name="email"
                  type="email"
                  value={passengerInfo.email}
                  onChange={handlePassengerInfoChange}
                />
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField
                  fullWidth
                  label="Phone"
                  name="phone"
                  value={passengerInfo.phone}
                  onChange={handlePassengerInfoChange}
                />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </DialogContent>
      <DialogActions>
        <Button onClick={onClose} disabled={loadingBookings}>Cancel</Button>
        <Button
          variant="contained"
          onClick={() => setShowConfirmation(true)}
          disabled={loadingBookings || selectedSeat === null || !passengerInfo.name || !passengerInfo.email}
        >
          Proceed to Payment
        </Button>
      </DialogActions>

      <Dialog open={showConfirmation} onClose={() => setShowConfirmation(false)}>
        <DialogTitle>Confirm Booking</DialogTitle>
        <DialogContent>
          <Typography>
            You are about to book seat {selectedSeat} for flight {flight?.ID}.
            Total amount: ${flight?.Price}
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setShowConfirmation(false)} disabled={loadingBookings}>Cancel</Button>
          <Button variant="contained" onClick={handleBooking} disabled={loadingBookings}>
            Confirm Booking
          </Button>
        </DialogActions>
      </Dialog>
    </Dialog>
  );
};

export default Booking; 