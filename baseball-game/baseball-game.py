class Solution:
    def calPoints(self, ops: List[str]) -> int:
        while ops.count("C"):
            ops.pop(ops.index("C")-1)
            ops.remove("C")
        for i in range(len(ops)):
            if ops[i] == "D":
                ops[i] = int(ops[i-1])*2
            elif ops[i] == "+":
                ops[i] = int(ops[i-1]) + int(ops[i-2])
        ops = list(map(int,ops))
        return sum(ops)
    