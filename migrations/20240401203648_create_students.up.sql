CREATE TABLE students
(
  student_id           INTEGER NOT NULL PRIMARY KEY,
  student_first_name   VARCHAR(20) NOT NULL,
  student_middle_name  VARCHAR(20) NOT NULL,
  student_last_name    VARCHAR(20) NOT NULL,
  birth_date           DATE NOT NULL,
  achievements         INTEGER NOT NULL DEFAULT 0,
  passport             VARCHAR(20) NOT NULL UNIQUE
);

CREATE UNIQUE INDEX passport_student ON students
(
  passport
);
