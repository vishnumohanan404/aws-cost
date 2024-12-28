# AWS Cost Analyser
`aws-cost` is a command-line tool to fetch AWS cost data and display a breakdown by service for a given date range.

## Installation
1. Clone the repository:

   ```bash
   git clone https://github.com/vishnumohanan404/aws-cost.git
   cd aws-cost
   ```

2. Build the tool:

   ```bash
   go build -o aws-cost
   ```

3. Set up AWS credentials:

   If you haven't already, configure your AWS credentials using:

   ```bash
   aws configure
   ```

## Usage

Run the following command to get the AWS cost breakdown for a specific date range:

```bash
./aws-cost --start-date "2024-12-01" --end-date "2024-12-31"
```

#### Flags
--start-date: Start date of the cost period (inclusive).  
--end-date: End date of the cost period (exclusive).

Dates should be in the format ```YYYY-MM-DD```.

## License
This project is licensed under the MIT License - see the LICENSE file for detail

