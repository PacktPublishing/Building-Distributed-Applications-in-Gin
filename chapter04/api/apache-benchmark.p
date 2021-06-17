set terminal png
set output "benchmark.png"
set title "Cache benchmark"
set size 1,0.7
set grid y
set xlabel "request"
set ylabel "response time (ms)"
plot "with-cache.data" using 9 smooth sbezier with lines title "with cache", "without-cache.data" using 9 smooth sbezier with lines title "without cache"