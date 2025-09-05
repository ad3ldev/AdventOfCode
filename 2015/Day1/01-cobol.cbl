       IDENTIFICATION DIVISION.
       PROGRAM-ID. MAIN.

       ENVIRONMENT DIVISION.
       INPUT-OUTPUT SECTION.
       FILE-CONTROL.
           SELECT INPUT-FILE ASSIGN TO 'input.txt'
               ORGANIZATION IS LINE SEQUENTIAL
               FILE STATUS IS FILE-STATUS.

       DATA DIVISION.
       FILE SECTION.
       FD INPUT-FILE.
       01 INPUT-RECORD         PIC X(10).

       WORKING-STORAGE SECTION.
       01 FILE-STATUS          PIC XX.
       01 FLOOR-COUNT          PIC S9(9) VALUE ZERO.
       01 COUNTER              PIC 9(3) VALUE ZERO.
       01 EOF-FLAG             PIC X VALUE 'N'.
           88 END-OF-FILE      VALUE 'Y'.

       PROCEDURE DIVISION.
       MAIN-LOGIC.
           OPEN INPUT INPUT-FILE
           PERFORM READ-RECORD
           PERFORM UNTIL END-OF-FILE
               PERFORM PROCESS-RECORD
               PERFORM READ-RECORD
           END-PERFORM
           CLOSE INPUT-FILE
           DISPLAY FLOOR-COUNT
           STOP RUN.

       READ-RECORD.
           READ INPUT-FILE INTO INPUT-RECORD
           EVALUATE FILE-STATUS
               WHEN '00'
                   CONTINUE
               WHEN '10'
                   SET END-OF-FILE TO TRUE
           END-EVALUATE.

       PROCESS-RECORD.
           PERFORM VARYING COUNTER FROM 1 BY 1
               UNTIL COUNTER > 10
               EVALUATE INPUT-RECORD(COUNTER:1)
                   WHEN '('
                       ADD 1 TO FLOOR-COUNT
                   WHEN ')'
                       SUBTRACT 1 FROM FLOOR-COUNT
               END-EVALUATE
           END-PERFORM.
