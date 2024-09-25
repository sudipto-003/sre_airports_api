variable "bucket_name" {
  type    = string
  default = ""
}

variable "bucket_tags" {
  type    = map(string)
  default = {}
}

variable "bucket_policy_allow_principals" {
  type    = list(map(any))
  default = []
}
