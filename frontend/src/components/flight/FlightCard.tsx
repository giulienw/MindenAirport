/**
 * FlightCard Component
 * 
 * A reusable card component for displaying flight information in a structured format.
 * Shows comprehensive flight details including departure/arrival times, route information,
 * status, and additional flight details like gate and terminal information.
 * 
 * Features:
 * - Responsive design with clean visual hierarchy
 * - Status-based color coding for quick identification
 * - Support for both scheduled and actual times
 * - Route visualization with city names
 * - Additional details like gate, terminal, and baggage claim
 * 
 * @param props - Component props
 * @param props.flight - Flight data to display
 */

import { type FlightDisplayInfo } from '@/types';
import { getFlightStatusColor as getStatusColor } from '@/lib/utils';

interface FlightCardProps {
  /** Flight data to display in the card */
  flight: FlightDisplayInfo;
}

export function FlightCard({ flight }: FlightCardProps) {
  const formatTime = (dateString: string) => {
    try {
      return new Date(dateString).toLocaleTimeString('en-US', {
        hour: '2-digit',
        minute: '2-digit',
        hour12: false,
      });
    } catch {
      return 'N/A';
    }
  };

  const formatDate = (dateString: string) => {
    try {
      return new Date(dateString).toLocaleDateString('en-US', {
        month: 'short',
        day: 'numeric',
      });
    } catch {
      return 'N/A';
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-md hover:shadow-lg transition-shadow duration-200 p-6 border border-gray-200">
      {/* Header with Flight Status */}
      <div className="flex justify-between items-start mb-4">
        <div className="flex items-center space-x-2">
          <div className="text-lg font-semibold text-gray-900">
            {flight.fromAirport?.city || flight.from} → {flight.toAirport?.city || flight.to}
          </div>
        </div>
        <span className={`px-3 py-1 rounded-full text-sm font-medium ${getStatusColor(flight.statusInfo?.name)}`}>
          {flight.statusInfo?.name || 'Unknown'}
        </span>
      </div>

      {/* Flight Details */}
      <div className="grid grid-cols-2 gap-4 mb-4">
        <div>
          <div className="text-sm text-gray-500 mb-1">Departure</div>
          <div className="font-semibold text-lg text-gray-900">
            {formatTime(flight.scheduledDeparture)}
          </div>
          <div className="text-sm text-gray-600">
            {formatDate(flight.scheduledDeparture)}
          </div>
          {flight.actualDeparture && flight.actualDeparture !== flight.scheduledDeparture && (
            <div className="text-sm text-red-600">
              Actual: {formatTime(flight.actualDeparture)}
            </div>
          )}
        </div>
        <div>
          <div className="text-sm text-gray-500 mb-1">Arrival</div>
          <div className="font-semibold text-lg text-gray-900">
            {formatTime(flight.scheduledArrival)}
          </div>
          <div className="text-sm text-gray-600">
            {formatDate(flight.scheduledArrival)}
          </div>
          {flight.actualArrival && flight.actualArrival !== flight.scheduledArrival && (
            <div className="text-sm text-red-600">
              Actual: {formatTime(flight.actualArrival)}
            </div>
          )}
        </div>
      </div>

      {/* Additional Info */}
      <div className="flex justify-between items-center text-sm text-gray-600 pt-4 border-t border-gray-200">
        <div className="flex space-x-4">
          {flight.gate && (
            <span>
              <span className="font-medium">Gate:</span> {flight.gate}
            </span>
          )}
          {flight.terminalId && (
            <span>
              <span className="font-medium">Terminal:</span> {flight.terminalId}
            </span>
          )}
        </div>
        {flight.baggageClaim && (
          <span>
            <span className="font-medium">Baggage:</span> {flight.baggageClaim}
          </span>
        )}
      </div>
    </div>
  );
}
