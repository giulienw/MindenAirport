import type { BaggageItem } from '@/types';
import { API_BASE_URL } from '@/config';

export const baggageService = {
  async getUserBaggage(): Promise<BaggageItem[]> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/baggage/my`, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch baggage');
      }

      const result = await response.json();
      const items: BaggageItem[] = result.data || [];
      
      return items;
    } catch (error) {
      console.error('Failed to fetch baggage:', error);
      throw error;
    }
  },

  /*async getBaggageTracking(baggageId: string): Promise<any> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/baggage/${baggageId}/tracking`, {
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to fetch tracking information');
      }

      const result = await response.json();
      return result.data;
    } catch (error) {
      console.error('Failed to fetch baggage tracking:', error);
      throw error;
    }
  },*/

  async reportLostBaggage(baggageId: string, description: string): Promise<void> {
    const token = localStorage.getItem('token');
    if (!token) {
      throw new Error('No authentication token found');
    }

    try {
      const response = await fetch(`${API_BASE_URL}/baggage/${baggageId}/report-lost`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ description }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to report lost baggage');
      }
    } catch (error) {
      console.error('Failed to report lost baggage:', error);
      throw error;
    }
  },
};
