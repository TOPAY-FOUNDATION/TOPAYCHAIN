# TOPAYCHAIN API Documentation

This document outlines the API endpoints available in **TOPAYCHAIN**. These endpoints enable interaction with the blockchain, wallets, transactions, and nodes.

---

## Base URL

The API base URL:
```
http://localhost:8080
```

---

## Endpoints

### 1. **Blockchain**

#### **GET /blocks**
Retrieve the current blockchain.

- **URL**: `/blocks`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "index": 0,
      "timestamp": "2024-01-01T00:00:00Z",
      "transactions": [],
      "nonce": 100,
      "previous_hash": "0",
      "hash": "abc123..."
    },
    ...
  ]
  ```

#### **POST /blocks/mine**
Mine a new block.

- **URL**: `/blocks/mine`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "miner_address": "tpy1qxyz..."
  }
  ```
- **Response**:
  ```json
  {
    "message": "Block mined successfully",
    "block": {
      "index": 1,
      "timestamp": "2024-01-01T01:00:00Z",
      "transactions": [...],
      "nonce": 200,
      "previous_hash": "abc123...",
      "hash": "def456..."
    }
  }
  ```

---

### 2. **Transactions**

#### **POST /transactions**
Create a new transaction.

- **URL**: `/transactions`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "sender": "tpy1qabc...",
    "receiver": "tpy1qxyz...",
    "amount": 100.0,
    "signature": "abcdef..."
  }
  ```
- **Response**:
  ```json
  {
    "message": "Transaction created successfully",
    "transaction": {
      "sender": "tpy1qabc...",
      "receiver": "tpy1qxyz...",
      "amount": 100.0,
      "signature": "abcdef..."
    }
  }
  ```

#### **GET /transactions/pending**
Retrieve all pending transactions.

- **URL**: `/transactions/pending`
- **Method**: `GET`
- **Response**:
  ```json
  [
    {
      "sender": "tpy1qabc...",
      "receiver": "tpy1qxyz...",
      "amount": 100.0,
      "signature": "abcdef..."
    },
    ...
  ]
  ```

---

### 3. **Wallets**

#### **POST /wallets**
Create a new wallet.

- **URL**: `/wallets`
- **Method**: `POST`
- **Response**:
  ```json
  {
    "address": "tpy1qabc...",
    "private_key": "abcdef...",
    "public_key": "123456..."
  }
  ```

#### **GET /wallets/{address}**
Retrieve wallet details.

- **URL**: `/wallets/{address}`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "address": "tpy1qabc...",
    "balance": 1000.0,
    "transactions": [...]
  }
  ```

---

### 4. **Node Management**

#### **GET /nodes**
Retrieve a list of connected nodes.

- **URL**: `/nodes`
- **Method**: `GET`
- **Response**:
  ```json
  [
    "http://127.0.0.1:8081",
    "http://127.0.0.1:8082"
  ]
  ```

#### **POST /nodes/register**
Register a new node to the network.

- **URL**: `/nodes/register`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "nodes": [
      "http://127.0.0.1:8081",
      "http://127.0.0.1:8082"
    ]
  }
  ```
- **Response**:
  ```json
  {
    "message": "Nodes added successfully",
    "nodes": [
      "http://127.0.0.1:8081",
      "http://127.0.0.1:8082"
    ]
  }
  ```

#### **GET /nodes/consensus**
Resolve conflicts and synchronize the blockchain.

- **URL**: `/nodes/consensus`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "message": "Blockchain updated to the longest chain",
    "chain": [...]
  }
  ```

---

## Error Handling

All API responses include HTTP status codes to indicate success or failure:

- `200 OK`: Request was successful.
- `400 Bad Request`: The request was invalid or missing required data.
- `404 Not Found`: The requested resource does not exist.
- `500 Internal Server Error`: An error occurred on the server.

Example error response:
```json
{
  "error": "Invalid transaction signature"
}
```

---

## Future Enhancements

- Support for smart contracts.
- Additional analytics endpoints for blockchain statistics.
- WebSocket support for real-time updates.

---

## Contact
For further assistance, visit [Topay Foundation](https://www.topayfoundation.com).

