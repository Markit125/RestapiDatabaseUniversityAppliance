CREATE TABLE department
(
  id                   INTEGER NOT NULL PRIMARY KEY,
  department_name      VARCHAR(20) NOT NULL,
  budget_places        INTEGER NOT NULL DEFAULT 0,
  paid_places          INTEGER NOT NULL DEFAULT 0,
  subject_1_id         INTEGER NOT NULL,
  subject_2_id         INTEGER NOT NULL,
  subject_3_id         INTEGER NOT NULL
);