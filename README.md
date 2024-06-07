## Installation

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
