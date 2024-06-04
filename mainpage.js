import React, { useState, useEffect } from 'react';
import './style_specific_mint.css'; // Ensure this path is correct
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHeart, faComment, faShareSquare, faUser, faSearch } from '@fortawesome/free-solid-svg-icons';

const App = () => {
 const [posts, setPosts] = useState([]);

 useEffect(() => {
   fetch('/posts.json')
     .then(response => response.json())
     .then(data => setPosts(data))
     .catch(error => console.error('Error fetching posts:', error));
 }, []);

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
   const username = 'Guest';
   if (comment) {
     const newPosts = [...posts];
     newPosts[index].comments.push({ username, comment });
     setPosts(newPosts);
   }
 };

 return (
   <div>
     <div className="header">
       <img src="/SSlogo.png" alt="Study Spotter Logo" className="logo" />
       <form action="search.html" method="get" className="search-bar">
         <input type="text" placeholder="Search..." name="query" />
         <button type="submit">
           <FontAwesomeIcon icon={faSearch} />
         </button>
       </form>
       <div className="right-buttons">
         <button onClick={() => window.location.href = 'profile.html'}>
           <FontAwesomeIcon icon={faUser} /> Profile
         </button>
       </div>
     </div>
     <div className="content">
       {posts.map((post, index) => (
         <div className="post" key={index}>
           <div className="uploader">{post.uploader}</div>
           <img src={post.imgSrc} alt="Post" />
           <div className="buttons">
             <button onClick={() => likePost(index)}>
               <FontAwesomeIcon icon={post.liked ? faHeart : "fa-regular fa-heart"} /> Like <span className="like-count">{post.likes}</span>
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
       ))}
     </div>
   </div>
 );
};

export default App;
