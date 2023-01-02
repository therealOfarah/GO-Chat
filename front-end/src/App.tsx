import { useState } from 'react'
import { connect, sendMsg } from './api'
import { Header } from './components/Header/Header'
import './App.css'
import { ChatHistory } from './components/ChatHistory/ChatHistory'
import ChatInput from './components/ChatInput/ChatInput'

function App() {
  const[chatHistory, setChatHistory] = useState<any>()
  // connect()
  function send(event:any){
    if(event.keyCode === 13) {
    sendMsg(event.target.value);
    event.target.value = "";
  }
  }

  connect((msg:any)=>{
    return setChatHistory([msg])
  })
  console.log(chatHistory)
  return (
    <div className="App">
      <Header/>
      <ChatInput send={send}/>
      <ChatHistory chatHistory={chatHistory}/>
    </div>
  )
}

export default App
