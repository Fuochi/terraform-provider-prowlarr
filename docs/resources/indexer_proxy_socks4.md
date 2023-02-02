---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "prowlarr_indexer_proxy_socks4 Resource - terraform-provider-prowlarr"
subcategory: "Indexer Proxies"
description: |-
  Indexer Proxy Socks4 resource.
  For more information refer to Indexer Proxy https://wiki.servarr.com/prowlarr/settings#indexer-proxies and Socks4 https://wiki.servarr.com/prowlarr/supported#socks4.
---

# prowlarr_indexer_proxy_socks4 (Resource)

<!-- subcategory:Indexer Proxies -->Indexer Proxy Socks4 resource.
For more information refer to [Indexer Proxy](https://wiki.servarr.com/prowlarr/settings#indexer-proxies) and [Socks4](https://wiki.servarr.com/prowlarr/supported#socks4).

## Example Usage

```terraform
resource "prowlarr_indexer_proxy_socks4" "example" {
  name     = "Example"
  host     = "localhost"
  port     = 8080
  username = "User"
  password = "Pass"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) host.
- `name` (String) Indexer Proxy name.
- `password` (String, Sensitive) Password.
- `port` (Number) Port.
- `username` (String) Username.

### Optional

- `tags` (Set of Number) List of associated tags.

### Read-Only

- `id` (Number) Indexer Proxy ID.

## Import

Import is supported using the following syntax:

```shell
# import using the API/UI ID
terraform import prowlarr_indexer_proxy_socks4.example 1
```