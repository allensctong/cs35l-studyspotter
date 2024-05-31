import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'


function App() {
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

  const handleUpload = () => {
    if(selectedImage==null){
      setError("Please upload an image before submitting.");
    }
    console.log("User Input:", userInput);
    console.log("Selected Image:", selectedImage);
 
    
    


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
        {error && <p className="error-message">{error}</p>}
        {imageURL && <img src={imageURL} alt="Selected" className="uploaded-image" />}
        <div className="input-container">
          <label htmlFor="userInput">Enter your text: </label>
          <input id="userInput" type="text" value={userInput} onChange={handleInputChange}></input>
        </div>
        <button className="uploadButton" onClick={handleUpload}> Upload</button>
        <button className="setProfileButton" onClick={handleProfile}> Set as Profile Picture</button>
      </div>

    </>

   /* <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.jsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </> */
  )
}

export default App
