import type { FlightDisplayInfo } from './flight';

export interface Ticket {
  id: string;
  airportUserID: string; // Match backend field name
  flight: string;
  seatNumber?: string;
  travelClass?: string;
  price?: number;
  bookingDate?: string;
  status?: 'CONFIRMED' | 'CANCELLED' | 'CHECKED_IN';
  from?: string;
  to?: string;
  gate?: string;
  baggageClaim?: string;
  departureTime?: string;
}

export interface TicketWithFlight extends Ticket {
  flightInfo?: FlightDisplayInfo;
}

export interface UserDashboard {
  user: {
    id: string;
    firstName: string;
    lastName: string;
    email?: string; // Make email optional to match backend User type
  };
  tickets: TicketWithFlight[];
  upcomingFlights: TicketWithFlight[];
  pastFlights: TicketWithFlight[];
}
