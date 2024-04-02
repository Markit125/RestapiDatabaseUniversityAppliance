CREATE TABLE applience
(
  id                   INTEGER NOT NULL PRIMARY KEY,
  student_id           INTEGER NOT NULL,
  department_id        INTEGER NOT NULL,
  exam_subject_1_id    INTEGER NOT NULL,
  exam_subject_2_id    INTEGER NOT NULL,
  exam_subject_3_id    INTEGER NOT NULL,
  FOREIGN KEY (department_id) REFERENCES department(id),
  FOREIGN KEY (exam_subject_1_id) REFERENCES exam_subject(id),
  FOREIGN KEY (exam_subject_2_id) REFERENCES exam_subject(id),
  FOREIGN KEY (exam_subject_3_id) REFERENCES exam_subject(id)
);
