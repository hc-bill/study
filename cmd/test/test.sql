select sname
from student
         left join
     (select student_id
      from score
               left join
           (select cid
            from course
                     left join
                 (select tid
                  from teacher
                  where tname != "xxx"
                 ) teacher_table
                 on course.teacher_id = teacher_table.tid
           ) course_table
           on score.course_id = course_table.cid
     ) score_table
     on student.sid = score_table.student_id;








