# Objective
https://github.com/WolffunService/parser-exercise

# Summary
This project is a blockchain parser built with Go for monitoring specific Ethereum addresses and retrieving their transaction history in real-time. The parser connects to the Ethereum blockchain using JSON-RPC, stores relevant transactions in MongoDB, and exposes a RESTful API using the Echo framework.

The application continuously polls for new blocks every 10 seconds. It fetches transactions within new blocks and checks if they involve any subscribed addresses. Relevant transactions are then stored in MongoDB.

# Features
- Monitor Ethereum Addresses: Watch specific addresses for transactions.
- Fetch Transactions: Collect and store transactions for monitored addresses in MongoDB.
- RESTful API: Provides endpoints for subscribing to addresses, retrieving transactions, and getting the current block number.

# Setup and Installation
1. Environment Configuration: Create a .env file in the root directory and add your environment variables:
```
ALCHEMY_NODE_SERVICES = "https://eth-mainnet.g.alchemy.com/v2/ECXieLFmzlBVVaY3nA5FlZMzGb52td-r" <- Here my 
                                                                                        ALCHEMY_PROJECT_ID 

DATABASE_URL = "mongodb://localhost:27017"
```

2. Docker and MongoDB Setup: Run MongoDB with Docker Compose:
```
docker-compose up -d
```
- Access MongoDB Shell:
```
docker exec -it mongodb mongo
```
- Connect with MongoDB Compass: Connect using `mongodb://localhost:27017`

!!! Tới đây em cũng ko rõ  xài như nào anh có gì anh chỉ với
(At this point, I’m not really sure how to use it either. If you could, please help guide me through it.)

3. Run the Application: Install dependencies and start the application:
```
go mod tidy
go run main.go
```

# API Endpoints
1. Get the Current Block Number
- GET `/blockNumber`
- Returns the current block number from the Ethereum blockchain.
2. Subscribe to an Address
- POST `/subscribe`
- Body: `{ "address": "0xYourEthereumAddress" }`
- Subscribes an Ethereum address for monitoring.
3. Get Transactions for a Subscribed Address
- GET `/transactions?address=0xYourEthereumAddress`
- Returns all transactions for a subscribed address.