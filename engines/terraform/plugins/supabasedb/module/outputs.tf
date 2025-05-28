output "db_url" {
  sensitive = true
  value = "postgresql://postgres:${random_password.password.result}@${supabase_project.project.id}.${var.region}.supabase.co:5432/postgres"
}
