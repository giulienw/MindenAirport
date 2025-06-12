export interface Flight {
  id: string;
  from: string;
  to: string;
  date: string;
  pilotId: string;
  planeId: string;
  airlineId: string;
  terminalId?: string;
  statusId?: string;
  scheduledDeparture: string;
  actualDeparture?: string;
  scheduledArrival: string;
  actualArrival?: string;
  gate?: string;
  baggageClaim?: string;
}

export interface FlightStatus {
  id: string;
  name: string;
  description?: string;
}

export interface Airline {
  id: string;
  name: string;
  country: string;
  logo: string;
  active: boolean;
}

export interface Airport {
  id: string;
  name?: string;
  country: string;
  city: string;
  timezone?: string;
  elevation?: number;
  numberOfTerminals?: number;
  latitude?: number;
  longitude?: number;
}

export interface FlightDisplayInfo extends Flight {
  fromAirport?: Airport;
  toAirport?: Airport;
  statusInfo?: FlightStatus;
  airline?: Airline;
}
