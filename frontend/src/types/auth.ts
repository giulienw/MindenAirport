export interface User {
  id: string
  email: string
  firstName: string
  lastName: string
  createdAt: string
  updatedAt: string
}

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
