import {Nav} from '@/components/nav'
import Footer from '@/components/footer'
import Payment from '@/components/payment'

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen justify-between">
      <Nav />
      <main className="container mx-auto mb-auto mt-8">
        <Payment />
      </main>
      <Footer />
    </div>
  )
}
