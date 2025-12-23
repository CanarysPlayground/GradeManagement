-- schema
CREATE TABLE IF NOT EXISTS grades (
  id SERIAL PRIMARY KEY,
  student_id INTEGER NOT NULL,
  course TEXT NOT NULL,
  score NUMERIC(5,2) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

-- sample data
INSERT INTO grades (student_id, course, score) VALUES
(1, 'Math', 95.0),
(2, 'Science', 88.5),
(3, 'History', 72.0);
