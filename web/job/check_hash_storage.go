package job

import (
	"github.com/h1151449095/3x-ui-mogaikai/v3/web/service"
)

// CheckHashStorageJob periodically cleans up expired hash entries from the Telegram bot's hash storage.
type CheckHashStorageJob struct {
	tgbotService service.Tgbot
}

// NewCheckHashStorageJob creates a new hash storage cleanup job instance.
func NewCheckHashStorageJob() *CheckHashStorageJob {
	return new(CheckHashStorageJob)
}

// Run removes expired hash entries from the Telegram bot's hash storage.
func (j *CheckHashStorageJob) Run() {
	storage := j.tgbotService.GetHashStorage()
	if storage == nil {
		return
	}
	storage.RemoveExpiredHashes()
}
