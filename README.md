## Google Drive query builder 

This package helps build search queries for Google Drive API.

See the URLs below for Drive `files.list` API and search syntax documentation:
 - https://developers.google.com/drive/api/v3/reference/files/list
 - https://developers.google.com/drive/api/v3/search-parameters.


Example:

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
	var query = q.Query(stm1).And(stm2).Or(q.Raw(`name = "Bar"`))

	fmt.Println(q.Stringize(query), "\n")

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


