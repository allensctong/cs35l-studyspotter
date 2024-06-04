import React from 'react'
import {useState} from 'react'
import ReactDOM from 'react-dom/client'
import studySpotLogo from './assets/study-spotter.jpg'
import './login.css'

function Login() {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	function handleChangeUsername(Event) {
		setUsername(Event.target.value);
	}

	function handleChangePassword(Event) {
		setPassword(Event.target.value);
	}
	async function handleLogin() {
		let response = await fetch("http://localhost:8080/api/login", {
			method: 'POST',
			headers: {
				'content-type': 'application/json',
			},
			credentials: 'include',
			body: JSON.stringify({
				'username': username,
				'password': password
			}),
		}
		);
		if(await response.status !== 200) {
			alert("Login failed!");
			return;
    		}
		window.location.href = 'http://localhost:5173/main';
	}

	async function handleSignup() {
		let response = await fetch("http://localhost:8080/api/signup", {
			method: 'POST',
			headers: {
				'content-type': 'application/json',
			},
			credentials: 'include',
			body: JSON.stringify({
				'username': username,
				'password': password
			}),
		}
		);

		if(await response.status !== 201) {
			alert("OHNOES");
			return;
    } else {
			alert("poggers");
			return;
		}
	}

	const onSubmit = (e) => {
		e.preventDefault();
    console.log("refresh prevented");
  };
	return (
		<>
			<div>
				<img src={studySpotLogo} className="logo" alt="Study Spotter Logo"/>
			</div>
			<div>
				<form onSubmit={onSubmit}>
					Username: <input name="Username" value={username} onChange={handleChangeUsername} /><br/>
					Password: <input name="Password" value={password} onChange={handleChangePassword} /><br />
					<button type="Submit" onClick={handleLogin}>Login</button><br />
					<button type="Submit" onClick={handleSignup}>Create Account</button>
				</form>
			</div>
		</>
	)
}

export default Login
