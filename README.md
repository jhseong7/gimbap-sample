# Gimbap Sample

This is a sample project to demonstrate the use of GIMBAP.

each executive in the `cmd` folder is a sample of how to use GIMBAP.

## Description

List of samples:

1. Echo server engine example
   - This is a simple example of how to use the echo server engine.
2. Gin server engine example
   - This is a simple example of how to use the gin server engine.
3. Fiber server engine example
   - This is a simple example of how to use the fiber server engine.
4. Microservice example
   - This is a simple example of how to use the microservice feature.

## How to run the samples

The commands to run the samples are managed in the `Makefile` file. So if your environment does not have `make` installed, you can run the commands manually.

> See all the commands in the `Makefile` file.

```bash
make help
```

Then run the appropriate command.

> Example: To run the echo server engine example

```bash
make run-sample-echo
```

## Future sample plans

- DB connection example
  - some of the code is in but no command to run it yet.
- Middleware example
- etc.
