# Logrus Hook for Aliyun SLS

[![Build Status](https://travis-ci.org/innopals/sls-logrus-hook.svg?branch=master)](https://travis-ci.org/innopals/sls-logrus-hook)
[![](https://goreportcard.com/badge/github.com/innopals/sls-logrus-hook)](https://goreportcard.com/report/github.com/innopals/sls-logrus-hook)

Simple yet powerful logrus hook for aliyun simple logging service.

## Features

+ Batch sending logs asynchronously
+ Dump huge logs (exceeds sls service limit) to stdout
+ Fallback dumping logs to stdout when sls api not available

## Getting Start

Simply add a hook to global logrus instance or any logrus instance.

```golang
slsLogrusHook, err := hook.NewSlsLogrusHook("<project>.<region>.log.aliyuncs.com", "access_key", "access_secret", "logstore", "topic")
logrus.AddHook(slsLogrusHook)
```

Ensure logs are flushed to sls before program exits
```golang
slsLogrusHook.Flush(5 * time.Second)
```

## Performance Tuning

Disable processing logs for default output.

```golang
logrus.SetFormatter(&hook.NoopFormatter{})
logrus.SetOutput(ioutil.Discard)
```

Disable locks.

```golang
logrus.SetNoLock()
```

Setting send interval if necessary.

```golang
slsLogrusHook.SetSendInterval(100 * time.Millisecond) // defaults to 300 * time.Millisecond
```

## Contributing

This project welcomes contributions from the community. Contributions are accepted using GitHub pull requests. If you're not familiar with making GitHub pull requests, please refer to the [GitHub documentation "Creating a pull request"](https://help.github.com/articles/creating-a-pull-request/).
