# Terraform Provider Hashicups

This repo is a companion repo to the [Call APIs with Terraform Providers](https://learn.hashicorp.com/collections/terraform/providers)
Learn collection. 

In the collection, you will use the HashiCups provider as a bridge between Terraform and the HashiCups API.
Then, extend Terraform by recreating the HashiCups provider.
By the end of this collection, you will be able to take these intuitions to create your own custom Terraform provider. 

Visit the [`boilerplate`](https://github.com/hashicorp/terraform-provider-hashicups/tree/boilerplate) branch of this repository for this
Terraform provider's specific starter template. The [Terraform Provider Scaffold](https://github.com/hashicorp/terraform-provider-scaffolding)
is a quick-start repository for creating a Terraform provider. Use this GitHub template when you're ready to create your own custom provider.

## Build provider

Run the following command to build the provider

```bash
make build
```

## Local release build

```bash
make release
```

You will find the releases in the `/dist` directory.
You will need to rename the provider binary to `terraform-provider-hashicups` and move the binary into
[the appropriate subdirectory within the user plugins directory](https://learn.hashicorp.com/tutorials/terraform/provider-use?in=terraform/providers#install-hashicups-provider).

## Test sample configuration

```bash
# build and install the provider
make install
# navigate to the `examples` directory
cd examples
# initialize the workspace
terraform init
# apply the sample configuration
terraform apply
```
