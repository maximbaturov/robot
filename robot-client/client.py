import websocket
import _thread
import time
import rel
import json
from robot.robot import *

def on_message(ws, message):
    try:
        json_message = json.loads(message)

        if json_message["action"] == "keydown":
            if json_message["button"] == "ArrowUp":
                do_forward()

            if json_message["button"] == "ArrowDown":
                do_backward()

            if json_message["button"] == "ArrowLeft":
                do_left()

            if json_message["button"] == "ArrowRight":
                do_right()
        
        if json_message["action"] == "keyup":
            do_stop()

    except BaseException as e:
        do_stop()

def on_error(ws, error):
    print(error)

def on_close(ws, close_status_code, close_msg):
    print("### closed ###")

def on_open(ws):
    ws.send(json.dumps({
            "r": "robot_test_1"
        }))
    run()

if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("ws://192.168.0.115:9098/websocket",
                              on_open=on_open,
                              on_message=on_message,
                              on_error=on_error,
                              on_close=on_close)

    ws.run_forever(dispatcher=rel, reconnect=5)  # Set dispatcher to automatic reconnection, 5 second reconnect delay if connection closed unexpectedly
    rel.signal(2, rel.abort)  # Keyboard Interrupt
    rel.dispatch()