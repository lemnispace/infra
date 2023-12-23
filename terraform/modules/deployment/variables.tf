variable "api_id" {
  description = "API Gateway ID"
  type        = string
}

variable "route_hashes" {
  description = "values of sha1sum of the routes"
  type        = list(string)
}
