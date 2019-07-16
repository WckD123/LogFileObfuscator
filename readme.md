# Log File Obfuscator

This GoLang program returns obfuscated values for certain fields when accessed by a user

## APIs

```bash
                /upload  
		/user/:key/:id   (Use Key and Id to search)
		/admin/:key/:id  (Use Key and Id to search)
		/uploadwithpath/:path  (Path is the path of the file)
```

## CONCURRENCY TEST

To test concurrency: 

```bash
http://localhost:3000/api/uploadwithpath/test.json
```

## Make changes to UploadUsingConcurrency method in services.go file. Change the variable "numberOfWorkers"

```bash
numberOfWorkers := 30

for w := 1; w <= numberOfWorkers; w++ {
	go uploadHelper(w, logfiles, jobs)
}
```

## Results

### Test 1:

11,500 inputs into DB

10 Workers

8.01 seconds


### Test 2:

11,500 inputs into DB

30 Workers

12.96 seconds (Figure Out Why??)

[WORKED WELL SECOND TIME, SO PROBABLY SOME ANOMALY]

### Test 3 :

Same as test 1 but with no print statements

2.72 seconds

### Test 4 :

Same as test 2 but with no print statements

2.68 seconds

## Update on concurrency:

Rather than passing indexes related to the slice, I have now created a channel of the "Logfile" model itself. This way, I don't have to pass the Logfile slice to every go routine making it even faster.

### Test 5:

Same as Test 3

2.61 milli seconds. (This is a 1000x increase!!!!)

### Test 6:

Same as Test 4 
 
1.21 milli seconds (This is a 1000x increase!!!!)