curl -H "Content-Type: application/json" \
         -H "Authorization: Bearer {{.AccessToken}}" \
         -d '{
               "files": {{.Files}}
             }' \
https://firebasehosting.googleapis.com/v1beta1/sites/{{.SiteId}}/versions/{{.VersionId}}:populateFiles