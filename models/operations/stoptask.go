// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
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
