/**
 * FlightManagement Component
 * 
 * Administrative interface for managing flight operations including viewing,
 * updating, and monitoring all flights in the system. Provides comprehensive
 * flight management capabilities for airport administrators.
 * 
 * Features:
 * - Real-time flight data display with enriched information
 * - Flight status update functionality with modal interface
 * - Comprehensive flight details (route, times, gate, terminal)
 * - Error handling and loading states
 * - Responsive grid layout for multiple flights
 * - Status-based visual indicators and color coding
 * 
 * Access: Requires admin authentication and appropriate permissions
 * 
 * @returns JSX element containing the flight management interface
 */

import React, { useState, useEffect } from 'react';
import { Plane, Clock, MapPin, Users, AlertTriangle } from 'lucide-react';
import type { FlightManagement as FlightManagementType } from '@/types';
import { adminService, flightService } from '@/services';
import { getFlightType, getFlightStatusColor as getStatusColor } from '@/lib/utils';

export const FlightManagement: React.FC = () => {
  // Flight management data state
  const [flights, setFlights] = useState<FlightManagementType[]>([]);
  // Loading state for UI feedback
  const [loading, setLoading] = useState(true);
  // Error state for error handling and display
  const [error, setError] = useState<string | null>(null);
  // Selected flight for status update operations
  const [selectedFlight, setSelectedFlight] = useState<FlightManagementType | null>(null);
  // Modal visibility state for status updates
  const [showStatusModal, setShowStatusModal] = useState(false);

  // Fetch flight data on component mount
  useEffect(() => {
    fetchFlights();
  }, []);

  /**
   * Fetches flight management data from the service layer
   * Handles loading states and error management
   */
  const fetchFlights = async () => {
    try {
      setLoading(true);
      setError(null);
      const result = await flightService.getEnrichedFlights();
      setFlights(result as FlightManagementType[]);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to fetch flights');
    } finally {
      setLoading(false);
    }
  };

  /**
   * Updates flight status through the admin service
   * 
   * @param flightId - ID of the flight to update
   * @param newStatus - New status ID to assign to the flight
   */
  const handleUpdateStatus = async (flightId: string, newStatus: string) => {
    try {
      await adminService.updateFlightStatus(flightId, Number(newStatus));
      await fetchFlights();
      setShowStatusModal(false);
      setSelectedFlight(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update flight status');
    }
  };

  /**
   * Returns appropriate icon component based on flight status
   * 
   * @param status - Flight status string
   * @returns Lucide icon component for the status
   */
  const getStatusIcon = (status: string) => {
    switch (status.toLowerCase()) {
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

  /**
   * Formats time string for display in user-friendly format
   * 
   * @param dateString - ISO date string to format
   * @returns Formatted time string (HH:MM AM/PM)
   */
  const formatTime = (dateString: string) => {
    return new Date(dateString).toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  /**
   * Formats date string for display in user-friendly format
   * 
   * @param dateString - ISO date string to format
   * @returns Formatted date string (MM/DD/YYYY)
   */
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      month: 'short',
      day: 'numeric',
    });
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
          console.log(flight)
          
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
                    {/*<p className="text-sm text-gray-600">{flight.airline?.name || 'Unknown Airline'}</p> */}
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
                  <option value="5">Delayed</option>
                  <option value="2">Boarding</option>
                  <option value="3">Departed</option>
                  <option value="6">Cancelled</option>
                  <option value="4">Arrived</option> 
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
