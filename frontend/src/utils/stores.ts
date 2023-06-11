import {persistentMap} from '@nanostores/persistent'

export type CartStorage = {
  products: string
}

export const cartStorage = persistentMap<CartStorage>('cart:', {
  products: JSON.stringify([]),
})
