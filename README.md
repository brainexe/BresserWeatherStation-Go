[![Build Status](https://travis-ci.org/brainexe/BresserWeatherStation-Go.svg?branch=master)](https://travis-ci.org/brainexe/BresserWeatherStation-Go)

# Installation
## Install rtl-sdr, e.g. via https://gist.github.com/floehopper/99a0c8931f9d779b0998

```
git clone git://git.osmocom.org/rtl-sdr.git
go get github.com/brainexe/BresserWeatherStation-Go
```


# Usage
Fetch data from USB device and pipe data in ./weatherStation. Additionally log the output into out.bin:

```
rtl_fm -M am -f 868.300M -s 48k -g 50 | tee out.bin | ./weatherStation -noise=700

# 2nd example: stop at first of prefetched data
./weatherStation -noise=850 -stop-at-first < out.bin

```

