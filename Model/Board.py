import numpy as np

class Board():
    def __init__(self):
        self.state = np.zeros((15, 15), dtype=int)
    
    def step(self, player: int, x: int, y: int):
        if self.state[x][y] == 0:
            self.state[x][y] = player
    
    def recv_candiates(self, packet: bytearray) -> np.ndarray:
        length = len(packet)
        candiates_num = int(packet[1])
        assert(length == candiates_num * 2 + 2)
        
        index = 2
        while index < length:
            candiate = self.state.copy()
            x = packet[index]
            y = packet[index+1]
            candiate[x][y] = player
            index = index + 2