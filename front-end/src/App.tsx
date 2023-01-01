import { useState } from 'react'
import { connect, sendMsg } from './api'
import { Header } from './components/Header/Header'
import './App.css'
import { ChatHistory } from './components/ChatHistory/ChatHistory'

function App() {
  const[chatHistory, setChatHistory] = useState()
  // connect()
  function send(){
    sendMsg("hello")
  }
  function mounted(){
    connect((msg:any)=>{
      setChatHistory(msg)
    })
  }
  return (
    <div className="App">
      <Header/>
      <ChatHistory chatHistory={chatHistory}/>
      <button onClick={send}></button>
    </div>
  )
}

export default App
