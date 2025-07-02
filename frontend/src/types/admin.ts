import type { Airline, Airport, FlightDisplayInfo } from './flight';
import type { BaggageItem } from './baggage';

export interface AdminDashboardStats {
  totalFlights: number;
  totalPassengers: number;
  totalBaggage: number;
  activeAirlines: number;
  revenue: number;
}

export interface SystemAlert {
  id: string;
  type: 'ERROR' | 'WARNING' | 'INFO';
  message: string;
  timestamp: string;
  resolved: boolean;
  category: 'FLIGHT' | 'BAGGAGE' | 'SYSTEM' | 'SECURITY';
}

export interface FlightManagement extends FlightDisplayInfo {
  passengerCount: number;
  capacity: number;
  revenue: number;
  baggageCount: number;
  checkedInCount: number;
}

export interface BaggageManagement extends BaggageItem {
  passengerName: string;
  flightNumber: string;
  issueReported: boolean;
  handlingNotes?: string;
}

export interface AdminActivity {
  id: string;
  adminId: string;
  adminName: string;
  action: string;
  details: string;
  timestamp: string;
  category: 'USER' | 'FLIGHT' | 'BAGGAGE' | 'SYSTEM';
}

export interface AdminDashboard {
  stats: AdminDashboardStats;
  airlines: Airline[];
  airports: Airport[];
  flights: FlightDisplayInfo[];
}
