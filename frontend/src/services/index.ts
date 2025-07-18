/**
 * Services Module Export
 * 
 * This module serves as the central export point for all service modules
 * in the Minden Airport Frontend application. Services handle API communication,
 * data transformation, and business logic operations.
 * 
 * Available services:
 * - authService: User authentication and session management
 * - flightService: Flight data retrieval and management
 * - baggageService: Baggage tracking and reporting
 * - adminService: Administrative operations and management
 */

export { authService } from './authService'
export { flightService } from './flightService'
export { baggageService } from './baggageService'
export { adminService } from './adminService'
