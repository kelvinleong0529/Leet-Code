class Solution:
    def judgeCircle(self, moves: str) -> bool:
        if "U" or "D" in moves:
            if moves.count("U") != moves.count("D"):
                return 0
        if "L" or " R" in moves:
            if moves.count("L") != moves.count("R"):
                return 0
        return 1
