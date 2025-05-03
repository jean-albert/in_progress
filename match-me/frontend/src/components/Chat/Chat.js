import React, { useState, useEffect } from 'react';
import Message from './Message';
import { useAuth } from '../../auth';
import * as api from '../../api';

function Chat({ connectionId }) {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');
  const { token } = useAuth();

  useEffect(() => {
    const fetchMessages = async () => {
      const data = await api.getMessages(connectionId, token);
      setMessages(data);
    };
    fetchMessages();
  }, [connectionId, token]);

  const handleSend = async () => {
    if (!newMessage.trim()) return;
    await api.sendMessage(connectionId, newMessage, token);
    setMessages([...messages, {
      id: Date.now(),
      content: newMessage,
      sender_id: 'me',
      created_at: new Date().toISOString()
    }]);
    setNewMessage('');
  };

  return (
    <div className="chat-container">
      <div className="messages">
        {messages.map(msg => (
          <Message key={msg.id} message={msg} />
        ))}
      </div>
      <div className="message-input">
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          onKeyPress={(e) => e.key === 'Enter' && handleSend()}
          placeholder="Type a message..."
        />
        <button onClick={handleSend} className="primary-btn">Send</button>
      </div>
    </div>
  );
}

export default Chat;