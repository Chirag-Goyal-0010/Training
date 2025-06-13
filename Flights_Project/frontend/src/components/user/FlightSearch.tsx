import React, { useState, useEffect } from 'react';
import { useFlights } from '../../context/FlightsContext';
import { IFlight } from '../../context/FlightsContext';
import { TextField, Button, Grid, Paper, Typography } from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';

interface SearchCriteria {
  origin?: string;
  destination?: string;
  date?: Date | null;
}

const FlightSearch: React.FC = () => {
  const { flights, loadingFlights, flightsError, fetchFlights } = useFlights();
  const [searchCriteria, setSearchCriteria] = useState<SearchCriteria>({
    origin: '',
    destination: '',
    date: null,
  });

  const handleSearch = () => {
    const formattedCriteria = {
      ...searchCriteria,
      date: searchCriteria.date ? searchCriteria.date.toISOString().split('T')[0] : undefined,
    };
    fetchFlights(formattedCriteria);
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDateFns}>
      <Paper elevation={3} sx={{ p: 3, mb: 3 }}>
        <Typography variant="h6" gutterBottom>
          Search Flights
        </Typography>
        <Grid container spacing={2}>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Origin"
              value={searchCriteria.origin}
              onChange={(e) => setSearchCriteria({ ...searchCriteria, origin: e.target.value })}
            />
          </Grid>
          <Grid item xs={12} sm={4}>
            <TextField
              fullWidth
              label="Destination"
              value={searchCriteria.destination}
              onChange={(e) => setSearchCriteria({ ...searchCriteria, destination: e.target.value })}
            />
          </Grid>
          <Grid item xs={12} sm={4}>
            <DatePicker
              label="Date"
              value={searchCriteria.date}
              onChange={(date) => setSearchCriteria({ ...searchCriteria, date })}
              renderInput={(params) => <TextField {...params} fullWidth />}
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              variant="contained"
              color="primary"
              onClick={handleSearch}
              disabled={loadingFlights}
            >
              Search
            </Button>
          </Grid>
        </Grid>
      </Paper>

      {loadingFlights && <Typography>Loading flights...</Typography>}
      {flightsError && <Typography color="error">{flightsError}</Typography>}
      
      <Grid container spacing={2}>
        {flights.map((flight: IFlight) => (
          <Grid item xs={12} key={flight.id}>
            <Paper elevation={2} sx={{ p: 2 }}>
              <Typography variant="h6">
                {flight.flightCode}: {flight.origin} → {flight.destination}
              </Typography>
              <Typography>
                Departure: {new Date(flight.departureTime).toLocaleString()}
              </Typography>
              <Typography>
                Arrival: {new Date(flight.arrivalTime).toLocaleString()}
              </Typography>
              <Typography>
                Price: ${flight.price}
              </Typography>
            </Paper>
          </Grid>
        ))}
      </Grid>
    </LocalizationProvider>
  );
};

export default FlightSearch; 