# 🚀 Enhanced TCP Echo Server in Go

A production-ready TCP server with concurrency, logging, command protocol, and custom responses. Built for Network Programming 2 (Test #2).

## Features
- ✅ Concurrent client handling (goroutines)  
- 📝 Logging connections/messages to files  
- ⏱️ 30-second inactivity timeout  
- 🛡️ Overflow protection (1024-byte limit)  
- 🤖 Personality mode (`hello` → `Hi there!`)  
- 🔌 Command protocol (`/time`, `/quit`, `/echo`)  

---

## 🛠️ How to Run

### Prerequisites
- `netcat` (for testing)  

### Steps
1. **Clone and enter repo**:
   ```
   git clone https://github.com/ibigona/Improved-TCP-Echo-Server.git
   cd tcp-echo-server
Start the server (default: port 4000):
go run . --port 4000

Connect with netcat:
nc localhost 4000

## Test Commands
Command	Response
hello	Hi there!
bye	Goodbye! (disconnects)
/time	Current server time
/echo <msg>	Repeats <msg>
/quit	Closes connection

## 🎓 Educationally Enriching Features
Task	Key Learnings
Concurrency	Goroutines vs. threads, race condition avoidance
Logging	File I/O, timestamp formatting, defer for cleanup
Graceful Disconnects	Error handling (io.EOF), connection lifecycle management
Byte Parsing	Manual buffer management without bufio
Inactivity Timeout	SetReadDeadline, net.Error type assertions
Command Protocol	Prefix-based parsing (/command), stateful responses

## 💡 Lessons Learned
1. Concurrency Pitfalls
go
go handleConnection(conn) // Unbounded goroutines!
Fix: Added a semaphore pattern to limit concurrent connections.

2. Deadline Gotchas
go
conn.SetReadDeadline(time.Now().Add(30 * time.Second))
Lesson: Deadlines apply per operation, not per connection. Reset after each read.

3. Logging Efficiency
go
file.WriteString(msg) // Slow for high throughput!
Optimization: Switched to buffered writes with bufio.NewWriter.

4. Command Security
go
msg = strings.TrimSpace(input) // Prevent padding attacks
Risk: Naive splitting (strings.Split) could allow command injection.

