import React, { useState, useEffect } from 'react';
import {
  Container,
  Paper,
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
  Typography,
  TextField,
  Alert,
  FormControl, InputLabel, Select, MenuItem, Box, Checkbox, FormControlLabel
} from '@mui/material';
import axios from 'axios';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import API_BASE_URL from '../api';

const TITLES = ['Mr', 'Mrs', 'Ms', 'Dr', 'Master', 'Miss'];
const NATIONALITIES = ['India', 'Other'];

const FlightList = ({ searchParams }) => {
  const [flights, setFlights] = useState([]);
  const [selectedFlight, setSelectedFlight] = useState(null);
  const [open, setOpen] = useState(false);
  const [seats, setSeats] = useState(1);
  const [bookingError, setBookingError] = useState('');
  const [displayTotalPrice, setDisplayTotalPrice] = useState(0);
  const [isPremiumBooking, setIsPremiumBooking] = useState(false);
  const [travelClass, setTravelClass] = useState('Economy');
  const [pricePerSelectedClass, setPricePerSelectedClass] = useState(0);
  const [travellerDetails, setTravellerDetails] = useState([
    { title: '', first_name: '', last_name: '', dob: null, nationality: 'India' }
  ]);
  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(10);
  const [total, setTotal] = useState(0);
  const [showAll, setShowAll] = useState(false);
  const [showSurchargeDialog, setShowSurchargeDialog] = useState(false);
  const [pendingBooking, setPendingBooking] = useState(false);

  useEffect(() => {
    fetchFlights();
  }, [searchParams, page, showAll]);

  useEffect(() => {
    if (selectedFlight) {
      let currentBasePrice = 0;
      switch (travelClass) {
        case 'Economy':
          currentBasePrice = selectedFlight.economy_price;
          break;
        case 'PremiumEconomy':
          currentBasePrice = selectedFlight.premium_economy_price;
          break;
        case 'Business':
          currentBasePrice = selectedFlight.business_price;
          break;
        case 'FirstClass':
          currentBasePrice = selectedFlight.first_class_price;
          break;
        default:
          currentBasePrice = selectedFlight.economy_price;
      }
      setPricePerSelectedClass(currentBasePrice);

      const departureTime = new Date(selectedFlight.departure_time);
      const now = new Date();
      const timeUntilDepartureMs = departureTime.getTime() - now.getTime();
      const timeUntilDepartureMinutes = timeUntilDepartureMs / (1000 * 60);

      let calculatedPrice = currentBasePrice * seats;
      let premium = false;

      if (timeUntilDepartureMinutes < 15) {
        premium = false;
      } else if (timeUntilDepartureMinutes < 60 && timeUntilDepartureMinutes >= 15) {
        calculatedPrice = calculatedPrice * 1.30;
        premium = true;
      }

      setDisplayTotalPrice(calculatedPrice);
      setIsPremiumBooking(premium);
    }
  }, [selectedFlight, seats, travelClass]);

  useEffect(() => {
    // When seats changes, update travellerDetails array
    setTravellerDetails((prev) => {
      const arr = [...prev];
      while (arr.length < seats) arr.push({ title: '', first_name: '', last_name: '', dob: null, nationality: 'India' });
      while (arr.length > seats) arr.pop();
      return arr;
    });
  }, [seats]);

  const fetchFlights = async () => {
    try {
      const token = localStorage.getItem('token');
      let url = `${API_BASE_URL}/flights`;
      const query = new URLSearchParams();
      if (searchParams) {
        if (searchParams.from) query.append('origin', searchParams.from);
        if (searchParams.to) query.append('destination', searchParams.to);
        if (searchParams.departureDate) query.append('departure_time', searchParams.departureDate);
        if (searchParams.travelClass) query.append('travel_class', searchParams.travelClass);
      }
      if (!showAll) {
        query.append('page', page);
        query.append('limit', limit);
      } else {
        query.append('allFlights', 'true');
      }
      if (Array.from(query).length > 0) {
        url = `${url}?${query.toString()}`;
      }
      const response = await axios.get(url, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setFlights(response.data.data);
      setTotal(response.data.meta?.total || 0);
    } catch (error) {
      console.error('Error fetching flights:', error);
    }
  };

  const handleTravellerChange = (idx, field, value) => {
    setTravellerDetails((prev) => {
      const arr = [...prev];
      arr[idx] = { ...arr[idx], [field]: value };
      return arr;
    });
  };

  const validateTravellers = () => {
    for (let t of travellerDetails) {
      if (!t.title || !t.first_name || !t.last_name || !t.dob || !t.nationality) return false;
    }
    return true;
  };

  const handleBooking = async () => {
    if (!validateTravellers()) {
      setBookingError('Please fill all traveller details.');
      return;
    }
    // If premium surcharge applies, show confirmation dialog first
    if (isPremiumBooking && !pendingBooking) {
      setShowSurchargeDialog(true);
      setPendingBooking(true);
      return;
    }
    setPendingBooking(false);
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        setBookingError('Please login to book flights');
        return;
      }
      await axios.post(
        `${API_BASE_URL}/bookings`,
        {
          flight_id: selectedFlight.ID,
          seats: seats,
          travel_class: travelClass,
          travellers: travellerDetails,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      setOpen(false);
      setSeats(1);
      setTravelClass('Economy');
      setBookingError('');
      fetchFlights(); // Refresh flight list to update available seats
    } catch (error) {
      setBookingError(error.response?.data?.error || 'Failed to book flight');
    }
  };

  const handleClickOpen = (flight) => {
    setSelectedFlight(flight);
    setOpen(true);
    setSeats(1);
    setTravelClass('Economy');
    setBookingError('');
  };

  const handleClose = () => {
    setOpen(false);
    setSeats(1);
    setTravelClass('Economy');
    setBookingError('');
  };

  const getAvailableSeatsForClass = () => {
    if (!selectedFlight) return 0;
    switch (travelClass) {
      case 'Economy':
        return selectedFlight.economy_seats;
      case 'PremiumEconomy':
        return selectedFlight.premium_economy_seats;
      case 'Business':
        return selectedFlight.business_seats;
      case 'FirstClass':
        return selectedFlight.first_class_seats;
      default:
        return 0;
    }
  };

  const totalPages = Math.ceil(total / limit);

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        Available Flights
      </Typography>
      <FormControlLabel
        control={<Checkbox checked={showAll} onChange={e => { setShowAll(e.target.checked); setPage(1); }} />}
        label="Show All Flights"
        sx={{ mb: 2 }}
      />
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Flight Number</TableCell>
              <TableCell>Origin</TableCell>
              <TableCell>Destination</TableCell>
              <TableCell>Departure</TableCell>
              <TableCell>Arrival</TableCell>
              <TableCell>Price (₹)</TableCell>
              <TableCell>Available Seats</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {(flights
              .filter(flight => {
                const travellers = searchParams?.travellers ? parseInt(searchParams.travellers) : 1;
                const travelClass = searchParams?.travelClass || 'Economy';
                let availableSeats = 0;
                switch (travelClass) {
                  case 'Economy':
                    availableSeats = flight.economy_seats;
                    break;
                  case 'PremiumEconomy':
                    availableSeats = flight.premium_economy_seats;
                    break;
                  case 'Business':
                    availableSeats = flight.business_seats;
                    break;
                  case 'FirstClass':
                    availableSeats = flight.first_class_seats;
                    break;
                  default:
                    availableSeats = flight.economy_seats;
                }
                return availableSeats >= travellers;
              })
            ).map((flight) => (
              <TableRow key={flight.ID}>
                <TableCell>{flight.ID}</TableCell>
                <TableCell>{flight.origin}</TableCell>
                <TableCell>{flight.destination}</TableCell>
                <TableCell>
                  {new Date(flight.departure_time).toLocaleString()}
                </TableCell>
                <TableCell>
                  {new Date(flight.arrival_time).toLocaleString()}
                </TableCell>
                <TableCell>{flight.price}</TableCell>
                <TableCell>
                  {searchParams?.travelClass === 'Economy' && flight.economy_seats}
                  {searchParams?.travelClass === 'PremiumEconomy' && flight.premium_economy_seats}
                  {searchParams?.travelClass === 'Business' && flight.business_seats}
                  {searchParams?.travelClass === 'FirstClass' && flight.first_class_seats}
                  {!searchParams?.travelClass && flight.economy_seats}
                </TableCell>
                <TableCell>
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={() => handleClickOpen(flight)}
                    disabled={flight.seats === 0}
                  >
                    Book
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      {!showAll && (
        <Box sx={{ display: 'flex', alignItems: 'center', justifyContent: 'center', mt: 2 }}>
          <Button onClick={() => setPage(page - 1)} disabled={page === 1} sx={{ mr: 2 }}>Previous</Button>
          <span>Page</span>
          <TextField
            type="number"
            value={page}
            onChange={e => {
              let val = parseInt(e.target.value) || 1;
              if (val < 1) val = 1;
              if (val > totalPages) val = totalPages;
              setPage(val);
            }}
            size="small"
            sx={{ width: 60, mx: 1 }}
            inputProps={{ min: 1, max: totalPages }}
          />
          <span>of {totalPages}</span>
          <Button onClick={() => setPage(page + 1)} disabled={page === totalPages} sx={{ ml: 2 }}>Next</Button>
        </Box>
      )}

      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>Book Flight</DialogTitle>
        <DialogContent>
          <Typography variant="h6" gutterBottom>
            Flight Details
          </Typography>
          <Typography variant="body1" gutterBottom>
            Flight Number: {selectedFlight?.ID}
          </Typography>
          <Typography variant="body1" gutterBottom>
            From: {selectedFlight?.origin} To: {selectedFlight?.destination}
          </Typography>
          <Typography variant="body1" gutterBottom>
            Departure: {selectedFlight?.departure_time}
          </Typography>
          <Typography variant="body1" gutterBottom>
            Price per seat: ₹{pricePerSelectedClass.toFixed(2)}
          </Typography>

          <FormControl fullWidth margin="normal">
            <InputLabel id="travel-class-label">Travel Class</InputLabel>
            <Select
              labelId="travel-class-label"
              id="travel-class-select"
              value={travelClass}
              label="Travel Class"
              onChange={(e) => {
                setTravelClass(e.target.value);
                setSeats(1);
              }}
            >
              {selectedFlight?.economy_seats > 0 && selectedFlight?.economy_price > 0 && (
                <MenuItem value="Economy">Economy ({selectedFlight.economy_seats} available)</MenuItem>
              )}
              {selectedFlight?.premium_economy_seats > 0 && selectedFlight?.premium_economy_price > 0 && (
                <MenuItem value="PremiumEconomy">Premium Economy ({selectedFlight.premium_economy_seats} available)</MenuItem>
              )}
              {selectedFlight?.business_seats > 0 && selectedFlight?.business_price > 0 && (
                <MenuItem value="Business">Business ({selectedFlight.business_seats} available)</MenuItem>
              )}
              {selectedFlight?.first_class_seats > 0 && selectedFlight?.first_class_price > 0 && (
                <MenuItem value="FirstClass">First Class ({selectedFlight.first_class_seats} available)</MenuItem>
              )}
            </Select>
          </FormControl>

          <TextField
            label="Number of Seats"
            type="number"
            value={seats}
            onChange={(e) => setSeats(Math.max(1, Math.min(parseInt(e.target.value) || 1, getAvailableSeatsForClass())))}
            fullWidth
            margin="normal"
            InputProps={{ inputProps: { min: 1, max: getAvailableSeatsForClass() } }}
          />

          {/* Traveller Details */}
          <Typography variant="h6" sx={{ mt: 2 }}>Traveller Details</Typography>
          {travellerDetails.map((trav, idx) => (
            <Box key={idx} sx={{ border: '1px solid #eee', borderRadius: 2, p: 2, mb: 2 }}>
              <Typography variant="subtitle1">Traveller {idx + 1}</Typography>
              <FormControl fullWidth margin="dense">
                <InputLabel>Title</InputLabel>
                <Select
                  value={trav.title}
                  label="Title"
                  onChange={e => handleTravellerChange(idx, 'title', e.target.value)}
                >
                  {TITLES.map(t => <MenuItem key={t} value={t}>{t}</MenuItem>)}
                </Select>
              </FormControl>
              <TextField
                label="First & Middle Name"
                value={trav.first_name}
                onChange={e => handleTravellerChange(idx, 'first_name', e.target.value)}
                fullWidth
                margin="dense"
              />
              <TextField
                label="Last Name"
                value={trav.last_name}
                onChange={e => handleTravellerChange(idx, 'last_name', e.target.value)}
                fullWidth
                margin="dense"
              />
              <DatePicker
                label="Date of Birth"
                value={trav.dob}
                onChange={date => handleTravellerChange(idx, 'dob', date)}
                renderInput={params => <TextField {...params} fullWidth margin="dense" />}
              />
              <FormControl fullWidth margin="dense">
                <InputLabel>Nationality</InputLabel>
                <Select
                  value={trav.nationality}
                  label="Nationality"
                  onChange={e => handleTravellerChange(idx, 'nationality', e.target.value)}
                >
                  {NATIONALITIES.map(n => <MenuItem key={n} value={n}>{n}</MenuItem>)}
                </Select>
              </FormControl>
            </Box>
          ))}

          <Typography variant="h6" gutterBottom>
            Total Price: ₹{displayTotalPrice.toFixed(2)}
          </Typography>
          {isPremiumBooking && (
            <Typography variant="body2" color="textSecondary" sx={{ mt: 1 }}>
              (Includes 30% premium charge for last-minute booking)
            </Typography>
          )}
          {bookingError && (
            <Alert severity="error" sx={{ mt: 2 }}>
              {bookingError}
            </Alert>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleBooking} variant="contained" color="primary">
            Confirm Booking
          </Button>
        </DialogActions>
      </Dialog>
      {/* Surcharge confirmation dialog */}
      <Dialog open={showSurchargeDialog} onClose={() => { setShowSurchargeDialog(false); setPendingBooking(false); }}>
        <DialogTitle>Premium Surcharge Notice</DialogTitle>
        <DialogContent>
          <Typography gutterBottom>
            A 30% premium charge is being applied to your booking because it is a last-minute booking (within 15-60 minutes of departure).
          </Typography>
          <Typography gutterBottom>
            Total Price: ₹{displayTotalPrice.toFixed(2)}
          </Typography>
          <Typography variant="body2" color="textSecondary">
            Do you want to proceed with this booking?
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => { setShowSurchargeDialog(false); setPendingBooking(false); }}>Cancel</Button>
          <Button onClick={() => { setShowSurchargeDialog(false); handleBooking(); }} variant="contained" color="primary">Confirm</Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default FlightList;