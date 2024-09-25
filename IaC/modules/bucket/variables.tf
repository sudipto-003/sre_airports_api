variable "bucket_using_prefix" {
    type = bool
    default = false
}

variable "bucket_name" {
  type    = string
  default = ""
}

variable "bucket_force_destroy" {
  type    = bool
  default = false
}

variable "bucket_object_lock" {
  type    = bool
  default = false
}

variable "bucket_tags" {
  type    = map(string)
  default = {}
}

variable "bucket_policy_allow_principals" {
  type    = list(map(any))
  default = [ ]
}
