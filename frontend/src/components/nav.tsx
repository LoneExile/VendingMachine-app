'use client'
import {ItemsTable} from '@/components/items'
import Link from 'next/link'

import {useStore} from '@nanostores/react'
import {cartStorage} from '@/utils/stores'

interface Props {
  status: number
  location: string
}
interface ItemsProps {
  status: number
  // children: React.ReactNode
}

// function useTotal() {
//   const local = useStore(cartStorage).products
//   // const local = cartStorage.get().products
//   const $cart: Cart = JSON.parse(local)

//   let total = 0
//   $cart.items?.forEach((item) => {
//     total += item.quantity
//   })
//   return total
// }

export default function Nav({status, location}: Props) {
  const local = useStore(cartStorage).products
  // const local = cartStorage.get().products
  const $cart: Cart = JSON.parse(local)

  let total = 0
  $cart.items?.forEach((item) => {
    total += item.quantity
  })

  const ItemsNav: React.FC<ItemsProps> = ({status}) => {
    if (status === 1) {
      return (
        <div className="dropdown dropdown-end">
          <label tabIndex={0} className="btn btn-ghost btn-circle">
            <div className="indicator">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                />
              </svg>
              <span className="badge badge-sm indicator-item">{total}</span>
            </div>
          </label>
          <div
            tabIndex={0}
            className="mt-3 card card-compact dropdown-content w-[25rem] bg-base-100 shadow"
          >
            <ItemsTable />
          </div>
        </div>
      )
    } else {
      return <div></div>
    }
  }

  return (
    <>
      <div
        className="navbar bg-base-300"
        style={{position: 'relative', zIndex: 1}}
      >
        <div className="navbar-start">
          {/* <div className="drawer">
            <input id="my-drawer" type="checkbox" className="drawer-toggle" />
            <div className="drawer-content">
              <label
                htmlFor="my-drawer"
                className="btn btn-primary drawer-button"
              >
                Menu
              </label>
            </div>
            <div className="drawer-side">
              <label htmlFor="my-drawer" className="drawer-overlay"></label>
              <ul className="menu p-4 w-80 h-full bg-base-200 text-base-content">
                <li>
                  <a>Sidebar Item 1</a>
                </li>
                <li>
                  <a>Sidebar Item 2</a>
                </li>
              </ul>
            </div>
          </div> */}
          <Link href="/" className="btn btn-ghost normal-case text-xl">
            Logo
          </Link>
        </div>
        <div className="navbar-center">
          {/* <a className="btn btn-ghost normal-case text-xl hidden lg:block"> */}
          <a className="btn btn-ghost normal-case text-xl">{location}</a>
        </div>
        <div className="navbar-end">
          <ItemsNav status={status} />
        </div>
      </div>
    </>
  )
}
