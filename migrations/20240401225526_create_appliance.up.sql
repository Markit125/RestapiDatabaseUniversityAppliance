CREATE TABLE applience
(
  id                   INTEGER NOT NULL PRIMARY KEY,
  student_id           INTEGER NOT NULL,
  department_id        INTEGER NOT NULL FOREIGN KEY REFERENCES department(id),
  exam_subject_1_id    INTEGER NOT NULL FOREIGN KEY REFERENCES exam_subject(id),
  exam_subject_2_id    INTEGER NOT NULL FOREIGN KEY REFERENCES exam_subject(id),
  exam_subject_3_id    INTEGER NOT NULL FOREIGN KEY REFERENCES exam_subject(id)
);
