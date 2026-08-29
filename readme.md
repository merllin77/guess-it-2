# Guess it 2

A Go command-line program that reads numbers continuously from standard input and prints a prediction range for the next number after every input.

## How it predicts

Each received value is a y-coordinate and its zero-based input position is x. The program fits a linear-regression line to all received values and predicts the next y at x = n, where n is the number of received values.

- With one value, it uses a safe range of 50 below to 50 above that value.
- With two values, it uses a smaller safe range of 20 below to 20 above the regression prediction. Two points have zero residual error because one line joins them exactly.
- From three through 100 values, it measures residual standard deviation: the observed distance from the regression line. Low-noise histories narrow progressively as more data becomes available; higher-noise histories retain a wider multiplier.
- After 100 values, it rounds the regression prediction to the nearest integer and prints an interval from 6 below to 6 above it. This gives a stable integer width of 12, matching the supplied `linear-regr` reference's observed width policy.
- Bounds are emitted as two integers separated by one space.

This separates the prediction centre (linear regression) from the range width (early uncertainty, residual noise, and final integer scoring width).

## Input and output

Provide one number per line. For each received number, the program writes one line with two integers separated by a space: the lower and upper bounds for the next input.

Example:

```text
189
139 239
113
101 201
```

The exact predicted bounds change as more values arrive.

## Run locally

```sh
go run .
```

Then type numbers line by line. End the input with `Ctrl+D` on Linux/macOS.

To build an executable:

```sh
go build -o guess-it .
```

## Test

Run the Go checks with:

```sh
go test ./...
```

The supplied browser-based tester is in `guess-it-dockerized/`. It expects the compiled Linux executable and launcher at `guess-it-dockerized/student/`.

```sh
cd guess-it-dockerized
docker compose up --build
```

Open `http://localhost:3000` and compare the **Final Result** as well as coverage. A wider interval can improve correct guesses while lowering the score, so Final Result is the main measure for tuning.

For a reproducible comparison, score all five fixed files in both official folders, `data_sets/4/` and `data_sets/5/`, against `linear-regr` rather than relying on one browser click. The final audited policy won 4 of 5 Data 4 samples and 5 of 5 Data 5 samples.
