import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import uploadImages from "./setImage";

function App() {
 
  return (
    <>
      <div>
        <h1> Upload Page</h1>
        <h2>Add Image:</h2>
        <uploadImages></uploadImages>
        <div className="input-container">
          <label htmlFor="userInput">Enter your text: </label>
          <input id="userInput" type="text"></input>
        </div>
        <button className="uploadButton"> Upload</button>
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
