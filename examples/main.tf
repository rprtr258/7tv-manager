terraform {
  required_providers {
    seventv = {
      version = "0.0.1"
      source = "local/rprtr258/seventv"
    }
  }
}

provider "seventv" {
  username = "education"
  password = "test123"
}

resource "seventv_emoteset" "main_emoteset" {
  name = "testset2"
  emotes {
    id = "61801776e0801fb98788c028"
    name = "MMMM"
  }
  emotes {
    id = "60a1babb3c3362f9a4b8b33a"
    name = "catKiss"
  }
}

# output "sample_order" {
#   value = seventv_order.sample
# }

# # coffee_name = "Packer Spiced Latte"
# variable "coffee_name" {
#   type    = string
#   default = "Vagrante espresso"
# }

# data "seventv_coffees" "all" {}

# # Returns all coffees
# output "all_coffees" {
#   value = data.seventv_coffees.all.coffees
# }

# # Only returns packer spiced latte
# output "coffee" {
#   value = {
#     for coffee in data.seventv_coffees.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }

# # output "psl" {
# #   value = module.psl.coffee
# # }

# data "seventv_order" "order" {
#   id = 1
# }

# output "order" {
#   value = data.seventv_order.order
# }

# resource "seventv_order" "edu" {
#   items {
#     coffee {
#       id = 3
#     }
#     quantity = 2
#   }
#   items {
#     coffee {
#       id = 2
#     }
#     quantity = 3
#   }
# }

# output "edu_order" {
#   value = seventv_order.edu
# }


# data "seventv_order" "first" {
#   id = 1
# }

# output "first_order" {
#   value = data.seventv_order.first
# }
