import React, {useState, useEffect} from 'react';

const ProfileComponent = () => {
  const [profile, setProfile] = useState({
    profilePicture: './src/assets/Default_pfp.svg',
    profileName: '',
    bio: '',
    followerCount: 0,
    followingCount: 0
  });

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        // Fetch the current user
        const profileResponse = await fetch('http://localhost:8080/api/user/current', {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
          credentials: 'include',
        });

        if (profileResponse.status !== 200) {
          console.error('Failed to fetch profile data');
          return;
        }

        const profileData = await profileResponse.json();
        setProfile(profileData);
      } catch (error) {
        console.error('Error fetching profile data:', error);
      }
    };

    fetchProfile();
  }, []);

  
  return (
     <div className="profile">
      <div className="profile-picture">
        <img src={profile.profilePicture} alt="Profile" />
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