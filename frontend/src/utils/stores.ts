import {persistentMap} from '@nanostores/persistent'

export type Storage = {
  products: string
  pocket: string
  bills: string
}

export const cartStorage = persistentMap<Storage>('cart:', {
  products: JSON.stringify([]),
  pocket: JSON.stringify([]),
  bills: JSON.stringify([]),
})
