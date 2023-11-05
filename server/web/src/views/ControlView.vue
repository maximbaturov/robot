
<template>
  <main>
  </main>
</template>

<script>
    const clientID = 'test_client_1' //urlParams.get('c');

    const host = '127.0.0.1';   
    const port = '9097';
    
    const socket = new WebSocket(`ws://${host}:${port}/websocket`);

    const keys = [
        'ArrowUp',
        'ArrowDown',    
        'ArrowLeft',
        'ArrowRight',
    ];

    var lastKeyDown = null

    socket.addEventListener('open', function (event) {
        console.log("connected")
        socket.send(JSON.stringify({c: clientID}));
    });

    socket.onclose = function () {
        console.log("Connection closed");
        console.log('WebSocket disconnected');
    };

    socket.addEventListener('message', (event) => {
        console.log("message from server " + event.data);
    });

    document.addEventListener("keyup", (event) => {
        if(keys.indexOf(event.key) !== -1) {
            lastKeyDown = null;
            let data = {
                button: event.key,
                action: "keyup",
                c: clientID
            }

            console.log(data)

            socket.send(JSON.stringify(data));
        }
    })

    document.addEventListener("keydown",(event) => {

        if (keys.indexOf(event.key) !== -1 && lastKeyDown != event.key) {
            lastKeyDown = event.key
            let data = {
                button: event.key,
                action: "keydown",
                c: clientID
            }

            console.log(data)

            socket.send(JSON.stringify(data));
        }
    });

    

</script>
