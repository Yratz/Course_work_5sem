CREATE TABLE
    Students (
        idStudent SERIAL,
        name varchar(20) NOT NULL,
        secondName varchar(20) NOT NULL,
        PRIMARY KEY (idStudent)
    );

CREATE TABLE
    Answers (
        idAnswer SERIAL,
        variants varchar(150) NOT NULL,
        isTrue boolean NOT NULL,
        PRIMARY KEY (idAnswer)
    );

CREATE TABLE
    Exams (
        idExam SERIAL,
        title varchar(150) NOT NULL,
        timemin int NOT NULL,
        PRIMARY KEY (idExam)
    );

Create table
    Questions (
        idQuestion SERIAL,
        body varchar(150) not null,
        ball decimal NOT NULL,
        primary key (idQuestion)
    );

Create table
    QuestionsAnswers (
        idQa SERIAL,
        Answer int not null,
        Question int not null,
        primary key (idqa),
        FOREIGN KEY (Answer) references Answers (idAnswer),
        FOREIGN KEY (Question) references Questions (idQuestion)
    );

Create table
    QuestionsExam (
        idQe SERIAL,
        Question int not null,
        Exam int not null,
        primary key (idQe),
        FOREIGN KEY (Exam) references Exams (idExam),
        FOREIGN KEY (Question) references Questions (idQuestion)
    );

Create table
    StudentsResult (
        idSr SERIAL,
        Exam int not null,
        Student int not null,
        result decimal not null,
        maximum decimal not null,
        state int not NULL,
        FOREIGN KEY (Exam) references Exams (idExam),
        FOREIGN KEY (Student) references Students (idStudent)
    );