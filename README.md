# AutoDep 🚀

Hello, Gophers! 🖐 Welcome to **AutoDep** - a simple Go application that greets its visitors.

## What does it do? 🤔

Whenever you make a request, it returns a friendly greeting mentioning the requested URL path. Perfect for understanding the basics of Go web applications.

## Features 🌟

- Dockerized Go application 🐳
- Automatic build and test with GitHub Actions 🛠
- Simple and clean code for easy understanding 📖

## Setup and Installation 🛠

### Prerequisites:

- Ensure you have Docker installed on your system.
- Git must be installed for cloning purposes.

### Steps:

1. **Clone the repository** 📂:

   ```bash
   git clone https://github.com/ibilalkayy/AutoDep.git
   ```

2. **Navigate to the directory**:

   ```bash
   cd AutoDep
   ```

3. **Build the Docker image**:

   ```bash
   docker build -t autodep .
   ```

4. **Run the application**:

   ```bash
   docker run -p 8080:8080 autodep
   ```

5. Open a web browser or use a tool like `curl` to navigate to `http://localhost:8080`. You'll be greeted by our friendly app! 😊

## Contribute 🤝

Feel free to fork, improve, make pull requests or fill issues. I'll be glad to fix bugs you detected or improve the extension. Any feedback is welcome!

## Links 🌐

- [Check out the repository on GitHub](https://github.com/ibilalkayy/AutoDep)

Happy coding, Gophers! 🚀🎉
