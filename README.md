# Go
Go Tools
I'm taking the code as a building block from the Black Hat Go book by Tom Steele, Chris, Pattern, and Dan Kottmann from No Starch Press


I have noticed a few things with the example code that i will be looking at after each chapter or concept.

June 28, 2026
When testing the scanner everything works well I played with 64 - 2048 goroutines the book uses 100 as a starting point and scanned all 65535 ports. I did notice that while the program was running that no other network conneciton to/from my system was allowed. Wireshark showed all expected traffic during the scan to the end site (a local server in my home lab). I will be looking at this closer as I work on a more improved scanner version that I can use command line arguments for. 
