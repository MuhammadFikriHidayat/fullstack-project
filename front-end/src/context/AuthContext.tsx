"use client";

import React, {
  createContext,
  useContext,
  useEffect,
  useState,
  ReactNode,
} from "react";

import {
  login as loginService,
  logout as logoutService,
  getToken,
  getProfile,
  User,
  LoginRequest,
} from "@/services/auth";

interface AuthContextType {
  user: User | null;
  loading: boolean;
  isAuthenticated: boolean;

  login: (data: LoginRequest) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

interface Props {
  children: ReactNode;
}

export function AuthProvider({ children }: Props) {
  const [user, setUser] = useState<User | null>(null);

  const [loading, setLoading] = useState(true);

  const login = async (data: LoginRequest) => {
    const result = await loginService(data);

    setUser(result.user);
  };

  const logout = () => {
    logoutService();

    setUser(null);
  };

  useEffect(() => {
    const initialize = async () => {
      const token = getToken();

      if (!token) {
        setLoading(false);
        return;
      }

      try {
        const profile = await getProfile();

        setUser(profile);
      } catch (err) {
        console.error(err);

        logoutService();

        setUser(null);
      }

      setLoading(false);
    };

    initialize();
  }, []);

  return (
    <AuthContext.Provider
      value={{
        user,
        loading,
        isAuthenticated: user !== null,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error("useAuth must be used inside AuthProvider");
  }

  return context;
}