declare interface Product {
  id: string
  picture: string
  price: string
  product_name: string
  stock: number
}

declare interface ProductsResponse {
  products: Product[]
}

declare interface CartItem {
  products: string
  quantity: number
  price: number
  total: number
  picture: string
}

declare interface Cart {
  items: CartItem[]
  total: number
}
