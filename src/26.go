import numpy as np

class Go():
    def __init__(self):
        self.board = np.array([['1', '2', '3'],
                               ['4', '5', '6'],
                               ['7', '8', '9']])

    def move(self, player, position):
        row, col = int(position[0]), int(position[1]) - 1
        if self.board[row][col] == '.':
            return 'Game Over'
        else:
            self.board[row][col] = '.'
            winner = self.check_winner()
            if winner != 'Draw':
                return f'Player {player} wins!'
            elif player == 1:
                return 'Player 1 wins!'

    def check_winner(self):
        for i in range(3):
            row = 0
            while True:
                row += 1
                if (self.board[row][i] == self.board[row][i + 1] and self.board[row][i] != '.' or
                    (self.board[row][i + 1] == self.board[row][i + 2] and self.board[row][i + 1] != '.'
                     and row >= 0 and row < 3)):
                    return 'Player 1 wins!'
                if (row == 3):
                    break
        for i in range(3):
            col = 0
            while True:
                col += 1
                if (self.board[i][col] == self.board[i + 1][col] and self.board[i][col] != '.' or
                    (self.board[i + 1][col] == self.board[i + 2][col] and self.board[i + 1][col] != '.'
                     and col >= 0 and col < 3)):
                    return 'Player 2 wins!'
                if (col == 3):
                    break
        for i in range(3):
            if ((self.board[0][i] == self.board[1][i] and self.board[0][i] != '.' or
                 (self.board[1][i] == self.board[2][i] and self.board[1][i] != '.'
                  and i >= 0 and i < 3))):
                return 'Player 1 wins!'
        if ((self.board[0][0] == self.board[1][1] and self.board[0][0] != '.' or
             (self.board[1][1] == self.board[2][2] and self.board[1][1] != '.'
              and i >= 0 and i < 3))):
            return 'Player 1 wins!'
        if ((self.board[0][2] == self.board[1][1] and self.board[0][2] != '.' or
             (self.board[1][1] == self.board[2][0] and self.board[1][1] != '.'
              and i >= 0 and i < 3))):
            return 'Player 1 wins!'
        return 'Draw'

go = Go()
print(go.move(1, '5'))
