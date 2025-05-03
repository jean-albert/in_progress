import React, { createContext, useContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import * as api from './api';

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(localStorage.getItem('token'));
  const navigate = useNavigate();

  useEffect(() => {
    if (token) {
      const fetchUser = async () => {
        try {
          const profile = await api.getProfile(token);
          setUser(profile);
        } catch (error) {
          logout();
        }
      };
      fetchUser();
    }
  }, [token]);

  const login = async (email, password) => {
    const { token } = await api.login(email, password);
    localStorage.setItem('token', token);
    setToken(token);
    navigate('/dashboard');
  };

  const register = async (email, password) => {
    const { token } = await api.register(email, password);
    localStorage.setItem('token', token);
    setToken(token);
    navigate('/dashboard');
  };

  const logout = () => {
    localStorage.removeItem('token');
    setToken(null);
    setUser(null);
    navigate('/');
  };

  return (
    <AuthContext.Provider value={{ user, token, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}