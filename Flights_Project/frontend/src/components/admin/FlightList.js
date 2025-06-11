import React, { useState, useEffect } from 'react';
import {
  Box,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  IconButton,
  Typography,
} from '@mui/material';
import { DataGrid } from '@mui/x-data-grid';
import { Edit as EditIcon, Delete as DeleteIcon } from '@mui/icons-material';
import axios from 'axios';
import { format } from 'date-fns';

const FlightList = () => {
  const [flights, setFlights] = useState([]);
  const [open, setOpen] = useState(false);
  const [selectedFlight, setSelectedFlight] = useState(null);
  const [formData, setFormData] = useState({
    aircraft_id: '',
    departure_airport_id: '',
    arrival_airport_id: '',
    departure_time: '',
    arrival_time: '',
    distance: '',
    status: '',
  });

  const columns = [
    { field: 'id', headerName: 'ID', width: 70 },
    { field: 'aircraft_id', headerName: 'Aircraft ID', width: 100 },
    { field: 'departure_airport', headerName: 'Departure', width: 150 },
    { field: 'arrival_airport', headerName: 'Arrival', width: 150 },
    {
      field: 'departure_time',
      headerName: 'Departure Time',
      width: 180,
      valueFormatter: (params) => format(new Date(params.value), 'PPpp'),
    },
    {
      field: 'arrival_time',
      headerName: 'Arrival Time',
      width: 180,
      valueFormatter: (params) => format(new Date(params.value), 'PPpp'),
    },
    { field: 'distance', headerName: 'Distance (km)', width: 120 },
    { field: 'status', headerName: 'Status', width: 120 },
    {
      field: 'actions',
      headerName: 'Actions',
      width: 120,
      renderCell: (params) => (
        <Box>
          <IconButton onClick={() => handleEdit(params.row)}>
            <EditIcon />
          </IconButton>
          <IconButton onClick={() => handleDelete(params.row.id)}>
            <DeleteIcon />
          </IconButton>
        </Box>
      ),
    },
  ];

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/flights', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      });
      setFlights(response.data);
    } catch (error) {
      console.error('Error fetching flights:', error);
    }
  };

  const handleEdit = (flight) => {
    setSelectedFlight(flight);
    setFormData({
      aircraft_id: flight.aircraft_id,
      departure_airport_id: flight.departure_airport_id,
      arrival_airport_id: flight.arrival_airport_id,
      departure_time: flight.departure_time,
      arrival_time: flight.arrival_time,
      distance: flight.distance,
      status: flight.status,
    });
    setOpen(true);
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this flight?')) {
      try {
        await axios.delete(`http://localhost:8080/api/admin/flights/${id}`, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
        fetchFlights();
      } catch (error) {
        console.error('Error deleting flight:', error);
      }
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (selectedFlight) {
        await axios.put(
          `http://localhost:8080/api/admin/flights/${selectedFlight.id}`,
          formData,
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem('token')}`,
            },
          }
        );
      } else {
        await axios.post('http://localhost:8080/api/admin/flights', formData, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        });
      }
      setOpen(false);
      fetchFlights();
    } catch (error) {
      console.error('Error saving flight:', error);
    }
  };

  return (
    <Box sx={{ height: 600, width: '100%', p: 2 }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 2 }}>
        <Typography variant="h5">Flight Management</Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={() => {
            setSelectedFlight(null);
            setFormData({
              aircraft_id: '',
              departure_airport_id: '',
              arrival_airport_id: '',
              departure_time: '',
              arrival_time: '',
              distance: '',
              status: '',
            });
            setOpen(true);
          }}
        >
          Add Flight
        </Button>
      </Box>

      <DataGrid
        rows={flights}
        columns={columns}
        pageSize={10}
        rowsPerPageOptions={[10]}
        checkboxSelection
        disableSelectionOnClick
      />

      <Dialog open={open} onClose={() => setOpen(false)}>
        <DialogTitle>
          {selectedFlight ? 'Edit Flight' : 'Add New Flight'}
        </DialogTitle>
        <DialogContent>
          <Box component="form" onSubmit={handleSubmit} sx={{ mt: 2 }}>
            <TextField
              fullWidth
              label="Aircraft ID"
              value={formData.aircraft_id}
              onChange={(e) =>
                setFormData({ ...formData, aircraft_id: e.target.value })
              }
              margin="normal"
            />
            <TextField
              fullWidth
              label="Departure Airport ID"
              value={formData.departure_airport_id}
              onChange={(e) =>
                setFormData({ ...formData, departure_airport_id: e.target.value })
              }
              margin="normal"
            />
            <TextField
              fullWidth
              label="Arrival Airport ID"
              value={formData.arrival_airport_id}
              onChange={(e) =>
                setFormData({ ...formData, arrival_airport_id: e.target.value })
              }
              margin="normal"
            />
            <TextField
              fullWidth
              label="Departure Time"
              type="datetime-local"
              value={formData.departure_time}
              onChange={(e) =>
                setFormData({ ...formData, departure_time: e.target.value })
              }
              margin="normal"
              InputLabelProps={{ shrink: true }}
            />
            <TextField
              fullWidth
              label="Arrival Time"
              type="datetime-local"
              value={formData.arrival_time}
              onChange={(e) =>
                setFormData({ ...formData, arrival_time: e.target.value })
              }
              margin="normal"
              InputLabelProps={{ shrink: true }}
            />
            <TextField
              fullWidth
              label="Distance"
              type="number"
              value={formData.distance}
              onChange={(e) =>
                setFormData({ ...formData, distance: e.target.value })
              }
              margin="normal"
            />
            <TextField
              fullWidth
              label="Status"
              value={formData.status}
              onChange={(e) =>
                setFormData({ ...formData, status: e.target.value })
              }
              margin="normal"
            />
          </Box>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>Cancel</Button>
          <Button onClick={handleSubmit} variant="contained" color="primary">
            {selectedFlight ? 'Update' : 'Add'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default FlightList; 