package provider

import (
	"context"
	"fmt"
	"github.com/wangjiaxi90/terraform-provider-hashdata/internal/provider/cloudmgr"
	"log"
	"math/rand"
	_nethttp "net/http"
	"strconv"
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

func waitJobComplete(ctx context.Context, clt *cloudmgr.CoreJobServiceApiService, id string) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}
	for i := 1; i <= 500; i++ {
		var resp cloudmgr.CommonDescribeJobResponse
		var r *_nethttp.Response
		var err error
		resp, r, err = clt.DescribeJob(ctx, id).Execute()
		if err != nil {
			return err
		}
		if r.StatusCode != 200 {
			return fmt.Errorf("Bad response code with %s ", r.Status)
		}
		//TODO 判断是否被删除
		status := strings.ToLower(resp.GetStatus())
		if status == JOB_WAIT_PENDING || status == JOB_WAIT_RUNNING {
			time.Sleep(10 * time.Second)
			log.Println("Still waiting current state is " + status + ". [" + strconv.Itoa(i*10) + "s elapsed]")
			continue
		} else if status == JOB_FAILED_ABANDONED || status == JOB_FAILED_FAILURE {
			return fmt.Errorf("Error instance create failed, request id %s, status %s ", id, status)
		} else if status == JOB_SUCCESS {
			return nil
		} else {
			return fmt.Errorf("Error unknow status code %s ", status)
		}
	}
	return fmt.Errorf("Timeout while waiting for state become 'success'. ")
}

func startService(ctx context.Context, id string, apiClient *cloudmgr.APIClient) error {
	respStartService, rStartService, errStartService := apiClient.CoreServiceApi.StartService(ctx, id).Execute()
	if errStartService != nil {
		return fmt.Errorf("Error when calling `CoreServiceApi.StartService`: %v\v", errStartService)
	}
	if rStartService.StatusCode != 200 {
		return fmt.Errorf("Error when calling `CoreServiceApi.StartService`: %s\n", rStartService.Status)
	}
	if eWait := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStartService.GetId()); eWait != nil {
		return fmt.Errorf("Error when waiting start service %s\n", eWait.Error())
	}
	return nil
}

func stopService(ctx context.Context, id string, apiClient *cloudmgr.APIClient) error {
	respStopService, rStopService, errStopService := apiClient.CoreServiceApi.StopService(ctx, id).Execute()
	if errStopService != nil {
		return fmt.Errorf("Error when calling `CoreServiceApi.StopService`: %v\v", errStopService)
	}
	if rStopService.StatusCode != 200 {
		return fmt.Errorf("Error when calling `CoreServiceApi.StopService`: %s\n", rStopService.Status)
	}
	if eWait := waitJobComplete(ctx, apiClient.CoreJobServiceApi, respStopService.GetId()); eWait != nil {
		return fmt.Errorf("Error when waiting stop service %s\n", eWait.Error())
	}
	return nil
}
