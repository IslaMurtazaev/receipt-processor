# Receipt Processor Service

### Description
The Receipt Processor is a simple web service that processes receipts and calculates points based on a set of predefined rules. The service provides two endpoints for processing and retrieving receipt data.

### Requirements:
1. Install Go 1.19+ on the system
2. Clone the Repository and enter the Project Directory
   ```bash
   git clone https://github.com/IslaMurtazaev/receipt-processor.git
   cd receipt-processor
   ```
3. Run the Application
    ```bash
    go run main.go
   ```
4. Access the API on http://localhost:8080

Alternatively, you can use docker. Start with step 2 and run these commands:
```bash
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

### Testing
Unit tests are included to validate points calculation. Run the tests with:
```bash
go test ./...
```

### In-Memory Storage
Receipts are stored in memory.
