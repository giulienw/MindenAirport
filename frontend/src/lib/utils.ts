/**
 * Utility Functions Library
 * 
 * This module contains utility functions used throughout the Minden Airport
 * Frontend application. Includes UI utilities, data formatting functions,
 * and business logic helpers.
 * 
 * Key functions:
 * - cn: Tailwind CSS class name merging utility
 * - getFlightType: Determines if flight is departure or arrival
 * - getFlightStatusColor: Returns appropriate status colors for flights
 * - getCookie: Browser cookie retrieval utility
 */

import type { Flight } from '@/types'
import { type ClassValue, clsx } from 'clsx'
import { twMerge } from 'tailwind-merge'

/**
 * Combines and merges Tailwind CSS class names with conflict resolution
 * 
 * Uses clsx for conditional class handling and tailwind-merge for
 * resolving conflicting Tailwind classes (e.g., competing margins).
 * 
 * @param inputs - Class names, objects, or arrays to merge
 * @returns Merged and deduplicated class string
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

/**
 * Determines flight type based on flight data
 * 
 * Currently returns a placeholder value. In a full implementation,
 * this would analyze the flight's origin/destination relative to
 * the current airport to determine if it's a departure or arrival.
 * 
 * @param flight - Flight object to analyze
 * @returns Flight type string ('departure' or 'arrival')
 */
export function getFlightType(flight: Flight): string {
  if (flight.from === "MIN") {
    return "departure"
  } else {
    return "arrival"
  }
}

/**
 * Returns appropriate Tailwind CSS classes for flight status display
 * 
 * Maps flight status strings to appropriate text and background colors
 * for consistent visual representation throughout the application.
 * 
 * @param status - Flight status string (case-insensitive)
 * @returns Tailwind CSS classes for text and background colors
 */
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

/**
 * Retrieves a cookie value by key from the browser's cookie store
 * 
 * Parses the document.cookie string to find and return the value
 * associated with the specified cookie name.
 * 
 * @param key - The name of the cookie to retrieve
 * @returns The cookie value as a string, or undefined if not found
 */
export function getCookie(key: string) {
  const b = document.cookie.match("(^|;)\\s*" + key + "\\s*=\\s*([^;]+)");
  return b ? b.pop() : "";
}