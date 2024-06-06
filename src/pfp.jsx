import { useState } from 'react'
import './pfp.css'

function getCookieValue(name) 
    {
      const regex = new RegExp(`(^| )${name}=([^;]+)`)
      const match = document.cookie.match(regex)
      if (match) {
        return match[2]
      }
   }

function Upload() {
  const[selectedImage, setSelectedImage]=useState(null);
  const [imageURL, setImageURL] = useState(null);
  const [error, setError] = useState('');
  const [profilePicture, setProfileImage]=useState(null);

  const handleImageChange = (Event)=> {
    const file = Event.target.files[0];
    if (file) {
      const fileExtension = file.name.split('.').pop().toLowerCase();
      if (['jpg', 'jpeg', 'png'].includes(fileExtension)) {
        setSelectedImage(file);
        setImageURL(URL.createObjectURL(file));
        setError('');
      } else {
        setSelectedImage(null);
        setImageURL(null);
        setError('Invalid file type. Please upload an image file (.jpg, .jpeg, .png).');
      }
    }
  };

  const handleUpload = async () => {
    if (selectedImage == null) {
      setError("Please upload an image before submitting.");
    } else {
      setProfileImage(selectedImage);
      console.log("Profile Picture:", profilePicture);
      console.log("Username:", document.cookie);

      const formData = new FormData();
      formData.append('username', getCookieValue('Username'));
      formData.append('image', selectedImage);
      // Need to edit fetch request
      let response = await fetch('http://localhost:8080/api/user/' + getCookieValue('Username'), {
        method: 'PUT',
        credentials: 'include',
        body: formData,
      });
      if(await response.status !== 200) {
        alert("Upload failed!");
        return;
      }
      window.location.href = "../uploaded";
    }
  };
 
  return (
    <>
      <div>
        <h1> Upload Profile Picture</h1>
        <h2>Add Image:</h2>
        
        <input type="file" accept=".jpg,.jpeg,.png" onChange={handleImageChange} />
        {error && <p className="error-message">{error}</p>} <br />
        {imageURL && <img src={imageURL} alt="Selected" className="uploaded-image" />}
        {imageURL && <br />}
        <button className="setProfileButton" onClick={handleUpload}> Set as Profile Picture</button>
      </div>
    </>
  )
}

export default Upload