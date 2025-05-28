output "db_url" {
  sensitive = true
  value = "postgresql://postgres:${random_password.password.result}@${supabase_project.project.ref}.${var.region}.supabase.co:5432/postgres"
}
