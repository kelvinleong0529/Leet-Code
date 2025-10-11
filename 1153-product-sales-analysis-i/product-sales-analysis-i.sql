# Write your MySQL query statement below
select a.price, a.year, b.product_name
from sales a inner join product b on a.product_id = b.product_id;