# Example Output This will be populated in the output field of the step
# PULUMI_COMMAND_STDOUT
# {
#   "name": "sites/SITE_ID/versions/VERSION_ID",
#   "status": "CREATED",
#   "config": {
#     "headers": [{
#       "glob": "**",
#       "headers": {
#         "Cache-Control": "max-age=1800"
#       }
#     }]
#   }
# }
curl -H "Authorization: Bearer $ACCESS_TOKEN" \
     -X DELETE \
https://firebasehosting.googleapis.com/v1beta1/sites/{{.SiteId}}/versions/{{.VersionId}}