import { useState, useEffect } from 'react';
import type { BaggageItem } from '@/types';
import { baggageService } from '@/services';

export function useBaggage() {
  const [baggage, setBaggage] = useState<BaggageItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

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

  const refetchBaggage = () => {
    fetchBaggage();
  };

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
