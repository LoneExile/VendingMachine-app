import Link from 'next/link'

interface FooterProps {
  status: number
}

export default function Footer({status}: FooterProps) {
  const stepClass = (stepNumber: number) =>
    stepNumber <= status ? 'step step-primary' : 'step'

  return (
    <footer className="footer footer-center p-4 bg-base-300 text-base-content">
      <div>
        <ul className="steps steps-vertical lg:steps-horizontal ">
          <li className={stepClass(1)}>Choose product</li>
          <li className={stepClass(2)}>Checkout</li>
          <li className={stepClass(3)}>Purchase</li>
        </ul>
      </div>
    </footer>
  )
}

// interface StepProps {
//   stepNumber: number
//   href: string
//   children: React.ReactNode
// }

// export default function Footer({status}: FooterProps) {
//   const stepClass = (stepNumber: number) =>
//     stepNumber <= status ? 'step step-primary' : 'step'

//   const Step: React.FC<StepProps> = ({stepNumber, href, children}) => {
//     if (stepNumber <= status) {
//       return (
//         <Link href={href}>
//           <li className={stepClass(stepNumber)}>{children}</li>
//         </Link>
//       )
//     }
//     return <li className={stepClass(stepNumber)}>{children}</li>
//   }

//   return (
//     <footer className="footer footer-center p-4 bg-base-300 text-base-content">
//       <div>
//         <ul className="steps steps-vertical lg:steps-horizontal">
//           <Step stepNumber={1} href="/">
//             Choose product
//           </Step>
//           <Step stepNumber={2} href="/Checkout">
//             Checkout
//           </Step>
//           <Step stepNumber={3} href="/payment">
//             Purchase
//           </Step>
//         </ul>
//       </div>
//     </footer>
//   )
// }
