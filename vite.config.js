import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	build: {
		rollUpOptions: {
			input: {
				user: 'user.html',
				login: 'login.html',
				main: 'main.html',
				upload: 'upload.html',
				pfp: 'pfp.html',
				uploaded: 'uploaded.html'
			}
		}
	}
})
