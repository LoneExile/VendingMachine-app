'use client'
import Link from 'next/link'
import {useState} from 'react'

function ListItem() {
  const [count, setCount] = useState(0)

  const increase = () => {
    setCount((prevCount) => prevCount + 1)
  }

  const decrease = () => {
    if (count > 0) {
      setCount((prevCount) => prevCount - 1)
    }
  }
  return (
    <tr>
      <td>
        <div className="flex items-center space-x-3">
          <div className="avatar">
            <div className="mask mask-squircle w-12 h-12">🪙</div>
          </div>
          <div>
            <div className="font-bold">100THB</div>
          </div>
        </div>
      </td>
      <td>
        <div className="flex items-center space-x-2">
          <button className="btn btn-primary" onClick={decrease}>
            -
          </button>

          <input
            type="number"
            className="input input-bordered w-20"
            value={count}
            readOnly
          />

          <button className="btn btn-primary" onClick={increase}>
            +
          </button>
        </div>
      </td>
      <td>1000</td>
    </tr>
  )
}

export default function Payment() {
  return (
    <>
      <div className="collapse collapse-arrow bg-base-200 mb-2">
        <input type="radio" name="my-accordion-2" defaultChecked={true} />
        <div className="collapse-title text-xl font-medium">
          <div className="flex justify-between">
            <div>Banknote</div>
            <div>1000$</div>
          </div>
        </div>
        <div className="collapse-content">
          <div className="overflow-x-auto">
            <table className="table">
              <thead>
                <tr>
                  <th>Item</th>
                  <th>Quantity</th>
                  <th>Total</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {Array.from({length: 5}, (_, i) => (
                  <ListItem key={i} />
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <div className="collapse collapse-arrow bg-base-200">
        <input type="radio" name="my-accordion-1" defaultChecked={true} />
        <div className="collapse-title text-xl font-medium">
          <div className="flex justify-between">
            <div>Coin</div>
            <div>1000$</div>
          </div>
        </div>
        <div className="collapse-content">
          <div className="overflow-x-auto">
            <table className="table">
              <thead>
                <tr>
                  <th>Item</th>
                  <th>Quantity</th>
                  <th>Total</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                {Array.from({length: 3}, (_, i) => (
                  <ListItem key={i} />
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <div className="card shadow-lg bg-base-100 my-4">
        <div className="card-body">
          <div className="flex justify-between">
            <div>Total Payment</div>
            <div>1000$</div>
          </div>
          <div className="flex justify-between">
            <div>Curruent</div>
            <div>100$</div>
          </div>
          <div className="flex justify-between">
            <div>Change</div>
            <div>1100$</div>
          </div>
        </div>
      </div>
      <div className="flex justify-between m-4">
        <Link href="/checkout">
          <button className="btn btn-primary">Back</button>
        </Link>

        <button className="btn btn-primary">Pay</button>
      </div>
    </>
  )
}
