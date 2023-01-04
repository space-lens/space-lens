# Space-lens (WIP)

## Dropin replacement cleanmymac spacelens

> proposal : https://github.com/alfiankan/space-lens/issues/1



# Summary
Space-lens is a files analyzer that aims to create a list of deletions using filtering so that we can easily clean up storage, basically the same as the space-lens feature in the cleanmymac application, we can analyze by indexing all files on the disk or specified location and selecting files deletion . but the difference is that in this project, filters that are adapted to the general public or developers will be carried out such as large file filters, cache filters, log file filters, node_modules filters, .etc file extension filters. for the first plan in version 1.0.0 we will use TUI.

# Motivation
As a mac user using a base model mac with limited ssd (256) and maybe for other pc users with small storage need to clean their storage to keep their device up and running, popular way on mac is using clean my mac but its proprietary  app, even when we just want to use space lens to delete files manually.

in other case it is not easy to analyze detailed space used in unix virtual machine like, to do mass analysis and delete some tools can only show disk usage and presentation and we still need to wipe ourselves, so in this proposal I would like to start a project to make dropin replacement on TUI (Terminal user interface) first, and maybe will be gui in future.

# Goal
Create an easy-to-use yet fast tool to clean up storage on unix like os by analyzing and filtering files then creating delete lists and performing bulk deletes all at once.

# Detailed design

## Flow diagram
![diag](https://user-images.githubusercontent.com/40946917/210474478-e54b6636-1a72-483d-b4b3-6951bb8cd6d2.png)




## Indexing, filtering and sorting
This application is required to create a tree data structure for interpreting file systems, and perform filtering, searching and sorting on a single data structure to save memory usage and fast processing.

1. File/dir walking 
Scanning all folders and files below the path defined recursively.
2. Building tree
Inserting full path and file size to tree.

```go
	// prototype in go
	var tree Filetree

        // walking file and dir recursively
	filepath.Walk("/Users/john/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
                         // apply filter
			if filter(path, "node_modules") {
                                 // add to tree
				tree.add(path, info.Size())
			}
			return nil
		})
```
3. Shows a filtered list view to the user
4. User will select the desired files to delete
5. User confirm bulk deletion


# MLP (Minimum lovable product)
1. analyze 
analyzes files on disks or specific paths defined by the user, the result is a file tree with detailed sizes and dates.
here is an illustration using NERDTREE nvim :

     <img width="681" alt="Screenshot 2023-01-03 at 06 50 48" src="https://user-images.githubusercontent.com/40946917/210285594-94a9dc0a-b9b0-4c96-8fda-ba631b194ecd.png">


6. filter
filtering files while analyzing them is very helpful in my opinion, let's say we need to clean files related to node_modules or we need to clean big files or old files etc.

     <img width="824" alt="Screenshot 2023-01-03 at 06 54 09" src="https://user-images.githubusercontent.com/40946917/210285725-3aae55de-31af-434d-8a50-59c84829a3ea.png">


7. bulk deletion
bulk deleteion just like in clean mymac


