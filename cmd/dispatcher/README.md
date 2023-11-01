# sms/cmd/dispatcher
This project serves as an SMS dispatcher and invitation manager for my SMS RSVP service. Designed to be used in tandem with sms-backend.

## environment setup

### .env file
The dispatcher requires a .env file with the following variables: 
```
SECRET=xxx  # auth token for communicating with web-server
```

### redis
The dispatcher uses redis to manage message (task) queuing and handling.
It expects that redis is running on port **6379**


