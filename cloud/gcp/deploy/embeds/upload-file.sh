curl -H "Authorization: Bearer $ACCESS_TOKEN" \
     -H "Content-Type: application/octet-stream" \
     --data-binary @{{.FilePath}} \
https://upload-firebasehosting.googleapis.com/upload/sites/{{.SiteId}}/versions/{{.VersionId}}/files/{{.FileHash}}