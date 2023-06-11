'use client'
import Link from 'next/link'
import dynamic from 'next/dynamic'

const ItemsTable = dynamic(() => import('@/components/itemsTable'), {
  ssr: false,
})

function Items() {
  return (
    <>
      <ItemsTable />

      {/* <div className="card shadow-lg bg-base-100 my-4">
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
      </div> */}

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

export {Items, ItemsTable}
