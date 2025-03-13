curl -H "Content-Type: application/json" \
       -H "Authorization: Bearer $ACCESS_TOKEN" \
       -X PATCH \
       -d '{"status": "FINALIZED"}' \
https://firebasehosting.googleapis.com/v1beta1/sites/{{.SiteId}}/versions/{{.VersionId}}?update_mask=status