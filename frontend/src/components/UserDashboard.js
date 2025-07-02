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
  Alert,
  Box,
  Chip,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Checkbox,
  FormGroup,
  FormControlLabel,
  Divider,
} from '@mui/material';
import { format } from 'date-fns';

function UserDashboard() {
  const [bookings, setBookings] = useState([]);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [limit, setLimit] = useState(10);
  const [confirmOpen, setConfirmOpen] = useState(false);
  const [bookingToCancel, setBookingToCancel] = useState(null);
  const [selectedTravellers, setSelectedTravellers] = useState([]);
  const [dialogBooking, setDialogBooking] = useState(null);

  useEffect(() => {
    fetchBookings(page);
  }, [page]);

  const fetchBookings = async (pageNum = 1) => {
    try {
      const token = localStorage.getItem('token');
      const response = await axios.get(`${API_BASE_URL}/bookings?page=${pageNum}&limit=${limit}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setBookings(response.data.data);
      setTotal(response.data.meta?.total || 0);
    } catch (error) {
      setError('Failed to fetch bookings');
    }
  };

  const getStatusColor = (status) => {
    switch (status.toLowerCase()) {
      case 'confirmed':
        return 'success';
      case 'cancelled':
        return 'error';
      default:
        return 'default';
    }
  };

  const totalPages = Math.ceil(total / limit);

  const openConfirmDialog = (bookingId) => {
    const booking = bookings.find(b => b.ID === bookingId);
    setBookingToCancel(bookingId);
    setDialogBooking(booking);
    setSelectedTravellers(booking?.travellers?.map(t => t.id || t.ID) || []);
    setConfirmOpen(true);
  };

  const closeConfirmDialog = () => {
    setConfirmOpen(false);
    setBookingToCancel(null);
    setDialogBooking(null);
    setSelectedTravellers([]);
  };

  const handleTravellerToggle = (travellerId) => {
    setSelectedTravellers(prev =>
      prev.includes(travellerId)
        ? prev.filter(id => id !== travellerId)
        : [...prev, travellerId]
    );
  };

  const getDialogPriceInfo = () => {
    if (!dialogBooking || !dialogBooking.travellers) return { original: 0, fee: 0, refund: 0 };
    const totalSeats = dialogBooking.travellers.length;
    const perSeat = dialogBooking.total_price / totalSeats;
    const toCancel = selectedTravellers.length;
    const cancelTotal = perSeat * toCancel;
    const fee = 0.10 * cancelTotal;
    const refund = cancelTotal - fee;
    return {
      original: dialogBooking.total_price,
      fee,
      refund,
      toCancel,
      totalSeats,
      perSeat
    };
  };

  const handleCancel = async () => {
    try {
      const token = localStorage.getItem('token');
      await axios.delete(`${API_BASE_URL}/bookings/${bookingToCancel}`, {
        headers: { Authorization: `Bearer ${token}` },
        data: { traveller_ids: selectedTravellers },
      });
      setSuccess('Booking cancelled successfully');
      closeConfirmDialog();
      fetchBookings(page);
    } catch (error) {
      setError(error.response?.data?.error || 'Failed to cancel booking');
      closeConfirmDialog();
    }
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        My Bookings
      </Typography>

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
      {success && (
        <Alert severity="success" sx={{ mb: 2 }}>
          {success}
        </Alert>
      )}

      <Dialog open={confirmOpen} onClose={closeConfirmDialog} maxWidth="xs" fullWidth>
        <DialogTitle>Cancel Booking</DialogTitle>
        <DialogContent>
          {dialogBooking && dialogBooking.travellers && dialogBooking.travellers.length > 1 ? (
            <>
              <DialogContentText>Select the individual(s) to cancel:</DialogContentText>
              <FormGroup>
                {dialogBooking.travellers.map(trav => (
                  <FormControlLabel
                    key={trav.id || trav.ID}
                    control={
                      <Checkbox
                        checked={selectedTravellers.includes(trav.id || trav.ID)}
                        onChange={() => handleTravellerToggle(trav.id || trav.ID)}
                      />
                    }
                    label={`${trav.title} ${trav.first_name} ${trav.last_name}`}
                  />
                ))}
              </FormGroup>
              <Divider sx={{ my: 2 }} />
              <DialogContentText>
                <b>Original Price:</b> ₹{getDialogPriceInfo().original.toFixed(2)}<br />
                <b>Cancellation Fee (10%):</b> ₹{getDialogPriceInfo().fee.toFixed(2)}<br />
                <b>Refund Amount:</b> ₹{getDialogPriceInfo().refund.toFixed(2)}
              </DialogContentText>
            </>
          ) : (
            <DialogContentText>
              Are you sure you want to cancel this booking? A 10% cancellation fee will be applied and only 90% of the total booking amount will be refunded.
            </DialogContentText>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={closeConfirmDialog} color="primary">No</Button>
          <Button
            onClick={handleCancel}
            color="error"
            variant="contained"
            disabled={dialogBooking && dialogBooking.travellers && dialogBooking.travellers.length > 1 && selectedTravellers.length === 0}
          >
            Yes, Cancel
          </Button>
        </DialogActions>
      </Dialog>

      {bookings.length === 0 ? (
        <Box sx={{ mt: 4, textAlign: 'center' }}>
          <Typography variant="h6" color="text.secondary">
            You haven't made any bookings yet.
          </Typography>
        </Box>
      ) : (
        <>
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Booking ID</TableCell>
                <TableCell>Flight Number</TableCell>
                <TableCell>Route</TableCell>
                <TableCell>Travel Date</TableCell>
                <TableCell>Seats</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Booking Date</TableCell>
                <TableCell>Refund</TableCell>
                <TableCell>Action</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bookings.map((booking) => {
                return (
                <TableRow key={booking.ID}>
                  <TableCell>#{booking.ID}</TableCell>
                  <TableCell>{booking.Flight?.ID}</TableCell>
                  <TableCell>
                    {booking.Flight?.origin} → {booking.Flight?.destination}
                  </TableCell>
                  <TableCell>
                    {booking.Flight?.departure_time ? format(new Date(booking.Flight.departure_time), 'dd/MM/yyyy HH:mm') : 'N/A'}
                  </TableCell>
                  <TableCell>{booking.seats}</TableCell>
                  <TableCell>
                    <Chip
                      label={booking.status}
                      color={getStatusColor(booking.status)}
                      size="small"
                    />
                  </TableCell>
                  <TableCell>
                    {format(new Date(booking.booking_date), 'dd/MM/yyyy HH:mm')}
                  </TableCell>
                  <TableCell>
                    {(booking.refund_amount > 0) && (
                      <>
                        <div>Refund: ₹{booking.refund_amount?.toFixed(2) || 0}</div>
                        {booking.status === 'Confirmed' && (
                          <div style={{fontSize: '0.85em', color: '#888'}}>
                            Partial refund so far
                          </div>
                        )}
                        {booking.status === 'Cancelled' && (
                          <div style={{fontSize: '0.85em', color: '#888'}}>
                            {booking.cancellation_date ? `Cancelled: ${format(new Date(booking.cancellation_date), 'dd/MM/yyyy HH:mm')}` : ''}
                          </div>
                        )}
                      </>
                    )}
                  </TableCell>
                  <TableCell>
                    {booking.status === 'Confirmed' && (() => {
                      const departure = booking.Flight?.departure_time ? new Date(booking.Flight.departure_time) : null;
                      const now = new Date();
                      const twoHoursMs = 2 * 60 * 60 * 1000;
                      if (departure && (departure.getTime() - now.getTime() > twoHoursMs)) {
                        return (
                          <Button
                            variant="outlined"
                            color="error"
                            size="small"
                            onClick={() => openConfirmDialog(booking.ID)}
                          >
                            Cancel
                          </Button>
                        );
                      } else {
                        return null;
                      }
                    })()}
                  </TableCell>
                </TableRow>
              )})}
            </TableBody>
          </Table>
        </TableContainer>
        <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', mt: 2 }}>
          <button onClick={() => setPage(page - 1)} disabled={page === 1} style={{marginRight: 8}}>
            Previous
          </button>
          <span>Page {page} of {totalPages}</span>
          <button onClick={() => setPage(page + 1)} disabled={page === totalPages} style={{marginLeft: 8}}>
            Next
          </button>
        </Box>
        </>
      )}
    </Container>
  );
}

export default UserDashboard;