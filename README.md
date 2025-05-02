# Building a chat app in Go with WebSockets and Nitric

## What we'll be doing

1. Use Nitric to create a WebSocket API
2. Manage WebSocket connections using a Key-Value store
3. Handle WebSocket events:

- Register connections on connect
- Remove connections on disconnect
- Broadcast messages to all connected clients

4. Run locally for testing
5. Deploy to AWS
