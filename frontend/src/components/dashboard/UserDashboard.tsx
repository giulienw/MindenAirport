import { useState, useEffect } from 'react';
import type { UserDashboard } from '@/types';
import { TicketCard } from '@/components/ticket';
import { BaggageView } from '@/components/baggage';
import { useBaggage } from '@/hooks';
import { User, Calendar, Plane, Clock, LogOut, Package } from 'lucide-react';

interface UserDashboardComponentProps {
  dashboard: UserDashboard;
  onLogout: () => void;
}

export function UserDashboardComponent({ dashboard, onLogout }: UserDashboardComponentProps) {
  const [activeTab, setActiveTab] = useState<'upcoming' | 'past' | 'all' | 'baggage'>('upcoming');
  const [currentTime, setCurrentTime] = useState(new Date());
  const { baggage, loading: baggageLoading, refetch: refetchBaggage } = useBaggage();

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  const getDisplayTickets = () => {
    switch (activeTab) {
      case 'upcoming':
        return dashboard.upcomingFlights;
      case 'past':
        return dashboard.pastFlights;
      case 'all':
        return dashboard.tickets;
      case 'baggage':
        return []; // Baggage view doesn't use tickets
      default:
        return dashboard.upcomingFlights;
    }
  };

  const stats = {
    total: dashboard.tickets.length,
    upcoming: dashboard.upcomingFlights.length,
    past: dashboard.pastFlights.length,
    confirmed: dashboard.tickets.filter(t => t.status === 'CONFIRMED').length,
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <div className="bg-white shadow-sm border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-3">
              <div className="h-10 w-10 bg-blue-600 rounded-full flex items-center justify-center">
                <User className="h-6 w-6 text-white" />
              </div>
              <div>
                <h1 className="text-xl font-semibold text-gray-900">
                  Welcome, {dashboard.user.firstName} {dashboard.user.lastName}
                </h1>
                <p className="text-sm text-gray-600">{dashboard.user.email}</p>
              </div>
            </div>
            <div className="flex items-center space-x-4">
              <div className="text-right">
                <div className="text-sm font-medium text-gray-900">
                  {currentTime.toLocaleTimeString('en-US', {
                    hour: '2-digit',
                    minute: '2-digit',
                    hour12: false,
                  })}
                </div>
                <div className="text-xs text-gray-500">
                  {currentTime.toLocaleDateString('en-US', {
                    month: 'short',
                    day: 'numeric',
                  })}
                </div>
              </div>
              <button
                onClick={onLogout}
                className="flex items-center space-x-2 text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium transition-colors"
              >
                <LogOut className="h-4 w-4" />
                <span>Logout</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Stats */}
        <div className="grid grid-cols-2 md:grid-cols-5 gap-4 mb-8">
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Plane className="h-8 w-8 text-blue-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.total}</div>
            <div className="text-sm text-gray-600">Total Tickets</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Calendar className="h-8 w-8 text-green-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.upcoming}</div>
            <div className="text-sm text-gray-600">Upcoming</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Clock className="h-8 w-8 text-purple-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.past}</div>
            <div className="text-sm text-gray-600">Past Flights</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <User className="h-8 w-8 text-yellow-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{stats.confirmed}</div>
            <div className="text-sm text-gray-600">Confirmed</div>
          </div>
          <div className="bg-white rounded-lg shadow p-6 text-center">
            <div className="flex justify-center mb-2">
              <Package className="h-8 w-8 text-orange-600" />
            </div>
            <div className="text-2xl font-bold text-gray-900">{baggage?.length || 0}</div>
            <div className="text-sm text-gray-600">Baggage Items</div>
          </div>
        </div>

        {/* Tabs */}
        <div className="bg-white rounded-lg shadow mb-6">
          <div className="border-b border-gray-200">
            <nav className="flex space-x-8 px-6">
              <button
                onClick={() => setActiveTab('upcoming')}
                className={`py-4 px-1 border-b-2 font-medium text-sm transition-colors ${
                  activeTab === 'upcoming'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                Upcoming Flights ({stats.upcoming})
              </button>
              <button
                onClick={() => setActiveTab('past')}
                className={`py-4 px-1 border-b-2 font-medium text-sm transition-colors ${
                  activeTab === 'past'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                Past Flights ({stats.past})
              </button>
              <button
                onClick={() => setActiveTab('all')}
                className={`py-4 px-1 border-b-2 font-medium text-sm transition-colors ${
                  activeTab === 'all'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                All Tickets ({stats.total})
              </button>
              <button
                onClick={() => setActiveTab('baggage')}
                className={`py-4 px-1 border-b-2 font-medium text-sm transition-colors ${
                  activeTab === 'baggage'
                    ? 'border-blue-500 text-blue-600'
                    : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
                }`}
              >
                My Baggage ({baggage.length})
              </button>
            </nav>
          </div>
        </div>

        {/* Content */}
        <div className="space-y-6">
          {activeTab === 'baggage' ? (
            baggage ? (
              <BaggageView 
                baggage={baggage} 
                loading={baggageLoading} 
                onRefresh={refetchBaggage}
              />
            ) : (
              <div className="bg-white rounded-lg shadow p-8 text-center">
                <Package className="h-16 w-16 text-gray-300 mx-auto mb-4" />
                <h3 className="text-lg font-medium text-gray-900 mb-2">No baggage data</h3>
                <p className="text-gray-500">Loading baggage information...</p>
              </div>
            )
          ) : (
            <>
              {getDisplayTickets().length === 0 ? (
                <div className="bg-white rounded-lg shadow p-8 text-center">
                  <Plane className="h-16 w-16 text-gray-300 mx-auto mb-4" />
                  <h3 className="text-lg font-medium text-gray-900 mb-2">No tickets found</h3>
                  <p className="text-gray-500">
                    {activeTab === 'upcoming' 
                      ? "You don't have any upcoming flights."
                      : activeTab === 'past'
                      ? "You don't have any past flights."
                      : "You don't have any tickets yet."
                    }
                  </p>
                </div>
              ) : (
                getDisplayTickets().map(ticket => (
                  <TicketCard key={ticket.id} ticket={ticket} />
                ))
              )}
            </>
          )}
        </div>
      </div>
    </div>
  );
}
