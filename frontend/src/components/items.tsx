import Link from 'next/link'

interface ListItemProps {
  index: number
}
function ListItem({index}: ListItemProps) {
  return (
    <tr>
      <td>
        <div className="flex items-center space-x-3">
          <div className="avatar">
            <div className="mask mask-squircle w-12 h-12">
              <img
                src="https://illustoon.com/photo/11612.png"
                alt="Avatar Tailwind CSS Component"
              />
            </div>
          </div>
          <div>
            <div className="font-bold">{`Water ${index + 1}`}</div>
          </div>
        </div>
      </td>
      <td>50$</td>
      <td>20</td>
      <td>1000</td>
    </tr>
  )
}

export default function Items() {
  return (
    <>
      <div className="overflow-x-auto">
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
            {Array.from({length: 6}, (_, i) => (
              <ListItem key={i} index={i} />
            ))}
          </tbody>
        </table>
      </div>
      <div className="card shadow-lg bg-base-100 my-4">
        <div className="card-body">
          <div className="flex justify-between">
            <div>Subtotal</div>
            <div>1000$</div>
          </div>
          <div className="flex justify-between">
            <div>Tax</div>
            <div>100$</div>
          </div>
          <div className="flex justify-between">
            <div>Total</div>
            <div>1100$</div>
          </div>
        </div>
      </div>
      {/* go back*/}
      <div className="flex justify-between m-4">
        <Link href="/">
          <button className="btn btn-primary">Back</button>
        </Link>

        <Link href="/payment">
          <button className="btn btn-primary">Checkout</button>
        </Link>
      </div>
    </>
  )
}
