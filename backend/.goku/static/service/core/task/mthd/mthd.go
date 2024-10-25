package mthd

import (
	"context"
	"fmt"
	"regexp"

	"github.com/teejays/gokutil/errutil"

	app_client "ongoku/backend/.goku/generated/client"
	svccore_enttask_typ "ongoku/backend/.goku/generated/service/core/entity/task/typ"
)

func HookSavePre(ctx context.Context, c app_client.Client, in svccore_enttask_typ.Task) (svccore_enttask_typ.Task, error) {

	var err error

	// Validate the task

	/*

	   type Task struct {
	   	ID             scalars.ID             `json:"id"`
	   	Name           string                 `json:"name" validate:"required"`
	   	Description    string                 `json:"description" validate:"required"`
	   	MethodService  string                 `json:"methodService" validate:"required"`
	   	MethodName     string                 `json:"methodName" validate:"required"`
	   	MethodArgs     scalars.JSON           `json:"methodArgs" validate:"required"`
	   	CronExpression CronExpressionWithMeta `json:"cronExpression" validate:"required"`
	   	Enabled        bool                   `json:"enabled"`
	   	LastRunAt      scalars.Timestamp      `json:"lastRunAt" validate:"required"`
	   	NextRunAt      scalars.Timestamp      `json:"nextRunAt" validate:"required"`
	   	CreatedAt      scalars.Timestamp      `json:"createdAt"`
	   	UpdatedAt      scalars.Timestamp      `json:"updatedAt"`
	   	DeletedAt      *scalars.Timestamp     `json:"deletedAt"`
	   }
	*/

	// Validate the cron expression
	err = ValidateCronExpression(ctx, in)
	if err != nil {
		return in, err
	}

	// If args are provided, validate them
	if !in.MethodArgs.IsEmpty() {
		args := map[string]interface{}{}
		err = in.MethodArgs.JSONUnmarshal(&args)
		if err != nil {
			return in, fmt.Errorf("UMethodArgs seem invalid. Cannot unmarshal them into a map: %w", err)
		}
	}

	return in, nil
}

var (
	rgxSecondStr     = `^(\*|\/|\,|\-|([0-9]|[1-5][0-9]))$`
	rgxMinuteStr     = `^(\*|\/|\,|\-|([0-9]|[1-5][0-9]))$`
	rgxHourStr       = `^(\*|\/|\,|\-|([0-9]|1[0-9]|2[0-3]))$`
	rgxDayOfMonthStr = `^(\*|\/|\,|\-|\?|([0-9]|1[0-9]|2[0-9]|3[0-1]))$`
	rgxMonthStr      = `^(\*|\/|\,|\-|([0-9]|1[0-2])))$`
	rgxDayOfWeekStr  = `^(\*|\/|\,|\-|\?|([0-6])|?)$`

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

func ValidateCronExpression(ctx context.Context, in svccore_enttask_typ.Task) error {
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

	if !rgxSecond.MatchString(in.CronExpression.Second) {
		errs.AddNew("CronExpression.Second failed: %s")
	}
	if !rgxMinute.MatchString(in.CronExpression.Minute) {
		errs.AddNew("CronExpression.Minute failed: %s")
	}
	if !rgxHour.MatchString(in.CronExpression.Hour) {
		errs.AddNew("CronExpression.Hour failed: %s")
	}
	if !rgxDayOfMonth.MatchString(in.CronExpression.DayOfMonth) {
		errs.AddNew("CronExpression.DayOfMonth failed: %s")
	}
	if !rgxMonth.MatchString(in.CronExpression.Month) {
		errs.AddNew("CronExpression.Month failed: %s")
	}
	if !rgxDayOfWeek.MatchString(in.CronExpression.DayOfWeek) {
		errs.AddNew("CronExpression.DayOfWeek failed: %s")
	}

	return errs.ErrOrNil()

}
