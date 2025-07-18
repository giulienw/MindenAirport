/**
 * Flight Service
 * 
 * This service handles all flight-related data operations including retrieving
 * flight information, airport data, airline details, and flight statuses.
 * It provides enriched flight data by combining multiple data sources for
 * improved user experience.
 * 
 * Key features:
 * - Flight data retrieval and caching
 * - Airport and airline information management
 * - Flight status tracking
 * - Data enrichment by joining related entities
 * - Individual flight lookup by ID
 * 
 * @module FlightService
 */

import type { Flight, Airport, FlightStatus, Airline, FlightDisplayInfo } from '@/types';
import { API_BASE_URL } from '@/config';

/**
 * Flight service object containing all flight-related methods
 */
export const flightService = {
  /**
   * Retrieves all flights from the backend API
   * 
   * @returns Promise resolving to array of Flight objects
   * @throws Error if API request fails
   */
  async getFlights(): Promise<Flight[]> {
    try {
      const response = await fetch(`${API_BASE_URL}/flight`);
      if (!response.ok) {
        throw new Error('Failed to fetch flights');
      }
      return response.json();
    } catch (error) {
      console.warn('Backend not available:', error);
      return [];
    }
  },

  /**
   * Retrieves all airport information from the backend API
   * 
   * @returns Promise resolving to array of Airport objects
   * @throws Error if API request fails
   */
  async getAirports(): Promise<Airport[]> {
    try {
      const response = await fetch(`${API_BASE_URL}/airport`);
      if (!response.ok) {
        throw new Error('Failed to fetch airports');
      }
      return response.json();
    } catch (error) {
      console.warn('Backend not available:', error);
      return [];
    }
  },

  /**
   * Retrieves all flight status definitions from the backend API
   * 
   * @returns Promise resolving to array of FlightStatus objects
   * @throws Error if API request fails
   */
  async getFlightStatuses(): Promise<FlightStatus[]> {
    try {
      const response = await fetch(`${API_BASE_URL}/flightStatus`);
      if (!response.ok) {
        throw new Error('Failed to fetch flight statuses');
      }
      return response.json();
    } catch (error) {
      console.warn('Backend not available', error);
      return [];
    }
  },

  /**
   * Retrieves all airline information from the backend API
   * 
   * @returns Promise resolving to array of Airline objects
   * @throws Error if API request fails
   */
  async getAirlines(): Promise<Airline[]> {
    try {
      const response = await fetch(`${API_BASE_URL}/airline`);
      if (!response.ok) {
        throw new Error('Failed to fetch airlines');
      }
      return response.json();
    } catch (error) {
      console.warn('Backend not available', error);
      return [];
    }
  },

  /**
   * Retrieves enriched flight data by combining flights with related entities
   * 
   * This method fetches flights, airports, flight statuses, and airlines in parallel,
   * then combines them to create enriched flight objects with full relational data.
   * This reduces the need for multiple API calls in components and provides
   * complete flight information for display.
   * 
   * @returns Promise resolving to array of FlightDisplayInfo objects with enriched data
   * @throws Error if any API request fails
   */
  async getEnrichedFlights(): Promise<FlightDisplayInfo[]> {
    try {
      const [flights, airports, flightStatuses, airlines] = await Promise.all([
        this.getFlights(),
        this.getAirports(),
        this.getFlightStatuses(),
        this.getAirlines(),
      ]);

      const airportMap = new Map(airports.map(airport => [airport.id, airport]));
      const statusMap = new Map(flightStatuses.map(status => [status.id, status]));
      const airlineMap = new Map(airlines.map(airline => [airline.id, airline]));

      return flights.map(flight => ({
        ...flight,
        fromAirport: airportMap.get(flight.from),
        toAirport: airportMap.get(flight.to),
        statusInfo: statusMap.get(flight.statusId || -1),
        airline: airlineMap.get(flight.airlineId),
      }));
    } catch (error) {
      console.error('Failed to fetch enriched flight data:', error);
      throw error;
    }
  },

  /**
   * Retrieves a specific flight by its ID
   * 
   * @param flightId - The unique identifier of the flight to retrieve
   * @returns Promise resolving to FlightDisplayInfo object or null if not found
   * @throws Error if API request fails
   */
  async getFlightById(flightId: string): Promise<FlightDisplayInfo | null> {
    try {
      const response = await fetch(`${API_BASE_URL}/flight/${flightId}`);
      if (!response.ok) {
        throw new Error('Failed to fetch flight by ID');
      }
      return response.json();
    } catch (error) {
      console.error('Failed to fetch flight by ID:', error);
      throw error;
    }
  },
};
