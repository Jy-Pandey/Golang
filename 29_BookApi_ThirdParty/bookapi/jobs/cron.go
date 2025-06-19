package jobs

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type CronJob struct {
	DB *gorm.DB
}

func NewCronJob(db *gorm.DB) *CronJob {
	return &CronJob{DB: db}
}

func (cj *CronJob) StartCronJob() {
	// Creates a new cron scheduler instance.
	c := cron.New(cron.WithSeconds())

	// add schedule
	// delet books after 1 hour
	// return id and err
	_, err := c.AddFunc("0 0 * * * *", func() {
		fmt.Println("Running cleanup Job at : ", time.Now())
		err := cj.DB.Exec("TRUNCATE TABLE books").Error
		if err != nil {
			fmt.Println("Failed to truncate books:", err)
		} else {
			fmt.Println("All books truncated.")
		}
	})

	if err != nil {
		fmt.Println("Failed to schedule cron job:", err)
	}

	c.Start() // start the schedule
	fmt.Println("Cron job started")

}
