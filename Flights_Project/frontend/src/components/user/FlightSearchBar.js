import React, { useState } from 'react';
import {
  Box, Tabs, Tab, TextField, Button, Checkbox, FormControlLabel,
  ToggleButton, ToggleButtonGroup, Typography, Paper
} from '@mui/material';
import { DatePicker } from '@mui/lab';
import SearchIcon from '@mui/icons-material/Search';

const specialFares = [
  { label: 'Student', value: 'student' },
  { label: 'Senior Citizen', value: 'senior' },
  { label: 'Armed Forces', value: 'armed' },
];

export default function FlightSearchBar() {
  const [tripType, setTripType] = useState('oneway');
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [departure, setDeparture] = useState(null);
  const [returnDate, setReturnDate] = useState(null);
  const [travellers, setTravellers] = useState(1);
  const [travelClass, setTravelClass] = useState('Economy');
  const [special, setSpecial] = useState([]);
  const [addHotel, setAddHotel] = useState(false);

  return (
    <Paper elevation={3} sx={{ p: 2, borderRadius: 4, maxWidth: 1100, mx: 'auto', mt: 3 }}>
      <Tabs
        value={tripType}
        onChange={(_, v) => setTripType(v)}
        sx={{ mb: 2 }}
      >
        <Tab label="One Way" value="oneway" />
        <Tab label="Round Trip" value="roundtrip" />
      </Tabs>
      <Box display="flex" gap={2} alignItems="center" flexWrap="wrap">
        <TextField
          label="From"
          value={from}
          onChange={e => setFrom(e.target.value)}
          sx={{ minWidth: 150, flex: 1 }}
        />
        <TextField
          label="To"
          value={to}
          onChange={e => setTo(e.target.value)}
          sx={{ minWidth: 150, flex: 1 }}
        />
        <DatePicker
          label="Departure"
          value={departure}
          onChange={setDeparture}
          renderInput={(params) => <TextField {...params} sx={{ minWidth: 170 }} />}
        />
        {tripType === 'roundtrip' && (
          <DatePicker
            label="Return"
            value={returnDate}
            onChange={setReturnDate}
            renderInput={(params) => <TextField {...params} sx={{ minWidth: 170 }} />}
          />
        )}
        <TextField
          label="Travellers & Class"
          value={`${travellers} Traveller${travellers > 1 ? 's' : ''}, ${travelClass}`}
          InputProps={{
            readOnly: true,
          }}
          sx={{ minWidth: 200 }}
        />
        <Button
          variant="contained"
          color="warning"
          size="large"
          sx={{ px: 4, borderRadius: 3, fontWeight: 600 }}
          endIcon={<SearchIcon />}
        >
          Search
        </Button>
      </Box>
      <Box display="flex" alignItems="center" mt={2} gap={2} flexWrap="wrap">
        <Typography variant="body2" sx={{ fontWeight: 500 }}>
          Special Fares <span style={{ color: '#888' }}>(Optional)</span>:
        </Typography>
        <ToggleButtonGroup
          value={special}
          onChange={(_, v) => setSpecial(v)}
          aria-label="special fares"
          size="small"
        >
          {specialFares.map(fare => (
            <ToggleButton key={fare.value} value={fare.value}>
              {fare.label}
            </ToggleButton>
          ))}
        </ToggleButtonGroup>
        <FormControlLabel
          control={
            <Checkbox
              checked={addHotel}
              onChange={e => setAddHotel(e.target.checked)}
              sx={{ ml: 2 }}
            />
          }
          label={<span style={{ color: '#1976d2', fontWeight: 500 }}>Add hotel and save up to 50%</span>}
        />
      </Box>
    </Paper>
  );
} 