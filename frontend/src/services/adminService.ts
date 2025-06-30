import type { 
  AdminDashboard, 
  UserManagement, 
  FlightManagement, 
  BaggageManagement} from '@/types';
import { API_BASE_URL } from '@/config';

export const adminService = {
  async getAdminDashboard(): Promise<AdminDashboard> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/dashboard`, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch admin dashboard');
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error('Failed to fetch admin dashboard:', error);
      throw error;
    }
  },

  async getUsers(page = 1, limit = 20, search?: string): Promise<{ users: UserManagement[], total: number }> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const params = new URLSearchParams({
        page: page.toString(),
        limit: limit.toString(),
        ...(search && { search })
      });

      const response = await fetch(`${API_BASE_URL}/admin/users?${params}`, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch users');
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error('Failed to fetch users:', error);
      throw error;
    }
  },

  async updateUserStatus(userId: string, active: boolean): Promise<void> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/users/${userId}/status`, {
        method: 'PATCH',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ active }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to update user status');
      }
    } catch (error) {
      console.error('Failed to update user status:', error);
      throw error;
    }
  },

  async getFlightManagement(): Promise<FlightManagement[]> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/flights`, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch flight management data');
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error('Failed to fetch flight management data:', error);
      throw error;
    }
  },

  async updateFlightStatus(flightId: string, status: string, gate?: string): Promise<void> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/flights/${flightId}`, {
        method: 'PATCH',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ status, gate }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to update flight status');
      }
    } catch (error) {
      console.error('Failed to update flight status:', error);
      throw error;
    }
  },

  async getBaggageManagement(): Promise<BaggageManagement[]> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/baggage`, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch baggage management data');
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error('Failed to fetch baggage management data:', error);
      throw error;
    }
  },

  async resolveAlert(alertId: string): Promise<void> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/admin/alerts/${alertId}/resolve`, {
        method: 'PATCH',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to resolve alert');
      }
    } catch (error) {
      console.error('Failed to resolve alert:', error);
      throw error;
    }
  },

}