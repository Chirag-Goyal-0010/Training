import React, { createContext, useState, useContext, useCallback, ReactNode } from 'react';
import { bookingAPI, IBooking, ICreateBookingRequest, IUpdateBookingStatusRequest } from '../api';

interface IBookingsContext {
  bookings: IBooking[];
  loadingBookings: boolean;
  bookingsError: string | null;
  fetchUserBookings: (userId: number) => Promise<void>;
  fetchAllBookings: () => Promise<void>;
  createBooking: (bookingData: ICreateBookingRequest) => Promise<{ success: boolean; data?: IBooking; error?: string }>;
  cancelBooking: (id: number) => Promise<{ success: boolean; error?: string }>;
  updateBookingStatus: (id: number, statusData: IUpdateBookingStatusRequest) => Promise<{ success: boolean; data?: IBooking; error?: string }>;
}

interface BookingsProviderProps {
  children: ReactNode;
}

const BookingsContext = createContext<IBookingsContext | undefined>(undefined);

export const useBookings = () => {
  const context = useContext(BookingsContext);
  if (!context) {
    throw new Error('useBookings must be used within a BookingsProvider');
  }
  return context;
};

export const BookingsProvider = ({ children }: BookingsProviderProps) => {
  const [bookings, setBookings] = useState<IBooking[]>([]);
  const [loadingBookings, setLoadingBookings] = useState<boolean>(false);
  const [bookingsError, setBookingsError] = useState<string | null>(null);

  const fetchUserBookings = useCallback(async (userId: number) => {
    setLoadingBookings(true);
    setBookingsError(null);
    try {
      const response = await bookingAPI.getUserBookings(userId);
      setBookings(response.data);
    } catch (error: any) {
      setBookingsError(error.message || 'Failed to fetch user bookings');
    } finally {
      setLoadingBookings(false);
    }
  }, []);

  const fetchAllBookings = useCallback(async () => {
    setLoadingBookings(true);
    setBookingsError(null);
    try {
      const response = await bookingAPI.getAllBookings();
      setBookings(response.data);
    } catch (error: any) {
      setBookingsError(error.message || 'Failed to fetch all bookings');
    } finally {
      setLoadingBookings(false);
    }
  }, []);

  const createBooking = useCallback(async (bookingData: ICreateBookingRequest) => {
    setLoadingBookings(true);
    setBookingsError(null);
    try {
      const response = await bookingAPI.createBooking(bookingData);
      setBookings((prevBookings: IBooking[]) => [...prevBookings, response.data]);
      return { success: true, data: response.data };
    } catch (error: any) {
      setBookingsError(error.message || 'Failed to create booking');
      return { success: false, error: error.message };
    } finally {
      setLoadingBookings(false);
    }
  }, []);

  const cancelBooking = useCallback(async (id: number) => {
    setLoadingBookings(true);
    setBookingsError(null);
    try {
      await bookingAPI.cancelBooking(id);
      setBookings((prevBookings: IBooking[]) => prevBookings.filter((booking: IBooking) => booking.id !== id));
      return { success: true };
    } catch (error: any) {
      setBookingsError(error.message || 'Failed to cancel booking');
      return { success: false, error: error.message };
    } finally {
      setLoadingBookings(false);
    }
  }, []);

  const updateBookingStatus = useCallback(async (id: number, statusData: IUpdateBookingStatusRequest) => {
    setLoadingBookings(true);
    setBookingsError(null);
    try {
      const response = await bookingAPI.updateBookingStatus(id, statusData);
      setBookings((prevBookings: IBooking[]) =>
        prevBookings.map((booking: IBooking) => (booking.id === id ? response.data : booking))
      );
      return { success: true, data: response.data };
    } catch (error: any) {
      setBookingsError(error.message || 'Failed to update booking status');
      return { success: false, error: error.message };
    } finally {
      setLoadingBookings(false);
    }
  }, []);

  const value = {
    bookings,
    loadingBookings,
    bookingsError,
    fetchUserBookings,
    fetchAllBookings,
    createBooking,
    cancelBooking,
    updateBookingStatus,
  };

  return (
    <BookingsContext.Provider value={value}>
      {children}
    </BookingsContext.Provider>
  );
};

export default BookingsContext; 