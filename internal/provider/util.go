package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/wangjiaxi90/terraform-provider-hashdata/internal/provider/cloudmgr"
	"math/rand"
	_nethttp "net/http"
	"strings"
	"time"
)

func Int32(v int) *int32 {
	t := int32(v)
	return &t
}
func String(v string) *string {
	return &v
}
func Bool(v bool) *bool {
	return &v
}

func simpleRetry(fn func() error) error {
	return retry(MAX_RETRY_TIMES, 10*time.Second, fn)
}

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}

func isServerBusy(err error) error {
	if err != nil {
		return stop{err}
	}
	return nil
}

func InstanceTransitionStateRefresh(ctx context.Context, clt *cloudmgr.CoreJobServiceApiService, id string) (interface{}, error) {
	if id == "" {
		return nil, nil
	}
	refreshFunc := func() (interface{}, string, error) {
		var resp cloudmgr.CommonDescribeJobResponse
		var r *_nethttp.Response
		var err error
		resp, r, err = clt.DescribeJob(ctx, id).Execute()
		if err != nil {
			return nil, "", err
		}
		if r.StatusCode != 200 {
			return nil, "", fmt.Errorf("Bad response code with %s ", r.Status)
		}
		//TODO 判断是否被删除
		status := strings.ToLower(resp.GetStatus())
		if status == JOB_WAIT_PENDING || status == JOB_WAIT_RUNNING {
			return nil, status, nil
		} else if status == JOB_FAILED_ABANDONED || status == JOB_FAILED_FAILURE {
			return nil, status, fmt.Errorf("Error instance create failed, request id %s, status %s ", id, status)
		} else if status == JOB_SUCCESS {
			return nil, "", nil
		} else {
			return nil, status, fmt.Errorf("Error unknow status code %s ", status)
		}

	}
	stateConf := &resource.StateChangeConf{
		Pending:    []string{JOB_WAIT_PENDING, JOB_WAIT_RUNNING},
		Target:     []string{""},
		Refresh:    refreshFunc,
		Timeout:    waitJobTimeOutDefault * time.Second,
		Delay:      waitJobIntervalDefault * time.Second,
		MinTimeout: waitJobIntervalDefault * time.Second,
		NotFoundChecks: 600,
	}
	return stateConf.WaitForStateContext(ctx)
}
