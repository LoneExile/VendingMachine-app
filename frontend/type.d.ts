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

declare interface Denomination {
  id: string
  denomination_value: number
  stock: number
  typed: string
}

declare interface DenominationsResponse {
  denominations: Denomination[]
}

declare interface CartItem {
  id: string
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

declare interface PocketItem {
  id: string
  denomination_value: number
  stock: number
  typed: string
  quantity: number
  total: number
}

declare interface Pocket {
  items: PocketItem[]
  balance: number
}

declare interface BillItem {
  id: string
  denomination_value: number
  stock: number
  typed: string
  quantity: number
}
declare interface Chagne {
  change: BillItem[]
  message: string
  status: string
}

declare interface Bill {
  message: Chagne
}
