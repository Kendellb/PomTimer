# PomTimer
simple pomodoro timer in Go with bubble tea and lipgloss

## Prerequisites

Before you can build and run PomTimer, ensure you have the following installed:

- [Go 1.20+](https://golang.org/dl/)
- A terminal to run the program.

## Installation

1. **Clone the repository:**

   First, clone the PomTimer repository to your local machine:

   ```bash
   git clone https://github.com/<your-username>/PomTimer.git
   ```
   
2. Navigate to the project directory:

Change to the directory where you cloned the repository:

cd PomTimer

3. Build the project:

Next, build the Go program by running the following command:

```bash
go build -v ./...
```

4. Run the program:

After the build process completes, run the main.go file:

```bash
./PomTimer
```

The program will start running in your terminal, and you can interact with it according to the instructions displayed.

## Usage

Once the program starts, you can:
```
Press s to start the timer.
Press p to pause the timer.
Press q to quit the program.
```

The timer will display the remaining time for the current session, and the program will automatically switch between work sessions and break intervals.
