import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
   <div className='w-[100dvw] h-[90dvh] border border-red-200'>
    <h1 className='m-auto'>Ballot Box</h1>
   </div>
    </>
  )
}

export default App
