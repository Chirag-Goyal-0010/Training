import React, { createContext, useContext, useState, ReactNode } from 'react';
import { authAPI } from '../api';

interface AuthContextType {
  isAuthenticated: boolean;
  user: User | null;
  loadingAuth: boolean;
  authError: string | null;
  login: (username: string, password: string) => Promise<{ success: boolean }>;
  register: (username: string, email: string, password: string, role: string) => Promise<{ success: boolean }>;
  logout: () => void;
}

interface User {
  id: number;
  username: string;
  email: string;
  role: string;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(false);
  const [user, setUser] = useState<User | null>(null);
  const [loadingAuth, setLoadingAuth] = useState<boolean>(false);
  const [authError, setAuthError] = useState<string | null>(null);

  const login = async (username: string, password: string) => {
    try {
      setLoadingAuth(true);
      setAuthError(null);
      const response = await authAPI.login(username, password);
      setIsAuthenticated(true);
      setUser(response.user);
      localStorage.setItem('token', response.token);
      return { success: true };
    } catch (error) {
      setAuthError(error instanceof Error ? error.message : 'Login failed');
      return { success: false };
    } finally {
      setLoadingAuth(false);
    }
  };

  const register = async (username: string, email: string, password: string, role: string) => {
    try {
      setLoadingAuth(true);
      setAuthError(null);
      await authAPI.register(username, email, password, role);
      return { success: true };
    } catch (error) {
      setAuthError(error instanceof Error ? error.message : 'Registration failed');
      return { success: false };
    } finally {
      setLoadingAuth(false);
    }
  };

  const logout = () => {
    setIsAuthenticated(false);
    setUser(null);
    localStorage.removeItem('token');
  };

  const value = {
    isAuthenticated,
    user,
    loadingAuth,
    authError,
    login,
    register,
    logout,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export default AuthContext; 