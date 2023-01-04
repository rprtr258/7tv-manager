terraform {
  required_providers {
    hashicups = {
      version = "0.0.1"
      source = "local/rprtr258/seventv"
    }
  }
}

provider "hashicups" {
  username = "education"
  password = "test123"
}

resource "hashicups_order" "sample" {}

output "sample_order" {
  value = hashicups_order.sample
}

# coffee_name = "Packer Spiced Latte"
variable "coffee_name" {
  type    = string
  default = "Vagrante espresso"
}

data "hashicups_coffees" "all" {}

# Returns all coffees
output "all_coffees" {
  value = data.hashicups_coffees.all.coffees
}

# Only returns packer spiced latte
output "coffee" {
  value = {
    for coffee in data.hashicups_coffees.all.coffees :
    coffee.id => coffee
    if coffee.name == var.coffee_name
  }
}

# output "psl" {
#   value = module.psl.coffee
# }

data "hashicups_order" "order" {
  id = 1
}

output "order" {
  value = data.hashicups_order.order
}

resource "hashicups_order" "edu" {
  items {
    coffee {
      id = 3
    }
    quantity = 2
  }
  items {
    coffee {
      id = 2
    }
    quantity = 3
  }
}

output "edu_order" {
  value = hashicups_order.edu
}


data "hashicups_order" "first" {
  id = 1
}

output "first_order" {
  value = data.hashicups_order.first
}
