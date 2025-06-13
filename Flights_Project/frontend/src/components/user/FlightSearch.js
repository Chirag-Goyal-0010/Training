import React, { useState, useEffect } from 'react';
import {
  Box,
  Paper,
  Typography,
  Grid,
  Button,
  Card,
  CardContent,
  CardActions,
} from '@mui/material';
import SearchFilters from '../SearchFilters';
import Booking from '../Booking';
import { useFlights } from '../../context/FlightsContext';

const FlightSearch = () => {
  const { flights, loadingFlights, flightsError, fetchFlights } = useFlights();

  const [filters, setFilters] = useState({
    priceRange: [0, 1000],
    sortBy: 'price',
    departureTime: null,
    arrivalTime: null,
    origin: '',
    destination: '',
    date: '',
  });
  const [selectedFlight, setSelectedFlight] = useState(null);

  useEffect(() => {
    const searchCriteria = {
      origin: filters.origin,
      destination: filters.destination,
      date: filters.date ? filters.date.toISOString().split('T')[0] : '',
    };
    fetchFlights(searchCriteria);
  }, [filters, fetchFlights]);

  const handleFilterChange = (newFilters) => {
    setFilters((prevFilters) => ({ ...prevFilters, ...newFilters }));
  };

  const handleBookFlight = (flight) => {
    setSelectedFlight(flight);
  };

  const handleCloseBooking = () => {
    setSelectedFlight(null);
  };

  if (loadingFlights) {
    return (
      <Box sx={{ p: 3 }}>
        <Typography variant="h6">Loading flights...</Typography>
      </Box>
    );
  }

  if (flightsError) {
    return (
      <Box sx={{ p: 3 }}>
        <Typography variant="h6" color="error">Error: {flightsError}</Typography>
      </Box>
    );
  }

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Search Flights
      </Typography>

      <SearchFilters onFilterChange={handleFilterChange} />

      <Grid container spacing={3} sx={{ mt: 2 }}>
        {flights.length === 0 ? (
          <Grid item xs={12}>
            <Typography variant="h6" align="center">No flights found for your search criteria.</Typography>
          </Grid>
        ) : (
          flights.map((flight) => (
            <Grid item xs={12} md={6} key={flight.id}>
              <Card>
                <CardContent>
                  <Typography variant="h6" gutterBottom>
                    {flight.Origin} → {flight.Destination}
                  </Typography>
                  <Typography color="textSecondary">
                    Departure: {new Date(flight.DepartureTime).toLocaleString()}
                  </Typography>
                  <Typography color="textSecondary">
                    Arrival: {new Date(flight.ArrivalTime).toLocaleString()}
                  </Typography>
                  <Typography color="textSecondary">
                    Flight Code: {flight.FlightCode}
                  </Typography>
                  <Typography variant="h6" color="primary" sx={{ mt: 2 }}>
                    ${flight.Price}
                  </Typography>
                </CardContent>
                <CardActions>
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={() => handleBookFlight(flight)}
                  >
                    Book Now
                  </Button>
                </CardActions>
              </Card>
            </Grid>
          ))
        )}
      </Grid>

      {selectedFlight && (
        <Booking
          flight={selectedFlight}
          onClose={handleCloseBooking}
        />
      )}
    </Box>
  );
};

export default FlightSearch; 