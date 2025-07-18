/**
 * Ticket and Booking Type Definitions
 * 
 * This module contains all type definitions related to ticket management,
 * booking operations, and user travel information within the airport system.
 * Supports comprehensive ticket lifecycle from booking to travel completion.
 */

import type { FlightDisplayInfo } from './flight';

/**
 * Core ticket entity representing a flight booking
 */
export interface Ticket {
  /** Unique ticket identifier */
  id: string;
  /** ID of the user who owns this ticket (matches backend field name) */
  airportUserID: string;
  /** Associated flight ID */
  flight: string;
  /** Assigned seat number (optional) */
  seatNumber?: string;
  /** Travel class (economy, business, first, etc.) */
  travelClass?: string;
  /** Ticket price in appropriate currency (optional) */
  price?: number;
  /** Date when the ticket was booked (optional) */
  bookingDate?: string;
  /** Current status of the ticket */
  status?: 'CONFIRMED' | 'CANCELLED' | 'CHECKED_IN';
  /** Origin airport code (optional) */
  from?: string;
  /** Destination airport code (optional) */
  to?: string;
  /** Assigned gate number (optional) */
  gate?: string;
  /** Baggage claim area (optional) */
  baggageClaim?: string;
  /** Departure time for convenience (optional) */
  departureTime?: string;
}

/**
 * Enhanced ticket with enriched flight information
 * 
 * Extends the base Ticket interface with detailed flight data
 * for improved user experience and comprehensive ticket display.
 */
export interface TicketWithFlight extends Ticket {
  /** Enriched flight information associated with the ticket */
  flightInfo?: FlightDisplayInfo;
}

/**
 * User dashboard data structure containing user and travel information
 */
export interface UserDashboard {
  /** User profile information */
  user: {
    /** User's unique identifier */
    id: string;
    /** User's first name */
    firstName: string;
    /** User's last name */
    lastName: string;
    /** User's email address (optional to match backend User type) */
    email?: string;
  };
  /** All tickets associated with the user */
  tickets: TicketWithFlight[];
  /** Upcoming flights (future departure times) */
  upcomingFlights: TicketWithFlight[];
  /** Past flights (completed travel) */
  pastFlights: TicketWithFlight[];
}
