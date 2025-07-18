/**
 * useBaggage Hook
 * 
 * Custom React hook for managing user baggage data and operations.
 * Provides state management for baggage tracking, lost baggage reporting,
 * and data refresh capabilities with comprehensive error handling.
 * 
 * Features:
 * - Automatic baggage data fetching on mount
 * - Loading and error state management
 * - Lost baggage reporting functionality
 * - Manual data refresh capability
 * - User-specific baggage filtering
 * 
 * @returns Object containing baggage array, states, and action functions
 */

import { useState, useEffect } from 'react';
import type { BaggageItem } from '@/types';
import { baggageService } from '@/services';

export function useBaggage() {
  // Baggage items state
  const [baggage, setBaggage] = useState<BaggageItem[]>([]);
  // Loading state for UI feedback
  const [loading, setLoading] = useState(true);
  // Error state for error handling and display
  const [error, setError] = useState<string | null>(null);

  /**
   * Fetches user's baggage data from the service layer
   * Handles loading states and error management
   */
  const fetchBaggage = async () => {
    try {
      setLoading(true);
      setError(null);
      console.log('Loading baggage...');
      
      // For now, use mock data. In production, use the real API
      //const baggageData = baggageService.getMockBaggage();
      const baggageData = await baggageService.getUserBaggage();
      
      console.log('Baggage data loaded:', baggageData);
      setBaggage(baggageData);
    } catch (err) {
      console.error('Baggage loading error:', err);
      setError(err instanceof Error ? err.message : 'Failed to load baggage');
    } finally {
      setLoading(false);
    }
  };

  /**
   * Manually refetch baggage data
   * Useful for refresh operations and data synchronization
   */
  const refetchBaggage = () => {
    fetchBaggage();
  };

  /**
   * Report a baggage item as lost
   * 
   * @param baggageId - ID of the baggage item to report as lost
   * @param description - Description of the loss circumstances
   * @throws Error if reporting fails
   */
  const reportLostBaggage = async (baggageId: string, description: string) => {
    try {
      setError(null);
      await baggageService.reportLostBaggage(baggageId, description);
      // Refresh baggage data after reporting
      await fetchBaggage();
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to report lost baggage');
      throw err;
    }
    };

  // Fetch baggage data on component mount
  useEffect(() => {
    fetchBaggage();
  }, []);

  return {
    baggage,
    loading,
    error,
    refetch: refetchBaggage,
    reportLostBaggage,
  };
}