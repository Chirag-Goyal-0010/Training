import { toast } from 'react-toastify';
import { ERROR_MESSAGES, SUCCESS_MESSAGES } from './constants';

export const showSuccess = (message) => {
  toast.success(message, {
    position: 'top-right',
    autoClose: 3000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
  });
};

export const showError = (message) => {
  toast.error(message, {
    position: 'top-right',
    autoClose: 5000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
  });
};

export const showInfo = (message) => {
  toast.info(message, {
    position: 'top-right',
    autoClose: 3000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
  });
};

export const showWarning = (message) => {
  toast.warning(message, {
    position: 'top-right',
    autoClose: 4000,
    hideProgressBar: false,
    closeOnClick: true,
    pauseOnHover: true,
    draggable: true,
  });
};

export const handleApiError = (error) => {
  if (error.response) {
    switch (error.response.status) {
      case 401:
        showError(ERROR_MESSAGES.UNAUTHORIZED);
        break;
      case 403:
        showError(ERROR_MESSAGES.FORBIDDEN);
        break;
      case 404:
        showError(ERROR_MESSAGES.NOT_FOUND);
        break;
      case 422:
        showError(error.response.data.message || ERROR_MESSAGES.VALIDATION_ERROR);
        break;
      default:
        showError(ERROR_MESSAGES.SERVER_ERROR);
    }
  } else if (error.request) {
    showError(ERROR_MESSAGES.NETWORK_ERROR);
  } else {
    showError(error.message || ERROR_MESSAGES.SERVER_ERROR);
  }
};

export const showSuccessMessage = (type) => {
  const message = SUCCESS_MESSAGES[type];
  if (message) {
    showSuccess(message);
  }
};

export const showErrorMessage = (type) => {
  const message = ERROR_MESSAGES[type];
  if (message) {
    showError(message);
  }
}; 