select concat_ws(':',emp_id, concat_ws(' ', fname, lname), dept) from employees 
where emp_id = 1;

select concat_ws(':',emp_id, conact_ws(' ', fname, lname), dept) from employees 
where emp_id = 1;

select concat(left(dept, 1), emp_id), fname from employees;



select left(fname, 3) from employees;
select right(fname, 3) from employees;
select trim('  alright  ') as trimed_val;
select length(trim('  alright  ')) as len;
select position('om' in 'Thomas');

select * from employees
where salary = (select max(salary) from employees);

select fname, salary, salary * 0.1 as bonus from  employees;

select fname, salary,
case
	when salary >= 55000 then 'High'
	when salary between 48000 and 55000 then 'Mid'
	else 'Low'
end as sal_cat
from employees;

select
case
	when salary >= 55000 then 'High'
	when salary between 48000 and 55000 then 'Mid'
	else 'Low'
end as sal_cat, 
count(emp_id)
from employees 
group by sal_cat
order by sal_cat desc;