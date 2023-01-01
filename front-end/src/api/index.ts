let socket = new WebSocket("ws://localhost:8080/ws")
let connect= (cb:any)=>{
  console.log("connecting")
  socket.onopen=()=>{
    console.log("Connected")
  }
  socket.onmessage = msg=>{
    console.log(msg)
    cb(msg)
  }
  socket.onclose = event=>{
    console.log("Closed because of ", event)
  }
  socket.onerror = error=>{
    console.log("Error: ",error)
  }
}
let sendMsg= (msg: string | ArrayBufferLike | Blob | ArrayBufferView) =>{
  console.log("message: ",msg)
  socket.send(msg)
}
export{sendMsg, connect}