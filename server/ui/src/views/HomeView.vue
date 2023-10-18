<script setup>
import TheWelcome from '../components/TheWelcome.vue'
</script>

<template>
  <main>
    <!-- <TheWelcome /> -->
    <video id="video1" width="500" height="500" autoplay muted></video> <br />

    Logs<br />
    <div id="logs"></div>
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

    function log(msg) {
      document.getElementById('logs').innerHTML += msg + '<br>'
    }

    function startSession(sd) {
      // const sd = document.getElementById('remoteSessionDescription').value
      if (sd === '') {
        return alert('Session Description must not be empty')
      }

      try {
        pc.setRemoteDescription(JSON.parse(atob(sd)))
      } catch (e) {
        alert(e)
      }
    }

    (function () {
      const pc = new RTCPeerConnection({
        iceServers: [
          {
            urls: 'stun:stun.l.google.com:19302'
          }
        ]
      })
      pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
      pc.onicecandidate = event => {
        if (event.candidate === null) {
          let xhr = new XMLHttpRequest();
          xhr.open("POST", `http://${host}:${port}/stream`);
          xhr.setRequestHeader("Accept", "application/json");
          xhr.setRequestHeader("Content-Type", "application/json");

          xhr.onreadystatechange = function () {
            if (xhr.readyState === 4) {
              console.log(xhr.status);
              console.log(xhr.responseText);
            }};

          xhr.send(JSON.stringify({ "data": btoa(JSON.stringify(pc.localDescription)) }));

        }
      }

      pc.addTransceiver('video')
      pc.createOffer()
          .then(d => pc.setLocalDescription(d))
          .catch(log)

      pc.ontrack = function (event) {
        const el = document.getElementById('video1')
        el.srcObject = event.streams[0]
        el.autoplay = true
        el.controls = true
        el.style.width = "500px"
        el.style.height = "500px"
      }

      // startSession()

      // const btns = document.getElementsByClassName('createSessionButton')
      // for (let i = 0; i < btns.length; i++) {
      //   btns[i].style = 'display: none'
      // }
      //
      // document.getElementById('signalingContainer').style = 'display: block'
      //
    })()

</script>
