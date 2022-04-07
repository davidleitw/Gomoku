import numpy as np

BLACKCODE=1
WHITECODE=2    

PlayerCode = (BLACKCODE, WHITECODE)

class Board():
    def __init__(self):
        self.state = np.zeros((15, 15), dtype=int)
    
    def step(self, player: int, x: int, y: int):
        if self.state[x][y] == 0:
            self.state[x][y] = player
    
    def recv_candiates(self, packet: bytearray) -> np.ndarray:
        length = len(packet)
        pc = packet[0]
        
        index = 1
        candiates = np.zeros((3, 15, 15), dtype=int)
        while index < length:
            x, y = packet[index], packet[index+1]
            index += 2
            candiate = self.state.copy()
            candiate[x][y] = pc
            candiates[0] = candiate
        
        
        