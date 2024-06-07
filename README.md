# Study Spotter

35L Final Project group: Emacs Enthusiasts

## Group Member Email Mapping

Tony Jeon: tonysjeon@gmail.com

Yuxin Bao: 104787687+c13752hz@users.noreply.github.com

Raymond Shao: 44079695+RaymondShao777@users.noreply.github.com

Allen Tong: allensctong@gmail.com

Sky Wang: 88114746+WindskyVG@users.noreply.github.com

Antara Chugh: antara.chugh15@gmail.com

## Installation
### Frontend
1. Clone the repository to a local directory:
    ```sh
    git clone https://github.com/allensctong/cs35l-studyspotter.git
    ```
2. Navigate to the project directory:
    ```sh
    cd cs35l-studyspotter
    ```
3. Install frontend dependencies (NPM):
    ```sh
    npm install
    ```

4. Spin up frontend:
    ```sh
    npx vite dev
    ```

### Backend
5. Install backend dependencies:
Install Go based on your device: https://go.dev/doc/install
Execute the following commands on another terminal:
    ```sh
    go mod init studyspotter
    go get .
    ```
6. Spin up the backend server
    ```sh
    go run main.go
    ```

### Database
1. No set up is needed for the database, but use this if you ever need to reset the database:
    ```sh
    rm studyspotter.db
    ```

## Running the App
Commands 4. and 6. need to be running for the app to work. The frontend will be hosted at [http://localhost:5173](http://localhost:5173/) and backend is at [http://localhost:8080](http://localhost:8080/).
