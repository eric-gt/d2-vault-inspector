env "local" {
    src = "file://schema/schema.sql"
    dev = "docker://postgres/16/dev?search_path=public"
    url= "postgres://root:password1234@localhost:5432/database?search_path=public&sslmode=disable"
    migration {
        dir = "file://migrations"
    }
}
