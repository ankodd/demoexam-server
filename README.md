# Server for demoexam app
## Install and Run:
1. Clone repository: ```git clone https://github.com/ankodd/demoexam-server.git```
2. Install dependencies: ```cd demoexam-server && go get .```
3. For environment variables need create .env file and add path to storage, and port. Example: ```PORT=8080 STORAGE_PATH=storage/storage.db```
4. Create storage directory: ```mkdir storage```
5. Run: ```go run ./cmd/main.go``` Or build ```go build -o build/run_server .cmd/main.go```
