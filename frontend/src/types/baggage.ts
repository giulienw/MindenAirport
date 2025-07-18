/**
 * Baggage Management Type Definitions
 * 
 * This module contains all type definitions related to baggage tracking,
 * management, and reporting within the airport system. Supports comprehensive
 * baggage lifecycle management from check-in to delivery.
 */

/**
 * Core baggage item entity representing a single piece of luggage
 */
export interface BaggageItem {
  /** Unique baggage identifier */
  id: string;
  /** Associated flight ID for the baggage */
  flightId: string;
  /** Type of baggage (carry-on, checked, personal item) */
  type: BaggageType;
  /** Weight of the baggage in appropriate units (optional) */
  weight?: number;
  /** Current status of the baggage in the system */
  status: BaggageStatus;
  /** Unique tracking number for customer reference (optional) */
  trackingNumber?: string;
  /** Special handling instructions or requirements (optional) */
  specialHandling?: string;
}

/**
 * Baggage type enumeration defining categories of luggage
 */
export type BaggageType = 'CARRY_ON' | 'CHECKED' | 'PERSONAL';

/**
 * Baggage status enumeration tracking baggage lifecycle
 */
export type BaggageStatus = 'CHECKED' | 'IN_TRANSIT' | 'DELIVERED' | 'LOST';

/**
 * Baggage tracking event for location and status history
 */
export interface BaggageTracking {
  /** Unique tracking event identifier */
  id: string;
  /** Associated baggage ID */
  baggageId: string;
  /** Current location of the baggage */
  location: string;
  /** Timestamp of the tracking event */
  timestamp: string;
  /** Status at this tracking point */
  status: string;
  /** Additional notes or comments (optional) */
  notes?: string;
}

/**
 * User baggage summary containing aggregated information
 */
export interface UserBaggage {
  /** Array of baggage items belonging to the user */
  items: BaggageItem[];
  /** Total number of baggage items */
  totalItems: number;
  /** Number of checked baggage items */
  checkedBags: number;
  /** Number of carry-on baggage items */
  carryOnBags: number;
}
