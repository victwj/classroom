// Snippet from Satellite project to upload data to BigTable
// This script does not work on its own; shown here for reference
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"path"
	"strings"

	"cloud.google.com/go/bigtable"
)

func uploadModule() {

	// Set up connection
	projectID := configStr("modules.upload.project_id")  // The Google Cloud Platform project ID
	instanceID := configStr("modules.upload.instance_id") // The Google Cloud Bigtable instance ID
	tableID := configStr("modules.upload.table_id")       // The Google Cloud Bigtable table

	// Get filenames
	directory := path.Base(configStr("run_dir"))
	interferenceFileName := configStr("modules.detect.output.interference")
	taggedResolversFileName := configStr("modules.tag.output.tagged_resolvers")
	timestamp := strings.Replace(strings.Replace(directory, "-", "", -1), "T", "-", -1)

	// Init and file open
	fin := openFile(interferenceFileName)
	defer fin.Close()

	logInfo("Start: BigTable upload", logData{
		"projectID":        projectID,
		"instanceID":       instanceID,
		"tableID":          tableID,
		"upload_timestamp": timestamp,
	})

	// Load map of resolver countries
	resolverCountries := loadResolverCountries(taggedResolversFileName)

	// Set up Bigtable data operations client
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		logPanic("Could not create BigTable data operations client", err)
	}
	tbl := client.Open(tableID)

	// Initialize
	bulkDataSize := 90000
	mutations := make([]*bigtable.Mutation, bulkDataSize)
	rowKeys := make([]string, bulkDataSize)
	totalCounter := 0
	counter := 0

	// Scan interference data
	s := bufio.NewScanner(fin)
	for s.Scan() {

		// Unpack and parse data
		var data map[string]interface{}
		if err := json.Unmarshal(s.Bytes(), &data); err != nil {
			panic(err)
		}
		domain := data["query"].(string)
		resolverIP := data["resolver"].(string)
		passed := data["passed"].(bool)
		var country string
		if val, ok := resolverCountries[resolverIP]; ok {
			country = val
		} else {
			country = ""
		}

		// Create row key
		rowKey := timestamp + "#" + resolverIP + "#" + domain + "#" + country

		// Create mutation
		mut := bigtable.NewMutation()
		passedByte := make([]byte, 1)
		passedByte[0] = 0
		if passed {
			passedByte[0] = 1
		}
		mut.Set("passed", "bool", bigtable.Now(), passedByte)

		// Append
		mutations[counter] = mut
		rowKeys[counter] = rowKey

		counter++
		totalCounter++

		if counter == len(mutations) {
			// Apply mutation
			rowErrs, err := tbl.ApplyBulk(ctx, rowKeys, mutations)
			if err != nil {
				logPanic("Could not apply BigTable bulk row mutation", err)
			}
			if rowErrs != nil {
				logPanic("Could not apply BigTable mutation", rowErrs[0])
			}

			// Clear and reset counter
			mutations = make([]*bigtable.Mutation, bulkDataSize)
			rowKeys = make([]string, bulkDataSize)
			counter = 0
		}
	}

	if counter > 0 {
		// Apply mutation
		rowErrs, err := tbl.ApplyBulk(ctx, rowKeys[:counter], mutations[:counter])
		if err != nil {
			logPanic("Could not apply BigTable bulk row mutation", err)
		}
		if rowErrs != nil {
			logPanic("Could not apply BigTable mutation", rowErrs[0])
		}
	}

	// Close client
	if err = client.Close(); err != nil {
		logPanic("Failed to close BigTable client", err)
	}

	logInfo("Done: BigTable upload", logData{
		"num_uploaded": totalCounter,
	})
}

func loadResolverCountries(taggedResolversFile string) map[string]string {
	fin := openFile(taggedResolversFile)
	defer fin.Close()
	s := bufio.NewScanner(fin)
	result := make(map[string]string)
	for s.Scan() {
		var data map[string]string
		if err := json.Unmarshal(s.Bytes(), &data); err != nil {
			panic(err)
		}
		result[data["resolver"]] = strings.ToLower(strings.Replace(data["country"], " ", "", -1))
	}
	return result
}
