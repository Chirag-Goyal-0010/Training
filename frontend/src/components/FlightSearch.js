import React, { useState, useEffect } from 'react';
import {
  Container,
  Paper,
  Box,
  Typography,
  Button,
  TextField,
  ToggleButtonGroup,
  ToggleButton,
  Autocomplete,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  InputLabel,
  Select,
  MenuItem,
  FormControl,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import axios from 'axios';
import { format } from 'date-fns';
import API_BASE_URL from '../api';

const FlightSearch = ({ onSearch }) => {
  const [tripType, setTripType] = useState('oneWay'); // 'oneWay' or 'roundTrip'
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [departureDate, setDepartureDate] = useState(null);
  const [returnDate, setReturnDate] = useState(null);
  const [travellers, setTravellers] = useState(1);
  const [travelClass, setTravelClass] = useState('Economy');
  const [origins, setOrigins] = useState([]);
  const [destinations, setDestinations] = useState([]);
  const [openTravellerDialog, setOpenTravellerDialog] = useState(false);

  useEffect(() => {
    const fetchLocations = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`${API_BASE_URL}/locations`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setOrigins(response.data.origins);
        setDestinations(response.data.destinations);
      } catch (error) {
        console.error('Error fetching locations:', error);
      }
    };
    fetchLocations();
  }, []);

  const handleTripTypeChange = (event, newTripType) => {
    if (newTripType !== null) {
      setTripType(newTripType);
    }
  };

  const handleSearch = () => {
    onSearch({
      from,
      to,
      departureDate: departureDate ? format(departureDate, 'yyyy-MM-dd') : null,
      travellers,
      travelClass,
    });
  };

  const handleTravellerDialogClose = () => {
    setOpenTravellerDialog(false);
  };

  const handleTravellersChange = (event) => {
    const value = parseInt(event.target.value, 10);
    if (!isNaN(value) && value > 0) {
      setTravellers(value);
    } else if (event.target.value === '') {
      setTravellers('');
    }
  };

  const handleTravelClassChange = (event) => {
    setTravelClass(event.target.value);
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Paper elevation={3} sx={{ p: 3 }}>
        <Box display="flex" justifyContent="center" mb={3}>
          <ToggleButtonGroup
            value={tripType}
            exclusive
            onChange={handleTripTypeChange}
            aria-label="trip type"
          >
            <ToggleButton value="oneWay" aria-label="one way">
              One Way
            </ToggleButton>
            <ToggleButton value="roundTrip" aria-label="round trip">
              Round Trip
            </ToggleButton>
          </ToggleButtonGroup>
        </Box>

        <Box display="flex" gap={2} mb={3} flexWrap="wrap">
          <Autocomplete
            options={origins}
            value={from}
            onChange={(event, newValue) => {
              setFrom(newValue);
            }}
            renderInput={(params) => <TextField {...params} label="From" sx={{ flex: 1, minWidth: '200px' }} />}
            sx={{ flex: 1, minWidth: '200px' }}
          />
          <Autocomplete
            options={destinations}
            value={to}
            onChange={(event, newValue) => {
              setTo(newValue);
            }}
            renderInput={(params) => <TextField {...params} label="To" sx={{ flex: 1, minWidth: '200px' }} />}
            sx={{ flex: 1, minWidth: '200px' }}
          />
          <LocalizationProvider dateAdapter={AdapterDateFns}>
            <DatePicker
              label="Departure"
              value={departureDate}
              onChange={(newValue) => setDepartureDate(newValue)}
              renderInput={(params) => <TextField {...params} sx={{ flex: 1, minWidth: '200px' }} />}
            />
            {tripType === 'roundTrip' && (
              <DatePicker
                label="Return"
                value={returnDate}
                onChange={(newValue) => setReturnDate(newValue)}
                renderInput={(params) => <TextField {...params} sx={{ flex: 1, minWidth: '200px' }} />}
              />
            )}
          </LocalizationProvider>
          <Button
            variant="outlined"
            onClick={() => setOpenTravellerDialog(true)}
            sx={{ flex: 1, minWidth: '200px', height: '56px', borderColor: 'rgba(0, 0, 0, 0.23)', color: 'rgba(0, 0, 0, 0.87)' }}
          >
            {`${travellers} Traveller, ${travelClass}`}
          </Button>
        </Box>

        <Box display="flex" justifyContent="flex-end">
          <Button variant="contained" onClick={handleSearch} sx={{ px: 5, py: 1.5 }}>
            Search
          </Button>
        </Box>
      </Paper>

      <Dialog open={openTravellerDialog} onClose={handleTravellerDialogClose}>
        <DialogTitle>Select Travellers and Class</DialogTitle>
        <DialogContent>
          <TextField
            label="Travellers"
            type="number"
            value={travellers}
            onChange={handleTravellersChange}
            fullWidth
            margin="normal"
            inputProps={{ min: 1 }}
          />
          <FormControl fullWidth margin="normal">
            <InputLabel id="travel-class-label">Travel Class</InputLabel>
            <Select
              labelId="travel-class-label"
              value={travelClass}
              label="Travel Class"
              onChange={handleTravelClassChange}
            >
              <MenuItem value="Economy">Economy</MenuItem>
              <MenuItem value="PremiumEconomy">Premium Economy</MenuItem>
              <MenuItem value="Business">Business</MenuItem>
              <MenuItem value="FirstClass">First Class</MenuItem>
            </Select>
          </FormControl>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleTravellerDialogClose}>OK</Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default FlightSearch; 