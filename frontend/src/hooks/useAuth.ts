import { useState, useEffect } from 'react';
import type { User, UserDashboard, LoginCredentials } from '@/types';
import { authService } from '@/services/authService';

export function useAuth() {
  const [user, setUser] = useState<User | null>(null);
  const [dashboard, setDashboard] = useState<UserDashboard | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isInitialized, setIsInitialized] = useState(false);

  const loadDashboard = async () => {
    try {
      setLoading(true);
      setError(null); // Clear any previous errors
      console.log('Loading dashboard...');
      const dashboardData = await authService.getUserDashboard();
      console.log('Dashboard data loaded:', dashboardData);
      setDashboard(dashboardData);
    } catch (err) {
      console.error('Dashboard loading error:', err);
      setError(err instanceof Error ? err.message : 'Failed to load dashboard');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    const initAuth = async () => {
      // Check if user is already logged in
      const token = localStorage.getItem('token');
      const currentUser = authService.getCurrentUser();
      
      console.log('Auth check:', { token: !!token, currentUser: !!currentUser });
      console.log(isInitialized)
      
      if (token && currentUser) {
        console.log('User already authenticated, setting up auth state...');
        setUser(currentUser);
        setIsAuthenticated(true);
        // Load dashboard after setting authentication
        try {
          await loadDashboard();
        } catch (err) {
          console.error('Failed to load dashboard during init:', err);
        }
      } else {
        console.log('No existing authentication found');
        // Clear any stale auth data
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        setIsAuthenticated(false);
        setUser(null);
        setDashboard(null);
      }
      
      setIsInitialized(true);
    };

    initAuth();
  }, []);

  const login = async (credentials: LoginCredentials) => {
    try {
      setLoading(true);
      setError(null);
      
      const authResponse = await authService.login(credentials);
      
      // Store auth data
      //localStorage.setItem('token', authResponse.token);
      localStorage.setItem('user', JSON.stringify(authResponse.user));
      
      setUser(authResponse.user);
      setIsAuthenticated(true);
      
      // Load user dashboard
      await loadDashboard();
      
      return authResponse;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Login failed';
      setError(errorMessage);
      throw new Error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  const register = async (data: {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    birthdate?: string;
    phone?: string;
  }) => {
    try {
      setLoading(true);
      setError(null);
      
      const authResponse = await authService.register(data);
      
      // Store auth data
      //localStorage.setItem('token', authResponse.token);
      localStorage.setItem('user', JSON.stringify(authResponse.user));
      
      setUser(authResponse.user);
      setIsAuthenticated(true);
      
      // Load user dashboard
      await loadDashboard();
      
      return authResponse;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Registration failed';
      setError(errorMessage);
      throw new Error(errorMessage);
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    authService.logout();
    setUser(null);
    setDashboard(null);
    setIsAuthenticated(false);
    setError(null);
  };

  const refreshDashboard = async () => {
    setError(null); // Clear any existing errors
    await loadDashboard();
  };

  const clearError = () => {
    setError(null);
  };

  return {
    user,
    dashboard,
    loading,
    error,
    isAuthenticated,
    isInitialized,
    login,
    register,
    logout,
    refreshDashboard,
    clearError,
  };
}
