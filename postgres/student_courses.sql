delete from student_courses
where student_id = 0;

select * from student_courses;

select * from courses;
delete from courses
where id > 3;

select s.id, s.name, sc.*, c.title
from students s
join student_courses sc on s.id = sc.student_id
join courses c on c.id = sc.course_id
where c.Title = 'Math';


select c.title, count(*)
from courses c
join student_courses sc on c.id = sc.course_id
group by c.title;