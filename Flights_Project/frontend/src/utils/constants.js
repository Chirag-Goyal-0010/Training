export const FLIGHT_STATUS = {
  SCHEDULED: 'scheduled',
  DELAYED: 'delayed',
  DEPARTED: 'departed',
  ARRIVED: 'arrived',
  CANCELLED: 'cancelled',
};

export const FLIGHT_STATUS_LABELS = {
  [FLIGHT_STATUS.SCHEDULED]: 'Scheduled',
  [FLIGHT_STATUS.DELAYED]: 'Delayed',
  [FLIGHT_STATUS.DEPARTED]: 'Departed',
  [FLIGHT_STATUS.ARRIVED]: 'Arrived',
  [FLIGHT_STATUS.CANCELLED]: 'Cancelled',
};

export const FLIGHT_STATUS_COLORS = {
  [FLIGHT_STATUS.SCHEDULED]: '#1976d2', // Blue
  [FLIGHT_STATUS.DELAYED]: '#f57c00', // Orange
  [FLIGHT_STATUS.DEPARTED]: '#388e3c', // Green
  [FLIGHT_STATUS.ARRIVED]: '#388e3c', // Green
  [FLIGHT_STATUS.CANCELLED]: '#d32f2f', // Red
};

export const USER_ROLES = {
  ADMIN: 'admin',
  USER: 'user',
};

export const USER_ROLE_LABELS = {
  [USER_ROLES.ADMIN]: 'Administrator',
  [USER_ROLES.USER]: 'User',
};

export const API_ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    REGISTER: '/auth/register',
    VERIFY: '/auth/verify',
  },
  FLIGHTS: {
    LIST: '/flights',
    DETAIL: (id) => `/flights/${id}`,
    CREATE: '/admin/flights',
    UPDATE: (id) => `/admin/flights/${id}`,
    DELETE: (id) => `/admin/flights/${id}`,
  },
};

export const DATE_FORMATS = {
  DISPLAY: 'MMM dd, yyyy HH:mm',
  DATE_ONLY: 'MMM dd, yyyy',
  TIME_ONLY: 'HH:mm',
  API: "yyyy-MM-dd'T'HH:mm:ss.SSS'Z'",
};

export const ERROR_MESSAGES = {
  NETWORK_ERROR: 'Network error. Please check your connection.',
  UNAUTHORIZED: 'Unauthorized access. Please log in.',
  FORBIDDEN: 'You do not have permission to perform this action.',
  NOT_FOUND: 'The requested resource was not found.',
  SERVER_ERROR: 'Server error. Please try again later.',
  VALIDATION_ERROR: 'Please check your input and try again.',
};

export const SUCCESS_MESSAGES = {
  LOGIN: 'Successfully logged in.',
  REGISTER: 'Successfully registered.',
  LOGOUT: 'Successfully logged out.',
  FLIGHT_CREATED: 'Flight created successfully.',
  FLIGHT_UPDATED: 'Flight updated successfully.',
  FLIGHT_DELETED: 'Flight deleted successfully.',
}; 