import React from 'react'
import {useState, useEffect} from 'react'
import ReactDOM from 'react-dom/client'
import studySpotLogo from './assets/study-spotter.jpg'
import './login.css'

function getCookieValue(name) 
    {
      const regex = new RegExp(`(^| )${name}=([^;]+)`)
      const match = document.cookie.match(regex)
      if (match) {
        return match[2]
      }
   }

function Login() {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');
    useEffect(() => {
        if (getCookieValue('Username') !== undefined) {
        window.location.href = 'main';
        }
    }, []);
	function handleChangeUsername(Event) {
		setUsername(Event.target.value);
	}

	function handleChangePassword(Event) {
		setPassword(Event.target.value);
	}

	function handleEnteringPassword(Event) {
        const invalidChars = /[:~` ]/; // can add to this
        if (invalidChars.test(Event.key)) {
			// listen to the keys pressed on the keyboard and escape invalidchars
            Event.preventDefault();
			alert("Invalid characters entered!");
        }
    }

	function handleChangeVerifyPassword(event) {
        setVerifyPassword(event.target.value);
    }

	async function handleLogin() {
		if (username==='' || password==='') {
			alert("User name or password not enetered!");
		}

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
		if (username==='' || password==='' || verifyPassword==='') {
			alert("Username or password not enetered!");
		}

		if (password !== verifyPassword){
			alert("The passwords you entered don't match, try again!")
			return;
		}
		console.log("sign up successful") //remove for debugging
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
					<label>
					Username: <input name="Username" value={username} onChange={handleChangeUsername} />
					</label>
					<br/>
					<label>
					Password: <input name="Password" type='password' value={password} onChange={handleChangePassword} onKeyDown={handleEnteringPassword} />
					</label>
					<br/>
					{isSigningUp && (
						<>
						<label>
							Verify Password:
							<input
								name="VerifyPassword"
								type="password"
								value={verifyPassword}
								onChange={handleChangeVerifyPassword}
								onKeyDown={handleEnteringPassword}
							/>
						</label>
						<br />
						<button type="button" onClick={handleSignup}>Submit</button>
					</>
					)}
					{!isSigningUp && (
						<>
							<button type="Submit" onClick={handleLogin}>Login</button><br />
							<button type="Submit" onClick={() => setIsSigningUp(true)}>Create Account</button>
						</>
					)}
					
				</form>
			</div>
		</>
	)
}

export default Login
