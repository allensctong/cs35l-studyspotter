import React from 'react';

const ProfileComponent = () => {
  return (
     <div className="profile-header">
      <div className="profile-picture">
        {/* <img src="profile-picture-url.jpg" alt="Profile" /> */}
      </div>
      <div className="profile-info">
        <h1>Profile Name</h1>
        <p>Bio</p>
        <div className="counts">
          <span id="follower-count">Followers: 0</span> | 
          <span id="following-count">Following: 0</span>
        </div>
        <button id="friend-button">Add Friend</button>
      </div>
    </div>
  );
};

export default ProfileComponent;
