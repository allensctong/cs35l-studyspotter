import React, { useState } from 'react';
import './styles.css';

const ProfilePage = () => {
    const [followerCount, setFollowerCount] = useState(0);
    const [followingCount, setFollowingCount] = useState(0);
    const [isFriend, setIsFriend] = useState(false);
    const [isSelf, setIsSelf] = useState(false); // call setIsSelf when user logs in

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
                        <h1>Profile Name</h1>
                        <p>Bio</p>
                        <div className="counts">
                            <span id="follower-count">Followers: {followerCount}</span> |
                            <span id="following-count">Following: {followingCount}</span>
                        </div>
                        <button id="friend-button" onClick={handleAddFriend}>
                            {isFriend ? 'Unfriend' : 'Add Friend'}
                        </button>
                    </div>
                </div>
                <div className="gallery">
                    <div className="photo add-photo">
                        <button className="add-button">+</button>
                    </div>
                    <div className="photo"></div>
                    <div className="photo"></div>
                    <div className="photo"></div>
                    <div className="photo"></div>
                </div>
            </div>
        </div>
    );
};

export default ProfilePage;
