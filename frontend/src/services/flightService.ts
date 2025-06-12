import type { Flight, Airport, FlightStatus, Airline, FlightDisplayInfo } from '@/types';
import { API_BASE_URL } from '@/config';

export const flightService = {
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
        statusInfo: statusMap.get(flight.statusId || ''),
        airline: airlineMap.get(flight.airlineId),
      }));
    } catch (error) {
      console.error('Failed to fetch enriched flight data:', error);
      throw error;
    }
  },
};
