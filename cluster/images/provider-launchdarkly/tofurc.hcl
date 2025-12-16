provider_installation {
  filesystem_mirror {
    path    = "/opentofu/provider-mirror"
    include = ["*/*"]
  }
  direct {
    exclude = ["*/*"]
  }
}

