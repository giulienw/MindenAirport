import React, { useState, useEffect } from 'react';
import { Plane, Clock, MapPin, Users, DollarSign, Package, AlertTriangle, CheckCircle } from 'lucide-react';
import type { FlightDisplayInfo, FlightManagement as FlightManagementType } from '@/types';
import { adminService, flightService } from '@/services';
import { getFlightType, getFlightStatusColor as getStatusColor } from '@/lib/utils';

export const FlightManagement: React.FC = () => {
  const [flights, setFlights] = useState<FlightManagementType[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedFlight, setSelectedFlight] = useState<FlightManagementType | null>(null);
  const [showStatusModal, setShowStatusModal] = useState(false);

  useEffect(() => {
    fetchFlights();
  }, []);

  const fetchFlights = async () => {
    try {
      setLoading(true);
      setError(null);
      const result = await flightService.getEnrichedFlights();
      setFlights(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch flights');
    } finally {
      setLoading(false);
    }
  };

  const handleUpdateStatus = async (flightId: string, newStatus: string, gate?: string) => {
    try {
      await adminService.updateFlightStatus(flightId, newStatus, gate);
      await fetchFlights();
      setShowStatusModal(false);
      setSelectedFlight(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update flight status');
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status.toLowerCase()) {
      case 'on time':
        return CheckCircle;
      case 'delayed':
        return Clock;
      case 'cancelled':
        return AlertTriangle;
      case 'boarding':
        return Users;
      case 'departed':
        return Plane;
      default:
        return Clock;
    }
  };

  const formatTime = (dateString: string) => {
    return new Date(dateString).toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
    });
  };

  const calculateOccupancy = (passengerCount: number, capacity: number) => {
    return Math.round((passengerCount / capacity) * 100);
  };

  if (loading) {
    return (
      <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
        <div className="text-center">
          <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-600">Loading flights...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
        <div className="text-center">
          <AlertTriangle className="w-12 h-12 text-red-500 mx-auto mb-4" />
          <p className="text-gray-600 mb-4">{error}</p>
          <button
            onClick={fetchFlights}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            Try Again
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        <div className="flex justify-between items-center">
          <div>
            <h2 className="text-xl font-semibold text-gray-900">Flight Management</h2>
            <p className="text-gray-600">{flights.length} flights</p>
          </div>
          <button
            onClick={fetchFlights}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors"
          >
            Refresh
          </button>
        </div>
      </div>

      {/* Flights Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {flights.map((flight) => {
          const StatusIcon = getStatusIcon(flight.statusInfo?.name || '');
          const occupancy = calculateOccupancy(flight.passengerCount, flight.capacity);
          
          return (
            <div key={flight.id} className="bg-white rounded-lg shadow-sm border border-gray-200 hover:shadow-md transition-shadow">
              <div className="p-6">
                {/* Flight Header */}
                <div className="flex justify-between items-start mb-4">
                  <div>
                    <div className="flex items-center space-x-3 mb-2">
                      <h3 className="text-lg font-semibold text-gray-900">{flight.id}</h3>
                      <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${getStatusColor(flight.statusInfo?.name || '')}`}>
                        <StatusIcon className="w-3 h-3 mr-1" />
                        {(flight.statusInfo?.name || 'Unknown').replace('_', ' ').toUpperCase()}
                      </span>
                    </div>
                    <p className="text-sm text-gray-600">{flight.airline?.name || 'Unknown Airline'}</p>
                  </div>
                  <button
                    onClick={() => {
                      setSelectedFlight(flight);
                      setShowStatusModal(true);
                    }}
                    className="text-blue-600 hover:text-blue-700 text-sm font-medium"
                  >
                    Update Status
                  </button>
                </div>

                {/* Route Information */}
                <div className="grid grid-cols-2 gap-4 mb-4">
                  <div>
                    <p className="text-xs text-gray-500 uppercase tracking-wide mb-1">
                      {getFlightType(flight)}
                    </p>
                    <div className="flex items-center space-x-2">
                      <MapPin className="w-4 h-4 text-gray-400" />
                      <span className="font-medium text-gray-900">
                        {getFlightType(flight) === 'departure' ? flight.to : flight.from}
                      </span>
                    </div>
                    <p className="text-sm text-gray-600 mt-1">
                      {formatDate(flight.scheduledDeparture)} • {formatTime(flight.scheduledDeparture)}
                    </p>
                    {flight.actualDeparture && flight.actualDeparture !== flight.scheduledDeparture && (
                      <p className="text-sm text-orange-600">
                        Actual: {formatTime(flight.actualDeparture)}
                      </p>
                    )}
                  </div>
                  <div>
                    <p className="text-xs text-gray-500 uppercase tracking-wide mb-1">Gate</p>
                    <p className="font-medium text-gray-900">{flight.gate || 'TBD'}</p>
                    <p className="text-xs text-gray-500 uppercase tracking-wide mb-1 mt-2">Terminal</p>
                    <p className="font-medium text-gray-900">{flight.terminalId || 'N/A'}</p>
                  </div>
                </div>

                {/* Stats */}
                <div className="grid grid-cols-4 gap-4 pt-4 border-t border-gray-200">
                  <div className="text-center">
                    <div className="flex items-center justify-center mb-1">
                      <Users className="w-4 h-4 text-blue-500" />
                    </div>
                    <p className="text-xs text-gray-500">Passengers</p>
                    <p className="font-medium text-gray-900">{flight.passengerCount}/{flight.capacity}</p>
                    <p className="text-xs text-gray-600">{occupancy}%</p>
                  </div>
                  <div className="text-center">
                    <div className="flex items-center justify-center mb-1">
                      <Package className="w-4 h-4 text-purple-500" />
                    </div>
                    <p className="text-xs text-gray-500">Baggage</p>
                    <p className="font-medium text-gray-900">{flight.baggageCount}</p>
                  </div>
                  <div className="text-center">
                    <div className="flex items-center justify-center mb-1">
                      <CheckCircle className="w-4 h-4 text-green-500" />
                    </div>
                    <p className="text-xs text-gray-500">Checked In</p>
                    <p className="font-medium text-gray-900">{flight.checkedInCount}</p>
                  </div>
                  <div className="text-center">
                    <div className="flex items-center justify-center mb-1">
                      <DollarSign className="w-4 h-4 text-emerald-500" />
                    </div>
                    <p className="text-xs text-gray-500">Revenue</p>
                    <p className="font-medium text-gray-900">${(flight.revenue / 1000).toFixed(0)}k</p>
                  </div>
                </div>

                {/* Occupancy Bar */}
                <div className="mt-4">
                  <div className="flex justify-between text-xs text-gray-500 mb-1">
                    <span>Occupancy</span>
                    <span>{occupancy}%</span>
                  </div>
                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div
                      className={`h-2 rounded-full transition-all duration-300 ${
                        occupancy >= 90 ? 'bg-red-500' :
                        occupancy >= 70 ? 'bg-orange-500' :
                        'bg-green-500'
                      }`}
                      style={{ width: `${occupancy}%` }}
                    />
                  </div>
                </div>
              </div>
            </div>
          );
        })}
      </div>

      {/* Status Update Modal */}
      {showStatusModal && selectedFlight && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md mx-4">
            <h3 className="text-lg font-semibold text-gray-900 mb-4">
              Update Flight Status - {selectedFlight.id}
            </h3>
            
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-2">
                  Status
                </label>
                <select
                  onChange={(e) => {
                    if (e.target.value) {
                      handleUpdateStatus(selectedFlight.id, e.target.value);
                    }
                  }}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  defaultValue=""
                >
                  <option value="">Select new status...</option>
                  <option value="ON_TIME">On Time</option>
                  <option value="DELAYED">Delayed</option>
                  <option value="BOARDING">Boarding</option>
                  <option value="DEPARTED">Departed</option>
                  <option value="CANCELLED">Cancelled</option>
                </select>
              </div>
            </div>
            
            <div className="flex justify-end space-x-3 mt-6">
              <button
                onClick={() => {
                  setShowStatusModal(false);
                  setSelectedFlight(null);
                }}
                className="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}

      {flights.length === 0 && (
        <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-8">
          <div className="text-center">
            <Plane className="w-12 h-12 text-gray-400 mx-auto mb-4" />
            <p className="text-gray-600">No flights found</p>
          </div>
        </div>
      )}
    </div>
  );
};
