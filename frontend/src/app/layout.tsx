// 'use client'
import './globals.css'
import {Inter} from 'next/font/google'

const inter = Inter({subsets: ['latin']})

export const metadata = {
  title: 'Vendor Machine',
  description: 'Vendor Machine',
}

export default function RootLayout({children}: {children: React.ReactNode}) {
  return (
    <html lang="en" data-theme="light">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
