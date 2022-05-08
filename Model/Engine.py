import numpy as np

BLACKCODE=1
WHITECODE=2    
PlayerCode = (BLACKCODE, WHITECODE)

class ReversiEngine():
    def __init__(self, board_size: int, debug_mode: bool=False, epoch: int=100):
        self.epoch = epoch
        self.size = board_size
        self.debug_mode = debug_mode

    def reset_reversi_board(self):
        self.player = BLACKCODE
        self.remain = self.size * self.size
        self.state = np.zeros((self.size, self.size), dtype=int)

    def step(self, x: int, y: int):
        if self.state[x][y] == 0:
            self.state[x][y] = self.player
        else:
            return

        # 判斷下完棋之後翻面的邏輯
        self.judge(x, y, -1, 0)
        self.judge(x, y, -1, -1)
        self.judge(x, y, 0, -1)
        self.judge(x, y, 1, -1)
        self.judge(x, y, 1, 0)
        self.judge(x, y, 1, 1)
        self.judge(x, y, 0, 1)
        self.judge(x, y, -1, 1)
        self.change_player()
        self.remain -= 1
        
    def judge(self, x: int, y: int, dx: int, dy: int):
        route = []
        while True:
            x += dx
            y += dy

            if self.out_range(x) or self.out_range(y):
                return
            
            if self.state[x][y] == 0:
                return
            elif self.state[x][y] == self.player:
                break

            route.append({x, y})
        
        for x, y in route:
            self.state[x][y] = self.player

    def change_player(self):
        self.player = self.player%2 + 1

    def out_range(self, n: int):
        return n < 0 or n > self.size - 1

    def game_over(self) -> bool:
        return self.remain == 0

    def Run(self):
        for epoch in range(self.epoch):
            self.reset_reversi_board()
            while not self.game_over():
                # 所有可能的結果
                candiates = self.get_candiates()
                print(candiates.shape)
                # Model 選擇一個最好的棋步
                # x, y = model.select(candiates)
                # self.step(x, y)


    def get_candiates(self) -> np.ndarray:
        ptr = 0
        candiates = np.zeros((self.remain, self.size, self.size), dtype=int)
        for x in range(self.size):
            for y in range(self.size):
                if self.state[x][y] == 0:
                    candiate = self.state.copy()
                    candiate[x][y] = self.player
                    candiates[ptr] = candiate
                    ptr += 1
        return candiates
    