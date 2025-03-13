curl -H "Content-Type: application/json" \
       -H "Authorization: Bearer $ACCESS_TOKEN" \
       -d '{{.Config}}' \
https://firebasehosting.googleapis.com/v1beta1/sites/{{.SiteId}}/versions