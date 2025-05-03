import React from 'react';
import { useAuth } from '../auth';
import Recommendations from '../components/Recommendations';
import ConnectionList from '../components/Connections/List';
import ConnectionRequests from '../components/Connections/Requests';
import ProfileView from '../components/Profile/View';

function Dashboard() {
  const { user, logout } = useAuth();

  if (!user) {
    return <div>Please login to view dashboard</div>;
  }

  return (
    <div className="dashboard">
      <header>
        <h2>Welcome, {user.username}!</h2>
        <button onClick={logout} className="secondary-btn">Logout</button>
      </header>

      <div className="dashboard-content">
        <div className="profile-section">
          <ProfileView />
        </div>
        
        <div className="recommendations-section">
          <Recommendations />
        </div>
        
        <div className="connections-section">
          <ConnectionList />
          <ConnectionRequests />
        </div>
      </div>
    </div>
  );
}

export default Dashboard;