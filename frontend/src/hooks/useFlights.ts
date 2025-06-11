import { useState, useEffect } from 'react';
import { type FlightDisplayInfo } from '@/types';
import { flightService } from '@/services/flightService';

export function useFlights() {
  const [flights, setFlights] = useState<FlightDisplayInfo[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

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

  useEffect(() => {
    fetchFlights();
  }, []);

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
