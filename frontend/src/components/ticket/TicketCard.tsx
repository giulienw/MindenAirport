import type { TicketWithFlight } from '@/types';
import { Plane, Clock, MapPin, CreditCard } from 'lucide-react';

interface TicketCardProps {
  ticket: TicketWithFlight;
}

export function TicketCard({ ticket }: TicketCardProps) {
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
        weekday: 'short',
        month: 'short',
        day: 'numeric',
      });
    } catch {
      return 'N/A';
    }
  };

  const getStatusColor = (status?: string) => {
    switch (status?.toUpperCase()) {
      case 'CONFIRMED':
        return 'text-green-600 bg-green-100';
      case 'CHECKED_IN':
        return 'text-blue-600 bg-blue-100';
      case 'CANCELLED':
        return 'text-red-600 bg-red-100';
      default:
        return 'text-gray-600 bg-gray-100';
    }
  };

  const getTravelClassColor = (travelClass?: string) => {
    switch (travelClass?.toLowerCase()) {
      case 'first':
        return 'text-purple-600 bg-purple-100';
      case 'business':
        return 'text-blue-600 bg-blue-100';
      case 'economy':
        return 'text-gray-600 bg-gray-100';
      default:
        return 'text-gray-600 bg-gray-100';
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-lg border border-gray-200 overflow-hidden hover:shadow-xl transition-shadow duration-300">
      {/* Ticket Header */}
      <div className="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-2">
            <Plane className="h-5 w-5" />
            <span className="font-semibold text-lg">
              {ticket.flightInfo?.fromAirport?.city || ticket.from} → {ticket.flightInfo?.toAirport?.city || ticket.to}
            </span>
          </div>
          <span className={`px-3 py-1 rounded-full text-sm font-medium ${getStatusColor(ticket.status)}`}>
            {ticket.status}
          </span>
        </div>
      </div>

      {/* Ticket Body */}
      <div className="p-6">
        {/* Flight Details */}
        <div className="grid grid-cols-2 gap-6 mb-6">
          <div>
            <div className="text-sm text-gray-500 mb-1">Departure</div>
            <div className="font-bold text-xl text-gray-900">
              {ticket.flightInfo?.scheduledDeparture ? formatTime(ticket.flightInfo.scheduledDeparture) : formatTime(ticket.departureTime || '')}
            </div>
            <div className="text-sm text-gray-600">
              {ticket.flightInfo?.scheduledDeparture ? formatDate(ticket.flightInfo.scheduledDeparture) : formatDate(ticket.departureTime || '')}
            </div>
            <div className="text-sm text-gray-500 mt-1">
              {ticket.flightInfo?.fromAirport?.name || ticket.from}
            </div>
          </div>
          
          <div>
            <div className="text-sm text-gray-500 mb-1">Arrival</div>
            <div className="font-bold text-xl text-gray-900">
              {ticket.flightInfo?.scheduledArrival ? formatTime(ticket.flightInfo.scheduledArrival) : 'TBD'}
            </div>
            <div className="text-sm text-gray-600">
              {ticket.flightInfo?.scheduledArrival ? formatDate(ticket.flightInfo.scheduledArrival) : 'TBD'}
            </div>
            <div className="text-sm text-gray-500 mt-1">
              {ticket.flightInfo?.toAirport?.name || ticket.to}
            </div>
          </div>
        </div>

        {/* Ticket Information */}
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-6">
          <div className="text-center">
            <div className="text-sm text-gray-500 mb-1">Seat</div>
            <div className="font-semibold text-gray-900">{ticket.seatNumber || 'N/A'}</div>
          </div>
          <div className="text-center">
            <div className="text-sm text-gray-500 mb-1">Class</div>
            <span className={`px-2 py-1 rounded-full text-xs font-medium ${getTravelClassColor(ticket.travelClass)}`}>
              {ticket.travelClass || 'N/A'}
            </span>
          </div>
          <div className="text-center">
            <div className="text-sm text-gray-500 mb-1">Gate</div>
            <div className="font-semibold text-gray-900">{ticket.gate || ticket.flightInfo?.gate || 'TBD'}</div>
          </div>
          <div className="text-center">
            <div className="text-sm text-gray-500 mb-1">Terminal</div>
            <div className="font-semibold text-gray-900">{ticket.flightInfo?.terminalId || 'N/A'}</div>
          </div>
        </div>

        {/* Additional Information */}
        <div className="flex justify-between items-center pt-4 border-t border-gray-200">
          <div className="flex items-center space-x-4 text-sm text-gray-600">
            {ticket.baggageClaim && (
              <div className="flex items-center space-x-1">
                <MapPin className="h-4 w-4" />
                <span>Baggage: {ticket.baggageClaim}</span>
              </div>
            )}
            {ticket.bookingDate && (
              <div className="flex items-center space-x-1">
                <Clock className="h-4 w-4" />
                <span>Booked: {formatDate(ticket.bookingDate)}</span>
              </div>
            )}
          </div>
          {ticket.price && (
            <div className="flex items-center space-x-1 text-lg font-bold text-green-600">
              <CreditCard className="h-4 w-4" />
              <span>${ticket.price.toFixed(2)}</span>
            </div>
          )}
        </div>

        {/* Flight Status */}
        {ticket.flightInfo?.statusInfo && (
          <div className="mt-4 p-3 bg-gray-50 rounded-lg">
            <div className="flex items-center justify-between">
              <span className="text-sm font-medium text-gray-700">Flight Status:</span>
              <span className={`px-2 py-1 rounded-full text-xs font-medium ${getStatusColor(ticket.flightInfo.statusInfo.name)}`}>
                {ticket.flightInfo.statusInfo.name}
              </span>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
