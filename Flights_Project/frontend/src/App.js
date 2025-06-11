import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { ThemeProvider, CssBaseline, Box } from '@mui/material';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { AuthProvider, useAuth } from './context/AuthContext';
import Navigation from './components/common/Navigation';
import Loading from './components/common/Loading';
import ErrorBoundary from './components/common/ErrorBoundary';
import Login from './components/auth/Login';
import Register from './components/auth/Register';
import Dashboard from './components/user/Dashboard';
import FlightList from './components/admin/FlightList';
import theme from './theme';

const PrivateRoute = ({ children, adminOnly = false }) => {
  const { user, loading } = useAuth();

  if (loading) {
    return <Loading />;
  }

  if (!user) {
    return <Navigate to="/login" />;
  }

  if (adminOnly && user.role !== 'admin') {
    return <Navigate to="/dashboard" />;
  }

  return children;
};

const AppContent = () => {
  const { user } = useAuth();

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      {user && <Navigation />}
      <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
        <ErrorBoundary>
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route
              path="/dashboard"
              element={
                <PrivateRoute>
                  <Dashboard />
                </PrivateRoute>
              }
            />
            <Route
              path="/admin/flights"
              element={
                <PrivateRoute adminOnly>
                  <FlightList />
                </PrivateRoute>
              }
            />
            <Route path="/" element={<Navigate to="/dashboard" />} />
          </Routes>
        </ErrorBoundary>
      </Box>
    </Box>
  );
};

const App = () => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <AuthProvider>
        <Router>
          <AppContent />
          <ToastContainer />
        </Router>
      </AuthProvider>
    </ThemeProvider>
  );
};

export default App; 