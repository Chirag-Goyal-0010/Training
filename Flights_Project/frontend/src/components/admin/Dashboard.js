import React, { useState, useEffect } from 'react';
import { Routes, Route, Link, useNavigate } from 'react-router-dom';
import {
  AppBar,
  Toolbar,
  Typography,
  Container,
  Box,
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from '@mui/material';
import { useAuth } from '../../context/AuthContext';
import AddFlight from './AddFlight';
import EditFlight from './EditFlight';

const Dashboard = () => {
  const [flights, setFlights] = useState([]);
  const { logout } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/flights');
      const data = await response.json();
      setFlights(data);
    } catch (error) {
      console.error('Error fetching flights:', error);
    }
  };

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Admin Dashboard
          </Typography>
          <Button color="inherit" component={Link} to="/admin/add-flight">
            Add Flight
          </Button>
          <Button color="inherit" onClick={handleLogout}>
            Logout
          </Button>
        </Toolbar>
      </AppBar>

      <Container sx={{ mt: 4 }}>
        <Routes>
          <Route path="add-flight" element={<AddFlight onFlightAdded={fetchFlights} />} />
          <Route path="edit-flight/:id" element={<EditFlight onFlightUpdated={fetchFlights} />} />
          <Route
            path="/"
            element={
              <TableContainer component={Paper}>
                <Table>
                  <TableHead>
                    <TableRow>
                      <TableCell>ID</TableCell>
                      <TableCell>Departure Airport</TableCell>
                      <TableCell>Arrival Airport</TableCell>
                      <TableCell>Departure Time</TableCell>
                      <TableCell>Arrival Time</TableCell>
                      <TableCell>Status</TableCell>
                      <TableCell>Actions</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {flights.map((flight) => (
                      <TableRow key={flight.id}>
                        <TableCell>{flight.id}</TableCell>
                        <TableCell>{flight.departure_airport}</TableCell>
                        <TableCell>{flight.arrival_airport}</TableCell>
                        <TableCell>{flight.departure_time}</TableCell>
                        <TableCell>{flight.arrival_time}</TableCell>
                        <TableCell>{flight.status}</TableCell>
                        <TableCell>
                          <Button
                            component={Link}
                            to={`/admin/edit-flight/${flight.id}`}
                            size="small"
                          >
                            Edit
                          </Button>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            }
          />
        </Routes>
      </Container>
    </Box>
  );
};

export default Dashboard; 