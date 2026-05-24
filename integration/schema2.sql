-- Create the students table
CREATE TABLE students (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          email VARCHAR(255) UNIQUE NOT NULL
);

-- Create the courses table
CREATE TABLE courses (
                         id SERIAL PRIMARY KEY,
                         course_name VARCHAR(100) NOT NULL,
                         credits INTEGER NOT NULL CHECK (credits > 0)
);

-- Create the enrollments junction table
CREATE TABLE enrollments (
                             student_id INTEGER NOT NULL,
                             course_id INTEGER NOT NULL,
                             enrollment_date DATE NOT NULL DEFAULT CURRENT_DATE,
                             grade VARCHAR(5),
                             PRIMARY KEY (student_id, course_id),
                             FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
                             FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);