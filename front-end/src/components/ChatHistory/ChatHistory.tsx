import React from "react";
import '../chatHistory.scss'
import Message from "../Message/Message";
type Props={
  chatHistory: any
}
export function ChatHistory(props:Props){
  const history = props.chatHistory
  const message = history?.map((msg:any,index:number) => <><p key={index}>{msg.data}</p></>)
  return(
    <div className="ChatHistory">
      <Message message={message}/>
    </div>
  )
}