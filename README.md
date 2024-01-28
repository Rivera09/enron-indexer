# Steps to run the indexer:

1. Run the zincsearch engine. In my case, I'm running the zincsearch engine through Docker with the following command

```
docker run --user root -v %cd%/data:/data -e ZINC_DATA_PATH="/data" -p 4080:4080 -e ZINC_FIRST_ADMIN_USER=admin -e ZINC_FIRST_ADMIN_PASSWORD=admin --name zincsearch-enron zincsearch
```

2. (Optional) build the indexer

```
go build indexer.go
```

3. Run the indexer

```
go run indexer.go <enron-folder-path>

```

or if you created the build file:

```
./indexer <enron-folder-path>
```

# Steps to run the email client

For windows, you can just open the start-server.bat file, which will automatically build the vue project and serve the files with go.

For other users, go to mamuro, build the vue project with

```
npm run build
```

then serve the files with go

```
go run server.go
```
