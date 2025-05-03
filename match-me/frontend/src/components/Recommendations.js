import React, { useState, useEffect } from 'react';
import { useAuth } from '../auth';
import * as api from '../api';

function Recommendations() {
  const [recommendations, setRecommendations] = useState([]);
  const { token } = useAuth();

  useEffect(() => {
    const fetchRecommendations = async () => {
      const data = await api.getRecommendations(token);
      setRecommendations(data);
    };
    fetchRecommendations();
  }, [token]);

  const handleConnect = async (id) => {
    await api.sendConnectionRequest(id, token);
    setRecommendations(recommendations.filter(rec => rec.id !== id));
  };

  const handleDismiss = (id) => {
    setRecommendations(recommendations.filter(rec => rec.id !== id));
  };

  return (
    <div className="recommendations">
      <h3>Recommended Matches</h3>
      {recommendations.length === 0 ? (
        <p>No recommendations found. Update your profile to get better matches.</p>
      ) : (
        <div className="recommendation-list">
          {recommendations.map(rec => (
            <div key={rec.id} className="recommendation-card">
              <img 
                src={rec.profile_picture || '/default-profile.png'} 
                alt={rec.username}
              />
              <h4>{rec.username}</h4>
              <p>{rec.about_me?.substring(0, 100)}...</p>
              <div className="actions">
                <button 
                  onClick={() => handleConnect(rec.id)}
                  className="primary-btn"
                >
                  Connect
                </button>
                <button 
                  onClick={() => handleDismiss(rec.id)}
                  className="secondary-btn"
                >
                  Dismiss
                </button>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default Recommendations;