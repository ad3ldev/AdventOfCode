program main
    implicit none
    
    call read_file_as_string()
    
    contains

    subroutine read_file_as_string()
        CHARACTER(LEN=100000000):: file 
        CHARACTER(LEN=100000000):: line_buffer
        INTEGER :: io, stat, len, i, sum

        file = ''
        len = 0
        sum = 0
        
        OPEN(newunit=io, file='input.txt', status='old', action='read', iostat=stat)
        
        if (stat /= 0) then
            PRINT *, "Error opening file!"
            return
        end if
        
        do
            READ(io, '(A)', iostat=stat) line_buffer
            if (stat /= 0) exit
            
            if (len > 0) then
                file = file(1:len) // NEW_LINE('a') // TRIM(line_buffer)
                len = len + 1 + LEN_TRIM(line_buffer)
            else
                file = TRIM(line_buffer)
                len = LEN_TRIM(line_buffer)
            end if
        end do
        
        close(io)

        do i = 1, len
            if (file(i:i) == '(') then
                sum = sum + 1
            else
                sum = sum - 1
            end if
        end do

        print *, sum
        
    end subroutine read_file_as_string

end program main
