import { useState, useEffect } from 'react'
import { useRef} from 'react'
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
  const[postWidth, setPostWidth]=useState(0);
  const[postHeight, setPostHeight]=useState(0);

  useEffect(() => {
    if (!getCookieValue('Username')) {
      window.location.href = 'login';
      return;
    }
  }, []);

  const imgDimension={width: 1024/2, height:768/2 };
 
  const hiddenFileInput=useRef(null);

  const handleClick=(Event)=>{
    hiddenFileInput.current.click();
  }
  const handleInputChange= (Event)=> {
    setUserInput(Event.target.value);
  };

  const handleImageChange = (Event)=> {
    const file = Event.target.files[0];
    if (file) {
      const fileExtension = file.name.split('.').pop().toLowerCase();
      if (['jpg', 'jpeg', 'png'].includes(fileExtension)) {
        setImageURL(URL.createObjectURL(file));
        setSelectedImage(file);
        setError('');
        const img=new Image();
        img.onload = function(){
          var width=img.width;
          var height=img.height;
          if (width > height) {
            if (width > imgDimension.width) {
            height *= imgDimension.width / width;
            width = imgDimension.width;
            }
        } 
        else {
          if (height > imgDimension.height) {
            width *= imgDimension.height / height;
            height = imgDimension.height;
            }
        }

        setPostHeight(height);
        setPostWidth(width);
        
      }
      img.src=URL.createObjectURL(file);
      console.log(postHeight);
      console.log(postWidth);
      console.log(imageURL);
      console.log(selectedImage);
        //get dimensions, edit if nesecary, set postWidth and postHeight
        //do your feedback forms
        //do the weekly update
        

        /*var img=new Image();
        img.src=URL.createObjectURL(file);
        resizeImage(img,file);*/
       
      } 
    
      else {
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
   
    }

  

  const uploadInfo=new FormData();
  uploadInfo.append('username', getCookieValue('Username'));
  uploadInfo.append('caption', userInput);
  uploadInfo.append('image', selectedImage);
  uploadInfo.append('imageWidth', postWidth);
  uploadInfo.append('imageHeight', postHeight);
  

  let response = await fetch("http://localhost:8080/api/post", {
    method: 'POST',
    credentials: 'include',
    body: uploadInfo,
  });
  if(await response.status !== 200) {
    alert("Upload failed!");
    return;
  }
  window.location.href = "/uploaded";
  };

  const handleProfile=()=>{
    if(selectedImage==null){
      setError("Please upload an image before submitting.");
    }
    setProfileImage(selectedImage);
    console.log("Profile Picture:", profilePicture);
  };

  const handleCancel=()=>{
    window.location.href = "../user.html";
  }

 
 

 
  return (
    <>
      <div>
        <h1> Create Your Post!</h1>
        <h2>Show off how YOU study:</h2>

        {imageURL ? ( <img src={imageURL} width={postWidth} height={postHeight} alt="Selected" className="image-preview-container" />)
        :
        
       (<button className="add-button" onClick={handleClick}>Click to Add Your Image</button>)}
        < input type="file" accept=".jpg,.jpeg,.png" onChange={handleImageChange} ref={hiddenFileInput} style={{display: 'none'}}/>
        {error && <p className="error-message">{error}</p>}
        
        
        <div className="input-container">
        
          <label htmlFor="userInput">Caption (optional): </label>
          <input id="userInput" type="text" value={userInput} onChange={handleInputChange}></input>
        </div>
        <div>
        <button className="uploadButton" onClick={handleUpload}> Upload</button>
        <button className="uploadButton" onClick={handleClick}> Select A Different Image</button>
        <button className="uploadButton" onClick={handleCancel}> Cancel</button>
        </div>
       

      </div>
    </>
  )
}

export default Upload
