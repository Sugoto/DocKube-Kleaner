# Kube Kleaner 🧹

**Kube Kleaner** is a powerful garbage collection service designed to keep your Docker 🐳 and Kubernetes 🛠️ environments clean and efficient. It automates the cleanup of unused containers, images, and volumes in Docker, as well as idle pods and resources in Kubernetes. With Kube Kleaner, you can ensure your development and production environments remain clutter-free and optimized 🚀.

## Features 🌟

- **Docker Cleanup**:
  - Remove unused containers 🗑️
  - Clean up dangling images 🖼️
  - Delete unused volumes 💾

- **Kubernetes Cleanup**:
  - Remove idle pods 💤
  - Clean up unused resources 🧹

## Getting Started 🚀

### Prerequisites

- **Go**: [Download and install Go](https://golang.org/dl/)
- **Docker**: [Install Docker](https://docs.docker.com/get-docker/)
- **Kubernetes**: Use [Minikube](https://minikube.sigs.k8s.io/docs/start/) or [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/) for a local Kubernetes cluster

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/sugoto/kube-kleaner.git
   cd kube-kleaner
   ```

2. **Build the Project**:
   ```bash
   go build -o kube-kleaner ./cmd/main.go
   ```

3. **Run the Project**:
   ```bash
   ./kube-kleaner
   ```

### Running with Docker

1. **Build the Docker Image**:
   ```bash
   docker build -t kube-kleaner .
   ```

2. **Run the Docker Container**:
   ```bash
   docker run -v /var/run/docker.sock:/var/run/docker.sock \
              -v ~/.kube:/root/.kube \
              kube-kleaner
   ```

### Testing

Run the tests to ensure everything is working correctly:
```bash
go test -v
```

## Contributing 🤝

We welcome contributions! Please follow these steps:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Create a new Pull Request

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
Happy cleaning with Kube Kleaner! 🎉
