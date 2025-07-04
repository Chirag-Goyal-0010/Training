import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Container,
  Paper,
  TextField,
  Button,
  Typography,
  Box,
  MenuItem,
} from '@mui/material';

const AddFlight = ({ onFlightAdded }) => {
  const navigate = useNavigate();
  const [flight, setFlight] = useState({
    aircraft_id: '',
    departure_airport_id: '',
    arrival_airport_id: '',
    departure_time: '',
    arrival_time: '',
    distance: '',
    status: 'On time',
  });

  const handleChange = (e) => {
    setFlight({
      ...flight,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/api/admin/flights', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify(flight),
      });

      if (!response.ok) {
        throw new Error('Failed to add flight');
      }

      onFlightAdded();
      navigate('/admin');
    } catch (error) {
      console.error('Error adding flight:', error);
    }
  };

  return (
    <Container maxWidth="sm">
      <Paper elevation={3} sx={{ p: 4, mt: 4 }}>
        <Typography variant="h5" component="h2" gutterBottom>
          Add New Flight
        </Typography>
        <form onSubmit={handleSubmit}>
          <TextField
            fullWidth
            margin="normal"
            label="Aircraft ID"
            name="aircraft_id"
            type="number"
            value={flight.aircraft_id}
            onChange={handleChange}
            required
          />
          <TextField
            fullWidth
            margin="normal"
            label="Departure Airport ID"
            name="departure_airport_id"
            type="number"
            value={flight.departure_airport_id}
            onChange={handleChange}
            required
          />
          <TextField
            fullWidth
            margin="normal"
            label="Arrival Airport ID"
            name="arrival_airport_id"
            type="number"
            value={flight.arrival_airport_id}
            onChange={handleChange}
            required
          />
          <TextField
            fullWidth
            margin="normal"
            label="Departure Time"
            name="departure_time"
            type="datetime-local"
            value={flight.departure_time}
            onChange={handleChange}
            InputLabelProps={{
              shrink: true,
            }}
            required
          />
          <TextField
            fullWidth
            margin="normal"
            label="Arrival Time"
            name="arrival_time"
            type="datetime-local"
            value={flight.arrival_time}
            onChange={handleChange}
            InputLabelProps={{
              shrink: true,
            }}
            required
          />
          <TextField
            fullWidth
            margin="normal"
            label="Distance"
            name="distance"
            type="number"
            value={flight.distance}
            onChange={handleChange}
            required
          />
          <TextField
            fullWidth
            margin="normal"
            select
            label="Status"
            name="status"
            value={flight.status}
            onChange={handleChange}
            required
          >
            <MenuItem value="On time">On time</MenuItem>
            <MenuItem value="Delayed">Delayed</MenuItem>
            <MenuItem value="Cancelled">Cancelled</MenuItem>
          </TextField>
          <Box sx={{ mt: 3, display: 'flex', gap: 2 }}>
            <Button
              type="submit"
              variant="contained"
              color="primary"
              fullWidth
            >
              Add Flight
            </Button>
            <Button
              variant="outlined"
              color="secondary"
              fullWidth
              onClick={() => navigate('/admin')}
            >
              Cancel
            </Button>
          </Box>
        </form>
      </Paper>
    </Container>
  );
};

export default AddFlight; 