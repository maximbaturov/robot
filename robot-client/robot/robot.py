import RPi.GPIO as GPIO
import logging
from time import sleep

in1 = 17
in2 = 22
in3 = 23
in4 = 24

en = 25

temp1=1

def run():
    GPIO.setmode(GPIO.BCM)
    GPIO.setup(in1,GPIO.OUT)
    GPIO.setup(in2,GPIO.OUT)
    GPIO.setup(in3,GPIO.OUT)
    GPIO.setup(in4,GPIO.OUT)

    GPIO.setup(en,GPIO.OUT)

    GPIO.output(in1,GPIO.LOW)
    GPIO.output(in2,GPIO.LOW)
    GPIO.output(in3,GPIO.LOW)
    GPIO.output(in4,GPIO.LOW)

    p=GPIO.PWM(en,1000)
    p.start(25)

def do_forward():
    GPIO.output(in1,GPIO.HIGH)
    GPIO.output(in2,GPIO.LOW)

    GPIO.output(in3,GPIO.HIGH)
    GPIO.output(in4,GPIO.LOW)
    

def do_backward(): 
    GPIO.output(in1,GPIO.LOW)
    GPIO.output(in2,GPIO.HIGH)

    GPIO.output(in3,GPIO.LOW)
    GPIO.output(in4,GPIO.HIGH)


def do_left():
    GPIO.output(in1,GPIO.HIGH)
    GPIO.output(in2,GPIO.LOW)

    GPIO.output(in3,GPIO.LOW)
    GPIO.output(in4,GPIO.HIGH)


def do_right():
    GPIO.output(in1,GPIO.LOW)
    GPIO.output(in2,GPIO.HIGH)

    GPIO.output(in3,GPIO.HIGH)
    GPIO.output(in4,GPIO.LOW)

def do_stop():
    GPIO.output(in1,GPIO.LOW)
    GPIO.output(in2,GPIO.LOW)

    GPIO.output(in3,GPIO.LOW)
    GPIO.output(in4,GPIO.LOW)

def do_exit():
    GPIO.cleanup()

# class S(BaseHTTPRequestHandler):
#     def _set_response(self):
#         self.send_response(200)
#         self.send_header('Content-type', 'text/html')
#         self.end_headers()

#     def do_GET(self):
#         if self.path == '/forward':
#             do_forward(self)
#         elif self.path == '/backward':
#             do_backward(self)
#         elif self.path == '/left':
#             do_left(self)
#         elif self.path == '/right':
#             do_right(self)
#         elif self.path == '/stop':
#             do_stop(self)

# def run(server_class=HTTPServer, handler_class=S, port=8080):
#     logging.basicConfig(level=logging.INFO)
#     server_address = ('', port)
#     httpd = server_class(server_address, handler_class)
#     logging.info('Starting httpd...\n')
#     try:
#         httpd.serve_forever()
#     except KeyboardInterrupt:
#         pass
#     httpd.server_close()
#     logging.info('Stopping httpd...\n')


# run(port=8080)
