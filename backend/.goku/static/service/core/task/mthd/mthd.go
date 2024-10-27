package mthd

import (
	"context"
	"fmt"
	"regexp"

	"github.com/robfig/cron/v3"

	"github.com/teejays/gokutil/errutil"
	"github.com/teejays/gokutil/filter"
	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
	app_typhelper "sampleapp/backend/.goku/generated/typhelper"
)

func StartTaskRunner(ctx context.Context, c app_client.Client) error {
	crn := cron.New()

	// Get all the tasks
	tasks, err := c.Core().Task().List(ctx, svccore_enttask_typ.ListRequest{
		Filter: svccore_enttask_typ.TaskFilter{
			Enabled: filter.NewBoolCondition(filter.EQUAL, true),
		},
	})
	if err != nil {
		return fmt.Errorf("Getting tasks: %w", err)
	}

	for _, task := range tasks.Items {
		log.Info(ctx, "Adding task to cron", "task", task)

		// expression, err := cronExpressionTypeToString(ctx, task.CronExpression)
		// if err != nil {
		// 	return fmt.Errorf("Converting cron expression to string: %w", err)
		// }
		// _, err = crn.AddFunc(expression, func() {
		// 	err := StartTask(ctx, c, task)
		// 	if err != nil {
		// 		log.Error(ctx, "Starting task", "task", task, "error", err)
		// 	}
		// })
		// if err != nil {
		// 	return fmt.Errorf("Adding function to cron: %w", err)
		// }
	}

	go func() {
		log.Warn(ctx, "Starting cron goroutine")
		crn.Start()
	}()

	return nil
}

// func cronExpressionTypeToString(ctx context.Context, exp svccore_enttask_typ.CronExpressionWithMeta) (string, error) {
// 	return fmt.Sprintf("%s %s %s %s %s %s", exp.Second, exp.Minute, exp.Hour, exp.DayOfMonth, exp.Month, exp.DayOfWeek), nil
// }

func StartTask(ctx context.Context, c app_client.Client, task svccore_enttask_typ.Task) error {
	log.Info(ctx, "Starting task", "task", task)

	reqType, err := app_typhelper.GetNewRequestTypeForMethod(ctx, task.Method)
	if err != nil {
		return fmt.Errorf("Getting request type for method: %w", err)
	}

	err = task.MethodRequestData.LoadTo(&reqType)
	if err != nil {
		return fmt.Errorf("Unmarshalling method request JSON: %w", err)
	}

	// Add a Task Run
	_, err = c.Core().TaskRun().Add(ctx, svccore_enttaskrun_typ.AddRequest{
		Object: svccore_enttaskrun_typ.TaskRunInput{
			TaskID:            task.ID,
			MethodRequestData: task.MethodRequestData,
			TriggeredBy:       svccore_enttaskrun_typ.TriggeredBy_Cron,
		},
	})
	if err != nil {
		return fmt.Errorf("Adding task run: %w", err)
	}

	// Run the TaskRun
	resp, err := c.Core().TaskRun().ActionRun(ctx, svccore_enttaskrun_typ.StandardEntityRequest{
		ID: task.ID,
	})
	if err != nil {
		return fmt.Errorf("Running task: %w", err)
	}

	log.Info(ctx, "Task completed", "task", task, "taskRun", resp.Object)

	return nil
}

func HookSavePre(ctx context.Context, c app_client.Client, in svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error) {

	// var err error

	// Validate the task

	// // Validate the cron expression
	// err = ValidateCronExpression(ctx, in)
	// if err != nil {
	// 	return in, err
	// }

	// If args are provided, validate them
	if !in.MethodRequestData.IsEmpty() {
		// Get the method request_type
		reqTyp, err := app_typhelper.GetNewRequestTypeForMethod(ctx, in.Method)
		if err != nil {
			return in, fmt.Errorf("Getting request type for method: %w", err)
		}
		// Make sure that the request JSON can be unmarshalled into the request type
		err = in.MethodRequestData.LoadTo(&reqTyp)
		if err != nil {
			return in, fmt.Errorf("Method Request JSON value seems invalid. Cannot unmarshalling it into the method request type: %w", err)
		}
	}

	return in, nil
}

func ActionRun(ctx context.Context, c app_client.Client, req svccore_enttask_typ.StandardEntityRequest) (svccore_enttask_typ.StandardEntityResponse, error) {

	return svccore_enttask_typ.StandardEntityResponse{}, nil
}

var (
	rgxSecondStr     = `^(\*|\/|\,|\-|([0-9]|[1-5][0-9]))$`
	rgxMinuteStr     = `^(\*|\/|\,|\-|([0-9]|[1-5][0-9]))$`
	rgxHourStr       = `^(\*|\/|\,|\-|([0-9]|1[0-9]|2[0-3]))$`
	rgxDayOfMonthStr = `^(\*|\/|\,|\-|\?|([0-9]|1[0-9]|2[0-9]|3[0-1]))$`
	rgxMonthStr      = `^(\*|\/|\,|\-|([0-9]|1[0-2]))$`
	rgxDayOfWeekStr  = `^(\*|\/|\,|\-|\?|([0-6]))$`

	rgxSecond     *regexp.Regexp
	rgxMinute     *regexp.Regexp
	rgxHour       *regexp.Regexp
	rgxDayOfMonth *regexp.Regexp
	rgxMonth      *regexp.Regexp
	rgxDayOfWeek  *regexp.Regexp
)

func init() {
	rgxSecond = regexp.MustCompile(rgxSecondStr)
	rgxMinute = regexp.MustCompile(rgxMinuteStr)
	rgxHour = regexp.MustCompile(rgxHourStr)
	rgxDayOfMonth = regexp.MustCompile(rgxDayOfMonthStr)
	rgxMonth = regexp.MustCompile(rgxMonthStr)
	rgxDayOfWeek = regexp.MustCompile(rgxDayOfWeekStr)
}

func ValidateCronExpression(ctx context.Context, in svccore_typ.CronExpression) error {
	var errs = errutil.NewMultiErr()

	/*
		type CronExpressionInput struct {
			Second     string `json:"second" validate:"required"`
			Minute     string `json:"minute" validate:"required"`
			Hour       string `json:"hour" validate:"required"`
			DayOfMonth string `json:"dayOfMonth" validate:"required"`
			Month      string `json:"month" validate:"required"`
			DayOfWeek  string `json:"dayOfWeek" validate:"required"`
		}
	*/

	// Validate the cron expression

	// Second: Allowed characters: 0-59 * / , -

	if !rgxSecond.MatchString(in.Second) {
		errs.AddNew("CronExpression.Second failed: %s")
	}
	if !rgxMinute.MatchString(in.Minute) {
		errs.AddNew("CronExpression.Minute failed: %s")
	}
	if !rgxHour.MatchString(in.Hour) {
		errs.AddNew("CronExpression.Hour failed: %s")
	}
	if !rgxDayOfMonth.MatchString(in.DayOfMonth) {
		errs.AddNew("CronExpression.DayOfMonth failed: %s")
	}
	if !rgxMonth.MatchString(in.Month) {
		errs.AddNew("CronExpression.Month failed: %s")
	}
	if !rgxDayOfWeek.MatchString(in.DayOfWeek) {
		errs.AddNew("CronExpression.DayOfWeek failed: %s")
	}

	return errs.ErrOrNil()

}
