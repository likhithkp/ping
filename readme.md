# Ping - Distributed Real-Time Messaging Platform

A scalable, fault-tolerant, distributed messaging platform inspired by WhatsApp.

- One-to-One Messaging
- Group Messaging
- Online Presence
- Message Delivery Acknowledgements
- Media Sharing
- Push Notifications
- Audio Calling
- Video Calling
- Group Video Calling
- Multi-Device Synchronization

---

# Table of Contents

1. Problem Statement
2. Functional Requirements
3. Non-Functional Requirements
4. Capacity Estimation
5. Core APIs
6. Data Model
7. High-Level Architecture
8. WebSocket Architecture
9. Service Discovery
10. Presence Service
11. Message Routing
12. Online Message Delivery
13. Offline Message Delivery
14. Message Acknowledgements
15. Multi-Device Synchronization
16. Group Messaging
17. Media Sharing
18. Push Notification Service
19. Audio Calling
20. Video Calling
21. Group Video Calling
22. Scaling Strategy
23. Reliability
24. Security
25. Observability
26. Future Improvements

---

# 1. Problem Statement

Design a globally distributed messaging platform capable of supporting millions of users simultaneously.

Users should be able to:

- Send messages instantly
- Create groups
- Share media
- See online status
- Receive notifications
- Synchronize across multiple devices
- Make voice calls
- Make video calls

The system must provide:

- Low latency
- High availability
- High durability
- Horizontal scalability

---

# 2. Functional Requirements

## Messaging

Support:

- One-to-one messaging
- Group messaging
- Rich media messages
- Message persistence

---

## Presence

Users can view:

- Online
- Offline
- Last Seen

---

## Delivery Status

Messages support:

```text
SENT
DELIVERED
READ
```

---

## Notifications

Notify users about:

- New messages
- Missed calls
- Group invitations

---

## Media Sharing

Support:

- Images
- Videos
- Audio files
- Documents

---

## Voice Calls

Support:

- Initiate call
- Accept call
- Reject call
- End call
- Missed call tracking

---

## Video Calls

Support:

- One-to-one video calls
- Group video calls
- Camera controls
- Screen sharing (future)

---

# 3. Non-Functional Requirements

| Requirement     | Target     |
| --------------- | ---------- |
| Availability    | 99.99%     |
| Latency         | <100ms     |
| Scalability     | Horizontal |
| Durability      | High       |
| Reliability     | High       |
| Fault Tolerance | High       |

---

# 4. Capacity Estimation

Assumptions:

```text
100M Daily Active Users
10M Concurrent Connections
50B Messages / Day
```

Average Traffic:

```text
~600K Messages / Second
```

Peak Traffic:

```text
1M+ Messages / Second
```

---

# 5. Core APIs

## Send Message

```http
POST /messages
```

Request:

```json
{
  "senderId": "user1",
  "receiverId": "user2",
  "content": "Hello"
}
```

---

## Get Messages

```http
GET /conversations/{id}/messages
```

---

## Get Presence

```http
GET /users/{id}/presence
```

---

## Start Call

```http
POST /calls/start
```

---

## End Call

```http
POST /calls/end
```

---

# 6. Data Model

## Users

```sql
users
-----
id
phone_number
name
created_at
```

---

## Conversations

```sql
conversations
-------------
id
type
created_at
```

---

## Messages

```sql
messages
--------
id
conversation_id
sender_id
content
type
status
created_at
```

---

## Groups

```sql
groups
------
id
name
owner_id
created_at
```

---

## Group Members

```sql
group_members
-------------
group_id
user_id
role
```

---

## Calls

```sql
calls
-----
id
caller_id
receiver_id
type
status
started_at
ended_at
duration
```

---

# 7. High-Level Architecture

```text
                    Clients
                       │
                       ▼
                 API Gateway
                       │
                       ▼
                Load Balancer
                       │
 ┌──────────────┬──────────────┬──────────────┐
 │              │              │              │
 ▼              ▼              ▼              ▼

Chat Service Presence Service Call Service Notification Service

 │              │              │
 └──────────────┼──────────────┘
               ▼

          Kafka Cluster

               ▼

   ┌────────────┬────────────┐
   ▼            ▼            ▼

Postgres      Redis      Object Storage
```

---

# 8. WebSocket Architecture

Persistent WebSocket connections are used for:

- Real-time messaging
- Presence updates
- Typing indicators
- Delivery acknowledgements
- Call signaling

Benefits:

- Low latency
- Bi-directional communication
- Reduced connection overhead

---

# 9. Service Discovery

Responsibilities:

- Discover active chat servers
- Route users to correct nodes
- Support horizontal scaling

Possible solutions:

- Consul
- Kubernetes Service Discovery
- etcd

