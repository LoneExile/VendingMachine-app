import Nav from '@/components/nav'
import Footer from '@/components/footer'
import Payment from '@/components/payment'

const status = 3

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen justify-between">
      <Nav status={status} location="Payment" />
      <main className="container mx-auto mb-auto mt-8 px-12">
        <Payment />
      </main>
      <Footer status={status} />
    </div>
  )
}
