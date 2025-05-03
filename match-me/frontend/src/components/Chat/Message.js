import React from 'react';
import { useAuth } from '../../auth';

function Message({ message }) {
  const { user } = useAuth();
  const isMe = message.sender_id === user.id;

  return (
    <div className={`message ${isMe ? 'sent' : 'received'}`}>
      <p>{message.content}</p>
      <span className="timestamp">
        {new Date(message.created_at).toLocaleTimeString()}
      </span>
    </div>
  );
}

export default Message;