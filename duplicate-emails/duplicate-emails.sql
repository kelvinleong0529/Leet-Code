# Write your MySQL query statement below
SELECT DISTINCT p1.Email
FROM Person p1
INNER JOIN Person p2
ON p1.Id <> p2.Id AND p1.Email = p2.Email