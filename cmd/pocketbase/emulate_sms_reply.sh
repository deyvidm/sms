
generate_post_data()
{
  cat <<EOF
{
  "Type" : "Notification",
  "MessageId" : "f06ca0bc-75d0-5a29-8846-fc0d2c376830",
  "TopicArn" : "arn:aws:sns:us-east-1:642644262173:SMSEventManager",
  "Message" : "{\"originationNumber\":\"+16478653169\",\"destinationNumber\":\"+15876085626\",\"messageKeyword\":\"KEYWORD_642644262173\",\"messageBody\":\"Thanks\",\"previousPublishedMessageId\":\"90f1edf9-fe20-5f06-9243-b15320419b86\",\"inboundMessageId\":\"cb10b39f-e14c-4391-88d9-53942dc1d87b\"}",
  "Timestamp" : "2023-04-13T01:15:19.084Z",
  "SignatureVersion" : "1",
  "Signature" : "qBY09lVyYsZfNdqT8FF1NWFz4DzWRo54NUsj87rlfGrBJKeRfETCwF3ErYWcJpPZNwSbp7zRoCTwf2L/kGCGkufXTb7TaSownfanKdGpIH5zx0zi4JyZIdNyszK+YWkxfgz00i8qcVaWF+i1XVDB7oFzbrrO8JpUD1GZmGc9INnLyiH4Tl7voaaLBVqOTXZ/xC/RJVErPpN6DEB9LETLliE1b3grLXwVqAFBKtGW+M7qQdyDtweG1EDCl0bpv/9s4yAiO35sIC+tMhiOL/+jvVQkKxxBaRdDTEZ+6LXx0aaF3eFiW/cv10caiS9EUH2UbKvbm+QqPqSfYBRzAp8R2A==",
  "SigningCertURL" : "https://sns.us-east-1.amazonaws.com/SimpleNotificationService-56e67fcb41f6fec09b0196692625d385.pem",
  "UnsubscribeURL" : "https://sns.us-east-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-1:642644262173:SMSEventManager:ea9cd29c-8cc0-43c0-8e0b-8e37f447c46f"
}
EOF
}

# echo "'"$(generate_post_data)"'"
# echo 
# echo 
# echo "$(generate_post_data)"

  # -X POST --data "'"$(generate_post_data)"'"  http://127.0.0.1:8090/api/respond 

curl \
  -H 'Content-Type: application/json' \
  -H "Accept: application/json" \
  -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJfcGJfdXNlcnNfYXV0aF8iLCJleHAiOjE2ODI1NjU1MjcsImlkIjoicWdkY3AxdjZwOWptZ3lyIiwidHlwZSI6ImF1dGhSZWNvcmQifQ.uuQrSviPGz5gEd9EvTuM4gmWZ0FG9AUsc2-tbKh3y6w' \
  -X POST --data "$(generate_post_data)" http://127.0.0.1:8090/api/respond

# curl 'http://127.0.0.1:8090/api/respond' \
# -X POST \
# -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/110.0' \
# -H 'Accept: */*' \
# -H 'Accept-Language: en-US' \
# -H 'Accept-Encoding: gzip, deflate, br' \
# -H 'Content-Type: application/json' \
# -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJfcGJfdXNlcnNfYXV0aF8iLCJleHAiOjE2ODI1NjU1MjcsImlkIjoicWdkY3AxdjZwOWptZ3lyIiwidHlwZSI6ImF1dGhSZWNvcmQifQ.uuQrSviPGz5gEd9EvTuM4gmWZ0FG9AUsc2-tbKh3y6w' \
# -H 'Origin: http://localhost:5173' \
# -H 'DNT: 1' \
# -H 'Connection: keep-alive' \
# -H 'Sec-Fetch-Dest: empty' \
# -H 'Sec-Fetch-Mode: cors' \
# -H 'Sec-Fetch-Site: cross-site' --data-raw '{"first_name":"sdfg","last_name":"sadff","phone":"+11223222222","owner":"qgdcp1v6p9jmgyr","status":"Pending"}'
