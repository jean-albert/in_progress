import React from 'react';
import { Link } from 'react-router-dom';
import Login from '../components/Auth/Login';
import Register from '../components/Auth/Register';
import { useAuth } from '../auth';

function Home() {
  const { user } = useAuth();

  return (
    <div className="home">
      <header>
        <h1>Match-Me</h1>
        <p>Find your perfect match based on shared interests</p>
      </header>
      
      {user ? (
        <div className="auth-buttons">
          <Link to="/dashboard" className="primary-btn">Go to Dashboard</Link>
        </div>
      ) : (
        <div className="auth-options">
          <Login />
          <Register />
        </div>
      )}
    </div>
  );
}

export default Home;