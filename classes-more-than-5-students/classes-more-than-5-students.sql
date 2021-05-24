# Write your MySQL query statement below
Select distinct class
from courses
where class in (select class
From courses
group by class
having count(distinct student)>=5)