// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// PathParamTaskName - The name of the task to be started.
type PathParamTaskName string

const (
	PathParamTaskNameBackupDatabase            PathParamTaskName = "BackupDatabase"
	PathParamTaskNameBuildGracenoteCollections PathParamTaskName = "BuildGracenoteCollections"
	PathParamTaskNameCheckForUpdates           PathParamTaskName = "CheckForUpdates"
	PathParamTaskNameCleanOldBundles           PathParamTaskName = "CleanOldBundles"
	PathParamTaskNameCleanOldCacheFiles        PathParamTaskName = "CleanOldCacheFiles"
	PathParamTaskNameDeepMediaAnalysis         PathParamTaskName = "DeepMediaAnalysis"
	PathParamTaskNameGenerateAutoTags          PathParamTaskName = "GenerateAutoTags"
	PathParamTaskNameGenerateChapterThumbs     PathParamTaskName = "GenerateChapterThumbs"
	PathParamTaskNameGenerateMediaIndexFiles   PathParamTaskName = "GenerateMediaIndexFiles"
	PathParamTaskNameOptimizeDatabase          PathParamTaskName = "OptimizeDatabase"
	PathParamTaskNameRefreshLibraries          PathParamTaskName = "RefreshLibraries"
	PathParamTaskNameRefreshLocalMedia         PathParamTaskName = "RefreshLocalMedia"
	PathParamTaskNameRefreshPeriodicMetadata   PathParamTaskName = "RefreshPeriodicMetadata"
	PathParamTaskNameUpgradeMediaAnalysis      PathParamTaskName = "UpgradeMediaAnalysis"
)

func (e PathParamTaskName) ToPointer() *PathParamTaskName {
	return &e
}

func (e *PathParamTaskName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BackupDatabase":
		fallthrough
	case "BuildGracenoteCollections":
		fallthrough
	case "CheckForUpdates":
		fallthrough
	case "CleanOldBundles":
		fallthrough
	case "CleanOldCacheFiles":
		fallthrough
	case "DeepMediaAnalysis":
		fallthrough
	case "GenerateAutoTags":
		fallthrough
	case "GenerateChapterThumbs":
		fallthrough
	case "GenerateMediaIndexFiles":
		fallthrough
	case "OptimizeDatabase":
		fallthrough
	case "RefreshLibraries":
		fallthrough
	case "RefreshLocalMedia":
		fallthrough
	case "RefreshPeriodicMetadata":
		fallthrough
	case "UpgradeMediaAnalysis":
		*e = PathParamTaskName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PathParamTaskName: %v", v)
	}
}

type StopTaskRequest struct {
	// The name of the task to be started.
	TaskName PathParamTaskName `pathParam:"style=simple,explode=false,name=taskName"`
}

func (o *StopTaskRequest) GetTaskName() PathParamTaskName {
	if o == nil {
		return PathParamTaskName("")
	}
	return o.TaskName
}

type StopTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *StopTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *StopTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *StopTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}