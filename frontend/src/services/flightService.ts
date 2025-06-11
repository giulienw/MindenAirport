import type { Flight, Airport, FlightStatus, Airline, FlightDisplayInfo } from '@/types';

const API_BASE_URL = 'http://localhost:8080/api';

// Mock data for demonstration when backend is not available
const mockFlights: Flight[] = [
  {
    id: '1',
    from: 'LAX',
    to: 'JFK',
    date: new Date().toISOString(),
    pilotId: 'pilot1',
    planeId: 'plane1',
    terminalId: 'A',
    status: 'on-time',
    scheduledDeparture: new Date(Date.now() + 2 * 60 * 60 * 1000).toISOString(), // 2 hours from now
    scheduledArrival: new Date(Date.now() + 7 * 60 * 60 * 1000).toISOString(), // 7 hours from now
    gate: 'A12',
    baggageClaim: 'BC1',
  },
  {
    id: '2',
    from: 'ORD',
    to: 'MIA',
    date: new Date().toISOString(),
    pilotId: 'pilot2',
    planeId: 'plane2',
    terminalId: 'B',
    status: 'delayed',
    scheduledDeparture: new Date(Date.now() + 1 * 60 * 60 * 1000).toISOString(), // 1 hour from now
    actualDeparture: new Date(Date.now() + 1.5 * 60 * 60 * 1000).toISOString(), // 1.5 hours from now
    scheduledArrival: new Date(Date.now() + 4 * 60 * 60 * 1000).toISOString(), // 4 hours from now
    gate: 'B8',
    baggageClaim: 'BC2',
  },
  {
    id: '3',
    from: 'DEN',
    to: 'SEA',
    date: new Date().toISOString(),
    pilotId: 'pilot3',
    planeId: 'plane3',
    terminalId: 'C',
    status: 'boarding',
    scheduledDeparture: new Date(Date.now() + 30 * 60 * 1000).toISOString(), // 30 minutes from now
    scheduledArrival: new Date(Date.now() + 3 * 60 * 60 * 1000).toISOString(), // 3 hours from now
    gate: 'C15',
    baggageClaim: 'BC3',
  },
  {
    id: '4',
    from: 'ATL',
    to: 'PHX',
    date: new Date().toISOString(),
    pilotId: 'pilot4',
    planeId: 'plane4',
    terminalId: 'A',
    status: 'on-time',
    scheduledDeparture: new Date(Date.now() + 3 * 60 * 60 * 1000).toISOString(), // 3 hours from now
    scheduledArrival: new Date(Date.now() + 6 * 60 * 60 * 1000).toISOString(), // 6 hours from now
    gate: 'A9',
    baggageClaim: 'BC1',
  },
  {
    id: '5',
    from: 'BOS',
    to: 'SFO',
    date: new Date().toISOString(),
    pilotId: 'pilot5',
    planeId: 'plane5',
    terminalId: 'B',
    status: 'cancelled',
    scheduledDeparture: new Date(Date.now() + 4 * 60 * 60 * 1000).toISOString(), // 4 hours from now
    scheduledArrival: new Date(Date.now() + 10 * 60 * 60 * 1000).toISOString(), // 10 hours from now
    gate: 'B3',
    baggageClaim: 'BC2',
  },
];

const mockAirports: Airport[] = [
  { id: 'LAX', name: 'Los Angeles International Airport', city: 'Los Angeles', country: 'USA' },
  { id: 'JFK', name: 'John F. Kennedy International Airport', city: 'New York', country: 'USA' },
  { id: 'ORD', name: "O'Hare International Airport", city: 'Chicago', country: 'USA' },
  { id: 'MIA', name: 'Miami International Airport', city: 'Miami', country: 'USA' },
  { id: 'DEN', name: 'Denver International Airport', city: 'Denver', country: 'USA' },
  { id: 'SEA', name: 'Seattle-Tacoma International Airport', city: 'Seattle', country: 'USA' },
  { id: 'ATL', name: 'Hartsfield-Jackson Atlanta International Airport', city: 'Atlanta', country: 'USA' },
  { id: 'PHX', name: 'Phoenix Sky Harbor International Airport', city: 'Phoenix', country: 'USA' },
  { id: 'BOS', name: 'Logan International Airport', city: 'Boston', country: 'USA' },
  { id: 'SFO', name: 'San Francisco International Airport', city: 'San Francisco', country: 'USA' },
];

const mockFlightStatuses: FlightStatus[] = [
  { id: 'on-time', name: 'On Time' },
  { id: 'delayed', name: 'Delayed' },
  { id: 'boarding', name: 'Boarding' },
  { id: 'cancelled', name: 'Cancelled' },
  { id: 'departed', name: 'Departed' },
  { id: 'arrived', name: 'Arrived' },
];

const mockAirlines: Airline[] = [
  { id: '1', name: 'American Airlines', country: 'USA', logo: '', active: true },
  { id: '2', name: 'Delta Air Lines', country: 'USA', logo: '', active: true },
  { id: '3', name: 'United Airlines', country: 'USA', logo: '', active: true },
];

export const flightService = {
  async getFlights(): Promise<Flight[]> {
    try {
      const response = await fetch(`${API_BASE_URL}/flight`);
      if (!response.ok) {
        throw new Error('Failed to fetch flights');
      }
      return response.json();
    } catch (error) {
      console.warn('Backend not available, using mock data:', error);
      return mockFlights;
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
      console.warn('Backend not available, using mock data:', error);
      return mockAirports;
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
      console.warn('Backend not available, using mock data:', error);
      return mockFlightStatuses;
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
      console.warn('Backend not available, using mock data:', error);
      return mockAirlines;
    }
  },

  async getEnrichedFlights(): Promise<FlightDisplayInfo[]> {
    try {
      const [flights, airports, flightStatuses] = await Promise.all([
        this.getFlights(),
        this.getAirports(),
        this.getFlightStatuses(),
        this.getAirlines(),
      ]);

      const airportMap = new Map(airports.map(airport => [airport.id, airport]));
      const statusMap = new Map(flightStatuses.map(status => [status.id, status]));
      //const airlineMap = new Map(airlines.map(airline => [airline.id, airline]));

      return flights.map(flight => ({
        ...flight,
        fromAirport: airportMap.get(flight.from),
        toAirport: airportMap.get(flight.to),
        statusInfo: statusMap.get(flight.status || ''),
        // Note: Would need airline info associated with flights in the backend
        // airline: airlineMap.get(flight.airlineId),
      }));
    } catch (error) {
      console.error('Failed to fetch enriched flight data:', error);
      throw error;
    }
  },
};
