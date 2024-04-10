
CREATE TABLE exam_subject
(
  id              BIGSERIAL NOT NULL PRIMARY KEY,
  subject_name    VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE result
(
  id                   BIGSERIAL NOT NULL PRIMARY KEY,
  score                INTEGER NOT NULL,
  student_id           INTEGER NOT NULL,
  exam_subject_id      INTEGER NOT NULL,

  FOREIGN KEY (student_id) REFERENCES students(id),
  FOREIGN KEY (exam_subject_id) REFERENCES exam_subject(id)
);
