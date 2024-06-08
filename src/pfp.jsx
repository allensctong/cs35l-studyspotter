import { useState, useEffect } from 'react'
import './pfp.css'
import { useRef} from 'react'

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
  const hiddenFileInput=useRef(null);

  const handleClick = (Event)=>{
    hiddenFileInput.current.click();
  }

  useEffect(() => {
    if (!getCookieValue('Username')) {
      window.location.href = 'login';
      return;
    }
  }, []);


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
  }

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
      let response = await fetch('http://localhost:8080/api/user/' + getCookieValue('Username') + '/pfp', {
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
  }
 
  return (
    <>
      <div>
        <h1> Upload Profile Picture</h1>
        <h2>Add Image:</h2>
        {imageURL ? ( <img src={imageURL}  alt="Selected" />)
        :
        
        (<button className="add-button" onClick={handleClick}>Click to Add Your Image</button>)}
        < input type="file" accept=".jpg,.jpeg,.png" onChange={handleImageChange} ref={hiddenFileInput} style={{display: 'none'}}/>
        {error && <p className="error-message">{error}</p>}
        
        <div>
        <button className="setProfileButton" onClick={handleUpload}> Set as Profile Picture</button>
        </div>
      </div>
    </>
  );
}

export default Upload
