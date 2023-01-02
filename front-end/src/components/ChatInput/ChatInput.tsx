import React from 'react';
import '../chatinput.scss'
type Props ={
  send:(event:any)=>void
}
function ChatInput(props:Props) {
  return (
    <div>
      <input type="text" onKeyDown={props.send} />
    </div>
  );
}

export default ChatInput;