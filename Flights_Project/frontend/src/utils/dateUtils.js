import { format, parseISO } from 'date-fns';

export const formatDateTime = (dateString) => {
  if (!dateString) return '';
  try {
    const date = parseISO(dateString);
    return format(date, 'MMM dd, yyyy HH:mm');
  } catch (error) {
    console.error('Error formatting date:', error);
    return dateString;
  }
};

export const formatDate = (dateString) => {
  if (!dateString) return '';
  try {
    const date = parseISO(dateString);
    return format(date, 'MMM dd, yyyy');
  } catch (error) {
    console.error('Error formatting date:', error);
    return dateString;
  }
};

export const formatTime = (dateString) => {
  if (!dateString) return '';
  try {
    const date = parseISO(dateString);
    return format(date, 'HH:mm');
  } catch (error) {
    console.error('Error formatting time:', error);
    return dateString;
  }
};

export const calculateDuration = (departureTime, arrivalTime) => {
  if (!departureTime || !arrivalTime) return '';
  try {
    const departure = parseISO(departureTime);
    const arrival = parseISO(arrivalTime);
    const durationMs = arrival - departure;
    const hours = Math.floor(durationMs / (1000 * 60 * 60));
    const minutes = Math.floor((durationMs % (1000 * 60 * 60)) / (1000 * 60));
    return `${hours}h ${minutes}m`;
  } catch (error) {
    console.error('Error calculating duration:', error);
    return '';
  }
};

export const isDateInRange = (date, startDate, endDate) => {
  if (!date || !startDate || !endDate) return true;
  try {
    const checkDate = parseISO(date);
    const start = parseISO(startDate);
    const end = parseISO(endDate);
    return checkDate >= start && checkDate <= end;
  } catch (error) {
    console.error('Error checking date range:', error);
    return true;
  }
}; 