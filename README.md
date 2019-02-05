## Google Drive query builder 

This package helps build search queries for [Google Drive API](https://godoc.org/google.golang.org/api/drive/v3).

Drive `files.list` API and search syntax are documented here:
 - https://developers.google.com/drive/api/v3/reference/files/list
 - https://developers.google.com/drive/api/v3/search-parameters.


### Usage

The example below searches for documents matching the following criteria: 

 -  ( content-type is text/plain **AND** name includes "Foo" ) **OR** name is "pluto" 
- **AND** document was created less than 10 hours ago

```
package main

import (
	"context"
	"fmt"
	"log"
	q "gitlab.com/mxmz/google-drive-query-builder/query"
	"golang.org/x/oauth2/google"
	drive "google.golang.org/api/drive/v3"
)

func main() {

	var stm1 = q.MimeType().Equal("text/plain")
	var stm2 = q.Name().Contains("Foo")

	var query = q.
		Query(stm1).
		And(stm2).
		Or(q.Raw(`name = "pluto"`)).
		And(q.CreatedTime().After(time.Now().Add(-10 * time.Hour)))


	ctx := context.Background()
	client, _ := google.DefaultClient(ctx, drive.DriveScope)
	driveSvc, _ := drive.New(client)

	var svc = drive.NewFilesService(driveSvc)
	lst, _ := svc.List().
		Q(q.Stringize(query)).
		Fields("files(name,createdTime,properties)").
		OrderBy("createdTime").
		PageSize(1000).
		Do()

	for _, v := range lst.Files {

		log.Printf("%s %v \n", v.Name, v.CreatedTime)

	}


}

```


