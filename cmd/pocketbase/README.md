## Project is a Work-In-Progress 

# SMS Event Plannet -- backend
This is the backend API for my SMS-based RSVP service. [Frontend available here](https://github.com/deyvidm/sms-event-planner-front)


### What is it? 
This project allows users to create events, send invites, and collect responses from other users using SMS.

The motivation for the project is to help my friend organize pickup volleyball tournaments. 

### PreRequisites & Requirements
* GoLang 1.19
* configured & functional AWS auth
  * the easiest way to achieve this is to download and confgure the [AWS CLI](https://aws.amazon.com/cli/)

### Up and Running
Once you have sorted the prerequisites, navigate to the root of the repo and run
```
go mod download // only needed once, to download all go dependencies
make
```

<img width="563" alt="successful make" src="https://user-images.githubusercontent.com/16841661/233929269-0010ded0-c38d-4ba3-bdb0-c324e4cba567.png">


If setup was successful and the database is empty, then visiting [the admin UI](http://127.0.0.1:8090/_/) should produce this web page: 
<img width="578" alt="image" src="https://user-images.githubusercontent.com/16841661/233953936-3330d786-f80c-4722-80d7-9009b97108af.png">
