import type { AdminUser, UserRole } from './auth';
import type { FlightDisplayInfo } from './flight';
import type { BaggageItem } from './baggage';

export interface AdminDashboardStats {
  totalFlights: number;
  activeFlights: number;
  totalPassengers: number;
  totalBaggage: number;
  delayedFlights: number;
  lostBaggage: number;
  revenue: number;
  capacity: number;
}

export interface SystemAlert {
  id: string;
  type: 'ERROR' | 'WARNING' | 'INFO';
  message: string;
  timestamp: string;
  resolved: boolean;
  category: 'FLIGHT' | 'BAGGAGE' | 'SYSTEM' | 'SECURITY';
}

export interface UserManagement {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  role: UserRole;
  active: boolean;
  lastLogin?: string;
  createdAt: string;
  ticketCount: number;
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
  admin: AdminUser;
  stats: AdminDashboardStats;
  alerts: SystemAlert[];
  recentActivity: AdminActivity[];
  criticalFlights: FlightManagement[];
  problematicBaggage: BaggageManagement[];
}
