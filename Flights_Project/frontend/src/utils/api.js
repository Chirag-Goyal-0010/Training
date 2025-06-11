import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add a request interceptor to add the auth token to requests
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Add a response interceptor to handle common errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response) {
      // The request was made and the server responded with a status code
      // that falls out of the range of 2xx
      if (error.response.status === 401) {
        // Handle unauthorized access
        localStorage.removeItem('token');
        window.location.href = '/login';
      }
      return Promise.reject(error.response.data);
    } else if (error.request) {
      // The request was made but no response was received
      return Promise.reject({ message: 'No response from server' });
    } else {
      // Something happened in setting up the request that triggered an Error
      return Promise.reject({ message: error.message });
    }
  }
);

// Auth API calls
export const authAPI = {
  login: (username, password) =>
    api.post('/auth/login', { username, password }),
  register: (username, password, role) =>
    api.post('/auth/register', { username, password, role }),
  verify: () => api.get('/auth/verify'),
};

// Flight API calls
export const flightAPI = {
  getFlights: () => api.get('/flights'),
  getFlight: (id) => api.get(`/flights/${id}`),
  createFlight: (flightData) => api.post('/admin/flights', flightData),
  updateFlight: (id, flightData) => api.put(`/admin/flights/${id}`, flightData),
  deleteFlight: (id) => api.delete(`/admin/flights/${id}`),
};

export default api; 