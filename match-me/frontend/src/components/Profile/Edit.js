import React, { useState, useEffect } from 'react';
import { useAuth } from '../../auth';
import * as api from '../../api';

function ProfileEdit() {
  const { user, token } = useAuth();
  const [profile, setProfile] = useState({
    username: '',
    about_me: '',
    location: ''
  });

  useEffect(() => {
    const fetchProfile = async () => {
      const data = await api.getProfile(token);
      setProfile(data);
    };
    fetchProfile();
  }, [token]);

  const handleChange = (e) => {
    setProfile({...profile, [e.target.name]: e.target.value});
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    await api.updateProfile(profile, token);
    alert('Profile updated!');
  };

  return (
    <div className="profile-edit">
      <h2>Edit Profile</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Username:
          <input
            type="text"
            name="username"
            value={profile.username}
            onChange={handleChange}
            required
          />
        </label>
        <label>
          About Me:
          <textarea
            name="about_me"
            value={profile.about_me}
            onChange={handleChange}
          />
        </label>
        <label>
          Location:
          <input
            type="text"
            name="location"
            value={profile.location}
            onChange={handleChange}
          />
        </label>
        <button type="submit" className="primary-btn">Save</button>
      </form>
    </div>
  );
}

export default ProfileEdit;