import React, { createContext, useState, useContext, useEffect } from 'react';
import { authAPI } from '../utils/api';
import { storage } from '../utils/storage';
import { showSuccessMessage, showErrorMessage } from '../utils/notifications';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const verifyToken = async () => {
      const token = storage.getToken();
      if (token) {
        try {
          const response = await authAPI.verify();
          setUser(response.data);
          storage.setUser(response.data);
        } catch (error) {
          storage.clear();
        }
      }
      setLoading(false);
    };

    verifyToken();
  }, []);

  const login = async (username, password) => {
    try {
      const response = await authAPI.login(username, password);
      const { token, user } = response.data;
      storage.setToken(token);
      storage.setUser(user);
      setUser(user);
      showSuccessMessage('LOGIN');
      return { success: true };
    } catch (error) {
      showErrorMessage('LOGIN');
      return { success: false, error: error.message };
    }
  };

  const register = async (username, password, role) => {
    try {
      await authAPI.register(username, password, role);
      showSuccessMessage('REGISTER');
      return { success: true };
    } catch (error) {
      showErrorMessage('REGISTER');
      return { success: false, error: error.message };
    }
  };

  const logout = () => {
    storage.clear();
    setUser(null);
    showSuccessMessage('LOGOUT');
  };

  const value = {
    user,
    loading,
    login,
    register,
    logout,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export default AuthContext; 