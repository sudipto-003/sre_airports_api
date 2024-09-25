bucket_name = "airport-images"
bucket_tags = {
  "owner"       = "sre"
  "app_version" = "v2"
}

bucket_policy_allow_principals = [
  {
    "type"     = "AWS",
    "identity" = "arn:aws:iam::111122223333:user/app-user"
  }
]