CREATE TABLE applience
(
  id                   BIGSERIAL NOT NULL PRIMARY KEY,
  student_id           BIGSERIAL NOT NULL,
  department_id        BIGSERIAL NOT NULL,
  exam_subject_1_id    BIGSERIAL NOT NULL,
  exam_subject_2_id    BIGSERIAL NOT NULL,
  exam_subject_3_id    BIGSERIAL NOT NULL,
  FOREIGN KEY (department_id) REFERENCES department(id),
  FOREIGN KEY (exam_subject_1_id) REFERENCES exam_subject(id),
  FOREIGN KEY (exam_subject_2_id) REFERENCES exam_subject(id),
  FOREIGN KEY (exam_subject_3_id) REFERENCES exam_subject(id)
);
