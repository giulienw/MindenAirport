/**
 * Authentication Service
 * 
 * This service handles all authentication-related operations including user login,
 * registration, session management, and user profile operations. It manages JWT tokens
 * and provides secure communication with the backend authentication endpoints.
 * 
 * Key features:
 * - User login and registration
 * - JWT token management and validation
 * - User profile retrieval and management
 * - User dashboard data aggregation
 * - Session persistence and cleanup
 * 
 * @module AuthService
 */

import type {
  LoginCredentials,
  AuthResponse,
  User,
  Ticket,
  TicketWithFlight,
  UserDashboard,
} from "@/types";
import { flightService } from "./flightService";
import { API_BASE_URL } from "@/config";
import { getCookie } from "@/lib/utils";

/**
 * Authentication service object containing all auth-related methods
 */
export const authService = {
  /**
   * Authenticates a user with email and password
   * 
   * @param credentials - User login credentials
   * @returns Promise resolving to authentication response with user data and token
   * @throws Error if authentication fails
   */
  async login(credentials: LoginCredentials): Promise<AuthResponse> {
    try {
      const response = await fetch(`${API_BASE_URL}/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: credentials.email,
          password: credentials.password,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Login failed');
      }

      const result = await response.json();
      const authResponse: AuthResponse = {
        user: {
          id: result.data.user.id,
          firstName: result.data.user.firstName,
          lastName: result.data.user.lastName,
          email: result.data.user.email,
          phone: result.data.user.phone,
          birthdate: result.data.user.birthdate,
          active: result.data.user.active,
          role: result.data.user.role,
        },
        token: result.data.token,
      };
      
      localStorage.setItem("user", JSON.stringify(authResponse.user));
      return authResponse;
    } catch (error) {
      console.error('Login error:', error);
      throw error;
    }
  },

  async register(data: {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    birthdate?: string;
    phone?: string;
  }): Promise<AuthResponse> {
    try {
      const response = await fetch(`${API_BASE_URL}/auth/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          firstName: data.firstName,
          lastName: data.lastName,
          email: data.email,
          password: data.password,
          birthdate: data.birthdate || new Date().toISOString(),
          phone: data.phone || '',
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Registration failed');
      }

      const result = await response.json();
      const authResponse: AuthResponse = {
        user: {
          id: result.data.user.id,
          firstName: result.data.user.firstName,
          lastName: result.data.user.lastName,
          email: result.data.user.email,
          phone: result.data.user.phone,
          birthdate: result.data.user.birthdate,
          active: result.data.user.active,
          role: result.data.user.role,
        },
        token: result.data.token,
      };
      
      localStorage.setItem("user", JSON.stringify(authResponse.user));
      return authResponse;
    } catch (error) {
      console.error('Registration error:', error);
      throw error;
    }
  },

  async getUserTickets(): Promise<TicketWithFlight[]> {
    try {
      const response = await fetch(`${API_BASE_URL}/ticket/my`, {
        headers: {
          Authorization: `Bearer ${authService.getToken()}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to fetch tickets');
      }

      const result = await response.json();
      const tickets = result.data || [];

      // Enrich tickets with flight information
      const enrichedFlights = await flightService.getEnrichedFlights();
      const flightMap = new Map(enrichedFlights.map((f) => [f.id, f]));

      return tickets.map((ticket: Ticket) => ({
        ...ticket,
        flightInfo: flightMap.get(ticket.flight),
      }));
    } catch (error) {
      console.warn('Failed to fetch tickets:', error);
      return [];
    }
  },

  async getUserProfile(): Promise<User> {
    try {
      const response = await fetch(`${API_BASE_URL}/auth/profile`, {
        headers: {
          Authorization: `Bearer ${authService.getToken()}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to fetch user profile');
      }

      const result = await response.json();
      return {
        id: result.data.id,
        firstName: result.data.firstName,
        lastName: result.data.lastName,
        email: result.data.email,
        phone: result.data.phone,
        birthdate: result.data.birthdate,
        active: result.data.active,
        role: result.data.role,
      };
    } catch (error) {
      console.error('Failed to fetch user profile:', error);
      throw error;
    }
  },

  async getUserDashboard(): Promise<UserDashboard> {
    try {
      // Get user profile and tickets in parallel
      const [userProfile, tickets] = await Promise.all([
        this.getUserProfile(),
        this.getUserTickets(),
      ]);

      const now = new Date();
      const upcomingFlights = tickets.filter(ticket => {
        const departureTime = new Date(ticket.departureTime || '');
        return departureTime > now;
      });

      const pastFlights = tickets.filter(ticket => {
        const departureTime = new Date(ticket.departureTime || '');
        return departureTime <= now;
      });

      return {
        user: {
          id: userProfile.id,
          firstName: userProfile.firstName,
          lastName: userProfile.lastName,
          email: userProfile.email,
        },
        tickets,
        upcomingFlights,
        pastFlights,
      };
    } catch (error) {
      console.error('Failed to fetch user dashboard:', error);
      throw error;
    }
  },

  async logout() {
    try {
      const response = await fetch(`${API_BASE_URL}/auth/logout`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Logout failed');
      }
    } catch (error) {
      console.error('Logout error:', error);
      throw error;
    }
  },

  getCurrentUser(): User | null {
    const userString = localStorage.getItem("user");
    return userString ? JSON.parse(userString) : null;
  },

  isAuthenticated(): boolean {
    return !!getCookie("token");
  },

  getToken(): string | null {
    console.log(getCookie("token"))
    return getCookie("token") || null;
  }
};
