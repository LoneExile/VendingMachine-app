import Nav from '@/components/nav'
import Footer from '@/components/footer'
import {Items} from '@/components/items'

const status = 2

export default function Checkout() {
  return (
    <div className="flex flex-col min-h-screen justify-between">
      <Nav status={status} location="Checkout" />
      <main className="container mx-auto mb-auto mt-8 px-12">
        <Items />
      </main>
      <Footer status={status} />
    </div>
  )
}
