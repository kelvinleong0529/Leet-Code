# Write your MySQL query statement below
SELECT IFNULL(NULL,(SELECT DISTINCT(SALARY) FROM EMPLOYEE E ORDER BY SALARY DESC LIMIT 1 OFFSET 1)) AS SecondHighestSalary