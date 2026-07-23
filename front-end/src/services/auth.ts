import api from "./api";

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

export interface User {
  id: number;
  name: string;
  email: string;
  role: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

/**
 * Login
 */
export async function login(
  data: LoginRequest
): Promise<LoginResponse> {
  const response = await api.post("/login", data);

  const result: LoginResponse = response.data;

  localStorage.setItem("token", result.token);

  return result;
}

/**
 * Register
 */
export async function register(
  data: RegisterRequest
) {
  const response = await api.post("/register", data);

  return response.data;
}

/**
 * Logout
 */
export function logout() {
  localStorage.removeItem("token");
}

/**
 * Mendapatkan token JWT
 */
export function getToken(): string | null {
  return localStorage.getItem("token");
}

/**
 * Mengecek apakah user sudah login
 */
export function isAuthenticated(): boolean {
  return getToken() !== null;
}

/**
 * Mengambil data profile user
 */
export async function getProfile(): Promise<User> {
  const response = await api.get("/profile");

  return response.data;
}