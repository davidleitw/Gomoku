import numpy as np

BLACKCODE=1
WHITECODE=2    
PlayerCode = (BLACKCODE, WHITECODE)

class Board():
    def __init__(self):
        self.state = np.zeros((15, 15), dtype=int)
        self.pc = 0
    
    def step(self, x: int, y: int):
        if self.state[x][y] == 0:
            self.state[x][y] = self.pc
    
    def get_candiates(self, packet: bytes) -> np.ndarray:
        self.pc = packet[0] # player_code  (輪到誰下棋)
        cnt = packet[1]     # candiate_num (共有幾組候選)
        candiates = np.zeros((cnt, 15, 15), dtype=int)
        
        cid   = 0
        idx = 2
        while cid < cnt:
            candiate = self.state.copy()
            candiate[packet[idx]][packet[idx+1]] = self.pc
            candiates[cid] = candiate
            cid = cid + 1
            idx = idx + 2
        return candiates