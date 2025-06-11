export const validateUsername = (username) => {
  if (!username) {
    return 'Username is required';
  }
  if (username.length < 3) {
    return 'Username must be at least 3 characters long';
  }
  if (username.length > 50) {
    return 'Username must be less than 50 characters';
  }
  if (!/^[a-zA-Z0-9_]+$/.test(username)) {
    return 'Username can only contain letters, numbers, and underscores';
  }
  return '';
};

export const validatePassword = (password) => {
  if (!password) {
    return 'Password is required';
  }
  if (password.length < 6) {
    return 'Password must be at least 6 characters long';
  }
  if (password.length > 100) {
    return 'Password must be less than 100 characters';
  }
  return '';
};

export const validateFlightData = (flightData) => {
  const errors = {};

  if (!flightData.aircraftId) {
    errors.aircraftId = 'Aircraft ID is required';
  }

  if (!flightData.departureAirportId) {
    errors.departureAirportId = 'Departure airport is required';
  }

  if (!flightData.arrivalAirportId) {
    errors.arrivalAirportId = 'Arrival airport is required';
  }

  if (!flightData.departureTime) {
    errors.departureTime = 'Departure time is required';
  }

  if (!flightData.arrivalTime) {
    errors.arrivalTime = 'Arrival time is required';
  }

  if (flightData.departureTime && flightData.arrivalTime) {
    const departure = new Date(flightData.departureTime);
    const arrival = new Date(flightData.arrivalTime);

    if (departure >= arrival) {
      errors.arrivalTime = 'Arrival time must be after departure time';
    }
  }

  if (!flightData.distance) {
    errors.distance = 'Distance is required';
  } else if (flightData.distance <= 0) {
    errors.distance = 'Distance must be greater than 0';
  }

  if (!flightData.status) {
    errors.status = 'Status is required';
  }

  return errors;
};

export const validateSearchFilters = (filters) => {
  const errors = {};

  if (filters.departureDate && filters.arrivalDate) {
    const departure = new Date(filters.departureDate);
    const arrival = new Date(filters.arrivalDate);

    if (departure > arrival) {
      errors.arrivalDate = 'Arrival date must be after or equal to departure date';
    }
  }

  return errors;
}; 