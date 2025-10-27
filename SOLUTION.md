# Solution Steps

1. Create the project structure with three main files: main.go, service/healthchecker.go, and handler/health.go.

2. In handler/health.go, define a Result struct with URL, Status, and Err fields.

3. Implement the CheckURL function in handler/health.go to simulate checking the health of a URL by introducing predefined delays and timeouts for selected URLs using a goroutine and select statement.

4. In service/healthchecker.go, implement the CheckURLs function, which takes a list of URLs, spawns a goroutine for each CheckURL call, and collects results through a channel.

5. Use sync.WaitGroup in CheckURLs to synchronize all goroutines and close the results channel when work is complete.

6. In main.go, define a list of 10 URLs to check, start a timer, call CheckURLs, and print the status/error for each URL as results arrive.

7. After all results are printed, display the total elapsed time indicating the concurrent nature of the solution.


