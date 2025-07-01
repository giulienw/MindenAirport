import React, { useState } from 'react';
import { Clock, Users, Plane, Package, AlertTriangle, DollarSign, Activity, TrendingUp } from 'lucide-react';
import { AdminStatsCard } from './AdminStatsCard';
import { UserManagement } from './UserManagement';
import { FlightManagement } from './FlightManagement';
import { useAdmin } from '@/hooks/useAdmin';
import { authService } from '@/services';

type TabType = 'overview' | 'users' | 'flights' | 'baggage';

export const AdminDashboard: React.FC = () => {
  const [activeTab, setActiveTab] = useState<TabType>('overview');
  const { dashboard, loading, error, refreshDashboard } = useAdmin();

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="w-16 h-16 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-600">Loading admin dashboard...</p>
        </div>
      </div>
    );
  }

  if (error || !dashboard) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <AlertTriangle className="w-16 h-16 text-red-600 mx-auto mb-4" />
          <h2 className="text-xl font-semibold text-gray-900 mb-2">Failed to load dashboard</h2>
          <p className="text-gray-600 mb-4">{error || 'An unexpected error occurred'}</p>
          <button
            onClick={refreshDashboard}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  const tabs = [
    { id: 'overview' as const, label: 'Overview', icon: TrendingUp },
    { id: 'users' as const, label: 'Users', icon: Users },
    { id: 'flights' as const, label: 'Flights', icon: Plane },
    { id: 'baggage' as const, label: 'Baggage', icon: Package }
  ];

  const statsConfig = [
    {
      title: 'Total Flights',
      value: dashboard.stats.totalFlights,
      icon: Plane,
      color: 'blue' as const,
      change: '+5.2%'
    },
    {
      title: 'Active Flights',
      value: dashboard.stats.activeFlights,
      icon: Activity,
      color: 'green' as const,
      change: '+12%'
    },
    {
      title: 'Total Passengers',
      value: dashboard.stats.totalPassengers.toLocaleString(),
      icon: Users,
      color: 'purple' as const,
      change: '+8.1%'
    },
    {
      title: 'Revenue',
      value: `$${(dashboard.stats.revenue / 1000000).toFixed(1)}M`,
      icon: DollarSign,
      color: 'emerald' as const,
      change: '+15.3%'
    },
    {
      title: 'Delayed Flights',
      value: dashboard.stats.delayedFlights,
      icon: Clock,
      color: 'orange' as const,
      change: '-2.1%'
    },
    {
      title: 'Lost Baggage',
      value: dashboard.stats.lostBaggage,
      icon: Package,
      color: 'red' as const,
      change: '-5.5%'
    },
  ];

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <div className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center py-6">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">Admin Dashboard</h1>
              <p className="text-gray-600">Welcome back, {authService.getCurrentUser()?.firstName}</p>
            </div>
            <div className="flex items-center space-x-3">
              <button
                onClick={refreshDashboard}
                className="px-4 py-2 text-gray-600 hover:text-gray-900 transition-colors"
                title="Refresh dashboard"
              >
                <Activity className="w-5 h-5" />
              </button>
              <div className="text-sm text-gray-500">
                Last updated: {new Date().toLocaleTimeString()}
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Navigation Tabs */}
      <div className="bg-white border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <nav className="flex space-x-8">
            {tabs.map((tab) => {
              const Icon = tab.icon;
              return (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`flex items-center space-x-2 py-4 px-1 border-b-2 font-medium text-sm transition-colors ${
                    activeTab === tab.id
                      ? 'border-blue-500 text-blue-600'
                      : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                  }`}
                >
                  <Icon className="w-5 h-5" />
                  <span>{tab.label}</span>
                </button>
              );
            })}
          </nav>
        </div>
      </div>

      {/* Content */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {activeTab === 'overview' && (
          <div className="space-y-8">
            {/* Stats Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {statsConfig.map((stat, index) => (
                <AdminStatsCard
                  key={index}
                  title={stat.title}
                  value={stat.value}
                  icon={stat.icon}
                  color={stat.color}
                  change={stat.change}
                />
              ))}
            </div>
          </div>
        )}

        {activeTab === 'users' && <UserManagement />}
        {activeTab === 'flights' && <FlightManagement />}
        {activeTab === 'baggage' && (
          <div className="bg-white rounded-lg shadow p-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Baggage Management</h2>
            <p className="text-gray-600">Baggage management interface coming soon...</p>
          </div>
        )}
      </div>
    </div>
  );
};
