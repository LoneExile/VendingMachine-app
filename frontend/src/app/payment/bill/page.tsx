'use client'
import {useStore} from '@nanostores/react'
import {cartStorage} from '@/utils/stores'
import Link from 'next/link'
import Footer from '@/components/footer'
import Nav from '@/components/nav'

export default function Page() {
  cartStorage.setKey('products', JSON.stringify([]))
  cartStorage.setKey('pocket', JSON.stringify([]))

  // const local = useStore(cartStorage)
  const local = cartStorage.get()
  const bills: Bill = JSON.parse(local.bills)
  console.log('bills is: ', bills.message)
  const sum = bills.message?.change.reduce(
    (total, item) => total + item.denomination_value * item.quantity,
    0
  )

  function ChangeList() {
    return (
      <div className="overflow-x-auto">
        <table className="table">
          {/* head */}
          <thead>
            <tr>
              <th></th>
              <th>Denomination</th>
              <th>Quantity</th>
              <th>Total</th>
            </tr>
          </thead>
          <tbody>
            {bills.message?.change.map((item, index) => (
              <tr key={index}>
                <th>{index + 1}</th>
                <td>
                  {' '}
                  {item.typed === 'coin' ? '🪙' : '💵'}{' '}
                  {item.denomination_value}
                </td>
                <td>{item.quantity}</td>
                <td>{item.denomination_value * item.quantity}</td>
              </tr>
            ))}
          </tbody>
        </table>
        <div className="card-body">
          <span className="font-bold ">Total Change: {sum}</span>
        </div>
      </div>
    )
  }

  const status = 3

  return (
    <div className="flex flex-col min-h-screen justify-between">
      <Nav status={status} location="Payment" />
      <main className="container mx-auto mb-auto  px-12">
        <div className="hero min-h-screen bg-base-200">
          <div className="hero-content text-center">
            <div className="max-w-md">
              <h1
                className={
                  !(bills.message?.status === 'success')
                    ? 'text-5xl font-bold uppercase text-red-500'
                    : 'text-5xl font-bold uppercase text-green-500'
                }
              >
                {bills.message?.status}
              </h1>
              <p
                className={
                  !(bills.message?.status === 'success')
                    ? 'py-6 text-red-500'
                    : 'py-6 text-green-500'
                }
              >
                {bills.message?.status === 'success'
                  ? 'Thank you for your purchase!'
                  : 'Sorry, your payment was not successful.'}
              </p>
              <div>
                <ChangeList />
              </div>
              <Link href="/">
                <button className="btn btn-primary">Done</button>
              </Link>
            </div>
          </div>
        </div>
      </main>
      <Footer status={status} />
    </div>
  )
}
