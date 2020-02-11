# Example request
curl --request POST \
  --url http://127.0.0.1:8089/iam/service_token \
  --header 'authorization: Basic YWRtaW46TnV0YW5peC4xMjM=' \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data audience=xx \
  --data domain=yy \
  --data assertion=zz
