curl -H "Authorization: Bearer $ACCESS_TOKEN" \
       -X POST
https://firebasehosting.googleapis.com/v1beta1/sites/{{.SiteId}}/releases?versionName=sites/{{.SiteId}}/versions/{{.VersionId}}