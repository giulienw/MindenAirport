import { useEffect } from 'react';
import { useNavigate } from 'react-router';
import { useAuth } from '@/hooks';
import { UserDashboardComponent } from '@/components/dashboard';

export default function Dashboard() {
  const { dashboard, isAuthenticated, logout, loading, refreshDashboard, error, clearError, isInitialized } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    // Only check authentication after initialization is complete
    if (isInitialized) {
      if (!isAuthenticated) {
        console.log('User not authenticated, redirecting to login');
        navigate('/login');
      } else if (isAuthenticated && !dashboard && !loading && !error) {
        console.log('Dashboard not loaded, triggering refresh...');
        refreshDashboard();
      }
    }
  }, [isInitialized, isAuthenticated, dashboard, loading, error, navigate, refreshDashboard]);

  // Debug logging
  console.log('Dashboard component state:', {
    isInitialized,
    isAuthenticated,
    hasDashboard: !!dashboard,
    loading,
    error
  });

  // Show loading while auth is initializing
  if (!isInitialized) {
    return (
      <div className="min-h-[80vh] flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading...</p>
        </div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return null; // This will redirect to login
  }

  if (error) {
    return (
      <div className="min-h-[80vh] flex items-center justify-center">
        <div className="text-center">
          <div className="text-red-600 mb-4">
            <svg className="h-12 w-12 mx-auto" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 className="text-lg font-medium text-gray-900 mb-2">Error loading dashboard</h3>
          <p className="text-gray-600 mb-4">{error}</p>
          <button
            onClick={() => {
              clearError();
              refreshDashboard();
            }}
            className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  if (loading || !dashboard) {
    return (
      <div className="min-h-[80vh] flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">Loading your dashboard...</p>
        </div>
      </div>
    );
  }

  return <UserDashboardComponent dashboard={dashboard} onLogout={logout} />;
}
