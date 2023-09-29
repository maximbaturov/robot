import datetime
import tornado.httpserver
import tornado.websocket
import tornado.ioloop
import tornado.web
from robot.robot import *
import json

class WSHandler(tornado.websocket.WebSocketHandler):
    clients = []
    def open(self):
        print('new connection')
        self.write_message("Hello World")
        WSHandler.clients.append(self)

    def on_message(self, message):
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

        self.write_message('ECHO: ' + message)


    def on_close(self):
        print('connection closed')
        WSHandler.clients.remove(self)

    @classmethod
    def write_to_clients(cls):
        print("Writing to clients")
        for client in cls.clients:
            client.write_message("Hi there!")

    def check_origin(self, origin):
        return True



application = tornado.web.Application([
  (r'/websocket', WSHandler),
])

if __name__ == "__main__":
    do_exit()
    run()
    http_server = tornado.httpserver.HTTPServer(application)
    http_server.listen(9098)
    tornado.ioloop.IOLoop.instance().start()