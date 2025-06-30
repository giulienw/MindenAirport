import type { Flight } from '@/types'
import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function getFlightType(flight: Flight): string {
  if(flight.from === "MIN") {
    return "departure"
  } else {
    return "arrival"
  }
}