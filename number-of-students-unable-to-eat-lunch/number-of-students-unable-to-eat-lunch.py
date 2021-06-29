class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        student_count = students.count(0)
        sandwiches_count = sandwiches.count(0)
        diff = student_count - sandwiches_count
        if not diff:
            return 0
        elif diff > 0:
            student_count = students.count(1)
            while student_count > 0:
                index = sandwiches.index(1)
                sandwiches = sandwiches[index+1:]
                student_count -= 1
            return len(sandwiches[sandwiches.index(1):])
                
        else:
            while student_count > 0:
                index = sandwiches.index(0)
                sandwiches = sandwiches[index+1:]
                student_count -= 1
            return len(sandwiches[sandwiches.index(0):])