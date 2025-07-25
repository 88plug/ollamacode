// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode

import (
	"context"
	"net/http"
	"net/url"

	"github.com/sst/opencode-sdk-go/internal/apijson"
	"github.com/sst/opencode-sdk-go/internal/apiquery"
	"github.com/sst/opencode-sdk-go/internal/param"
	"github.com/sst/opencode-sdk-go/internal/requestconfig"
	"github.com/sst/opencode-sdk-go/option"
)

// FileService contains methods and other services that help with interacting with
// the opencode API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Read a file
func (r *FileService) Read(ctx context.Context, query FileReadParams, opts ...option.RequestOption) (res *FileReadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get file status
func (r *FileService) Status(ctx context.Context, opts ...option.RequestOption) (res *[]File, err error) {
	opts = append(r.Options[:], opts...)
	path := "file/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type File struct {
	Added   int64      `json:"added,required"`
	Path    string     `json:"path,required"`
	Removed int64      `json:"removed,required"`
	Status  FileStatus `json:"status,required"`
	JSON    fileJSON   `json:"-"`
}

// fileJSON contains the JSON metadata for the struct [File]
type fileJSON struct {
	Added       apijson.Field
	Path        apijson.Field
	Removed     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileJSON) RawJSON() string {
	return r.raw
}

type FileStatus string

const (
	FileStatusAdded    FileStatus = "added"
	FileStatusDeleted  FileStatus = "deleted"
	FileStatusModified FileStatus = "modified"
)

func (r FileStatus) IsKnown() bool {
	switch r {
	case FileStatusAdded, FileStatusDeleted, FileStatusModified:
		return true
	}
	return false
}

type FileReadResponse struct {
	Content string               `json:"content,required"`
	Type    FileReadResponseType `json:"type,required"`
	JSON    fileReadResponseJSON `json:"-"`
}

// fileReadResponseJSON contains the JSON metadata for the struct
// [FileReadResponse]
type fileReadResponseJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FileReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileReadResponseJSON) RawJSON() string {
	return r.raw
}

type FileReadResponseType string

const (
	FileReadResponseTypeRaw   FileReadResponseType = "raw"
	FileReadResponseTypePatch FileReadResponseType = "patch"
)

func (r FileReadResponseType) IsKnown() bool {
	switch r {
	case FileReadResponseTypeRaw, FileReadResponseTypePatch:
		return true
	}
	return false
}

type FileReadParams struct {
	Path param.Field[string] `query:"path,required"`
}

// URLQuery serializes [FileReadParams]'s query parameters as `url.Values`.
func (r FileReadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
