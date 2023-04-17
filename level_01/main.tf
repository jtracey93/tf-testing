terraform {
  required_version = ">= 1.4.0, < 2.0.0"
}

# terraform_data.example will always be created.
resource "terraform_data" "example" {
  input = "example"
}

# terraform_data.example_for_each will be created for each key/value pair in the map.
# The input attribute will be set to the value of the map.
resource "terraform_data" "example_for_each" {
  for_each = var.example_for_each
  input    = each.value
}

# terraform_data.example_condition will be created if the variable is set to true.
resource "terraform_data" "example_condition" {
  count = var.example_condition ? 1 : 0
  input = "example_condition"
}
