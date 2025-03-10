# learn_golang

Golang (or Go) is an open-source programming language developed by Google. It is designed for simplicity, efficiency, and scalability. Go is widely used for building web applications, cloud services, networking tools, and distributed systems. Some of its key features include:

Compiled and Statically Typed: Ensures fast execution and type safety.
Concurrency Support: Uses goroutines for lightweight and efficient parallel execution.
Garbage Collection: Manages memory automatically.
Strong Standard Library: Provides built-in support for networking, file handling, and more.
Cross-Platform: Works on Windows, macOS, and Linux.
How to Install Golang
Step 1: Download Go
Visit the official Go website: https://go.dev/dl/
Step 2: Install Go
Windows: Run the .msi installer and follow the setup instruction
Linux: Use the following commands (for Ubuntu/Debian-based systems)
.. wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
.. sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
.. echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
..  source ~/.bashrc
Step 3: Verify Installation
After installation, open a terminal or command prompt and run:
.. go version
Step 4: Set Up GOPATH (Optional)
For better project management, configure the Go workspace:
.. mkdir -p ~/go/{bin,src,pkg}
.. echo 'export GOPATH=$HOME/go' >> ~/.bashrc
.. echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
.. source ~/.bashrc
# learn_golang
