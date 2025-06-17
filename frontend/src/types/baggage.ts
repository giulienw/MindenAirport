export interface BaggageItem {
  id: string;
  flightId: string;
  type: BaggageType;
  weight?: number;
  status: BaggageStatus;
  trackingNumber?: string;
  specialHandling?: string;
}

export type BaggageType = 'CARRY_ON' | 'CHECKED' | 'PERSONAL';
export type BaggageStatus = 'CHECKED' | 'IN_TRANSIT' | 'DELIVERED' | 'LOST';


export interface BaggageTracking {
  id: string;
  baggageId: string;
  location: string;
  timestamp: string;
  status: string;
  notes?: string;
}

export interface UserBaggage {
  items: BaggageItem[];
  totalItems: number;
  checkedBags: number;
  carryOnBags: number;
}
