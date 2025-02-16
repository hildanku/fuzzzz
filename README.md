## Status
- on development

## how to use
1. Build the binary: `go build -o <binaryname>`
2. If, build is done, Exec the binary using:
```
./<binaryname> fuzzzz -t https://<target>.com -w <wordlist>.txt
```

## todo:
1. feature to export fuzzing to csv or text
2. improve concurrency
3. mechanism caching when scanning?
4. mechanism for filter http status when scanning
5. rotate proxy
6. add delay per req
7. support other wordlist extension

## done
1. Migrate from ordinary cli to cobra cli (DONE)

## disclaimer
- This repository was created for learning and security research purposes. Use this tool if you done to make Agreement/Permission to Stakeholders.

- I am not responsible for any misuse of this tool. DWYOR! 
