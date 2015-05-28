# bench_go

## write vs bufio(every Flush)

|            | real  | user |  sys  |
|------------|-------|------|-------|
| Direct     | 37.34 | 3.30 | 32.59 |
| bufio 4096 | 40.14 | 4.83 | 34.12 |
| bufio 8192 | 40.96 | 4.80 | 34.92 |
| bufio 16384| 40.70 | 4.84 | 34.43 |
