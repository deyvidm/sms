# sms-backend

Web server backend/API which communicates with SMS dispatcher and handles frontend API calls.

## .env setup

The backend web server requires a .env file with the following variables:
```
API_SECRET=xxx      # used for signing JWTs 
ASYNQ_SECRET=xxx    # authenticated the dispatcher -- needs to match dispatcher's .env value
DB_FILE=/xxx/yyy    # filepath to SQLite DB file
```