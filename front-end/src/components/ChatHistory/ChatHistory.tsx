import React from "react";
import '../chatHistory.css'
export function ChatHistory(chatHistory:any){
  const messages = chatHistory.map((msg:any, index:number) => (
    <p key={index}>{msg.data}</p>
  ));
  return(
    <div className="ChatHistory">
      {messages}
    </div>
  )
}