import React, { useEffect, useState } from 'react';
import { useAuth } from '../../auth';
import * as api from '../../api';

function ConnectionRequests() {
  const [requests, setRequests] = useState([]);
  const { token } = useAuth();

  useEffect(() => {
    const fetchRequests = async () => {
      const data = await api.getConnectionRequests(token);
      setRequests(data);
    };
    fetchRequests();
  }, [token]);

  const handleResponse = async (id, accept) => {
    await api.respondToRequest(id, accept, token);
    setRequests(requests.filter(req => req.id !== id));
  };

  return (
    <div className="requests">
      <h3>Connection Requests</h3>
      {requests.length === 0 ? (
        <p>No pending requests</p>
      ) : (
        <ul>
          {requests.map(req => (
            <li key={req.id}>
              <span>{req.username}</span>
              <button 
                onClick={() => handleResponse(req.id, true)}
                className="primary-btn"
              >
                Accept
              </button>
              <button 
                onClick={() => handleResponse(req.id, false)}
                className="secondary-btn"
              >
                Decline
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default ConnectionRequests;