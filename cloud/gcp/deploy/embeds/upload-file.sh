curl -H "Authorization: Bearer {{.AccessToken}}" \
       -H "Content-Type: application/octet-stream" \
       --data-binary @{{.FilePath}} \
https://upload-firebasehosting.googleapis.com/upload/sites/{{.SiteId}}/versions/{{.VersionId}}/files/{{.FileHash}}