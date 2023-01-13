# Exercise Go Fundamentals
### How to start
1. Fork this repository to your private group `batch-{number}/{your-name}`
2. Clone forked repo (repo on your private group) to your local machine
3. Open terminal and `cd` to this project
4. Install dependencies with `go install`
5. Fix the function in `pos.go` and `park_billing.go` file
6. You can use `main.go` to test the output, build and run main fn with command `go run .`
7. After all functions completed run `go test ./...`, make sure all the test is passed!

### Parking Lot Billing
Parking Price List

| Duration   | Motorcycle | Car       |
|------------|------------|-----------|
| First 1 hr | 3000       | 7000      |
| After 1 hr | 2000/hour  | 5000/hour |

If vehicle parked more than 24hrs then they will be an extra charged, as below:

| Vehicle Type | Extra Charge |
|--------------|--------------|
| Motorcycle   | +20000       |
| Car          | +50000       |

### Notes
* Assume all barcode in cart is exist in database
* You don't have to add new test cases, but if you feel something is missing then feel free to write it
* Don't rename the function name
