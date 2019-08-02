# Golang client library for bitcoinsv blockchain

This library will allow you to complete a basic set of functions with a bitcoinsv node.
Also it has the address package you can generate keys with and eventually sign transactions with.

THIS IS VERY NEW AND USED-TO BE A MULTICHAIN LIB. PLEASE BE PATIENT.

If you wish to contribute to flesh out the remaining API calls, please make pull requests.

## Testing

To fully test this package it is neccesary to have a full hot node running at the given parameters.

```

  host := flag.String("host", "localhost", "is a string for the hostname")
  port := flag.String("port", "80", "is a string for the host port")
  username := flag.String("username", "bitcoinsvrpc", "is a string for the username")
  password := flag.String("password", "12345678", "is a string for the password")

  flag.Parse()

  client := bitcoinsv.NewClient(
      *host,
      *port,
      *username,
      *password,
  )

  obj, err := client.GetInfo()
  if err != nil {
      panic(err)
  }

  fmt.Println(obj)

```
