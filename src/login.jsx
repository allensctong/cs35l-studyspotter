import React from 'react'
import ReactDOM from 'react-dom/client'
import studySpotLogo from './assets/study-spotter.jpg'
import './login.css'

function Login() {
	async function uwu(FormData) {
		console.log("hi");
		let response = await fetch("http://localhost:8080/api/user", {
			method: 'POST',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({
				user: FormData.get('Username'),
				pass: FormData.get('Password')
			}),
		}
		);
		let userInfo = await response.json();
		console.log(userInfo);
	}
	return (
		<>
			<div>
				<img src={studySpotLogo} className="logo" alt="Study Spotter Logo"/>
			</div>
			<div>
				<form onSubmit={uwu}>
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
