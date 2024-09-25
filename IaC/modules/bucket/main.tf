resource "aws_s3_bucket" "this" {
    bucket = var.bucket_name

    force_destroy = var.bucket_force_destroy
    object_lock_enabled = var.bucket_object_lock
    tags = var.bucket_tags
}

resource "aws_s3_bucket_policy" "this_bucket_policy" {
    bucket = aws_s3_bucket.this.id
    policy = data.aws_iam_policy_document.this_bucket_policy_doc.json
}

data "aws_iam_policy_document" "this_bucket_policy_doc" {
    statement {
        dynamic principals {
            for_each = var.bucket_policy_allow_principals

            content {
                type = principals.value["effect"]
                identifiers = principals.value["identity"]
            }
        }

        actions = ["s3:GetObject", "s3:ListObject", "s3:PutObject"]

        resources = [ aws_s3_bucket.this.arn ]
    }

}