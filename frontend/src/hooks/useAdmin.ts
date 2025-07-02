import { useState, useEffect } from 'react';
import type { AdminDashboard } from '@/types';
import { adminService } from '@/services';

interface UseAdminReturn {
  dashboard: AdminDashboard | null;
  loading: boolean;
  error: string | null;
  refreshDashboard: () => Promise<void>;
}

export const useAdmin = (): UseAdminReturn => {
  const [dashboard, setDashboard] = useState<AdminDashboard | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchDashboard = async () => {
    try {
      setLoading(true);
      setError(null);
      
      const data = await adminService.getAdminDashboard();
      setDashboard(data);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch admin dashboard';
      setError(errorMessage);
      console.error('Admin dashboard error:', err);
    } finally {
      setLoading(false);
    }
  };

  const refreshDashboard = async () => {
    await fetchDashboard();
  };

  useEffect(() => {
    fetchDashboard();
  }, []);

  return {
    dashboard,
    loading,
    error,
    refreshDashboard,
  };
};
