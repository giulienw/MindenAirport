import type { Flight } from '@/types'
import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function getFlightType(flight: Flight): string {
  if (flight.from === "MIN") {
    return "departure"
  } else {
    return "arrival"
  }
}

export const getFlightStatusColor = (status?: string) => {
  switch (status?.toLowerCase()) {
    case 'departed':
      return 'text-green-600 bg-green-100'
    case 'arrived':
      return 'text-green-600 bg-green-100';
    case 'delayed':
      return 'text-yellow-600 bg-yellow-100';
    case 'cancelled':
      return 'text-red-600 bg-red-100';
    case 'boarding':
      return 'text-blue-600 bg-blue-100';
    default:
      return 'text-gray-600 bg-gray-100';
  }
};

export function getCookie(key: string) {
  const b = document.cookie.match("(^|;)\\s*" + key + "\\s*=\\s*([^;]+)");
  return b ? b.pop() : "";
}