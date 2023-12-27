curl -d `{
  "buckets": [
    {
      "configuration": {
        "bucket": "ssssdfsgagdagadgag2estdgcssdafag",
        "tags": [
          { "key": "Environment", "value": "Dev" },
          { "key": "Project", "value": "ProjectX" }
        ]
      },
      "public_access_block": {
        "acl": "public-read",
        "block_public_acls": false,
        "block_public_policy": false,
        "ignore_public_acls": false,
        "restrict_public_buckets": false
      },
      "website_configuration": {
        "host_website": true,
        "index_document": "index.html",
        "error_document": "error.html",
        "routing_rules": [{
          "redirect": {
            "host_name": "documents/",
            "http_redirect_code":  204,
            "protocol": "https",
            "replace_key_prefix_with": "/docs"
          }
        }]
      },
    "cors_configuration": {
      "cors_rules_list": [{
        "allowed_headers": ["*"],
        "allowed_methods": ["GET", "PUT"],
        "allowed_origins": ["*"],
        "expose_headers": ["x-amz-server-side-encryption"],
        "max_age_seconds": 3000
      },
      {
        "allowed_headers": ["Authorization"],
        "allowed_methods": ["GET"],
        "allowed_origins": ["*"],
        "expose_headers": ["x-amz-server-side-encryption"],
        "max_age_seconds": 3000
      }]
    }
    }
  ]
}` -X POST http://localhost:8080/createbuckets
