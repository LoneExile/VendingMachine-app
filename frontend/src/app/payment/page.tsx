'use client'
import Nav from '@/components/nav'
import Footer from '@/components/footer'
import dynamic from 'next/dynamic'

const Payment = dynamic(() => import('@/components/payment'), {ssr: false})

const status = 3

async function getData() {
  const res = await fetch('http://localhost:8080/denomination', {
    cache: 'no-store',
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return res.json()
}

export default async function Home() {
  const data: DenominationsResponse = await getData()
  return (
    <div className="flex flex-col min-h-screen justify-between">
      <Nav status={status} location="Payment" />
      <main className="container mx-auto mb-auto mt-8 px-12">
        <Payment denominations={data?.denominations} />
      </main>
      <Footer status={status} />
    </div>
  )
}
