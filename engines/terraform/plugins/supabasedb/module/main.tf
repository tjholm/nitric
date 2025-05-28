# Create a random password
resource "random_password" "password" {
  length  = 16
  special = true
}

resource "supabase_project" "project" {
  organization_id   = var.organization_id
  name              = var.name
  database_password = random_password.password.result
  region            = var.region
  lifecycle {
    ignore_changes = [database_password]
  }
}
# Configure api settings for the linked project
resource "supabase_settings" "production" {
  project_ref = supabase_project.project.id
  api = jsonencode({
    db_schema            = "public,storage,graphql_public"
    db_extra_search_path = "public,extensions"
    max_rows             = 1000
  })
}