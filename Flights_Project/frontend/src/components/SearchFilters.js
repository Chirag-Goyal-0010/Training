import React, { useState } from 'react';
import {
  Box,
  Paper,
  Typography,
  Slider,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Grid,
  Button,
} from '@mui/material';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';

const SearchFilters = ({ onFilterChange }) => {
  const [filters, setFilters] = useState({
    origin: '',
    destination: '',
    date: null,
    priceRange: [0, 1000],
    sortBy: 'price',
    airlines: [],
  });

  const handleFilterChange = (event) => {
    const { name, value } = event.target;
    setFilters((prevFilters) => ({
      ...prevFilters,
      [name]: value,
    }));
  };

  const handleDateChange = (newValue) => {
    setFilters((prevFilters) => ({
      ...prevFilters,
      date: newValue,
    }));
  };

  const handlePriceRangeChange = (event, newValue) => {
    setFilters((prevFilters) => ({
      ...prevFilters,
      priceRange: newValue,
    }));
  };

  const handleSortChange = (event) => {
    setFilters((prevFilters) => ({
      ...prevFilters,
      sortBy: event.target.value,
    }));
  };

  const handleSearch = () => {
    onFilterChange(filters);
  };

  return (
    <Paper elevation={3} sx={{ p: 3, mb: 3 }}>
      <Typography variant="h6" gutterBottom>
        Search Filters
      </Typography>
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <TextField
            label="Origin"
            name="origin"
            value={filters.origin}
            onChange={handleFilterChange}
            fullWidth
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <TextField
            label="Destination"
            name="destination"
            value={filters.destination}
            onChange={handleFilterChange}
            fullWidth
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <LocalizationProvider dateAdapter={AdapterDateFns}>
            <DatePicker
              label="Departure Date"
              value={filters.date}
              onChange={handleDateChange}
              renderInput={(params) => <TextField {...params} fullWidth />}
            />
          </LocalizationProvider>
        </Grid>

        <Grid item xs={12} md={6}>
          <Typography gutterBottom>Price Range</Typography>
          <Slider
            value={filters.priceRange}
            onChange={handlePriceRangeChange}
            valueLabelDisplay="auto"
            min={0}
            max={1000}
            marks={[
              { value: 0, label: '$0' },
              { value: 1000, label: '$1000' },
            ]}
          />
        </Grid>
        
        <Grid item xs={12} md={6}>
          <FormControl fullWidth>
            <InputLabel>Sort By</InputLabel>
            <Select
              value={filters.sortBy}
              label="Sort By"
              onChange={handleSortChange}
            >
              <MenuItem value="price">Price (Low to High)</MenuItem>
              <MenuItem value="price-desc">Price (High to Low)</MenuItem>
              <MenuItem value="duration">Duration</MenuItem>
              <MenuItem value="departure">Departure Time</MenuItem>
            </Select>
          </FormControl>
        </Grid>

        <Grid item xs={12}>
          <Button variant="contained" color="primary" onClick={handleSearch} fullWidth>
            Search Flights
          </Button>
        </Grid>
      </Grid>
    </Paper>
  );
};

export default SearchFilters; 