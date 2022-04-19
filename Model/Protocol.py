import os
import socket
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
    
    def recv(self) -> bytes:
        return self.conn.recv(452)
    
    def send(self, decision: bytes) -> None:
        self.conn.send(decision)

    def run(self) -> None:
        while True:
            actions = self.board.get_candiates(self.recv())
            print(actions.shape)
        
    def demo(self) -> None:
        while True:
            actions = self.board.get_candiates(self.conn.recv(452))
            print(actions.shape)
            x, y = map(lambda x: int(x), input("Please input x, y:").split())
            self.board.step(x, y)
            self.send(bytes([x, y]))