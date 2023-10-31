Here we create two channels to control the massive upload to a S3 Bucket.

One channel control how many threads or simultaneous uploads we will have. Adding and removing an empty struct to the channel. 
The other channel is used to fill with the files that has any error to retry again. 

Is a simple program with some S3 Bucket configuration, but have a good use of Go routines and channels.

cmd/generator/main.go - generates a lot of files until you cancel. Take care! 
cmd/uploader/main.go - upload the files at the tmp dir to a S3 Bucket. 