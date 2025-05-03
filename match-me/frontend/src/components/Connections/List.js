import React, { useEffect, useState } from 'react';
import { useAuth } from '../../auth';
import * as api from '../../api';

function ConnectionList() {
  const [connections, setConnections] = useState([]);
  const { token } = useAuth();

  useEffect(() => {
    const fetchConnections = async () => {
      const data = await api.getConnections(token);
      setConnections(data);
    };
    fetchConnections();
  }, [token]);

  return (
    <div className="connections">
      <h3>Your Connections</h3>
      {connections.length === 0 ? (
        <p>No connections yet</p>
      ) : (
        <ul>
          {connections.map(conn => (
            <li key={conn.id}>
              <span>{conn.username}</span>
              <button className="secondary-btn">Chat</button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default ConnectionList;