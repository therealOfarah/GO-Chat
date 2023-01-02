import React from 'react';
import '../messages.scss'
type Props={
  message:string
}
function Message(props:Props) {
  // let temp = JSON.parse(props.message)
  return (
    <div>
      {props.message}
    </div>
  );
}

export default Message;