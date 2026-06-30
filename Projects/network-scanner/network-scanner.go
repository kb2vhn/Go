// John Wood
// June 28, 2026
// purpose is to take a hostname resolve and scan all ports.
// this will begin to get moved to a modular style now that the base concept is working
// the idea is to build my own standard nmap replacment nothing when it comes to the scripts
// that nmap can execute but just normal port status open closed filtered.
// ultimate goal is to understand Go and build some custom tools with it.
// network_scanner
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort" // Kept and used at the bottom for the final report
	"sync"
	"time"
)

func worker(ip string, ports, results chan int, rateLimiter <-chan time.Time, wg *sync.WaitGroup) {
	for p := range ports {
		<-rateLimiter // Enforce the speed limit

		address := fmt.Sprintf("%s:%d", ip, p)
		conn, err := net.DialTimeout("tcp", address, 300*time.Millisecond)

		if err != nil {
			// CLOSED PORT: Skip it completely. No locks, no channel bloat!
			continue
		}
		conn.Close()
		
		// OPEN PORT: Only write to the channel when we find a real open port
		results <- p
	}
	// Signal to main that this worker thread is totally done
	wg.Done()
}

func main() {
	hostFlag := flag.String("host", "", "The target domain name or IP address to scan (Required)")
	// default at 128 to keep network and port status extreamly stable.
	rateFlag := flag.Int("rate", 128, "Maximum connections per second")
	
	flag.Parse()

	if *hostFlag == "" {
		fmt.Println("Error: Missing required target host.")
		flag.Usage()
		os.Exit(1)
	}

	ips, err := net.LookupHost(*hostFlag)
	if err != nil || len(ips) == 0 {
		fmt.Printf("Failed to resolve host %s: %v\n", *hostFlag, err)
		return
	}
	targetIP := ips[0]
	
	fmt.Printf("[+] Target Resolved: %s (%s)\n", *hostFlag, targetIP)
	fmt.Printf("[+] Scanning all 65535 ports at %d conns/sec...\n", *rateFlag)
	fmt.Println("--------------------------------------------------")

	ports := make(chan int, 128)
	results := make(chan int)
	var openports []int

	ticker := time.NewTicker(time.Second / time.Duration(*rateFlag))
	defer ticker.Stop()

	var wg sync.WaitGroup

	// Spin up 128 workers
	for i := 0; i < 128; i++ {
		wg.Add(1) 
		go worker(targetIP, ports, results, ticker.C, &wg) // Passing the address of wg
	}

	// Queue up the ports to be scanned
	go func() {
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
		close(ports) // Triggers workers to break out of their loops when done
	}()

	// Orchestration: Wait for workers to finish, then close results channel
	go func() {
		wg.Wait()
		close(results) 
	}()

	// Read ONLY the open ports. This loop naturally finishes when 'results' is closed.
	for port := range results {
		openports = append(openports, port)
	}

	// PRINT THE FINAL REPORT
	fmt.Println("\n---------------- SCAN REPORT ----------------")
	sort.Ints(openports) // Using the sort package here so it compiles!
	if len(openports) == 0 {
		fmt.Println("No open ports found.")
	} else {
		fmt.Printf("Found %d open ports:\n", len(openports))
		for _, port := range openports {
			fmt.Printf(" -> Port %d is OPEN\n", port)
		}
	}
	fmt.Println("---------------------------------------------")
}

