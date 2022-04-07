import os
import socket
import numpy
import logging

from .Board import Board

HEADER_LEN = 1
HEADER = "H"

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
        self.board = Board()
        
        while True:
            data = self.conn.recv(452)
            self.get_candiates(bytearray(data))
            self.conn.send(data)
            
    def get_candiates(self, packet: bytearray) ->list:
        length = len(packet)
        num = int(packet[HEADER_LEN])
        print(length, num, packet)

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