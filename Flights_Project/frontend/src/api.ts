import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

interface LoginResponse {
  token: string;
  user: {
    id: number;
    username: string;
    email: string;
    role: string;
  };
}

interface RegisterResponse {
  message: string;
}

export const authAPI = {
  login: async (username: string, password: string): Promise<LoginResponse> => {
    const response = await axios.post(`${API_URL}/auth/login`, {
      username,
      password,
    });
    return response.data;
  },

  register: async (
    username: string,
    email: string,
    password: string,
    role: string
  ): Promise<RegisterResponse> => {
    const response = await axios.post(`${API_URL}/auth/register`, {
      username,
      email,
      password,
      role,
    });
    return response.data;
  },
};

export interface IFlight {
  id: number;
  flightCode: string;
  origin: string;
  destination: string;
  departureTime: string; // ISO string format for date-time
  arrivalTime: string; // ISO string format for date-time
  capacity: number;
  price: number;
}

interface ISearchCriteria {
  origin?: string;
  destination?: string;
  date?: string; // Date string (e.g., 'YYYY-MM-DD')
  id?: number;
}

interface CreateFlightRequest extends Omit<IFlight, 'id'> {}

interface UpdateFlightRequest extends Partial<IFlight> {}

export const flightAPI = {
  getFlights: async (searchCriteria?: ISearchCriteria): Promise<{ data: IFlight[] }> => {
    const response = await axios.get(`${API_URL}/flights`, {
      params: searchCriteria,
    });
    return response.data;
  },

  createFlight: async (flightData: CreateFlightRequest): Promise<{ data: IFlight }> => {
    const response = await axios.post(`${API_URL}/flights`, flightData);
    return response.data;
  },

  updateFlight: async (
    id: number,
    flightData: UpdateFlightRequest
  ): Promise<{ data: IFlight }> => {
    const response = await axios.put(`${API_URL}/flights/${id}`, flightData);
    return response.data;
  },

  deleteFlight: async (id: number): Promise<void> => {
    await axios.delete(`${API_URL}/flights/${id}`);
  },
};

export interface IBooking {
  id: number;
  flightId: number;
  userId: number;
  seatNumber: number;
  bookingStatus: string;
}

export interface ICreateBookingRequest extends Omit<IBooking, 'id' | 'bookingStatus'> {}

export interface IUpdateBookingStatusRequest {
  bookingStatus: string;
}

export const bookingAPI = {
  getUserBookings: async (userId: number): Promise<{ data: IBooking[] }> => {
    const response = await axios.get(`${API_URL}/bookings/user/${userId}`);
    return response.data;
  },
  getAllBookings: async (): Promise<{ data: IBooking[] }> => {
    const response = await axios.get(`${API_URL}/bookings`);
    return response.data;
  },
  createBooking: async (bookingData: ICreateBookingRequest): Promise<{ data: IBooking }> => {
    const response = await axios.post(`${API_URL}/bookings`, bookingData);
    return response.data;
  },
  cancelBooking: async (id: number): Promise<void> => {
    await axios.delete(`${API_URL}/bookings/${id}`);
  },
  updateBookingStatus: async (
    id: number,
    statusData: IUpdateBookingStatusRequest
  ): Promise<{ data: IBooking }> => {
    const response = await axios.patch(`${API_URL}/bookings/${id}/status`, statusData);
    return response.data;
  },
}; 