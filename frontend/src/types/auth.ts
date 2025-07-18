/**
 * Authentication and Authorization Type Definitions
 * 
 * This module contains all type definitions related to user authentication,
 * authorization, and user management within the Minden Airport system.
 * 
 * Key features:
 * - User role-based access control (USER, ADMIN, STAFF, MANAGER)
 * - Admin permission system for granular access control
 * - Authentication request/response types
 * - Error handling for auth operations
 */

/**
 * Core User entity representing a registered user in the system
 */
export interface User {
  /** Unique identifier for the user */
  id: string
  /** User's email address (optional for some auth methods) */
  email?: string
  /** User's first name */
  firstName: string
  /** User's last name */
  lastName: string
  /** User's date of birth (optional) */
  birthdate?: string
  /** User's password (not included in responses for security) */
  password?: string
  /** Whether the user account is active */
  active: boolean
  /** User's phone number (optional) */
  phone?: string
  /** User's role determining access level */
  role: UserRole
  /** Number of tickets associated with the user (optional) */
  ticketCount?: number
  /** Timestamp of user's last login (optional) */
  lastLogin?: string
}

/**
 * User role enumeration defining access levels within the system
 */
export type UserRole = 'USER' | 'ADMIN' | 'STAFF' | 'MANAGER'

/**
 * Admin permission types for granular access control
 * Used to determine what administrative functions a user can access
 */
export type AdminPermission = 
  | 'MANAGE_FLIGHTS'    // Create, update, delete flights
  | 'MANAGE_USERS'      // Manage user accounts and permissions
  | 'MANAGE_BAGGAGE'    // Handle baggage tracking and issues
  | 'MANAGE_TICKETS'    // Manage ticket bookings and cancellations
  | 'VIEW_ANALYTICS'    // Access to dashboard analytics and reports
  | 'SYSTEM_SETTINGS'   // Configure system-wide settings

/**
 * User login credentials for authentication
 */
export interface LoginCredentials {
  /** User's email address */
  email: string
  /** User's password */
  password: string
  /** Whether to keep the user logged in across sessions */
  rememberMe?: boolean
}

/**
 * User registration data for creating new accounts
 */
export interface RegisterData {
  /** User's email address (must be unique) */
  email: string
  /** User's chosen password */
  password: string
  /** User's first name */
  firstName: string
  /** User's last name */
  lastName: string
  /** Confirmation that user accepts terms and conditions */
  acceptTerms: boolean
}

/**
 * Response structure for successful authentication operations
 */
export interface AuthResponse {
  /** The authenticated user's information */
  user: User
  /** JWT token for subsequent API requests */
  token: string
}

/**
 * Standardized API error response structure
 */
export interface ApiError {
  /** Human-readable error message */
  message: string
  /** Optional error code for programmatic handling */
  code?: string
  /** HTTP status code */
  status?: number
}
