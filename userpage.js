import React, { useState } from 'react';
import './styles.css';

// we can see what the backend db is storing, but we should pass the loggedIn user ids, profile id, etc. in here 
const ProfilePage = ({ loggedInUserId, profileId, profileName, profileBio, initialFollowerCount, initialFollowingCount }) => {
    const [followerCount, setFollowerCount] = useState(initialFollowerCount);
    const [followingCount, setFollowingCount] = useState(initialFollowingCount);
    const [isFriend, setIsFriend] = useState(false);
    const [isSelf, setIsSelf] = useState(loggedInUserId===profileId); // call setIsSelf when user logs in

    const handleAddFriend = () => {
        setIsFriend(!isFriend);
        this.classList.add("hideButton"); // not display hide button
    };

    return (
        <div>
            <div className="top-bar">
                <img src="SSlogo.png" className="logo" alt="Logo" />
            </div>
            <div className="profile-container">
                <div className="profile-header">
                    <div className="profile-picture"></div>
                    <div className="profile-info">
                        <h1> {profileName} </h1>
                        <p> {profileBio} </p>
                        <div className="counts">
                            <span id="follower-count">Followers: {followerCount}</span> |
                            <span id="following-count">Following: {followingCount}</span>
                        </div>
                        {loggedInUserId !== profileId && (
                            <button id="friend-button" onClick={handleAddFriend}>
                                {isFriend ? 'Unfriend' : 'Add Friend'}
                            </button>
                        )}
                    </div>
                </div>
                <div className="gallery">
                    <div className="photo add-photo">
                        <a href="upload" class="add-button">+</a>
                    </div>
                    <div class="photo"></div>
                    <div class="photo"></div>
                    <div class="photo"></div>
                    <div class="photo"></div>
                    <div class="photo"></div>
                </div>
            </div>
        </div>
    );
};

export default ProfilePage;
