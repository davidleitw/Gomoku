import os
import logging
from pickletools import uint8
import socket

HEADER_LEN = 7
HEADER = "Headers"

class IpcServer():
    def __init__(self, sock_path: str):
        FORMAT = '%(asctime)s %(levelname)s: %(message)s'
        logging.basicConfig(level=logging.DEBUG, format=FORMAT)
        logging.info("Create unix socket server...")
        
        if os.path.exists(sock_path):
            os.remove(sock_path)

        sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
        if sock.fileno() < 0:
            print("socket create error!")
            return

        self.sock = sock
        self.sock.bind(sock_path)
        self.sock.listen(1)
        self.conn, _ = self.sock.accept()
        
        while True:
            data = self.conn.recv(572)
            self.get_candiates(bytearray(data))
            self.conn.send(data)
            
    def get_candiates(self, packet: bytearray) ->list:
        length = len(packet)
        if length < HEADER_LEN or packet[0:HEADER_LEN] == "Headers":
            return

        num = int(packet[HEADER_LEN])
        if length != HEADER_LEN + 1 + num * 2:
            return

        points = []
        index = HEADER_LEN + 1
        while index < length:
            x = packet[index]
            index = index + 1
            y = packet[index]
            index = index + 1
            points.append({x, y})

        print(points)

    def send_decision(self, x: int, y: int) -> None:
        pass