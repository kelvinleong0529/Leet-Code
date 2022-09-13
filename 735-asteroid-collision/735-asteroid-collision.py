class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        stack = []
        ans = []
        for a in asteroids:
            if a > 0:
                stack.append(a)
            else:
                same = False
                while stack and abs(a) >= stack[-1]:
                    if stack and abs(a) == stack[-1]:
                        stack.pop()
                        same = True
                        break
                    stack.pop()
                if not stack and not same:
                    ans.append(a)
            
        ans.extend(stack)
        return ans