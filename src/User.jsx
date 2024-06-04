import React, { useState, useEffect } from 'react';
import './User.css';

function getCookieValue(name) 
    {
      const regex = new RegExp(`(^| )${name}=([^;]+)`)
      const match = document.cookie.match(regex)
      if (match) {
        return match[2]
      }
   }

// we can see what the backend db is storing, but we should pass the loggedIn user ids, profile id, etc. in here 
function ProfilePage ({username = ''}) {
    const [followerCount, setFollowerCount] = useState(0);
    const [followingCount, setFollowingCount] = useState(0);
    const [isFriend, setIsFriend] = useState(false);
    const [isUser, setIsUser] = useState(true); 
    const [profileName, setProfileName] = useState('');
    const [profileBio, setProfileBio] = useState('');
    
    //call the first time the page is rendered
    useEffect(() => {
      let ignore = false;
      if (!ignore) { getUserInfo(); }
      return () => { ignore = true; };
    }, []);
    
    async function getUserInfo() {
      //if input is default get username from cookie
      const curUser = getCookieValue('Username');
      if (username === "") {
        setProfileName(curUser);
	username = curUser;
      } else {
        setProfileName(username);
      }
      setIsUser(profileName === curUser);
      
      //fetch profile info
      let response = await fetch('http://localhost:8080/api/user/' + username);
      if (await response.status !== 200) {
        //TODO handle error
        alert('ERROR');
        return;
      }
      
      response = await response.json();
      setProfileBio(response.bio);
      setFollowerCount(response.followers);
      setFollowingCount(response.following);

    }

    const handleAddFriend = () => {
        setIsFriend(!isFriend);
        this.classList.add("hideButton"); // not display hide button
    };

    return (
        <div>
            <div className="top-bar">
                <a href="/main" class="logo-link">
                    <img src="SSlogo.png" className="logo" alt="Logo" />
                </a>
            </div>
            <div className="profile-container">
                <div className="profile-header">
                    <div className="profile-picture"></div>
                    <div className="profile-info">
                        <h1> {profileName} </h1>
                        <p> {profileBio} </p>
                        <div className="counts">
                            <span id="follower-count">Followers: {followerCount}</span> |
                            <span id="following-count"> Following: {followingCount}</span>
                        </div>
                        {isUser && (
                            <button id="friend-button" onClick={handleAddFriend}>
                                {isFriend ? 'Unfriend' : 'Add Friend'}
                            </button>
                        )}
                    </div>
                </div>
                <div className="gallery">
                    {isUser && (
                        <div className="photo add-photo">
                            <a href="upload" className="add-button">+</a>
                        </div>
                    )}
                    <div className="photo"></div>
                    <div className="photo"></div>
                    <div className="photo"></div>
                    <div className="photo"></div>
                    <div className="photo"></div>
                </div>
            </div>
        </div>
    )
}

export default ProfilePage
