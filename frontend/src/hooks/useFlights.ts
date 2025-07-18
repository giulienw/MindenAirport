/**
 * useFlights Hook
 * 
 * Custom React hook for managing flight data state and operations.
 * Provides a centralized way to fetch, manage, and refetch flight information
 * with built-in loading states and error handling.
 * 
 * Features:
 * - Automatic flight data fetching on mount
 * - Loading state management
 * - Error handling and reporting
 * - Manual refetch capability
 * - Enriched flight data with related entities
 * 
 * @returns Object containing flights array, loading state, error state, and refetch function
 */

import { useState, useEffect } from 'react';
import { type FlightDisplayInfo } from '@/types';
import { flightService } from '@/services/flightService';

export function useFlights() {
  // Flight data state
  const [flights, setFlights] = useState<FlightDisplayInfo[]>([]);
  // Loading state for UI feedback
  const [loading, setLoading] = useState(true);
  // Error state for error handling and display
  const [error, setError] = useState<string | null>(null);

  /**
   * Fetches flight data from the service layer
   * Handles loading states and error management
   */
  const fetchFlights = async () => {
    try {
      setLoading(true);
      setError(null);
      const flightData = await flightService.getEnrichedFlights();
      setFlights(flightData);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch flights');
      console.error('Error fetching flights:', err);
    } finally {
      setLoading(false);
    }
  };

  // Fetch flights on component mount
  useEffect(() => {
    fetchFlights();
  }, []);

  /**
   * Manually refetch flight data
   * Useful for refresh buttons or data synchronization
   */
  const refetch = () => {
    fetchFlights();
  };

  return {
    flights,
    loading,
    error,
    refetch,
  };
}
