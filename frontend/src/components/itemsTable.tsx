import {useStore} from '@nanostores/react'
import {cartStorage} from '@/utils/stores'
import Image from 'next/image'

interface ListItemProps {
  data: CartItem
}
function ListItem({data}: ListItemProps) {
  return (
    <tr>
      <td>
        <div className="flex items-center space-x-3">
          <div className="avatar">
            <div className="mask mask-squircle w-12 h-12">
              {/*  <Image src="https://illustoon.com/photo/11612.png" alt="item" />  */}
              <Image src={'/water.webp'} alt="item" />
            </div>
          </div>
          <div>
            <div className="font-bold">{data.products}</div>
          </div>
        </div>
      </td>
      <td>{data.price}</td>
      <td>{data.quantity}</td>
      <td>{data.total}</td>
    </tr>
  )
}

function getData() {
  const local = cartStorage.get().products
  const $cart: Cart = JSON.parse(local)
  return $cart
}

function ItemsTable() {
  const $cart = getData()

  const sumItems = () => {
    let total = 0
    $cart?.items?.forEach((item) => {
      total += item.quantity
    })
    return total
  }

  return (
    <div>
      <div className="overflow-x-auto">
        {sumItems() !== 0 ? (
          <table className="table">
            <thead>
              <tr>
                <th>Item</th>
                <th>Price</th>
                <th>Quantity</th>
                <th>Total</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {$cart?.items?.map((item, index) => (
                <ListItem key={index} data={item} />
              ))}
            </tbody>
          </table>
        ) : (
          <div></div>
        )}
      </div>
      <div className="card-body">
        <span className="font-bold text-lg">{sumItems()} Items</span>
        <span className="text-info">TOTAL PRICE: {$cart.total}</span>
        {/* <div className="card-actions">
          <button className="btn btn-primary btn-block">View cart</button>
        </div> */}
      </div>
    </div>
  )
}

export default ItemsTable
