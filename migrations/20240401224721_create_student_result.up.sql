
CREATE TABLE exam_subject
(
  id                   INTEGER NOT NULL PRIMARY KEY,
  exam_subject_name    VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE result
(
  id                   INTEGER NOT NULL PRIMARY KEY,
  result_score         INTEGER NOT NULL,
  student_id           INTEGER NOT NULL FOREIGN KEY REFERENCES students(id),
  exam_subject_id      INTEGER NOT NULL FOREIGN KEY REFERENCES exam_subject(id)
);
