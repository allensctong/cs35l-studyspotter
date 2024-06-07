import React, { useState, useEffect } from 'react';
import './style_specific_mint.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHeart, faComment, faUser, faSearch } from '@fortawesome/free-solid-svg-icons';
import { faHeart as faHeartEmpty } from '@fortawesome/free-regular-svg-icons';

const App = () => {
  const [posts, setPosts] = useState([]);
  const [users, setUsers] = useState([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchResults, setSearchResults] = useState([]);
  const [isDropdownVisible, setIsDropdownVisible] = useState(false);

  useEffect(() => {
    fetch('http://localhost:8080/api/post', { credentials: 'include' })
      .then(response => response.json())
      .then(data => setPosts(data))
      .catch(error => console.error('Error fetching posts:', error));
  }, []);

  useEffect(() => {
    if (searchQuery.length > 0) {
      fetch('http://localhost:8080/api/user/search/' + searchQuery, {
        credentials: 'include',
      })
        .then(response => response.json())
        .then(data => {
          setUsers(data);
          setSearchResults(data);
          setIsDropdownVisible(true);
        })
        .catch(error => console.error('Error fetching users:', error));
    } else {
      setSearchResults([]);
      setIsDropdownVisible(false);
    }
  }, [searchQuery]);

  const handleSearch = (e) => {
    setSearchQuery(e.target.value);
  };

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    if (searchQuery.length > 0) {
      const results = users.filter(user => user.username.toLowerCase().includes(searchQuery.toLowerCase()));
      setSearchResults(results);
      setIsDropdownVisible(true);
    }
  };

  const likePost = (index) => {
    const newPosts = [...posts];
    if (newPosts[index].liked) {
      newPosts[index].likes -= 1;
      newPosts[index].liked = false;
    } else {
      newPosts[index].likes += 1;
      newPosts[index].liked = true;
    }
    setPosts(newPosts);
  };

  const commentPost = (index) => {
    const comment = prompt('Enter your comment:');
    const username = posts[index].username;
    fetch('http://localhost:8080/api/post/' + posts[index].post_id + '/comment', {
      credentials: 'include',
    });

    if (comment) {
      const newPosts = [...posts];
      newPosts[index].comments.push({ username, comment });
      setPosts(newPosts);
    }
  };

  const deleteAllCookies = () => {
    document.cookie.split(';').forEach(cookie => {
      const eqPos = cookie.indexOf('=');
      const name = eqPos > -1 ? cookie.substring(0, eqPos) : cookie;
      document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:00 GMT';
    });
    window.location.href = '/';
  };

  return (
    <div>
      <div className="header">
        <img src="/SSlogo.png" alt="Study Spotter Logo" className="logo" onClick={() => window.location.reload()} />
        <form method="get" className="search-bar" onSubmit={handleSearchSubmit}>
          <input
            type="text"
            placeholder="Search..."
            value={searchQuery}
            onChange={handleSearch}
          />
          <button type="submit">
            <FontAwesomeIcon icon={faSearch} />
          </button>
          {isDropdownVisible && searchResults.length > 0 && (
            <div className="search-results">
              {searchResults.map((user, index) => (
                <a href={`http://localhost:5173/user?u=${user.username}`} key={index} style={{ textDecoration: 'none' }}>
                  <div className="search-result-item">
                    <img src={user.pfp} alt={`${user.username}'s profile`} className="profile-picture" />
                    <span className="username">{user.username}</span>
                  </div>
                </a>
              ))}
            </div>
          )}
        </form>
        <div className="right-buttons">
          <button onClick={() => window.location.href = 'http://localhost:5173/user'}>
            <FontAwesomeIcon icon={faUser} /> Profile
          </button>
          <button onClick={deleteAllCookies}>
            Logout
          </button>
        </div>
      </div>
      <div className="content">
        {posts.length === 0 ? (
          <div style={{ fontSize: '24px'}}>No posts yet.</div>
        ) : (
          posts.map((post, index) => (
            <div className="post" key={index}>
              <div className="uploader">{post.username}</div>
              <img src={post.image_src} alt="Post" />
              <div className="buttons">
                <button onClick={() => likePost(index)}>
                  <FontAwesomeIcon icon={post.liked ? faHeart : faHeartEmpty} /> Like <span className="like-count">{post.likes}</span>
                </button>
                <button onClick={() => commentPost(index)}>
                  <FontAwesomeIcon icon={faComment} /> Comment
                </button>
              </div>
              <div className="caption">{post.caption}</div>
              <div className="comments">
                {post.comments.map((comment, idx) => (
                  <div className="comment" key={idx}><strong>{comment.username}:</strong> {comment.comment}</div>
                ))}
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default App;
