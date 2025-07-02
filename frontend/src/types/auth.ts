export interface User {
  id: string
  email?: string
  firstName: string
  lastName: string
  birthDate?: string
  password?: string
  active: boolean
  phone?: string
  role?: UserRole
  ticketCount?: number
  lastLogin?: string
}

export type UserRole = 'USER' | 'ADMIN' | 'STAFF' | 'MANAGER'


export type AdminPermission = 
  | 'MANAGE_FLIGHTS'
  | 'MANAGE_USERS' 
  | 'MANAGE_BAGGAGE'
  | 'MANAGE_TICKETS'
  | 'VIEW_ANALYTICS'
  | 'SYSTEM_SETTINGS'

export interface LoginCredentials {
  email: string
  password: string
  rememberMe?: boolean
}

export interface RegisterData {
  email: string
  password: string
  firstName: string
  lastName: string
  acceptTerms: boolean
}

export interface AuthResponse {
  user: User
  token: string
}

export interface ApiError {
  message: string
  code?: string
  status?: number
}
