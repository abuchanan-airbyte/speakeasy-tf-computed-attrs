resource "terraform_connection" "my_connection" {
  configurations = {
    streams = [
      {
        cursor_field = [
          "..."
        ]
        name = "...my_name..."
        primary_key = [
          {
            # ...
          }
        ]
      }
    ]
  }
  destination_id = "5725b342-2d43-4e6c-90a4-e500c954e591"
  name           = "...my_name..."
  source_id      = "b5b2b4a5-bba6-4c3f-b0ef-ab87b373f331"
}