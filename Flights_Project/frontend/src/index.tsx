import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { AuthProvider } from './context/AuthContext.tsx';
import { FlightsProvider } from './context/FlightsContext.tsx';
import { BookingsProvider } from './context/BookingsContext.tsx';
import { ThemeProvider } from '@mui/material/styles';
import theme from './theme';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <AuthProvider>
        <FlightsProvider>
          <BookingsProvider>
            <App />
          </BookingsProvider>
        </FlightsProvider>
      </AuthProvider>
    </ThemeProvider>
  </React.StrictMode>
); 