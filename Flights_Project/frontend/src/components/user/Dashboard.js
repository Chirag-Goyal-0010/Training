import React, { useState, useEffect } from 'react';
import {
  Box,
  Container,
  Typography,
  TextField,
  Grid,
  Card,
  CardContent,
  CardActions,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Paper,
} from '@mui/material';
import { format } from 'date-fns';
import axios from 'axios';

const Dashboard = () => {
  const [flights, setFlights] = useState([]);
  const [filteredFlights, setFilteredFlights] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [filters, setFilters] = useState({
    departureAirport: '',
    arrivalAirport: '',
    date: '',
    status: '',
  });

  useEffect(() => {
    fetchFlights();
  }, []);

  useEffect(() => {
    filterFlights();
  }, [flights, searchTerm, filters]);

  const fetchFlights = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/flights', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      });
      setFlights(response.data);
      setFilteredFlights(response.data);
    } catch (error) {
      console.error('Error fetching flights:', error);
    }
  };

  const filterFlights = () => {
    let filtered = [...flights];

    // Apply search term
    if (searchTerm) {
      filtered = filtered.filter(
        (flight) =>
          flight.departure_airport.toLowerCase().includes(searchTerm.toLowerCase()) ||
          flight.arrival_airport.toLowerCase().includes(searchTerm.toLowerCase())
      );
    }

    // Apply filters
    if (filters.departureAirport) {
      filtered = filtered.filter(
        (flight) => flight.departure_airport === filters.departureAirport
      );
    }
    if (filters.arrivalAirport) {
      filtered = filtered.filter(
        (flight) => flight.arrival_airport === filters.arrivalAirport
      );
    }
    if (filters.date) {
      filtered = filtered.filter(
        (flight) =>
          format(new Date(flight.departure_time), 'yyyy-MM-dd') === filters.date
      );
    }
    if (filters.status) {
      filtered = filtered.filter((flight) => flight.status === filters.status);
    }

    setFilteredFlights(filtered);
  };

  const handleSearch = (event) => {
    setSearchTerm(event.target.value);
  };

  const handleFilterChange = (event) => {
    setFilters({
      ...filters,
      [event.target.name]: event.target.value,
    });
  };

  const handleBookFlight = async (flightId) => {
    try {
      // Implement booking logic here
      console.log('Booking flight:', flightId);
    } catch (error) {
      console.error('Error booking flight:', error);
    }
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Typography variant="h4" gutterBottom>
        Available Flights
      </Typography>

      <Paper sx={{ p: 2, mb: 3 }}>
        <Grid container spacing={2}>
          <Grid item xs={12} md={4}>
            <TextField
              fullWidth
              label="Search Flights"
              value={searchTerm}
              onChange={handleSearch}
              variant="outlined"
            />
          </Grid>
          <Grid item xs={12} md={2}>
            <FormControl fullWidth>
              <InputLabel>Departure</InputLabel>
              <Select
                name="departureAirport"
                value={filters.departureAirport}
                onChange={handleFilterChange}
                label="Departure"
              >
                <MenuItem value="">All</MenuItem>
                {Array.from(new Set(flights.map((f) => f.departure_airport))).map(
                  (airport) => (
                    <MenuItem key={airport} value={airport}>
                      {airport}
                    </MenuItem>
                  )
                )}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12} md={2}>
            <FormControl fullWidth>
              <InputLabel>Arrival</InputLabel>
              <Select
                name="arrivalAirport"
                value={filters.arrivalAirport}
                onChange={handleFilterChange}
                label="Arrival"
              >
                <MenuItem value="">All</MenuItem>
                {Array.from(new Set(flights.map((f) => f.arrival_airport))).map(
                  (airport) => (
                    <MenuItem key={airport} value={airport}>
                      {airport}
                    </MenuItem>
                  )
                )}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={12} md={2}>
            <TextField
              fullWidth
              type="date"
              name="date"
              label="Date"
              value={filters.date}
              onChange={handleFilterChange}
              InputLabelProps={{ shrink: true }}
            />
          </Grid>
          <Grid item xs={12} md={2}>
            <FormControl fullWidth>
              <InputLabel>Status</InputLabel>
              <Select
                name="status"
                value={filters.status}
                onChange={handleFilterChange}
                label="Status"
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value="On time">On Time</MenuItem>
                <MenuItem value="Delayed">Delayed</MenuItem>
                <MenuItem value="Cancelled">Cancelled</MenuItem>
              </Select>
            </FormControl>
          </Grid>
        </Grid>
      </Paper>

      <Grid container spacing={3}>
        {filteredFlights.map((flight) => (
          <Grid item xs={12} md={6} lg={4} key={flight.id}>
            <Card>
              <CardContent>
                <Typography variant="h6" gutterBottom>
                  {flight.departure_airport} → {flight.arrival_airport}
                </Typography>
                <Typography color="textSecondary" gutterBottom>
                  Departure: {format(new Date(flight.departure_time), 'PPpp')}
                </Typography>
                <Typography color="textSecondary" gutterBottom>
                  Arrival: {format(new Date(flight.arrival_time), 'PPpp')}
                </Typography>
                <Typography color="textSecondary" gutterBottom>
                  Distance: {flight.distance} km
                </Typography>
                <Typography
                  color={flight.status === 'On time' ? 'success.main' : 'error.main'}
                >
                  Status: {flight.status}
                </Typography>
              </CardContent>
              <CardActions>
                <Button
                  size="small"
                  color="primary"
                  onClick={() => handleBookFlight(flight.id)}
                >
                  Book Flight
                </Button>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Container>
  );
};

export default Dashboard; 