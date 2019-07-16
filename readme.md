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
