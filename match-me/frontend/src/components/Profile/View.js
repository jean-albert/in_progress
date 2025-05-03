import React from 'react';
import { useAuth } from '../../auth';

function ProfileView() {
  const { user } = useAuth();

  return (
    <div className="profile-view">
      <h2>Your Profile</h2>
      <div className="profile-info">
        <img 
          src={user.profile_picture || '/default-profile.png'} 
          alt={user.username}
          className="profile-pic"
        />
        <h3>{user.username}</h3>
        <p>{user.about_me}</p>
        <p>Location: {user.location}</p>
      </div>
    </div>
  );
}

export default ProfileView;