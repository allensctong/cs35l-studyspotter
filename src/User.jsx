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
function ProfilePage () {
    const [followerCount, setFollowerCount] = useState(0);
    const [followingCount, setFollowingCount] = useState(0);
    const [isFriend, setIsFriend] = useState(false);
    const [isUser, setIsUser] = useState(true); 
    const [profileName, setProfileName] = useState('');
    const [profileBio, setProfileBio] = useState('');
    const [pfpSrc, setPfpSrc] = useState("http://localhost:8080/assets/default-pfp.jpg");
    const [postSrcs, setPostSrcs] = useState([]);
    const [isEditingBio, setIsEditingBio] = useState(false);
    const [newBio, setNewBio] = useState('');
    const params = new URL(window.location.href).searchParams;
    
    //call the first time the page is rendered
    useEffect(() => {
      let ignore = false;
      if (!ignore) { getUserInfo(); }
      return () => { ignore = true; };
    }, []);
    
    async function getUserInfo() {
      //if input is default get username from cookie
      const curUser = getCookieValue('Username');
      var username = params.get('u');
      if (!username) {
        setProfileName(curUser);
        username = curUser;
      } else {
        setProfileName(username);
      }
      
      //fetch profile info
      let response = await fetch('http://localhost:8080/api/user/' + username, {
        credentials: 'include',
      });
      if (await response.status == 401) {
        window.location.href = 'login';
        return;
      }
      if (await response.status !== 200) {
        //TODO handle error
        alert('ERROR');
        return;
      }

      setIsUser(username === curUser);
      
      response = await response.json();
      setProfileBio(response.bio);
      setPfpSrc(response.pfp);
      setPostSrcs(response.posts);
      setIsUser(username === curUser);
      setFollowerCount(response.followers);
      setFollowingCount(response.following);
    }

    const handleAddFriend = () => {
        setIsFriend(!isFriend);
        this.classList.add("hideButton"); // not display hide button
    };

    const handleEditBio = () => {
        setIsEditingBio(true);
        setNewBio(profileBio);
    };

    const handleBioChange = (bio) => {
        setNewBio(bio.target.value);
    };

    const handleSubmitBio = async () => {
        let response = await fetch('http://localhost:8080/api/user/' + profileName + '/bio', {
            method: 'PUT',
            credentials: 'include',
            headers: {
                'content-type': 'application/json'
            },
            body: JSON.stringify({ 'bio': newBio }),
        });
        
        if (await response.status === 200) {
            setProfileBio(newBio);
            setIsEditingBio(false);
        } else {
            alert('ERROR');
        }
    };

    const handleCloseModal = () => {
        setIsEditingBio(false);
    };

    return (
        <div>
            <div className="top-bar">
                <a href="/main" className="logo-link">
                    <img src="SSlogo.png" className="logo" alt="Logo" />
                </a>
            </div>
            <div className="profile-container">
                <div className="profile-header">
                    <div className="profile-picture" src={pfpSrc}>
                    	<img className="profile-picture" src={pfpSrc} />
                        {isUser && (
                            <a href="/pfp" className="edit-profile-link">
                                <img src="../pencil.png" alt="Edit Profile" className="pencil-icon" />
                            </a>
                        )}
                    </div>
                    <div className="profile-info">
                        <h1> {profileName} </h1>
                        <p> {profileBio} </p>
                        <div className="counts">
                            <span id="follower-count">Followers: {followerCount}</span> |
                            <span id="following-count"> Following: {followingCount}</span>
                        </div>
                        {!isUser && (
                            <button id="friend-button" 
                                className={isFriend ? 'unadd-friend' : 'add-friend'}
                                onClick={handleAddFriend}>
                                {isFriend ? 'Unadd Friend' : 'Add Friend'}
                            </button>
                        )}
                        {isUser && (
                            <button onClick={handleEditBio}>Edit Bio</button>
                        )}
                    </div>
                </div>
                <div className="gallery">
                    {isUser && (
                        <div className="photo add-photo">
                            <a href="upload" className="add-button">+</a>
                        </div>
                    )}
                    {postSrcs.map((element, index) => (
                        <img className="photo" src={element} key={index}/>
                    ))}
                </div>
            </div>
            {isEditingBio && (
                <div className="modal">
                    <div className="modal-content">
                        <span className="close-button" onClick={handleCloseModal}>&times;</span>
                        <h2>Edit Bio</h2>
                        <textarea 
                            value={newBio} 
                            onChange={handleBioChange} 
                            rows="4" 
                            cols="50"
                        /> 
                        <button onClick={handleSubmitBio}>Submit</button>
                    </div>
                </div>
            )}
        </div>
    )
}

export default ProfilePage
