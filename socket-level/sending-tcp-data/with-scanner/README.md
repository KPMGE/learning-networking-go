## Scanner 

Reading data from a tcp connection using fixed-size buffers is ok, but it means, we must make sense 
of the data we read. Oftentimes, that's not so simple. Fortunately, there's a well-tested implementation
that helps us reading data from streams, the bufio.Scanner.
