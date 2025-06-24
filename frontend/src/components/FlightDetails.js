import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import {
  Container,
  Typography,
  Paper,
  Box,
  CircularProgress,
  Alert,
  List,
  ListItem,
  ListItemText,
} from '@mui/material';
import { format } from 'date-fns';
import API_BASE_URL from '../api';

function FlightDetails() {
  const { id } = useParams();
  const [flight, setFlight] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchFlightDetails = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          setError('Authentication token not found. Please log in.');
          setLoading(false);
          return;
        }

        const response = await axios.get(`${API_BASE_URL}/admin/flights/${id}`, {
          headers: { Authorization: `Bearer ${token}` },
        });
        setFlight(response.data.data);
        setLoading(false);
      } catch (err) {
        console.error('Error fetching flight details:', err.response?.data || err.message);
        setError(err.response?.data?.error || 'Failed to load flight details.');
        setLoading(false);
      }
    };

    fetchFlightDetails();
  }, [id]);

  if (loading) {
    return (
      <Container maxWidth="md" sx={{ mt: 4, textAlign: 'center' }}>
        <CircularProgress />
        <Typography variant="h6">Loading flight details...</Typography>
      </Container>
    );
  }

  if (error) {
    return (
      <Container maxWidth="md" sx={{ mt: 4 }}>
        <Alert severity="error">{error}</Alert>
      </Container>
    );
  }

  if (!flight) {
    return (
      <Container maxWidth="md" sx={{ mt: 4 }}>
        <Alert severity="info">Flight details not found.</Alert>
      </Container>
    );
  }

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Paper elevation={3} sx={{ p: 4 }}>
        <Typography variant="h4" gutterBottom>
          Flight Details: {flight.flight_number || flight.ID}
        </Typography>
        <Box sx={{ mt: 3 }}>
          <List>
            <ListItem>
              <ListItemText primary="Origin" secondary={flight.origin} />
            </ListItem>
            <ListItem>
              <ListItemText primary="Destination" secondary={flight.destination} />
            </ListItem>
            <ListItem>
              <ListItemText
                primary="Departure Time"
                secondary={format(new Date(flight.departure_time), 'dd/MM/yyyy HH:mm')}
              />
            </ListItem>
            <ListItem>
              <ListItemText
                primary="Arrival Time"
                secondary={format(new Date(flight.arrival_time), 'dd/MM/yyyy HH:mm')}
              />
            </ListItem>
            <ListItem>
              <ListItemText primary="Total Seats" secondary={flight.total_flight_seats} />
            </ListItem>
            <ListItem>
              <ListItemText primary="Economy Price" secondary={`₹${flight.economy_price}`} />
            </ListItem>
            <ListItem>
              <ListItemText primary="Premium Economy Price" secondary={`₹${flight.premium_economy_price}`} />
            </ListItem>
            <ListItem>
              <ListItemText primary="Business Price" secondary={`₹${flight.business_price}`} />
            </ListItem>
            <ListItem>
              <ListItemText primary="First Class Price" secondary={`₹${flight.first_class_price}`} />
            </ListItem>
            {/* Economy Class Details */}
            <ListItem>
              <ListItemText
                primary="Economy Seats"
                secondary={`Total: ${flight.total_economy_seats}, Booked: ${flight.booked_economy_seats}, Available: ${flight.available_economy_seats}`}
              />
            </ListItem>
            {/* Premium Economy Class Details */}
            <ListItem>
              <ListItemText
                primary="Premium Economy Seats"
                secondary={`Total: ${flight.total_premium_economy_seats}, Booked: ${flight.booked_premium_economy_seats}, Available: ${flight.available_premium_economy_seats}`}
              />
            </ListItem>
            {/* Business Class Details */}
            <ListItem>
              <ListItemText
                primary="Business Seats"
                secondary={`Total: ${flight.total_business_seats}, Booked: ${flight.booked_business_seats}, Available: ${flight.available_business_seats}`}
              />
            </ListItem>
            {/* First Class Details */}
            <ListItem>
              <ListItemText
                primary="First Class Seats"
                secondary={`Total: ${flight.total_first_class_seats}, Booked: ${flight.booked_first_class_seats}, Available: ${flight.available_first_class_seats}`}
              />
            </ListItem>
          </List>
        </Box>
        {/* Bookings and Travellers Section */}
        {flight.bookings && flight.bookings.length > 0 && (
          <Box sx={{ mt: 4 }}>
            <Typography variant="h5" gutterBottom>All Bookings & Travellers</Typography>
            {flight.bookings.map((booking, idx) => (
              <Paper key={booking.ID || idx} sx={{ p: 2, mb: 3, background: '#f9f9f9' }}>
                <Typography variant="subtitle1" sx={{ fontWeight: 'bold' }}>
                  Booking #{booking.ID} | Class: {booking.travel_class} | Seats: {booking.seats} | Status: {booking.status}
                </Typography>
                <Typography variant="body2" sx={{ mb: 1 }}>
                  Booked on: {booking.booking_date ? format(new Date(booking.booking_date), 'dd/MM/yyyy HH:mm') : 'N/A'}
                </Typography>
                {booking.travellers && booking.travellers.length > 0 ? (
                  <List dense>
                    {booking.travellers.map((trav, tIdx) => (
                      <ListItem key={trav.id || tIdx} sx={{ pl: 4 }}>
                        <ListItemText
                          primary={`${trav.title} ${trav.first_name} ${trav.last_name}`}
                          secondary={`DOB: ${trav.dob ? format(new Date(trav.dob), 'dd/MM/yyyy') : ''} | Nationality: ${trav.nationality}`}
                        />
                      </ListItem>
                    ))}
                  </List>
                ) : (
                  <Typography variant="body2" color="text.secondary">No traveller details.</Typography>
                )}
              </Paper>
            ))}
          </Box>
        )}
      </Paper>
    </Container>
  );
}

export default FlightDetails; 