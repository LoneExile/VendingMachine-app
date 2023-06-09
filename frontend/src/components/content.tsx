import Link from 'next/link'

function Card() {
  return (
    <div className="card card-compact w-52 m-4 bg-base-100 shadow-xl">
      <figure>
        <img src="https://illustoon.com/photo/11612.png" alt="Shoes" />
      </figure>
      <div className="card-body">
        <h2 className="card-title">Water</h2>
        <div className="card-actions justify-end">
          <button className="btn btn-primary">Add</button>
        </div>
      </div>
    </div>
  )
}
export default function Connent() {
  return (
    <>
      <div className="flex flex-wrap justify-center">
        {Array.from({length: 12}, (_, i) => (
          <Card key={i} />
        ))}
      </div>
      <div className="flex justify-end m-4">
        <Link href="/checkout">
          <button className="btn btn-primary">Next</button>
        </Link>
      </div>
    </>
  )
}