---

# 10. Presence Service

Tracks:

```text
Online
Offline
Last Seen
```

Redis stores:

```text
user:{id}:presence
```

Example:

```json
{
  "status": "online",
  "lastSeen": "2026-06-08T10:00:00Z"
}
```

---

# 11. Message Routing

Flow:

```text
Sender
  │
  ▼
WebSocket
  │
  ▼
Chat Service
  │
  ▼
Receiver Connection
```

If receiver is offline:

```text
Store Message
Send Notification
Deliver Later
```

---

# 12. Online Message Delivery

```text
Sender
  │
  ▼
Chat Service
  │
  ▼
Receiver Online
  │
  ▼
Instant Delivery
```

---

# 13. Offline Message Delivery

```text
Sender
  │
  ▼
Chat Service
  │
  ▼
Database
  │
  ▼
Receiver Reconnects
  │
  ▼
Pending Messages Delivered
```

---

# 14. Message Acknowledgements

States:

```text
SENT
DELIVERED
READ
```

Flow:

```text
Message Sent
   │
   ▼
Delivered
   │
   ▼
Read
```

---

# 15. Multi-Device Synchronization

Users may have:

- Mobile
- Tablet
- Desktop

Message fan-out:

```text
User
 ├── Phone
 ├── Tablet
 └── Desktop
```

All active devices receive updates.

---

# 16. Group Messaging

Components:

- Groups
- Members
- Admins

Flow:

```text
Sender
  │
  ▼
Group Service
  │
  ▼
Member Fanout
```

---

# 17. Media Sharing

Media is stored separately.

```text
Client
  │
Upload
  │
  ▼
Object Storage
  │
  ▼
URL Generated
  │
  ▼
Message Contains URL
```

Store:

- Images
- Videos
- Audio
- Documents

Possible storage:

- AWS S3
- MinIO

---

# 18. Push Notification Service

Responsible for:

- New messages
- Missed calls
- Group invitations

Platforms:

- FCM
- APNS

---

# 19. Audio Calling

## Architecture

```text
Caller
   │
   ▼
Signaling Service
   │
   ▼
Receiver

        WebRTC
```

Components:

- WebRTC
- STUN Server
- TURN Server
- Signaling Service

Call Flow:

```text
1. Call initiated
2. Receiver notified
3. SDP Exchange
4. ICE Exchange
5. Peer Connection Established
6. Audio Streaming Starts
```

---

# 20. Video Calling

Uses:

```text
WebRTC
```

Additional Components:

- Video Encoder
- Video Decoder
- Adaptive Bitrate Controller

Flow:

```text
Caller
   │
   ▼
Signaling
   │
   ▼
Receiver
   │
   ▼
WebRTC Session
```

---

# 21. Group Video Calling

Mesh architecture does not scale.

Instead use:

```text
SFU (Selective Forwarding Unit)
```

Architecture:

```text
Participant A
Participant B
Participant C
       │
       ▼
      SFU
       ▲
Participant D
Participant E
```

Possible SFU Solutions:

- LiveKit
- Janus
- mediasoup
- Jitsi

---

# 22. Scaling Strategy

## Stateless Services

All services remain stateless.

---

## Redis Cluster

Used for:

- Presence
- Sessions
- Caching

---

## Kafka

Used for:

- Message events
- Notification events
- Call events

---

## Database Sharding

Shard by:

```text
User ID
```

or

```text
Conversation ID
```

---

# 23. Reliability

Techniques:

- Retries
- Dead Letter Queues
- Replication
- Backpressure
- Circuit Breakers

---

# 24. Security

Authentication:

```text
JWT
```

Authorization:

```text
Role-Based Access Control
```

Future:

```text
End-to-End Encryption
```

---

# 25. Observability

Metrics:

- Active Connections
- Messages/sec
- Message Delivery Latency
- Active Calls
- Video Sessions
- Kafka Lag
- Redis Latency

Tools:

- Prometheus
- Grafana
- OpenTelemetry

---

# 26. Future Improvements

- End-to-End Encryption
- Status / Stories
- Message Reactions
- Live Location Sharing
- Message Search
- AI Assistant
- Screen Sharing
- Voice Notes
- Community Groups

---

# Tech Stack

| Component     | Technology |
| ------------- | ---------- |
| Backend       | Go         |
| API           | gRPC       |
| Realtime      | WebSocket  |
| Voice/Video   | WebRTC     |
| Database      | PostgreSQL |
| Cache         | Redis      |
| Queue         | Kafka      |
| Storage       | S3 / MinIO |
| Monitoring    | Prometheus |
| Visualization | Grafana    |
| Containers    | Docker     |
| Orchestration | Kubernetes |

---

# License

MIT License
