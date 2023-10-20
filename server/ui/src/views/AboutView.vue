<style>
  textarea {
    width: 500px;
    min-height: 75px;
  }
  video {
    width: 500px;
    height: 500px;
  }
</style>

<template>
  About
</template>
<!--
	SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
	SPDX-License-Identifier: MIT
-->
<!-- <div id="signalingContainer" style="display: none">
  Browser base64 Session Description<br />
  <textarea id="localSessionDescription" readonly="true"></textarea> <br />
  <button>
    Copy browser SDP to clipboard
  </button>
  <br />
  <br />

  Golang base64 Session Description<br />
  <textarea id="remoteSessionDescription"></textarea> <br/>
  <button onclick="window.startSession()"> Start Session </button><br />
</div>

<br />

Video<br />
<video id="video1" width="160" height="120" autoplay muted></video> <br /> -->

<!-- <button class="createSessionButton" onclick="window.createSession(true)"> Publish a Broadcast </button>
<button class="createSessionButton" onclick="window.createSession(false)"> Join a Broadcast </button><br /> -->

<!-- <br />

Logs<br />
<div id="logs"></div> -->

<script>
  /* eslint-env browser */

  // SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
  // SPDX-License-Identifier: MIT

  // const log = msg => {
  //   document.getElementById('logs').innerHTML += msg + '<br>'
  // }

  // window.createSession = isPublisher => {
  //   const pc = new RTCPeerConnection({
  //     iceServers: [
  //       {
  //         urls: 'stun:stun.l.google.com:19302'
  //       }
  //     ]
  //   })
  //   pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
  //   pc.onicecandidate = event => {
  //     if (event.candidate === null) {
  //       const host = '127.0.0.1';
  //       const port = '9097';

  //       let xhr = new XMLHttpRequest();

  //       xhr.open("POST", `http://${host}:${port}/stream`);
  //       xhr.setRequestHeader("Accept", "application/json");
  //       xhr.setRequestHeader("Content-Type", "application/json");

  //       xhr.onreadystatechange = function () {
  //         if (xhr.readyState === 4) {
  //           console.log(xhr.status);
  //           console.log(xhr.responseText);
  //         }};

  //       xhr.send(JSON.stringify({ "data": btoa(JSON.stringify(pc.localDescription)) }));

  //       document.getElementById('localSessionDescription').value = btoa(JSON.stringify(pc.localDescription))
  //     }
  //   }

  //   if (isPublisher) {
  //     navigator.mediaDevices.getUserMedia({ video: true, audio: false })
  //         .then(stream => {
  //           stream.getTracks().forEach(track => pc.addTrack(track, stream))
  //           document.getElementById('video1').srcObject = stream
  //           pc.createOffer()
  //               .then(d => pc.setLocalDescription(d))
  //               .catch(log)
  //         }).catch(log)
  //   } else {
  //     pc.addTransceiver('video')
  //     pc.createOffer()
  //         .then(d => pc.setLocalDescription(d))
  //         .catch(log)

  //     pc.ontrack = function (event) {
  //       const el = document.getElementById('video1')
  //       el.srcObject = event.streams[0]
  //       el.autoplay = true
  //       el.controls = true
  //       el.style.width = "500px"
  //       el.style.height = "500px"
  //     }
  //   }

  //   window.startSession = () => {
  //     const sd = document.getElementById('remoteSessionDescription').value
  //     if (sd === '') {
  //       return alert('Session Description must not be empty')
  //     }

  //     try {
  //       pc.setRemoteDescription(JSON.parse(atob(sd)))
  //     } catch (e) {
  //       alert(e)
  //     }
  //   }


  //   const btns = document.getElementsByClassName('createSessionButton')
  //   for (let i = 0; i < btns.length; i++) {
  //     btns[i].style = 'display: none'
  //   }

  //   document.getElementById('signalingContainer').style = 'display: block'
  // }

</script>