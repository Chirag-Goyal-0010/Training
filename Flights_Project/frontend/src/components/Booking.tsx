import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useBookings } from '../../context/BookingsContext.tsx';
import { useFlights } from '../../context/FlightsContext.tsx';
import { useAuth } from '../../context/AuthContext.tsx';
import { IFlight } from '../../api';
import {
  Paper,
  Typography,
  Grid,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Box,
} from '@mui/material';

interface BookingFormData {
  seatNumber: number;
}

const Booking: React.FC = () => {
  const { flightId } = useParams<{ flightId: string }>();
  const navigate = useNavigate();
  const { user } = useAuth();
  const { flights, loadingFlights, flightsError, fetchFlights } = useFlights();
  const { createBooking, loadingBookings, bookingsError } = useBookings();
  const [selectedFlight, setSelectedFlight] = useState<IFlight | null>(null);
  const [formData, setFormData] = useState<BookingFormData>({
    seatNumber: 1,
  });

  useEffect(() => {
    if (flightId) {
      fetchFlights({ id: parseInt(flightId) });
    }
  }, [flightId, fetchFlights]);

  useEffect(() => {
    if (flights.length > 0) {
      setSelectedFlight(flights[0]);
    }
  }, [flights]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!user || !selectedFlight) return;

    try {
      const result = await createBooking({
        flightId: selectedFlight.id,
        userId: user.id,
        seatNumber: formData.seatNumber,
      });

      if (result.success) {
        navigate('/dashboard');
      } else {
        // Handle error, e.g., display a message to the user
      }
    } catch (error) {
      // Handle error, e.g., display a message to the user
    }
  };

  if (loadingFlights || loadingBookings) {
    return <Typography>Loading...</Typography>;
  }

  if (flightsError || bookingsError) {
    return <Typography color="error">{flightsError || bookingsError}</Typography>;
  }

  if (!selectedFlight) {
    return <Typography>Flight not found</Typography>;
  }

  return (
    <Grid container spacing={3}>
      <Grid item xs={12}>
        <Paper elevation={3} sx={{ p: 3 }}>
          <Typography variant="h5" gutterBottom>
            Book Flight
          </Typography>

          <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography variant="h6">
                  Flight Details
                </Typography>
                <Typography>
                  Flight Code: {selectedFlight.flightCode}
                </Typography>
                <Typography>
                  Route: {selectedFlight.origin} → {selectedFlight.destination}
                </Typography>
                <Typography>
                  Departure: {new Date(selectedFlight.departureTime).toLocaleString()}
                </Typography>
                <Typography>
                  Arrival: {new Date(selectedFlight.arrivalTime).toLocaleString()}
                </Typography>
                <Typography>
                  Price: ${selectedFlight.price}
                </Typography>
              </Grid>

              <Grid item xs={12}>
                <FormControl fullWidth>
                  <InputLabel>Seat Number</InputLabel>
                  <Select
                    value={formData.seatNumber}
                    label="Seat Number"
                    onChange={(e) => setFormData({ ...formData, seatNumber: Number(e.target.value) })}
                  >
                    {Array.from({ length: selectedFlight.capacity }, (_, i) => (
                      <MenuItem key={i + 1} value={i + 1}>
                        Seat {i + 1}
                      </MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={12}>
                <Button
                  type="submit"
                  variant="contained"
                  color="primary"
                  fullWidth
                  disabled={loadingBookings}
                >
                  Confirm Booking
                </Button>
              </Grid>
            </Grid>
          </Box>
        </Paper>
      </Grid>
    </Grid>
  );
};

export default Booking; 