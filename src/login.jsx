import React from 'react'
import ReactDOM from 'react-dom/client'
import studySpotLogo from './assets/study-spotter.jpg'
import './login.css'

function Login() {
	return (
		<>
			<div>
				<img src={studySpotLogo} className="logo" alt="Study Spotter Logo"/>
			</div>
			<div>
				<form>
					<label> Username: </label><br/>			
					<input name="Username" /><br/>
					<label> Password: </label><br/>			
					<input name="Password" /><br />
					<button type="submit">Submit</button>
				</form>
			</div>
		</>
	)
}

export default Login

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Login />
  </React.StrictMode>,
)
