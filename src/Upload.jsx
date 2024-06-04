import { useState } from 'react'
import './Upload.css'

function getCookieValue(name) 
    {
      const regex = new RegExp(`(^| )${name}=([^;]+)`)
      const match = document.cookie.match(regex)
      if (match) {
        return match[2]
      }
   }

function Upload() {
  const [userInput, setUserInput]=useState('');
  const[selectedImage, setSelectedImage]=useState(null);
  const [imageURL, setImageURL] = useState(null);
  const [error, setError] = useState('');
  const [profilePicture, setProfileImage]=useState(null);


  const handleInputChange= (Event)=> {
    setUserInput(Event.target.value);
  };

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
    if(selectedImage==null){
      setError("Please upload an image before submitting.");
    }else{
      console.log("User Input:", userInput);
      console.log("Selected Image:", selectedImage);
      console.log("Username:", document.cookie);

      const formData = new FormData();
      formData.append('username', getCookieValue('Username'));
      formData.append('caption', userInput);
      formData.append('image', selectedImage);
      let response = await fetch("http://localhost:8080/api/post", {
        method: 'POST',
        credentials: 'include',
        body: formData,
      });
      if(await response.status !== 200) {
        alert("Upload failed!");
        return;
      }
      
      window.location.href = "../uploadSuccess.html";
    }
 
    
    


  };
  const handleProfile=()=>{
    if(selectedImage==null){
      setError("Please upload an image before submitting.");
    }
    setProfileImage(selectedImage);
    console.log("Profile Picture:", profilePicture);
  };
 
 

 
  return (
    <>
      <div>
        <h1> Upload Page</h1>
        <h2>Add Image:</h2>
        
        <input type="file" accept=".jpg,.jpeg,.png" onChange={handleImageChange} />
        {error && <p className="error-message">{error}</p>} <br />
        {imageURL && <img src={imageURL} alt="Selected" className="uploaded-image" />}
        <div className="input-container">
          <label htmlFor="userInput">Caption (optional): </label>
          <input id="userInput" type="text" value={userInput} onChange={handleInputChange}></input>
        </div>
        <button className="uploadButton" onClick={handleUpload}> Upload</button>
        <button className="setProfileButton" onClick={handleProfile}> Set as Profile Picture</button>
      </div>

    </>
  )
}

export default Upload
