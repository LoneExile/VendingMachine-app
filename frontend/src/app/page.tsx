import Footer from '@/components/footer'
import dynamic from 'next/dynamic'

const Nav = dynamic(() => import('@/components/nav'), {ssr: false})
const Connent = dynamic(() => import('@/components/content'), {ssr: false})

const status = 1

async function getData() {
  const res = await fetch(process.env.NEXT_PUBLIC_SERVER_IP + '/products', {
    cache: 'no-store',
    credentials: 'include',
  })
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
  return res.json()
}

export default async function Page() {
  const data: ProductsResponse = await getData()
  return (
    <div className="flex flex-col min-h-screen justify-between">
      <Nav status={status} location="Choose product" />
      <main className="mt-2">
        <Connent products={data?.products} />
      </main>
      <Footer status={status} />
    </div>
  )
}
