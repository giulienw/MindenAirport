/**
 * Flight Management Type Definitions
 * 
 * This module contains all type definitions related to flight operations,
 * including flights, airlines, airports, and flight status management.
 * These types support the core aviation functionality of the airport system.
 */

/**
 * Core Flight entity representing a scheduled flight
 */
export interface Flight {
  /** Unique flight identifier */
  id: string
  /** Origin airport code */
  from: string
  /** Destination airport code */
  to: string
  /** Flight date in ISO format */
  date: string
  /** ID of the assigned pilot */
  pilotId: string
  /** ID of the assigned aircraft */
  planeId: string
  /** ID of the operating airline */
  airlineId: string
  /** ID of the assigned terminal (optional) */
  terminalId?: string
  /** Current status ID (references FlightStatus) */
  statusId?: number
  /** Scheduled departure time in ISO format */
  scheduledDeparture: string
  /** Actual departure time in ISO format (if departed) */
  actualDeparture?: string
  /** Scheduled arrival time in ISO format */
  scheduledArrival: string
  /** Actual arrival time in ISO format (if arrived) */
  actualArrival?: string
  /** Assigned gate number (optional) */
  gate?: string
  /** Baggage claim carousel number (optional) */
  baggageClaim?: string
}

/**
 * Flight status definition for tracking flight states
 */
export interface FlightStatus {
  /** Unique status identifier */
  id: number
  /** Status name (e.g., "On Time", "Delayed", "Cancelled") */
  name: string
  /** Optional detailed description of the status */
  description?: string
}

/**
 * Airline information for flight operations
 */
export interface Airline {
  /** Unique airline identifier */
  id: string
  /** Airline name (e.g., "Delta Airlines") */
  name: string
  /** Country where airline is based */
  country: string
  /** URL to airline logo image */
  logo: string
  /** Whether the airline is currently active */
  active: boolean
}

/**
 * Airport information for flight routing
 */
export interface Airport {
  /** Unique airport code (e.g., "LAX", "JFK") */
  id: string
  /** Full airport name (optional) */
  name?: string
  /** Country where airport is located */
  country: string
  /** City where airport is located */
  city: string
  /** Airport timezone (optional) */
  timezone?: string
  /** Airport elevation in feet (optional) */
  elevation?: number
  /** Number of terminals at the airport (optional) */
  numberOfTerminals?: number
  /** Airport latitude coordinate (optional) */
  latitude?: number
  /** Airport longitude coordinate (optional) */
  longitude?: number
}

/**
 * Enhanced flight information with enriched data
 * 
 * This extends the base Flight interface with additional relational data
 * for improved user experience and detailed flight information display.
 */
export interface FlightDisplayInfo extends Flight {
  /** Enriched origin airport information */
  fromAirport?: Airport
  /** Enriched destination airport information */
  toAirport?: Airport
  /** Current flight status information */
  statusInfo?: FlightStatus
  /** Operating airline information */
  airline?: Airline
}
