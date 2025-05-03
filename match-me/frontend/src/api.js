const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

export const register = async (email, password) => {
  const response = await fetch(`${API_URL}/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  return await response.json();
};

export const login = async (email, password) => {
  const response = await fetch(`${API_URL}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  return await response.json();
};

export const getProfile = async (token) => {
  const response = await fetch(`${API_URL}/profile`, {
    headers: { 'Authorization': token }
  });
  return await response.json();
};

// Add other API calls similarly