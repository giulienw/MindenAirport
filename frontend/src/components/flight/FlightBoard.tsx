import { useState, useMemo } from 'react';
import type { FlightDisplayInfo } from '@/types';
import { FlightCard } from './FlightCard';

interface FlightBoardProps {
  flights: FlightDisplayInfo[];
  loading?: boolean;
}

export function FlightBoard({ flights, loading }: FlightBoardProps) {
  const [filter, setFilter] = useState<'all' | 'departures' | 'arrivals'>('all');
  const [statusFilter, setStatusFilter] = useState<string>('all');

  const filteredFlights = useMemo(() => {
    let filtered = flights;

    // Filter by flight type
    if (filter === 'departures') {
      // Assuming departures are flights starting from this airport
      // You might need to adjust this logic based on your specific airport
      filtered = filtered.filter(flight => {
        const now = new Date();
        const departure = new Date(flight.scheduledDeparture);
        return departure >= now;
      });
    } else if (filter === 'arrivals') {
      // Assuming arrivals are flights coming to this airport
      filtered = filtered.filter(flight => {
        const now = new Date();
        const arrival = new Date(flight.scheduledArrival);
        return arrival >= now;
      });
    }

    // Filter by status
    if (statusFilter !== 'all') {
      filtered = filtered.filter(flight => 
        flight.statusInfo?.name?.toLowerCase() === statusFilter.toLowerCase()
      );
    }

    // Sort by departure time
    return filtered.sort((a, b) => 
      new Date(a.scheduledDeparture).getTime() - new Date(b.scheduledDeparture).getTime()
    );
  }, [flights, filter, statusFilter]);

  const statusOptions = useMemo(() => {
    const statuses = new Set(flights.map(f => f.statusInfo?.name).filter(Boolean));
    return Array.from(statuses);
  }, [flights]);

  if (loading) {
    return (
      <div className="space-y-4">
        {[...Array(6)].map((_, i) => (
          <div key={i} className="bg-white rounded-lg shadow-md p-6 animate-pulse">
            <div className="flex justify-between items-start mb-4">
              <div className="h-6 bg-gray-200 rounded w-1/3"></div>
              <div className="h-6 bg-gray-200 rounded w-20"></div>
            </div>
            <div className="grid grid-cols-2 gap-4 mb-4">
              <div>
                <div className="h-4 bg-gray-200 rounded w-16 mb-2"></div>
                <div className="h-6 bg-gray-200 rounded w-20 mb-1"></div>
                <div className="h-4 bg-gray-200 rounded w-16"></div>
              </div>
              <div>
                <div className="h-4 bg-gray-200 rounded w-16 mb-2"></div>
                <div className="h-6 bg-gray-200 rounded w-20 mb-1"></div>
                <div className="h-4 bg-gray-200 rounded w-16"></div>
              </div>
            </div>
          </div>
        ))}
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Filters */}
      <div className="bg-white rounded-lg shadow-md p-4">
        <div className="flex flex-wrap gap-4 items-center">
          <div className="flex space-x-2">
            <button
              onClick={() => setFilter('all')}
              className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                filter === 'all'
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              }`}
            >
              All Flights
            </button>
            <button
              onClick={() => setFilter('departures')}
              className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                filter === 'departures'
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              }`}
            >
              Departures
            </button>
            <button
              onClick={() => setFilter('arrivals')}
              className={`px-4 py-2 rounded-lg font-medium transition-colors ${
                filter === 'arrivals'
                  ? 'bg-blue-600 text-white'
                  : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
              }`}
            >
              Arrivals
            </button>
          </div>

          <div className="flex items-center space-x-2">
            <label htmlFor="status-filter" className="text-sm font-medium text-gray-700">
              Status:
            </label>
            <select
              id="status-filter"
              value={statusFilter}
              onChange={(e) => setStatusFilter(e.target.value)}
              className="border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="all">All Statuses</option>
              {statusOptions.map(status => (
                <option key={status} value={status}>
                  {status}
                </option>
              ))}
            </select>
          </div>
        </div>
      </div>

      {/* Flight Count */}
      <div className="text-sm text-gray-600">
        Showing {filteredFlights.length} of {flights.length} flights
      </div>

      {/* Flights List */}
      {filteredFlights.length === 0 ? (
        <div className="bg-white rounded-lg shadow-md p-8 text-center">
          <div className="text-gray-500 text-lg mb-2">No flights found</div>
          <div className="text-gray-400">Try adjusting your filters</div>
        </div>
      ) : (
        <div className="space-y-4">
          {filteredFlights.map(flight => (
            <FlightCard key={flight.id} flight={flight} />
          ))}
        </div>
      )}
    </div>
  );
}
