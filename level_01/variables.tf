variable "example_for_each" {
  type        = map(string)
  default     = {}
  description = "This variable will control the creation of the terraform_data.example_for_each resource."
}

variable "example_condition" {
  type        = bool
  default     = false
  description = "This variable will control the creation of the terraform_data.example_condition resource."
}
