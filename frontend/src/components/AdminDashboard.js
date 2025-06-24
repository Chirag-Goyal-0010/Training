import React, { useState, useEffect } from 'react';
import axios from 'axios';
import API_BASE_URL from '../api';
import {
  Container,
  Paper,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Alert,
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Snackbar,
} from '@mui/material';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { format } from 'date-fns';
import { useNavigate } from 'react-router-dom';

function AdminDashboard() {
  const [flights, setFlights] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [openDialog, setOpenDialog] = useState(false);
  const [showAllFlights, setShowAllFlights] = useState(false);
  const [filterOrigin, setFilterOrigin] = useState('');
  const [filterDestination, setFilterDestination] = useState('');
  const [filterStatus, setFilterStatus] = useState('');
  const [filterDate, setFilterDate] = useState(null);
  const [originOptions, setOriginOptions] = useState([]);
  const [destinationOptions, setDestinationOptions] = useState([]);
  const [snackbarOpen, setSnackbarOpen] = useState(false);

  const [newFlight, setNewFlight] = useState({
    origin: '',
    destination: '',
    departure_time: new Date(),
    arrival_time: new Date(),
    economy_price: '',
    premium_economy_price: '',
    business_price: '',
    first_class_price: '',
    economy_seats: '',
    premium_economy_seats: '',
    business_seats: '',
    first_class_seats: '',
  });

  const navigate = useNavigate();

  useEffect(() => {
    fetchFlights();
  }, [showAllFlights]);

  useEffect(() => {
    const fetchLocations = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`${API_BASE_URL}/locations`, {
          headers: { Authorization: `Bearer ${token}` },
        });
        setOriginOptions(response.data.origins);
        setDestinationOptions(response.data.destinations);
      } catch (error) {
        console.error('Failed to fetch locations:', error);
      }
    };
    fetchLocations();
  }, []);

  const fetchFlights = async () => {
    try {
      const token = localStorage.getItem('token');
      let url = `${API_BASE_URL}/flights`;
      const query = new URLSearchParams();

      if (showAllFlights) {
        query.append('all_flights', 'true');
      }

      if (filterOrigin) query.append('origin', filterOrigin);
      if (filterDestination) query.append('destination', filterDestination);
      if (filterStatus) query.append('status', filterStatus);
      if (filterDate) query.append('departure_date', format(filterDate, 'yyyy-MM-dd'));

      if (query.toString()) {
        url = `${url}?${query.toString()}`;
      }

      const response = await axios.get(url, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setFlights(response.data.data);
    } catch (error) {
      setError('Failed to fetch flights');
    }
  };

  const handleClearFilters = () => {
    setFilterOrigin('');
    setFilterDestination('');
    setFilterStatus('');
    setFilterDate(null);
    fetchFlights();
  };

  const handleInputChange = (field) => (event) => {
    setNewFlight({
      ...newFlight,
      [field]: event.target.value,
    });
  };

  const handleDateChange = (field) => (date) => {
    setNewFlight({
      ...newFlight,
      [field]: date,
    });
  };

  const handleAddFlight = async () => {
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        setError('Please login to add flights');
        setSnackbarOpen(true);
        console.error('Error: No token found. User not authenticated.');
        return;
      }

      if (newFlight.origin.trim().toLowerCase() === newFlight.destination.trim().toLowerCase()) {
        setError('Origin and destination cannot be the same.');
        setSnackbarOpen(true);
        return;
      }

      const flightData = {
        ...newFlight,
        economy_price: parseFloat(newFlight.economy_price),
        premium_economy_price: parseFloat(newFlight.premium_economy_price),
        business_price: parseFloat(newFlight.business_price),
        first_class_price: parseFloat(newFlight.first_class_price),
        economy_seats: parseInt(newFlight.economy_seats, 10),
        premium_economy_seats: parseInt(newFlight.premium_economy_seats, 10),
        business_seats: parseInt(newFlight.business_seats, 10),
        first_class_seats: parseInt(newFlight.first_class_seats, 10),
        departure_time: newFlight.departure_time.toISOString(),
        arrival_time: newFlight.arrival_time.toISOString(),
      };

      if (
        !flightData.origin ||
        !flightData.destination ||
        isNaN(flightData.economy_price) || flightData.economy_price < 0 ||
        isNaN(flightData.premium_economy_price) || flightData.premium_economy_price < 0 ||
        isNaN(flightData.business_price) || flightData.business_price < 0 ||
        isNaN(flightData.first_class_price) || flightData.first_class_price < 0 ||
        (isNaN(flightData.economy_seats) &&
          isNaN(flightData.premium_economy_seats) &&
          isNaN(flightData.business_seats) &&
          isNaN(flightData.first_class_seats)) ||
        (flightData.economy_seats < 0 ||
          flightData.premium_economy_seats < 0 ||
          flightData.business_seats < 0 ||
          flightData.first_class_seats < 0) ||
        (flightData.economy_seats === 0 &&
          flightData.premium_economy_seats === 0 &&
          flightData.business_seats === 0 &&
          flightData.first_class_seats === 0)
      ) {
        setError(
          'Please fill all required fields and ensure prices/seats are valid numbers (at least one seat class must have available seats, and prices cannot be negative).'
        );
        setSnackbarOpen(true);
        console.error('Validation Error: Missing or invalid flight data.', flightData);
        return;
      }

      console.log('Sending flight data:', flightData);

      await axios.post(`${API_BASE_URL}/admin/flights`, flightData, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setSuccess('Flight added successfully');
      setOpenDialog(false);
      fetchFlights();
      setTimeout(() => setSuccess(''), 3000);
      setError('');
    } catch (error) {
      let msg = error.response?.data?.error || 'Failed to add flight. Check browser console for details.';
      if (
        msg.includes("ArrivalTime") &&
        (msg.includes("gtfield") || msg.toLowerCase().includes("arrival time") || msg.toLowerCase().includes("after departure"))
      ) {
        msg = 'Arrival time must be after departure time.';
      }
      setError(msg);
      setSnackbarOpen(true);
    }
  };

  const handleDeleteFlight = async (id) => {
    try {
      const token = localStorage.getItem('token');
      await axios.delete(`${API_BASE_URL}/admin/flights/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setSuccess('Flight deleted successfully');
      fetchFlights();
      setTimeout(() => setSuccess(''), 3000);
    } catch (error) {
      setError(error.response?.data?.error || 'Failed to delete flight');
    }
  };

  const indianCities = [
    'Mumbai', 'Delhi', 'Bangalore', 'Hyderabad', 'Chennai',
    'Kolkata', 'Ahmedabad', 'Pune', 'Jaipur', 'Lucknow',
    'Kanpur', 'Nagpur', 'Indore', 'Thane', 'Bhopal',
    'Visakhapatnam', 'Pimpri-Chinchwad', 'Patna', 'Vadodara', 'Ghaziabad'
  ];

  return (
    <LocalizationProvider dateAdapter={AdapterDateFns}>
      <Container maxWidth="lg" sx={{ mt: 4 }}>
        <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
          <Typography variant="h4">Manage Flights</Typography>
          <Box>
            <Button
              variant="contained"
              color="primary"
              onClick={() => setOpenDialog(true)}
              sx={{ mr: 2 }}
            >
              Add New Flight
            </Button>
            <Button
              variant="outlined"
              color="secondary"
              onClick={() => setShowAllFlights(!showAllFlights)}
            >
              {showAllFlights ? 'Show Paginated Flights' : 'Show All Flights'}
            </Button>
          </Box>
        </Box>

        <Box mb={3} display="flex" gap={2} alignItems="center">
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <InputLabel>Filter by Origin</InputLabel>
            <Select
              value={filterOrigin}
              label="Filter by Origin"
              onChange={(e) => setFilterOrigin(e.target.value)}
            >
              <MenuItem value="">All</MenuItem>
              {originOptions.map((option) => (
                <MenuItem key={option} value={option}>
                  {option}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <InputLabel>Filter by Destination</InputLabel>
            <Select
              value={filterDestination}
              label="Filter by Destination"
              onChange={(e) => setFilterDestination(e.target.value)}
            >
              <MenuItem value="">All</MenuItem>
              {destinationOptions.map((option) => (
                <MenuItem key={option} value={option}>
                  {option}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <FormControl size="small" sx={{ minWidth: 120 }}>
            <InputLabel>Status</InputLabel>
            <Select
              value={filterStatus}
              label="Status"
              onChange={(e) => setFilterStatus(e.target.value)}
            >
              <MenuItem value="">All</MenuItem>
              <MenuItem value="Scheduled">Scheduled</MenuItem>
              <MenuItem value="In Air">In Air</MenuItem>
              <MenuItem value="Landed">Landed</MenuItem>
              <MenuItem value="Departing Soon">Departing Soon</MenuItem>
            </Select>
          </FormControl>

          <DatePicker
            label="Filter by Date"
            value={filterDate}
            onChange={(newValue) => setFilterDate(newValue)}
            renderInput={(params) => <TextField {...params} size="small" sx={{ width: 200 }} />}
            inputFormat="yyyy-MM-dd"
          />

          <Button
            variant="contained"
            onClick={() => fetchFlights()}
          >
            Apply Filters
          </Button>
          <Button
            variant="outlined"
            onClick={handleClearFilters}
          >
            Clear Filters
          </Button>
        </Box>

        <Snackbar
          open={snackbarOpen && !!error}
          autoHideDuration={4000}
          onClose={() => setSnackbarOpen(false)}
          message={error}
          anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
        />

        {success && (
          <Alert severity="success" sx={{ mb: 2 }}>
            {success}
          </Alert>
        )}

        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Flight Number</TableCell>
                <TableCell>Route</TableCell>
                <TableCell>Departure</TableCell>
                <TableCell>Arrival</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Action</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {flights.map((flight) => {
                let status = 'Scheduled';
                if (flight.is_in_air) {
                  status = 'In Air';
                } else if (flight.is_landed) {
                  status = 'Landed';
                } else if (flight.is_departing_soon) {
                  status = 'Departing Soon';
                }
                return (
                  <TableRow
                    key={flight.ID}
                    hover
                    sx={{ cursor: 'pointer' }}
                    onClick={() => navigate(`/admin/flights/${flight.ID}`)}
                  >
                    <TableCell>{flight.ID}</TableCell>
                    <TableCell>{flight.origin} → {flight.destination}</TableCell>
                    <TableCell>
                      {format(new Date(flight.departure_time), 'dd/MM/yyyy HH:mm')}
                    </TableCell>
                    <TableCell>
                      {format(new Date(flight.arrival_time), 'dd/MM/yyyy HH:mm')}
                    </TableCell>
                    <TableCell>{status}</TableCell>
                    <TableCell>
                      <Button
                        variant="contained"
                        color="error"
                        onClick={(e) => {
                          e.stopPropagation();
                          handleDeleteFlight(flight.ID);
                        }}
                      >
                        Delete
                      </Button>
                    </TableCell>
                  </TableRow>
                );
              })}
            </TableBody>
          </Table>
        </TableContainer>

        <Dialog open={openDialog} onClose={() => setOpenDialog(false)} maxWidth="sm" fullWidth>
          <DialogTitle>Add New Flight</DialogTitle>
          <DialogContent>
            <FormControl fullWidth margin="normal">
              <InputLabel id="origin-select-label">Select Origin</InputLabel>
              <Select
                labelId="origin-select-label"
                id="origin-select"
                value={newFlight.origin}
                label="Select Origin"
                onChange={handleInputChange('origin')}
              >
                <MenuItem value="" disabled>Select Origin</MenuItem>
                {indianCities.map((city) => (
                  <MenuItem key={city} value={city}>
                    {city}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl fullWidth margin="normal">
              <InputLabel id="destination-select-label">Select Destination</InputLabel>
              <Select
                labelId="destination-select-label"
                id="destination-select"
                value={newFlight.destination}
                label="Select Destination"
                onChange={handleInputChange('destination')}
              >
                <MenuItem value="" disabled>Select Destination</MenuItem>
                {indianCities.map((city) => (
                  <MenuItem key={city} value={city}>
                    {city}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <DateTimePicker
              label="Departure Time"
              value={newFlight.departure_time}
              onChange={handleDateChange('departure_time')}
              renderInput={(params) => (
                <TextField {...params} fullWidth margin="normal" />
              )}
            />
            <DateTimePicker
              label="Arrival Time"
              value={newFlight.arrival_time}
              onChange={handleDateChange('arrival_time')}
              renderInput={(params) => (
                <TextField {...params} fullWidth margin="normal" />
              )}
            />
            <TextField
              label="Economy Price (₹)"
              type="number"
              value={newFlight.economy_price}
              onChange={handleInputChange('economy_price')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="Premium Economy Price (₹)"
              type="number"
              value={newFlight.premium_economy_price}
              onChange={handleInputChange('premium_economy_price')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="Business Price (₹)"
              type="number"
              value={newFlight.business_price}
              onChange={handleInputChange('business_price')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="First Class Price (₹)"
              type="number"
              value={newFlight.first_class_price}
              onChange={handleInputChange('first_class_price')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="Economy Seats"
              type="number"
              value={newFlight.economy_seats}
              onChange={handleInputChange('economy_seats')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="Premium Economy Seats"
              type="number"
              value={newFlight.premium_economy_seats}
              onChange={handleInputChange('premium_economy_seats')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="Business Seats"
              type="number"
              value={newFlight.business_seats}
              onChange={handleInputChange('business_seats')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
            <TextField
              label="First Class Seats"
              type="number"
              value={newFlight.first_class_seats}
              onChange={handleInputChange('first_class_seats')}
              fullWidth
              margin="normal"
              InputProps={{ inputProps: { min: 0 } }}
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setOpenDialog(false)}>Cancel</Button>
            <Button onClick={handleAddFlight} variant="contained" color="primary">
              Add Flight
            </Button>
          </DialogActions>
        </Dialog>
      </Container>
    </LocalizationProvider>
  );
}

export default AdminDashboard;