import { Observable } from "rxjs"

export enum Role {
  ADMIN = 'ADMIN',
  CLIENT = 'CLIENT'
}

export interface UserRequest {
  name: string
  email: string
  password: string
  role: Role
}

export interface User {
  id: string
  name: string
  email: string
  is_active: boolean
  role: Role
  created_at: string
  updated_at: string
}

export interface UserResponse {
  token: string
  user: User
}

export interface UserValidated {
  token: string
  id: string
  email: string
}

export interface Jwt {
  token: string
}

export interface VerificationCodeRequest {
  code: string
  userId: string
}

export interface ResendVerificationCodeRequest {
  email: string
}

export interface ResendVerificationCodeResponse {
  message: string
}

export interface SendResetPasswordRequest {
  frontBaseUrl: string
  email: string
}

export interface SendResetPasswordResponse {
  message: string
}

export interface ResetPasswordRequest {
  token: string
  email: string
  password: string
}

export interface ResetPasswordResponse {
  message: string
}

export interface AuthService {
  CreateUser(data: UserRequest): Observable<User>;
  ValidateUser(data: UserRequest): Observable<UserValidated>;
  JwtParse(data: Jwt): Observable<User>;
  ActivateUser(data: VerificationCodeRequest): Observable<UserResponse>;
  ResendVerificationCode(data: ResendVerificationCodeRequest): Observable<ResendVerificationCodeResponse>;
  SendResetPassword(data: SendResetPasswordRequest): Observable<SendResetPasswordResponse>;
  ResetPassword(data: ResetPasswordRequest): Observable<ResetPasswordResponse>;
}

export const AUTH_PACKAGE = 'auth'
export const USER_SERVICE_NAME = 'AuthService'